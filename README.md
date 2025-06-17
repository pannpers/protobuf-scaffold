# Protobuf Scaffold

A repository for managing application entity and API definitions using Protocol Buffers.

## Prerequisites

- [mise](https://mise.jdx.dev/) must be installed
- [pre-commit](https://pre-commit.com/) must be installed

## Setup

1. Install dependencies:

```bash
mise install
```

2. Verify buf is available:

```bash
buf --version
```

3. Install pre-commit hooks:

```bash
pre-commit install
pre-commit install --hook-type pre-push
```

## Usage

### Generate Code

```bash
buf generate
```

This will generate:

- Go code in `gen/go/`
- TypeScript code in `gen/typescript/`

### Lint

```bash
buf lint
```

### Check Breaking Changes

```bash
buf breaking --against '.git#branch=main'
```

### Format

```bash
buf format -w
```

## Pre-commit Hooks

This repository uses pre-commit hooks to ensure code quality:

### Commit-time hooks:

- **buf-lint**: Checks Protocol Buffers files for style and best practices
- **buf-format**: Automatically formats Protocol Buffers files
- **buf-breaking**: Detects breaking changes in Protocol Buffers
- **prettier**: Formats TypeScript/JavaScript files
- **trailing-whitespace**: Removes trailing whitespace
- **end-of-file-fixer**: Ensures files end with a newline
- **check-yaml**: Validates YAML files
- **check-added-large-files**: Prevents large files from being committed

### Push-time hooks:

- **buf-generate-push**: Generates code from Protocol Buffers files before push
- **auto-commit-generated-push**: Automatically commits generated files before push

The commit hooks run automatically on every commit. The push hooks run before `git push` to ensure generated code is up-to-date.

**Note**: The push hooks will automatically generate code and commit any changes to the `gen/` directory before pushing to the repository.

## Directory Structure

```
.
├── .mise/           # mise configuration
├── proto/           # Protocol Buffers definitions
│   ├── entity/      # Entity definitions
│   ├── api/         # API definitions
│   └── common/      # Common definitions
├── gen/             # Generated code (gitignored)
│   ├── go/          # Generated Go code
│   └── typescript/  # Generated TypeScript code
├── buf.yaml         # buf configuration
├── buf.gen.yaml     # Code generation configuration
├── .pre-commit-config.yaml  # pre-commit configuration
└── README.md        # This file
```

## Development Flow

1. Create new `.proto` files in the `proto/` directory
2. Commit changes (commit hooks will run automatically)
3. Push to repository (push hooks will run automatically)
   - `buf generate` will be executed
   - Generated files will be automatically committed
   - Then the push will proceed
