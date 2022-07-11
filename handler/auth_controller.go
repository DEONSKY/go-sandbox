package handler

import (
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/DEONSKY/go-sandbox/utils"
	"github.com/gofiber/fiber/v2"
)

//AuthController interface is a contract what this controller can do
type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type authHandler struct {
	jwtService service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(jwtService service.JWTService) AuthHandler {
	return &authHandler{
		jwtService: jwtService,
	}
}

/*
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}*/

// Login
// @Summary Login
// @Description Login
// @Tags auth
// @Accept json
// @Produce json
// @Param Auth body request.LoginRequest true "loginRequest"
// @Success 200 {object} helper.Response{data=model.User}
// @Failure 400 {object} helper.Response{data=helper.EmptyObj}
// @Router /api/auth/login [post]
func (c *authHandler) Login(ctx *fiber.Ctx) error {
	var loginDTO request.LoginRequest
	if err := ctx.BodyParser(&loginDTO); err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Request is incorrect", []string{})
	}

	errors := utils.ValidateStruct(loginDTO)
	if errors != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Validation error", errors)
	}

	authResult := service.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse("OK!", v)
		return ctx.Status(http.StatusOK).JSON(response)

	}
	return utils.ReturnErrorResponse(fiber.StatusUnauthorized, "Please check again your credential", []string{"Invalid Credential"})
}

// Register
// @Summary Register
// @Description Regişter
// @Tags auth
// @Accept json
// @Produce json
// @Param Auth body request.RegisterRequest true "registerRequest"
// @Success 200 {object} helper.Response{data=model.User}
// @Failure 400 {object} helper.Response{data=helper.EmptyObj}
// @Router /api/auth/register [post]
func (c *authHandler) Register(ctx *fiber.Ctx) error {
	var registerDTO request.RegisterRequest
	errDTO := ctx.BodyParser(&registerDTO)
	if errDTO != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Failed to process request", []string{errDTO.Error()})
	}

	errors := utils.ValidateStruct(registerDTO)
	if errors != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, "Validation error", errors)
	}

	if !service.IsDuplicateEmail(registerDTO.Email) {
		return utils.ReturnErrorResponse(fiber.StatusConflict, "Duplicate email", []string{})
	} else {
		createdUser := service.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse("OK!", createdUser)
		return ctx.Status(http.StatusCreated).JSON(response)
	}
}
