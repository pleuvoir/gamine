
## 语言相关



### 字符串操作


- IsBlank 是否为空

```go

func TestIsBlank(t *testing.T) {
	t.Log(IsBlank(""))
	t.Log(IsBlank("pleuvoir"))
}
```

输出：

```
true
false
```


- IsAnyBlank 任意一个是否为空

```go

func TestIsAnyBlank(t *testing.T) {
    t.Log(IsAnyBlank("1",""))
}
```

输出：

```
true
```

- If 模拟的三元运算符

因为返回的是`any`类型因此需要强制类型转换

```go
func TestIf(t *testing.T) {
	var number int
	a := If(rand.Int() > 1, 12, 13)
	number = a.(int)
	t.Log(number)
}
```

- ToUint64

```go
func TestToUint64(t *testing.T) {
	val := "3"
	if value, err := ToUint64(val); err == nil {
		t.Log(value)
	} 
}
```

输出：

```
3
```

- ToFloat64 ToFloat64 字符串转换为float64

```go
func TestToFloat64(t *testing.T) {
	val := "3.14"
	if value, err := ToFloat64(val); err == nil {
		t.Log(value)
	}
}
```

输出：
```go
3.14
```

输出：12或者13

### 反射操作

- GetRealType 解指针获取真实类型

```go
type demo struct {
	Name string
}

func TestGetRealType(t *testing.T) {
	realType := GetRealType(demo{})
	t.Log(realType)
}
```

输出：

```
helper_lang.demo
```


- MakeInstance 反射构建实例

通过真实类型构建一个空的实例

```go
func TestMakeInstance(t *testing.T) {
	instance := MakeInstance(GetRealType(demo{Name: "pleuvoir"}))
	t.Logf(fmt.Sprintf("%T", instance))

	if _, ok := instance.(TemplateSub); !ok {
		t.Log("不是这个接口的实现")
	}

	makeInstance := MakeInstance(GetRealType(SubImpl{}))
	t.Logf(fmt.Sprintf("%T", makeInstance))
	if templateSub, ok := makeInstance.(TemplateSub); ok {
		t.Logf(fmt.Sprintf("%T", templateSub))
	}
}
```

输出：

```
*helper_lang.demo
不是这个接口的实现
*helper_lang.SubImpl
```