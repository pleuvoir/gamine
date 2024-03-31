

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

- GetWdQuiet 安静的获取当前目录

**注意：**   这个获取的实际上是`helper_os.go`文件所在的目录。因此，该函数的示例意义大于实际使用。

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

**注意：**   这个获取的实际上是`helper_os.go`文件所在文件的的根目录。因此，该函数的示例意义大于实际使用。


```go
func TestRootPath(t *testing.T) {
	t.Log(RootPath())
}
```


输出：
```
/Users/pleuvoir/dev/space/git/gamine/helper <nil>
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