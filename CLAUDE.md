# CLAUDE.md

## Project Overview

This is a **Lagrangian Reconstruction** project implementing Shamir's Secret Sharing using Go. It reconstructs cryptographic secrets from polynomial evaluation points stored in various numeric bases using Lagrange interpolation with arbitrary precision arithmetic.

## Development Environment

This project uses **Nix Flakes** for reproducible development environments:

```bash
# Enter development shell (auto-activated with direnv)
nix develop

# Format code (once files are committed to git)
nix fmt

# Check configuration (once files are committed to git)
nix flake check
```

## Core Commands

### Running the Application
```bash
# Run with test cases
go run main.go test1.json    # Simple case (n=4, k=3)
go run main.go test2.json    # Complex case (n=10, k=7, large numbers)
```

### Go Development
```bash
# Format Go code
go fmt ./...

# Check for issues
go vet ./...

# Manage dependencies
go mod tidy

# Build executable
go build
```

## Code Architecture

- **Single-file design**: All logic in `main.go` for simplicity
- **Arbitrary precision**: Uses `math/big` package for large number handling
- **Core functions**:
  - `convertFromBase()` - Converts values from any numeric base to big.Int
  - `lagrangeInterpolation()` - Implements mathematical reconstruction algorithm
  - `Point` struct - Represents (x,y) coordinates with big.Int precision

## Input Format

JSON test files structure:
```json
{
  "keys": {"n": 4, "k": 3},           // n=total shares, k=minimum needed
  "1": {"base": "10", "value": "4"},   // Share 1: base-10 value "4"
  "2": {"base": "2", "value": "111"}   // Share 2: base-2 value "111"
}
```

## Code Style

- **Indentation**: 2 spaces (enforced by .editorconfig)
- **Go conventions**: CamelCase functions, descriptive variable names
- **Error handling**: Explicit error returns with descriptive messages
- **Comments**: Document algorithm steps, especially mathematical operations

## Development Workflow

1. Make changes to Go code
2. Run `go fmt ./...`
3. Test with `go run main.go test1.json`
4. Verify complex case with `go run main.go test2.json`
5. Git commit triggers pre-commit hooks automatically (trailing whitespace, security checks, etc.)

## Nix Configuration Structure

```
flake/
├── devshells.nix    # Go toolchain, LSPs (gopls, nixd)
├── formatters.nix   # alejandra, deadnix, statix
├── checks.nix       # Pre-commit hooks configuration
└── actions/         # GitHub Actions integration
```

The project enforces code quality through pre-commit hooks including private key detection, trailing whitespace removal, and case conflict checking.
