package render

import (
	"encoding/json"

	"github.com/Olivia-Vasquez/chgsmart/internal/model" // import the model package to use model.Options and model.Item
)

// JSONRender renders the options as a JSON string.
func JSONRender(opt model.Options) string {
	title := opt.Title
	if title == "" {
		title = "Unreleased"
	}

	out := struct {
		Title string       `json:"title"`
		Items []model.Item `json:"items"`
	}{
		Title: title,
		Items: opt.Items,
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(data)
}
