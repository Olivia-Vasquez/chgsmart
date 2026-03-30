package render

import (
	"encoding/json"

	"github.com/Olivia-Vasquez/chgsmart/internal/model" // import the model package to use model.Options and model.Item
)

// JSONRender renders the options as a JSON string, including the suggested version bump.
func JSONRender(opt model.Options, suggestedBump string) string {
	title := opt.Title
	if title == "" {
		title = "Unreleased"
	}

	out := struct {
		Title         string       `json:"title"`
		Items         []model.Item `json:"items"`
		SuggestedBump string       `json:"suggested_bump,omitempty"`
	}{
		Title:         title,
		Items:         opt.Items,
		SuggestedBump: suggestedBump,
	}

	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		return "{}"
	}

	return string(data)
}
