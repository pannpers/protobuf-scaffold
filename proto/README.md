# Protocol Buffers Definitions

This directory contains entity and API definitions for the application.

## Directory Structure

```
proto/
├── entity/          # Entity definitions
├── api/             # API definitions
└── common/          # Common definitions
```

## Usage

1. Create new proto files
2. Generate code with `buf generate`
3. Check code quality with `buf lint`
4. Check breaking changes with `buf breaking`
