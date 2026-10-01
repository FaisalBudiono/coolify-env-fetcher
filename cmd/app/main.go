package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"FaisalBudiono/coolify-env-fetcher/internal/coolify"
	"FaisalBudiono/coolify-env-fetcher/internal/mapper"
)

func main() {
	baseFlag := flag.String("base", "", "Base URL")
	accessTokenFlag := flag.String("access", "", "Coolify Access Token")
	appIDFlag := flag.String("app", "", "Coolify App ID")
	isPreviewFlag := flag.String("is-preview", "false", "(optional) Is Preview?")
	flag.Parse()

	if *baseFlag == "" || *accessTokenFlag == "" || *appIDFlag == "" {
		fmt.Println("Should complete all parameters")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var isPreview bool
	if isPreviewFlag != nil && strings.ToLower(*isPreviewFlag) == "true" {
		isPreview = true
	}

	res, err := coolify.ParseENV(*baseFlag, *appIDFlag, *accessTokenFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := ".env"
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := mapper.NewDotENV()
	err = m.WriteFile(f, res, isPreview)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
