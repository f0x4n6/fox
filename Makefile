GO=go
BIN=/usr/local/bin

.PHONY: all clean

all: install

install: build
	chmod +x bin/fox
	cp -v bin/fox ${BIN}/fox

build:
	mkdir -p bin
	${GO} build -o bin/fox main.go

clean:
	rm -rf bin
	${GO} clean
