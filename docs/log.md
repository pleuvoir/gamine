
## log 日志组件

```yaml
log-engine:
  log-configs:
    default:
      level: info
      path: '/Users/pleuvoir/dev/space/git/gamine/test/test_data/'
      filename: default
      maxAge: 1440h
      rotationTime: 24h
    bak:
      level: info
      path: '/Users/pleuvoir/dev/space/git/gamine/test/test_data/'
      filename: bak
      maxAge: 1440h
      rotationTime: 24h
```

```go
//设置工作目录，工作目录下读取配置文件
gamine.SetWorkDir("/Users/pleuvoir/dev/space/git/gamine/test")
gamine.SetEnvName("dev")
gamine.InstallComponents(&Instance{})
GetDefault().Infoln("default work")
Get("bak").Infoln("bak work")
```

如果配置的是拆分文件，则可以这样获取：

```go
log.Get("bak").Infoln("work")
log.Get("default").Infoln("work")
```

`name`和配置文件中的键保持一致即可。