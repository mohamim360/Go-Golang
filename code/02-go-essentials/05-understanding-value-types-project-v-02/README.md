

## 📌 Go Types, Conversions & Power Calculation (Short Notes)

### **1. Go is statically typed**

* Every value has a **specific type** (int, float64, string, etc.).
* Types matter because Go **doesn’t allow mixing incompatible types** in calculations.

### **2. Variable types**

```go
var investmentAmount = 1000        // int → whole number
var expectedReturnRate = 5.5       // float64 → decimal number
```

### **3. Why type errors happen**

You **cannot** mix `int` and `float64` in math:

```go
investmentAmount * (1 + expectedReturnRate/100) // ❌ type mismatch
```

### **4. Fix: Convert int → float64**

Use Go’s built-in conversion:

```go
float64(investmentAmount)
```

This makes the calculation valid:

```go
futureValue := float64(investmentAmount) * (1 + expectedReturnRate/100)
```

---

## 📌 Adding Power Calculation with `math.Pow`

Go has **no power operator**, so we use:

```go
import "math"
```

`math.Pow(base, exponent)` requires **float64** values.

### **5. Apply the investment formula**

```
futureValue = P × (1 + r/100)^t
```

Go version:

```go
futureValue := float64(investmentAmount) * 
    math.Pow(1 + expectedReturnRate/100, float64(years))
```

### **6. Output the result**

```go
fmt.Println(futureValue)
```

---
