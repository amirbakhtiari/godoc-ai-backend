# GoDoc AI

GoDoc AI is a production-oriented Retrieval-Augmented Generation (RAG) chatbot for technical documentation.

The goal of this project is to build an AI-powered documentation assistant that can retrieve relevant information from a large documentation base and generate accurate, grounded answers with source references.

## Architecture

```text
Source Files
    ↓
Document Loader
    ↓
Markdown Parser
    ↓
Chunker
    ↓
PostgreSQL
    ↓
Embedding
    ↓
pgvector
    ↓
Retrieval
    ↓
Reranker
    ↓
Context Builder
    ↓
LLM
    ↓
Answer + Sources
```

## Current Progress

- [x] Project structure
- [x] PostgreSQL + pgvector
- [x] Document domain model
- [x] Markdown document loader
- [x] Document scanner
- [x] PostgreSQL repository foundation
- [x] Markdown section parser
- [x] Markdown-aware chunking
- [x] Chunk metadata
- [ ] Production-grade chunking
- [ ] Idempotent document ingestion
- [ ] Embedding generation
- [ ] Vector storage
- [ ] Vector retrieval
- [ ] Keyword search
- [ ] Hybrid retrieval
- [ ] Reranking
- [ ] Context building
- [ ] LLM generation
- [ ] Source citations
- [ ] Chat API
- [ ] Evaluation
- [ ] Production deployment

## Tech Stack

### Backend

- Go
- PostgreSQL
- pgvector
- pgx

### AI / RAG

- Embeddings
- Vector Search
- Keyword Search
- Hybrid Retrieval
- Reranking
- LLM

### Infrastructure

- Docker
- Docker Compose

## Project Structure

```text
godoc-ai/
├── backend/
│   ├── cmd/
│   │   └── api/
│   ├── internal/
│   │   ├── domain/
│   │   ├── application/
│   │   ├── infrastructure/
│   │   └── interfaces/
│   ├── migrations/
│   ├── docs/
│   ├── go.mod
│   └── go.sum
├── frontend/
├── docker-compose.yml
└── README.md
```

## Development

Start PostgreSQL with pgvector:

```bash
docker compose up -d
```

Run the backend:

```bash
cd backend
go run ./cmd/api
```

## Status

This project is currently under active development.

The current focus is building a reliable document ingestion and chunking pipeline before implementing embeddings and retrieval.