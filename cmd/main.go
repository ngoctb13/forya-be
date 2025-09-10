package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ngoctb13/forya-be/config"
	"github.com/ngoctb13/forya-be/infra"
	"github.com/ngoctb13/forya-be/infra/repos"
	classUC "github.com/ngoctb13/forya-be/internal/domains/class/usecases"
	userUC "github.com/ngoctb13/forya-be/internal/domains/user/usecases"
	"github.com/ngoctb13/forya-be/server"
	"github.com/ngoctb13/forya-be/setting"
	"go.uber.org/zap"
)

func main() {
	var configFile, port string
	flag.StringVar(&configFile, "config-file", "", "Specify config file path")
	flag.StringVar(&port, "port", "", "Specify port")
	flag.Parse()

	defer setting.WaitOSSignal()

	cfg, err := config.Load(configFile)
	if err != nil {
		zap.S().Errorf("Error loading config: %v", err)
		panic(err)
	}

	go setting.ConnectDatabase(cfg.DB)

	db, err := infra.InitPostgres(cfg.DB)
	rp := repos.NewSQLRepo(db, cfg.DB)
	userRepo := rp.Users()
	classRepo := rp.Classes()
	userUsecases := userUC.NewUser(userRepo)
	classUsescases := classUC.NewClass(classRepo)

	s := server.NewServer(cfg)
	domains := &server.Domains{
		User:  userUsecases,
		Class: classUsescases,
	}
	s.InitCORS()
	s.InitRouter(domains)

	serverErr := make(chan error, 1)
	go func() {
		if err := s.ListenHTTP(); err != nil {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	select {
	case <-stop:
		zap.S().Info("Received shutdown signal")
	case err := <-serverErr:
		zap.S().Errorf("Server error: %v", err)
	}

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		zap.S().Errorf("Failed to shutdown server: %v", err)
	} else {
		zap.S().Info("Server shutdown completed")
	}
}
