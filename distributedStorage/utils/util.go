package utils

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strconv"
	strings2 "strings"
	"time"
)

// 随机数生成
func Random(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("the length of the parameter strings should not be less than 0")
	}

	if length <= 0 || len(strings) <= length {
		return "", errors.New("the size of the parameter length illegal")
	}

	for i := len(strings) - 1; i > 0; i-- {
		rand.Seed(time.Now().Unix())
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i] + ","
	}
	str = strings2.Trim(str, ",")
	return str, nil
}

// 去重
func RemoveRep(s []int) []int {
	result := []int{}
	m := make(map[int]bool) //map的值不重要
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	return result
}

// 移除数组元素
func RemoveParam(sli []string, n string) []string {
	for i := 0; i < len(sli); i++ {
		if sli[i] == n {
			if i == 0 {
				sli = sli[1:]
			} else if i == len(sli)-1 {
				sli = sli[:i]
			} else {
				sli = append(sli[:i], sli[i+1:]...)
			}
			i-- // 如果索引i被去掉后，原来索引i+1的会移动到索引i
		}
	}
	return sli
}

// interface 转换string
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
