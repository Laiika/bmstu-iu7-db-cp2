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

type artifactRoutes struct {
	artifactService service.Artifact
	authService     service.Auth
	log             *logger.Logger
}

func newArtifactRoutes(gr *gin.RouterGroup, artifactService service.Artifact, authService service.Auth, log *logger.Logger) {
	r := &artifactRoutes{
		artifactService: artifactService,
		authService:     authService,
		log:             log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/:location_id", r.getByLocationId)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
}

func (r *artifactRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("artifactRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("artifactRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	artifact, err := r.artifactService.GetArtifactById(ctx, client, id)
	if err != nil {
		r.log.Errorf("artifactRoutes getById: artifactService.GetArtifactById %v", err)
		if errors.Is(err, service.ErrArtifactNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"artifact": artifact})
}

func (r *artifactRoutes) getByLocationId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("artifactRoutes getByLocationId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	locationId, err := strconv.Atoi(ctx.Param("location_id"))
	if err != nil {
		r.log.Errorf("artifactRoutes getByLocationId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	artifacts, err := r.artifactService.GetLocationArtifacts(ctx, client, locationId)
	if err != nil {
		r.log.Errorf("artifactRoutes getByLocationId: artifactService.GetLocationArtifacts %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"artifacts": artifacts})
}

func (r *artifactRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("artifactRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	artifacts, err := r.artifactService.GetAllArtifacts(ctx, client)
	if err != nil {
		r.log.Errorf("artifactRoutes getAll: artifactService.GetAllArtifacts %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"artifacts": artifacts})
}

func (r *artifactRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("artifactRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateArtifactInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("artifactRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.artifactService.CreateArtifact(ctx, client, &input)
	if err != nil {
		r.log.Errorf("artifactRoutes create: artifactService.CreateArtifact %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}
