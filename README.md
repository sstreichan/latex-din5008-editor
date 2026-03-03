# 📄 LaTeX DIN 5008 Brief Editor

Online-Editor für Briefe nach **DIN 5008** – vollständig im Browser via **Go WebAssembly**.

🔗 **[Live Demo → sstreichan.github.io/latex-din5008-editor](https://sstreichan.github.io/latex-din5008-editor/)**

> **Hinweis:** Nach dem ersten Push muss GitHub Pages einmalig auf **"GitHub Actions"** als Quelle umgestellt werden.
> `Settings → Pages → Source: GitHub Actions`

## Features

- **Go WebAssembly** – Formularlogik & LaTeX-Generierung läuft nativ im Browser
- **CodeMirror Editor** – Syntax-Highlighting für LaTeX mit `Ctrl+Enter` zum Kompilieren
- **DIN 5008A Vorlage** – KOMA-Script `scrlttr2` mit `DIN5008A`-Option
- **PDF Kompilierung** – via [latex.ytotech.com](https://latex.ytotech.com) (kostenlos, kein Account)
- **PDF Vorschau & Download** – direkt im Browser
- **GitHub Pages** – vollständig statisch hostbar

## Tech Stack

| Komponente | Technologie |
|---|---|
| UI-Logik & Template | Go 1.22 → WebAssembly (`syscall/js`) |
| LaTeX Editor | CodeMirror 5 (stex mode) |
| PDF Kompilierung | latex.ytotech.com REST API |
| Styling | Vanilla CSS (Dark Theme) |
| Hosting | GitHub Pages |
| CI/CD | GitHub Actions |

## Lokale Entwicklung

```bash
# Voraussetzung: Go >= 1.22

# WASM bauen & wasm_exec.js kopieren
make build

# Lokaler Webserver
make serve
# → http://localhost:8080

# Aufräumen
make clean
```

## Projektstruktur

```
├── main.go                      # Go WASM – Formular → LaTeX Generierung
├── go.mod
├── Makefile
├── docs/
│   ├── index.html               # Single-Page App
│   ├── style.css                # Dark Theme
│   ├── main.wasm                # (via GitHub Actions / make build)
│   └── wasm_exec.js             # (via GitHub Actions / make build)
└── .github/workflows/
    └── deploy.yml               # Build WASM + Deploy to Pages
```

## LaTeX-Sonderzeichen

Da die Formularfelder direkt in das LaTeX-Dokument eingesetzt werden, müssen
Sonderzeichen in Eingabefeldern (außer dem Brieftext) ggf. manuell escaped werden:

| Zeichen | LaTeX | Zeichen | LaTeX |
|---------|-------|---------|-------|
| `&` | `\&` | `%` | `\%` |
| `$` | `\$` | `#` | `\#` |
| `_` | `\_` | `~` | `\textasciitilde{}` |

Im **Brieftext**-Feld können direkt LaTeX-Befehle verwendet werden.

## Lizenz

MIT
