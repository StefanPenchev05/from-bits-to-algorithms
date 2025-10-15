# 🧩 Challenge 003: Recursive Maze Solver

## Overview
**Difficulty:** ⭐⭐⭐ (3/5)  
**Time Estimate:** 1-2 hours  
**Focus:** Recursion, Backtracking, 2D Array Navigation

This challenge is designed to strengthen your understanding of **recursive algorithms** and **backtracking techniques**. You'll navigate through a 2D maze using pure recursion—no loops allowed!

---

## Problem Description

You are given a 2D grid (matrix) representing a maze where:
- **🟢 `1`** = Open path (walkable)
- **🔴 `0`** = Wall (blocked)

### Objective
Find if a path exists from the **entrance** to the **exit**:
- **🚪 Entrance:** Top-left corner `(0, 0)`
- **🏁 Exit:** Bottom-right corner `(n-1, m-1)`

### Movement Rules
- ✅ Move **up, down, left, right** (4-directional)
- ❌ **No diagonal moves**
- ❌ **Cannot pass through walls** (`0`)
- ❌ **Cannot revisit cells** (avoid infinite loops)

### 🔄 Implementation Constraint
**RECURSION ONLY** — No `for`, `while`, or any iterative loops permitted!

---

## Rules & Constraints

### ✅ Allowed
- **Standard packages:** `fmt`, `os`, `time`
- **Recursion and backtracking**
- **Manual 2D array definition**

### 🚫 Forbidden
- ❌ **Loops** (`for`, `while`, `range`)
- ❌ **Global variables** for maze state
- ❌ **External libraries**
- ❌ **Built-in pathfinding functions**

---

## Example

### Input Maze:
```go
maze := [][]int{
    {1, 0, 1, 1},
    {1, 1, 0, 1},
    {0, 1, 1, 0},
    {0, 0, 1, 1},
}
```

### Visual Representation:
```
🟢 🔴 🟢 🟢
🟢 🟢 🔴 🟢
🔴 🟢 🟢 🔴
🔴 🔴 🟢 🟢
```

### Expected Path (if exists):
```
🚪 🔴 ⬜ ⬜
➡️ ⬇️ 🔴 ⬜
⬜ ➡️ ➡️ 🔴
⬜ ⬜ ⬇️ 🏁
```

---

## 🎯 Deliverables

1. **`main.go`** — Complete recursive maze solver
2. **`ANALYSIS.md`** — Algorithm analysis including:
   - Recursive approach explanation
   - Time/Space complexity
   - Backtracking strategy
   - Test cases and edge conditions

---

## 🧪 Test Cases to Consider

- **✅ Solvable maze** with multiple paths
- **❌ Unsolvable maze** (no path exists)
- **🔄 Single path** maze (unique solution)
- **📏 Edge cases:** 1x1, entrance/exit blocked
- **🎯 Corner cases:** Maze with cycles

---

## 🚀 Stretch Goals (Optional)

1. **📍 Path Tracking:** Return the actual path coordinates
2. **🎨 Visual Output:** Print the maze with the solution path
3. **⚡ Optimization:** Implement memoization for visited cells
4. **📊 Analytics:** Count total recursive calls made

---

**Good luck, and may your recursion be ever in your favor!** 🍀