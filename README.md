# Utilcanvas

Utilcanvas is a package which provides some utilities to interface with the Canvas LMS Api.

## Example usage

```go
c := utilcanvas.NewClient("https://canvas.instructure.com", "<OAuth2 Token>")

course := c.Course(12345)

calendarEvents, err := c.ListCalendarEvents(utilcanvas.ApiParameter{
	"context_codes[]": course.CtxCode()})
if err != nil {
	// handle err
}

fmt.Println(calendarEvents)
```