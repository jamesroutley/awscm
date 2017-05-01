package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

var cmdFile = flag.String("file", "", "file to write commands to")

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"use: awscm COMMAND")
	}
}

func throw(msg string) {
	flag.Usage()
	fmt.Println("Error: " + msg)
	os.Exit(1)
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		throw("no command supplied")
	}

	switch flag.Arg(0) {
	case "init":
		cliInit()
	case "ls":
		cliLs()
	case "use":
		cliUse()
	default:
		throw("unrecognised command")

	}
}

func cliInit() {
	text, err := ioutil.ReadFile("awscm.sh")
	if err != nil {
		throw("could not read the file 'awscm.sh'")
	}
	fmt.Println("# Load awscm automatically by appending")
	fmt.Println("# the following to your shell startup file:")
	fmt.Println("")
	fmt.Println(string(text))

}

func cliLs() {
	fmt.Println(strings.Join(profileNames(), "\n"))
}

func cliUse() {
	if flag.NArg() == 1 {
		msg := fmt.Sprintf("specify a valid profile:\n%s",
			strings.Join(profileNames(), "\n"))
		throw(msg)
	}
	profile := flag.Arg(1)
	profiles := profiles()
	if !profiles[profile] {
		msg := fmt.Sprintf("profile not found, specify a valid profile:\n%s",
			strings.Join(profileNames(), "\n"))
		throw(msg)
	}
	cmds := []string{
		fmt.Sprintf("export AWS_PROFILE=%s\n", profile),
		fmt.Sprintf("export AWS_DEFAULT_PROFILE=%s\n", profile),
	}
	write(cmds)
}

func profileNames() []string {
	names := []string{}
	for k, _ := range profiles() {
		names = append(names, k)
	}
	return names
}

func profiles() map[string]bool {
	profiles := make(map[string]bool)

	credCfg, err := ini.Load("/Users/helsinki/.aws/credentials")
	if err != nil {
		fmt.Println(err)
	}

	for _, section := range credCfg.SectionStrings() {
		if section == "DEFAULT" {
			continue
		}
		profiles[section] = true
	}

	configCfg, err := ini.Load("/Users/helsinki/.aws/config")
	if err != nil {
		fmt.Println(err)
	}

	for _, section := range configCfg.SectionStrings() {
		if section == "DEFAULT" {
			continue
		}
		profiles[strings.TrimLeft(section, "profile ")] = true
	}

	return profiles
}

func write(cmds []string) {
	f, err := os.Create(*cmdFile)
	if err != nil {
		throw(fmt.Sprintf("Could not open '%s'\n", *cmdFile))
	}
	defer f.Close()
	for _, cmd := range cmds {
		// fmt.Fprintf(f, cmd)
		_, err = f.Write([]byte(cmd))
		if err != nil {
			throw(fmt.Sprintf("could not write to '%s': %v\n", *cmdFile, err))
		}
	}
}
