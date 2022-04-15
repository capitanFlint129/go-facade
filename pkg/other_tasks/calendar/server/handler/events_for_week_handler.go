package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForWeekServer struct {
	transport      eventsForWeekTransport
	calendar       service
	errorTransport errorTransport
}

func (c *eventsForWeekServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data   types.DateHandlerData
		events []types.Event
		err    error
	)
	if r.Method != http.MethodPost {
		c.errorTransport.EncodeError(w, err, http.StatusMethodNotAllowed)
		return
	}

	data, err = c.transport.DecodeRequest(r)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	events, err = c.calendar.EventsForWeek(ctx, data)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = c.transport.EncodeResponse(w, events); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewEventsForWeekServer(transport eventsForWeekTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &eventsForWeekServer{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}