package stream

import (
	"errors"
	"github.com/cloudintheking/go-utils/common"
	"reflect"
	"sort"
)

var Streams *streams

func init() {
	Streams = &streams{}
}

//流接口实现
type streams struct {
	data []interface{} //切片数据
}

func (s *streams) Of(arr interface{}) Stream {
	data, err := common.Interface2Slice(arr)
	if err != nil {
		panic(err)
	}
	ns := &streams{
		data: data,
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

func (s *streams) Map(convert func(interface{})interface{}) Stream {
	s.requireNonNil()
	newData := make([]interface{}, 0)
	for _, d := range s.data {
		newData = append(newData, convert(d))
	}
	s.data = newData
	return s
}

func (s *streams) FlatMap(flat func(t interface{}) []interface{}) Stream {
	s.requireNonNil()
	newData := make([]interface{}, 0)
	for _, d := range s.data {
		newData = append(newData, flat(d)...)
	}
	s.data =newData
	return s
}

func (s *streams) Filter(filter func(interface{}) bool) Stream {
	s.requireNonNil()
	newData := make([]interface{}, 0)
	for _, d := range s.data {
		if filter(d) {
			newData = append(newData, d)
		}
	}
	s.data = newData
	return s
}

func (s *streams) Sort(compare func(data []interface{}, i, j int) bool) Stream {
	s.requireNonNil()
	td:=s.data
	sort.SliceStable(td, func(i, j int) bool {
		return compare(td, i, j)
	})
	return s
}

func (s *streams) Distinct(mapperKey func(interface{})interface{}) Stream {
	s.requireNonNil()
	dataMap := make(map[interface{}]interface{}, 0)
	for _, d := range s.data {
		key := mapperKey(d)
		if _, ok := dataMap[key]; !ok {
			dataMap[key] = d
		}
	}
	newData := make([]interface{}, 0)
	for _, v := range dataMap {
		newData = append(newData, v)
	}
	s.data = newData
	return s
}

func (s *streams) ToSlice() interface{} {
	return s.data
}

func (s *streams) CollectToMap(mapperKey func(interface{}) interface{}, collect func(interface{}) interface{}) map[interface{}][]interface{} {
	s.requireNonNil()
	dataMap := make(map[interface{}][]interface{}, 0)
	for _, d := range s.data {
		key := mapperKey(d)
		if _, ok := dataMap[key]; !ok {
			dataMap[key] = make([]interface{}, 0)
		}
		dataMap[key] = append(dataMap[key], collect(d))
	}
	return dataMap
}
func (s *streams) Foreach(foreach func(int, interface{})) {
	s.requireNonNil()
	for i, d := range s.data {
		foreach(i, d)
	}
}
