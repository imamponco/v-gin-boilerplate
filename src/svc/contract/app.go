package contract

import (
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vhttp"
	"github.com/nbs-go/nlogger"
)

type AppContract struct {
	Config          *Config
	Services        ServiceContract
	Repositories    RepositoryContract
	Adapters        AdapterContract
	Log             nlogger.Logger
	ResponseHandler vhttp.ResponseHandler
}

type ServiceContract struct {
	// TODO: Set your service contract
}

type RepositoryContract struct {
	// TODO: Set your repository contract
}

type AdapterContract struct {
	// TODO: Set your adapter contract
}
