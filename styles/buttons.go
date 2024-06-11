package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

// Buttons styles for the stylesheet.
func (ss *StyleSheet) Buttons() Styles {
	return Styles{
		MergeStyles(
			gcss.Style{
				Selector: ".button",
				Props: gcss.Props{
					AlignItems:     props.AlignItemsCenter,
					BorderRadius:   variables.Size1H,
					Display:        props.DisplayInlineFlex,
					FontSize:       variables.Size3H,
					FontWeight:     props.FontWeightMedium,
					Height:         variables.Size10,
					JustifyContent: props.JustifyContentCenter,
					LineHeight:     variables.Size5,
					PaddingTop:     variables.Size2,
					PaddingRight:   variables.Size4,
					PaddingBottom:  variables.Size2,
					PaddingLeft:    variables.Size4,
				},
			},
			transitionColors,
		),
	}
}

// Buttons styles for the theme.
func (t *Theme) Buttons() Styles {
	return Styles{
		{
			Selector: ".button-primary",
			Props: gcss.Props{
				BackgroundColor: t.Primary,
				Color:           t.PrimaryForeground,
			},
		},
		{
			Selector: ".button-primary:hover",
			Props: gcss.Props{
				BackgroundColor: t.Primary.Alpha(230),
			},
		},
		{
			Selector: ".button-outline",
			Props: gcss.Props{
				Border: props.Border{
					Color: t.Primary,
					Style: props.BorderStyleSolid,
					Width: props.UnitPx(1),
				},
			},
		},
		{
			Selector: ".button-outline:hover",
			Props: gcss.Props{
				BorderColor: t.Primary.Alpha(204),
			},
		},
	}
}
