package handlers

import "github.com/aterip/agata/internal/logging"

type handler struct {
	logger logging.Logger
}

func NewHandler(logger logging.Logger) *handler {
	return &handler{
		logger: logger,
	}
}
