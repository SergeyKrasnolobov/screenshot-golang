package generate

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	types "screenshot/internal/pkg"
	"screenshot/internal/pkg/chrome"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type VParams struct {
	Width  *int64
	Height *int64
}

type Handler struct {
	logger *slog.Logger
	chrome chrome.Service
}

func New(logger *slog.Logger, chrome chrome.Service) *Handler {
	return &Handler{
		logger: logger,
		chrome: chrome,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() {
		_ = r.Body.Close()
	}()

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	params := types.InputBody{}
	h.logger.Debug("params", slog.String("params", fmt.Sprintf("%+v", params)))

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		h.logger.Error("faild to decode body", err.Error(), "")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	devtoolsWsURL, err := h.chrome.GetWsURL()
	if err != nil || devtoolsWsURL == "" {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte(fmt.Errorf("failed to get devtoolsWsURL").Error())); err != nil {
			h.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	allocatorContext, allocatorCancel := chromedp.NewRemoteAllocator(ctx, devtoolsWsURL)
	defer allocatorCancel()

	chCtx, chCancel := chromedp.NewContext(allocatorContext)
	defer chCancel()

	buf, err := makeScreenshot(chCtx, params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")

	if _, err := w.Write(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func makeScreenshot(ctx context.Context, params types.InputBody) ([]byte, error) {
	var buf []byte

	if params.Source != "" {
		if err := chromedp.Run(ctx, rawHtmlScreenshot(params.Source, &params.Viewport.Height, &params.Viewport.Width, &buf)); err != nil {
			_, err := fmt.Printf("failed to make rawHtmlScreenshot: %v", err)
			return nil, err
		}
	}

	return buf, nil
}

func rawHtmlScreenshot(html string, height *int64, width *int64, res *[]byte) chromedp.Tasks {
	var viewPort VParams
	var defaultHeight int64 = int64(768)
	var defaultWidth int64 = int64(1024)

	// test html for quick testing
	htmltest := `
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Title</title>
            <link rel="preload" href="http://127.0.0.1:3001/assets/fonts/manrope/manrope-bold.woff2" as="font" type="font/woff2" crossorigin>
			<style>
				*, *::after, *::before {
					animation-delay: -0.0001s !important;
					animation-duration: 0s !important;
					animation-play-state: paused !important;
					transition-delay: 0s !important;
					transition-duration: 0s !important;
					caret-color: transparent !important;
				}

				@font-face {
                        font-family: Manrope;
                        src: url('http://127.0.0.1:3001/assets/fonts/manrope/manrope-bold.woff2') format('woff2');
                        font-weight: 700 800;
                }


				.container {
					width: 100%;
				}

				.button {
					width: 200px;
					height: 60px;
					background: rebeccapurple;
					color: white;
					border-radius: 10px;
				}

				.test {
					font-family: "Manrope";
					font-weight: 800;
					font-size: 20px;
				}
			</style>
		</head>
		<body>
			<div class="contaner">
				<button class="button">
				<span class="test">Hellow world</span>
				</button>
			</div>
		</body>
	</html>
`

	extraHeaders := map[string]interface{}{
		"Host": "http://127.0.0.1:3001",
	}

	if viewPort.Height = &defaultHeight; height != nil {
		viewPort.Height = height
	}

	if viewPort.Width = &defaultWidth; width != nil {
		viewPort.Width = width
	}

	scale := chromedp.EmulateScale(1)

	return chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(extraHeaders)),
		chromedp.EmulateViewport(*viewPort.Width, *viewPort.Height, scale),
		chromedp.Navigate("about:blank"),
		chromedp.PollFunction(`(html) => {
			document.open();
			document.write(html);
			// document.close(); // just for testing
			// return true // just for testing
		}`, nil, chromedp.WithPollingArgs(htmltest)),
		chromedp.CaptureScreenshot(res),
	}
}
