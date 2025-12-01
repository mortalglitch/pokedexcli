package main
import (
	"fmt"
	"bufio"
	"os"

)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		if scanner.Scan() != false{
			cleanText := cleanInput(scanner.Text())		
			fmt.Println("Your command was:", cleanText[0])
		}
	}
}

