package utils

import "time"

/**
 * @Description: 获取传入时间的所在天的开始时间
 * @param date
 * @return time.Time
 */
func GetBeginOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

/**
 * @Description: 获取传入时间的所在天的结束时间
 * @param date
 * @return time.Time
 */
func GetEndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999, date.Location())
}

/**
 * @Description:  获取传入年份的起始时间
 * @param year
 * @return time.Time
 */
func GetBeginOfYear(year int) time.Time {
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
}

/**
 * @Description:  获取传入年份的结束时间
 * @param year
 * @return time.Time
 */
func GetEndOfYear(year int) time.Time {
	return time.Date(year, 12, 31, 23, 59, 59, 999, time.Local)
}

/**
 * @Description:  获取传入年份的起始时间
 * @param year
 * @return time.Time
 */
func GetBeginOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

/**
 * @Description:  获取传入年份的结束时间
 * @param year
 * @return time.Time
 */
func GetEndOfMonth(year int, month time.Month) time.Time {
	return GetEndOfDay(GetBeginOfMonth(year, month)).AddDate(0, 1, -1)
}

/**
 * @Description: 判断两个时间段是否重叠
 * @param begin1
 * @param end1
 * @param begin2
 * @param end2
 */
func IsOverLap(begin1 time.Time, end1 time.Time, begin2 time.Time, end2 time.Time) bool {
	return begin1.Before(end2) && end1.After(begin2)
}

/**
 * @Description: 取最小值
 * @param time1
 * @param time2
 * @return time.Time
 */
func MinTime(time1 time.Time, time2 time.Time) time.Time {
	if time1.Before(time2) {
		return time1
	} else {
		return time2
	}

}

/**
 * @Description:  判断时间是否在某一时间段内
 * @param date
 * @param start
 * @param end
 * @return bool
 */
func IsInTimeRange(date time.Time, start time.Time, end time.Time) bool {
	if CompareTime(date, start) >= 0 && CompareTime(date, end) <= 0 {
		return true
	}
	return false
}

/**
 * @Description:  比较时间大小（精确到秒）
 * @param u
 * @param t
 * @return int64
 */
func CompareTime(u time.Time, t time.Time) int64 {
	us := u.Unix()
	ts := t.Unix()
	if us < ts {
		return -1
	} else if us > ts {
		return 1
	} else {
		return 0
	}
}

/**
 * @Description:  相差天数
 * @param start
 * @param end
 * @return int64
 */
func DiffDays(start time.Time, end time.Time) float64 {
	return float64((end.Unix() - start.Unix())) / 86400
}

/**
 * @Description:  相差纳秒
 * @param start
 * @param end
 * @return int64
 */
func DiffNanos(start time.Time, end time.Time) float64 {
	return float64((end.UnixNano() - start.UnixNano()))
}