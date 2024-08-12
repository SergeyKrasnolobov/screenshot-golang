package screenshoter

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	loggerWrapper "screenshot/internal/logger"
	types "screenshot/internal/pkg"

	resty "github.com/go-resty/resty/v2"
)

var ErrGen = errors.New("error while generating screenshot")

func (s *service) GenerateScreenshot(ctx context.Context, loggerr *slog.Logger, body types.InputBody, restyClient *resty.Client) ([]byte, error) {
	resp, err := restyClient.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(s.generateScreenshot())
	if err != nil {
		loggerWrapper.New(ctx, loggerr).Error(err)
		return nil, err
	}

	if !resp.IsSuccess() {
		loggerWrapper.New(ctx, loggerr).Error(resp.String())
		return nil, err
	}

	return resp.Body(), nil
}

func (s *service) generateScreenshot() string {
	return fmt.Sprintf("http://%s:%d/chrome/screenshot/", s.chromeServerHost, s.chromeServerPort)
}
