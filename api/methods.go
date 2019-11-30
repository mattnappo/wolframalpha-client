package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/loggo"
)

// getHistory returns the search history of the current session's user.
func (api *API) getHistory(ctx *gin.Context) {}

// search makes a search using the scraping engine.
func (api *API) search(ctx *gin.Context) {
	// searchText struct for request validation
	type searchText struct {
		SearchText string `form:"search_text" json:"search_text"`
	}

	// Load the request into the validation struct
	var newSearchText searchText
	err := ctx.ShouldBindJSON(&newSearchText)
	if err != nil {
		// Log and respond with the error
		api.Logger.Logf(loggo.ERROR, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new search object
	searchObject, err := api.Engine.NewSearchObject(newSearchText.SearchText)
	if err != nil {
		// Log and respond with the error
		api.Logger.Logf(loggo.ERROR, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Execute the search
	err = searchObject.Search(api.CWD)
	if err != nil {
		// Log and respond with the error
		api.Logger.Logf(loggo.ERROR, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	searchResult := searchObject.Result.String()                 // Get the search result
	ctx.JSON(http.StatusOK, gin.H{"searchResult": searchResult}) // Respond success
}
