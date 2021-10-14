package slices

import (
	"errors"
	"fmt"
	"github.com/cloudintheking/go-utils/common"
	"github.com/cloudintheking/go-utils/stream"
	"reflect"
	"time"
)

const (
	YEAR = 0 + iota
	MONTH
	DAY
)

//时间填充接口
type DateFillInterface interface {
	GetFillDate() *time.Time //获取填充时间
	SetFillDate(*time.Time)  //设置填充时间
}

func FillBlankDate(data interface{}, fillType int) (interface{}, error) {
	if data == nil {
		return nil, nil
	}
	//将data转换为 []interface
	sData, err := common.Interface2Slice(data)
	if err != nil || len(sData) == 0 {
		return nil, err
	}
	iType := reflect.TypeOf((*DateFillInterface)(nil)).Elem()
	newData := make([]DateFillInterface, 0)
	//判断元素是否都实现了DateFillInterface接口
	for _, d := range sData {
		dT := reflect.TypeOf(d)
		if !dT.Implements(iType) {
			panic(errors.New(fmt.Sprintf("expect element type: %s,but get element type: %s ", iType.Name(), dT.Name())))
		}
		//转为[]DateFillInterface
		newData = append(newData, d.(DateFillInterface))
	}
	elementType := reflect.TypeOf(sData[0]) //元素类型
	for {
		if elementType.Kind() == reflect.Ptr {
			elementType = elementType.Elem()
		} else {
			break
		}
	}
	switch fillType {
	case YEAR:
		return FillYearBlankDate(newData, elementType)
	case MONTH:
		return FillMonthBlankDate(newData, elementType)
	case DAY:
		return FillDayBlankDate(newData, elementType)
	default:
		return data, nil
	}
}

/**
 *  @Description:  按年填充
 *  @param data
 *  @param elementType
 *  @return interface{}
 */
func FillYearBlankDate(data []DateFillInterface, elementType reflect.Type) (interface{}, error) {
	//将data转换为[]int切片yearSlice(去重并升序排序)
	yearSlice := stream.Streams.Of(data).
		Map(func(t interface{}) interface{} {
			date := t.(DateFillInterface)
			return date.GetFillDate().Year()
		}).
		Distinct(func(t interface{}) interface{} {
			return t.(int)
		}).
		Sort(func(data []interface{}, i, j int) bool {
			iD := data[i].(int)
			jD := data[j].(int)
			return iD < jD
		}).
		ToSlice().([]stream.T)

	length := len(yearSlice)
	//拿到最大值maxYear,遍历yearSlice: year < maxYear
	minYear := yearSlice[0].(int)
	maxYear := yearSlice[length-1].(int)

	fillYear := make([]DateFillInterface, 0)
	for i := minYear; i <= maxYear; i++ {
		//构建一个空元素对象
		emptyYear := reflect.New(elementType)
		tempDate := time.Date(i, time.January, 1, 0, 0, 0, 0, time.Local)
		emptyYear.MethodByName("SetFillDate").Call([]reflect.Value{reflect.ValueOf(&tempDate)})
		//判断year是否存在data中,不存在则创建空DateFillInterface
		existsYear, ok := Contains(data, nil, func(dist interface{}, iter interface{}) bool {
			iterYear := iter.(DateFillInterface)
			return iterYear.GetFillDate().Year() == i
		})
		if ok {
			fillYear = append(fillYear, existsYear.(DateFillInterface))
		} else {
			fillYear = append(fillYear, emptyYear.Interface().(DateFillInterface))
		}
	}
	return fillYear, nil
}

/**
 *  @Description: 按月填充
 *  @param data
 *  @param elementType
 *  @return interface{}
 */
func FillMonthBlankDate(data []DateFillInterface, elementType reflect.Type) (interface{}, error) {
	//将data转换为[]time.Time切片monthSlice(去重并升序排序)
	//拿到最大值maxMonth,遍历monthSlice: month < maxMonth
	//判断month是否存在data中,不存在则创建空DateFillInterface
	return nil, nil
}

/**
 *  @Description: 按天填充
 *  @param data
 *  @param elementType
 *  @return interface{}
 */
func FillDayBlankDate(data []DateFillInterface, elementType reflect.Type) (interface{}, error) {
	//将data转换为[]time.Time切片daySlice(去重并升序排序)
	//拿到最大值maxDay,遍历daySlice: day < maxDay
	//判断day是否存在data中,不存在则创建空DateFillInterface
	return nil, nil
}
