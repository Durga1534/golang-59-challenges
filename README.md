# Go Engineering Path 🚀

A structured repository dedicated to mastering Go through specification-driven development, moving beyond tutorial-based learning.

## 🧠 Why I Pivoted
I previously maintained a "59-Day Challenge" repository focused on daily consistency. However, I chose to retire that repository to prioritize **Engineering Quality over Daily Quantity**. 

**Why I made the switch:**
*   **Active vs. Passive Learning:** I moved away from "follow-along" YouTube tutorials to focus on building from technical specifications.
*   **Architectural Debt:** I wanted a clean slate to implement professional folder structures (`cmd/`, `internal/`) from the first commit.
*   **Security & Best Practices:** This new repository ensures 100% compliance with secret management (no leaked tokens) and optimized performance for resource-constrained environments.

## 🏗️ The Methodology: Spec-Driven Learning
Every project in this repository is built using the following professional workflow:
1.  **Requirement Analysis:** Breaking down a spec (e.g., from [roadmap.sh](https://roadmap.sh/projects)) into actionable Go structs.
2.  **Architecture Design:** Defining clear boundaries between business logic and I/O.
3.  **Implementation:** Writing clean, idiomatic Go with an emphasis on the standard library.
4.  **Observability:** Using structured logging and proper error handling.

## 📂 Level-Based Roadmap

| Level | Project | Core Concepts | Status |
| :--- | :--- | :--- | :--- |
| **Beginner** | [Task Tracker](./01-beginner/task-tracker) | JSON Persistence, CLI `os.Args`, File I/O | 🚧 In Progress |
| **Intermediate** | URL Shortener | Redis Caching, HTTP Middleware, Context | ⏳ Planned |
| **Advanced** | [Analytics Ingestor](https://github.com/Durga1534) | Redis Streams, PostgreSQL, Concurrency | ✅ Completed |

## 🛠️ Tech Stack & Standards
*   **Language:** Go (Golang)
*   **Patterns:** Clean Architecture, Dependency Injection (via Interfaces)
*   **Tooling:** Makefiles for automation, Go Workspaces for monorepo management
*   **Hardware Constraint:** Optimized for 8GB RAM environments to ensure high-performance, low-memory footprints.

## 🚀 Getting Started
This repository uses **Go Workspaces**.
```bash
git clone [https://github.com/Durga1534/go-engineering-path.git](https://github.com/Durga1534/go-engineering-path.git)
cd 01-beginner/task-tracker
make run
