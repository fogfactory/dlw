//go:build debug
// +build debug

package main

import (
    "github.com/fogfactory/dlw"
    "github.com/fogfactory/dlw/config"
)

func init() {
    runMain = dlw.Wrap(config.Default(),
        config.WithPort(1122),
        config.WaitDebugger())(runMain)
}
