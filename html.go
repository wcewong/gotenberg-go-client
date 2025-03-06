package gotenberg

// HTMLRequest facilitates HTML conversion
// with the Gotenberg API.
type HTMLRequest struct {
	index  Document
	assets []Document

	*chromeRequest
}

// NewHTMLRequest create HTMLRequest.
func NewHTMLRequest(index Document) *HTMLRequest {
	return &HTMLRequest{index, []Document{}, newChromeRequest()}
}

// Assets sets assets form files.
func (req *HTMLRequest) Assets(assets ...Document) {
	req.assets = assets
}

func (req *HTMLRequest) postURL() string {
	return "/convert/html"
}

func (req *HTMLRequest) formFiles() map[string]Document {
	files := make(map[string]Document)
	files["index.html"] = req.index
	if req.header != nil {
		files["header.html"] = req.header
	}
	if req.footer != nil {
		files["footer.html"] = req.footer
	}
	for _, asset := range req.assets {
		files[asset.Filename()] = asset
	}
	return files
}

// Compile-time checks to ensure type implements desired interfaces.
var (
	_ = Request(new(HTMLRequest))
)

// SetFormValue sets a custom form value for the HTMLRequest.
// This method allows you to add or modify form values that are sent to the Gotenberg API.
// For example, you can use this to set margins (e.g., "marginTop", "marginBottom") or other
// configuration options supported by the Gotenberg API.
//
// Parameters:
//   - key: The form field key (e.g., "marginTop", "marginBottom").
//   - value: The value to set for the form field (e.g., "0", "10mm").
//
// Example:
//
//	req := NewHTMLRequest(index)
//	req.SetFormValue("marginTop", "0")
//	req.SetFormValue("marginBottom", "0")
func (req *HTMLRequest) SetFormValue(key, value string) {
	if req.chromeRequest != nil {
		req.chromeRequest.values[key] = value
	}
}
