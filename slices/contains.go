package slices

import "github.com/cloudintheking/go-utils/common"

//判断对象是否相等
type EqualsFunc func(dist interface{}, iter interface{}) bool

/**
 *  @Description: 是否包含元素
 *  @param arr
 *  @param dist
 *  @param equalsFunc
 *  @return interface{} 被包含的元素
 *  @return bool 包含标志
 */
func Contains(arr interface{}, dist interface{}, equalsFunc EqualsFunc) (interface{}, bool) {
	arrD, err := common.Interface2Slice(arr)
	if err != nil {
		panic(err)
	}
	for _, a := range arrD {
		if equalsFunc(dist, a) {
			return a, true
		}
	}
	return nil, false
}
