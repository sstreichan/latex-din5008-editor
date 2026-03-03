package main

import (
	"strings"
	"syscall/js"
)

const latexTemplate = `\documentclass[
  fontsize=12pt,
  version=last,
  parskip=half,
]{scrlttr2}

\usepackage[utf8]{inputenc}
\usepackage[T1]{fontenc}
\usepackage[ngerman]{babel}
\LoadLetterOption{DIN5008A}

\setkomavar{fromname}{<<FROMNAME>>}
\setkomavar{fromaddress}{<<FROMSTREET>>\\<<FROMCITY>>}
\setkomavar{fromphone}{<<FROMPHONE>>}
\setkomavar{fromemail}{<<FROMEMAIL>>}

\KOMAoptions{
  foldmarks=true,
  fromalign=right,
  fromrule=aftername,
  fromphone=true,
  fromemail=true,
  backaddress=true,
}

\begin{document}
\begin{letter}{<<TONAME>>\\<<TOSTREET>>\\<<TOCITY>>}

\setkomavar{place}{<<PLACE>>}
\setkomavar{date}{\today}
\setkomavar{subject}{<<SUBJECT>>}

\opening{<<OPENING>>}

<<BODY>>

\closing{<<CLOSING>>}

<<ENCL>>

\end{letter}
\end{document}`

func generateLatex(this js.Value, args []js.Value) interface{} {
	doc := js.Global().Get("document")
	get := func(id string) string {
		el := doc.Call("getElementById", id)
		if el.IsNull() || el.IsUndefined() {
			return ""
		}
		return el.Get("value").String()
	}

	encl := strings.TrimSpace(get("encl"))
	enclLine := ""
	if encl != "" {
		parts := strings.Split(encl, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		enclLine = `\encl{` + strings.Join(parts, ` \\ `) + `}`
	}

	latex := latexTemplate
	latex = strings.ReplaceAll(latex, "<<FROMNAME>>", get("fromName"))
	latex = strings.ReplaceAll(latex, "<<FROMSTREET>>", get("fromStreet"))
	latex = strings.ReplaceAll(latex, "<<FROMCITY>>", get("fromCity"))
	latex = strings.ReplaceAll(latex, "<<FROMPHONE>>", get("fromPhone"))
	latex = strings.ReplaceAll(latex, "<<FROMEMAIL>>", get("fromEmail"))
	latex = strings.ReplaceAll(latex, "<<TONAME>>", get("toName"))
	latex = strings.ReplaceAll(latex, "<<TOSTREET>>", get("toStreet"))
	latex = strings.ReplaceAll(latex, "<<TOCITY>>", get("toCity"))
	latex = strings.ReplaceAll(latex, "<<PLACE>>", get("place"))
	latex = strings.ReplaceAll(latex, "<<SUBJECT>>", get("subject"))
	latex = strings.ReplaceAll(latex, "<<OPENING>>", get("opening"))
	latex = strings.ReplaceAll(latex, "<<BODY>>", get("body"))
	latex = strings.ReplaceAll(latex, "<<CLOSING>>", get("closing"))
	latex = strings.ReplaceAll(latex, "<<ENCL>>", enclLine)

	editor := js.Global().Get("editor")
	if !editor.IsNull() && !editor.IsUndefined() {
		editor.Call("setValue", latex)
	}
	return latex
}

func getEditorContent(this js.Value, args []js.Value) interface{} {
	editor := js.Global().Get("editor")
	if editor.IsNull() || editor.IsUndefined() {
		return ""
	}
	return editor.Call("getValue").String()
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("goGenerateLaTeX", js.FuncOf(generateLatex))
	js.Global().Set("goGetEditorContent", js.FuncOf(getEditorContent))
	<-done
}
