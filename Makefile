.PHONY: build-demo

bin:
	mkdir -p bin

build-demo: bin

	# CGO_ENABLED=0: is used to disable CGO, which is required for building static binaries.
	# go build: is the Go compiler command used to compile the source code.
	#   -trimpath: is used to remove absolute paths to the source code from the compiled binary.
	#   -pgo=auto: is used to enable Profile-Guided Optimization (PGO) for auto-profiling.
	#   -ldflags="-s -w": is used to pass linker flags to the Go linker.
	#     -s: is used to remove the symbol table from the compiled binary.
	#     -w: is used to remove the debug information from the compiled binary.
	# -o bin/demo: is used to specify the output file name.
	# ./cmd/demo: is the path to the main package.
	CGO_ENABLED=0 go build -trimpath -pgo=auto -ldflags="-s -w" -o bin/demo ./cmd/demo

