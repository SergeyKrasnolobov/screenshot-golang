package chrome

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/chromedp/chromedp"
)

func (s *service) NewWorkerAllocator(ctx context.Context) (context.Context, context.CancelFunc) {
	options := chromedp.DefaultExecAllocatorOptions[:]
	options = append(options,
		chromedp.Flag("headless", false), // to check if static resources are loaded
		chromedp.Flag("no-sandbox", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("remote-debugging-address", s.headlessHost),
		chromedp.Flag("remote-debugging-port", fmt.Sprintf("%d", s.headlessPort)),
	)

	allocatorCtx, allocatorCancel := chromedp.NewExecAllocator(ctx, options...)

	return allocatorCtx, allocatorCancel
}

func (s *service) GetChromeContext() (context.Context, error) {
	if s.chrContext != nil {
		return s.chrContext, nil
	}
	return nil, fmt.Errorf("could not find context")
}

func (s *service) SetChromeContext(ctx context.Context) error {
	if ctx == nil {
		return fmt.Errorf("could not find context: %v", ctx)
	}
	s.chrContext = ctx
	return nil
}

func (s *service) GetWsURL() (string, error) {
	if s.webSocketDebuggerUrl != "" {
		return s.webSocketDebuggerUrl, nil
	}
	return "", fmt.Errorf("could not find context")
}

// Вытаскиваем адресс webSocketDebuggerUrl
// Алгоритм описан тут https://stackoverflow.com/questions/56846538/how-to-obtain-the-wschromeendpointurl-on-windows
func (s *service) WSDebuggerURL() (string, error) {
	if s.webSocketDebuggerUrl != "" {
		return s.webSocketDebuggerUrl, nil
	}

	urlAddr := fmt.Sprintf("http://%s:%d/json/version", s.headlessHost, s.headlessPort)

	resp, err := http.Get(forceIP(urlAddr))
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var result map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if v, vOk := result["webSocketDebuggerUrl"]; vOk {
		if wsURL, ok := v.(string); ok {
			s.webSocketDebuggerUrl = wsURL
			return wsURL, nil
		}
	}

	return "", fmt.Errorf("could not find webSocketDebuggerUrl value")
}

func (s *service) GetDebugInfo() ([]byte, error) {
	urlAddr := fmt.Sprintf("http://%s:%d/json/version", s.headlessHost, s.headlessPort)

	resp, err := http.Get(forceIP(urlAddr))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func forceIP(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}

	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return urlStr
	}

	addr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return urlStr
	}

	u.Host = net.JoinHostPort(addr.IP.String(), port)

	return u.String()
}
