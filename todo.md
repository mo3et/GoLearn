# 需要加深理解和解决的点
- interface
- struct
- slice 二维slice
- map 
- 匿名函数
- select的使用



# 知识点
- slice 是不限制数组长度的array  s:=make([]string,3)
- slice[low:high]low为起始位置 high为slice容量
- Channel是在goroutine之间传输数据 ch:=make(chan DataType) 例如int interface{}