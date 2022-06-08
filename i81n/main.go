package main

import (
	"github.com/BurntSushi/toml"
	"github.com/abadojack/whatlanggo"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main(){
      /* Set language translation */

      // Create a new i18n bundle with default language.
      bundle := i18n.NewBundle(language.English)

      // Register a toml unmarshal function for i18n bundle.
      bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

      // Load translations from toml files for non-default languages.
      bundle.MustLoadMessageFile("./lang/active.ar.toml")
      bundle.MustLoadMessageFile("./lang/active.es.toml")

      var lang string
  
        input := "Hi"
        // input := "هاي"
  
				info := whatlanggo.Detect(input)
				fmt.Println("Language:", info.Lang.String(), " Script:", whatlanggo.Scripts[info.Script], " Confidence: ", info.Confidence)

				switch whatlanggo.Scripts[info.Script] {
				case "Arabic":
					lang = "ar"
          fmt.Println("Arabic")
				case "Latin":
          fmt.Println("Latin")
				}

				// Create a new localizer.
				localizer := i18n.NewLocalizer(bundle, lang)
				// Set title message.
				helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: &i18n.Message{
						ID:    "HelloPerson",     // set translation ID
						Other: "Hello {{.Name}}", // set default translation
					},
					TemplateData: map[string]string{
            "Name": "Hasan",
					},
					PluralCount: nil,
				})

				fmt.Println(helloPerson)
  
}
