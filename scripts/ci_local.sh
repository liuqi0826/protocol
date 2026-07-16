#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

echo "=== Unit tests ==="
go test ./...

echo "=== Generate all languages ==="
(cd example && go run main.go)

echo "=== Vet/build generated Go ==="
go vet ./example/export/go/...
go build -o /dev/null ./example/export/go/...

echo "=== Go round-trip ==="
(cd example/tests && go run test_go.go)

echo "=== JavaScript round-trip ==="
(cd example && node test_javascript.js)

echo "=== All CI checks passed ==="
