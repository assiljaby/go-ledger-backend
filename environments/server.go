package environments

import "sync"

type Server struct {
	Port        string
	Environment string
}

var (
	server     Server
	serverOnce sync.Once
)

func GetServer() Server {
	serverOnce.Do(func() {
		server = Server{
			Port:        getenv("PORT", "8080"),
			Environment: getenv("ENVIRONMENT", "development"),
		}
	})

	return server
}
