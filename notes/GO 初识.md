关于 go 的介绍，这里就不在详细讲解，百度查阅即可。 go 的命名方式及规范和其他部分语言相似，也不做过多的介绍。我们直接来了解 go 的一些其他东西。  

在学习 go 中，我们都知道 go 以其简单、高效而受到开发者的青睐，那么我就来详细了解一下 go 是如何简单的呢？下面列举了 Go 代码中会使用到的 25 个关键字或保留字：  
```
break       default     func    interface   select
case        defer       go      map         struct
chan        else        goto    package     switch
const       fallthrough if      range       type
continue    for         import  return      var
```
##### 关键字说明：
fallthrough：Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。  

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符，其中包含了基本类型的名称和一些基本的内置函数，如下：  
```
append  bool    byte    cap         close   complex     complex64   complex128  uint16
copy    false   float32 float64     imag    int         int8        int16       uint32
int32   int64   iota    len         make    new         nil         panic       uint64
print   println real    recover     string  true        uint        uint8       uintptr
```
程序一般由关键字、常量、变量、运算符、类型和函数组成。  
程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。  
程序中可能会使用到这些标点符号：.、,、;、: 和 …。  

如果记住以上的关键字和预定义标识符，那么你的 go 基本入门。但是，没有必要死记硬背，后面我们都会涉及到，这里只需要有个印象就好。
