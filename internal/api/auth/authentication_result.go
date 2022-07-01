package auth

import (
	"time"

	"github.com/trino-network/pay-srv/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
