package main

import (
	"fmt"
	"time"
)

func speak(person, text string, qntt int) {
	for i := 0; i < qntt; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interation %d)\n", person, text, i+1)
	}
}

func main() {
	//speak("John", "O tempo andou mexendo com a gente, sim?", 3)
	//speak("Belchior", "OH NO, NO NO NO", 1)

	//go speak("John", "Got it", 500)
	//go speak("Belchior", "Lets move on", 500)

	go speak("John", "That's it", 10)
	speak("Belchias", "Alright", 5)
}
