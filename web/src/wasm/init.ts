export let module: WebAssembly.Module;
export let instance: WebAssembly.Instance;

export const go = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then( async (result) => {
    module = result.module;
    instance = result.instance;
    go.run(instance)
    console.log("WASM initialized")
});

