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
	gr.GET("/:leader_id", r.getByLeaderId)
	gr.GET("/:member_id", r.getByMemberId)
	gr.GET("/:curator_id", r.getByCuratorId)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.PATCH("/:id", r.update)
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

func (r *expeditionRoutes) getByLeaderId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLeaderId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	leaderId, err := strconv.Atoi(ctx.Param("leader_id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLeaderId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditions, err := r.expeditionService.GetLeaderExpeditions(ctx, client, leaderId)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLeaderId: expeditionService.GetLeaderExpeditions %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expeditions": expeditions})
}

func (r *expeditionRoutes) getByMemberId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByMemberId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	memberId, err := strconv.Atoi(ctx.Param("member_id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes getByMemberId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditions, err := r.expeditionService.GetMemberExpeditions(ctx, client, memberId)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLeaderId: expeditionService.GetMemberExpeditions %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expeditions": expeditions})
}

func (r *expeditionRoutes) getByCuratorId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByCuratorId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	curatorId, err := strconv.Atoi(ctx.Param("curator_id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes getByCuratorId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditions, err := r.expeditionService.GetCuratorExpeditions(ctx, client, curatorId)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByCuratorId: expeditionService.GetCuratorExpeditions %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expeditions": expeditions})
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

func (r *expeditionRoutes) update(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes update: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes update: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}
	// ???
	var input entity.UpdateExpeditionInput
	input.Id = id
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("expeditionRoutes update: dates error %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.expeditionService.UpdateExpeditionDates(ctx, client, &input)
	if err != nil {
		r.log.Errorf("expeditionRoutes update: expeditionService.UpdateExpeditionDates %v", err)
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
