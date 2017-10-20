package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/buger/jsonparser"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	r := gofight.New()

	r.GET("/myendpoint").
		SetHeader(gofight.H{
			"Authorization": "authToken",
		}).
		Run(SetupServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			body := r.Body.Bytes()

			method, _ := jsonparser.GetString(body, "method")
			assert.Equal(t, "GET", method)

			path, _ := jsonparser.GetString(body, "path")
			assert.Equal(t, "/myendpoint", path)

			rbody, _ := jsonparser.GetString(body, "body")
			assert.Empty(t, rbody)

			auth, _ := jsonparser.GetString(body, "headers", "Authorization")
			assert.Equal(t, "authToken", auth)
		})
}
func TestPOST(t *testing.T) {
	r := gofight.New()

	r.POST("/anotherendpoint").
		SetHeader(gofight.H{
			"Authorization": "authToken",
		}).
		SetJSON(gofight.D{
			"a": 1,
			"b": "aString",
		}).
		Run(SetupServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			body := r.Body.Bytes()

			method, _ := jsonparser.GetString(body, "method")
			assert.Equal(t, "POST", method)

			path, _ := jsonparser.GetString(body, "path")
			assert.Equal(t, "/anotherendpoint", path)

			bodyA, _ := jsonparser.GetInt(body, "body", "a")
			assert.Equal(t, 1, int(bodyA))

			bodyB, _ := jsonparser.GetString(body, "body", "b")
			assert.Equal(t, "aString", bodyB)

			auth, _ := jsonparser.GetString(body, "headers", "Authorization")
			assert.Equal(t, "authToken", auth)
		})
}
func TestInvalid(t *testing.T) {
	r := gofight.New()

	r.POST("/anotherendpoint").
		SetBody("{a:aze}"). // invalid body
		Run(SetupServer(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusOK, r.Code)
			body := r.Body.Bytes()

			bodyA, _, _, _ := jsonparser.Get(body, "body")
			assert.Empty(t, bodyA)
		})
}
