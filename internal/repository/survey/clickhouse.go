package survey

import (
	"context"
	"fmt"
	"innovatex-app/internal/models"
)

type viaClickhouse struct {
}

func newViaClickhouse() *viaClickhouse {
	return &viaClickhouse{}
}

func (r *viaClickhouse) Save(ctx context.Context, survey *models.Survey) error {
	fmt.Printf("%+v", survey)
	return nil
}
