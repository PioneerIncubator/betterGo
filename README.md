# betterGo
Better Go implement parts that I think Golang missed

#### Real Generic
Provide the real interface{} to you so that you can use it in your code.
Before deployment, just use translator to generate specify type code, in which way will not affect your performance.

Here are all generic functions:
* enum.Reduce
* enum.Map


#### Implementation
Use go ast to analyse your code where using generic functions, generate specify function for your types and replace your original call.

#### What I actually do
![](https://pic1.zhimg.com/50/v2-dd2dc3bc72b058b85774ee804a521165_hd.webp)



I do this shit for you :P

# betterGo

### Join Us



<div align=center><img width="350" height="450" src="https://user-images.githubusercontent.com/51999056/84882050-43bf2b00-b0c1-11ea-9142-f1a7ec774dbc.png"/></div>

### Background

现在的Go语言不支持泛型（像C++中的template、Java中的interface）

目前，为实现泛型的需求，在Go语言中往往有如下几种方式[<sup>1</sup>](#refer-anchor-1)：

> 1. Copy & paste
>
>    尽管这是一个看起来很笨的方法，但是结合实际应用情况，也许大多数情况下你只需要一两个类型就足够了，太早想到『优化』，带来的可能和你预期的有所出入。
>
>    优点：无需三方库，利用一些 IDE 或者编辑器插件，完成功能迅速。
>
>    缺点：代码有些臃肿，不符合 [Dry](https://link.zhihu.com/?target=https%3A//en.wikipedia.org/wiki/Don%27t_repeat_yourself) 编程原则。
>
> 2. Interface （with method）
>
>    优点：无需三方库，代码干净而且通用。
>
>    缺点：需要一些额外的代码量，以及也许没那么夸张的运行时开销。
>
> 3. Use type assertions
>
>    优点：无需三方库，代码干净。
>
>    缺点：需要执行类型断言，接口转换的运行时开销，没有编译时类型检查。
>
> 4. Reflection
>
>    优点：干净
>
>    缺点：相当大的运行时开销，没有编译时类型检查。
>
> 5. Code generation
>
>    优点：非常干净的代码(取决工具)，编译时类型检查（有些工具甚至允许编写针对通用代码模板的测试），没有运行时开销。
>
>    缺点：构建需要第三方工具，如果一个模板为不同的目标类型多次实例化，编译后二进制文件较大。

betterGo就是通过code generation来实现泛型

### 技术思路

1. 导入需要操作的文件，可以是文件/目录

2. 通过AST进行语法分析

   AST能分析出每条语句的性质，如：

   - `GenDecl` (一般声明)：包括import、常量声明、变量声明、类型声明
   - `AssignStmt`(赋值语句)：包括赋值语句和短的变量声明(a := 1)
   - `FuncDecl`(函数声明)
   - `TypeAssertExpr`(类型断言)
   - `CallExpr`(函数调用语句)

3. 当分析到包含变量的值/类型的语句时(`AssignStmt`、`FuncDecl`)会对变量的值和类型进行记录，并建立二者之间的映射关系，以便于在后续环节中能够通过变量名获取变量的类型

4. 当发现函数调用语句(`CallExpr`)时，会检查该函数是否为我们提供的函数，如果是，则通过上一步中记录的参数名对应的类型生成专门处理该类型的一份代码，并存储到指定路径下（如果之前已经生成过相同类型的代码则不重复生成）

5. 将原代码中的原来的函数调用语句替换成新的函数调用语句，使其调用上一步中新生成的函数，并更新import的包

### Reference

<div id="refer-anchor-1"></div>- [1] [Go有什麽泛型的实现方法? - 达的回答 - 知乎](https://www.zhihu.com/question/62991191/answer/342121627)
