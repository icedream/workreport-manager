package main

import (
	"log"
	"os"

	"github.com/icedream/workreportmgr/export/latex"
	"github.com/icedream/workreportmgr/project"

	"github.com/alecthomas/kingpin"
)

var (
	cli = kingpin.New("workreport-manager", "Manage all of your workreports through an easy to read and manage format.")

	flagProjectFile = cli.Flag("file", "Defines the project filename from which to parse all workreports.").Short('f').Default("workreports.yml").File()
	flagVerbose     = cli.Flag("verbose", "Print extra information about what's happening.").Short('v').Bool()
	flagLocale      = cli.Flag("locale", "The locale to use for exports.").Default("en-US").String()

	cmdInit = cli.Command("init", "Creates a new workreports project file.")

	cmdAdd = cli.Command("add", "Add new information to the workreport.")

	cmdAddActivity = cmdAdd.Command("activity", "Adds a new work activity to the workreport for the current date.")

	cmdAddPeriod = cmdAdd.Command("period", "Adds a new school period to the workreport for the current date.")

	cmdAddActivityDetails          = cmdAdd.Command("details", "Sets the detail information for a work activity for the current date.")
	cmdAddActivityDetailsFlagInput = cmdAddActivityDetails.Flag("input", "The file from which to read the information. Defaults to stdin.").Short('i').Default(os.Stdin.Name()).ExistingFile()

	cmdExportToLatex            = cli.Command("latex", "Exports the current workreport project to Latex.")
	cmdExportToLatexFlagOutput  = cmdExportToLatex.Flag("output", "The file to write the generated code to. Defaults to stdout.").Short('o').Default(os.Stdout.Name()).String()
	cmdExportToLatexFlagInput   = cmdExportToLatex.Flag("input", "Use this to write additional \\input statements in the resulting LaTeX file.").Short('i').Strings()
	cmdExportToLatexFlagProgram = cmdExportToLatex.Flag("program", "The LaTeX compiler program to use, for example \"lualatex\". Leave empty if you don't know what to put here.").Short('p').String()

	version = "master"

	currentProject *project.Project
)

func main() {
	cli.Version(version)
	if *flagVerbose {
		log.Printf("%s v%s", cli.Name, version)
	}

	command := kingpin.MustParse(cli.Parse(os.Args[1:]))

	switch command {
	case cmdInit.FullCommand():
		log.Println("Not yet implemented") // TODO
	case cmdAddActivity.FullCommand():
		log.Println("Not yet implemented") // TODO
	case cmdAddActivityDetails.FullCommand():
		log.Println("Not yet implemented") // TODO
	case cmdAddPeriod.FullCommand():
		log.Println("Not yet implemented") // TODO
	case cmdExportToLatex.FullCommand():
		parseProject()
		if *flagVerbose {
			log.Println("Going to export to Latex file...")
		}
		w, err := os.Create(*cmdExportToLatexFlagOutput)
		if err != nil {
			log.Print("Failed to create output file.")
			os.Exit(1)
		}
		log.Println("Created output file.")
		defer w.Close()
		exporter := new(latex.Exporter)
		if flagLocale != nil {
			exporter.Locale = *flagLocale
		}
		if cmdExportToLatexFlagInput != nil {
			exporter.Inputs = *cmdExportToLatexFlagInput
		}
		if cmdExportToLatexFlagProgram != nil {
			exporter.Marker.Program = *cmdExportToLatexFlagProgram
		}
		log.Println("Now exporting...")
		if err := exporter.Export(currentProject, w); err != nil {
			log.Print("Failed to generate LaTeX output:", err)
			os.Exit(1)
		}
		log.Println("Done.")
	default:
		log.Println("Unknown command.")
	}
}

func parseProject() {
	var err error
	defer func() {
		if err != nil {
			log.Print("Failed to parse project file:", err)
			os.Exit(1)
		}
	}()

	currentProject, err = project.DecodeFromFile(*flagProjectFile)
	if *flagVerbose {
		log.Println("Parsed project file successfully.")
		log.Printf("Resulting project: %+v", currentProject)
	}
}
