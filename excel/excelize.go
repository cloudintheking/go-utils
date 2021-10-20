/**
 * @Author: cloudintheking
 * @Description:
 * @File: excel
 * @Version: 1.0.0
 * @Date: 2021/10/11 16:23
 */
package excel

import (
	"github.com/xuri/excelize/v2"
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

/**
 *  @Description:  批量合并单元格
 *  @param excel excel操作对象
 *  @param sheetName 表名
 *  @param axisList 位置集合 [0][1]代表左上角位置  [2][3]代表右下角位置
 *  @return error
 */
func BatchMergeCells(excel *excelize.File, sheetName string, axisList [][]int) error {
	for _, axis := range axisList {
		if err := excel.MergeCell(sheetName, GetCellAxis(axis[0], axis[1]), GetCellAxis(axis[2], axis[3])); err != nil {
			return err
		}
	}
	return nil
}
