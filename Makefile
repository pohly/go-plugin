all: main plugin.so

main: main.go
	go build -o $@ .

plugin.so:
	go build -buildmode=plugin -o $@ ./plugin

clean:
	rm -f main plugin.so
