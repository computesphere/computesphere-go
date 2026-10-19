#!/usr/bin/env bash
# The Go example in README.md must build against this checkout of the SDK.
# Extracts the README's ```go block, builds it with `go vet` and `go build` in
# a throwaway module that replaces the SDK with this directory, and checks it
# is gofmt-formatted. Run from the repository root: scripts/check-readme-example.sh
set -euo pipefail

root=$(pwd)
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT

blocks=$(grep -c '^```go$' README.md || true)
if [ "$blocks" -ne 1 ]; then
  echo "README.md: expected one go block, found $blocks" >&2
  exit 1
fi
awk '/^```go$/{f=1;next} /^```$/{f=0} f' README.md > "$work/main.go"

cat > "$work/go.mod" <<MOD
module readme

go $(awk '/^go /{print $2; exit}' "$root/go.mod")

require github.com/computesphere/computesphere-go v0.0.0

replace github.com/computesphere/computesphere-go => $root
MOD

cd "$work"
unformatted=$(gofmt -l .)
if [ -n "$unformatted" ]; then
  echo "README.md go example is not gofmt-formatted" >&2
  exit 1
fi
GOFLAGS=-mod=mod go mod tidy
go vet ./...
go build ./...
echo "README.md go example builds against this checkout"
