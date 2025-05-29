package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	webgames "webgames/app"
)

func main() {
	app := webgames.CreateApp()
	port := ":5000"
	url := fmt.Sprintf("http://localhost%s", port)

	srv := &http.Server{
		Addr:    port,
		Handler: app.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("starting server on %s\n", url)

	if runtime.GOOS == "windows" {
		go func() {
			time.Sleep(3 * time.Second)
			if err := exec.Command("cmd", "/C", "start", url).Run(); err != nil {
				log.Printf("failed to open '%s': %v\n", url, err)
			}
		}()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("server shutdown:", err)
	}
	<-ctx.Done()
	log.Println("server exiting")
}
