package screenshoter

import (
	"github.com/pkg/errors"
)

type service struct {
	chromeServerHost string
	chromeServerPort int64
}

// Must create new pdf_generator service object with options from arguments or throw panic
func Must(options ...Option) *service {
	s, err := New(options...)
	if err != nil {
		panic(err)
	}
	return s
}

// New return new pdf_generator service object with options from arguments
func New(options ...Option) (*service, error) {
	s := &service{}

	for i, configure := range options {
		if err := configure(s); err != nil {
			return nil, errors.Wrapf(err, "screenshot_handler: invalid option %d", i)
		}
	}
	if s.chromeServerHost == "" {
		return nil, errors.New("screenshot_handler: please provide chromeServerHost")
	}
	if s.chromeServerPort == 0 {
		return nil, errors.New("screenshot_handler: please provide chromeServerPort")
	}

	return s, nil
}
