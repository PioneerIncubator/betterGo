English / [中文](https://github.com/PioneerIncubator/betterGo/blob/master/README_CN.md)

# betterGo

betterGo implement parts that I think Golang missed

## Real Generic

Provide the real interface{} to you so that you can use it in your code.
Before deployment, just use translator to generate specific type code, in which way will not affect your performance.

Here are all generic functions:
* `enum.Reduce`
* `enum.Map`
* `enum.Delete`: Delete slice's first element for which fun returns a truthy value.
* `enum.Find`: Returns slice's first element for which fun returns a truthy value.

### Implementation

Use go ast to analyse your code where using generic functions, generate specific function for your types and replace your original call expressions.

### What I actually do

![](https://pic1.zhimg.com/50/v2-dd2dc3bc72b058b85774ee804a521165_hd.webp)



I do this shit for you :P

### Background

The current Go doesn't have generics (like template in c++, interface in Java).

The following ways are often used in Go in order to implement generics:

1. Interface (with method)

   Pros: No third-party libraries required, clean and universal code.

   Cons: Requires some additional amount of code, and perhaps a less dramatic runtime overhead.

2. Use type assertions

   Pros: No third-party libraries required, clean and universal code.

   Cons: Requires execution of type assertions, runtime overhead for interface conversions, and no compile-time type checking.

3. Reflection

   Pros: Clean code.

   Cons: Considerable runtime overhead, and no compile-time type checking.

4. Code generation
   
Pros: Extremely clean code, compile-time type checking, no runtime overhead.
   
   Cons: Requires third-party libraries, larger compiled binaries.

`betterGo` is a generic implementation of `code generation`.

### Usage

If you want to use `betterGo` to implement generics by automatically generating code, have a look at the following example:

There are test cases in the project, for example, the code that needs to be generic is `test/map/map.go`, if you want to use the `interface{}` function, just `enum.Map`.

If you want to generate a function of a specific type, run this command: `go run main.go -w -f test/map/map.go`

Then you'll find that `test/map/map.go` has changed, `enum.Map` has become `enum.MapOriginFn(origin, fn)`.

After that you can see that the project directory generates: `utils/enum/map.go`, which is a function of a specific type.

### Contributing

If you want to work with us on this project, you can look directly at the code, find the package related to `AST` and try to understand the relevant function so that you can easily understand the project and the code.

If you want to start with theory, you can find some information on `AST` and study it briefly.

### Implementation method

1. Load the file/directory to be manipulated

2. Syntactic analysis by `AST`

   `AST` can analyze the nature of each statement, such as:

   - `GenDecl` (General Declaration): Includes import, constant declaration, variable declaration, type declaration.
   - `AssignStmt`(Assignment Statement): Includes assignment statements and short variable declarations (a := 1).
   - `FuncDecl`(Function Declaration): 
   - `TypeAssertExpr`(Type Assertion Expression)
   - `CallExpr`( Call Expression)

3. When a statement containing the value/type of a variable  (`AssignStmt`, `FuncDecl`) is analyzed, the value and type of the variable are recorded and a mapping between them is established so that the type of the variable can be obtained from the variable name in subsequent sessions.

4. When a function call expression (`CallExpr`) is found, it is checked whether the function is provided by us, and if it is, the code that deals specifically with that type is generated from the type corresponding to the argument name recorded in the previous step and stored in the specified path (if code of the same type has already been generated before, it is not repeated).

5. Replaces the original function call expression in the original code with a new function call expression that calls the newly generated function from the previous step, and updates the import package.

