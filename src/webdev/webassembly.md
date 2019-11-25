# [WebAssembly · golang/go Wiki · GitHub](https://github.com/golang/go/wiki/WebAssembly)

# Links

# where is wasm_exec.js

* [go/wasm_exec.js at master · golang/go](https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js)

and how to load and run compiled Go wasm module:

* [Load and Run Go WebAssembly Module](https://siongui.github.io/2018/10/06/load-and-run-golang-wasm-code/)

```javascript
<script src="wasm_exec.js"></script>
<script>
      const go = new Go();
      let mod, inst;
      WebAssembly.instantiateStreaming(
              fetch("mymodule.wasm", {cache: 'no-cache'}), go.importObject).then((result) => {
              mod = result.module;
              inst = result.instance;
              run();
      });
      async function run() {
              await go.run(inst);
      };
</script>
```

