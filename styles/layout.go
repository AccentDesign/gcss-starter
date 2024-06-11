package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
)

// Layout styles for the stylesheet.
func (ss *StyleSheet) Layout() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				MinHeight: variables.FullScreenHeight,
			},
		},
		{
			Selector: "main",
			Props: gcss.Props{
				Display: props.DisplayGrid,
			},
		},
	}
}

// Layout styles for the media.
func (m *Media) Layout() Styles {
	return Styles{
		{
			Selector: "main",
			Props: gcss.Props{
				Padding: m.Padding,
				RowGap:  m.RowGap,
			},
		},
	}
}

// Layout styles for the theme.
func (t *Theme) Layout() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				BackgroundColor: t.Background,
				Color:           t.Foreground,
			},
		},
	}
}
