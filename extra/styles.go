package extra

import (
	"fmt"
	"image/color"
	"reflect"
)

type StyleValueType int

const (
	SVTColor StyleValueType = iota
	SVTInt
	SVTString
	SVTFloat32
	SVTFloat64
)

func ConvertToSVT(svt StyleValueType, key string, val any) any {
	t := reflect.TypeOf(val)

	switch svt {
	case SVTColor:
		if t == reflect.TypeOf(color.RGBA{}) {
			return val
		} else {
			panic(fmt.Sprintf("invalid style type for key %s: %s (expected color.RGBA{})", key, t.Name()))
		}
	case SVTString:
		switch val.(type) {
		case string:
			return val
		default:
			panic(fmt.Sprintf("invalid style type for key %s: %s (expected string)", key, t.Name()))
		}
	case SVTInt:
		switch val.(type) {
		case int, int8, int16, int32, int64:
			return val
		default:
			panic(fmt.Sprintf("invalid style type for key %s: %s (expected int, int8, int16, int32 or int64)", key, t.Name()))
		}
	case SVTFloat32:
		switch val.(type) {
		case float32:
			return val
		case float64:
			return float32(val.(float64))
		case int:
			return float32(val.(int))
		default:
			panic(fmt.Sprintf("invalid style type for key %s: %s (expected float32, float64 or int)", key, t.Name()))
		}
	case SVTFloat64:
		switch val.(type) {
		case float64:
			return val
		case float32:
			return float64(val.(float32))
		case int:
			return float64(val.(int))
		default:
			panic(fmt.Sprintf("invalid style type for key %s: %s (expected float32, float64 or int)", key, t.Name()))
		}
	default:
		panic(fmt.Sprintf("couldn't convert %s into any valid style type", t.Name()))
	}
}

func GetSVT(val any) StyleValueType {
	t := reflect.TypeOf(val)
	if t == reflect.TypeOf(color.RGBA{}) {
		return SVTColor
	}

	switch val.(type) {
	case string:
		return SVTString
	case int, int8, int16, int32, int64:
		return SVTInt
	case float32:
		return SVTFloat32
	case float64:
		return SVTFloat64
	}

	panic(fmt.Sprintf("%s isn't any valid style value type", t.Name()))
}
