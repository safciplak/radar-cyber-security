package json

type JSONResponse struct {
	JSONItem []JSONItem `json:"response"`
}

type JSONItem struct {
	Ts       int64   `json:"ts"`
	SourceIP string  `json:"source_ip"`
	URLItem  URLItem `json:"url_item"`
	Size     int     `json:"size"`
	Note     string  `json:"note"`
}

type URLItem struct {
	Scheme string `json:"schema"`
	Host   string `json:"host"`
	Path   string `json:"path"`
	Opaque string `json:"opaque"`
}
