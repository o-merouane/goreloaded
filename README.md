Text Completion/Editing/Auto-Correction Tool

Introduction
This project is a text completion/editing/auto-correction tool written in Go. The tool reads a text file, applies various modifications to the text based on specific patterns, and writes the modified text to a new file. This project will be reviewed by peer auditors, and you will also act as an auditor for other students' projects.

Features
The tool supports the following modifications:
Hexadecimal to Decimal Conversion: Replace a hexadecimal number (followed by "(hex)") with its decimal equivalent.
Example: "1E (hex) files were added" -> "30 files were added"

Binary to Decimal Conversion: Replace a binary number (followed by "(bin)") with its decimal equivalent.
Example: "It has been 10 (bin) years" -> "It has been 2 years"

Uppercase Conversion: Convert the preceding word to uppercase (indicated by "(up)").
Example: "Ready, set, go (up) !" -> "Ready, set, GO !"

Lowercase Conversion: Convert the preceding word to lowercase (indicated by "(low)").
Example: "I should stop SHOUTING (low)" -> "I should stop shouting"

Capitalized Conversion: Capitalize the preceding word (indicated by "(cap)").
Example: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge"

Multiple Word Case Conversion: Convert multiple preceding words to uppercase, lowercase, or capitalized based on a specified number.
Example: "This is so exciting (up, 2)" -> "This is SO EXCITING"

Punctuation Formatting: Ensure punctuation marks (, . ! ? : ;) are placed correctly without space before them and with space after them.

Example: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!"
Quotation Formatting: Ensure single quotation marks are placed correctly around words.

Example: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'"Article Correction: Replace "a" with "an" if the following word starts with a vowel or 'h'.

Example: "There it was. A amazing rock!" -> "There it was. An amazing rock!"


Project Structure
main.go: The main entry point of the application.
src/libs: Contains functions for each type of text modification.

Usage
To run the tool, use the following command:
go run main.go <input_file> <output_file>
Where <input_file> is the name of the file containing the text to be modified, and <output_file> is the name of the file where the modified text will be written.
