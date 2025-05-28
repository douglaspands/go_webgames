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

	srv := &http.Server{
		Addr:    port,
		Handler: app.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("starting server on http://localhost%s\n", port)

	if runtime.GOOS == "windows" {
		go func() {
			time.Sleep(3 * time.Second)
			exec.Command(fmt.Sprintf("powershell.exe -c 'start http://localhost%s'", port))
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
	log.Println("timeout of 5 seconds.")
	log.Println("server exiting")
}
