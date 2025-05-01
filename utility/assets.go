package utility

import (
	"os"
	"runtime"
)

func LoadAssetFrom(path string) string {
	_, isFlatpak := os.LookupEnv("container")

	var basePath string
	if isFlatpak && runtime.GOOS == "linux" {
		basePath = "/app/bin/assets/"
	} else {
		basePath = "assets/"
	}

	return basePath + path
}
