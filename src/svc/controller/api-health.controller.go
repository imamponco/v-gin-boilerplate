package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	vgin_http "github.com/imamponco/v-gin-boilerplate/src/pkg/vhttp"
	"github.com/imamponco/v-gin-boilerplate/src/svc/contract"
	"github.com/imamponco/v-gin-boilerplate/src/svc/dto"
	"time"
)

type GetHealthController struct {
	handlerResponse vgin_http.ResponseHandler
	accessTime      time.Time
}

func NewGetHealthController(app *contract.AppContract, accessTime time.Time) *GetHealthController {
	// Init Get Health Controller
	h := GetHealthController{
		accessTime:      accessTime,
		handlerResponse: app.ResponseHandler,
	}

	return &h
}

// @Summary Get Health API
// @Description  Get Health API Check
// @Tags         Common
// @Accept       json
// @Produce json
// @Success 200 {object} dto.GetHealth_Result
// @Router / [get]
func (h *GetHealthController) GetHealthAPI(c *gin.Context) {
	h.handlerResponse.Success(c, dto.GetHealth_Result{
		AppVersion:     "dev:latest",
		BuildSignature: uuid.NewString(),
		Uptime:         time.Since(h.accessTime).String(),
		ServerTime:     time.Now(),
	})
}
