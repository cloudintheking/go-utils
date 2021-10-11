/**
 * @Author: cloudintheking
 * @Description:
 * @File: 0_init
 * @Version: 1.0.0
 * @Date: 2021/10/11 14:12
 */
package common

import (
	"errors"
	"reflect"
)

/**
 * @Description:  接口转切片
 * @param slice
 * @return []interface{}
 */
func Interface2Slice(slice interface{}) ([]interface{}, error) {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return nil, errors.New("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret, nil
}
