
## sqlite

使用`gorm`，需要准备如下的配置文件。路径可以是相对或者绝对路径。

```yaml
sqlite:
  path: '../../test/test_data/sqlite-yml.db'
```

### 使用配置文件从组件加载

```go
type account struct {
    gorm.Model
    Name string
}

func TestInstallComponents(t *testing.T) {
    dbPath, _ := filepath.Abs("../../test/test_data/sqlite-yml.db")
    t.Log(dbPath)

    defer os.Remove(dbPath)
    
    gamine.SetWorkDir("../../test/")
    gamine.InstallComponents(&Instance{})
    err := Get().AutoMigrate(&account{})
    assert.NoError(t, err)
    acc := account{Name: "gamine"}
    Get().Create(&acc)
    acc2 := account{}
    Get().First(&acc2, "name=?", acc.Name)
    assert.Equal(t, acc.Name, acc2.Name)
}
```


### 直接运行组件

```go
func TestInstance_Run(t *testing.T) {
	dbPath, _ := filepath.Abs("../../test/test_data/sqlite-test.db")
	t.Log(dbPath)

	defer os.Remove(dbPath)
	gamine.RunComponents(&Instance{Path: dbPath})
	err := Get().AutoMigrate(&account{})
	assert.NoError(t, err)
	acc := account{Name: "gamine"}
	Get().Create(&acc)
	acc2 := account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)
}
```



