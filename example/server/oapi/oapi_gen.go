// Package oapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package oapi

import (
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi"
	"net/http"
)

// ArtOfWOrkDisplayItem defines model for ArtOfWOrkDisplayItem.
type ArtOfWOrkDisplayItem struct {

	// 作品の形式
	MimeType *string `json:"mime_type,omitempty"`

	// 作品のURL
	Url *string `json:"url,omitempty"`
}

// ArtOfWork defines model for ArtOfWork.
type ArtOfWork struct {

	// 作品
	DisplayItems ArtOfWOrkDisplayItem `json:"display_items"`

	// 作品ID
	Id *string `json:"id,omitempty"`

	// 作品詳細情報
	Information ArtOfWorkInformation `json:"information"`
}

// ArtOfWorkInformation defines model for ArtOfWorkInformation.
type ArtOfWorkInformation struct {

	// 作品の紹介文
	Introduction string `json:"introduction"`

	// 作品タイトル
	Title string `json:"title"`
}

// Artist defines model for Artist.
type Artist struct {

	// 作家ID
	Id    string        `json:"id"`
	Links *[]ArtistLink `json:"links,omitempty"`

	// 代表作
	Masterpiece ArtistMasterpiece `json:"masterpiece"`

	// 作家情報
	Profile ArtistProfile `json:"profile"`
}

// ArtistLink defines model for ArtistLink.
type ArtistLink struct {

	// 表示名
	DisplayName string `json:"display_name"`

	// サービスアイコンのURL
	Icon *string `json:"icon,omitempty"`

	// URL
	Url string `json:"url"`
}

// ArtistMasterpiece defines model for ArtistMasterpiece.
type ArtistMasterpiece struct {

	// 作品タイトル
	Title string `json:"title"`

	// 作品の画像URL
	// 解像度: 150x150(適当)
	Url string `json:"url"`
}

// ArtistProfile defines model for ArtistProfile.
type ArtistProfile struct {

	// アバターのURL
	Avator *string `json:"avator,omitempty"`

	// 作家紹介
	Introduction string `json:"introduction"`

	// 作家名
	Name string `json:"name"`
}

// Error defines model for Error.
type Error struct {

	// HTTPステータスコード
	Code int32 `json:"code"`

	// エラーリスト
	Errors []ErrorDetail `json:"errors"`

	// エラーメッセージ概要 画面上にエラーを表示することが想定されている
	Message string `json:"message"`
}

// ErrorDetail defines model for ErrorDetail.
type ErrorDetail struct {

	// エラーの詳細内容 画面上に表示される形式で有ることは保証されない
	Message string `json:"message"`

	// エラー種別
	Reason string `json:"reason"`
}

// GetArtist defines model for GetArtist.
type GetArtist Artist

// ListArtOfWork defines model for ListArtOfWork.
type ListArtOfWork struct {
	ArtOfWorks *[]ArtOfWork `json:"art_of_works,omitempty"`

	// 次ページ取得のためのトークン
	NextPageToken *string `json:"next_page_token,omitempty"`

	// 総作家数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// ListArtist defines model for ListArtist.
type ListArtist struct {
	Artists *[]Artist `json:"artists,omitempty"`

	// 次ページ取得のためのトークン
	NextPageToken *string `json:"next_page_token,omitempty"`

	// 総作家数
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NotFoundError defines model for NotFoundError.
type NotFoundError Error

// ServerError defines model for ServerError.
type ServerError Error

// ListArtistParams defines parameters for ListArtist.
type ListArtistParams struct {

	// ページトークンを指定する
	//
	// 一覧を取得する際に、レスポンスの `next_page_token` の値を設定してリクエストを送信することで
	// 次のページの一覧を取得することができる
	PageToken *string `json:"page_token,omitempty"`

	// 一度のリクエストで取得するリストの数を指定する
	PageSize *int32 `json:"page_size,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 作家一覧
	// (GET /artists)
	ListArtist(w http.ResponseWriter, r *http.Request, params ListArtistParams)
	// 作家詳細情報取得
	// (GET /artists/{artist_id})
	GetArtist(w http.ResponseWriter, r *http.Request, artistId string)
	// 作品一覧取得
	// (GET /artists/{artist_id}/artofworks)
	ListArtOfWork(w http.ResponseWriter, r *http.Request, artistId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ListArtist operation middleware
func (siw *ServerInterfaceWrapper) ListArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params ListArtistParams

	// ------------- Optional query parameter "page_token" -------------
	if paramValue := r.URL.Query().Get("page_token"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_token", r.URL.Query(), &params.PageToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page_token: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "page_size" -------------
	if paramValue := r.URL.Query().Get("page_size"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "page_size", r.URL.Query(), &params.PageSize)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page_size: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ListArtist(w, r.WithContext(ctx), params)
}

// GetArtist operation middleware
func (siw *ServerInterfaceWrapper) GetArtist(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "artist_id" -------------
	var artistId string

	err = runtime.BindStyledParameter("simple", false, "artist_id", chi.URLParam(r, "artist_id"), &artistId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter artist_id: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.GetArtist(w, r.WithContext(ctx), artistId)
}

// ListArtOfWork operation middleware
func (siw *ServerInterfaceWrapper) ListArtOfWork(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "artist_id" -------------
	var artistId string

	err = runtime.BindStyledParameter("simple", false, "artist_id", chi.URLParam(r, "artist_id"), &artistId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter artist_id: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ListArtOfWork(w, r.WithContext(ctx), artistId)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Get("/artists", wrapper.ListArtist)
	})
	r.Group(func(r chi.Router) {
		r.Get("/artists/{artist_id}", wrapper.GetArtist)
	})
	r.Group(func(r chi.Router) {
		r.Get("/artists/{artist_id}/artofworks", wrapper.ListArtOfWork)
	})

	return r
}
