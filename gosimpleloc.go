package gosimpleloc

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"gopkg.in/yaml.v2"
)

// LString - function was written for very specific needs
var LString = func(l, s string, a ...interface{}) string {
	defaultLang := language.Russian
	var lang language.Tag
	switch l {
	case "KZ", "KK":
		lang = language.Make("KK")
	case "EN":
		lang = language.Make("EN")
	default:
		lang = language.Make("RU")
	}
	trans := message.NewPrinter(lang).Sprintf(s, a...)
	if l != "RU" && trans == s {
		trans = message.NewPrinter(defaultLang).Sprintf(s, a...)
	}
	return trans
}

func InitDictionary(localizationDir string) (err error) {
	files, err := ioutil.ReadDir(localizationDir)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, f := range files {
		yamlFile, err := ioutil.ReadFile(localizationDir + "/" + f.Name())
		if err != nil {
			return err
		}
		data := map[string]string{}
		err = yaml.Unmarshal(yamlFile, &data)
		if err != nil {
			return err
		}

		langID := language.Make(strings.Split(f.Name(), ".")[0])
		for k, v := range data {
			message.SetString(langID, k, v)
		}
	}
	return nil
}
