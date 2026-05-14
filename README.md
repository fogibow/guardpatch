# GuardPatch

**Defensive AI Patch Assistant for secure vulnerability remediation**

GuardPatch is an experimental open-source tool that helps developers and security engineers analyze vulnerable code patterns, generate candidate security patches, create defensive regression test ideas, and validate fixes in a controlled workflow.

GuardPatch is designed for **defensive security use only**.

---

## Why GuardPatch?

Security teams often find vulnerabilities faster than development teams can safely remediate them. GuardPatch aims to reduce the time between finding a vulnerability and preparing a reviewed fix.

It helps answer:

- Where is the risky code?
- What is the likely secure fix?
- Why is the fix safer?
- What test should prove the issue is fixed?
- What should a human reviewer check before accepting the patch?

GuardPatch does **not** replace developers, security engineers, or code review. It assists them.

---

## Core Principle

> GuardPatch generates candidate fixes that must be reviewed, tested, and approved by humans before use.

GuardPatch is **human-in-the-loop** by design.

---

## What GuardPatch Does

- Analyzes local source code for risky security patterns
- Detects common insecure coding patterns
- Uses AI to suggest candidate patches
- Explains why a patch may reduce risk
- Suggests defensive regression tests
- Produces remediation reports
- Supports local-first AI workflows using Ollama
- Encourages safe, review-based vulnerability remediation

---

## What GuardPatch Does Not Do

GuardPatch does not:

- Attack third-party systems
- Generate weaponized exploit code
- Automatically patch production systems
- Replace secure code review
- Guarantee that generated patches are correct
- Bypass responsible disclosure processes
- Run against systems without authorization

---

## Current MVP Scope

The first version focuses on simple local source-code analysis and AI-assisted patch suggestions.

### Supported languages in MVP

- Python
- C

### Initial vulnerability patterns

Python:

- `subprocess` with `shell=True`
- unsafe `eval()`
- unsafe `exec()`
- unsafe `pickle.load()` / `pickle.loads()`

C:

- `strcpy()`
- `strcat()`
- `gets()`
- `sprintf()`

### Initial AI provider

- Ollama local models

Planned later:

- OpenAI provider
- Anthropic provider
- Go analyzer
- JavaScript/TypeScript analyzer
- Java analyzer
- Rust analyzer
- Docker sandbox validation
- SARIF output
- GitHub Actions integration
- Pull request workflow

---

## Example Use Cases

GuardPatch can be used for:

- Secure coding practice
- DevSecOps remediation support
- Security research
- Vulnerability management workflows
- Internal AppSec enablement
- Security training labs
- AI-assisted secure code review experiments

Example:

A Python application contains:

```python
subprocess.run("ping -c 1 " + host, shell=True)
