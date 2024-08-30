package v1

import (
	"db_cp_6/internal/entity"
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authRoutes struct {
	authService service.Auth
	log         *logger.Logger
}

func newAuthRoutes(gr *gin.RouterGroup, authService service.Auth, log *logger.Logger) {
	r := &authRoutes{
		authService: authService,
		log:         log,
	}

	gr.POST("/login", r.login)
	gr.POST("/logout", r.logout)
}

// @Summary	Log in to the server
// @Description	log in to the server
// @Tags common
// @Param data body	entity.LoginInput true "Authentication request"
// @Accept json
// @Produce	json
// @Success	200	{object} map[string]interface{}
// @Failure	400	{object} map[string]interface{}
// @Router /login [post]
func (r *authRoutes) login(ctx *gin.Context) {
	var data entity.LoginInput

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		r.log.Errorf("authRoutes login: %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	token, err := r.authService.Login(ctx, &data)
	if err != nil {
		r.log.Errorf("authRoutes login: authService.Login %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"Token": token})
}

// @Summary	Log out from the server
// @Description	log out from the server
// @Tags common
// @Param token query string true "User authentication token"
// @Success	200
// @Failure	400	{object} map[string]interface{}
// @Router /logout [post]
func (r *authRoutes) logout(ctx *gin.Context) {
	token := ctx.Query("token")

	err := r.authService.Logout(token)
	if err != nil {
		r.log.Errorf("authRoutes logout: authService.Logout %v", err)
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
