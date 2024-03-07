package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Controls if the game is running
var gameRunning = true

// Controls the article selection
var choosingArticle = true

// color codes
var cyan = "\033[36m"
var red = "\033[31m"
var reset = "\033[0m"

/*
Author: Aidan Scott
Main() launches the application
*/
func main() {
	fmt.Println(cyan,
		"You are one of the many tens of thousands of lone Asteroid miners out in a distant section\n",
		"of the belt. You grew up on the Martian colonies, but you didnt like the couped up life of the domned habitats.\n"+
			"You wanted to be an explorer like the people who founded your home colony, so you decided to use your savings\n"+
			"and move out into the belt where you can explore and profit heavily from doing what you love...\n",
		"At least that is what you thought, before you were again coupled up in a little metal can with a bunch of\n"+
			"'Space Bumpkins' on the coms and local news channel. Your ship is an old throwaway leftover with the internet ingress\n"+
			"point ripped out by its previous owners to make an extra buck. You are now stuck bored between mining missions.\n",
		"Wonder what the 'Space Bumpkins' are up to, so you log onto your dialup-like internet and search buffered news articles...\n",
		"------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(reset, "Do you want to check local or system wide space news?")
	// start main for loop that controls the repo selection
	for gameRunning {
		fmt.Println("Article Repositories: system or local: ")
		// uses the reader() function to scan user input
		var text = reader()
		// if user wants to quit, stop the game from running
		if text == "quit" {
			fmt.Println("Goodbye!")
			gameRunning = false // stops the game
			break
			// If the user input local, then list out articles and aliases for stories
		} else if text == "local" {
			fmt.Println("--List of buffered news articles from your local area--")
			fmt.Println("		> alias: title")
			for alias := range getLocalNews() {
				// list articles and aliases
				fmt.Println(cyan, "		>", alias, ":", getLocalNews()[alias][0])
			}
			fmt.Println(reset) // resets color to white
			choosingArticle = true
			// loop allowing user to check multiple articles
			for choosingArticle {
				fmt.Println("Type Article alias you wish you read!")
				var text = reader()
				// checks to see if value is in the map
				body, ok := getLocalNews()[text]
				// allows for quiting
				if text == "quit" {
					fmt.Println("Exiting Article Selection")
					choosingArticle = false
					// if article is in map
				} else if ok {
					// print article
					fmt.Println("--Article '", body[0], "' reading--")
					fmt.Println(cyan, body[1])
					fmt.Print(reset)
				} else {
					// tell the user there is an issue in red
					fmt.Println(red, text, "is not a valid alias or command, to quit article selection use 'quit'")
					fmt.Print(reset)
				}

			}
		} else if text == "system" {
			fmt.Println("--List of buffered news articles from the settled systems--")
			fmt.Println("		> alias: title")
			for alias := range getSystemNews() {
				// list articles and aliases
				fmt.Println(cyan, "		>", alias, ":", getSystemNews()[alias][0])
			}
			fmt.Println(reset) // resets color to white
			choosingArticle = true
			// loop allowing user to check multiple articles
			for choosingArticle {
				fmt.Println("Type Article alias you wish you read!")
				var text = reader()
				// checks to see if value is in the map
				body, ok := getSystemNews()[text]
				// allows for quiting
				if text == "quit" {
					fmt.Println("Exiting Article Selection")
					choosingArticle = false
					// if article is in map
				} else if ok {
					// print article
					fmt.Println("--Article '", body[0], "' reading--")
					fmt.Println(cyan, body[1])
					fmt.Print(reset)
				} else {
					// tell the user there is an issue in red
					fmt.Println(red, text, "is not a valid alias or command, to quit article selection use 'quit'")
					fmt.Print(reset)
				}

			}
		} else {
			// if repo is not found, print message in red
			fmt.Println(red, text, "was not found as an article repository. please type a valid repository or 'quit'")
			fmt.Print(reset)
		}
	}

}

func reader() string {
	// Initiate user input reader
	reader := bufio.NewReader(os.Stdin)
	// make sure user understand they can type by inputting ->
	fmt.Print("-> ")
	key, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occurred:", err)
		panic(err)
	}
	// small formatting, making all text to be lowercase and trimming the end of the string to get rid of \r\n and spaces
	return strings.ToLower(strings.TrimSpace(key))
}
