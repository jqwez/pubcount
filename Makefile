BIN=bin
MAIN=main.go
TARGET=pubcount
TWIND=third_party/twind
TWIND_OUT=./web/static/css/styles.css

all: build tests

build-run: build run

dev: build
	@./bin/pubcount && $(TWIND) --watch

run:
	@./bin/pubcount

build: header tailwind gobin
	@echo " 🎉 🎉 🎉  Build Complete 🎉 🎉 🎉"

.PHONY: build

header:
	@echo "~~~~~Building PubCount~~~~~"
	@echo "....."

.PHONY: header

gobin: 
	@echo "🚀 🚀 🚀 Compiling Go Executable 🚀 🚀 🚀"
	@echo "....."
	@go build -o bin/$(TARGET) $(MAIN)
	@echo "PubCount Ready in: bin/$(TARGET)"

.PHONY:	gobin

tailwind:
	@echo "✒️ ✒️ ✒️  Compiling TailwindCSS ✒️ ✒️ ✒️"
	@echo "....."
	@ $(TWIND) -o $(TWIND_OUT) --minify

.PHONY: tailwind

tests:
	@echo "Running go tests......"
	@go test ./...

.PHONY: tests

clean:
	@echo "🗑️ 🗑️ 🗑️   Cleaning 🗑️ 🗑️ 🗑️" 
	@echo "Cleaning $(BIN)"
	@rm $(BIN)/*  
	@echo "Cleaning Styles"
	@rm $(TWIND_OUT)
	@echo "🪅 🪅 🪅  Done!🪅 🪅 🪅"

.PHONY: clean
