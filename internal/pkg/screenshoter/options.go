package screenshoter

// Option type describe option for client object
type Option func(*service) error

func WithChromeServerHost(host string) Option {
	return func(service *service) error {
		service.chromeServerHost = host
		return nil
	}
}

func WithChromeServerPort(port int64) Option {
	return func(service *service) error {
		service.chromeServerPort = port
		return nil
	}
}
