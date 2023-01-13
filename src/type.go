package src

// Body struct is used to store the url and the fielname from the request
type Body struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}

// Response is used to send the response
type Response struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	SourceUrl string `json:"source_url"`
}

