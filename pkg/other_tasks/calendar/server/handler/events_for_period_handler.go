package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"time"
)

type eventsForPeriodServer struct {
	transport      eventsForPeriodTransport
	calendar       service
	errorTransport errorTransport
}

func (c *eventsForPeriodServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data   types.DateIntervalHandlerData
		events []types.Event
		err    error
	)

	mainCtx := r.Context()
	ctx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	switch r.Method {
	case http.MethodGet:
		data, err = c.transport.DecodeRequest(r)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
			return
		}
		events, err = c.calendar.EventsForPeriod(r.Context(), data)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
			return
		}
		if err = c.transport.EncodeResponse(w, events); err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
			return
		}
	default:
		c.errorTransport.EncodeError(w, err, http.StatusMethodNotAllowed)
		return
	}
}

func NewEventsForPeriodServer(transport eventsForPeriodTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &eventsForPeriodServer{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
