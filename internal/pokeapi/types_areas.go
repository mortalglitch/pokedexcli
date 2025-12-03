package pokeapi

// RespShallowLocations -
type RespShallowAreas struct {
	ID               int                `json:"id"`
	Name             string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounter struct {
	Pokemon       NamedApiResource    `json:"pokemon"`
}
