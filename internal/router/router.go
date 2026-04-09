package router

import (
	"io/fs"
	"log/slog"
	"net/http"

	"github.com/Rephobia/green-api-test-task/internal/handler"
	"github.com/Rephobia/green-api-test-task/internal/middleware"
)

func New(logger *slog.Logger, frontFiles fs.FS) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.FS(frontFiles)))

	mux.HandleFunc("/api/settings",
		middleware.Validate(handler.GetSettings),
	)
	mux.HandleFunc("/api/state",
		middleware.Validate(handler.GetStateInstance),
	)
	mux.HandleFunc("/api/send-message",
		middleware.Validate(handler.SendMessage),
	)
	mux.HandleFunc("/api/send-file",
		middleware.Validate(handler.SendFileByUrl),
	)

	return middleware.Chain(
		mux,
		middleware.Logging(logger),
		middleware.Recovery(logger),
	)
}
