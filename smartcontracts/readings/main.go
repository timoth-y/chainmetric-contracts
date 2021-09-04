package main

import (
	"syscall"

	"github.com/timoth-y/chainmetric-network/smartcontracts/shared/core"
	"github.com/ztrue/shutdown"
)

func init() {
	core.InitCore()
	core.InitLevelDB()
}

func main() {
	go core.BootstrapChaincodeServer(NewReadingsContract())

	shutdown.Add(core.CloseLevelDB)
	shutdown.Listen(syscall.SIGINT, syscall.SIGTERM)
}
