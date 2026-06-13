# Rag-From-Scratch

A simple Retrieval-Augmented Generation (RAG) pipeline built from scratch in Go — no LangChain, no LlamaIndex, no magic abstractions. Just Go, a vector database, and an LLM API.

The goal is to understand every layer of a RAG system by building each piece manually: document loading, text chunking, semantic retrieval, prompt augmentation, and response generation.

\---

## Why from scratch?

Frameworks like LangChain are great for shipping fast, but they hide what's actually happening. This project is about understanding the internals — writing the chunker, wiring the embedding API, talking to the vector DB directly, and constructing prompts by hand. Every component you see here is intentional.

\---

## How it works

```
Indexing (runs once)
─────────────────────────────────────────────────────────
  Load document → Split into chunks → Generate embeddings
  → Store vectors + metadata in Qdrant

Querying (runs on every user question)
─────────────────────────────────────────────────────────
  Embed the query → Search Qdrant for top-k similar chunks
  → Build a prompt with retrieved context
  → Send to LLM → Stream the answer
```

\---

## Stack

|Layer|Choice|Why|
|-|-|-|
|Language|Go|Fast, explicit, great for understanding what's happening|
|Embeddings|OpenAI `text-embedding-3-small`|Simple REST API, good quality|
|Vector DB|Qdrant (local via Docker)|Free, open source, official Go client|
|LLM|Gemini|Strong reasoning, good context handling|
|File parsing|Standard library + pdf lib|No framework needed|

\---

## Project structure

```
rag-from-scratch/
├── cmd/
│   ├── index/        # CLI to ingest and index a document
│   └── query/        # CLI to ask a question
├── internal/
│   ├── chunker/      # Text splitting logic
│   ├── embedder/     # Calls embedding API
│   ├── store/        # Qdrant client wrapper
│   ├── retriever/    # Similarity search
│   └── generator/    # Prompt builder + LLM call
├── docs/             # Sample documents to test with
├── go.mod
└── README.md
```

\---

## Getting started

### Prerequisites

* Go 1.21+
* Docker (for Qdrant)
* An OpenAI API key (embeddings)
* An Anthropic API key (generation)

### Run Qdrant locally

```bash
docker run -p 6333:6333 qdrant/qdrant
```

### Set environment variables

```bash
export OPENAI\_API\_KEY=your\_key\_here
export GEMINI\_API\_KEY=your\_key\_here
```

### Index a document

```bash
go run cmd/index/main.go --file docs/sample.txt
```

### Ask a question

```bash
go run cmd/query/main.go --question "What is this document about?"
```

\---

## What I built manually

* **Chunker** — fixed-size splitting with configurable overlap, no library
* **Embedder** — raw HTTP calls to the embedding API, no SDK
* **Vector store client** — direct REST calls to Qdrant
* **Retriever** — cosine similarity ranking, top-k selection
* **Prompt builder** — injects retrieved chunks into a structured prompt
* **LLM caller** — streams response from Claude API

\---

## Concepts covered

* Why RAG exists (private data, up-to-date info, reducing hallucination)
* How embeddings represent semantic meaning as vectors
* Why chunking strategy matters for retrieval quality
* How vector databases index and search high-dimensional space
* The difference between indexing (offline) and retrieval (real-time)
* How augmented prompts give the LLM grounded context to answer from

\---

## Roadmap

* \[x] Basic end-to-end pipeline
* \[ ] PDF support
* \[ ] Hybrid search (BM25 + vector)
* \[ ] Re-ranking pass
* \[ ] Metadata filtering
* \[ ] Simple web UI

\---

## Learning resources

* [CampusX RAG playlist](https://www.youtube.com/@campusx-official) — conceptual foundation
* [Qdrant docs](https://qdrant.tech/documentation/)
* [OpenAI Embeddings guide](https://platform.openai.com/docs/guides/embeddings)

