package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosscode/easy-user/infrastructure/logger"
)

type (
	Port int64

	Server interface {
		Listen(ctx context.Context, wg *sync.WaitGroup)
	}

	webServer struct {
		r       GinRouter
		port    int64
		ctx     context.Context
		timeout time.Duration
	}
)

func NewWebServer(
	r GinRouter,
	port int64,
	timeout time.Duration,
) *webServer {
	return &webServer{r: r, port: port, timeout: timeout}
}

func (w *webServer) Listen(ctx context.Context, wg *sync.WaitGroup) {
	gin.SetMode(gin.ReleaseMode)

	w.r.SetAppHandlers()

	ctx, cancel := context.WithCancel(ctx)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", w.port),
		Handler:      w.r.GetRouter(),
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-stop
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("❌ Erro ao iniciar o servidor: %v", err)
		}
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), w.timeout)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Fatal("❌ Erro ao encerrar o servidor: %v", err)
		} else {
			logger.Info("✅ Servidor encerrado com sucesso.")
		}
	}()
}
