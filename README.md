
# QuadHealth Core (Go) – CI/CD + DevSecOps

Serviço base em **Golang** com endpoint `/healthz`, acompanhado de **pipeline GitHub Actions** com *gates* de segurança: **gosec/semgrep (SAST)**, **govulncheck (SCA)**, **Trivy (FS+Imagem)**, **OWASP ZAP (DAST)**, **Gitleaks (segredos)** e **SBOM CycloneDX**.

## Rodar local
```bash
cd services/core-go
go mod tidy
go run ./...
# http://localhost:8080/healthz
```
