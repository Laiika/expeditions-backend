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

	gr.GET("/leader/:expedition_id", r.getByExpeditionId)
	gr.GET("/leaders", r.getAll)
	gr.POST("/leader", r.create)
	gr.POST("/leader_expedition", r.createLeaderExpedition)
	gr.DELETE("/leader/:id", r.delete)
}

// @Summary	Show leaders of specified expedition
// @Description	return all leaders of specified expedition
// @Tags member
// @Param token query string true "User authentication token"
// @Param expedition_id path string true "Id of expedition"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /leader/{expedition_id} [get]
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

// @Summary	Show all leaders
// @Description	return all leaders
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /leaders [get]
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

// @Summary	Add new leader
// @Description	add new leader
// @Tags admin
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateLeaderInput true "Information about stored leader"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /leader [post]
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

type newLeaderExpedition struct {
	LeaderId     int `json:"leader_id"`
	ExpeditionId int `json:"expedition_id"`
}

// @Summary	Add leader to expedition
// @Description	add leader to expedition
// @Tags admin
// @Param token	query string true "User authentication token"
// @Param input body newLeaderExpedition true "Information about leader and expedition"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /leader_expedition [post]
func (r *leaderRoutes) createLeaderExpedition(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("leaderRoutes createLeaderExpedition: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input newLeaderExpedition
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("leaderRoutes createLeaderExpedition: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.leaderService.CreateLeaderExpedition(ctx, client, input.LeaderId, input.ExpeditionId)
	if err != nil {
		r.log.Errorf("leaderRoutes createLeaderExpedition: leaderService.CreateLeaderExpedition %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// @Summary Delete leader by id
// @Description delete leader by id
// @Tags admin
// @Param token query string true "User authentication token"
// @Param id path string true "Leader id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /leader/{id} [delete]
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
