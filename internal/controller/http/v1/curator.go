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

type curatorRoutes struct {
	curatorService service.Curator
	authService    service.Auth
	log            *logger.Logger
}

func newCuratorRoutes(gr *gin.RouterGroup, curatorService service.Curator, authService service.Auth, log *logger.Logger) {
	r := &curatorRoutes{
		curatorService: curatorService,
		authService:    authService,
		log:            log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/:expedition_id", r.getByExpeditionId)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.DELETE("/:id", r.delete)
}

func (r *curatorRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("curatorRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	curator, err := r.curatorService.GetCuratorById(ctx, client, id)
	if err != nil {
		r.log.Errorf("curatorRoutes getById: curatorService.GetCuratorById %v", err)
		if errors.Is(err, service.ErrCuratorNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"curator": curator})
}

func (r *curatorRoutes) getByExpeditionId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes getByExpeditionId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditionId, err := strconv.Atoi(ctx.Param("expedition_id"))
	if err != nil {
		r.log.Errorf("curatorRoutes getByExpeditionId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	curators, err := r.curatorService.GetExpeditionCurators(ctx, client, expeditionId)
	if err != nil {
		r.log.Errorf("curatorRoutes getByExpeditionId: curatorService.GetExpeditionCurators %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"curators": curators})
}

func (r *curatorRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	curators, err := r.curatorService.GetAllCurators(ctx, client)
	if err != nil {
		r.log.Errorf("curatorRoutes getAll: curatorService.GetAllCurators %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"curators": curators})
}

func (r *curatorRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateCuratorInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("curatorRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.curatorService.CreateCurator(ctx, client, &input)
	if err != nil {
		r.log.Errorf("curatorRoutes create: curatorService.CreateCurator %v", err)
		if errors.Is(err, service.ErrCuratorAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

func (r *curatorRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("curatorRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.curatorService.DeleteCurator(ctx, client, id)
	if err != nil {
		r.log.Errorf("curatorRoutes delete: curatorService.DeleteCurator %v", err)
		if errors.Is(err, service.ErrCuratorNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
