package controller

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gitlab.com/kongrentian-group/tianyi/entity"
)

type LifecycleController interface {
	// listen for requests, blocks the goroutine
	Listen()
	// gracefully shutdown
	Shutdown(exitCode int)
	// gracefully shutdown when an interruption signal is sent
	// blocks the goroutine untill the signal is recieved
	// listens for os.Interrupt and syscall.SIGTERM
	ShutdownOnInterruptionSignal()
}

type lifecycleController struct {
	router       *fiber.App
	config       *entity.ConfigServer
	sessionStore *session.Store
}

func NewLifecycleContoller(router *fiber.App, sessionStore *session.Store) LifecycleController {
	return &lifecycleController{router: router, sessionStore: sessionStore}
}

func (controller *lifecycleController) Listen() {
	err := controller.router.Listen(controller.config.Port)
	if err == nil {
		return
	}
	log.Println("Runtime error:\n", err)
	controller.Shutdown(1)
}

func (controller *lifecycleController) Shutdown(code int) {
	log.Println("Shutting down...")
	errors := make([]error, 0)
	errors = append(errors, controller.router.Shutdown())
	errors = append(errors, controller.sessionStore.Storage.Close())
	for _, err := range errors {
		if err == nil {
			continue
		}
		log.Println("Shutdown error:\n", err)
	}
	if len(errors) != 0 {
		code = 1
	}
	os.Exit(code)
}

func (controller *lifecycleController) ShutdownOnInterruptionSignal() {
	// channel for the interruption signal
	channel := make(chan os.Signal, 1)
	// when an interruption or termination signal is sent, shutdown
	signal.Notify(channel, os.Interrupt, syscall.SIGTERM)
	// wait for it and shutdown
	<-channel
	controller.Shutdown(0)
}
