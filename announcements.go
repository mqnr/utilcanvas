package utilcanvas

// ListCalendarEvents retrieves a list of announcements for the user.
func (c *CanvasClient) ListAnnouncements(options ApiParameter) ([]DiscussionTopic, error) {
	var announcements []DiscussionTopic

	err := c.Get(ListAnnouncementsEndpoint, options, &announcements)

	return announcements, err
}
