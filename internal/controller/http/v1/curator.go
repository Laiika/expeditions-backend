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

	gr.GET("/curator/:expedition_id", r.getByExpeditionId)
	gr.GET("/curators", r.getAll)
	gr.POST("/curator", r.create)
	gr.POST("/curator_expedition", r.createCuratorExpedition)
	gr.DELETE("/curator/:id", r.delete)
}

// @Summary	Show curators of specified expedition
// @Description	return all curators of specified expedition
// @Tags member
// @Param token query string true "User authentication token"
// @Param expedition_id path string true "Id of expedition"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /curator/{expedition_id} [get]
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

// @Summary	Show all curators
// @Description	return all curators
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /curators [get]
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

// @Summary	Add new curator
// @Description	add new curator
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateCuratorInput true "Information about stored curator"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /curator [post]
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

type newCuratorExpedition struct {
	CuratorId    int `json:"curator_id"`
	ExpeditionId int `json:"expedition_id"`
}

// @Summary	Add curator to expedition
// @Description	add curator to expedition
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body newCuratorExpedition true "Information about curator and expedition"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /curator_expedition [post]
func (r *curatorRoutes) createCuratorExpedition(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("curatorRoutes createCuratorExpedition: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input newCuratorExpedition
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("curatorRoutes createCuratorExpedition: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.curatorService.CreateCuratorExpedition(ctx, client, input.CuratorId, input.ExpeditionId)
	if err != nil {
		r.log.Errorf("curatorRoutes createCuratorExpedition: curatorService.CreateCuratorExpedition %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// @Summary Delete curator by id
// @Description delete curator by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param id path string true "Curator id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /curator/{id} [delete]
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
