package slices



/**
 *  @Description:  int64切片去重
 *  @param slice
 *  @return []interface{}
 */
func DistinctInt64Slice(slice []*int64) []*int64 {
	temp := make(map[int64]*int64)
	result := make([]*int64, 0)
	for _, s := range slice {
		if _, ok := temp[*s]; !ok {
			temp[*s] = s
			result = append(result, s)
		}
	}
	return result
}

/**
 *  @Description: 字符串切片求交集
 *  @param s1
 *  @param s2
 *  @return []string
 */
func IntersectStringSlice(s1, s2 []string) []string {
	m := make(map[string]interface{})
	for _, v := range s1 {
		m[v] = struct{}{}
	}

	intersection := make([]string, 0)
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

/**
 *  @Description: 字符串切片求差集
 *  @param s1
 *  @param s2
 *  @return []string
 */
func DiffStringSlice(s1Slice, s2Slice []string) []string {
	res := make([]string, 0)
	for _, s1 := range s1Slice {
		var exist = false
		for _, s2 := range s2Slice {
			if s1 == s2 {
				exist = true
				break
			}
		}
		if !exist {
			res = append(res, s1)
		}
	}
	return res
}

/**
 *  @Description:  两个float64切片对应相加
 *  @param slice1
 *  @param slice2
 *  @return slice3
 */
func SliceTwoFloatAdd(slice1 []float64, slice2 []float64) []float64 {
	len1 := len(slice1)
	len2 := len(slice2)
	len := len1
	if len1 < len2 {
		len = len2
	}
	slice3 := make([]float64, len, len)
	for i := 0; i < len; i++ {
		if i < len1 && i < len2 {
			slice3[i] = slice1[i] + slice2[i]
		}
		if i > len1 {
			slice3[i] = slice2[i]
		}
		if i > len2 {
			slice3[i] = slice1[i]
		}
	}
	return slice3
}

/**
 *  @Description: 任意数量float64切片对应相加相加
 *  @param sli
 *  @return []float64
 */
func SliceMuchFloatAdd(sli ...[]float64) []float64 {
	if sli == nil || len(sli) == 0 {
		return nil
	}
	origin := sli[0]
	if len(sli) == 1 {
		return origin
	}
	sli = append(sli[:0], sli[1:]...)
	for _, s := range sli {
		origin = SliceTwoFloatAdd(origin, s)
	}
	return origin
}
