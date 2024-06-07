package main

import (
	"bufio"
	"fmt"
	"goreloaded/src/libs"
	"os"
	"strings"
)

func main() {
	// Check for the correct number of arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <inputfile> <outputfile>")
		return
	}

	// Open the file for reading
	file0, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file0.Close()

	// Open the file for writing
	file1, errs := os.Create(os.Args[2])
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		return
	}
	defer file1.Close()

	scanner := bufio.NewScanner(file0)

	var myword string

	for scanner.Scan() {
		line := scanner.Text()
		// Removing space from (low, n), (up, n), (cap, n)
		line = libs.FAdjustPunctuation(line)
		words := strings.Fields(line)

		var tab [2]string
		t := 0
		for k, word := range words {
			tab[0], tab[1] = tab[1], word

			switch {
			case strings.HasPrefix(word, "(up"):
				count := 1
				if strings.Contains(word, ",") {
					fmt.Sscanf(word, "(up,%d)", &count)

					w := strings.Fields(myword)
					myword = strings.Join(w[:k-count-t], " ") + " "

					for i := count; i > 1; i-- {
						myword += libs.ToUpper(words[k-i]) + " "
					}
				}
				myword += libs.ToUpper(tab[0]) + " "
				t++
			case strings.HasPrefix(word, "(low"):
				count := 1
				if strings.Contains(word, ",") {
					fmt.Sscanf(word, "(low,%d)", &count)

					w := strings.Fields(myword)
					myword = strings.Join(w[:k-count-t], " ") + " "

					for i := count; i > 1; i-- {
						myword += libs.ToLower(words[k-i]) + " "
					}
				}
				myword += libs.ToLower(tab[0]) + " "
				t++
			case strings.HasPrefix(word, "(cap"):
				count := 1
				if strings.Contains(word, ",") {
					fmt.Sscanf(word, "(cap,%d)", &count)

					w := strings.Fields(myword)
					myword = strings.Join(w[:k-count-t], " ") + " "

					for i := count; i > 1; i-- {
						myword += strings.Title(libs.ToLower(words[k-i])) + " "
					}
				}
				myword += strings.Title(libs.ToLower(tab[0])) + " "
				t++
			case word == "(hex)":
				myword += libs.Hex(tab[0]) + " "
				t++
			case word == "(bin)":
				myword += libs.Bin(tab[0]) + " "
				t++
			default:
				if !strings.HasPrefix(tab[0], "(") && tab[0] != "" {
					myword += tab[0] + " "
				}
			}
		}

		// Add the last word
		myword += tab[1]

	}

	// Adjust punctuation
	myword = libs.AdjustPunctuation(myword)

	// Adjust indefinite articles
	myword = libs.AdjustIndefiniteArticles(myword)

	// Adjust single quotation marks
	myword = libs.AdjustSingleQuotes(myword)

	// Write into the new file
	_, errs = file1.WriteString(strings.TrimSpace(myword))
	if errs != nil {
		fmt.Println("Failed to write to file:", errs)
		return
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print that the process succeeded
	fmt.Println("File processed successfully!")
}
