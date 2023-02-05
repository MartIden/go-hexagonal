package main

import (
	"fmt"

	"github.com/MartIden/go-hexagonal/cmd"
	"github.com/MartIden/go-hexagonal/core/settings"
)

func main() {
	config, err := settings.GetAppSettings()
	if err != nil {
		return
	}
	fmt.Println(config)
	cmd.Start(config)
}