package health

//go:generate mockgen -destination=./healthcheck_mock_test.go -package=health_test github.com/sazzer/wyrdwest/service/internal/health Healthcheck

// Healthcheck represents a way to check the health of the system
type Healthcheck interface {
	// CheckHealth checks the health of the system
	CheckHealth() error
}
