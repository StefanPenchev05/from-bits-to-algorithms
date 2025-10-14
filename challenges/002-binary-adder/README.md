# Binary Adder

## 🧠 Description
This challenge focuses on understanding how addition works at the binary level — exactly how CPUs do it.  
You will implement a program that adds two integers **using only bitwise operations**, without using any arithmetic operators like `+` or `-`.

You will simulate the behavior of a digital adder circuit that handles both the **sum bits** and the **carry bits**, repeating the process until the carry is gone.

---

## ⚙️ Rules
- ❌ You **cannot** use:
  - `+`, `-`, `*`, `/`, `%`
  - any math library or package
- ✅ You **can** use:
  - `fmt` for I/O
  - bitwise operators: `&`, `|`, `^`, `<<`, `>>`
  - loops and conditionals

---

## 🧩 Example 1

**Input:**
```text
a = 5
b = 3
```

**Binary:**
```text
a = 0101
b = 0011
```

**Process:**
```text
0101
+0011

1000
```

---

## 🎯 Learning Objectives
By completing this challenge, you will:
- Understand **bitwise addition** and **carry propagation**
- Learn how CPUs handle arithmetic logic without math operators
- Strengthen understanding of **binary representation**
- Practice **loop convergence** and **bit-level reasoning**

---

## 🧪 Post-Challenge Tests
Use these to check correctness after implementation:

| a  | b  | Expected Decimal | Expected Binary |
|----|----|------------------|-----------------|
| 1  | 1  | 2                | 10              |
| 2  | 2  | 4                | 100             |
| 5  | 7  | 12               | 1100            |
| 9  | 6  | 15               | 1111            |
| 10 | 15 | 25               | 11001           |

---

## 🧱 Repository Structure
```text
002-binary-adder/
│
├── main.go
├── README.md
└── examples/
├── example1.txt
└── example2.txt
```
---

## 💡 Stretch Goals
- Implement support for **negative numbers** using **two’s complement** logic.
- Write a reusable function:
  ```go
  func BinaryAdd(a, b int) int
  ```

and test it with multiple pairs.
	•	Visualize the carry propagation step-by-step for each bit.
	•	Compare performance vs. the regular + operator (just for curiosity).

## ⏱ Estimated Time

45–60 minutes if you’re comfortable with bitwise operations.
Add 30–45 minutes if this is your first time working with binary logic.