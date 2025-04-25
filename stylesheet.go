package vex

import (
	"image/color"
)

type VStyleSheet struct {
	widgetSpecificStyles map[*VWidget]map[string]interface{}
	styleClasses         map[string]interface{}
}

func newStyleSheet() *VStyleSheet {
	return &VStyleSheet{
		make(map[*VWidget]map[string]interface{}),
		make(map[string]interface{}),
	}
}

func (s *VStyleSheet) GetKeyRaw(widget *VWidget, key string) interface{} {
	if val, ok := s.widgetSpecificStyles[widget][key]; ok {
		return val
	}

	return s.styleClasses[key]
}

func (s *VStyleSheet) SetKey(widget *VWidget, key string, value interface{}) {
	if s.widgetSpecificStyles[widget] == nil {
		s.widgetSpecificStyles[widget] = make(map[string]interface{})
	}

	s.widgetSpecificStyles[widget][key] = value
}

func (s *VStyleSheet) GetKeyAsColor(widget *VWidget, key string) color.RGBA {
	return s.GetKeyRaw(widget, key).(color.RGBA)
}
