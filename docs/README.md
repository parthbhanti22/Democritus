# Democritus Documentation

This directory contains the technical documentation for the Democritus distributed Monte Carlo simulation engine, written in LaTeX.

## Documents

| File | Description |
|------|-------------|
| [design.tex](design.tex) | System architecture and design decisions |
| [api.tex](api.tex) | gRPC API specification and message formats |
| [implementation.tex](implementation.tex) | Implementation details for Worker and Scheduler |

## Building the Documentation

### Prerequisites

You need a LaTeX distribution installed:

- **Windows**: [MiKTeX](https://miktex.org/) or [TeX Live](https://tug.org/texlive/)
- **macOS**: [MacTeX](https://tug.org/mactex/)
- **Linux**: `sudo apt install texlive-full` (Debian/Ubuntu)

### Compile to PDF

#### Option 1: Compile Individual Files

```bash
cd docs
pdflatex design.tex
pdflatex api.tex
pdflatex implementation.tex
```

#### Option 2: Create a Master Document

Create a `main.tex` file that includes all sections:

```latex
\documentclass{article}
\usepackage{amsmath}
\usepackage{hyperref}

\title{Democritus: Technical Documentation}
\author{Democritus Contributors}
\date{\today}

\begin{document}
\maketitle
\tableofcontents
\newpage

\input{design.tex}
\newpage
\input{api.tex}
\newpage
\input{implementation.tex}

\end{document}
```

Then compile:
```bash
pdflatex main.tex
pdflatex main.tex  # Run twice for table of contents
```

### Using VS Code

1. Install the [LaTeX Workshop](https://marketplace.visualstudio.com/items?itemName=James-Yu.latex-workshop) extension
2. Open any `.tex` file
3. Press `Ctrl+Alt+B` to build
4. Press `Ctrl+Alt+V` to view PDF

### Using Overleaf

1. Create a new project on [Overleaf](https://www.overleaf.com/)
2. Upload all `.tex` files from this directory
3. Create a `main.tex` as shown above
4. Click "Recompile" to generate PDF

## Document Structure

### design.tex
- High-level architecture diagram
- Component responsibilities (Scheduler, Worker, Dashboard)
- Design decisions (Why gRPC? Why stateless workers?)
- Data flow and networking
- Scalability and security considerations

### api.tex
- gRPC service definition
- RPC methods (RegisterWorker, GetTask, SubmitResult)
- Message structures (TaskPayload, TaskResult)
- Reproducibility through seeded RNG

### implementation.tex
- Worker implementation (Strategy Pattern, Simulator interface)
- Random Walk physics and mathematics
- Scheduler implementation (Mutex, Reaper mechanism)
- Containerization and deployment
- Prometheus metrics

## Contributing to Documentation

When adding new documentation:

1. Follow the existing LaTeX style
2. Use `\section{}` for main topics
3. Use `\subsection{}` and `\subsubsection{}` for hierarchy
4. Include code examples in `\begin{verbatim}...\end{verbatim}`
5. Use math mode for equations: `$inline$` or `\[ display \]`

### Style Guidelines

```latex
% Good: Clear section hierarchy
\section{New Feature}
\subsection{Overview}
\subsection{Implementation}

% Good: Code examples
\begin{verbatim}
func Example() error {
    return nil
}
\end{verbatim}

% Good: Mathematical notation
The position at step $n$ is: $P_n = P_{n-1} + \vec{\delta}$
```

## License

This documentation is part of the Democritus project and is released under the MIT License.
