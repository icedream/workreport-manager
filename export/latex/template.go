package latex

import (
	"log"
	"os"
	"text/template"

	"git.dekart811.net/icedream/workreportmgr/export/latex/stringutil"
	"git.dekart811.net/icedream/workreportmgr/project"
	"github.com/jinzhu/now"
	"github.com/nicksnyder/go-i18n/i18n"
)

const dateFormat = "02.01.2006" // TODO - localize!

var exportTemplate = template.Must(template.
	New("latex_export").
	Funcs(template.FuncMap{
		"T": i18n.IdentityTfunc,
		"beginofweek": func(date project.Date) project.Date {
			return project.Date{Time: now.New(date.Time).BeginningOfWeek()}
		},
		"endofweek": func(date project.Date) project.Date {
			return project.Date{Time: now.New(date.Time).EndOfWeek()}
		},
		"escape": stringutil.TexEscape,
		"add": func(a, b int) int {
			return a + b
		},
	}).
	Delims("<", ">").
	Parse(`% !TeX
<- with .TexMarker.Program> program = <.><end>

\documentclass[11pt,a4paper,oneside]{article}

\usepackage{fancyhdr}
\usepackage{tabularx}
\usepackage[left=2cm,right=2cm,top=2cm,bottom=6cm,includeheadfoot]{geometry}
\usepackage{csquotes}
\usepackage{fontspec}
\usepackage{ifluatex}
\ifluatex
\else
	\usepackage[utf8]{inputenc}
\fi

\pagestyle{fancy}
\fancyhf{}
\setlength{\headheight}{2cm}

\newcommand{\Name}{<.Project.Name>}
\newcommand{\Department}{<.Project.Department>}

\newcommand{\wrSigningField}[0]{
	\begin{tabularx}{\textwidth}{| X | X | X |}
		\hline
		<T "trainee"> &
		<T "legal_representative"> &
		<T "instructor"> \\[2cm]
		\hline
	\end{tabularx}
}

\newenvironment{weeklyreport}{
}{
	\fancyfoot[C]{\wrSigningField}
}

\newcommand{\preweeklyreporthead}[3]{}

\newcommand{\weeklyreporthead}[3]{
	\preweeklyreporthead{#1}{#2}{#3}
	\fancyhead[R]{
		\begin{tabularx}{8cm}{rl}
			\textbf{<T "name">}: & \Name \\
			\textbf{<T "department">}: & \Department \\
			\textbf{<T "time_period">}: & #2 - #3 \\
		\end{tabularx}
	}
	\fancyfoot{}
	\newpage
	\setcounter{section}{#1}
	\setcounter{subsection}{0}
	\section*{<T "proof_of_education" "#1">}
	\addcontentsline{toc}{section}{<T "proof_of_education" "#1">}
}

\newcommand{\weeklyreportsection}[1]{
	\subsection*{#1}
}

<with .TexInputs>
<range .>
\input{<.>}
<end>
<end>

\begin{document}
\tableofcontents

<range $index, $week := .Project.Weeks>
\begin{weeklyreport}
	\weeklyreporthead{<add $index 1>}{<beginofweek $week.Date>}{<endofweek $week.Date>}

	\weeklyreportsection{<T "operational_activities">}
	<with $week.WorkActivities>
	\begin{itemize}
		<range .>
		\item <. | escape>
		<end>
	\end{itemize}
	<end>

	\weeklyreportsection{<T "operational_instruction">}
	<$week.WorkActivityDetails | escape>

	\weeklyreportsection{<T "professional_school">}
	<if eq (len $week.Periods) 0>
		<T "no_school_periods_this_week" | escape>
	<else>
	\begin{itemize}
		<range $week.Periods>
		\item{
			<- .Subject | escape ->
			<- with .Topics ->:
			<range $index, $topic := . ->
			<- if ne $index 0 ->, <end -><- $topic | escape ->
			<- end ->
			<- end ->
		}
		<end>
	\end{itemize}
	<end>

\end{weeklyreport}
<end>

\end{document}
`))

func initTemplate() {
	log.Println("Initializing template for LaTeX export...")

	var err error
	defer func() {
		if err != nil {
			log.Fatal("Failed at initializing template.", err)
			os.Exit(0xFF)
		}
	}()
}
