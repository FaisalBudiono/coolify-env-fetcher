package main

import (
	"FaisalBudiono/coolify-env-fetcher/internal/coolify"
	"FaisalBudiono/coolify-env-fetcher/internal/mapper"
	"flag"
	"fmt"
	"os"
)

func main() {
	base := flag.String("base", "", "Base URL")
	accessToken := flag.String("access", "", "Coolify Access Token")
	projectID := flag.String("project", "", "Coolify Project ID")
	flag.Parse()

	if *base == "" || *accessToken == "" || *projectID == "" {
		fmt.Println("Should complete all parameters")
		flag.PrintDefaults()
		os.Exit(1)
	}

	res, err := coolify.ParseENV(*base, *projectID, *accessToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := fmt.Sprintf("out/%s", ".env")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := mapper.NewDotENV()

	err = m.WriteFile(f, res)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}