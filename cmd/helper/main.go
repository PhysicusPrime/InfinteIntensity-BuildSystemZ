package main

import (
	"fmt"

	"github.com/username/firmware-builder/internal/build"
)

func main() {
	fmt.Println("Helper-Tool gestartet")
	build.SetupToolchain()
	build.BuildBusybox()
	build.BuildRootfs()
}
