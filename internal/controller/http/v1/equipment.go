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

type equipmentRoutes struct {
	equipmentService service.Equipment
	authService      service.Auth
	log              *logger.Logger
}

func newEquipmentRoutes(gr *gin.RouterGroup, equipmentService service.Equipment, authService service.Auth, log *logger.Logger) {
	r := &equipmentRoutes{
		equipmentService: equipmentService,
		authService:      authService,
		log:              log,
	}

	gr.GET("/:id", r.getById)
	gr.GET("/:expedition_id", r.getByExpeditionId)
	gr.GET("/", r.getAll)
	gr.POST("/", r.create)
	gr.DELETE("/:id", r.delete)
}

func (r *equipmentRoutes) getById(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("equipmentRoutes getById: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("equipmentRoutes getById: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	equipment, err := r.equipmentService.GetEquipmentById(ctx, client, id)
	if err != nil {
		r.log.Errorf("equipmentRoutes getById: equipmentService.GetEquipmentById %v", err)
		if errors.Is(err, service.ErrEquipmentNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"equipment": equipment})
}

func (r *equipmentRoutes) getByExpeditionId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("equipmentRoutes getByExpeditionId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditionId, err := strconv.Atoi(ctx.Param("expedition_id"))
	if err != nil {
		r.log.Errorf("equipmentRoutes getByExpeditionId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	equipments, err := r.equipmentService.GetExpeditionEquipments(ctx, client, expeditionId)
	if err != nil {
		r.log.Errorf("equipmentRoutes getByExpeditionId: equipmentService.GetExpeditionEquipments %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"equipments": equipments})
}

func (r *equipmentRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("equipmentRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	equipments, err := r.equipmentService.GetAllEquipments(ctx, client)
	if err != nil {
		r.log.Errorf("equipmentRoutes getAll: equipmentService.GetAllEquipments %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"equipments": equipments})
}

func (r *equipmentRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("equipmentRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateEquipmentInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("equipmentRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.equipmentService.CreateEquipment(ctx, client, &input)
	if err != nil {
		r.log.Errorf("equipmentRoutes create: equipmentService.CreateEquipment %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

func (r *equipmentRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("equipmentRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("equipmentRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.equipmentService.DeleteEquipment(ctx, client, id)
	if err != nil {
		r.log.Errorf("equipmentRoutes delete: equipmentService.DeleteEquipment %v", err)
		if errors.Is(err, service.ErrEquipmentNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
