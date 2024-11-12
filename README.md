# fucli: A CLI Tool for Text Mutation and Fuzzing

## Description

fucli is a command-line tool designed for text mutation and fuzzing. It allows you to apply various transformations to input text, generating a wide range of variations for testing, analysis, or creative purposes.

## Features

* **Basic Mutations:**
    * **Case Change:**  Convert text to uppercase, lowercase, or toggle case.
    * **Vowel Change:** Replace vowels with other vowels or random characters.
    * **Fake Word Generation:** Create plausible-looking words based on the input text.
    * **Random Word Insertion:** Insert random words from a provided wordlist.
* **Advanced Mutations:**
    * **Insertion:** Insert letters, numbers, or special characters at random positions.
    * **Deletion:** Delete letters, numbers, or special characters at random positions.
    * **Replacement:** Replace letters, numbers, or special characters with random alternatives.
    * **Swapping:** Swap characters at random positions.
    * **Duplication:** Duplicate characters at random positions.
    * **Reversal:** Reverse the input text.
    * **Picking:** Select a random subset of characters from the input text.
    * **HTML Element Insertion:** Insert HTML elements at random positions.
    * **Boundary Value Insertion:** Insert boundary values (e.g., minimum, maximum) at random positions.
    * **Random JSON Generation:** Generate random JSON objects.
* **Output Options:**
    * **Direct Output:** Print the mutated text to the console.
    * **File Output:** Save the mutated text to a file.
    * **Network Output:** Send the mutated text over TCP or UDP.
    * **Bluetooth Output:** Send the mutated text over Bluetooth to a specified device.

## Installation

```bash
# Install Go
# ...

# Install fucli
go get github.com/jaavier/fucli
```

## Usage

```bash
# Basic usage
fucli -sd "example" -il 3 -dl 2

# Use a wordlist for random word insertion
fucli -sd "example" -rw 2 -w wordlist.txt

# Generate random JSON objects
fucli -rj 5

# Send mutated text over TCP
fucli -sd "example" -il 3 -tcp localhost:8080

# Send mutated text over UDP
fucli -sd "example" -il 3 -udp localhost:8080

# Send mutated text over Bluetooth
fucli -sd "example" -il 3 -bt-mac 00:11:22:33:44:55

# Scan for Bluetooth devices
fucli -bt-scan
```

## Examples

**Example 1: Inserting and Deleting Characters**

```bash
fucli -sd "example" -il 3 -dl 2
```

This command will insert 3 random letters and delete 2 random characters from the seed text "example".

**Example 2: Generating Random JSON Objects**

```bash
fucli -rj 5
```

This command will generate 5 random JSON objects.

**Example 3: Sending Mutated Text over TCP**

```bash
fucli -sd "example" -il 3 -tcp localhost:8080
```

This command will insert 3 random letters into the seed text "example" and send the mutated text to the TCP server at localhost:8080.

## Use Cases

* **Fuzzing:** Generate a wide range of input variations for testing software robustness.
* **Security Testing:** Create malicious input for penetration testing.
* **Data Generation:** Generate synthetic data for testing or analysis.
* **Creative Writing:** Experiment with text transformations for creative writing projects.
* **Code Generation:** Generate variations of code for testing or analysis.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.