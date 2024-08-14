package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	generate "screenshot/internal/api/handler/screenshot"
	"screenshot/internal/pkg/chrome"

	"github.com/chromedp/chromedp"
)

const (
	chromeHeadlessHost = "127.0.0.1"
	chromeHeadlessPort = 9222
)

func main() {
	portStr := flag.String("port", "3001", "chrome server port")
	flag.Parse()

	if portStr == nil || *portStr == "" {
		println("port is empty")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(*portStr, 10, 64)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	chromeNew := chrome.Must(
		chrome.WithHeadlessHost(chromeHeadlessHost),
		chrome.WithHeadlessPort(chromeHeadlessPort),
	)

	allocatorCtx, allocatorCancel := chromeNew.NewWorkerAllocator(ctx)
	defer allocatorCancel()

	tCtx, tCancel := chromedp.NewContext(allocatorCtx)

	if errRun := chromedp.Run(tCtx); errRun != nil {
		tCancel()
		os.Exit(1)
	}

	if _, err := chromeNew.WSDebuggerURL(); err != nil {
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// serve assets files
	mux.Handle("/assets/", http.StripPrefix("/assets", staticFileServer()))

	mux.Handle("/chrome/screenshot/", generate.New(log, chromeNew))

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func staticFileServer() http.Handler {
	fs := http.FileServer(http.Dir("./public/assets/"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientRemoteAddr := r.RemoteAddr
		clientHOST := r.URL.Host
		log.Printf("clientRemoteAddr: %s", clientRemoteAddr)
		log.Printf("clientHOST: %s", clientHOST)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		fs.ServeHTTP(w, r)
	})
}
