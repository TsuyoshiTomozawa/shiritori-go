package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("GAME START\n")
	word := "り"
	start(word)

}

func start(word string) {
	fmt.Printf("最初の文字: %s\n", word)
	text := input()
	judge(text, word)
}

func input() string {

	fmt.Printf("文字を入力してください\n")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}

func judge(text, word string) {

	textRune := []rune(text)
	size := len(textRune)
	firstWord := string(textRune[0])

	if firstWord != word {
		fmt.Printf("先頭の文字が違います\n")
		fmt.Printf("もう一度入力してね\n")
		start(word)
	}

	lastWord := string([]rune(text)[size-1])
	fmt.Printf("最後の文字: %s\n", lastWord)
	if lastWord == "ん" {
		end()
	}

	start(lastWord)
}

func end() {
	fmt.Printf("GAME OVER\n")
	os.Exit(0)
}
