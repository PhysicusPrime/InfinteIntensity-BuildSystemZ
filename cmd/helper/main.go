package main

import (
	"fmt"

	"github.com/PhysicusPrime/InfinteIntensity-BuildSystemZ/internal/build"
)

func main() {
	fmt.Println("Helper-Tool gestartet")
	build.SetupToolchain()
	build.BuildBusybox()
	build.BuildRootfs()
}
