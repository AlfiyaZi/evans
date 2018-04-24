package main

import (
	"os"

	"github.com/ktr0731/evans/adapter/controller"
	"github.com/ktr0731/evans/meta"
)

func main() {
	os.Exit(run(controller.NewCLI(meta.AppName, meta.Version.String())))
}

func run(runnable controller.Runnable) int {
	return runnable.Run(os.Args[1:])
}
