# jsonly

`jsonly` is a command-line tool written in Go that filters URLs from input and only prints those that match the following criteria:
1. The URL must be a valid HTTP/HTTPS URL.
2. The URL must have a `.js` file extension.

## Features

- Reads URLs from `stdin`.
- Filters only URLs that end with `.js`.
- Displays usage if no parameters or input are provided.

## Usage

### Basic Usage

```bash
jsonly < input_file
```

This command will read from the `input_file`, filter out any non-`.js` URLs, and print the matching ones to `stdout`.

### Example

Given the following input (`input.txt`):

```txt
http://example.com/script.js
http://example.com/index.html
https://example.com/other.js
https://example.com/script.js
```

Running the tool like this:

```bash
cat input.txt | jsonly
```

Would output:

```txt
http://example.com/script.js
https://example.com/other.js
https://example.com/script.js
```

### No Input Provided

If no input is provided (i.e., the program is run without a pipe or `input_file`), it will display the usage message and exit:

```bash
Usage: jsonly < input_file
Filters URLs that end with .js from stdin input.
```

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/0xsudomode/jsonly.git
   cd jsonly
   ```

2. Build the binary:

   ```bash
   go build -v jsonly.go
   ```

3. Run the binary:

   ```bash
   ./jsonly < input_file
   ```

   Or pipe input directly to it:

   ```bash
   linkfinder -o cli -i "http://example.com" | jsonly
   ```

