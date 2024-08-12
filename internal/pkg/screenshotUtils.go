package types

type InputBodyViewport struct {
	Width  int64 `json:"width" form:"width"`
	Height int64 `json:"height" form:"height"`
}
type InputBody struct {
	Source   string             `json:"source" form:"source"` // Строка в виде сырого html
	Viewport *InputBodyViewport `json:"viewport,omitempty" form:"viewport,omitempty"`
}
