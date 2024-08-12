package screenshoter

import (
	"context"

	"screenshot/internal/logger"
	types "screenshot/internal/pkg"

	"github.com/go-resty/resty/v2"
)

type Service interface {
	GenerateScreenshot(ctx context.Context, logger *logger.Logger, body types.InputBody, restyClient *resty.Client) ([]byte, error)
	GetDebugInfo(ctx context.Context, logger *logger.Logger, restyClient *resty.Client) ([]byte, error) // not implemented for now
}
