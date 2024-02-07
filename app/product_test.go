package app_test

import (
	"testing"

	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestAppProductIfBusinessRuleInFunctionEnabledIsWorking(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than 0 to enable product", err.Error())
}

func TestAppProductIfBusinessRuleInFunctionDisabledIsWorking(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Price must be zero greater than 0 in order to have product disabled", err.Error())
}

func TestAppProductIfBusinessRuleInFunctionIsValidIsWorking(t *testing.T) {
	product := app.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "Status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "Price must be greater than 0 or equal 0", err.Error())
}
