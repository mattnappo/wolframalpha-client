package common

const (
	// APIServerPort is the port that the API server will run on.
	APIServerPort = 8080

	// SeleniumPort is the port that the selenium server will run on.
	SeleniumPort = 8081
)

// AllowOrigins is the list of allowed CORS origins.
var AllowOrigins = []string{
	"http://localhost:8002",
}
