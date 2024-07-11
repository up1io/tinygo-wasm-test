# Directories
BUILD_DIR := .build/js
TINYGOROOT := $(shell tinygo env TINYGOROOT)

.PHONY: clean-build build-wasm run

build-wasm: | $(BUILD_DIR)
	#@tinygo build -o $(BUILD_DIR)/main.wasm -target wasm -no-debug ./cmd/editor/
	@tinygo build -o $(BUILD_DIR)/main.wasm -target wasm ./cmd/wasm/
	@cp $(TINYGOROOT)/targets/wasm_exec.js $(BUILD_DIR)/wasm_exec_tiny.js
	@uglifyjs $(BUILD_DIR)/wasm_exec_tiny.js -o $(BUILD_DIR)/wasm_exec_tiny.js --compress --mangle

build-html: | $(BUILD_DIR)
	@cp ./static/index.html .build

$(BUILD_DIR):
	@mkdir -p $@

dev: build-html build-wasm run

run:
	@go run ./cmd/server/main.go
