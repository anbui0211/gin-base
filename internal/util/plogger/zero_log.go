package plogger

import (
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func requestID() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), err
}

func GetLogger() (*zerolog.Logger, string, error) {
	requestID, err := requestID()
	if err != nil {
		return nil, "", err
	}

	var logger zerolog.Logger
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000000"}
	logger = zerolog.
		New(output).
		With().
		Timestamp().
		Caller().
		Str("requestID", requestID).Logger()
	return &logger, requestID, nil
}
