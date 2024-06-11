package styles

import (
	"github.com/AccentDesign/gcss"
	"io"
	"reflect"
)

var (
	transitionColors = gcss.Style{
		Selector: ".transition-colors",
		CustomProps: []gcss.CustomProp{
			{Attr: "transition-property", Value: "color, background-color, border-color, text-decoration-color, fill, stroke"},
			{Attr: "transition-timing-function", Value: "cubic-bezier(0.4, 0, 0.2, 1)"},
			{Attr: "transition-duration", Value: "150ms"},
		},
	}
)

var utilities = Styles{}

// UtilityCSS Writes the CSS for the utilities to the writer.
func UtilityCSS(w io.Writer) error {
	for _, style := range utilities {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	return nil
}

// MergeProps merges multiple gcss.Props structs into a single struct.
func MergeProps(props ...gcss.Props) gcss.Props {
	merged := gcss.Props{}
	for _, p := range props {
		v := reflect.ValueOf(p)
		vType := v.Type()
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if !field.IsZero() {
				reflect.ValueOf(&merged).Elem().FieldByName(vType.Field(i).Name).Set(field)
			}
		}
	}
	return merged
}

// MergeStyles merges multiple gcss.Style structs into a single struct.
// The first style is used as the base style.
func MergeStyles(styles ...gcss.Style) gcss.Style {
	merged := gcss.Style{Selector: styles[0].Selector}

	for _, s := range styles {
		merged.Props = MergeProps(merged.Props, s.Props)
		merged.CustomProps = append(merged.CustomProps, s.CustomProps...)
	}

	return merged
}
