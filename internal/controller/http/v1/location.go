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

type locationRoutes struct {
	locationService service.Location
	authService     service.Auth
	log             *logger.Logger
}

func newLocationRoutes(gr *gin.RouterGroup, locationService service.Location, authService service.Auth, log *logger.Logger) {
	r := &locationRoutes{
		locationService: locationService,
		authService:     authService,
		log:             log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.DELETE("/:id", r.delete)
}

func (r *locationRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("locationRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	location, err := r.locationService.GetLocationById(ctx, client, id)
	if err != nil {
		r.log.Errorf("locationRoutes getById: locationService.GetLocationById %v", err)
		if errors.Is(err, service.ErrLocationNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"location": location})
}

func (r *locationRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	locations, err := r.locationService.GetAllLocations(ctx, client)
	if err != nil {
		r.log.Errorf("locationRoutes getAll: locationService.GetAllLocations %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"locations": locations})
}

func (r *locationRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateLocationInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("locationRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.locationService.CreateLocation(ctx, client, &input)
	if err != nil {
		r.log.Errorf("locationRoutes create: locationService.CreateLocation %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

func (r *locationRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("locationRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.locationService.DeleteLocation(ctx, client, id)
	if err != nil {
		r.log.Errorf("locationRoutes delete: locationService.DeleteLocation %v", err)
		if errors.Is(err, service.ErrLocationNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
