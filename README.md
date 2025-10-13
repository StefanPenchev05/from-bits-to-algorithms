# 🧠 From Bits to Algorithms

Welcome to **From Bits to Algorithms** — a personal reasoning dojo designed to forge deep computational intuition from the ground up.

> This repo is not about writing code fast.  
> It’s about understanding every operation, every trade-off, and every byte that moves through your program.

Each challenge is handcrafted to train **reasoning**, **problem-solving**, and **algorithmic clarity** — not memorization.

---

## 🎯 Purpose

Train across seven fundamental axes of computing:

1. **Go (Golang)** — primary language for systems and backend exercises  
2. **Algorithms** — core data structures, optimization, and proofs  
3. **Security** — defensive thinking, safe design, cryptographic reasoning  
4. **Backend (Go / JS)** — protocols, APIs, concurrency, durability *(JS backend allowed)*  
5. **Low-level Programming (Go/C)** — memory, pointers, CPU model, concurrency  
6. **Mathematical Machine Learning** — implement the math yourself  
7. **Time & Memory Complexity** — measure, reason, and optimize

This repository grows **challenge by challenge**, from bit-level logic to multi-system design.

---

## ⚙️ Rules

### ✅ Allowed
- **Only the standard library** (e.g., `fmt`, `os`, `bufio`, `strings`, `math`, `time`, etc.)
- Your own testing harnesses written from scratch
- Paper/pencil math, personal notes
- Conceptual discussion, proofs, references (no code)

### 🚫 Not Allowed
- **No external libraries/packages** (Go or JS) — **standard library only**
- **No AI-generated code** or autocomplete “solutions”
- **No copying code** from the web or books
- **No frameworks** that hide core logic (unless a challenge explicitly permits)

---

## 🧩 How Challenges Work

Each challenge arrives in a new folder under `challenges/`, e.g. `0001-<short-title>/`, and includes:

- **Problem Description** — clear goals and behavior
- **Input/Output Specs** — with examples
- **Constraints & Edge Cases** — adversarial conditions to consider
- **Allowed Standard Packages** — explicitly listed
- **Complexity Targets** — desired time/space bounds
- **Difficulty** — 1 (warm-up) to 5 (expert)
- **Suggested Time** — e.g., 30–60 min, 1–2 h, half-day, 1–2 days, 2–4 days
- **Stretch Goals** — optional deeper optimizations or variants
- **Deliverables** — code + `ANALYSIS.md` (reasoning, invariants, complexity)

---

## 📦 Repository Structure

```
from-bits-to-algorithms/
├── README.md
├── challenges/
│   └── 0001-/
│       ├── README.md        # Challenge spec
│       ├── src/             # Your implementation (std lib only)
│       ├── tests/           # Your test cases
│       ├── inputs/          # Provided inputs
│       ├── expected/        # Expected outputs for samples
│       └── ANALYSIS.md      # Reasoning, invariants, complexity, test plan
├── tracking/
│   ├── PROGRESS.md          # Dates, attempts, results, reflections
│   └── LEARNINGS.md         # Key takeaways after each challenge
├── .challenges/
│   └── index.json           # Challenge counter & metadata (auto-maintained)
└── common/
    └── docs/                # Big-O sheets, math notes, reference summaries
```

---

## 🧪 Submission & Feedback

1. Implement the challenge using **only the standard library**.
2. Include `ANALYSIS.md` describing:
   - Approach & invariants
   - Correctness argument (brief proof/intuition)
   - Time & space complexity (Big-O + practical notes)
   - Test plan and notable cases
3. Submit your code for evaluation.

**Feedback policy:**  
I will run your solution on visible and hidden tests and return **only the wrong outputs** (input → your output vs expected).  
No line numbers, no direct fixes. Ask for **conceptual** hints if needed.

---

## 📊 Difficulty & Time Guide

| Level | Description        | Suggested Time     |
|------:|--------------------|--------------------|
| 1     | Warm-up / Logic    | 30–60 min          |
| 2     | Core Algorithm     | 1–2 hours          |
| 3     | Tough Challenge    | Half-day           |
| 4     | Advanced System    | 1–2 days           |
| 5     | Expert / Research  | 2–4 days           |

---

## 🧭 Roadmap

- Stage 1: Core algorithms & data structures in Go  
- Stage 2: Performance analysis and profiling (std lib only)  
- Stage 3: Security & backend design (Go/JS backend)  
- Stage 4: Low-level reasoning (Go/C, memory model)  
- Stage 5: Math-first ML implementations (no ML libs)

---

## 📚 Reference Materials (no code copying)

- Go Docs (Standard Library): https://pkg.go.dev/std  
- Effective Go: https://go.dev/doc/effective_go  
- CLRS — *Introduction to Algorithms*  
- CS:APP — *Computer Systems: A Programmer’s Perspective*  
- TAOCP — *The Art of Computer Programming* (Knuth)  
- Bishop — *Pattern Recognition and Machine Learning*

Use these to understand concepts and proofs — not to copy implementations.

---

## 🧾 Honor Code

This is a training ground for **your** mind.  
Every mistake and fix is part of your craft.

- Write your own code  
- Cite conceptual inspirations  
- Prefer simplicity first; optimize with evidence

**Welcome to the dojo.** Every bit counts — every algorithm is earned.