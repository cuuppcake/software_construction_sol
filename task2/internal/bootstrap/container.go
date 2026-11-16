package bootstrap

import (
	"fmt"
	"pos/internal/app"
	"pos/internal/domain"
	"pos/internal/export"
	"pos/internal/storage"
)

type Container struct {
	Service *app.Service
}

func NewContainer(policyName string) (*Container, error) {
	var policy domain.ResultPolicy
	switch policyName {
	case "weighted":
		policy = domain.WeightedPolicy{}
	case "equal":
		policy = domain.EqualPolicy{}
	default:
		return nil, fmt.Errorf("неизвестное правило расчёта итога %q", policyName)
	}
	store := storage.NewMemoryStore()
	service := app.NewService(store, store, store, export.CSVExporter{}, domain.NewCalculator(policy))
	return &Container{Service: service}, nil
}
