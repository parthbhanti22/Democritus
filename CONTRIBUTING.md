# Contributing to Democritus

Thank you for your interest in contributing to Democritus! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Making Changes](#making-changes)
- [Coding Standards](#coding-standards)
- [Submitting a Pull Request](#submitting-a-pull-request)
- [Reporting Issues](#reporting-issues)

## Code of Conduct

Please be respectful and considerate in all interactions. We welcome contributors of all backgrounds and experience levels.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/democritus.git
   cd democritus
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/parthbhanti22/democritus.git
   ```

## Development Setup

### Prerequisites

- **Go 1.24+** - For backend services
- **Docker & Docker Compose** - For containerized development
- **Node.js 18+ & npm** - For the dashboard
- **Protocol Buffers Compiler** (protoc) - For gRPC development

### Backend Setup

```bash
# Install Go dependencies
go mod download

# Run the scheduler locally (without Docker)
go run scheduler/main.go

# Run a worker locally
go run worker/main.go
```

### Dashboard Setup

```bash
cd dashboard
npm install
npm run dev
```

### Full Stack (Docker)

```bash
# Start everything with 5 workers
docker compose up --build --scale worker=5
```

### Regenerating Protobuf Files

If you modify `proto/democritus.proto`:

```bash
# Install protoc plugins (one-time)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Regenerate Go files
protoc --go_out=. --go-grpc_out=. proto/democritus.proto
```

## Project Structure

```
democritus/
├── scheduler/          # Master node (task distribution)
│   ├── main.go         # Entry point
│   ├── server/         # gRPC server implementation
│   └── store/          # In-memory task store
├── worker/             # Compute nodes
│   ├── main.go         # Entry point
│   └── simulation/     # Physics simulation engines
├── proto/              # Protocol Buffer definitions
├── pkg/                # Shared packages
│   └── physics/        # Physics interfaces
├── dashboard/          # Next.js frontend
├── config/             # Configuration files
├── docs/               # LaTeX documentation
└── docker-compose.yml  # Container orchestration
```

## Making Changes

### Branch Naming

Use descriptive branch names:
- `feature/add-ising-model` - New features
- `fix/worker-timeout-bug` - Bug fixes
- `docs/update-api-docs` - Documentation
- `refactor/simplify-store` - Code refactoring

### Workflow

1. **Create a branch** from `main`:
   ```bash
   git checkout main
   git pull upstream main
   git checkout -b feature/your-feature
   ```

2. **Make your changes** with clear, atomic commits

3. **Test your changes**:
   ```bash
   # Run Go tests
   go test ./...
   
   # Test with Docker
   docker compose up --build --scale worker=3
   
   # Test dashboard
   cd dashboard && npm run lint && npm run build
   ```

4. **Push and create PR**:
   ```bash
   git push origin feature/your-feature
   ```

## Coding Standards

### Go Code

- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` for formatting
- Add comments for exported functions
- Handle errors explicitly (no silent failures)

```go
// Good
func (s *Store) GetTask(workerID string) (*Task, error) {
    s.mu.Lock()
    defer s.mu.Unlock()
    // ...
}

// Bad - missing mutex, no error handling
func (s *Store) GetTask(workerID string) *Task {
    return s.tasks[0]
}
```

### TypeScript/React Code

- Use TypeScript strict mode
- Prefer functional components with hooks
- Use descriptive variable names
- Follow ESLint rules (`npm run lint`)

```tsx
// Good
interface ParticleCloudProps {
    count: number;
    color?: string;
}

export function ParticleCloud({ count, color = '#8b5cf6' }: ParticleCloudProps) {
    // ...
}
```

### Commit Messages

Follow conventional commits:
```
feat: add Ising model simulation
fix: resolve worker timeout race condition
docs: update API documentation
refactor: simplify task queue logic
test: add unit tests for scheduler
```

## Submitting a Pull Request

1. **Ensure all tests pass** locally
2. **Update documentation** if needed
3. **Fill out the PR template** completely
4. **Link related issues** using keywords (`Fixes #123`)
5. **Request review** from maintainers

### PR Checklist

- [ ] Code follows project style guidelines
- [ ] Tests added/updated for changes
- [ ] Documentation updated if needed
- [ ] Commit messages follow conventions
- [ ] Branch is up-to-date with main

## Reporting Issues

### Bug Reports

Include:
- Clear description of the bug
- Steps to reproduce
- Expected vs actual behavior
- Environment (OS, Go version, Docker version)
- Relevant logs or screenshots

### Feature Requests

Include:
- Clear description of the feature
- Use case / motivation
- Proposed implementation (optional)

## Areas for Contribution

Looking for something to work on? Here are some ideas:

### Good First Issues
- Improve documentation
- Add unit tests
- Fix typos or clarify comments

### Features
- Add new simulation types (Ising Model, Lattice QCD)
- Implement worker health checks
- Add authentication to gRPC
- Create Kubernetes deployment manifests

### Infrastructure
- Add CI/CD pipeline (GitHub Actions)
- Improve Docker build caching
- Add integration tests

## Questions?

Feel free to open an issue with the `question` label or reach out to the maintainers.

---

Thank you for contributing to Democritus! 🎉
