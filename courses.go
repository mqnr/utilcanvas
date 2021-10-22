package utilcanvas

import "fmt"

// Course returns a course given its ID.
// As of this moment, it is a mock function.
func (c *CanvasClient) Course(id int) *Course {
	return &Course{Id: id}
}

// CtxCode returns a valid context code for the specified course.
func (co *Course) CtxCode() string {
	return fmt.Sprintf("%v_%v", ContextCodeTypeCourse, co.Id)
}
