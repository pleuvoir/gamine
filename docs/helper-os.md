

## 系统相关

### 环境变量操作

- GetEnvOrDefault 获取环境变量，获取失败则返回默认值
- GetEnv 获取环境变量，获取失败返回空字符串
- SetEnvQuiet 安静的设置环境变量，不抛出异常

```go
func TestEnv(t *testing.T) {
	SetEnvQuiet("key2", "pleuvoir")
	env := GetEnv("key2")
	t.Logf(env)

	envOrDefault := GetEnvOrDefault("key", "default-key-value")
	t.Log(envOrDefault)
}
```

输出：

```
os_test.go:24: pleuvoir
os_test.go:27: default-key-value
```


### 系统目录相关

- NormalizePath 将路径转换为绝对路径，如果路径是相对路径的话


```go
func TestNormalizePath(t *testing.T) {
	t.Log(NormalizePath("../"))
}
```


输出：


```
/Users/pleuvoir/dev/space/git/gamine/helper <nil>
```

- ChdirQuietly 安静的切换目录

**注意：**   切换工作目录后，后面获取路径都是以切换后的为基础，比如相对路径都会从当前切换的路径进行计算。。

```go
func TestChdirQuietly(t *testing.T) {
    t.Log(GetWdQuiet())
    ChdirQuietly("../helper_lang") //切换到上一级
    t.Log(GetWdQuiet())
    ChdirQuietly("../") //切换到上一级
    t.Log(GetWdQuiet())
    t.Log(RootPath())
}
```

输出：


```
/Users/pleuvoir/dev/space/git/gamine/helper/helper_os
/Users/pleuvoir/dev/space/git/gamine/helper/helper_lang
/Users/pleuvoir/dev/space/git/gamine/helper
/Users/pleuvoir/dev/space/git/gamine <nil>
```


- GetWdQuiet 安静的获取当前目录

**注意：**   这个获取的实际上是`helper_os.go`文件所在的目录。在不切换工作路径的情况下，该函数的示例意义大于实际使用。

```go

func TestGetWdQuiet(t *testing.T) {
	dir := GetWdQuiet()
	t.Log(dir)
}

```

输出：
```
/Users/pleuvoir/dev/space/git/gamine/helper/helper_os
```

- CurrentExecutePath 获取当前的执行文件所在的目录

```go
func TestCurrentPath(t *testing.T) {
	currentPath, err := CurrentExecutePath()
	if err != nil {
		panic(err)
	}
	t.Log(currentPath)
}
```

输出：
```
/private/var/folders/b_/0j5tbqk55h9bstjczsrnbt000000gn/T/GoLand
```


- RootPath 获取项目根路径

**注意：**   这个获取的实际上是`helper_os.go`文件所在文件的的根目录。在不切换工作路径的情况下，该函数的示例意义大于实际使用。


```go
func TestRootPath(t *testing.T) {
	t.Log(RootPath())
}
```


输出：
```
/Users/pleuvoir/dev/space/git/gamine/helper <nil>
```


- Abs 获取绝对路径

获取失败时，返回原来输入的相对路径

**注意：**   这个获取的实际上是以`helper_os.go`文件计算推导出的路径。在不切换工作路径的情况下，该函数的示例意义大于实际使用。


```go
func TestAbs(t *testing.T) {
	t.Log(Abs("../test"))
}
```

输出：

```
/Users/pleuvoir/dev/space/git/gamine/helper/test
```



### 文件相关

- FileExists 判断文件是否存在

```go
func TestFileExists(t *testing.T) {
	filePath := path.Join(GetWdQuiet(), "os.go")
	t.Log(filePath)
	exists := FileExists(filePath)
	t.Log(exists)
}
```

输出：
```
/Users/pleuvoir/dev/space/git/gamine/helper/helper_os/os.go
true
```


- FolderExists 文件夹是否存在，或者是不是文件夹

```go
func TestFolderExists(t *testing.T) {
    t.Log(FolderExists("../../"))
}
```

输出：

```
true
```


- GetHomeDir 获取系统home路径

```go
func TestGetHomeDir(t *testing.T) {
	t.Log(GetHomeDir())
}

```

输出：

```
/Users/pleuvoir
```


- CloseQuietly 安静的调用Close()
```go
type closeImpl struct {
}

func (c *closeImpl) Close() error {
	panic("implement me")
}

func TestCloseQuietly(t *testing.T) {
	CloseQuietly(&closeImpl{})
}
```

输出：

```
panic: implement me
```

### 系统操作相关

- WaitQuit 阻塞进程，等待退出信号

```go
func TestWaitQuit(t *testing.T) {
    WaitQuit()
}
```

