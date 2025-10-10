# GoCalc CLI Calculator

A **command-line calculator** built with Go, demonstrating real-world Go development concepts including CLI arguments, error handling, file I/O, and deferred cleanup. This project is designed for learning Go by **building practical projects**.

---

## 🧰 Features

* Perform basic arithmetic operations: `+`, `-`, `*`, `/`
* Proper error handling for invalid inputs and division by zero
* Save calculation history in a `history.txt` file
* View history of past calculations using a single command
* Built with clean Go practices using `defer`, `strconv`, and `os` packages

---

## ⚡ Concepts Learned

### 1. Go Setup

* Install Go and verify installation:

```bash
go version
```

* Initialize project module:

```bash
go mod init gocalc
```

### 2. Go Basics

* Variables, types, and constants
* Functions and return values
* Structs and methods (basic understanding)
* Arrays, slices, and maps

### 3. Command-Line Arguments

* Access arguments via `os.Args`
* Parse strings to numbers using `strconv.ParseFloat`
* Handle missing or extra arguments gracefully

### 4. Error Handling

* Go’s idiomatic error checking using `if err != nil`
* Handling invalid user input or runtime errors

### 5. Control Flow

* Switch statements for operations
* Early return on invalid input or errors

### 6. File I/O

* Open/create file using `os.OpenFile` with flags:

  * `os.O_CREATE` → create if not exists
  * `os.O_APPEND` → append to file
  * `os.O_WRONLY` → write only
* Write strings to file with `WriteString`
* Read file contents with `os.ReadFile`

### 7. Deferred Cleanup

* Use `defer f.Close()` to ensure files are always closed, avoiding resource leaks

---

## 🚀 Installation & Usage

1. Clone the repository:

```bash
git clone https://github.com/VINODvoid/go_projects
cd go_projects/gocalc
```

2. Run a calculation:

```bash
go run main.go 10 + 5
```

3. View calculation history:

```bash
go run main.go history
```

---

## 📝 Example Usage

```bash
go run main.go 12 * 3
# Output: Result: 36

go run main.go 100 / 0
# Output: Error: division by zero

go run main.go history
# Output:
# 10.000000 + 5.000000 = 15.000000
# 12.000000 * 3.000000 = 36.000000
```

---

## 📚 Packages & Functions Used

* `fmt` → Printing to console and formatting strings
* `os` → Accessing CLI args, reading/writing files
* `strconv` → Converting strings to float64
* `defer` → Ensuring cleanup (closing files)

---

## ✅ Key Takeaways

* Always check errors (`err != nil`) for robust programs
* CLI programs in Go can be simple yet powerful
* File I/O and `defer` are essential for real-world apps
* Structuring small projects with modular code helps future scalability

---
**Author:** Kalki
**Project Type:** Learning / Practice Go Project
**Language
