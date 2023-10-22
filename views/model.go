package views

type ViewsData struct {
	Card        string
	Title       string
	ImageURL    string `default:""`
	VideoURL    string `default:""`
	Width       int    `default:"400"`
	Height      int    `default:"400"`
	URL         string
	Description string
}