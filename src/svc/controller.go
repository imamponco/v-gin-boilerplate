package svc

import (
	"github.com/imamponco/v-gin-boilerplate/src/svc/controller"
	"time"
)

type Controllers struct {
	GetHealth *controller.GetHealthController
}

func NewController(a *App) *Controllers {
	return &Controllers{
		GetHealth: controller.NewGetHealthController(a.AppContract, time.Now()),
	}
}
