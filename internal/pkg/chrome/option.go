package chrome

// Option type describe option for client object
type Option func(*service) error

func WithHeadlessHost(host string) Option {
	return func(service *service) error {
		service.headlessHost = host
		return nil
	}
}

func WithHeadlessPort(port int64) Option {
	return func(service *service) error {
		service.headlessPort = port
		return nil
	}
}
