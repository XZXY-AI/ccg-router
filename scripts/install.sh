#!/usr/bin/env bash
set -euo pipefail

REPO="ccg-labs/ccg-router"
VERSION="${CCG_VERSION:-latest}"
DEST="${CCG_DEST:-$HOME/.local/bin}"

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$os" in
  linux|darwin) ;;
  *) echo "unsupported os: $os" >&2; exit 1 ;;
esac

arch="$(uname -m)"
case "$arch" in
  x86_64|amd64) arch=amd64 ;;
  arm64|aarch64) arch=arm64 ;;
  *) echo "unsupported arch: $arch" >&2; exit 1 ;;
esac

if [ "$VERSION" = "latest" ]; then
  VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep -oE '"tag_name":[[:space:]]*"v[^"]+"' \
    | head -1 \
    | sed -E 's/.*"v([^"]+)".*/\1/')"
fi

archive="ccg-router_${VERSION}_${os}_${arch}.tar.gz"
url="https://github.com/${REPO}/releases/download/v${VERSION}/${archive}"
sums="https://github.com/${REPO}/releases/download/v${VERSION}/checksums.txt"

tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "$url" -o "$tmp/$archive"
curl -fsSL "$sums" -o "$tmp/checksums.txt"

want="$(grep -E " ${archive}$" "$tmp/checksums.txt" | awk '{print $1}')"
got="$(shasum -a 256 "$tmp/$archive" | awk '{print $1}')"
if [ "$want" != "$got" ]; then
  echo "checksum mismatch: want=$want got=$got" >&2
  exit 1
fi

mkdir -p "$DEST"
tar -xzf "$tmp/$archive" -C "$tmp"
install -m 0755 "$tmp/ccg-router" "$DEST/ccg-router"
echo "installed: $DEST/ccg-router"
echo "add $DEST to PATH if it is not already."
