package repositories

import (
	"context"
	"fmt"
	"github.com/dscamargo/go_app_template/internal/models"
	"github.com/uptrace/bun"
)

type ExampleRepository struct {
	db *bun.DB
}

func NewExampleRepository(db *bun.DB) *ExampleRepository {
	return &ExampleRepository{db}
}

func (r *ExampleRepository) GetAll(ctx context.Context) ([]models.Example, error) {
	items := make([]models.Example, 0)
	err := r.db.NewSelect().Model(&items).Scan(ctx)
	if err != nil {
		return []models.Example{}, fmt.Errorf("ExampleRepository.GetAll - %w", err)
	}
	return items, nil
}
