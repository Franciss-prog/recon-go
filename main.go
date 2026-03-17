package main

import (
	"fmt"
	"os"
)

func main() {
	// variables needed
	var link string
	var infoAction int

	// clear the cli
	os.Stdout.Write([]byte("\033[H\033[2J"))
	fmt.Print(`
 _ __ ___  ___ ___  _ __         __ _  ___  
| '__/ _ \/ __/ _ \| '_ \ _____ / _` + "`" + ` |/ _ \ 
| | |  __/ (_| (_) | | | |_____| (_| | (_) |
|_|  \___|\___\___/|_| |_|      \__, |\___/ 
                                |___/   
` + "\n")
	fmt.Println("A tool that automatically scans a domain and finds its infrastructure.")

	// ask the user for link
	fmt.Println("Enter the link of the domain to scan: ")
	fmt.Scanln(&link)

	// SAMPLE VALIDATION
	if len(link) <= 0 {
		fmt.Println("Please enter a valid link.")
		return
	}

	// clear the cli
	os.Stdout.Write([]byte("\033[H\033[2J"))

	// start the scanning
	// SAMPLE SCANNING
	fmt.Println("Scanning: " + link + "\n")

	// after scanning show the info of the link inputted
	fmt.Println("INFO: ")

	//  ask the user what to do in the info
	fmt.Println("Are you done reading the Info of " + link + "?")
	fmt.Println("What do you want to do to the Info of " + link + "?")
	fmt.Println("1. Open in browser")
	fmt.Println("2. Save to file")
	fmt.Println("3. Exit")
	fmt.Scanln(&infoAction)

}
