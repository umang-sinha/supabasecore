package supabasecore

import (
	"net/http"

	"github.com/umang-sinha/supabasecore/auth"
	"github.com/umang-sinha/supabasecore/internal/logger"
)

type Client struct {
	supabaseURL string
	supabaseKey string
	HTTPClient  *http.Client
	Auth        *auth.Service
}

func EnableLogging(level string) {
	logger.Init(level)
}
