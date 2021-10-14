package stream

type T interface{} //原始类型
type U interface{} //转换类型
//同步流
type Stream interface {
	Of(interface{}) Stream                                                   //创造流
	Map(func(interface{}) interface{}) Stream                                          //转换
	FlatMap(func(interface{}) []interface{}) Stream                                    //拍平
	Filter(func(interface{}) bool) Stream                                    //过滤
	Sort(func(data []interface{}, i, j int) bool) Stream                     //排序
	Distinct(mapperKey func(interface{}) interface{}) Stream                           //去重
	ToSlice() interface{}                                          //输出切片
	CollectToMap(mapperKey func(interface{})interface{}, collect func(interface{}) interface{}) map[interface{}][]interface{} //转换为map
	Foreach(func(int, interface{}))                                          //遍历
}

func convertT(in []interface{}) []interface{} {
	ts := make([]interface{}, 0)
	for _, i := range in {
		ts = append(ts, interface{}(i))
	}
	return ts
}

func convertT2U(ts []T) []U {
	us := make([]U, 0)
	for _, t := range ts {
		us = append(us, U(t))
	}
	return us
}

func convertU2T(us []U) []T {
	ts := make([]T, 0)
	for _, u := range us {
		ts = append(ts, T(u))
	}
	return ts
}
