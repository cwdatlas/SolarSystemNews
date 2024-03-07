package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var gameRunning = true
var choosingArticle = true
var cyan = "\033[36m"
var red = "\033[31m"
var reset = "\033[0m"

func main() {
	fmt.Println(cyan, "You are one of the many tens of thousands of lone Asteroid miners out in a distant section\n",
		"of the belt. You grew up on the Martian colonies, but you didnt like the couped up life of the domned habitats.\n"+
			"You wanted to be an explorer like the people who founded your home colony, so you decided to use your savings\n"+
			"and move out into the belt where you can explore and profit heavily from doing what you love...\n",
		"At least that is what you thought, before you were again coupled up in a little metal can with a bunch of\n"+
			"'Space Bumpkins' on the coms and local news channel. Your ship is an old throwaway leftover with the internet ingress\n"+
			"point ripped out by its previous owners to make an extra buck. You are now stuck bored between mining missions.\n",
		"Wonder what the 'Space Bumpkins' are up to, so you log onto your dialup-like internet and search buffered news articles...\n",
		"------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(reset, "Do you want to check local or system wide space news?")
	for gameRunning {
		fmt.Println("Article Repositories: system or local: ")
		var text = reader()
		if text == "quit" {
			fmt.Println("Goodbye!")
			gameRunning = false
			break
		} else if text == "local" {
			fmt.Println("--List of buffered news articles from your local area--")
			fmt.Println("		> alias: title")
			for alias := range getLocalNews() {
				fmt.Println(cyan, "		>", alias, ":", getLocalNews()[alias][0])
			}
			fmt.Println(reset)
			choosingArticle = true
			for choosingArticle {
				fmt.Println("Type Article alias you wish you read!")
				var text = reader()
				body, ok := getLocalNews()[text]
				if text == "quit" {
					fmt.Println("Exiting Article Selection")
					choosingArticle = false
				} else if ok {
					fmt.Println("--Article '", body[0], "' reading--")
					fmt.Println(cyan, body[1])
					fmt.Print(reset)
				} else {
					fmt.Println(red, text, "is not a valid alias or command, to quit article selection use 'quit'")
					fmt.Print(reset)
				}

			}
		} else if text == "system" {
			fmt.Println("List of buffered news articles from across the settled systems")
			fmt.Println("		> alias: Title")
			for alias := range getSystemNews() {
				fmt.Println(cyan, "		>", alias, ":", getSystemNews()[alias][0])
			}
			fmt.Println(reset)
			choosingArticle = true
			for choosingArticle {
				fmt.Println("Type Article alias you wish you read!")
				var text = reader()
				body, ok := getSystemNews()[text]
				if text == "quit" {
					fmt.Println("Exiting Article Selection")
					choosingArticle = false
				} else if ok {
					fmt.Println("--Article '", body[0], "' reading--")
					fmt.Println(cyan, body[1])
					fmt.Print(reset)
				} else {
					fmt.Println(red, text, "is not a valid article or command, to quit article selection use 'quit'")
					fmt.Print(reset)
				}

			}
		} else {
			fmt.Println(red, text, "was not found as an article repository. please type a valid repository or 'quit'")
			fmt.Print(reset)
		}
	}

}

func reader() string {
	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	key, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occurred:", err)
		panic(err)
	}
	return strings.ToLower(strings.TrimSpace(key))
}
