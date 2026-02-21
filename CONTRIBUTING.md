# AGATA Contributor Guidelines

Thank you for your interest in contributing to the AGATA project! We welcome contributions of all kinds: bug fixes, new features, documentation improvements, and more. Please take a moment to review this guide to make the contribution process smooth and transparent.

## Table of Contents
- [Code of Conduct](#code-of-conduct)
- [Reporting Bugs and Feature Requests](#reporting-bugs-and-feature-requests)
- [Development Process](#development-process)
- [Code Requirements](#code-requirements)
- [Licensing and CLA](#licensing-and-cla)
- [Setting Up Your Development Environment](#setting-up-your-development-environment)

## Code of Conduct
This project and everyone participating in it is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Reporting Bugs and Feature Requests
- Use [GitHub Issues](https://github.com/your-username/agata/issues) to report bugs and suggest features.
- Before creating a new issue, please check if a similar one already exists.
- For bugs, provide detailed steps to reproduce, expected vs. actual behavior, and environment details (OS, component versions).
- For feature requests, use the `enhancement` label.

## Development Process
1. **Fork** the repository on GitHub.
2. **Clone** your fork locally.
3. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/short-description
   
  Branch names should reflect the change (e.g., feature/add-mongo-adapter or fix/scheduler-panic).
4. Make your changes, adhering to the code requirements (see below).
5. Add tests if applicable. Ensure all existing tests pass.
6. Commit with a clear message. We encourage Conventional Commits:

feat: add MongoDB adapter
fix: prevent scheduler panic on empty schedule
docs: update README
7. Push your branch to your fork.
8. Open a Pull Request (PR) against the main repository. Describe your changes in detail.
9. Participate in the discussion – respond to reviewer comments and make adjustments as needed.

## Code Requirements

  - Go:
	- Code must be formatted with gofmt.
	- Follow Effective Go guidelines.
	- Before submitting, run go mod tidy and go test ./....
	- Use golangci-lint for style checks (configuration in the root).

  - Frontend (React/TypeScript):
	- Code must follow ESLint and Prettier settings in the web/ directory.
	- Use functional components and hooks.

  - Python (analytics microservice):
	- Follow PEP8; format code with black and isort.
	- Use type hints.
	- Add docstrings for public functions.

All changes should be covered by tests (unit, integration). Use Go's testing package, Python's pytest, or Jest/React Testing Library for frontend.

## Licensing and CLA

The AGATA project uses different licenses for its components:
- Go core and frontend – Apache License 2.0.
- Python Analytics API – GNU Affero General Public License v3.

Before we can accept your contribution, you must sign a Contributor License Agreement (CLA). This ensures that we have the necessary rights to use your code in the project, including for commercial versions of AGATA.

The full CLA text is in AGATA-CLA.md. To sign it:

1. When you open a Pull Request, our bot (CLA Assistant) will leave a comment with a link.
2. Follow the link, authenticate via GitHub, and click to sign.
3. Once signed, the PR status will update, and the CLA check will pass.

You only need to sign the CLA once; it applies to all your future contributions.
## Setting Up Your Development Environment

    Clone the repository
      git clone https://github.com/your-username/agata.git
      cd agata

    Start dependencies (PostgreSQL, ClickHouse, Redis)
      docker-compose up -d postgres clickhouse redis

    Download Go dependencies
      go mod download

    Run the Go server (with air for hot-reload)
      make run-go

    In a separate terminal, run the Python service
      make run-python

    In another terminal, run the frontend
      make run-web

For local development, use Docker Compose to spin up required services:

See the Makefile and README.md for more commands and details.
Thank you again for contributing! We appreciate your help in making AGATA better.
