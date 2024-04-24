package controller

import (
	"clean-architecture/util"
	"net/http"

	"clean-architecture/dto/auth"
	"clean-architecture/pkg"
	"clean-architecture/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	VerifyToken(context *gin.Context)
	//RefreshToken(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JwtService
	logger      *logger.Logger
}

func NewAuthController(
	authService service.AuthService,
	jwtService service.JwtService,
	logger *logger.Logger,
) *authController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
		logger:      logger,
	}
}

func (controller *authController) Login(context *gin.Context) {
	var loginDto dto.Login
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("End of time")
		return
	}
	isValidCredential, userId := controller.authService.VerifyCredential(loginDto.Email, loginDto.Password)
	if !isValidCredential {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse("invalid credential"))
		controller.logger.Error().Msg("invalid credential")
		return
	}
	tokenPair, err := controller.jwtService.GenerateTokenPair(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
		return
	}
	context.JSON(http.StatusOK, util.GetResponse(tokenPair))
	return
}

func (controller *authController) Register(context *gin.Context) {
	var userDto dto.Register
	err := context.ShouldBindJSON(&userDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("Logger controller")
		return
	}

	if err := controller.authService.Register(userDto); err != nil {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	context.JSON(http.StatusOK, util.GetResponse([]int{}))
}

func (controller *authController) VerifyToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
		controller.logger.Error().Err(err).Msg("")
		return
	}
	token, _ := util.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(
			http.StatusBadRequest, util.GetErrorResponse("Invalid Token"))
		controller.logger.Error().Msg("Invalid Token")
		return
	}
	context.JSON(http.StatusOK, util.GetResponse(gin.H{"is_valid": true}))
}

//func (controller *authController) RefreshToken(context *gin.Context) {
//	tokenDto := dto.Token{}
//	if err := context.ShouldBindJSON(&tokenDto); err != nil {
//		context.JSON(http.StatusBadRequest, util.GetErrorResponse(err.Error()))
//		controller.logger.Error().Err(err).Msg("")
//		return
//	}
//	token, err := util.ValidateToken(tokenDto.Token)
//	if token == nil || !token.Valid {
//		context.AbortWithStatusJSON(
//			http.StatusBadRequest, util.GetErrorResponse(err.Error()))
//		controller.logger.Error().Err(err).Msg("")
//		return
//	}
//	if claims, ok := token.Claims.(jwt.MapClaims); ok {
//		context.JSON(
//			http.StatusOK, controller.jwtService.GenerateTokenPair(claims["user_id"]))
//	} else {
//		context.AbortWithStatusJSON(
//			http.StatusBadRequest, util.GetErrorResponse("Failed to claim token"))
//		controller.logger.Error().Msg("Failed to claim token")
//	}
//}
