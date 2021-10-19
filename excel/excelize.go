/**
 * @Author: cloudintheking
 * @Description:
 * @File: excel
 * @Version: 1.0.0
 * @Date: 2021/10/11 16:23
 */
package excel

import (
	"strconv"
)

// maxCharCount 最多26个字符A-Z
const maxCharCount = 26

/**
 *  @Description: 获取列名
 *  @param column 列号
 *  @return string
 */
func GetColumnName(column int) []byte {
	const A = 'A'
	if column < maxCharCount {
		// 第一次就分配好切片的容量
		slice := make([]byte, 0)
		return append(slice, byte(A+column))
	} else {
		// 递归生成类似AA,AB,AAA,AAB这种形式的列名
		return append(GetColumnName(column/maxCharCount-1), byte(A+column%maxCharCount))
	}
}

/**
 *  @Description: 获取单元格坐标轴
 *  @param column 列位置
 *  @param row 行位置
 *  @return string
 */
func GetCellAxis(column, row int) string {
	if column < 0 || row < 0 {
		panic("column or row should grater than 0.")
	}
	return string(GetColumnName(column)) + strconv.Itoa(row)
}
