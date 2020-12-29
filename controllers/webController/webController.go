package webController

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/vivkpatl/bread-puns/controllers/webController/impl"
)

type WebController struct {
	address string
	engine  *gin.Engine
	server  http.Server
	stopCh  chan os.Signal
}

func NewWebController(port string, sigClose chan os.Signal) *WebController {
	wc := WebController{}

	wc.address = fmt.Sprint(":", port)
	wc.stopCh = sigClose
	wc.engine = gin.Default()
	wc.server = http.Server{
		Addr:    wc.address,
		Handler: wc.engine,
	}
	wc.SetupRoutes()

	return &wc
}

/*
Setup basic REST API routes for HTTP handler
*/
func (wc *WebController) SetupRoutes() {
	wc.engine.GET("/pun", impl.GetPun)

	// TODO: More API stuff for registration???? SNS stuff too? Cron procedure?
}

func (wc *WebController) Start() {
	go func() {
		err := wc.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			// TODO: Proper error handling
			panic(err.Error())
		}
	}()

	go func() {
		for {
			select {
			case <-wc.stopCh:
				wc.Shutdown()
			}
		}
	}()
}

func (wc *WebController) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := wc.server.Shutdown(ctx)
	if err != nil {
		log.Fatal("WebController could not gracefully shutdown!", err)
	}

	log.Println("WebController successfully shut down")
}
