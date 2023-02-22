package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// ミドルウェアの登録
	mux.Use(middleware.Recoverer)

	// ルートの登録
	mux.Get("/", app.Home)

	// 静的資産
	return mux
}
