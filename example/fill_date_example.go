/**
 * @Author: cloudintheking
 * @Description:
 * @File: fill_date_test
 * @Version: 1.0.0
 * @Date: 2021/10/11 14:24
 */
package example

import (
	"fmt"
	"github.com/cloudintheking/go-utils/slices"
	"time"
)

type DateFillTest struct {
	CreateTime *time.Time
}

func (t *DateFillTest) GetFillDate() *time.Time {
	return t.CreateTime
}

func (t *DateFillTest) SetFillDate(date *time.Time) {
	t.CreateTime = date
}

func FillDateTest() {
	data := make([]*DateFillTest, 0)
	y2001 := time.Date(2001, time.January, 1, 0, 0, 0, 0, time.Local)
	y2005 := y2001.AddDate(4, 0, 0)
	y2010 := y2005.AddDate(5, 0, 0)
	data = append(data, &DateFillTest{
		CreateTime: &y2010,
	})
	data = append(data, &DateFillTest{
		CreateTime: &y2001,
	})
	data = append(data, &DateFillTest{
		CreateTime: &y2005,
	})
	iData, _ := slices.FillBlankDate(data, slices.YEAR)

	fd := iData.([]slices.DateFillInterface)
	for _, d := range fd {
		fmt.Printf("%v", d)
		fmt.Println()
	}
}
