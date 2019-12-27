# betterGo
Better Go implement parts that I think Golang missed

#### Real Generic
Provide the real interface{} to you so that you can use it in your code.
Before deployment, just use translator to generate specify type code, in which way will not affect your performance.

Here are all generic functions:
* enum.Reduce


#### Implementation
Use go ast to analyse your code where using generic functions, generate specify function for your types and replace your original call.

#### What I actually do
![](https://pic1.zhimg.com/50/v2-dd2dc3bc72b058b85774ee804a521165_hd.webp)
I do this shit for you :P


