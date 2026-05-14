## Example Use Cases

GuardPatch can be used for:

- Secure coding practice
- DevSecOps remediation support
- Security research
- Vulnerability management workflows
- Internal AppSec enablement
- Security training labs
- AI-assisted secure code review experiments

Example vulnerable Python code:

    subprocess.run("ping -c 1 " + host, shell=True)

GuardPatch can detect the risky pattern and suggest a safer approach such as avoiding shell execution and passing arguments as a list.

---

## Installation

### Prerequisites

- Go 1.21+
- Git
- Ollama, for local AI patch suggestions

Install Ollama on macOS:

    brew install ollama

Start Ollama:

    ollama serve

Pull a coding model:

    ollama pull qwen2.5-coder

Alternative model:

    ollama pull codellama

---

## Clone and Build

    git clone https://github.com/YOURUSERNAME/guardpatch.git
    cd guardpatch
    go mod tidy
    go build -o bin/guardpatch ./cmd/guardpatch

Replace `YOURUSERNAME` with your GitHub username.

---

## Usage

### Analyze code

    ./bin/guardpatch analyze --target ./examples

Example output:

    Detected 2 finding(s):

    - python-command-injection-shell-true
      File: examples/python-command-injection/vulnerable.py:6
      Severity: high
      Code: subprocess.run("ping -c 1 " + host, shell=True)

    - c-unsafe-strcpy
      File: examples/c-buffer-overflow/vulnerable.c:8
      Severity: high
      Code: strcpy(buffer, input);

### Generate candidate patches

    ./bin/guardpatch patch --target ./examples --model qwen2.5-coder

Reports are saved in:

    reports/

---

## Recommended Workflow

1. Analyze local code
2. Review findings
3. Generate candidate patch
4. Review AI output manually
5. Run tests
6. Add or update regression tests
7. Validate behaviour in a safe environment
8. Create pull request
9. Complete code review and security review
10. Merge only after approval

---

## Project Structure

    guardpatch/
    ├── cmd/
    │   └── guardpatch/
    │       └── main.go
    ├── internal/
    │   ├── analyzer/
    │   ├── llm/
    │   ├── patcher/
    │   ├── validator/
    │   ├── reporter/
    │   └── safety/
    ├── pkg/
    │   └── types/
    ├── examples/
    ├── patches/
    ├── reports/
    └── docs/

---

## Safety and Responsible Use

GuardPatch must only be used on code you own or are authorized to test.

For sensitive or proprietary source code, prefer local LLMs such as Ollama.

Do not upload confidential source code to cloud LLM providers unless your organization has approved that use.

Generated patches may be wrong, incomplete, insecure, or unsuitable for production. Always review and test them.

---

## Security Model

GuardPatch follows these design principles:

- Defensive use only
- Local-first AI support
- No weaponized exploit generation
- Human approval required
- Sandbox validation planned
- Clear reporting
- Transparent patch explanation
- Safe-by-default documentation

---

## Roadmap

### v0.1.0

- Project skeleton
- Python analyzer
- C analyzer
- Ollama integration
- Candidate patch generation
- Markdown and JSON reports
- Example vulnerable projects
- Safe usage documentation

### v0.2.0

- Better patch parsing
- Unified diff generation
- Docker sandbox runner
- Basic validation runner
- Unit tests
- GitHub Actions CI

### v0.3.0

- Go analyzer
- JavaScript analyzer
- SARIF output
- Semgrep integration
- OpenAI and Anthropic optional providers
- Patch scoring

### v0.4.0

- Pull request generation
- GitHub App mode
- Policy engine
- Enterprise safe configuration
- Local-only mode enforcement

---

## Contributing

Contributions are welcome.

Good first contributions:

- Add new insecure pattern detectors
- Improve patch parsing
- Add examples
- Add tests
- Improve documentation
- Add sandbox validation
- Add support for more languages

Before contributing, please read:

- `CONTRIBUTING.md`
- `SECURITY.md`
- `SAFE_USAGE.md`

---

## Security Policy

If you find a security issue in GuardPatch, please report it responsibly.

Do not create a public GitHub issue for sensitive security vulnerabilities.

Use the contact method described in `SECURITY.md`.

---

## Disclaimer

GuardPatch is experimental software.

Generated outputs are suggestions only. They require human review, testing, and approval before use.

The maintainers are not responsible for misuse, incorrect patches, data exposure, or damage caused by generated outputs.

---

## License

Apache License 2.0
