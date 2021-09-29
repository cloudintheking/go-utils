package stream

type T interface{} //原始类型
type U interface{} //转换类型

type Stream interface {
	Of(T) Stream                                                   //创造流
	Map(func(T) U) Stream                                          //转换
	FlatMap(func(T) []U) Stream                                    //拍平
	Filter(func(T) bool) Stream                                    //过滤
	Sort(func(data []T, i, j int) bool) Stream                     //排序
	Distinct(mapperKey func(T) U) Stream                           //去重
	ToSlice() []T                                                  //输出切片
	CollectToMap(mapperKey func(T) U, collect func(T) U) map[U][]U //转换为map
	Foreach(func(int, T))                                          //遍历
}

func convertT(in []interface{}) []T {
	ts := make([]T, 0)
	for _, i := range in {
		ts = append(ts, T(i))
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
