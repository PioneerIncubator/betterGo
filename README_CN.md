[English](https://github.com/PioneerIncubator/betterGo/blob/master/README.md) / 中文

# betterGo

betterGo实现了我认为Go所缺失的部分

## Real Generic

为用户提供了可以直接用在代码中的真正的`interface{}`。

在部署之前，仅需要使用`translator`生成确定类型的代码，这种方式并不会影响你的代码性能。

下面是已经实现的所有泛型函数：

* `enum.Reduce`
* `enum.Map`
* `enum.Delete`: Delete slice's first element for which fun returns a truthy value.
* `enum.Find`: Returns slice's first element for which fun returns a truthy value.

### 实现

使用Go AST来分析你使用泛型函数的代码，生成确定类型的函数并替换掉你原先的调用语句

### 实际上所做的事

![](https://pic1.zhimg.com/50/v2-dd2dc3bc72b058b85774ee804a521165_hd.webp)

### [视频介绍](https://www.bilibili.com/video/BV1oT4y1j7L2?from=search&seid=10341891677310558379)

### 背景

现在的Go语言不支持泛型（像C++中的template、Java中的interface）

目前，为实现泛型的需求，在Go语言中往往有如下几种方式[<sup>1</sup>](#refer-anchor-1)：

> 1. Interface （with method）
>    优点：无需三方库，代码干净而且通用。
>    缺点：需要一些额外的代码量，以及也许没那么夸张的运行时开销。
> 2. Use type assertions
>    优点：无需三方库，代码干净。
>    缺点：需要执行类型断言，接口转换的运行时开销，没有编译时类型检查。
> 3. Reflection
>    优点：干净
>    缺点：相当大的运行时开销，没有编译时类型检查。
> 4. Code generation
>    优点：非常干净的代码(取决工具)，编译时类型检查（有些工具甚至允许编写针对通用代码模板的测试），没有运行时开销。
>    缺点：构建需要第三方工具，如果一个模板为不同的目标类型多次实例化，编译后二进制文件较大。

`betterGo`就是通过`code generation`来实现泛型

### 如何使用

如果你想使用betterGo来通过自动生成代码的方式实现泛型，可以看下面的例子：

在项目中包含了测试用例，例如，需要使用泛型的代码是`test/map/map.go`，如果想用`interface{}` 的函数就是`enum.Map` 这样子用。

如果想生成具体类型的函数，就运行这行命令：`go run main.go -w -f test/map/map.go`

然后你发现 `test/map/map.go` 改变了，`enum.Map` 变成了: `enum.MapOriginFn(origin, fn)`

然后你看项目目录下生成了： `utils/enum/map.go`，就是具体类型的函数

### 参与项目

如果想和我们一起完成项目的开发，可以直接看代码，找到`AST`相关的包，尝试理解相关函数的作用，很容易就可以理解这个项目以及代码了。

如果想从理论出发的话，可以简单看看这本书：https://github.com/chai2010/go-ast-book ，其实他也就是把`AST`包里的代码简单讲讲。

想参与具体开发可以参考项目接下来的[TODO List](https://github.com/PioneerIncubator/betterGo/issues/31)

### 技术思路

1. 导入需要操作的文件/目录

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

<div id="refer-anchor-1"></div><a href="https://www.zhihu.com/question/62991191/answer/342121627">[1] Go有什麽泛型的实现方法? - 达的回答 - 知乎</a>
