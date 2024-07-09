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

type leaderRoutes struct {
	leaderService service.Leader
	authService   service.Auth
	log           *logger.Logger
}

func newLeaderRoutes(gr *gin.RouterGroup, leaderService service.Leader, authService service.Auth, log *logger.Logger) {
	r := &leaderRoutes{
		leaderService: leaderService,
		authService:   authService,
		log:           log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/:expedition_id", r.getByExpeditionId)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.DELETE("/:id", r.delete)
}

func (r *leaderRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("leaderRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	leader, err := r.leaderService.GetLeaderById(ctx, client, id)
	if err != nil {
		r.log.Errorf("leaderRoutes getById: leaderService.GetLeaderById %v", err)
		if errors.Is(err, service.ErrLeaderNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"leader": leader})
}

func (r *leaderRoutes) getByExpeditionId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes getByExpeditionId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditionId, err := strconv.Atoi(ctx.Param("expedition_id"))
	if err != nil {
		r.log.Errorf("leaderRoutes getByExpeditionId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	leaders, err := r.leaderService.GetExpeditionLeaders(ctx, client, expeditionId)
	if err != nil {
		r.log.Errorf("leaderRoutes getByExpeditionId: leaderService.GetExpeditionLeaders %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"leaders": leaders})
}

func (r *leaderRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	leaders, err := r.leaderService.GetAllLeaders(ctx, client)
	if err != nil {
		r.log.Errorf("leaderRoutes getAll: leaderService.GetAllLeaders %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"leaders": leaders})
}

func (r *leaderRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateLeaderInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("leaderRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.leaderService.CreateLeader(ctx, client, &input)
	if err != nil {
		r.log.Errorf("leaderRoutes create: leaderService.CreateLeader %v", err)
		if errors.Is(err, service.ErrLeaderAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

func (r *leaderRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("leaderRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.leaderService.DeleteLeader(ctx, client, id)
	if err != nil {
		r.log.Errorf("leaderRoutes delete: leaderService.DeleteLeader %v", err)
		if errors.Is(err, service.ErrLeaderNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
