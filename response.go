package alpr

// Response :
type Response struct {
	Plate     string `json:"plate"`
	Region    string `json:"region"`
	Color     string `json:"color"`
	Make      string `json:"make"`
	BodyType  string `json:"body_type"`
	MakeModel string `json:"make_model"`
}

var instanceResponse *Response

// NewResponse : inisialisasi struct Response
func NewResponse() *Response {
	if instanceResponse == nil {
		instanceResponse = new(Response)
	}
	return instanceResponse
}
