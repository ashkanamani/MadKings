package cmd

import (
	"fmt"
	"github.com/ashkanamani/madkings/internal/repository"
	"github.com/ashkanamani/madkings/internal/repository/redis"
	"github.com/ashkanamani/madkings/internal/service"
	"github.com/ashkanamani/madkings/internal/telegram"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	_ = godotenv.Load()
	fmt.Println("serve called")

	// setup repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		logrus.WithError(err).Fatalln("could not connect to redis server")
	}
	accountRepository := repository.NewAccountRedisRepository(redisClient)

	// setup app
	app := service.NewApp(service.NewAccountService(accountRepository))

	// setup telegram

	tg, err := telegram.NewTelegram(app, os.Getenv("BOT_TOKEN"))
	if err != nil {
		logrus.WithError(err).Fatalln("could not create telegram client")
	}
	tg.Start()
}

func init() {
	rootCmd.AddCommand(serveCmd)

}
