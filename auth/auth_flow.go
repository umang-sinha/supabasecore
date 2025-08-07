package auth

import (
	"encoding/json"

	"github.com/umang-sinha/supabasecore/internal/logger"
)

func (s *Service) Signup(email, password string) (*SignUpResponse, error) {
	reqBody := SignUpRequest{Email: email, Password: password}

	bodyBytes, err := json.Marshal(reqBody)

	if err != nil {
		logger.Error()
		return nil, err
	}
}
