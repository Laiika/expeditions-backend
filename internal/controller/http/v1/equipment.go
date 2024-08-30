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

	gr.GET("/equipment/:expedition_id", r.getByExpeditionId)
	gr.GET("/equipments", r.getAll)
	gr.POST("/equipment", r.create)
	gr.DELETE("/equipment/:id", r.delete)
}

// @Summary	Show equipment of specified expedition
// @Description	return all equipment of specified expedition
// @Tags member
// @Param token query string true "User authentication token"
// @Param expedition_id path string true "Id of expedition"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /equipment/{expedition_id} [get]
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

// @Summary	Show all equipment
// @Description	return all equipment
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /equipments [get]
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

// @Summary	Add new equipment
// @Description	add new equipment
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateEquipmentInput true "Information about stored equipment"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /equipment [post]
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

// @Summary Delete equipment by id
// @Description delete equipment by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param id path string true "Equipment id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /equipment/{id} [delete]
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
