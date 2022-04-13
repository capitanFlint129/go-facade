package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
)

type loggingMiddleware struct {
	service       service
	logDateFormat string
	logger        *logrus.Logger
}

func (l *loggingMiddleware) CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error) {
	createdEvent, err := l.service.CreateEvent(ctx, createEventData)
	l.logger.WithFields(logrus.Fields{
		"user_id": createEventData.UserId,
		"name":    createdEvent.Name,
		"date":    createdEvent.Date.Format(l.logDateFormat),
	}).Info()
	return createdEvent, err
}

func NewLoggingMiddleware(service service, logger *logrus.Logger, logDateFormat string) service {
	return &loggingMiddleware{
		service:       service,
		logger:        logger,
		logDateFormat: logDateFormat,
	}
}