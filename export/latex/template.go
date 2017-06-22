package latex

import (
	"log"
	"os"
	"text/template"

	"github.com/icedream/workreport-manager/export/latex/stringutil"
	"github.com/icedream/workreport-manager/project"
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
		"add": func(a, b int) int {
			return a + b
		},
		"latexize": stringutil.Latexize,
	}).
	Delims("<", ">").
	Parse(`% !TeX
<- with .TexMarker.Program> program = <.><end>
\documentclass[11pt,a4paper,oneside]{article}

\usepackage[left=2cm,right=2cm,top=2cm,bottom=6cm,includeheadfoot]{geometry}
\usepackage{csquotes}
\usepackage{fancyhdr}
\usepackage{hyperref}
\usepackage{tabularx}
\usepackage{titling}

\pagestyle{fancy}
\fancyhf{}
\setlength{\headheight}{2cm}

\newcommand{\department}[1]{\global\def\wrDepartment{#1}}
\newcommand{\workreportNumber}[1]{\global\def\wrNumber{#1}}
\newcommand{\workreportDateFrom}[1]{\global\def\wrDateFrom{#1}}
\newcommand{\workreportDateTo}[1]{\global\def\wrDateTo{#1}}

\author{<.Project.Name>}
\department{<.Project.Department>}

\newcommand{\wrSigningField}[0]{
	<if .Project.HideLegalRepSignField>
	\begin{tabularx}{\textwidth}{| X | X | X |}
	<else>
	\begin{tabularx}{\textwidth}{| X | X |}
	<end>
		\hline
		<T "trainee"> &
		<if not .Project.HideLegalRepSignField>
		<T "legal_representative"> &
		<end>
		<T "instructor"> \\[2cm]
		\hline
	\end{tabularx}
}

\fancypagestyle{weeklyreport} {
	\fancyfoot[C]{\wrSigningField}
}

\newenvironment{weeklyreport}[3]{
	\newpage
	\thispagestyle{weeklyreport}
	\fancyhead[R]{
		\begin{tabularx}{8cm}{rl}
			\textbf{<T "name">}: & \theauthor \\
			\textbf{<T "department">}: & \wrDepartment \\
			\textbf{<T "time_period">}: & \wrDateFrom~-~\wrDateTo \\
		\end{tabularx}
	}

	\workreportNumber{#1}
	\workreportDateFrom{#2}
	\workreportDateTo{#3}

	\setcounter{section}{\wrNumber}
	\setcounter{subsection}{0}
	\section*{<T "proof_of_education_prefix"> \wrNumber}
	\addcontentsline{toc}{section}{<T "proof_of_education_prefix"> \wrNumber}
}{
}

\newcommand{\weeklyreporthead}[3]{}

\newcommand{\weeklyreportsection}[1]{
	\subsection*{#1}
}

<with .TexInputs>

% Includes

<range .>
\input{<.>}
<end>
<end>

\begin{document}
\tableofcontents

<range $index, $week := .Project.Weeks>
\begin{weeklyreport}{<index $.WeekNumbers $index>}{<beginofweek $week.Date>}{<endofweek $week.Date>}

	\weeklyreportsection{<T "operational_activities">}
	<with $week.WorkActivities>
	\begin{itemize}
		<range .>
		\item <. | latexize>
		<end>
	\end{itemize}
	<end>

	\weeklyreportsection{<T "operational_instruction">}
	<$week.WorkActivityDetails | latexize>

	\weeklyreportsection{<T "professional_school">}
	<if eq (len $week.Periods) 0>
		<T "no_school_periods_this_week">
	<else>
	\begin{itemize}
		<range $week.Periods>
		\item{
			<- .Subject | latexize ->
			<- with .Topics ->:
			<range $index, $topic := . ->
			<- if ne $index 0 ->, <end -><- $topic | latexize ->
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
