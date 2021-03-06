package generators

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	configs "github.com/crowdeco/bima/v2/configs"
)

const MODULES_FILE = "swaggers/modules.json"

type Swagger struct {
}

func (g *Swagger) Generate(template *configs.Template, modulePath string, packagePath string, templatePath string) {
	workDir, _ := os.Getwd()
	modules, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", workDir, MODULES_FILE))
	modulesJson := []configs.ModuleJson{}

	json.Unmarshal(modules, &modulesJson)
	modulesJson = append(modulesJson, configs.ModuleJson{
		Name: template.Module,
		Url:  fmt.Sprintf("./%s.swagger.json", template.ModuleLowercase),
	})

	modulesJson = g.makeUnique(modulesJson)
	modules, _ = json.Marshal(modulesJson)

	err := ioutil.WriteFile(fmt.Sprintf("%s/%s", workDir, MODULES_FILE), modules, 0644)
	if err != nil {
		panic(err)
	}
}

func (g *Swagger) makeUnique(modules []configs.ModuleJson) []configs.ModuleJson {
	occured := make(map[string]bool)
	var result []configs.ModuleJson
	for _, e := range modules {
		if occured[e.Name] != true {
			occured[e.Name] = true

			result = append(result, e)
		}
	}

	return result
}
