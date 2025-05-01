mkdir -p windows-build
CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -o windows-build/open_breaker.exe -ldflags "-s -w" cmd/main.go 
mkdir -p windows-build/assets
cp -r assets/* windows-build/assets/
cp raylib.dll windows-build/