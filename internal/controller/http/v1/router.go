package v1

import (
	"db_cp_6/internal/service"
	"db_cp_6/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(handler *gin.Engine, services *service.Services, log *logger.Logger) {
	gin.DisableConsoleColor()

	handler.Use(gin.LoggerWithWriter(log.Writer()))
	handler.Use(gin.RecoveryWithWriter(log.Writer()))

	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mainGroup := handler.Group("/api/v1")
	newAuthRoutes(mainGroup, services.Auth, log)

	authMiddleware := &AuthMiddleware{
		services.Auth,
		log,
	}
	withAuth := mainGroup.Group("", authMiddleware.SessionCheck())
	{
		newAdminRoutes(withAuth, services.Admin, services.Auth, log)
		newLeaderRoutes(withAuth, services.Leader, services.Auth, log)
		newMemberRoutes(withAuth, services.Member, services.Auth, log)
		newCuratorRoutes(withAuth, services.Curator, services.Auth, log)
		newLocationRoutes(withAuth, services.Location, services.Auth, log)
		newExpeditionRoutes(withAuth, services.Expedition, services.Auth, log)
		newArtifactRoutes(withAuth, services.Artifact, services.Auth, log)
		newEquipmentRoutes(withAuth, services.Equipment, services.Auth, log)
	}
}
