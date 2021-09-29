package slice

//时间填充接口
type DateFillInterface interface {
	GetYear() int                  //获取年份
	GetMonth() int                 //获取月份
	GetDay() int                   //获取天份
	SetYear(year int)              //设置年份
	SetMonth(month int)            //设置月份
	SetDay(day int)                //设置天份
	SetExtDefault(ext interface{}) //设置额外初始值
}


