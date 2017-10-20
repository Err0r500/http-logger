package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupServer()
	r.Run(":8080")
}

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.Use(
		cors.New(
			cors.Config{
				AllowAllOrigins:  true,
				AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
				AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
				ExposeHeaders:    []string{},
				MaxAge:           50 * time.Second,
				AllowCredentials: true,
			},
		),
	)

	router.Use(RequestLogger())

	return router
}

type resp struct {
	Method  string            `json:"method,omitempty"`
	Path    string            `json:"path,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    *interface{}      `json:"body,omitempty"`
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(
			http.StatusOK,
			resp{
				Method:  c.Request.Method,
				Path:    c.Request.URL.Path,
				Headers: headerToMap(c.Request.Header),
				Body:    getBody(c.Request.Body),
			},
		)
	}
}

func getBody(rC io.ReadCloser) *interface{} {
	var body interface{}
	bodyBytes, _ := ioutil.ReadAll(rC)

	if len(bodyBytes) > 0 {
		d := json.NewDecoder(bytes.NewReader(bodyBytes))
		if err := d.Decode(&body); err != nil {
			return nil
		}
		return &body
	}
	return nil
}

func headerToMap(header http.Header) map[string]string {
	res := map[string]string{}

	for name, values := range header {
		for _, value := range values {
			res[name] = value
		}
	}
	return res
}
