package services

import (
	"context"
	"github.com/dscamargo/go_app_template/internal/models"
	"github.com/dscamargo/go_app_template/internal/repositories"
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
