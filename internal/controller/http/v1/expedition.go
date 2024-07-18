package v1

import (
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type expeditionRoutes struct {
	expeditionService service.Expedition
	authService       service.Auth
	log               *logger.Logger
}

func newExpeditionRoutes(gr *gin.RouterGroup, expeditionService service.Expedition, authService service.Auth, log *logger.Logger) {
	r := &expeditionRoutes{
		expeditionService: expeditionService,
		authService:       authService,
		log:               log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.PATCH("/:id", r.updateDates)
	gr.DELETE("/:id", r.delete)
}

func (r *expeditionRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expedition, err := r.expeditionService.GetExpeditionById(ctx, client, id)
	if err != nil {
		r.log.Errorf("expeditionRoutes getById: expeditionService.GetExpeditionById %v", err)
		if errors.Is(err, service.ErrExpeditionNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expedition": expedition})
}

func (r *expeditionRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditions, err := r.expeditionService.GetAllExpeditions(ctx, client)
	if err != nil {
		r.log.Errorf("expeditionRoutes getAll: expeditionService.GetAllExpeditions %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expeditions": expeditions})
}

func (r *expeditionRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateExpeditionInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("expeditionRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.expeditionService.CreateExpedition(ctx, client, &input)
	if err != nil {
		r.log.Errorf("expeditionRoutes create: expeditionService.CreateExpedition %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

type newExpeditionDates struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (r *expeditionRoutes) updateDates(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes updateDates: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes updateDates: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input newExpeditionDates
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("expeditionRoutes updateDates: dates error %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.expeditionService.UpdateExpeditionDates(ctx, client, id, input.StartDate, input.EndDate)
	if err != nil {
		r.log.Errorf("expeditionRoutes updateDates: expeditionService.UpdateExpeditionDates %v", err)
		if errors.Is(err, service.ErrExpeditionNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (r *expeditionRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.expeditionService.DeleteExpedition(ctx, client, id)
	if err != nil {
		r.log.Errorf("expeditionRoutes delete: expeditionService.DeleteExpedition %v", err)
		if errors.Is(err, service.ErrExpeditionNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
