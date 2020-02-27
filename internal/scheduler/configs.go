package scheduler

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Conf struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	}
	MSG struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func NewConf() *Conf {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/configs/application.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := Conf{}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &c
}
