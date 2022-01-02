package api

import (
	"fmt"
	"github.com/labstack/echo"
	obj "goprojects/urlshortener/config/db"
	"goprojects/urlshortener/model"
	"net/http"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base int64 = 62

func APIShortURL(echo echo.Context) error {

	var urlToShort model.URL
	if err := echo.Bind(&urlToShort); err != nil {
		return err
	}

	if urlToShort.URL == "" {
		return echo.JSON(http.StatusBadRequest, model.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "please provide a valid url",
		})
	}

	row, err := obj.Db.Query("call shorturl.insertURLAndGetID(?)", urlToShort.URL)
	if err != nil {
		return err
	}

	var id int64
	if row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return err
		}

		shortCode := toBinary(id)
		shortURL := getShortURL(shortCode)
		_, err := obj.Db.Query("call shorturl.insertShortURL(?, ?)", shortURL, id)
		if err != nil {
			return err
		}

		var mShortURL model.SuccessResponse
		mShortURL.Data = fmt.Sprintf("%s%s", "http://localhost:8000/", shortURL)
		return echo.JSON(http.StatusOK, model.SuccessResponse{
			Status: http.StatusOK,
			Data:   mShortURL,
		})

	}

	return echo.JSON(http.StatusInternalServerError, model.ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
	})

}

func APIRedirectToURL(echo echo.Context) error {

	shortURL := echo.Param("path")

	row, err := obj.Db.Query("call shorturl.getOriginalURL(?)", shortURL)
	if err != nil {
		return err
	}

	if row.Next() {

		urlToRedirect := ""
		err = row.Scan(&urlToRedirect)
		if err != nil {
			return err
		}

		echo.Redirect(http.StatusPermanentRedirect, urlToRedirect)

	}

	return echo.JSON(http.StatusInternalServerError, model.ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "no url found",
	})

}

//function for decimal to binary conversion
func toBinary(id int64) []int64 {

	var baseEquivalent []int64

	if id < base {
		baseEquivalent = append(baseEquivalent, id)
		return baseEquivalent
	}

	for id >= base {

		//if id < base {
		//	baseEquivalent = append(baseEquivalent, id)
		//	return baseEquivalent
		//}

		baseEquivalent = append(baseEquivalent, id%base)
		id /= base

	}

	baseEquivalent = append(baseEquivalent, id%base)

	i, j := 0, len(baseEquivalent)-1

	for {

		temp := baseEquivalent[i]
		baseEquivalent[i] = baseEquivalent[j]
		baseEquivalent[j] = temp
		i++
		j--

		if i >= j {
			break
		}

	}

	return baseEquivalent

}

func getShortURL(shortCode []int64) string {

	if len(shortCode) == 1 {
		return string(chars[shortCode[0]])
	}

	var url []byte
	for _, j := range shortCode {
		url = append(url, chars[j])
	}

	return string(url)

}
