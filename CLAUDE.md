# go-primitives

> ___

## What this repo is

- **Domain**: Go library — blockchain primitives and type definitions for Trust Wallet backends.
- **Route here for**: blockchain coin registry (coin IDs, handles, symbols, decimals, explorer URLs); canonical transaction/token/asset types (`Tx`, `Token`, `TokenType`, `Collection`); asset ID encoding/decoding (`c<coinID>_t<tokenID>` format); EVM address utilities (EIP-55 checksum); decimal/satoshi number conversions; slice batch/chunking utilities.
- **Do not route here for**: transaction signing, private keys, or mnemonics (wallet-core); mobile platform code (mobile-monorepo); RPC clients or node management (wallet-kit-go); swap or DeFi logic; UI components.
- **Consumers**: Any Trust Wallet Go backend service that needs blockchain type definitions or the coin registry. Imports as `github.com/trustwallet/go-primitives`.
- **Ships**: Go module `github.com/trustwallet/go-primitives` (library, no binary). Key exports: `coin.Coins` map, `coin.Ethereum()` etc. accessors, `types.Tx`, `types.Token`, `types.TokenType`, `asset.ParseID`/`BuildID`, `address.EIP55Checksum`, `numbers.ToDecimal`.
- **Agent map**: new chain → `knowledge/architecture/coin-registry.md`; token types → `knowledge/architecture/token-type-registry.md`; transaction model → `knowledge/architecture/transaction-model.md`; asset ID → `knowledge/architecture/asset-id-encoding.md`; number conversions → `knowledge/libs/numbers.md`; build/CI → `knowledge/build/overview.md`; tests → `knowledge/tests/testing-strategy.md`.

## Knowledge Map

For the structured knowledge base, see [knowledge/constitution.md](knowledge/constitution.md).

- [code-conventions](knowledge/code-conventions/index.md) — Code conventions, style rules, and decision records
- [libs](knowledge/libs/index.md) — Core libraries and shared utilities
- [patterns](knowledge/patterns/index.md) — Coding patterns, recipes, and proven approaches

- [architecture](knowledge/architecture/index.md) — Architecture
- [build](knowledge/build/index.md) — Build
- [features](knowledge/features/index.md) — Features
- [tests](knowledge/tests/index.md) — Tests

## Learnings

This repo may keep a living archive of incident-derived rules in ~~[`learnings/`](learnings/)~~ — each file a postmortem of a real bug or a non-obvious pattern that bit once and would bite again: root cause, the rule that prevents recurrence, and tags for matching. The folder is **optional and may be absent** — create it the first time you have a learning worth saving.

**Before** investigating any bug, regression, or "weird behavior", *if a `learnings/` directory exists*:

1. Search the frontmatter directly — it's the source of truth and always present:
   - `grep -ril "<keyword>" learnings/` — matches the frontmatter `tags:`/`summary:` + body.
   - `ls learnings/ | grep -i "<keyword>"` — matches the slug-style filename.
   - Skim each match's `summary:` line to decide whether to read the full body.
2. For a topic-organized ToC (grouped by surface + a tag index), open `learnings/index.md`. It is a **generated** artifact that garden **always regenerates** from frontmatter — never hand-edit it (any edit is discarded next run). Depending on the repo it's either gitignored (a derived artifact) or committed; either way it can be stale if a learning file changed without a regen, so prefer reading the learning files' frontmatter over trusting it blindly.
3. Found a match? **Read it before forming a hypothesis** — a 30-second read can turn a 2-hour investigation into a 5-minute fix.
4. Every file has frontmatter (`title`, `date`, `area`, `files`, `symptom`, `tags`, `summary`; `pr` when tied to a specific PR). `area` drives the index's surface grouping; `tags` drive its tag index.

**After** any fix, feature, or non-trivial change — if you learned something not already obvious from the code:

1. Add a new file `learnings/<slug>.md` with the frontmatter above, then a body covering: the symptom, the root cause (the actual mechanism, not just "the bug"), why prior fixes weren't enough if applicable, the rule going forward, and any regression guards. Create the `learnings/` folder if it doesn't exist yet.
2. If the learning extends an existing entry, edit that file instead of creating a duplicate.
3. Make the new file's `area`, `tags`, and `summary` accurate — those drive both `grep` and the generated index (`area` → its surface grouping, `tags` → its tag index, `summary` → its hook). **Never hand-edit `learnings/index.md`** — it's generated and always regenerated; edit the learning file's frontmatter instead.
4. Commit the learning **in the same PR as the fix** — never as a follow-up.

The bar: would a future agent save time by reading this before touching the same surface? If yes, write it; if it would just say "read the diff," skip it. Don't ask which learnings to capture — commit every candidate that clears the bar.

## Repository Knowledge Scope

This repo's `knowledge/` covers: **code-conventions, libs, patterns**

Topics NOT documented locally: architecture, build, ci, conventions, core-libs, decisions, design, features, git-conventions, guides, observability, product, quality, security, tests, workflows, brand, business, legal, hr, prompts, api, specs, components, references

## Constraints

- [TODO: Add project-specific constraints]

<!-- sdd-knowledge-generated -->
