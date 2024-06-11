<img src="https://github.com/AccentDesign/gcss/raw/main/banner.jpg" alt="banner" style="width: 100%; height: auto;">

# gcss starter

Includes the bare-bones basics to generate a working stylesheet with:

- [x] CSS resets.
- [x] Media queries for screen layout.
- [x] Theme for light and dark mode.
- [x] Button styles.
- [x] Useful helper functions to merge styles and props a bit like css mixins, see [buttons.go](./styles/buttons.go) and [utility.go](./styles/utility.go) for an example.
- [x] Hot reload with [air](https://github.com/cosmtrek/air).
- [x] Example HTML and CSS.
- [x] Mutex to cache the stylesheet to prevent multiple builds.

## Installation

Hot reload is enabled by using [air](https://github.com/cosmtrek/air):

```bash
go install github.com/cosmtrek/air@latest
```

golang deps:

```bash
go mod tidy
```

## Server

```bash
air
```

## Usage

For an example of adding form styles.

1. Create a new `form.go` file. and add the following code (or add what you need):

```go
// Form returns the styles for buttons for the base stylesheet.
func (ss *StyleSheet) Form() Styles {
	return Styles{}
}

// Form returns the styles for the layout for the media.
func (m *Media) Form() Styles {
	return Styles{}
}

// Form returns the styles for the layout for the theme.
func (t *Theme) Form() Styles {
	return Styles{}
}
```

2. Then add the functions to the `CSS` function in `StyleSheet`, `Media` and `Theme`.


## HTML

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta content="width=device-width, initial-scale=1" name="viewport"/>
  <meta content="Example starting point written in gcss." name="description"/>
  <title>Theme</title>
  <link rel="stylesheet" href="/stylesheet.css">
</head>
<body>
<main>
  <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
  <div>
    <button class="button button-primary">Click me</button>
    <button class="button button-outline">Click me</button>
  </div>
</main>
</body>
</html>
```

## CSS

```css
/* formatted for visual clarity */

/* resets */
*, ::after, ::before, ::backdrop, ::file-selector-button {
  border: 0 solid;
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

/* ... other resets ... */

/* base */
body {
  min-height: 100vh;
}

main {
  display: grid;
}

.button {
  align-items: center;
  border-radius: 0.375rem;
  display: inline-flex;
  font-size: 0.875rem;
  font-weight: 500;
  height: 2.500rem;
  justify-content: center;
  line-height: 1.250rem;
  padding-bottom: 0.500rem;
  padding-left: 1.000rem;
  padding-right: 1.000rem;
  padding-top: 0.500rem;
  transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}

@media (max-width: 768px) {
  main {
    padding: 2.000rem;
    row-gap: 1.500rem;
  }
}

@media (min-width: 769px) {
  main {
    padding: 4.000rem;
    row-gap: 2.000rem;
  }
}

@media (prefers-color-scheme: light) {
  body {
    background-color: rgba(255, 255, 255, 1.00);
    color: rgba(23, 23, 23, 1.00);
  }

  .button-primary {
    background-color: rgba(23, 23, 23, 1.00);
    color: rgba(255, 255, 255, 1.00);
  }

  .button-primary:hover {
    background-color: rgba(23, 23, 23, 0.90);
  }

  .button-outline {
    border: 1px solid rgba(23, 23, 23, 1.00);
  }

  .button-outline:hover {
    border-color: rgba(23, 23, 23, 0.80);
  }
}

@media (prefers-color-scheme: dark) {
  body {
    background-color: rgba(23, 23, 23, 1.00);
    color: rgba(245, 245, 245, 1.00);
  }

  .button-primary {
    background-color: rgba(255, 255, 255, 1.00);
    color: rgba(23, 23, 23, 1.00);
  }

  .button-primary:hover {
    background-color: rgba(255, 255, 255, 0.90);
  }

  .button-outline {
    border: 1px solid rgba(255, 255, 255, 1.00);
  }

  .button-outline:hover {
    border-color: rgba(255, 255, 255, 0.80);
  }
}
```