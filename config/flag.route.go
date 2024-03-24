package config

import (
	"flag"
	"os"
)

func RouteFlag() (*flag.FlagSet, *string, *string) {
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)
	withFlag := installCmd.String("w", "", "Specify package in format username/repo (e.g., xodity/deviter)")
	withPath := installCmd.String("d", "", "Specify path for the package (optional)")
	installCmd.Parse(os.Args[2:])

	return installCmd, withFlag, withPath
}
