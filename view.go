package ksatriya

type View struct {
	renderer     Renderer
	renderConfig *RenderConfig
	renderArgs   RenderArgs
}

func NewView() *View {
	return &View{
		renderConfig: NewRenderConfig(),
		renderArgs:   RenderArgs{},
	}
}

func (v *View) Renderer() Renderer {
	return v.renderer
}
func (v *View) SetRenderer(val Renderer) {
	v.renderer = val
}

func (v *View) RenderConfig() *RenderConfig {
	return v.renderConfig
}

func (v *View) RenderArgs() RenderArgs {
	return v.renderArgs
}
func (v *View) SetRenderArg(key string, val interface{}) {
	v.renderArgs[key] = val
}

func (v *View) Render() string {
	return v.Renderer().Render(v.RenderConfig(), v.RenderArgs())
}
