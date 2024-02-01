package app_test

import (
	"testing"

	"github.com/Sotnasjeff/hexagonal-architecture-api/app"
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
