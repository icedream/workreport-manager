# Work Report Manager

This tool simplifies management of work reports down to the point where all you
have to do as an apprentice is writing down the actual content, then just sending
it through the tool to generate the document you're going to print.

Right now only output to a LaTeX document is supported but it's enough to get
the work done and get beautiful output from it. The report can easily be styled
through separate LaTeX code which then can be imported.

## Example

`workreports.yml`:

```yaml
Begin: 2016-09-01

Name: Elena Example
Department: Web
First day is monday: true
Only show work days: true

Weeks:
    -   Date: 2016-09-01
        Operational activities:
            - Introduction to the company
            - Setting up the desktop
        Operational instructions: |
            The company Example Inc. programs most of their software in C which
            is a programming language that […].
        Professional school:
            -   Subject: General
                Topics:
                    - Introduction to the school
                    - Rules
            -   Subject: Financial management
                Topics:
                    - […]
                    - […]
            -   Subject: Software programming
                Topics:
                    - […]
                    - […]
```

`styleguide.tex`:

```tex
\usepackage{tabularx}
\usepackage[english]{babel}

% Use Tahoma as main font (this only works thanks to the `fontspec` package).
\setmainfont{tahoma}[
	BoldFont = tahomabd,
	Extension = .ttf,
	Ligatures = TeX,
	Path = ./fonts/
]

% Company logo on the left side.
% Assumes there is a logo.png, logo.jpg or similar in the same folder.
\fancyhead[L]{
	\begin{tabularx}{6cm}{l}
		\includegraphics[width=6cm]{logo} \\
	\end{tabularx}
}
```

Assuming you have the two files above you can compile your work report as a
LaTeX document using this command:

    workreportmgr latex -i styleguide
