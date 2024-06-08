package styles

import (
	"io"
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
