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

	gr.GET("/expeditions", r.getAll)
	gr.GET("/expedition/:location_id", r.getByLocationId)
	gr.POST("/expedition", r.create)
	gr.PATCH("/expedition/:id", r.updateDates)
	gr.DELETE("/expedition/:id", r.delete)
}

// @Summary	Show all expeditions
// @Description	return all expeditions
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /expeditions [get]
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

// @Summary	Show expeditions of specified location
// @Description	return expeditions of specified location
// @Tags member
// @Param token query string true "User authentication token"
// @Param location_id path string true "Id of location"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /expedition/{location_id} [get]
func (r *expeditionRoutes) getByLocationId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLocationId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	locationId, err := strconv.Atoi(ctx.Param("location_id"))
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLocationId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditions, err := r.expeditionService.GetLocationExpeditions(ctx, client, locationId)
	if err != nil {
		r.log.Errorf("expeditionRoutes getByLocationId: expeditionService.GetLocationExpeditions %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"expeditions": expeditions})
}

// @Summary	Add new expedition
// @Description	add new expedition
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateExpeditionInput true "Information about stored expedition"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /expedition [post]
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

// @Summary Update expedition dates by id
// @Description update expedition dates by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param input body newExpeditionDates true "New expedition dates"
// @Param id path int true	"Expedition id"
// @Accept json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expedition/{id} [patch]
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

// @Summary Delete expedition by id
// @Description delete expedition by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param id path string true "Expedition id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /expedition/{id} [delete]
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
