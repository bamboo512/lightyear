package asynq

import (
	"fmt"
	"lightyear/core/global"
	"lightyear/pkg/tasks"
	"log"

	"github.com/hibiken/asynq"
)

func InitServer() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Network:  "tcp",
			Addr:     global.Config.Redis.Address,
			Username: global.Config.Redis.Username,
			Password: global.Config.Redis.Password,
			DB:       global.Config.Redis.DB,
		},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 4,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 2,
				"default":  1,

				"low": 1,
			},
			// See the godoc for other configuration options
		},
	)

	fmt.Printf("%s@%s:%s\n", global.Config.Redis.Username, global.Config.Redis.Password, global.Config.Redis.Address)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.Handle(tasks.TypeImageTranscode, tasks.NewImageProcessor())
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
