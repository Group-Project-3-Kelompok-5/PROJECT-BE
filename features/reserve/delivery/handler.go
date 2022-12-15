package delivery

import (
	"be13/project/features/reserve"
	"be13/project/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
)

type resDelivery struct {
	resUsecase reserve.ServiceEntities
}

func New(e *echo.Echo, service reserve.ServiceEntities) {
	handler := &resDelivery{
		resUsecase: service,
	}
	e.POST("/reserves", handler.PostData, middlewares.JWTMiddleware())
	e.GET("/reserves", handler.GetAll, middlewares.JWTMiddleware())

}

func (delivery *resDelivery) PostData(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	if idToken == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Failed Token Not Found",
		})
	}

	var dataCart Booking
	errBind := c.Bind(&dataCart)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error bind data",
		})
	}

	dataCart.UserID = uint(idToken)
	dataCart.Quantity = 1
	_, err := delivery.resUsecase.AddReserv(toCore(dataCart))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "error insert data",
		})
	}
	// if row == 2 {
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message": "product add to cart again",
	// 	})
	// } else if row == 3 {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "failed add to cart",
	// 	})
	// }
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "succes add to reservation",
	})
}

func (delivery *resDelivery) GetAll(c echo.Context) error {
	idToken := middlewares.ExtractTokenUserId(c)
	data, err := delivery.resUsecase.GetByToken(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed get cart",
		})
	} else if len(data) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "you dont have product in cart",
		})
	}

	dataRes, totalRes := fromCoreList(data)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Your Reserve",
		"res":     dataRes,
		"total":   totalRes,
	})
}
