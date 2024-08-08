package main

import (
	"Sentinel/lib"
	"fmt"
	"os"
)

func main() {
	args := lib.CliParser()
	failHandler := &lib.VersionHandler{}
	httpClient := lib.ClientInit()
	localVersion := lib.GetCurrentLocalVersion(failHandler)
	repoVersion := lib.GetCurrentRepoVersion(httpClient, failHandler)
	fmt.Printf(" ===[ Sentinel, v%s ]===\n\n", localVersion)
	// Compare local and github repo versions
	lib.VersionCompare(repoVersion, localVersion)
	if err := lib.CreateOutputDir(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	lib.DisplayCount = 0
	if len(args.WordlistPath) == 0 {
		fmt.Println("[*] Using passive enum method")
		lib.PassiveEnum(&args, httpClient)
	} else {
		fmt.Println("[*] Using direct enum method")
		if err := lib.DirectEnum(&args, httpClient); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}
