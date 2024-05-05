package services

import (
	"context"
	"github.com/dscamargo/template_go_w_cli/internal/models"
	"github.com/dscamargo/template_go_w_cli/internal/repositories"
)

type ExampleService struct {
	exampleRepository *repositories.ExampleRepository
}

func NewExampleService(socioHistoryRepo *repositories.ExampleRepository) *ExampleService {
	return &ExampleService{socioHistoryRepo}
}

func (svc *ExampleService) GetAll(ctx context.Context) ([]models.Example, error) {
	return svc.exampleRepository.GetAll(ctx)
}
