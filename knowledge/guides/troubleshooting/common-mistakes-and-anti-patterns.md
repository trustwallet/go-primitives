---
category: guides
subcategory: troubleshooting
confidence: low
documentType: how-to
scope: org
contentHash: e726341c6c75
tags: [anti-pattern, troubleshooting, faq]
source: (synthesized)
verified: 2026-07-22
synthetic: synthesized-faq
---

## Common Mistakes and Anti-Patterns

<!-- sdd-knowledge-synthesized -->

> Auto-generated from anti-patterns found across the knowledge base.

### Architecture

**From [Coin Registry Architecture](../../architecture/coin-registry-architecture.md):**
- `coin/gen.go` is a program guarded by `//go:build coins` (never compiled into the
  library itself — only executed when you run `make generate-coins`).
- The Makefile target runs `goimports` after generation to keep the output tidy.

**From [Coin Registry Architecture](../../architecture/coin-registry-architecture.md):**
- **Never hand-edit `coins.go`** — changes will be overwritten by the next `make
  generate-coins` run.

**From [Coin Registry Architecture](../../architecture/coin-registry-architecture.md):**
This is the canonical way to check if a coin is an EVM chain — do not compare `ChainID`
or check `Blockchain` directly in calling code.

**From [Central Symbols — God-Node Analysis](../../architecture/central-symbols-god-node-analysis.md):**
key semantic: callers do not need to know which representation was stored.

**Invariants:**

**From [Token Type Registry](../../architecture/token-type-registry.md):**
**Never reuse a version integer**; always assign the next increment when adding a new
`TokenType`.

### Build

**From [Build and CI Overview](../../build/build-and-ci-overview.md):**
`//go:build coins` guards the generator so it is never compiled into the library.

### Code-conventions

**From [Anti-Patterns (failed approaches)](../../code-conventions/anti-patterns-failed-approaches.md):**
<!-- Add failed approaches here. Each anti-pattern should include:
- **type**: anti-pattern
- **discovered**: YYYY-MM-DD

**From [Go Code Conventions](../../code-conventions/go-code-conventions.md):**
smallest unit (Wei, Satoshis, etc.). It is NEVER a floating-point representation.
The `numbers` package provides `ToDecimal` and `FromDecimal` to convert for display.

