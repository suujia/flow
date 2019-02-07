// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"sync/atomic"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/suujia/flow/api/gen/models"
	"github.com/suujia/flow/api/gen/restapi/operations"
	"github.com/suujia/flow/api/gen/restapi/operations/spots"
)

// the variables we need throughout our implementation
var testspots = make(map[int64]*models.Spots)
var lastID int64

func addSpot(spot *models.Spots) error {
	if spot == nil {
		return errors.New(500, "spot must be present")
	}
	newID := atomic.AddInt64(&lastID, 1)
	spot.ID = newID
	testspots[newID] = spot
	return nil
}

func deleteSpot(spotID int64) error {
	_, exist := testspots[spotID]
	if !exist {
		return errors.NotFound("not found: spot %d", spotID)
	}
	delete(testspots, spotID)
	return nil
}

func getSpots(since int64, limit int32) (results []*models.Spots) {
	results = make([]*models.Spots, len(testspots))
	for id, spot := range testspots {
		if len(results) >= int(limit) {
			return
		}
		if since == 0 || id > since {
			results = append(results, spot)
		}
	}
	return
}

func getSpot(spotID int64) (result *models.Spots) {
	item, exist := testspots[spotID]
	if exist {
		result = item
	}
	return
}

func updateSpot(id int64, spot *models.Spots) error {
	if spot == nil {
		return errors.New(500, "spot must be present")
	}

	_, exist := testspots[id]
	if !exist {
		return errors.NotFound("not found: item %d", id)
	}
	spot.ID = id
	testspots[id] = spot
	return nil
}

//go:generate swagger generate server --target ../../gen --name Spotter --spec ../../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.SpotterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SpotterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.SpotsAddSpotHandler = spots.AddSpotHandlerFunc(func(params spots.AddSpotParams) middleware.Responder {
		if err := addSpot(params.Body); err != nil {
			return spots.NewAddSpotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return spots.NewAddSpotCreated().WithPayload(params.Body)
	})

	api.SpotsDeleteSpotHandler = spots.DeleteSpotHandlerFunc(func(params spots.DeleteSpotParams) middleware.Responder {
		if err := deleteSpot(params.ID); err != nil {
			return spots.NewDeleteSpotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return spots.NewDeleteSpotNoContent()
	})

	api.SpotsUpdateSpotHandler = spots.UpdateSpotHandlerFunc(func(params spots.UpdateSpotParams) middleware.Responder {
		if err := updateSpot(params.ID, params.Body); err != nil {
			return spots.NewUpdateSpotDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return spots.NewUpdateSpotOK().WithPayload(params.Body)
	})

	api.SpotsGetHandler = spots.GetHandlerFunc(func(params spots.GetParams) middleware.Responder {
		mergedParams := spots.NewGetParams()
		mergedParams.Since = swag.Int64(0)
		if params.Since != nil {
			mergedParams.Since = params.Since
		}
		if params.Limit != nil {
			mergedParams.Limit = params.Limit
		}
		return spots.NewGetOK().WithPayload(getSpots(*mergedParams.Since, *mergedParams.Limit))
	})

	api.SpotsFindSpotHandler = spots.FindSpotHandlerFunc(func(params spots.FindSpotParams) middleware.Responder {
		mergedParams := spots.NewFindSpotParams()
		if params.ID >= 0 {
			mergedParams.ID = params.ID
		}
		return spots.NewFindSpotOK().WithPayload(getSpot(mergedParams.ID))
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
