package tool

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

// FormatInt
// interface{} 转换为整形
func FormatInt(data interface{}) (int, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(data).Uint()), nil
	case string:
		return strconv.Atoi(data.(string))
	}
	return 0, errors.New("invalid value type")
}

// FormatInt64
// 转换成int64
func FormatInt64(data interface{}) (int64, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		return i, err
	case float32, float64:
		return int64(reflect.ValueOf(data).Float()), nil
	case int, int8, int16, int32, int64:
		return int64(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(reflect.ValueOf(data).Uint()), nil
	case string:
		return strconv.ParseInt(data.(string), 10, 64)
	}
	return 0, errors.New("invalid value type")
}

// FormatFloat
// interface{} 转换为浮点
func FormatFloat(data interface{}) (float64, error) {
	switch data.(type) {
	case json.Number:
		return data.(json.Number).Float64()
	case float32, float64:
		return reflect.ValueOf(data).Float(), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(data).Uint()), nil
	case string:
		return strconv.ParseFloat(data.(string), 64)
	}
	return 0, errors.New("invalid value type")
}

// FormatString
// interface{} 转换为字符串
func FormatString(data interface{}) (string, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(i, 10), nil
	case float32, float64:
		fdata := reflect.ValueOf(data).Float()
		return strconv.FormatFloat(fdata, 'f', -1, 64), nil
	case int, int8, int16, int32, int64:
		idata := reflect.ValueOf(data).Int()
		return strconv.FormatInt(idata, 10), nil
	case uint, uint8, uint16, uint32, uint64:
		udata := reflect.ValueOf(data).Uint()
		return strconv.FormatUint(udata, 10), nil
	case string:
		return data.(string), nil
	}
	return "", errors.New("invalid value type")
}

// FormatMapStringString
// 转换为map类型其中key,value都是字符类型的
func FormatMapStringString(data interface{}) (map[string]string, error) {
	if mapData, ok := data.(map[string]string); ok {
		return mapData, nil
	} else {
		return nil, errors.New("invalid map type")
	}
}

// FormatMapString
// 转换为map，key为字符串类型
func FormatMapString(data interface{}) (map[string]interface{}, error) {
	if mapData, ok := data.(map[string]interface{}); ok {
		return mapData, nil
	} else {
		return nil, errors.New("invalid map type")
	}
}

// FormatArrayString
// 转成字符串数组
func FormatArrayString(data interface{}) ([]string, error) {
	if arrData, ok := data.([]string); ok {
		return arrData, nil
	} else {
		return nil, errors.New("invalid map type")
	}
}

// FormatArrayInterface
// 转成字符串数组
func FormatArrayInterface(data interface{}) ([]interface{}, error) {
	if arrData, ok := data.([]interface{}); ok {
		return arrData, nil
	} else {
		return nil, errors.New("invalid map type")
	}
}

// FormatArrayMap
// 转换为map的数组
func FormatArrayMap(data interface{}) ([]map[string]interface{}, error) {
	var tmp []interface{}
	var tmpMap []map[string]interface{}

	var ok = false
	if tmpMap, ok = data.([]map[string]interface{}); ok {
		return tmpMap, nil
	}

	if tmp, ok = data.([]interface{}); ok {
		list := []map[string]interface{}{}
		for _, valIf := range tmp {
			var valMap map[string]interface{}
			if valMap, ok = valIf.(map[string]interface{}); !ok {
				return nil, errors.New("invalid array map type")
			}
			list = append(list, valMap)
		}
		return list, nil
	}

	return nil, errors.New("invalid array map type")
}
