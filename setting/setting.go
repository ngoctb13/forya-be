package setting

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ngoctb13/forya-be/config"
	"github.com/ngoctb13/forya-be/infra"
	"go.uber.org/zap"
)

const (
	migrationFile = "file://./migrations/sql"
)

func ConnectDatabase(cfg *config.PostgresConfig) {
	infra.CreateDBAndMigrate(cfg, migrationFile)
}

func WaitOSSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	zap.S().Infof("Receive %v signal", s.String())
}
