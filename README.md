# CopyPasta

A small web app for sharing text snippets, written in Go using only the standard library.

> **Status:** early development. The routes below exist but return placeholder text; there is no storage or HTML rendering yet.

## Requirements

- Go 1.27.1 or later

## Running

```sh
go run .
```

The server listens on [http://localhost:4000](http://localhost:4000).

## Routes

| Path              | Description                              |
| ----------------- | ---------------------------------------- |
| `/`               | Home page (exact match only)             |
| `/snippet/view`   | Display a specific snippet               |
| `/snippet/create` | Display a form for creating a new snippet |

## Project layout

```
.
├── main.go          # HTTP server, router and handlers
├── templates/       # HTML templates (index.html is a stub for now)
└── go.mod
```
