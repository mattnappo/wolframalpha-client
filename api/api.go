package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juju/loggo"
	"github.com/juju/loggo/loggocolor"
	"github.com/xoreo/wolframalpha-client/common"
	"github.com/xoreo/wolframalpha-client/core"
	"github.com/xoreo/wolframalpha-client/engine"
)

// API represents an API instance.
type API struct {
	Router *gin.Engine `json:"router"` // The HTTP router

	DB *bolt.DB `json:"db"` // The database

	CWD    *core.ChromeWebDriver `json:"cwd"` // The web driver / scraper
	Engine engine.Engine         // The scraping engine

	Root string `json:"root"` // API root route

	Logger *loggo.Logger `json:"logger"` // The logger

	Client *http.Client `json:"client"` // The HTTP client
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

	// Create the Chrome WebDriver
	cwd, err := core.NewChromeWebDriver(common.SeleniumPort)
	if err != nil {
		return nil, err
	}

	// Create new API instance
	api := &API{
		Router: router,             // Set router
		DB:     nil,                // Nil for now
		CWD:    cwd,                // The Chrome WebDriver (the scraper)
		Root:   "/api/",            // Set root route
		Client: http.DefaultClient, // Set the HTTP client to the default HTTP client
	}

	// Initialize the API's logger
	err = api.initLogger()
	if err != nil {
		return nil, err
	}

	return api, nil
}

// StartServing starts serving the API.
func (api *API) StartServing(port int64) error {
	err := api.setupRoutes() // Setup API routes
	if err != nil {          // Check for errors
		return nil // No error occurred, return nil
	}

	// Start the router
	err = api.Router.Run(
		fmt.Sprintf("0.0.0.0:%s", strconv.FormatInt(port, 10)),
	)
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

// reportError will log and return an error.
func reportError(err error, api *API, ctx *gin.Context) {}
