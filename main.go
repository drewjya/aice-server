package main

import (
	"aice-server/config"
	"aice-server/controller"
	"aice-server/dto"
	"aice-server/middleware"
	"aice-server/repository"
	"aice-server/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

var (
	db                  *gorm.DB                       = config.SetupDatabaseConfig()
	userRepo            repository.UserRepository      = repository.NewUserRepository(db)
	transaksiRepo       repository.TransaksiRepository = repository.NewTransaksiRepository(db)
	authService         services.AuthService           = services.NewAuthService(userRepo)
	transaksiService    services.TransaksiService      = services.NewTransaksiService(transaksiRepo)
	jwtService          services.JWTService            = services.NewJWTService()
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	gin.SetMode(gin.ReleaseMode)
	validate := validator.New()

	validate.RegisterStructValidation(dto.FotoNormalDTOValidation, dto.FotoNormalDTO{})
	validate.RegisterStructValidation(dto.FotoSuperhyperDTOValidation, dto.FotoSuperHyperDTO{})

	r := gin.Default()
	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}
	midle := middleware.Authorization(jwtService)

	transactionROute := r.Group("api/transaksi").Use(midle)
	{
		transactionROute.GET("/detail/:id", transaksiController.GetTransactionHistoryDetail)
		transactionROute.GET("/historyToday", transaksiController.GetTransaksiHistoryToday)
		transactionROute.GET("/historyThisWeek", transaksiController.GetTransaksiHistoryTHisWeek)
		transactionROute.POST("/normal", func(ctx *gin.Context) {
			transaksiController.TambahTransaksiNormal(ctx, validate)
		})
		transactionROute.POST("/superHyper", transaksiController.TambahTransaksiSuperHyper)
	}

	r.Static("/images", os.Getenv("OUTPUT_IMAGE"))

	r.Run()
}
