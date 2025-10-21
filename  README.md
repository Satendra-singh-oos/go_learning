# Golang 


## **Advantages of Golang (Go)**

1. **Build Time** – Go compiles very fast. In CI/CD (Continuous Integration/Deployment) pipelines, this means you can build, test, and deploy your application quickly without long waits.

2. **Fast Startup** – Go programs start almost instantly because the code is compiled directly to machine code. This is great when auto-scaling, for example, creating new app instances during traffic spikes.

3. **Performance & Efficiency** – Go apps generally use less CPU and memory. Thanks to **goroutines** (Go’s lightweight threads), it can handle many tasks at the same time without hogging resources.

4. **Concurrency Model** – Go can run multiple tasks at once by fully using all CPU cores. This is like having multiple hands working together instead of just one.

5. **Static Typing & Compilation** – Go checks your code for errors before running it, which reduces bugs and makes your app more reliable.

6. **Small Executable Size** – When you compile a Go app, the final executable file is small, so it’s easy to store, transfer, and deploy.




## ** Some GO CMD ** 

-  go run main.go
-  go build main.go
- go fmt ./...




## **Variable & Constant Mapping Against JS**

| JavaScript                 | Go Syntax                         | Scope                                                      | Notes                                                                                             |
| -------------------------- | --------------------------------- | ---------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `let name = "John"`        | `name := "John"`                  | Block scope                                                | Only works **inside functions** in Go (short declaration). Type is inferred.                      |
| `const PI = 3.14`          | `const PI = 3.14`                 | Block scope                                                | Works inside and outside functions. Must be assigned at declaration.                              |
| `var age = 25`             | `var age = 25`                    | Function scope (in JS) / Package or function scope (in Go) | In Go, type is inferred here, but you can also declare with an explicit type: `var age int = 25`. |
| `var city; city = "Delhi"` | `var city string; city = "Delhi"` | Function scope (in JS) / Package or function scope (in Go) | Go requires you to define the type if not assigning immediately.                                  |
| `let active = true`        | `var active bool = true`          | Block scope                                                | Can also use `active := true` inside a function.                                                  |

---

💡 **Key Go Rules to Remember**:

1. **Unused variables cause compile errors** — must be used or removed.
2. **Short declaration (`:=`)** can’t be used outside functions.
3. **Type inference** works like JS `let` without a type, but Go still enforces static typing after compile time.
4. **Constants (`const`)** can only hold values known at compile time (no runtime calculation).

---

