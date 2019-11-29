package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"github.com/xoreo/wolframalpha-client/common"
	//"github.com/boltdb/bolt"
)

// API represents an API instance.
type API struct {
	Router *gin.Engine `json:"router"` // The HTTP router

	Root string // API root route

	Logger *loggo.Logger // The logger

	Client *http.Client // The HTTP client
}

// NewAPI initializes a new API instance.
func NewAPI() (*API, error) {

	// Initialize the API router
	router := gin.Default()

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: common.AllowOrigins,
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Access-Control-Allow-Origin",
			"Content-Type",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == common.AllowOrigins[0]
		},
		MaxAge: 12 * time.Hour,
	}))

	// Create new API instance
	api := &API{
		Router: router,             // Set router
		Root:   "/api/",            // Set root route
		Client: http.DefaultClient, // Set the HTTP client to the default HTTP client
	}

	// Initialize the API's logger
	err := api.initLogger()
	if err != nil {
		return nil, err
	}

	return api, nil
}

// StartServing starts serving the API.
func (api *API) StartServing(port int64) error {
	err := api.SetupRoutes() // Setup API routes
	if err != nil {          // Check for errors
		return nil // No error occurred, return nil
	}

	// Start the router
	err = api.Router.Run("0.0.0.0:" + strconv.FormatInt(port, 10))
	if err != nil {
		return err
	}

	return nil // No error occurred, return nil
}

// initLogger initializes the API's logger.
func (api *API) initLogger() error {
	logger := loggo.GetLogger("API") // Get logger
	// Create log dir
	err := common.CreateDirIfDoesNotExist(
		filepath.FromSlash(common.LogsDir),
	)
	if err != nil { // Check for errors
		return err // Return found error
	}

	// Create the log file
	logFile, err := os.OpenFile(filepath.FromSlash(fmt.Sprintf(
		"%s/logs_%s.txt", common.LogsDir,
		time.Now().Format("2006-01-02_15-04-05"))),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil { // Check for errors
		return err // Return found error
	}

	// Enabled colored output
	_, err = loggo.ReplaceDefaultWriter(loggocolor.NewWriter(os.Stderr))
	if err != nil {
		return err // Return found error
	}

	// Register file writer
	err = loggo.RegisterWriter("logs", loggo.NewSimpleWriter(
		logFile, loggo.DefaultFormatter),
	)
	if err != nil { // Check for errors
		return err // Return found error
	}

	api.Logger = &logger // Get a pointer to the logger

	return nil // Everything is fine, how are you?
}
