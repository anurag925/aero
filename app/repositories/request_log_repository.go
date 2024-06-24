package repositories

import (
	"context"

	"github.com/anurag925/aero/app/models"
)

type RequestLog interface {
	Log(ctx context.Context, log models.RequestLog) error
}
