package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/xylonx/s3-gateway/config"
	"github.com/xylonx/s3-gateway/internal/router"
	"github.com/xylonx/s3-gateway/util"
)

var rootCmd = &cobra.Command{
	Use:   "s3-gateway",
	Short: "a gateway that expose the standard S3 API and proxy to the real storage location",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return setupProject()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return server()
	},
}

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "./settings.yaml", "specify config file path")
	rootCmd.AddCommand(rootCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("Error: execute rootCmd failed: %v\n", err)
		return err
	}
	return nil
}

// TODO: init the config of the project
func setupProject() error {
	if err := config.SetupConfig(cfgFile); err != nil {
		return err
	}

	if err := util.SetupLogger(); err != nil {
		return err
	}

	return nil
}

func server() error {
	r := mux.NewRouter()

	// add s3 standard API routes
	router.SetupRouter(r)

	srv := &http.Server{
		Addr: config.Config.Application.Host + ":" + config.Config.Application.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.Config.Application.ReadTimeout,
		ReadTimeout:  config.Config.Application.WriteTimeout,
		IdleTimeout:  config.Config.Application.IdleTimeout,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), config.Config.Application.ShutdownWait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return nil
}
