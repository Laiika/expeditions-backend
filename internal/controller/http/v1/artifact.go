package v1

import (
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
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

	gr.GET("/artifact/:location_id", r.getByLocationId)
	gr.GET("/artifacts", r.getAll)
	gr.POST("/artifact", r.create)
}

// @Summary	Show artifacts of specified location
// @Description	return all artifacts of specified location
// @Tags member
// @Param token query string true "User authentication token"
// @Param location_id path string true "Id of location"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /artifact/{location_id} [get]
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

// @Summary	Show all artifacts
// @Description	return all artifacts
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /artifacts [get]
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

// @Summary	Add new artifact
// @Description	add new artifact
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateArtifactInput true "Information about stored artifact"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /artifact [post]
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
