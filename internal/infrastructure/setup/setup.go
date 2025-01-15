package setup

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rodrigosscode/easy-user/configs"
	usecase "github.com/rodrigosscode/easy-user/internal/application/usecase/user"
	infraDb "github.com/rodrigosscode/easy-user/internal/infrastructure/db"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/http"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/logger"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/mapper"
	"github.com/rodrigosscode/easy-user/internal/infrastructure/repository"
	"github.com/rodrigosscode/easy-user/internal/interface/delivery/controller"
)

type configuration struct {
	configApp *configs.AppConfig
	webServer http.Server
	db        *repository.UserRepository
	router    http.GinRouter
}

func NewConfig() *configuration {
	return &configuration{}
}

func (c *configuration) InitLogger() *configuration {
	logger.NewZapLogger()

	logger.Info("Log has been successfully configured")
	return c
}

func (c *configuration) WithAppConfig() *configuration {
	var err error
	c.configApp, err = configs.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}
	return c
}

func (c *configuration) WithDB() *configuration {
	db, err := infraDb.NewDbConnection(c.configApp.MySQLHostDsn)
	if err != nil {
		logger.Fatal(err)
	}

	userMapper := mapper.NewUserMapper()
	userErrorMapper := mapper.NewUserErrorMapper()
	c.db = repository.NewUserRepository(db, userMapper, userErrorMapper)
	logger.Info("DB has been successfully configured")
	return c
}

func (c *configuration) WithRouter() *configuration {

	fuc := controller.NewFindUserByIdController(usecase.NewFindByIdUseCase(c.db))
	duc := controller.NewDeleteUserByIdController(usecase.NewDeleteByIdUseCase(c.db))
	suc := controller.NewSaveUserController(usecase.NewSaveUseCase(c.db))
	uuc := controller.NewUpdateUserController(usecase.NewUpdateUseCase(c.db))
	fupc := controller.NewFindUsersPagingController(usecase.NewFindByPageUseCase(c.db))

	c.router = http.NewGinEngine(gin.Default(), duc, fuc, fupc, suc, uuc)
	return c
}

func (c *configuration) WithWebServer() *configuration {
	intPort, err := strconv.ParseInt(c.configApp.ServerPort, 10, 64)
	if err != nil {
		logger.Fatal(err)
	}

	intDuration, err := time.ParseDuration(c.configApp.ServerTimeout + "s")
	if err != nil {
		logger.Fatal(err)
	}

	c.webServer = http.NewWebServer(c.router, intPort, intDuration*time.Second)
	logger.Info("Web server has been successfully configurated")
	return c
}

func (c *configuration) Start(ctx context.Context, wg *sync.WaitGroup) {
	logger.Info("App running on port %s", c.configApp.ServerPort)
	c.webServer.Listen(ctx, wg)

}
