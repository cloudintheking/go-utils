/**
 * @Author: cloudintheking
 * @Description:
 * @File: excelizeWrap
 * @Version: 1.0.0
 * @Date: 2021/10/11 16:23
 */
package excelizeWrap

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"time"
)

// 表格写入(单个赋值)
func TestWriteExcelBySetCellValue() {
	excel := excelize.NewFile()
	// 创建一个工作表
	sheet := excel.NewSheet("学校")
	// 设置新建工作表(学校)的内容
	_ = excel.SetCellValue("学校", "A2", "北京大学")
	_ = excel.SetCellValue("学校", "A3", "南京大学")
	// 设置sheet1的内容(默认创建)
	excel.SetCellValue("Sheet1", "A1", "张三")
	excel.SetCellValue("Sheet1", "A2", "小明")
	// 设置默认工作表
	excel.SetActiveSheet(sheet)
	// 保存表格
	if err := excel.SaveAs("target/test.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行完成")
}

// 表格写入(按行写入)
func TestWriteByLine() {
	excel := excelize.NewFile()
	excelStream, _ := excel.NewStreamWriter("sheet")
	fmt.Println(excelStream)
	// 写入标题
	titleSlice := []interface{}{"序号", "姓名", "年龄", "性别"}
	_ = excel.SetSheetRow("Sheet1", "A1", &titleSlice)
	data := []interface{}{
		[]interface{}{1, "张三", 19, "男"},
		[]interface{}{2, "小丽", 18, "女"},
		[]interface{}{3, "小明", 20, "男"},
	}
	// 遍历写入数据
	for key, datum := range data {
		axis := fmt.Sprintf("A%d", key+2)
		// 利用断言，转换类型
		tmp, _ := datum.([]interface{})
		_ = excel.SetSheetRow("Sheet1", axis, &tmp)
	}
	excel.MergeCell("Sheet1", "A2", "A4")
	style, _ := excel.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		}})
	excel.SetCellStyle("Sheet1", "A2", "A4", style)
	mergeCells, _ := excel.GetMergeCells("Sheet1")
	fmt.Println("合并单元格有:", mergeCells)

	excel.GetRows("Sheet1")
	// 保存表格
	if err := excel.SaveAs("target/line.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行完成")
}

// 表格写入(按行流式写入)
func TestWriteByStream() {
	excel := excelize.NewFile()
	// 获取流式写入器
	streamWriter, err := excel.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println("获取流式写入器失败: " + err.Error())
		return
	}
	// 按行写入
	if err := streamWriter.SetRow("A1", []interface{}{"序号", "商品码", "价格"}); err != nil {
		fmt.Println("获取流式写入器失败: " + err.Error())
		return
	}

	// 制作数据
	// 设置随机因子
	rand.Seed(time.Now().Unix())
	for i := 2; i < 500000; i++ {
		tmp := []interface{}{
			i,
			fmt.Sprintf("P-%d", rand.Intn(100000000)),
			fmt.Sprintf("%.2f", float64(rand.Intn(10))+rand.Float64()),
		}
		_ = streamWriter.SetRow("A"+strconv.Itoa(i), tmp)
	}
	// 调用 Flush 函数来结束流式写入过程
	if err = streamWriter.Flush(); err != nil {
		fmt.Println("结束流式写入失败: " + err.Error())
		return
	}
	// 保存表格
	if err := excel.SaveAs("target/stream.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("执行完成")
}

// 读取表格
func TestRead() {
	// 打开表格文件
	openFile, err := excelize.OpenFile("target/line.xlsx")
	if err != nil {
		fmt.Println("打开表格文件失败: " + err.Error())
		return
	}
	// 读取指定工作表所有数据
	rows, err := openFile.GetRows("Sheet1")
	cells, _ := openFile.GetMergeCells("Sheet1")
	fmt.Println(cells)
	if err != nil {
		fmt.Println("读取失败: " + err.Error())
		return
	}
	for _, row := range rows {
		fmt.Printf("%+v\n", row)
	}
	fmt.Println("执行完成!")
}
