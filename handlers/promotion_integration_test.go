//go:build integration

package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/repositories"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {
	amount := 100
	expected := 80

	promoRepo := repositories.NewPromotionRepositoryMock()
	promoRepo.On("GetPromotion").Return(repositories.Promotion{
		ID:              1,
		Purchasemin:     100,
		DiscountPercent: 20,
	}, nil)

	prromoService := services.NewPromotionService(promoRepo)
	promoHandler := handlers.NewPromotionHanler(prromoService)

	//http://localhost:8000/calculate?amount=100
	app := fiber.New()
	app.Get("/calculate", promoHandler.CalculateDiscount)

	req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

	//Act
	res, _ := app.Test(req)
	defer res.Body.Close()

	//Assert
	if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
		body, _ := io.ReadAll(res.Body)
		assert.Equal(t, strconv.Itoa(expected), string(body))
	}

}
