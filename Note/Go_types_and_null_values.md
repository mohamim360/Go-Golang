

# 📘 Go Types & Null (Zero) Values

Go is a **statically typed** language, meaning every value has a fixed type.
Below are the most important built-in types and their default (zero) values.

---

## 🔢 Basic Types

### **int**

* Whole numbers (no decimals)
* Examples: `-5`, `10`, `42`

### **float64**

* Numbers with decimal places
* Examples: `-5.2`, `10.123`, `12.9`

### **string**

* Text values
* Created using:

  * Double quotes → `"Hello World"`
  * Backticks → `` `Hi everyone` ``

### **bool**

* Logical values: `true` or `false`

---

## 🧩 Additional (Less Common) Types

### **uint**

* Unsigned integer (non-negative only)
* Example: `0`, `10`, `255`

### **int32**

* 32-bit signed integer
* Range: **−2,147,483,648** to **2,147,483,647**

### **rune**

* Alias for **int32**
* Represents a **Unicode code point** (a single character)
* Examples: `'a'`, `'ñ'`, `'世'`

### **uint32**

* 32-bit unsigned integer
* Range: **0** to **4,294,967,295**

### **int64**

* 64-bit signed integer
* Much larger range
* Examples: `-9000000000000000000`, `9000000000000000000`

> 💡 There are also smaller versions: `int8`, `uint8` (same idea, smaller range).

---

## 🟦 Zero (Null) Values in Go

Every Go type has a **zero value**, used when no value is assigned.

Example:

```go
var age int
fmt.Println(age) // 0 → zero value of int
```

### **Zero values for common types:**

| Type    | Zero Value          |
| ------- | ------------------- |
| int     | `0`                 |
| float64 | `0.0`               |
| string  | `""` (empty string) |
| bool    | `false`             |

Zero values ensure that **all variables are always initialized**, preventing undefined behavior.

---
