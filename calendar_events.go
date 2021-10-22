package utilcanvas

import "fmt"

// ListCalendarEvents retrieves a list of calendar events for the user.
func (c *CanvasClient) ListCalendarEvents(options ApiParameter) ([]CalendarEvent, error) {
	var calendarEvents []CalendarEvent

	err := c.Get(ListCalendarEventsEndpoint, options, &calendarEvents)

	return calendarEvents, err
}

// ListCalendarEventsForUser retrieves a list of calendar events
// or assignments for the specified user.
// To view calendar events for a user other than that of the client,
// you must either be an observer of that user or an administrator.
func (c *CanvasClient) ListCalendarEventsForUser(userId int, options ApiParameter) ([]CalendarEvent, error) {
	var calendarEvents []CalendarEvent

	err := c.Get(fmt.Sprintf("/api/v1/users/%v/calendar_events", userId), options, &calendarEvents)

	return calendarEvents, err
}

// GetSingleCalendarEventOrAssignment gets a singleevent or assigment
// by its specified ID.
func (c *CanvasClient) GetSingleCalendarEventOrAssignment(id int) ([]CalendarEvent, error) {
	var calendarEvents []CalendarEvent

	options := ApiParameter{}
	err := c.Get(fmt.Sprintf("/api/v1/calendar_events/%v", id), options, &calendarEvents)

	return calendarEvents, err
}
