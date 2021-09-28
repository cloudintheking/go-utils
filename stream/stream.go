package stream

import (
	"errors"
	"github.com/cloudintheking/go-utils/slice"
	"reflect"
	"sort"
)

type T interface{} //原始类型
type U interface{} //转换类型

var Streams *streams

func init() {
	Streams = &streams{}
}

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

//流接口实现
type streams struct {
	data []T //切片数据
}

func (s *streams) Of(arr T) Stream {
	data := slice.Interface2Slice(arr)
	ns := &streams{
		data: convertT(data),
	}
	reflectTypeMap := make(map[reflect.Type]interface{}, 0)
	//判断每个元素类型是否一致
	for _, d := range ns.data {
		reflectTypeMap[reflect.TypeOf(d)] = nil
	}
	if len(reflectTypeMap) != 1 {
		panic(errors.New("element type should be same"))
	}
	return ns
}

func (s *streams) requireNonNil() {
	if s.data == nil || len(s.data) == 0 {
		panic(errors.New("data can not be empty"))
	}
}

func (s *streams) Map(convert func(T) U) Stream {
	s.requireNonNil()
	newData := make([]T, 0)
	for _, d := range s.data {
		newData = append(newData, convert(d))
	}
	s.data = newData
	return s
}

func (s *streams) FlatMap(flat func(t T) []U) Stream {
	s.requireNonNil()
	newData := make([]U, 0)
	for _, d := range s.data {
		newData = append(newData, flat(d)...)
	}
	s.data = convertU2T(newData)
	return s
}

func (s *streams) Filter(filter func(T) bool) Stream {
	s.requireNonNil()
	newData := make([]T, 0)
	for _, d := range s.data {
		if filter(d) {
			newData = append(newData, d)
		}
	}
	s.data = newData
	return s
}

func (s *streams) Sort(compare func(data []T, i, j int) bool) Stream {
	s.requireNonNil()
	sort.Slice(s.data, func(i, j int) bool {
		return compare(s.data, i, j)
	})
	return s
}

func (s *streams) Distinct(mapperKey func(T) U) Stream {
	s.requireNonNil()
	dataMap := make(map[U]T, 0)
	for _, d := range s.data {
		key := mapperKey(d)
		if _, ok := dataMap[key]; !ok {
			dataMap[key] = d
		}
	}
	newData := make([]T, 0)
	for _, v := range dataMap {
		newData = append(newData, v)
	}
	s.data = newData
	return s
}

func (s *streams) ToSlice() []T {
	return s.data
}

func (s *streams) CollectToMap(mapperKey func(T) U, collect func(T) U) map[U][]U {
	s.requireNonNil()
	dataMap := make(map[U][]U, 0)
	for _, d := range s.data {
		key := mapperKey(d)
		if _, ok := dataMap[key]; !ok {
			dataMap[key] = make([]U, 0)
		}
		dataMap[key] = append(dataMap[key], collect(d))
	}
	return dataMap
}
func (s *streams) Foreach(foreach func(int, T)) {
	s.requireNonNil()
	for i, d := range s.data {
		foreach(i, d)
	}
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
