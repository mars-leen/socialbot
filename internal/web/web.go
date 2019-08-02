package web

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"socialbot/internal/common/app"
	"socialbot/internal/web/setting"
	"socialbot/internal/web/wblogger"
	"time"
)

func Run() {
	// run web server
	cfg := setting.Cfg.Server
	gin.SetMode(cfg.Mode)
	handler := gin.New()
	handler = RegisterRouter(handler)

	server := &http.Server{
		Addr:         cfg.Listen,
		Handler:      handler,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		/*TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},*/
	}

	/*server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}*/
	//go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

	//log.Fatal(server.ListenAndServeTLS("", ""))

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			wblogger.Log.Error(err)
		}
	}()

	app.RegShutdownCallback(func() {
		fmt.Println("[INFO] shutdown http server start")
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.ShutdownTimeout)*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("[ERROR] shutdown http server error %v", errors.WithStack(err))
		}
		fmt.Println("[INFO] shutdown http server end")
	})

}
