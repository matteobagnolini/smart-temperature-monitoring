package http

type JsonData struct {
	Temps             []float32
	Avg               float32
	Max               float32
	Min               float32
	CurrState         string
	WindowOpeningPerc int
}
