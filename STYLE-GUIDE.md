# Web Style Guide

How to use the Catppuccin palette on a **website**.

```html
<!-- pick your theme -->
<link rel="stylesheet" href="theme-mocha.css">
```

```css
body { 
    background: var(--bg); 
    color: var(--text); 
}
a { color: var(--link); }
a:hover { color: var(--link-hover); }
::selection { background: var(--selection-bg); }
```

> [!IMPORTANT]
> Text colors are guidelines. Some cases need deviations - e.g. `--text` on a colored
> background. Legibility always comes first, so use your own judgement.

### Background Colors

| Variable | Palette color | Use for |
| --- | --- | --- |
| `--bg` | Base | Main page / app background |
| `--bg-secondary` | Mantle | Secondary panes (sidebars, headers) |
| `--bg-tertiary` | Crust | Deepest panes (gutters, wells) |
| `--surface-0` | Surface 0 | Raised surfaces - cards, inputs |
| `--surface-1` | Surface 1 | Hover / borders on surfaces |
| `--surface-2` | Surface 2 | Active / stronger borders |
| `--overlay-0` | Overlay 0 | Faint UI lines, disabled fills |
| `--overlay-1` | Overlay 1 | Subtle UI elements |
| `--overlay-2` | Overlay 2 | Stronger overlay accents |

### Typography & UI

| Variable | Palette color | Use for |
| --- | --- | --- |
| `--text` | Text | Body copy and headlines |
| `--label` | Subtext 1 | Sub-headlines, labels |
| `--label-muted` | Subtext 0 | De-emphasized labels, captions |
| `--subtle` | Overlay 1 | Subtle / placeholder text |
| `--accent` | Blue | Primary action color (buttons, active state) |
| `--accent-hover` | Sky | Accent hover state |
| `--on-accent` | Base | Text/icons **on top of** an accent fill (e.g. a colored button) |
| `--link` | Blue | Links, URLs, tags, pills |
| `--link-hover` | Sky | Link hover state |
| `--success` | Green | Success messages, valid state |
| `--warning` | Yellow | Warnings |
| `--error` | Red | Errors, invalid state |
| `--info` | Teal | Informational messages |
| `--cursor` | Rosewater | Text cursor / caret |
| `--selection-bg` | Overlay 2 @ 25% | Text-selection background |

> [!NOTE]
> `--selection-bg` is a `color-mix()` of Overlay 2 at 25% opacity, per the palette
> guideline of 20–30% for selection backgrounds. Adjust the percentage in the template
> if you want it lighter or stronger.
