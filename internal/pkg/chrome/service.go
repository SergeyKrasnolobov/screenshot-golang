package chrome

import (
	"context"

	"github.com/pkg/errors"
)

type service struct {
	headlessHost         string
	headlessPort         int64
	webSocketDebuggerUrl string
	chrContext           context.Context
}

// Must create new service object with options from arguments or throw panic
func Must(options ...Option) *service {
	s, err := New(options...)
	if err != nil {
		panic(err)
	}
	return s
}

// New return new service object with options from arguments
func New(options ...Option) (*service, error) {
	s := &service{}

	for i, configure := range options {
		if err := configure(s); err != nil {
			return nil, errors.Wrapf(err, "chrome: invalid option %d", i)
		}
	}

	if s.headlessHost == "" {
		return nil, errors.New("chrome: please provide headlessHost")
	}
	if s.headlessPort == 0 {
		return nil, errors.New("chrome: please provide headlessPort")
	}

	return s, nil
}
