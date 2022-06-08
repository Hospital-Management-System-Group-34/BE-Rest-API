package routes

import (
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/handler"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/delivery/http/echo/middleware"
	repository "github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/repository/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/bcrypt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/jwt"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/service/postgres"
	"github.com/Hospital-Management-System-Group-34/BE-Rest-API/src/usecase/authentication"
	"github.com/labstack/echo/v4"
)

func AuthenticationRoutes(e *echo.Echo) {
	postgresDB := postgres.Connect()

	authenticationRepository := repository.NewAuthenticationRepository(postgresDB)
	staffRepository := repository.NewStaffRepository(postgresDB)

	bcryptPasswordHash := bcrypt.NewBcryptPasswordHash()
	jwtTokenManager := jwt.NewJWTTokenManager()

	staffLoginUseCase := authentication.NewStaffLoginUseCase(
		staffRepository,
		bcryptPasswordHash,
		jwtTokenManager,
		authenticationRepository,
	)
	staffLogoutUsecase := authentication.NewStaffLogoutUseCase(authenticationRepository)
	updateAuthenticationUseCase := authentication.NewUpdateAuthenticationUseCase(
		jwtTokenManager,
		authenticationRepository,
	)

	authenticationHandler := handler.NewAuthenticationHandler(
		staffLoginUseCase,
		staffLogoutUsecase,
		updateAuthenticationUseCase,
	)

	e.POST("/login", authenticationHandler.PostStaffLoginHandler)
	e.POST("/logout", authenticationHandler.PostStaffLogoutHandler, middleware.JWTMiddleware())
	e.PUT("/authentications", authenticationHandler.PutAuthenticationHandler)
}
