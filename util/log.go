package util

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func SetupLogger(opts ...zap.Option) (err error) {
	Logger, err = zap.NewProduction(opts...)
	if err != nil {
		log.Printf("Error: init zap logger failed: %v\n", err)
		return err
	}
	return nil
}
