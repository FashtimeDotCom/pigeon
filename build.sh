# Linux64位
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -x -o edm64_linux main.go

# Linux32位
GOOS=linux GOARCH=386 go build -ldflags "-s -w" -x -o edm32_linux main.go

# Windows64位
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -x -o edm64.exe main.go

# Windows32位
GOOS=windows GOARCH=386 go build -ldflags "-s -w" -x -o edm32.exe main.go

# Mac
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -x -o edm_mac main.go