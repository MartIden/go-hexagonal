package main

import (
	"github.com/MartIden/go-hexagonal/cmd"
	"github.com/MartIden/go-hexagonal/core/settings"
)

func main() {
	if appSettings, err := settings.GetAppSettings(); err == nil {
		cmd.Start(appSettings)
	}
}