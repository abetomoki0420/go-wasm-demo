# How to Compile Go Code to WebAssembly (Wasm)

This guide explains how to compile Go code to WebAssembly (Wasm) format using the specified script.

## Prerequisites
Before proceeding with the compilation process, make sure you have the following requirements in place:
- Go programming language installed on your system.
- The necessary dependencies for your Go code.

## Compilation Command
To compile your Go code into WebAssembly, run the following command in your terminal:

```shell
sh compile.sh
```

This command sets the `GOOS` and `GOARCH` environment variables to `js` and `wasm`, respectively, ensuring that the Go code is compiled specifically for the WebAssembly target. The resulting Wasm file will be named `main.wasm` and saved in the `../web` directory.

Adjust the command according to your project structure and requirements. Make sure to specify the correct file path for your Go source code (`./src/main.go` in this example).

Once the compilation process is complete, you will have the `main.wasm` file ready to use for web applications or other compatible environments.

Feel free to explore the repository for additional resources and examples related to WebAssembly and Go.