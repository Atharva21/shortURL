package main

import (
	"net/http"
	"net/url"

	"github.com/Atharva21/shortURL/handler"
	"github.com/Atharva21/shortURL/handler/encoder"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	URL string `json:"url" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.POST("/encode", func(c *gin.Context) {
		var requestBody RequestBody
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "please enter 'url' in request body",
			})
			return
		}
		_, err = url.ParseRequestURI(requestBody.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "please enter a valid URL",
			})
			return
		}

		present, err := handler.CheckIfLongURLPresent(requestBody.URL)
		if present && err == nil {
			shortURL, err := handler.GetShortURL(requestBody.URL)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"encodedURL": shortURL,
				})
				return
			}
		}

		shortURL, err := encoder.URLEncoder.Encode(requestBody.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
			return
		}

		err = handler.Link(shortURL, requestBody.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"encodedURL": shortURL,
		})
	})

	router.GET("/:shortURL", func(c *gin.Context) {
		shortURL, _ := c.Params.Get("shortURL")

		longURL, err := handler.GetLongURL(shortURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Not found",
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, longURL)

	})

	router.Run()
}
