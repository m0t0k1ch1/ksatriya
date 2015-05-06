package ksatriya

type View struct {
	renderer     Renderer
	renderArgs   RenderArgs
	renderConfig *RenderConfig
}

func NewView() *View {
	return &View{
		renderArgs:   RenderArgs{},
		renderConfig: NewRenderConfig(),
	}
}

func (v *View) Renderer() Renderer {
	return v.renderer
}
func (v *View) SetRenderer(val Renderer) {
	v.renderer = val
}

func (v *View) RenderArgs() RenderArgs {
	return v.renderArgs
}
func (v *View) SetRenderArg(key string, val interface{}) {
	v.renderArgs[key] = val
}

func (v *View) RenderConfig() *RenderConfig {
	return v.renderConfig
}

func (v *View) Render() string {
	return v.Renderer().Render(v.RenderConfig(), v.renderArgs)
}
