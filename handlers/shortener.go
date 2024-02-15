package handlers

import (
	"log"
	"net/http"

	"github.com/Jaraxzus/url-shortener/services"
	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	URL string `form:"url"`
}
type RedirectParams struct {
	Code string `uri:"code"`
}

func ShortenHandler(c *gin.Context) {
	var queryParams QueryParams
	if c.ShouldBind(&queryParams) == nil {
		log.Println(queryParams.URL)
	}

	url := queryParams.URL
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL cannot be empty"})
		return
	}

	urlService, ok := c.Get("urlService")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get URLService from context"})
		return
	}
	urlServiceTyped, ok := urlService.(services.URLService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to convert URLService to the correct type"})
		return
	}

	code, err := urlServiceTyped.ShortenURL(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func RedirectHandler(c *gin.Context) {
	var redirectParams RedirectParams
	if err := c.ShouldBindUri(&redirectParams); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	code := redirectParams.Code

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code cannot be empty"})
		return
	}

	redisService, ok := c.Get("redisService")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Println("Unable to get redisService from context")
		return
	}
	redisServiceTyped, ok := redisService.(services.RedisService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		log.Println("Unable to convert redisService to the correct type")
		return
	}
	url, err := redisServiceTyped.GetValue(code)
	if (err != nil) || (url == "") {
		log.Println("from urlService")
		urlService, ok := c.Get("urlService")
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			log.Println("Unable to get URLService from context")
			return
		}
		urlServiceTyped, ok := urlService.(services.URLService)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			log.Println("Unable to convert URLService to the correct type")
			return
		}
		url, err = urlServiceTyped.GetURL(code)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		err = redisServiceTyped.SaveValue(code, url)
		if err != nil {
			log.Println("error cached value")
		}
	}
	c.Redirect(http.StatusFound, url)
}
