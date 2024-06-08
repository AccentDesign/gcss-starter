package styles

import (
	"bytes"
	"fmt"
	"github.com/AccentDesign/gcss"
	"io"
	"slices"
	"sync"
	"time"
)

type (
	Styles     []gcss.Style
	StyleSheet struct {
		Media  []*Media
		Themes []*Theme
		css    bytes.Buffer
		mutex  sync.Mutex
	}
)

// CSS Writes the CSS for the stylesheet to the writer.
func (ss *StyleSheet) CSS(w io.Writer) error {
	defer func(start time.Time) {
		fmt.Printf("CSS: %s\n", time.Since(start))
	}(time.Now())

	ss.mutex.Lock()
	defer ss.mutex.Unlock()

	// If the CSS has already been written, return it.
	if ss.css.Len() > 0 {
		_, err := w.Write(ss.css.Bytes())
		return err
	}

	// Write the CSS for the base styles.
	for _, style := range slices.Concat(
		ss.Resets(),
		ss.Layout(),
		ss.Buttons(),
	) {
		if err := style.CSS(&ss.css); err != nil {
			return err
		}
	}

	// Write the CSS for the media queries.
	for _, media := range ss.Media {
		if err := media.CSS(&ss.css); err != nil {
			return err
		}
	}

	// Write the CSS for the themes.
	for _, theme := range ss.Themes {
		if err := theme.CSS(&ss.css); err != nil {
			return err
		}
	}

	// Write the CSS for the utilities.
	if err := UtilityCSS(&ss.css); err != nil {
		return err
	}

	// Write the CSS to the writer.
	_, err := w.Write(ss.css.Bytes())
	return err
}

// NewStyleSheet returns a new stylesheet. It includes the media queries and themes.
func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		Media:  []*Media{mobileMedia, desktopMedia},
		Themes: []*Theme{lightTheme, darkTheme},
	}
}
