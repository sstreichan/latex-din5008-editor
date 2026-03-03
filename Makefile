.PHONY: build clean serve

build:
	GOOS=js GOARCH=wasm go build -o docs/main.wasm .
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" docs/

clean:
	rm -f docs/main.wasm docs/wasm_exec.js

serve:
	cd docs && python3 -m http.server 8080
