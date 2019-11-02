package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io"
	"log"
	"net/http"
	"nomal-gin/cache"
	"nomal-gin/utils/errlog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const configPath = "./config/app.yaml"

var (
	config *Config
)

type app struct {
	g     *gin.Engine
	cache cache.ICache
	db    *sql.DB
}

func New() *app {
	app := new(app)
	app.loadConfig()
	app.loadLog()
	app.connectDB()
	app.initRouter()
	return app
}

// 加载配置文件
func (a *app) loadConfig() {
	config = loadConfig(configPath)
	// todo 可加载多套配置文件，根据开发环境或生产环境进行选择

}

// 加载日志配置
func (a *app) loadLog() {
	// todo 根据配置文件,加载日志配置
	a.loadAccessLog()
	a.loadErrorLog()
}

// 连接数据库
func (a *app) connectDB() {
	db, err := sql.Open("mysql", "xxxxxxxxxxx")
	if err != nil {
		errlog.Fatalf("Failed to connect to database, err is: %v", err)
	}
	a.db = db
}

func (a *app) connectCache() {
	var (
		redisHost string
		redisPort string
		err       error
	)
	if redisHost = os.Getenv("ENV_REDIS_HOST"); redisHost != "" {
		config.Redis.Host = redisHost
	}
	if redisPort = os.Getenv("ENV_REDIS_PORT"); redisPort != "" {
		redisPort = fmt.Sprintf("%d", config.Redis.Port)
		config.Redis.Port, err = strconv.ParseInt(redisPort, 10, 64)
		if err != nil {
			errlog.Fatalf("Unknown Redis Port , %v", err)
		}
	}
	a.cache = cache.NewRedisCache(
		config.Redis.Host,
		int(config.Redis.Port),
		config.Redis.Password,
		config.Redis.Db,
		time.Duration(config.Redis.DialTimeOut)*time.Second,
		time.Duration(config.Redis.ReadTimeOut)*time.Second,
		time.Duration(config.Redis.WriteTimeOut)*time.Second,
		time.Duration(config.Redis.PoolTimeOut)*time.Second,
		1,
	)
}

func (a *app) Run() {
	addr := fmt.Sprintf(":%s", config.App.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: a.g,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errlog.Fatalf("listen: %s\n", err)
		}
	}()
	a.gracefulClose(srv)
	a.Close()
}

func (a *app) gracefulClose(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	errlog.Info("Shutdown Server ...")
	// 先关闭gin服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		errlog.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-ctx.Done():
		errlog.Info("timeout of 1 seconds.")
	}
	errlog.Info("Server exiting")
}

func (a *app) Close() {
	// 异步关闭redis
	go func() {
		errlog.Info("closing cache....")
		err := a.cache.Close()
		if err != nil {
			errlog.Error("cache does not close normally ", err)
		} else {
			errlog.Info("cache closed success")
		}
	}()
	// 同步关闭mysql
	defer func() {
		errlog.Info("closing database....")
		err := a.db.Close()
		if err != nil {
			errlog.Error("database does not close normally ", err)
		} else {
			errlog.Info("database closed success")
		}
	}()

}

func (a *app) loadAccessLog() {
	if config.Env {
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.
		f, err := os.Create(config.Log.Path + config.Log.AccessName)
		if err != nil {
			log.Fatalln(err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.ForceConsoleColor()
	}
}

func (a *app) loadErrorLog() {
	var env = config.Env
	if product := os.Getenv("ENV_PRODUCT"); product == "true" {
		env = true
	}
	errlog.NewErrLog(config.Log.Path+config.Log.ErrorName, config.Log.ErrLevel, env)
	errlog.Debug("log started")
}
