package bootstrap

import (
	"fmt"
	"pos/internal/app"
	"pos/internal/calculation"
	"pos/internal/export"
	"pos/internal/storage"
)

type Container struct {
	Service *app.Service
}

func NewContainer(policyName string) (*Container, error) {
	var policy calculation.ResultPolicy
	switch policyName {
	case "weighted":
		policy = calculation.WeightedPolicy{}
	case "equal":
		policy = calculation.EqualPolicy{}
	default:
		return nil, fmt.Errorf("неизвестное правило расчёта итога %q", policyName)
	}
	store := storage.NewMemoryStore()
	service := app.NewService(store, store, store, export.CSVExporter{}, calculation.NewCalculator(policy))
	return &Container{Service: service}, nil
}
