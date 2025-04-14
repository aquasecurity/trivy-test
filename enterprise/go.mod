module github.com/aquasecurity/trivy-enterprise

go 1.24.2

require github.com/aquasecurity/trivy v0.61.0

require (
	github.com/aquasecurity/trivy-db v0.0.0-20250227071930-8bd8a9b89e2d // indirect
	github.com/google/go-containerregistry v0.20.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/package-url/packageurl-go v0.1.3 // indirect
	github.com/samber/lo v1.49.1 // indirect
	golang.org/x/text v0.23.0 // indirect
	golang.org/x/xerrors v0.0.0-20240716161551-93cc26a95ae9 // indirect
)

replace github.com/aquasecurity/trivy => ../trivy