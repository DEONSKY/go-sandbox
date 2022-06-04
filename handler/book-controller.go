package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DEONSKY/go-sandbox/dto/request"
	"github.com/DEONSKY/go-sandbox/helper"
	"github.com/DEONSKY/go-sandbox/model"
	"github.com/DEONSKY/go-sandbox/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

//BookController is a ...
type BookHandler interface {
	All(context *fiber.Ctx) error
	FindByID(context *fiber.Ctx) error
	Insert(context *fiber.Ctx) error
	Update(context *fiber.Ctx) error
	Delete(context *fiber.Ctx) error
}

type bookHandler struct {
	bookService service.BookService
	jwtService  service.JWTService
}

//NewBookController create a new instances of BoookController
func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookHandler {
	return &bookHandler{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
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
func (c *bookHandler) All(context *fiber.Ctx) error {
	var books []model.Book = c.bookService.All()
	res := helper.BuildResponse(true, "OK", books)
	return context.Status(http.StatusOK).JSON(res)
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
func (c *bookHandler) FindByID(context *fiber.Ctx) error {
	id, err := strconv.ParseUint(context.Params("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	}

	var book model.Book = c.bookService.FindByID(id)
	if (book == model.Book{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		return context.Status(http.StatusNotFound).JSON(res)
	} else {
		res := helper.BuildResponse(true, "OK", book)
		return context.Status(http.StatusOK).JSON(res)
	}
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
func (c *bookHandler) Insert(context *fiber.Ctx) error {
	var bookCreateDTO request.BookCreateRequest
	errDTO := context.BodyParser(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)
	} else {
		user := context.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		convertedUserID, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		return context.Status(http.StatusCreated).JSON(response)
	}
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
func (c *bookHandler) Update(context *fiber.Ctx) error {
	var bookUpdateDTO request.BookUpdateRequest
	errDTO := context.BodyParser(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(res)

	}

	authHeader := context.Request().Header.Peek("Authorization")
	token, errToken := c.jwtService.ValidateToken(string(authHeader))
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		return context.Status(http.StatusOK).JSON(response)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		return context.Status(http.StatusForbidden).JSON(response)
	}
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
func (c *bookHandler) Delete(context *fiber.Ctx) error {
	var book model.Book
	id, err := strconv.ParseUint(context.Params("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		return context.Status(http.StatusBadRequest).JSON(response)
	}
	book.ID = id
	authHeader := context.Request().Header.Peek("Authorization")
	token, errToken := c.jwtService.ValidateToken(string(authHeader))
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) {
		c.bookService.Delete(book)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		return context.Status(http.StatusOK).JSON(res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		return context.Status(http.StatusForbidden).JSON(response)
	}
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
func (c *bookHandler) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
