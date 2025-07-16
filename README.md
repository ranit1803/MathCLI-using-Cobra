# MathCLI - Command Line Math Toolkit in Go

MathCLI is a fun, beginner-friendly command-line application written in Go using the [Cobra](https://github.com/spf13/cobra) library. It's built to help learn and visualize essential math logic problems — all while mastering CLI development in Go.

From checking if a number is a palindrome, to reversing digits or generating the Fibonacci sequence — this CLI does it all in seconds.

---

## Features

✅ Generate a Fibonacci sequence  
✅ Check if a number is **odd or even**  
✅ Check if a number is a **palindrome**  
✅ Reverse the digits of a number  
✅ Built-in flag support and error handling

---

## Project Structure

```
mathcli/
│
├── cmd/                    # Contains all Cobra subcommands
│   ├── fibbonacci.go      # Fibonacci logic
│   ├── oddeven.go         # Odd/Even logic
│   ├── palindrome.go      # Palindrome logic
│   └── reverse.go         # Reverse number logic
│
├── main.go                # Entry point
├── go.mod / go.sum        # Go module files
├── LICENSE                # MIT License
└── README.md              # You're here!
```

---

## 🛠️ Installation

1. **Clone the repository**
```bash
git clone https://github.com/your-username/mathcli.git
cd mathcli
```

2. **Build the binary**
```bash
go build
```

3. **Run the CLI**
```bash
./mathcli --help
```

---

## Commands & Examples

### Fibonacci
Generate Fibonacci series of `n` terms:

```bash
./mathcli fibbonacci -l 6
# Output: [0 1 1 2 3 5]
```

### Palindrome
Check if a number is a palindrome:

```bash
./mathcli palindrome -n 121
# Output: Palindrome
# Output: true
```

### Reverse
Reverse the digits of a number:

```bash
./mathcli reverse -n 1234
# Output: The original number is: 1234
# Output: The reversed number is: 4321
```

### Odd or Even
Determine if a number is odd or even:

```bash
./mathcli oddeven -p 42
# Output: The Number is Even
```

---

## 🧠 Why This Project?

This CLI was built as a personal learning project to:
* Practice coding common math logic problems
* Master the Cobra CLI framework in Go
* Get comfortable with project structure, flags, and commands in Golang

---

## License

This project is licensed under the MIT License.

---

## 🤝 Contributing

This is a solo learning project for now, but contributions are welcome in the form of suggestions, issues, or pull requests.

---
**Built by Ranit Santra**
