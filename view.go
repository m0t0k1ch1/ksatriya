package ksatriya

type View struct {
	renderer Renderer
}

func (v *View) Renderer() Renderer {
	return v.renderer
}
func (v *View) SetRenderer(val Renderer) {
	v.renderer = val
}

func (v *View) Render() (string, error) {
	return v.Renderer().Render()
}

func NewView() *View {
	return &View{}
}
