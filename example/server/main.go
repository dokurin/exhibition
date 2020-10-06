//go:generate oapi-codegen -generate "types,chi-server" -o oapi/oapi_gen.go -package oapi ../../api/openapi.yaml

package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/dokurin/exhibition/example/server/oapi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	// routing settings
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Route("/api", func(r chi.Router) {
		oapi.HandlerFromMux(new(handler), r)
	})

	if err := listhenAndServe(8080, router); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// listhenAndServe starts http server with graceful shutdown
func listhenAndServe(port int, h http.Handler) error {
	log.Printf("start server on http://localhost:%d", port)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: h,
	}

	errChan := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		log.Println("shutdown...")
		errChan <- srv.Shutdown(context.Background())
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return <-errChan
}

type handler struct{}

// 作家一覧
// (GET /artists)
func (h *handler) ListArtist(w http.ResponseWriter, r *http.Request, params oapi.ListArtistParams) {
	render.JSON(w, r, oapi.ListArtist{
		Artists:       &artists,
		NextPageToken: strPtr("xxxxxxxxxx"),
		TotalCount:    int32Ptr(1),
	})
}

// 作家詳細情報取得
// (GET /artists/{artist_id})
func (h *handler) GetArtist(w http.ResponseWriter, r *http.Request, artistId string) {
	render.JSON(w, r, artists[0])
}

// 作品一覧取得
// (GET /artists/{artist_id}/artofworks)
func (h *handler) ListArtOfWork(w http.ResponseWriter, r *http.Request, artistId string) {
	render.JSON(w, r, oapi.ListArtOfWork{
		ArtOfWorks:    &artOfWorks,
		NextPageToken: strPtr("xxxxxxxxxx"),
		TotalCount:    int32Ptr(1),
	})
}

func linkPtr(l []oapi.ArtistLink) *[]oapi.ArtistLink { return &l }
func strPtr(s string) *string                        { return &s }
func int32Ptr(i int32) *int32                        { return &i }

// dummy datas
var (
	artists = []oapi.Artist{
		{
			Id: "artist1",
			Links: linkPtr([]oapi.ArtistLink{
				{
					DisplayName: "github",
					Icon:        strPtr("https://aaaa.png"),
					Url:         "https://hogehoge.com",
				},
				{
					DisplayName: "twitter",
					Icon:        strPtr("https://aaaa.png"),
					Url:         "https://hogehoge.com",
				},
			}),
			Masterpiece: oapi.ArtistMasterpiece{
				Title: "テスト",
				Url:   "https://xxxxxxxx.png",
			},
			Profile: oapi.ArtistProfile{
				Avator:       strPtr("https://hoge.png"),
				Introduction: "自己紹介でぃす",
				Name:         "ryutah",
			},
		},
	}

	artOfWorks = []oapi.ArtOfWork{
		{
			Id: strPtr("art1"),
			DisplayItems: oapi.ArtOfWOrkDisplayItem{
				MimeType: strPtr("image/png"),
				Url:      strPtr("https://xxxx"),
			},
			Information: oapi.ArtOfWorkInformation{
				Introduction: "作品の説明",
				Title:        "素晴らしい作品",
			},
		},
	}
)
