

## 基础

组件运行方式，设置工作目录，传入对应的结构体，完成配置文件到结构体的转换。

```go
gamine.SetWorkDir(".")
gamine.SetEnvName("dev")
gamine.InstallComponents(log.&Instance{})
```

框架会以`workDir`目录为基础，依次寻找以下路径的配置文件，如果都没有找到则报错。

- .
- ./bin
- ../bin
- ./configs
- ../configs


寻找`gamine.{envName}.yml` 文件，其中`envName`的可选值为：`dev prod`，当不设置`envName`时默认为`dev`。
所有的组件都是一个单例对象，初始化完成后可在程序任意位置获取。


以上，对组件的了解可以从 [hello](https://github.com/pleuvoir/gamine/blob/main/docs/hello.md) 开始。