package svc

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vhttp"
	v_sqlx "github.com/imamponco/v-gin-boilerplate/src/pkg/vsqlx"
	"github.com/imamponco/v-gin-boilerplate/src/pkg/vtype"
	"github.com/imamponco/v-gin-boilerplate/src/svc/contract"
	"github.com/nbs-go/nlogger"

	"net/http"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

type App struct {
	*contract.AppContract
}

func NewAPI(config *contract.Config) (*App, error) {
	os.Setenv("TZ", "Asia/Jakarta")
	// Init database
	_, err := v_sqlx.InitDatabase(config)
	if err != nil {
		return nil, err
	}

	// Init logger
	logger := nlogger.Get()

	return &App{
		AppContract: &contract.AppContract{
			Config:       config,
			Repositories: contract.RepositoryContract{
				// TODO: Init repository contract here
			},
			Services: contract.ServiceContract{
				// TODO: Ini service contract here
			},
			Adapters: contract.AdapterContract{},
			Log:      nlogger.Get(),
			ResponseHandler: vhttp.NewResponseHandler(vhttp.ResponseHandlerOptions{
				Debug:  vtype.ParseBooleanFallback(config.Debug, false),
				Logger: logger,
			}),
		},
	}, nil
}

func (a *App) Boot() error {
	// Init service
	err := InitStruct(&a.Services, a.initService)
	if err != nil {
		return err
	}

	// Init repository
	err = InitStruct(&a.Repositories, a.initRepository)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) InitRouter(workDir string) (*gin.Engine, error) {
	// Init router gin
	router := gin.Default()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	// Set Strict Secure
	router = app.SetSecureHTTP(router, app.Config)

	// Set trust proxies
	router, err := app.SetTrustProxies(router, app.Config.ServerTrustProxy)
	if err != nil {
		return nil, err
	}

	// Set cors
	router, err = app.CORS(router, app.Config.CORS)
	if err != nil {
		return nil, err
	}

	// Init controllers
	controllers := NewController(app)

	// Setup routes
	r := SetupRoute(router, controllers)

	// Setup static
	staticDir := path.Join(workDir, "/public")
	r.Use(static.ServeRoot("/assets", staticDir))

	return r, nil
}

func (app *App) SetSecureHTTP(r *gin.Engine, config *contract.Config) *gin.Engine {
	// Generalize env
	stage := strings.ToLower(config.Stage)

	schema := "http"
	if vtype.ParseBooleanFallback(config.ServerSecure, false) {
		schema = "https"
	}

	strictSecure := secure.New(secure.Config{
		SSLRedirect:           true,
		IsDevelopment:         true,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		IENoOpen:              true,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": schema},
	})

	if stage == "production" || config.ServerSecure {
		strictSecure = secure.New(secure.Config{
			AllowedHosts:          []string{},
			SSLRedirect:           true,
			SSLHost:               "",
			STSSeconds:            315360000,
			STSIncludeSubdomains:  true,
			FrameDeny:             true,
			ContentTypeNosniff:    true,
			BrowserXssFilter:      true,
			ContentSecurityPolicy: "default-src 'self'",
			IENoOpen:              true,
			ReferrerPolicy:        "strict-origin-when-cross-origin",
			SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": schema},
		})
	}

	// Set secure based on env
	r.Use(strictSecure)

	return r
}

func (app *App) SetTrustProxies(r *gin.Engine, trustProxy string) (*gin.Engine, error) {
	// Set trusted proxies if allow all
	if trustProxy == "*" {
		return r, nil
	}

	// Set trusted proxies
	if trustProxy != "" && trustProxy != "*" {
		proxies := strings.Split(trustProxy, ",")

		err := r.SetTrustedProxies(proxies)
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	// Set false trust all proxies
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (app *App) CORS(r *gin.Engine, configCORS string) (*gin.Engine, error) {
	// Set cors for all
	if configCORS == "*" {
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodOptions,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				"Accept",
				"Content-Type",
				"Authorization",
				"Origin",
			},
			ExposeHeaders: []string{
				"Content-Length",
				"Content-Type",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		return r, nil
	}

	// Set cors if sets
	if configCORS != "" && configCORS != "*" {
		corsDomains := strings.Split(configCORS, ",")
		r.Use(cors.New(cors.Config{
			AllowOrigins: corsDomains,
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodOptions,
				http.MethodDelete,
			},
			AllowHeaders: []string{
				"Accept",
				"Content-Type",
				"Authorization",
				"Origin",
			},
			ExposeHeaders: []string{
				"Content-Length",
				"Content-Type",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		return r, nil
	}

	return r, nil
}

// / InitStruct reflect fields in struct and run initializer function
func InitStruct(s interface{}, initFn func(name string, i interface{}) error) error {
	// Reflect on struct element
	rv := reflect.ValueOf(s).Elem()

	// Iterate fields
	for i := 0; i < rv.NumField(); i++ {
		// Get field interface
		fieldValue := rv.Field(i).Interface()
		fieldName := rv.Type().Field(i).Name

		err := initFn(fieldName, fieldValue)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initRepository(name string, i interface{}) error {
	// Check interface
	r, ok := i.(contract.RepositoryInitializer)
	if !ok {
		return fmt.Errorf("repository '%s' does not implement repository.Initializer interface", name)
	}

	// Init repository
	err := r.Init(&a.Adapters)
	if err != nil {
		return err
	}

	a.Log.Debugf("Repositories.%s has been initialized", name)

	return nil
}

func (a *App) initService(name string, i interface{}) error {
	// Check interface
	r, ok := i.(contract.ServiceInitializer)
	if !ok {
		return fmt.Errorf("service '%s' does not implement ServiceInitializer interface", name)
	}

	// Init service
	err := r.Init(a.AppContract)
	if err != nil {
		return err
	}

	a.Log.Debugf("Services.%s has been initialized", name)

	return nil
}
