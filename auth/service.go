package auth

import "net/http"

type Service struct {
	supabaseURL string
	supabaseKey string
	httpClient  *http.Client
}

func NewService(supabaseURL, supabaseKey string, httpClient *http.Client) *Service {
	return &Service{
		supabaseURL: supabaseURL + "/auth/v1",
		supabaseKey: supabaseKey,
		httpClient:  httpClient,
	}
}
