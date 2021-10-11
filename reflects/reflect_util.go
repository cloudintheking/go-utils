package reflects

import (
	"reflect"
	"unsafe"
)

/**
 *  @Description: 获取未导出的字段值
 *  @param field 反射字段
 *  @return interface{}
 */
func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

/**
 *  @Description:  未导出字段设置
 *  @param field 反射字段
 *  @param val
 */
func SetUnexportedField(field reflect.Value, val interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(val))
}
