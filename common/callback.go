/**
 * @Author: cloudintheking
 * @Description: 解决包循环引用
 * @File: callback
 * @Version: 1.0.0
 * @Date: 2021/10/9 10:18
 */
package common

import (
	"fmt"
	"reflect"
)

var callBackMap map[string]interface{}

func init() {
	callBackMap = make(map[string]interface{})
}

/**
 *  @Description:  注册回调函数
 *  @param key
 *  @param callBack
 */
func RegisterCallBack(key string, callBack interface{}) {
	callBackMap[key] = callBack
}

/**
 *  @Description:  调用回调函数
 *  @param key
 *  @param args
 *  @return []interface{}
 */
func CallBackFunc(key string, args ...interface{}) []interface{} {
	if callbackFunc, ok := callBackMap[key]; ok {
		in := make([]reflect.Value, len(args))
		for i, arg := range args {
			in[i] = reflect.ValueOf(arg)
		}
		outList := reflect.ValueOf(callbackFunc).Call(in)
		result := make([]interface{}, len(outList))
		for i, out := range outList {
			result[i] = out.Interface()
		}
		return result
	} else {
		panic(fmt.Errorf("callBack(%s) not found", key))
	}
}
