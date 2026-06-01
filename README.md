<h3 align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/logos/exports/1544x1544_circle.png" width="100" alt="Logo"/><br/>
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
	Catppuccin for Websites.
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
</h3>

<p align="center">
	<a href="https://github.com/catppuccin/template/stargazers"><img src="https://img.shields.io/github/stars/catppuccin/template?colorA=363a4f&colorB=b7bdf8&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/template/issues"><img src="https://img.shields.io/github/issues/catppuccin/template?colorA=363a4f&colorB=f5a97f&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/template/contributors"><img src="https://img.shields.io/github/contributors/catppuccin/template?colorA=363a4f&colorB=a6da95&style=for-the-badge"></a>
</p>

## Themes

* [theme-frappe.css](https://github.com/Martin-Ludwig/catppuccin-css/blob/main/dist/theme-frappe.css)
* [theme-mocha.css](https://github.com/Martin-Ludwig/catppuccin-css/blob/main/dist/theme-mocha.css)
* [theme-latte.css](https://github.com/Martin-Ludwig/catppuccin-css/blob/main/dist/theme-latte.css)
* [theme-macchiato.css](https://github.com/Martin-Ludwig/catppuccin-css/blob/main/dist/theme-macchiato.css)

## Usage

Pick a flavor and pull in the stylesheet, then use the
variables (`--bg`, `--text`, `--link`, …) for your colors. See
[`STYLE-GUIDE.md`](STYLE-GUIDE.md) for the full variable reference.

```html
<link rel="stylesheet" href="theme-mocha.css">
```

```css
body      { background: var(--bg); color: var(--text); }
a         { color: var(--link); }
a:hover   { color: var(--link-hover); }
```

> [!TIP]
> To make the theme easy to override, import it into a low-priority
> [cascade layer](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/At-rules/@layer).
> Unlayered styles always beat layered ones, so your own `:root` rules win:
>
> ```css
> @import "theme-mocha.css" layer(theme);
> ```

<p align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" />
</p>

<p align="center">
	Copyright &copy; 2021-present <a href="https://github.com/catppuccin" target="_blank">Catppuccin Org</a>
</p>

<p align="center">
	<a href="https://github.com/catppuccin/catppuccin/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=for-the-badge&label=License&message=MIT&logoColor=d9e0ee&colorA=363a4f&colorB=b7bdf8"/></a>
</p>
