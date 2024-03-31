package helper_lang

import (
	"fmt"
	"testing"
)

type demo struct {
	Name string
}

type TemplateSub interface {
	Name() string
}

type SubImpl struct {
}

func (s SubImpl) Name() string {
	panic("implement me")
}

func TestGetRealType(t *testing.T) {
	realType := GetRealType(demo{})
	t.Log(realType)
}

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
