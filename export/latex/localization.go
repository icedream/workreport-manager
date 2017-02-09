package latex

import (
	"log"

	"github.com/nicksnyder/go-i18n/i18n"
)

//go:generate go-bindata -pkg latex -o localization_assets.go localization/

func initLocalization() {
	log.Println("Initializing localization for LaTeX export...")

	files, err := AssetDir("localization")
	if err != nil {
		log.Fatal("Failed to browse embedded asset directory.", err)
		panic(err)
	}

	for _, file := range files {
		if localizationBytes, err := Asset("localization/" + file); err != nil {
			log.Fatal("Failed to read localization file.", err)
			panic(err)
		} else if err := i18n.ParseTranslationFileBytes(file, localizationBytes); err != nil {
			log.Fatal("Failed to parse localization file.", err)
			panic(err)
		}
	}
}
