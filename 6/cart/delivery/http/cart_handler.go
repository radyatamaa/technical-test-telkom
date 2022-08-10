package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"technical-test-telkom/cart"
	"technical-test-telkom/helper/logger"
	"technical-test-telkom/models"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CartHandler  represent the http handler for country
type CartHandler struct {
	CartUseCase cart.UseCase
	Log         logger.Logger
}

// NewsampleModuleHandler will initialize the countrys/ resources endpoint
func NewCartHandler(gapis *echo.Group,e *echo.Echo, us cart.UseCase, log logger.Logger) {
	handler := &CartHandler{
		Log:         log,
		CartUseCase: us,
	}

	e.GET("/", handler.Documentation)
	gapis.GET("/cart", handler.List)
	gapis.POST("/cart", handler.Create)
	gapis.DELETE("/cart/:id", handler.Delete)
}
func (a *CartHandler) Documentation(c echo.Context) error{
	http.Redirect(c.Response(), c.Request(), "/swagger/index.html", 301)
	return nil
}
// List godoc
// @Summary List.
// @Description List.
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CartProduk
// @Router /api/cart [get]
func (a *CartHandler) List(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	response, err := a.CartUseCase.Find(ctx)
	if err != nil {
		a.Log.Error("cart.delivery.http.CartHandler.List: %s , line : %s", err.Error, log.Lshortfile)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response)
}

// Create godoc
// @Summary Create.
// @Description Create.
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param request body models.CartProduk true "request"
// @Success 200 {object} interface{}
// @Router /api/cart [post]
func (a *CartHandler) Create(c echo.Context) error {
	request := new(models.CartProduk)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err := a.CartUseCase.Store(ctx, *request)
	if err != nil {
		a.Log.Error("cart.delivery.http.CartHandler.List: %s , line : %s", err.Error, log.Lshortfile)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, request)
}


// Delete godoc
// @Summary Delete.
// @Description Delete.
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param id path string true "kode produk"
// @Success 200 {object} interface{}
// @Router /api/cart/{id} [delete]
func (a *CartHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := a.CartUseCase.Delete(ctx, c.Param("id"))
	if err != nil {
		a.Log.Error("cart.delivery.http.CartHandler.List: %s , line : %s", err.Error, log.Lshortfile)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}