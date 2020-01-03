package main

import "time"

// AnnotationsReq encodes the information provided by Grafana in its requests.
type AnnotationsReq struct {
	Range      Range      `json:"range"`
	Annotation Annotation `json:"annotation"`
}

// Range specifies the time range the request is valid for.
type RangeRaw struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type Range struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
	Raw  RangeRaw  `json:"raw"`
}
// Annotation is the object passed by Grafana when it fetches annotations.
//
// http://docs.grafana.org/plugins/developing/datasources/#annotation-query
type Annotation struct {
	// Name must match in the request and response
	Name string `json:"name"`

	Datasource string `json:"datasource"`
	IconColor  string `json:"iconColor"`
	Enable     bool   `json:"enable"`
	ShowLine   bool   `json:"showLine"`
	Query      string `json:"query"`
}

// AnnotationResponse contains all the information needed to render an
// annotation event.
//
// https://github.com/grafana/simple-json-datasource#annotation-api
type AnnotationResponse struct {
	// The original annotation sent from Grafana.
	Annotation Annotation `json:"annotation"`
	// Time since UNIX Epoch in milliseconds. (required)
	Time int64 `json:"time"`
	// The title for the annotation tooltip. (required)
	Title string `json:"title"`
	// Tags for the annotation. (optional)
	Tags string `json:"tags"`
	// Text for the annotation. (optional)
	Text string `json:"text"`
}
type SearchReq struct {
	Target      string      `json:"target"`
}
type SearchResponse struct {

}

type Target struct {
	Target string     `json:"target"`
	RefId  string     `json:"refId"`
	Type   string     `json:"type"`
}
type QueryReq struct {
	PanelId int64     `json:"panelId"`
	RequestId string   `json:"requestId"`
	Range      Range   `json:"range"`
	RangeRaw RangeRaw  `json:"rangeRaw"`
	Interval string    `json:"interval"`
	IntervalMs int64   `json:"intervalMs"`
	Targets []Target   `json:"targets`
	Format string      `json:"format"`
	MaxDataPoints int64   `json:"maxDataPoints"`
}

type QueryTimeSerieResponse struct {
	Target     string          `json:"target"`
	Datapoints [][]interface{} `json:"datapoints"`
}
type QueryTableResponse struct {
	Columns []struct {
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"columns"`
	Rows [][]interface{} `json:"rows"`
	Type string          `json:"type"`
}