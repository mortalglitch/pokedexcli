package pokecache

import (
	"testing"
	"time"
)

// TestAddAndGet tests the basic Add and Get functionality.
func TestAddAndGet(t *testing.T) {
	const testInterval = time.Second * 5 // Interval doesn't matter for this test
	cache := NewCache(testInterval)
	key := "test_key"
	val := []byte("test_value")

	cache.Add(key, val)

	retrievedVal, found := cache.Get(key)
	if !found {
		t.Fatalf("Expected key %s to be found, but it was not", key)
	}

	if string(retrievedVal) != string(val) {
		t.Fatalf("Expected value %s, got %s", string(val), string(retrievedVal))
	}

	// Test a missing key
	_, found = cache.Get("missing_key")
	if found {
		t.Fatalf("Expected missing_key not to be found, but it was")
	}
}

// TestReap tests the reapLoop functionality to ensure old entries are removed.
func TestReap(t *testing.T) {
	// Set a very short interval so the test runs quickly
	const baseInterval = time.Millisecond * 50
	cache := NewCache(baseInterval)
	key := "old_entry"
	val := []byte("old_data")

	// 1. Add an entry and artificially set its creation time to be expired.
	cache.Add(key, val)

	// Access the underlying map data to change the createdAt field.
	// We need to lock since we are manually manipulating the map outside of a method's lock.
	cache.mu.Lock()
	// Set the creation time far in the past, making it older than the baseInterval
	expiredTime := time.Now().Add(-baseInterval * 2)
	cache.data[key] = cacheEntry{
		createdAt: expiredTime,
		val:       val,
	}
	cache.mu.Unlock()

	// 2. Wait for the reapLoop to run at least once (wait slightly longer than the interval)
	time.Sleep(baseInterval + time.Millisecond*20) 

	// 3. Check if the entry has been removed (reaped)
	_, found := cache.Get(key)
	if found {
		t.Fatalf("Expected key %s to be reaped and not found, but it was present", key)
	}

	// Check a valid entry (not expired)
	freshKey := "fresh_entry"
	cache.Add(freshKey, []byte("fresh_data"))
	time.Sleep(time.Millisecond * 10) // Wait a short time
	_, found = cache.Get(freshKey)
	if !found {
		t.Fatalf("Expected key %s to be present, but it was not", freshKey)
	}
}
