

## 📌 Go Basics – Variables, Naming, and Basic Calculations

### **1. Variables in Go**

* Use the `var` keyword to declare a variable.
* Variables act like **data containers** that store values for later use.
* Naming convention: **camelCase** (e.g., `investmentAmount`, `expectedReturnRate`).

Example:

```go
var expectedReturnRate = 5.5
```

---

### **2. Why use variables?**

* Makes the program **readable**, **organized**, and **easy to update**.
* Instead of hard-coding values everywhere, store them once and reuse them.

---

### **3. Type Mismatch Issue**

* Go is **strongly typed**, meaning each value has a fixed type.
* `investmentAmount` → integer (`int`)
* `expectedReturnRate` → decimal (`float64`)

Mixing `int` and `float64` in math causes an error:

```
mismatched types int and float64
```

Solution (later in project):

* Convert types (e.g., `float64(investmentAmount)`),
  OR
* Declare everything as `float64` from the start.

---