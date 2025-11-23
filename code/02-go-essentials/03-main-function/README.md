
## 🏁 The `main()` Function in Go

### **1. Why `main()` is required**

* In Go, the **entry point of an executable program** is a function named `main` inside the `main` package.
* Go will automatically call this function when the program starts.
* Without `main()`, Go **does not know which code to execute**.

---

### **2. Code must be wrapped in a function**

* Unlike JavaScript or some other languages, Go does **not execute code directly from top to bottom**.
* Your program logic must be inside functions (most commonly `main()` for the starting code).
* Exceptions: `package` and `import` statements, and a few other declarations, can be outside functions.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

---

### **3. Only one `main()` per package**

* You can have multiple files in the `main` package.
* **But only one `main()` function** is allowed across all files in the same package.
* Having two `main()` functions in the same package will cause a compile-time error:

```
main redeclared in this block
```

---

### **4. Libraries vs Executable Programs**

* If you’re building a **library/package** (like `fmt`):

  * No `main()` function is needed.
  * The package provides reusable functions to other programs.
* If you’re building an **executable program**:

  * `main()` is mandatory in `package main`.

---

### **5. Summary**

* **`package main`** → defines the program entry point package.
* **`func main()`** → defines the first code that runs.
* Only **one `main()` function** per package.
* Packages intended as libraries **do not require `main()`**.

---