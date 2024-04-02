package helper_lang

import "strconv"

func IsBlank(val string) bool {
	if len(val) == 0 {
		return true
	}
	return false
}

func IsAnyBlank(val ...string) bool {
	for _, s := range val {
		if IsBlank(s) {
			return true
		}
	}
	return false
}

// If 模拟的三元运算符
//
//	condition: 条件表达式
//	trueVal: 表达式为true时返回的值
//	falseVal: 表达式为false时返回的值
//
// return: 根据表达式的true/false，返回trueVal/falseVal
//
//	注意，由于返回的类型是interface{}，需要转换成trueVal/falseVal对应的类型
func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

// ToFloat64 字符串转换为float64
func ToFloat64(val string) (value float64, err error) {
	float, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}
	return float, nil
}

// ToUint64 字符串转换为uint64
func ToUint64(val string) (value uint64, err error) {
	u, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return u, nil
}

// IntToString int转为String
func IntToString(number int) (value string) {
	return strconv.Itoa(number)
}
