package vex

import (
	"fmt"
	"github.com/snackbag/vex/extra"
	"image/color"
)

type VStyleSheet struct {
	widgetSpecificStyles map[*VWidget]map[string]interface{}
	styleClasses         map[string]interface{}
	typeRegistry         map[string]extra.StyleValueType
}

func newStyleSheet() *VStyleSheet {
	sheet := &VStyleSheet{
		make(map[*VWidget]map[string]interface{}),
		make(map[string]interface{}),
		make(map[string]extra.StyleValueType),
	}

	// Generic
	sheet.ReserveType("color", extra.SVTColor)
	sheet.ReserveType("background-color", extra.SVTColor)

	// Font
	sheet.ReserveType("font-size", extra.SVTInt)

	// Border
	sheet.ReserveType("border-roundness", extra.SVTFloat32)
	sheet.ReserveType("border-width", extra.SVTInt)
	sheet.ReserveType("border-color", extra.SVTColor)
	sheet.ReserveType("border-segments", extra.SVTInt)

	return sheet
}

func (s *VStyleSheet) GetKeyRaw(widget *VWidget, key string) interface{} {
	if val, ok := s.widgetSpecificStyles[widget][key]; ok {
		return val
	}

	return s.styleClasses[key]
}

func (s *VStyleSheet) SetKey(widget *VWidget, key string, value interface{}) {
	resTemp := s.GetKeyTypeReservation(key)
	var res extra.StyleValueType
	if resTemp == nil {
		res = extra.GetSVT(value)
		s.ReserveType(key, res)
	} else {
		res = *resTemp
	}

	if res != extra.GetSVT(value) {
		if res != extra.GetSVT(extra.ConvertToSVT(res, key, value)) {
			panic(fmt.Sprintf("cannot set key %s, because it is reserved for a different type", key))
		}
	}

	if s.widgetSpecificStyles[widget] == nil {
		s.widgetSpecificStyles[widget] = make(map[string]interface{})
	}

	s.widgetSpecificStyles[widget][key] = extra.ConvertToSVT(res, key, value)
}

func (s *VStyleSheet) GetKeyTypeReservation(key string) *extra.StyleValueType {
	if typ, ok := s.typeRegistry[key]; ok {
		return &typ
	}
	return nil
}

func (s *VStyleSheet) ReserveType(key string, typ extra.StyleValueType) {
	s.typeRegistry[key] = typ
}

func (s *VStyleSheet) GetKeyAsColor(widget *VWidget, key string) color.RGBA {
	return s.GetKeyRaw(widget, key).(color.RGBA)
}

func (s *VStyleSheet) GetKeyAsString(widget *VWidget, key string) string {
	return s.GetKeyRaw(widget, key).(string)
}

func (s *VStyleSheet) GetKeyAsInt(widget *VWidget, key string) int {
	return s.GetKeyRaw(widget, key).(int)
}

func (s *VStyleSheet) GetKeyAsFloat32(widget *VWidget, key string) float32 {
	return s.GetKeyRaw(widget, key).(float32)
}

func (s *VStyleSheet) GetKeyAsFloat64(widget *VWidget, key string) float64 {
	return s.GetKeyRaw(widget, key).(float64)
}
