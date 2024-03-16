package httpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"shortener/internal/app"
	"shortener/internal/errs"
	"shortener/internal/models"
	"strings"
)

// api contains utilities for correct work handler functions
type api struct {
	ctx       context.Context
	shortener app.Shortener
}

// routeHandlers represents a simple router
func (a api) routeHandlers(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		a.shortenUrl(w, r)
	case http.MethodGet:
		a.GetShortUrl(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		urlErrorResponce(w, errs.ErrNotAllowedMethod)
	}
}

// shortenUrl expects POST request and returns short url
//
//	@Summary		Shorten url
//	@Description	Shorten url provided in body and save it to storage
//	@Accept			json
//	@Produce		json
//	@Param			url	body	models.UserRequest	true	"Original url with protocol included"
//	@Router			/ [post]
//
//	@Success		200	{object}	UserResponse "Url shortened successfully"
//
//	@Failure		400	{object}	UserResponse "Json is invalid"
//	@Failure		422	{object}	UserResponse "Key 'url' is invalid or not provided"
//	@Failure		500	{object}	UserResponse "Short url creation caused error"
func (a api) shortenUrl(w http.ResponseWriter, r *http.Request) {
	var ur models.UserRequest
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := d.Decode(&ur)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		urlErrorResponce(w, err)
		log.Println(err)
		return
	}

	if isValidUrl(ur.Url) != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		urlErrorResponce(w, errs.ErrInvalidUrl)
		return
	}

	shortUrl, err := a.shortener.CreateUrl(a.ctx, ur.Url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		urlErrorResponce(w, err)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	urlSuccessResponse(w, shortUrl)
}

// GetShortUrl expects GET request and returns origin url by short form
// - 400, if short url not found in storage

// @Summary		Get original url
// @Description	Returns origin url by short form
// @Accept			json
// @Produce		json
// @Param			short_url	query	string	true	"Short url hash"
// @Router			/{hash} [get]
//
// @Success		200	{object}	UserResponse "Short url exists in storage"
//
// @Failure		400	{object}	UserResponse "Short url not found in storage"
func (a api) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := strings.TrimPrefix(r.URL.Path, "/")
	url, err := a.shortener.GetUrl(a.ctx, shortUrl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		urlErrorResponce(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	urlSuccessResponse(w, url)
}

func isValidUrl(u string) error {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return err
	}
	return nil
}

// urlSuccessResponse returns success response
func urlSuccessResponse(w http.ResponseWriter, url string) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(UserResponse{Url: url, Err: ""})
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Write(jsonResponse)
}

// urlErrorResponce returns error response
func urlErrorResponce(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(UserResponse{Url: "", Err: err.Error()})
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Write(jsonResponse)
}
