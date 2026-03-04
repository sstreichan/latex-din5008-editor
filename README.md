# 📄 LaTeX DIN 5008 Brief Editor

Online-Editor für Briefe nach **DIN 5008** – vollständig im Browser via **Go WebAssembly** + **SwiftLaTeX**.

🔗 **[Live Demo → sstreichan.github.io/latex-din5008-editor](https://sstreichan.github.io/latex-din5008-editor/)**

> **Hinweis:** Nach dem ersten Push muss GitHub Pages einmalig auf **"GitHub Actions"** als Quelle umgestellt werden.
> `Settings → Pages → Source: GitHub Actions`

## Features

- **100% Browser-basiert** – keine externen Server, keine API-Calls
- **Go WebAssembly** – Formularlogik & LaTeX-Generierung läuft nativ im Browser
- **SwiftLaTeX (WASM)** – Vollständiger PdfTeX-Compiler als WebAssembly
- **CodeMirror Editor** – Syntax-Highlighting für LaTeX mit `Ctrl+Enter` zum Kompilieren
- **DIN 5008A Vorlage** – KOMA-Script `scrlttr2` mit `DIN5008A`-Option
- **PDF Vorschau & Download** – direkt im Browser
- **Offline-fähig** – funktioniert nach dem ersten Laden auch ohne Internet
- **GitHub Pages** – vollständig statisch hostbar

## Tech Stack

| Komponente | Technologie |
|---|---|
| UI-Logik & Template | Go 1.22 → WebAssembly (`syscall/js`) |
| LaTeX Kompilierung | SwiftLaTeX 1.5.1 (PdfTeX via WASM) |
| LaTeX Editor | CodeMirror 5 (stex mode) |
| Styling | Vanilla CSS (Dark Theme) |
| Hosting | GitHub Pages |
| CI/CD | GitHub Actions |

## Warum keine externen Dienste?

**SwiftLaTeX** ist ein vollständiger TeX-Compiler, der in WebAssembly kompiliert wurde [web:28][web:29]. Die gesamte LaTeX-Kompilierung erfolgt lokal im Browser des Benutzers – es werden keine Daten an externe Server gesendet. Dies garantiert:

- ✅ **Volle Datensouveränität** – LaTeX-Dokumente verlassen niemals den Browser
- ✅ **Offline-Nutzung** – funktioniert nach dem ersten Laden ohne Internet
- ✅ **Keine API-Limits** – unbegrenzte Kompilierungen
- ✅ **Schnelle Performance** – keine Netzwerk-Latenz

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
│   ├── index.html               # Single-Page App (SwiftLaTeX Integration)
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

## Technische Details: SwiftLaTeX

SwiftLaTeX basiert auf einer vollständigen WebAssembly-Portierung von **PdfTeX** und **XeTeX** [web:29]. Die Engine lädt beim ersten Start (~3-5 MB), danach ist die Kompilierung rein lokal. Das Projekt wird aktiv gepflegt und wird von vielen LaTeX-Editoren als Basis genutzt [web:28][web:31].

## Lizenz

MIT
