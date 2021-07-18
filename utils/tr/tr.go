package tr

/*
这个模块是用于简单的数据转换
*/

import "strconv"

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
