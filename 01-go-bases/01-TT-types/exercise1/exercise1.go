package main

import "fmt"

func showCharCounter(word string) {
	fmt.Println("Length of word:", len(word))
}
func showEachChar(word string) {
	for _, r := range word {
		fmt.Println(string(r))
	}

}
func main() {
	showCharCounter("cat")
	showEachChar("cat")
	showCharCounter("computer")
	showEachChar("computer")

}
