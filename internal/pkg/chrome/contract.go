package chrome

import "context"

type Service interface {
	NewWorkerAllocator(ctx context.Context) (context.Context, context.CancelFunc)
	WSDebuggerURL() (string, error)
	GetWsURL() (string, error)
	GetDebugInfo() ([]byte, error)
	SetChromeContext(ctx context.Context) error
	GetChromeContext() (context.Context, error)
}
