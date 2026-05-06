# go install ccg-router@latest Problem

Install ccg-router with Go:

```bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@latest
ccg-router --version
```

## Checklist

- ccg-router requires a recent Go toolchain. The module can switch toolchains automatically when needed.
- Check that `$(go env GOPATH)/bin` is in your `PATH`.
- If `@latest` is delayed by the Go module proxy, install a specific version such as `@v0.1.2`.
- Use Homebrew or the shell installer if you do not want to manage Go binaries directly.

```bash
go install github.com/XZXY-AI/ccg-router/cmd/ccg-router@v0.1.2
```

## Related Pages

- Public page: <https://xzxy-ai.github.io/ccg-router/errors/go-install-latest/>
- Install page: <https://xzxy-ai.github.io/ccg-router/install/>
- Go module: <https://pkg.go.dev/github.com/XZXY-AI/ccg-router>
