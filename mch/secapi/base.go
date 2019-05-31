package secapi

import (
	"encoding/json"
	"reflect"
	"strconv"
)

func RequestToMap(req interface{}) map[string]string {
	m1 := make(map[string]string)

	beanValue := reflect.ValueOf(req)
	beanIndirectValue := reflect.Indirect(beanValue)
	beanIndirectValueType := beanIndirectValue.Type()

	for i := 0; i < beanIndirectValueType.NumField(); i++ {
		tag := beanIndirectValueType.Field(i).Tag
		xmlTagStr := tag.Get("xml")
		//if xmlTagStr == "xml" {
		//	continue
		//}
		if beanIndirectValue.Field(i).IsValid() {
			m1[xmlTagStr] = GetValue(beanIndirectValue.Field(i))
		}
	}
	return m1
}

func ResponseFromMap(m2 map[string]string, resp interface{}) {
	beanRespValue := reflect.ValueOf(resp)
	beanRespIndirectValue := reflect.Indirect(beanRespValue)
	beanRespIndirectTypeValue := beanRespIndirectValue.Type()
	for i := 0; i < beanRespIndirectTypeValue.NumField(); i++ {
		tag := beanRespIndirectTypeValue.Field(i).Tag
		xmlTagStr := tag.Get("xml")
		//if xmlTagStr == "xml" {
		//	continue
		//}
		if val, has := m2[xmlTagStr]; has {
			beanRespIndirectValue.Field(i).SetString(val)
		}
	}
}

func ToString(v interface{}) string {
	switch value := v.(type) {
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(value, 'g', -1, 64)
	case string:
		return value
	case []byte:
		return string(value)
	case bool:
		if value {
			return "true"
		} else {
			return "false"
		}
	default:
		data, err := json.Marshal(v)
		if err != nil {
			return "{}"
		}
		return string(data)
	}
}

func GetValue(value reflect.Value) string {
	if !value.CanSet() {
		return ""
	}
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return ToString(value.Int())
	case reflect.Float32, reflect.Float64:
		return ToString(value.Float())
	case reflect.String:
		return value.String()
	default:
		return ""
	}
}
