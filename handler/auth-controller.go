package handler

import (
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/gofiber/fiber/v2"
)

//AuthController interface is a contract what this controller can do
type AuthHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
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

// Login is a function to get all books data from database
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 503
// @Router /v1/books [get]
func (c *authHandler) Login(ctx *fiber.Ctx) error {
	var loginDTO dto.LoginDTO
	if err := ctx.BodyParser(&loginDTO); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(model.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		return ctx.Status(http.StatusOK).JSON(response)

	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{})
	return ctx.Status(http.StatusUnauthorized).JSON(response)

}

// GetAllBooks is a function to get all books data from database
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 503
// @Router /v1/books [get]
func (c *authHandler) Register(ctx *fiber.Ctx) error {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.BodyParser(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return ctx.Status(http.StatusBadRequest).JSON(response)
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{})
		return ctx.Status(http.StatusConflict).JSON(response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		return ctx.Status(http.StatusCreated).JSON(response)
	}
}
