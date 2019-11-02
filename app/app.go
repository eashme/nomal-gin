package app

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const configPath = "./config/app.yaml"

var (
	config *Config
)

type app struct {
	g         *gin.Engine
	redisPool *redis.Pool
	db        *sql.DB

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
}

// 连接数据库
func (a *app) connectDB() {
	db, err := sql.Open("mysql", "xxxxxxxxxxx")
	if err != nil {
		errlog.Fatalf("Failed to connect to database, err is: %v", err)
	}
	a.db = db
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
			log.Fatalf("listen: %s\n", err)
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
	log.Println("Shutdown Server ...")
	// 先关闭gin服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}

func (a *app) Close() {
	// 异步关闭redis
	go func() {
		errlog.Info("closing redis....")
		err := a.redisPool.Close()
		if err != nil {
			errlog.Error("redis pool does not close normally", err)
		} else {
			errlog.Info("redis closed success")
		}
	}()
	// 同步关闭mysql
	defer func() {
		errlog.Info("closing mysql....")
		err := a.db.Close()
		if err != nil {
			errlog.Error("mysql does not close normally", err)
		} else {
			errlog.Info("mysql closed success")
		}
	}()

}
