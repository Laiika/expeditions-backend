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

type adminRoutes struct {
	adminService service.Admin
	authService  service.Auth
	log          *logger.Logger
}

func newAdminRoutes(gr *gin.RouterGroup, adminService service.Admin, authService service.Auth, log *logger.Logger) {
	r := &adminRoutes{
		adminService: adminService,
		authService:  authService,
		log:          log,
	}

	gr.GET("/admins", r.getAll)
	gr.POST("/admin", r.create)
	gr.DELETE("/admin/:id", r.delete)
}

// @Summary	Show all admins
// @Description	return all admins
// @Tags admin
// @Param token query string true "User authentication token"
// @Produce	json
// @Success	200 {object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /admins [get]
func (r *adminRoutes) getAll(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("adminRoutes getAll: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	admins, err := r.adminService.GetAllAdmins(ctx, client)
	if err != nil {
		r.log.Errorf("adminRoutes getAll: adminService.GetAllAdmins %v", err)
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"admins": admins})
}

// @Summary	Add new admin
// @Description	add new admin
// @Tags admin
// @Param token	query string true "User authentication token"
// @Param input body entity.CreateAdminInput true "Information about stored admin"
// @Accept json
// @Success	201
// @Failure	400	{object} map[string]interface{}
// @Failure	500	{object} map[string]interface{}
// @Router /admin [post]
func (r *adminRoutes) create(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("adminRoutes create: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	var input entity.CreateAdminInput
	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		r.log.Errorf("adminRoutes create: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := r.adminService.CreateAdmin(ctx, client, &input)
	if err != nil {
		r.log.Errorf("adminRoutes create: adminService.CreateAdmin %v", err)
		if errors.Is(err, service.ErrAdminAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, map[string]interface{}{"Id": id})
}

// @Summary Delete admin by id
// @Description delete admin by id
// @Tags admin
// @Param token query string true "User authentication token"
// @Param id path string true "Admin id"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /admin/{id} [delete]
func (r *adminRoutes) delete(ctx *gin.Context) {
	token := ctx.Query("token")
	client, err := r.authService.GetClient(token)
	if err != nil {
		r.log.Errorf("adminRoutes delete: authService.GetClient %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		r.log.Errorf("adminRoutes delete: Atoi id %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	err = r.adminService.DeleteAdmin(ctx, client, id)
	if err != nil {
		r.log.Errorf("adminRoutes delete: adminService.DeleteAdmin %v", err)
		if errors.Is(err, service.ErrAdminNotFound) {
			ctx.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
