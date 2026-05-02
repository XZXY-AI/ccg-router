# Security Policy

## Scope

In scope:

- `ccg-router` daemon
- Local HTTP API
- Registry schema and signature verification
- Local ledger storage behavior

Out of scope:

- Third-party registries published by others
- User-provided upstream endpoints
- Misconfigured local environments

## Reporting

Use GitHub Security Advisories for `XZXY-AI/ccg-router`.

Expected timeline:

- Acknowledgement within 72 hours
- Fix or mitigation target within 90 days

## Verifying Release Archives

Release archives and `checksums.txt` are signed with cosign keyless signing from GitHub Actions.

```bash
cosign verify-blob \
  --certificate-identity-regexp 'https://github.com/XZXY-AI/ccg-router/.*' \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  --signature <archive>.sig <archive>
```
