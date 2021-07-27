package tr

/*
这个模块是用于简单的数据转换
*/

import (
	"fmt"
	"strconv"
)

// Int64ToString 是int64 转 string
func Int64ToString(i int64)string{
	return strconv.FormatInt(i,10)
}

func StringToInt(i string)int{
	in,_ := strconv.Atoi(i)
	return in
}

func FloatToInt64(f float64) int64{
	return int64(f)
}

func FloatToInt(f float64) int{
	return int(FloatToInt64(f))
}

func InterfaceToMIntInt64(d interface{})(r map[int]int64){
	t := d.(map[string]interface{})
	r = make(map[int]int64)
	for i,k := range t {
		r[StringToInt(i)] = FloatToInt64(k.(float64))
	}
	return
}

func InterfaceToMStringFace(d interface{})(r map[string]interface{}){
	t := d.(map[string]interface{})
	for k,v := range t {
		fmt.Println(k,v)
	}
	r = t
	return
}

func InterfaceToInt(d interface{}) (r int) {
	t := d.(float64)
	r = int(FloatToInt64(t))
	return
}

func InterfaceToInt64(d interface{}) (r int64) {
	t := d.(float64)
	r = FloatToInt64(t)
	return
}

func InterfaceToIntList(d interface{}) []int {
	t := d.([]interface{})
	r := make([]int, 0)
	for _, v := range t {
		r = append(r, InterfaceToInt(v))
	}
	return r
}
