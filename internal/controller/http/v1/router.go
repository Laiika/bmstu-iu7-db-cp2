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
	newAuthRoutes(auth, services.Auth)

	authMiddleware := &AuthMiddleware{
		services.Auth,
		log,
	}
	withAuth := mainGroup.Group("", authMiddleware.SessionCheck())
	{
		newLeaderRoutes(withAuth.Group("/leaders"), services.Leader, services.Auth, log)
		newMemberRoutes(withAuth.Group("/members"), services.Member, services.Auth, log)
		newCuratorRoutes(withAuth.Group("/curators"), services.Curator, services.Auth, log)
		newLocationRoutes(withAuth.Group("/locations"), services.Location, services.Auth, log)
		newExpeditionRoutes(withAuth.Group("/expeditions"), services.Expedition, services.Auth, log)
		newArtifactRoutes(withAuth.Group("/artifacts"), services.Artifact, services.Auth, log)
		newEquipmentRoutes(withAuth.Group("/equipments"), services.Equipment, services.Auth, log)
	}
}
