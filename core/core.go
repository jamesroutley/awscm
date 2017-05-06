package core

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
	"os"
	"path"
	"sort"
	"strings"
)

// Regions is a slice of all available AWS regions
var Regions = []string{
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-south-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"ca-central-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"sa-east-1",
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
}

var Outputs = []string{
	"json",
	"table",
	"text",
}

var (
	credFile, configFile string
)

func init() {
	hd, err := homedir.Dir()
	if err != nil {
		fmt.Println("Could not find user's home directory")
		os.Exit(1)
	}
	credFile = path.Join(hd, ".aws", "credentials")
	configFile = path.Join(hd, ".aws", "config")
}

func Profiles() []string {
	profiles := make(map[string]bool)

	credProfiles := iniSections(credFile)
	configProfiles := iniSections(configFile)

	for _, section := range credProfiles {
		if section == "DEFAULT" {
			continue
		}
		profiles[section] = true
	}

	for _, section := range configProfiles {
		if section == "DEFAULT" {
			continue
		}
		profiles[strings.TrimLeft(section, "profile ")] = true
	}
	return sortedKeys(profiles)
}

func iniSections(path string) []string {
	config, err := ini.Load(path)
	if err != nil {
		// TODO: improve
		fmt.Println(err)
	}
	return config.SectionStrings()
}

func sortedKeys(m map[string]bool) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
