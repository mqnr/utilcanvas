package utilcanvas

// Various context code types that the API uses.
const (
	ContextCodeTypeCourse = "course"
	ContextCodeTypeUser   = "user"
)

// A CanvasClient holds the API URL and the API key over which
// subsequent calls to the Canvas API will be made.
type CanvasClient struct {
	ApiUrl string
	ApiKey string
}

// A ContextCode represents a context that the Canvas API uses
// to present relevant information to the user.
type ContextCode struct {
	Type string
	Id   string
}

type GroupDiscussion struct {
	Id      int `json:"id"`
	GroupId int `json:"group_id"`
}

// A Course holds information for a course. Only an extremely
// limited subset of the fields managed by Canvas has been implemented.
type Course struct {
	Id            int    `json:"id"`
	SisCourseId   string `json:"sis_course_id"`
	Uuid          string `json:"uuid"`
	IntegrationId string `json:"integration_id"`
	SisImportId   int    `json:"sis_import_id"`
	Name          string `json:"name"`
	CourseCode    string `json:"course_code"`
}

// A CalendarEvent holds information for either a discussion
// topic or a course announcement.
type DiscussionTopic struct {
	Id                      int               `json:"id"`
	Title                   string            `json:"title"`
	Message                 string            `json:"message"`
	HtmlUrl                 string            `json:"html_url"`
	PostedAt                string            `json:"posted_at"`
	LastReplyAt             string            `json:"last_reply_at"`
	RequireInitialPost      bool              `json:"require_initial_post"`
	UserCanSeePosts         bool              `json:"user_can_see_posts"`
	DiscussionSubentryCount int               `json:"discussion_subentry_count"`
	ReadState               string            `json:"read_state"`
	UnreadCount             int               `json:"unread_count"`
	Subscribed              string            `json:"subscribed"`
	SubscriptionHold        string            `json:"subscription_hold"`
	AssignmentId            int               `json:"assignment_id"`
	DelayedPostAt           string            `json:"delayed_post_at"`
	Published               bool              `json:"published"`
	LockAt                  string            `json:"lock_at"`
	Locked                  bool              `json:"locked"`
	Pinned                  bool              `json:"pinned"`
	LockedForUser           bool              `json:"locked_for_user"`
	LockInfo                string            `json:"lock_info"`
	LockExplanation         string            `json:"lock_explanation"`
	UserName                string            `json:"user_name"`
	GroupTopicChildren      []GroupDiscussion `json:"group_topic_children"`
	RootTopicId             int               `json:"root_topic_id"`
	PodcastUrl              string            `json:"podcast_url"`
	DiscussionType          string            `json:"discussion_type"`
	GroupCategoryId         int               `json:"group_category_id"`
	Attachments             string            `json:"attachments"`
	Permissions             string            `json:"permissions"`
	AllowRating             bool              `json:"allow_rating"`
	OnlyGradersCanRate      bool              `json:"only_graders_can_rate"`
	SortByRating            bool              `json:"sort_by_rating"`
}

// A CalendarEvent holds information for a calendar event.
// The fields in this object are a subset of those that the
// API returns.
type CalendarEvent struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	StartAt         string `json:"start_at"`
	EndAt           string `json:"end_at"`
	Description     string `json:"description"`
	LocationName    string `json:"location_name"`
	LocationAddress string `json:"location_address"`
	ContextCode     string `json:"context_code"`
	ContextName     string `json:"context_name"`
	AllContextCodes string `json:"all_context_codes"`
	Hidden          bool   `json:"hidden"`
	Url             string `json:"url"`
	HtmlUrl         string `json:"html_url"`
	AllDayDate      string `json:"all_day_date"`
	AllDay          bool   `json:"all_day"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}

/*
func NewContextCodeList() []ContextCode {
	f
}
*/
