package helper_lang

import (
	"fmt"
	"reflect"
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

func TestGetRealType1(t *testing.T) {
	type args struct {
		any interface{}
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRealType(tt.args.any); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRealType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeInstance1(t *testing.T) {
	type args struct {
		p reflect.Type
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeInstance(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}
