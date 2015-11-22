package ksatriya

type View struct {
	renderer Renderer
}

func NewView() *View {
	return &View{}
}

func (v *View) Renderer() Renderer {
	return v.renderer
}
func (v *View) SetRenderer(val Renderer) {
	v.renderer = val
}

func (v *View) Render() string {
	return v.Renderer().Render()
}
