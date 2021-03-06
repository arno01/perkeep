// Code generated by reactGen. DO NOT EDIT.

package selectallbutton

import "myitcv.io/react"

type SelectAllBtnElem struct {
	react.Element
}

func (s SelectAllBtnDef) ShouldComponentUpdateIntf(nextProps react.Props, prevState, nextState react.State) bool {
	res := false

	{
		res = s.Props() != nextProps.(SelectAllBtnProps) || res
	}
	return res
}

func buildSelectAllBtn(cd react.ComponentDef) react.Component {
	return SelectAllBtnDef{ComponentDef: cd}
}

func buildSelectAllBtnElem(props SelectAllBtnProps, children ...react.Element) *SelectAllBtnElem {
	return &SelectAllBtnElem{
		Element: react.CreateElement(buildSelectAllBtn, props, children...),
	}
}

// IsProps is an auto-generated definition so that SelectAllBtnProps implements
// the myitcv.io/react.Props interface.
func (s SelectAllBtnProps) IsProps() {}

// Props is an auto-generated proxy to the current props of SelectAllBtn
func (s SelectAllBtnDef) Props() SelectAllBtnProps {
	uprops := s.ComponentDef.Props()
	return uprops.(SelectAllBtnProps)
}

func (s SelectAllBtnProps) EqualsIntf(val react.Props) bool {
	return s == val.(SelectAllBtnProps)
}

var _ react.Props = SelectAllBtnProps{}
