package utilcanvas

import (
	"net/url"
	"strconv"
)

// The ApiParameter type is used to map API options
// to their desired values.
type ApiParameter map[string]interface{}

// Process adds all parameters in the ApiParameter object
// to the Values map provided by package "url".
func (p ApiParameter) Process() (url.Values, error) {
	urlValues := url.Values{}
	// Properly formats all legal key-value pairs.
	for key, value := range p {
		switch t := value.(type) {
		case string:
			urlValues.Add(key, t)
		case bool:
			urlValues.Add(key, strconv.FormatBool(t))
		case *[]ContextCode:
			if len(*t) > 0 {
				// If the context codes slice isn't empty,
				// adds all of them to the request.
				for _, e := range *t {
					urlValues.Add("context_codes[]", e.Type+"_"+e.Id)
				}
			}
		}
	}
	return urlValues, nil
}
