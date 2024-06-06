package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

// ToLower converts a string to lowercase
func ToLower(s string) string {
    str := []rune(s)
    for i := 0; i < len(s); i++ {
        if s[i] >= 65 && s[i] <= 90 {
            str[i] = str[i] + 32
        }
    }
    return string(str)
}

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
    str := []rune(s)
    for i := 0; i < len(s); i++ {
        if s[i] >= 97 && s[i] <= 122 {
            str[i] = str[i] - 32
        }
    }
    return string(str)
}

// processWord applies the transformation function to a given number of words before the specified index
func processWord(words []string, index int, count int, transform func(string) string) {
    for i := 0; i < count; i++ {
        if index-i >= 0 {
            words[index-i] = transform(words[index-i])
        }
    }
}
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
    writer := bufio.NewWriter(file1)
    defer writer.Flush()

    var myword string

    for scanner.Scan() {
        line := scanner.Text()  // Read each line separately
        words := strings.Fields(line)

        var tab [2]string
		
		for i, word := range words {
            tab[0], tab[1] = tab[1], word

            switch {
            case strings.HasPrefix(word, "(up"):
                count := 1
                if strings.Contains(word, ",") {
                    fmt.Sscanf(word, "(up,%d)", &count)
                }
                processWord(words, i-1, count, ToUpper)
            case strings.HasPrefix(word, "(low"):
                count := 1
                if strings.Contains(word, ",") {
                    fmt.Sscanf(word, "(low,%d)", &count)
                }
                processWord(words, i-1, count, ToLower)
            case strings.HasPrefix(word, "(cap"):
                count := 1
                if strings.Contains(word, ",") {
                    fmt.Sscanf(word, "(cap,%d)", &count)
                }
                processWord(words, i-1, count, func(s string) string {
                    return strings.Title(ToLower(s))
                })
            case word == "(hex)":
                d, err := strconv.ParseInt(tab[0], 16, 64)
                if err != nil {
                    panic(err)
                }
                s := strconv.FormatInt(d, 10)
                words[i-1] = s
            case word == "(bin)":
                d, err := strconv.ParseInt(tab[0], 2, 64)
                if err != nil {
                    panic(err)
                }
                s := strconv.FormatInt(d, 10)
                words[i-1] = s
            default:
                if !strings.HasPrefix(tab[0], "(") && tab[0] != "" {
                    myword += tab[0] + " "
                }
            }
        }
		// Join words for the current line and adjust punctuation on this line only
        lineWithAdjustedPunctuation := adjustPunctuation(strings.Join(words, " "))  // Adjust punctuation for each line
        myword += lineWithAdjustedPunctuation + "\n"  // Preserve newline characters
    }

    // Adjust indefinite articles
    myword = adjustIndefiniteArticles(myword)

    // Adjust single quotation marks
    myword = adjustSingleQuotes(myword)

    _, errs = writer.WriteString(myword)
    if errs != nil {
        fmt.Println("Failed to write to file:", errs)
        return
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Println("File processed successfully!")
}
// adjustPunctuation ensures punctuation marks are formatted correctly without affecting newlines
func adjustPunctuation(text string) string {
    punctuations := []string{".", ",", "!", "?", ":", ";"}
    for _, p := range punctuations {
        re := regexp.MustCompile(`\s*` + regexp.QuoteMeta(p) + `\s*`)
        text = re.ReplaceAllString(text, p+" ")
    }

    // Handle groups of punctuation like ...
    text = strings.ReplaceAll(text, " . . . ", "...")
    text = strings.ReplaceAll(text, " ! ? ", "!?")

    return text
}
// adjustIndefiniteArticles ensures correct usage of "a" and "an"
func adjustIndefiniteArticles(text string) string {
    words := strings.Fields(text)
    for i := 0; i < len(words)-1; i++ {
        if words[i] == "a" || words[i] == "A" {
            nextWord := words[i+1]
            if strings.ContainsAny(string(nextWord[0]), "aeiouAEIOU") || strings.ToLower(string(nextWord[0])) == "h" {
                words[i] = "an"
            }
        }
    }
    return strings.Join(words, " ")
}


// adjustSingleQuotes formats single quotes around words correctly
func adjustSingleQuotes(text string) string {
    // This regular expression will match single quotes around one or more words, allowing for spaces around them
    re := regexp.MustCompile(`'\s*([^']+?)\s*'`)
    text = re.ReplaceAllStringFunc(text, func(match string) string {
        // Remove spaces after the opening quote and before the closing quote
        return "'" + strings.TrimSpace(match[1:len(match)-1]) + "'"
    })
    return text
}
