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

type locationRoutes struct {
	locationService service.Location
	authService     service.Auth
	log             *logger.Logger
}

func newLocationRoutes(gr *gin.RouterGroup, locationService service.Location, authService service.Auth, log *logger.Logger) {
	r := &locationRoutes{
		locationService: locationService,
		authService:     authService,
		log:             log,
	}

	gr.GET("/locations", r.getAll)
	gr.POST("/location", r.create)
	gr.DELETE("/location/:id", r.delete)
}

// @Summary	Show all locations
// @Description	return all locations
// @Tags member
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /locations [get]
func (r *locationRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	locations, err := r.locationService.GetAllLocations(ctx, client)
	if err != nil {
		r.log.Errorf("locationRoutes getAll: locationService.GetAllLocations %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"locations": locations})
}

// @Summary	Add new location
// @Description	add new location
// @Tags leader
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateLocationInput true "Information about stored location"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /location [post]
func (r *locationRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateLocationInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("locationRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.locationService.CreateLocation(ctx, client, &input)
	if err != nil {
		r.log.Errorf("locationRoutes create: locationService.CreateLocation %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// @Summary Delete location by id
// @Description delete location by id
// @Tags leader
// @Param token query string true "User authentication token"
// @Param id path string true "Location id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /location/{id} [delete]
func (r *locationRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("locationRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("locationRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.locationService.DeleteLocation(ctx, client, id)
	if err != nil {
		r.log.Errorf("locationRoutes delete: locationService.DeleteLocation %v", err)
		if errors.Is(err, service.ErrLocationNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
