# Text Modification Tool in Go

## Introduction

This project implements a text completion, editing, and auto-correction tool in Go. The tool processes a text file input, applies specified modifications, and outputs the modified text to another file. It incorporates various text manipulation functions to handle conversions, case adjustments, punctuation formatting, and more.

## Features

- **Decimal Conversion:**
  - Converts hexadecimal numbers preceded by "(hex)" to their decimal equivalents.
  - Converts binary numbers preceded by "(bin)" to their decimal equivalents.

- **Case Conversion:**
  - Converts words preceded by "(up)" to uppercase.
  - Converts words preceded by "(low)" to lowercase.
  - Converts words preceded by "(cap)" to capitalized form.
  - Supports optional conversion of a specified number of words following directives like "(up, 2)".

- **Punctuation Formatting:**
  - Ensures punctuation marks (. , ! ? : ;) are correctly placed with respect to adjacent words.
  - Handles exceptions like "..." and "!?" for proper spacing.

- **Quote Handling:**
  - Ensures single quotes (' ') are correctly placed around quoted phrases.

- **Indefinite Article Adjustment:**
  - Changes "a" to "an" before words starting with a vowel or "h".

## Implementation

- **File Handling:** Utilizes Go's file system API to read from and write to files.
- **String Manipulation:** Implements various string manipulation techniques for text modifications.
- **Error Handling:** Includes robust error handling for file operations and data processing.
- **Testing:** Comprehensive unit tests ensure functionality across different scenarios.

## Usage

To use the tool, follow these steps:
1. Compile the Go program.
2. Run the executable with two arguments:
    ```cmd
    ./text-tool input.txt output.txt
    ```
    - **input.txt**: File containing the text to be modified.
    - **output.txt**: File where modified text will be saved.

## Example

### Input (input.txt)
```cmd
1E (hex) files were added. It has been 10 (bin) years.
Ready, set, go (up) ! I should stop SHOUTING (low).
Welcome to the Brooklyn bridge (cap).
"There it was. A amazing rock!"
```

### Output (output.txt)
```cmd
30 files were added. It has been 2 years.
Ready, set, GO ! I should stop shouting.
Welcome to the Brooklyn Bridge.
"There it was. An amazing rock!"
```

## Conclusion

This project demonstrates proficiency in Go programming, file system manipulation, and text processing techniques. It serves as a robust tool for automating text modifications based on specified rules and enhances understanding of string and number manipulation in Go.

---

## License
This project is licensed under the [MIT License](LICENSE).