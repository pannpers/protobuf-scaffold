# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Essential Commands

### Development Setup
```bash
mise install          # Install dependencies (buf, pre-commit)
pre-commit install     # Install commit hooks
pre-commit install --hook-type pre-push  # Install push hooks
```

### Code Generation
```bash
buf generate           # Generate Go and TypeScript code from protobuf files
```

### Validation and Formatting
```bash
buf lint               # Lint Protocol Buffers files
buf format -w          # Format protobuf files in place
buf breaking --against '.git#branch=main'  # Check for breaking changes
```

### Verification
```bash
buf --version          # Verify buf is available
```

## Architecture Overview

This is a Protocol Buffers scaffold repository that generates code for both Go and TypeScript from protobuf definitions. The architecture follows a clear separation between entity definitions and API service definitions.

### Key Structure
- **Entity Layer**: Core business entities defined in `pannpers/entity/v1/` (e.g., Post, User)
- **API Layer**: Service definitions in `pannpers/api/v1/` that use entity types
- **Generated Code**: Go code in `gen/go/` and TypeScript in `gen/typescript/`

### Code Generation Flow
The repository uses Buf v2 with remote plugins:
- `buf.build/protocolbuffers/go` - Standard Go protobuf generation
- `buf.build/connectrpc/go` - Connect RPC Go bindings
- `buf.build/bufbuild/es` - TypeScript protobuf generation
- `buf.build/bufbuild/connect-es` - Connect RPC TypeScript bindings

### Proto Package Structure
- `pannpers.entity.v1` - Core entities (Post, User, value objects)
- `pannpers.api.v1` - Service definitions that reference entities
- Uses proper import paths and Go package declarations

### Automated Workflow
The repository uses pre-commit hooks for quality control:
- **Commit hooks**: buf lint, format, breaking change detection, prettier
- **Push hooks**: Automatic code generation and commit of generated files

### Key Files
- `buf.yaml` - Buf v2 configuration with googleapis dependencies
- `buf.gen.yaml` - Code generation plugin configuration
- `.pre-commit-config.yaml` - Hook definitions for quality gates
- `.mise/config.toml` - Tool version management

## Protobuf Design Guidelines

When designing entities and RPC interfaces, follow these established standards:

### Reference Documentation
- [Buf Style Guide](https://buf.build/docs/best-practices/style-guide/)
- [Protobuf Files and Packages](https://buf.build/docs/reference/protobuf-files-and-packages/)
- [Google AIP-190: Protobuf Design](https://google.aip.dev/190)
- [Google AIP General Guidelines](https://google.aip.dev/general)

### Naming Conventions (Buf Style Guide)
- **Packages**: Use `lower_snake_case` with version suffix (e.g., `pannpers.entity.v1`)
- **Messages**: Use `PascalCase` for message names
- **Fields**: Use `lower_snake_case` for field names
- **Services**: Use `PascalCase` with `Service` suffix (e.g., `PostService`)
- **RPCs**: Use `PascalCase` with VerbNoun pattern (e.g., `GetPost`, `CreatePost`)
- **Enums**: Names in `PascalCase`, values in `UPPER_SNAKE_CASE`
- **Zero values**: Should end with `_UNSPECIFIED`

### File Organization
- Package structure should have at least 3 components: `{org}.{purpose}.{version}`
- Directory structure should mirror package hierarchy
- One package per directory
- Use descriptive filenames in `lower_snake_case.proto`

### RPC Interface Design (Google AIP)
- **Standard Methods**: Follow standard CRUD patterns
  - `Get{Resource}` for single resource retrieval
  - `List{Resource}` for collection retrieval  
  - `Create{Resource}` for resource creation
  - `Update{Resource}` for resource modification
  - `Delete{Resource}` for resource removal
- **Request/Response Messages**: Always create custom messages, avoid `Empty`
- **Resource-Oriented Design**: Structure APIs around resources, not actions
- **Field Masks**: Use for partial updates in `Update` operations
- **Pagination**: Implement for `List` operations using `page_size` and `page_token`

### Entity Design Patterns
- **Value Objects**: Create dedicated message types for domain concepts (e.g., `PostId`, `PostTitle`)
- **Composition**: Reference other entities by ID, not nested objects
- **Versioning**: Plan for schema evolution with proper field numbering
- **Documentation**: Over-document with complete sentences using `//` comments

### Breaking Change Management
- Always check breaking changes against main branch
- Reserve field numbers for future use
- Use `optional` fields for new additions to existing messages
- Follow semantic versioning for major breaking changes

## Development Workflow Notes

1. Proto files should be modified directly in the `proto/` directory
2. Generated code in `gen/` is automatically managed by pre-commit hooks
3. Breaking changes are checked against the main branch
4. All protobuf files are automatically formatted and linted on commit
5. Code generation happens automatically on push, not on local development