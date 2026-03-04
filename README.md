# 📄 LaTeX DIN 5008 Brief Editor

Online-Editor für Briefe nach **DIN 5008** mit **Go WebAssembly** für LaTeX-Template-Generierung.

🔗 **[Live Demo → sstreichan.github.io/latex-din5008-editor](https://sstreichan.github.io/latex-din5008-editor/)**

> **Hinweis:** Nach dem ersten Push muss GitHub Pages einmalig auf **"GitHub Actions"** als Quelle umgestellt werden.
> `Settings → Pages → Source: GitHub Actions`

## Features

- **Go WebAssembly** – Formularlogik & LaTeX-Generierung läuft nativ im Browser
- **CodeMirror Editor** – Syntax-Highlighting für LaTeX
- **DIN 5008A Vorlage** – KOMA-Script `scrlttr2` mit `DIN5008A`-Option
- **.tex Export** – Download der generierten LaTeX-Datei
- **100% Datensouveränität** – Alle Daten bleiben lokal im Browser
- **GitHub Pages** – vollständig statisch hostbar

## PDF-Kompilierung

Der Editor generiert LaTeX-Code, den Sie auf drei Arten zu PDF kompilieren können:

### 1. Lokaler Compiler (empfohlen)

Kompilieren Sie die heruntergeladene `.tex`-Datei mit Ihrer lokalen LaTeX-Installation:

```bash
pdflatex brief.tex
```

**Voraussetzung:** TeX Live, MiKTeX oder MacTeX

### 2. Online-Dienste

Laden Sie die `.tex`-Datei in einen dieser Dienste hoch:

- [Overleaf](https://www.overleaf.com) (kostenlos, Account erforderlich)
- [Papeeria](https://papeeria.com) (kostenlos)
- [LaTeX Base](https://latexbase.com) (kein Account)

### 3. Cloud-Compiler API

Verwenden Sie einen dieser kostenlosen APIs (senden Daten an externe Server):

- [latex.ytotech.com](https://latex.ytotech.com)
- [LaTeX.Online](https://latexonline.cc)

## Warum keine Browser-Kompilierung?

**SwiftLaTeX** (LaTeX als WebAssembly) wäre technisch die ideale Lösung für eine vollständig browser-basierte Kompilierung. Allerdings:

- Keine zuverlässige CDN-Distribution verfügbar
- npm-Paket ist veraltet und nicht CDN-kompatibel
- Selbst-Hosting würde ~50 MB WASM-Binary + TeX-Pakete erfordern

Für maximale Datensouveränität empfehlen wir daher einen **lokalen Compiler**.

## Tech Stack

| Komponente | Technologie |
|---|---|
| UI-Logik & Template | Go 1.22 → WebAssembly (`syscall/js`) |
| LaTeX Editor | CodeMirror 5 (stex mode) |
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
