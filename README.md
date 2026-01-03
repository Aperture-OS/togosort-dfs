# Togosort - DFS

A small, **independent Go library made specifically for package managers** to help developers with dependency resolution and handling.

It provides exactly what a package manager needs when dealing with dependencies:

- **Circular dependency detection** using **Depth-First Search (DFS)**
- **Dependency ordering** from **least dependent to most dependent** using **Topological Sort** with *Kahn's Algorithm*

## Purpose

`togosort` is built for real-world dependency resolution, following standard package manager semantics.

If a package **depends on** another package, that dependency must be processed **first**.

The library enforces a clear and explicit convention:

dependency -> dependent

Example:

libc -> bash
bash depends on libc

## What it does

- Validates dependency graphs by detecting cycles (DFS)
- Produces a safe install / build order
- Works with multiple roots (e.g. user-requested packages)
- Has **no external dependencies**
- Designed to be embedded directly into package managers

## What it does NOT do

- It does not resolve versions
- It does not fetch packages
- It does not install anything
- It does not try to be “smart”
- Absolutely **NO AI SLOP**

It only solves **dependency correctness**.

## Typical workflow

1. Build the dependency graph
2. Run **DFS** to detect circular dependencies
3. Run **TopoSort** to get install order
4. Install from first -> last using the provided array

Skipping step 2 is a bug as if you don't detect cycles before sorting it can cause problems

## License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.
Copyright © ApertureOS 2025-2026
