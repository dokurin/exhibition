//go:generate oapi-codegen -generate "types,chi-server" -o oapi/oapi_gen.go -package oapi ../../api/openapi.yaml

package main

import (
	"context"
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
	var srv http.Server

	// Graceful shutdown
	waitShutdown := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		log.Println("shutdown server...")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown server: %v", err)
		}
		waitShutdown <- struct{}{}
	}()

	// routing settings
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Route("/api", func(r chi.Router) {
		oapi.HandlerFromMux(new(handler), r)
	})

	srv.Handler = router
	srv.Addr = ":8080"
	log.Println("start server in http://localhost:8080")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("failed to listen serve: %v", err)
	}
	<-waitShutdown
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
