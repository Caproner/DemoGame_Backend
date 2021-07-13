package tr

/*
这个模块是用于简单的数据转换
*/

import "strconv"

// Int64ToString 是int64 转 string
func Int64ToString(i int64)string{
	return strconv.FormatInt(i,10)
}
