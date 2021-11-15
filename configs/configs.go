package configs

import (
	"github.com/jianghaibo12138/TMZ/pkg/tools"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path"
	"runtime"
)

type YamlSetting struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     int16  `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Debug    bool   `yaml:"debug"`
		Flavor   string `yaml:"flavor"`
	}
}

var Settings = new(YamlSetting)

func init() {
	AppConfig()
}

func GetHomePath() string {
	if runtime.GOOS == "windows" {
		return ".\\"
	} else {
		return "./"
	}
}

func AppConfig() {
	homePath := GetHomePath()
	yamlPath := path.Join(homePath, "configs", "dev.conf.yml")
	if !tools.IsFile(yamlPath) {
		yamlPath = path.Join(homePath, "configs", "conf.yml")
	}
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, Settings)
	if err != nil {
		panic(err)
	}
}
