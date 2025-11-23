

# 🚀 Go Basics — Packages, Modules & Build

*A short summary for future reference*

## 📦 1. Packages in Go

* Every Go file **must start** with a `package` declaration.
* Packages organize your code into logical groups.
* A project can have:

  * multiple **packages**
  * each package split across multiple **files**
* Without a package statement, Go shows:
  **“expected ‘package’, found ‘import’”**

---

## 🏁 2. Why `package main` is special

* You can name packages anything (e.g., `app`, `utils`, `db`)
* But **`package main`** is **reserved by Go**.
* It tells Go:

  > “This is the entry point of the application.”
* Only code inside `package main` + `func main()`
  can produce an **executable file**.

Example:

```go
package main

func main() {
    // program starts here
}
```

If you rename it (e.g., `package app`),
`go build` **will not generate an executable**.

---

## 📥 3. Importing Standard Library Packages

Example:

```go
import "fmt"
```

* Allows use of `fmt.Print`, `fmt.Println`, etc.
* `fmt` is part of Go’s huge standard library.

---

# 📦 4. Go Modules (`go.mod`)

Go Modules are required in all modern Go projects.
Creating a module marks your project as a proper Go application.

```sh
go mod init module-name
```

Example module name:

```
example.com/first-app
```

This creates the **go.mod** file, which declares the module path.

---

## 🧩 Why We Use Go Modules (Very Important)

### **1. Defines your project**

A module tells Go:

> “This folder is a Go project.”

Without it, Go cannot build properly.

---

### **2. Manages dependencies**

If your project imports packages from GitHub or elsewhere,
`go.mod` tracks:

* package names
* versions
* updates

It ensures your project is **stable and reproducible**.

---

### **3. Enables `go build` & `go run`**

Without a module:

* `go build` may fail
* Go doesn't know your project’s structure
* import paths may break

---

### **4. Required for sharing your code**

Your module name becomes your import path:

```
github.com/yourname/project
```

---

### **5. Makes the project portable**

`go.mod` + `go.sum` allow the exact same build anywhere, anytime—
even years later.

---

# 🛠 5. Running & Building Go Programs

### ▶ Development (Run a file directly)

```sh
go run main.go
```

### 🏗 Production (Build an executable)

```sh
go build
```

Output:

* **Windows:** `first-app.exe`
* **Mac/Linux:** `first-app` (run using `./first-app`)

Executables run **without Go installed**.

---

# ⚠ When `go build` Fails

No executable will be created if:

❌ `package main` is missing
❌ No `main()` function
❌ No `go.mod` file
❌ Module path not initialized

---

# ✔ Quick Memory Notes

* `package main` → required for making an executable
* `func main()` → program entry point
* `import` → bring features from other packages
* `go mod init` → create module/project
* `go run` → run directly
* `go build` → create executable

---

