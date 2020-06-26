all: clean mac linux win
	echo "Building all binaries"
clean:
	rm -f main_linux main_win main_mac
mac: 
	env GOOS='darwin' GOARCH='amd64' go build -o main_mac
linux:
	env GOOS='linux' GOARCH='amd64' go build -o main_linux
win:
	echo "Warning not fully tested"
	env GOOS='windows' GOARCH='amd64' go build -o main_win

