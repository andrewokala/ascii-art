# ASCII ART GENERATOR
A simple command-line ASCII art generator written in Go. This program takes a string as input and prints it as ASCII art using different banner styles.

### Features
1. Convert test into ASCII art
2. Support for multiple banner styles
3. Handles multi-line input
4. Input validation and error handling

### How It Works
1. Validate command-line arguments
2. Load the selected banner file
3. Split the input into lines
4. Generates ASCII art character by character
5. Render the final output to the terminal

## Installation
### Clone the repository
`git clone <repository-url>`
`cd asacii-art`

### Run the Program
`go run . "Hello"`

### Using a Banner Style
`go run . "Hello" shadow`

### Available Banner Styles
1. shadow
2. standar
3. thinketoy

### Technologies Used
1. Go
2. File handling
3. String manipulation
4. Modular package structure

# Author
Built with Go for learning and practicing backend and text-processing concepts