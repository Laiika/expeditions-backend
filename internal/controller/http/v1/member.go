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

type memberRoutes struct {
	memberService service.Member
	authService   service.Auth
	log           *logger.Logger
}

func newMemberRoutes(gr *gin.RouterGroup, memberService service.Member, authService service.Auth, log *logger.Logger) {
	r := &memberRoutes{
		memberService: memberService,
		authService:   authService,
		log:           log,
	}

	gr.GET("/member/:expedition_id", r.getByExpeditionId)
	gr.GET("/members", r.getAll)
	gr.POST("/member", r.create)
	gr.POST("/member_expedition", r.createMemberExpedition)
	gr.DELETE("/member/:id", r.delete)
}

// @Summary	Show members of specified expedition
// @Description	return all members of specified expedition
// @Tags member
// @Param token query string true "User authentication token"
// @Param expedition_id path string true "Id of expedition"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /member/{expedition_id} [get]
func (r *memberRoutes) getByExpeditionId(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("memberRoutes getByExpeditionId: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	expeditionId, err := strconv.Atoi(ctx.Param("expedition_id"))
	if err != nil {
		r.log.Errorf("memberRoutes getByExpeditionId: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	members, err := r.memberService.GetExpeditionMembers(ctx, client, expeditionId)
	if err != nil {
		r.log.Errorf("memberRoutes getByExpeditionId: memberService.GetExpeditionMembers %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"members": members})
}

// @Summary	Show all members
// @Description	return all members
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /members [get]
func (r *memberRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("memberRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	members, err := r.memberService.GetAllMembers(ctx, client)
	if err != nil {
		r.log.Errorf("memberRoutes getAll: memberService.GetAllMembers %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"members": members})
}

// @Summary	Add new member
// @Description	add new member
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateMemberInput true "Information about stored member"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /member [post]
func (r *memberRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("memberRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateMemberInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("memberRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.memberService.CreateMember(ctx, client, &input)
	if err != nil {
		r.log.Errorf("memberRoutes create: memberService.CreateMember %v", err)
		if errors.Is(err, service.ErrMemberAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

type newMemberExpedition struct {
	MemberId     int `json:"member_id"`
	ExpeditionId int `json:"expedition_id"`
}

// @Summary	Add member to expedition
// @Description	add member to expedition
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body newMemberExpedition true "Information about member and expedition"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /member_expedition [post]
func (r *memberRoutes) createMemberExpedition(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("memberRoutes createMemberExpedition: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input newMemberExpedition
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("memberRoutes createMemberExpedition: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.memberService.CreateMemberExpedition(ctx, client, input.MemberId, input.ExpeditionId)
	if err != nil {
		r.log.Errorf("memberRoutes createMemberExpedition: memberService.CreateMemberExpedition %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// @Summary Delete member by id
// @Description delete member by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param id path string true "Member id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /member/{id} [delete]
func (r *memberRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("memberRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("memberRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.memberService.DeleteMember(ctx, client, id)
	if err != nil {
		r.log.Errorf("memberRoutes delete: memberService.DeleteMember %v", err)
		if errors.Is(err, service.ErrMemberNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
