// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ueauthentication

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/free5gc/logger_util"
	"github.com/free5gc/udm/logger"
)

var HttpLog *logrus.Entry

func init() {
	HttpLog = logger.HttpLog
}

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	AddService(router)
	return router
}

func genAuthDataHandlerFunc(c *gin.Context) {
	c.Params = append(c.Params, gin.Param{Key: "supiOrSuci", Value: c.Param("supi")})
	if strings.ToUpper("Post") == c.Request.Method {
		HttpGenerateAuthData(c)
		return
	}

	c.String(http.StatusNotFound, "404 page not found")
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nudm-ueau/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}

	genAuthDataPath := "/:supi/security-information/generate-auth-data"
	group.Any(genAuthDataPath, genAuthDataHandlerFunc)

	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"ConfirmAuth",
		strings.ToUpper("Post"),
		"/:supi/auth-events",
		HTTPConfirmAuth,
	},
}

var specialRoutes = Routes{
	{
		"GenerateAuthData",
		strings.ToUpper("Post"),
		"/:supiOrSuci/security-information/generate-auth-data",
		HttpGenerateAuthData,
	},
}
