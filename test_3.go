package main

import (
	"fmt"
	"slices"
	"strings"
)

func main(){
	brackets := []string{"{ [ ( ) ] }", "{ [ ( ] ) }", "{ ( ( [ ] ) [ ] ) [ ] }"}

	for i, bracket := range brackets{
		fmt.Println("CASE ", i + 1)
		fmt.Println("------------------------------")
		fmt.Println("Input: ", bracket)
		result := solution(bracket)
		fmt.Printf("Result: %s\n\n", result)
	}
}

func solution(s string) string {
	splitted_brackets := strings.Split(s, "")
	bracket_pairs := map[string]string{
		"(" : ")",
		"{" : "}",
		"[" : "]",
	}

	// delete white space
	for i, v := range splitted_brackets{
		if v == " "{
			splitted_brackets = slices.Delete(splitted_brackets, i, i + 1)
		}
	}
	for len(splitted_brackets) > 0{
		for i := 0; i < len(splitted_brackets); i++{
			if bracket_pairs[splitted_brackets[i]] == ""{
				return "NO"
			}

			if splitted_brackets[i + 1] == bracket_pairs[splitted_brackets[i]]{
				splitted_brackets = slices.Delete(splitted_brackets, i, i + 2)
				break
			}

			if (i == len(splitted_brackets) - 1){
				return "NO"
			}
		}
	}

	return "YES"
}