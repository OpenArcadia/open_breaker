go build -o open_breaker -ldflags="-s -w" -tags=x11 cmd/main.go
flatpak-builder build-dir --force-clean open_breaker.flatpak.yaml 
flatpak build-export repo build-dir
flatpak build-bundle repo open_breaker.flatpak com.openarchadia.open_breaker
