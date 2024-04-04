main.wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm

clean:
	rm main.wasm
