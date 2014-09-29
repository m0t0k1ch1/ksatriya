package ksatriya

const (
	TmplDirPathDefault  = "view"
	BaseTmplPathDefault = ""
)

type RenderArgs map[string]interface{}

type RenderConfig struct {
	TmplDirPath  string
	BaseTmplPath string
}

func NewRenderConfig() *RenderConfig {
	return &RenderConfig{
		TmplDirPath:  TmplDirPathDefault,
		BaseTmplPath: BaseTmplPathDefault,
	}
}

type View struct {
	Renderer     Renderer
	RenderArgs   RenderArgs
	RenderConfig *RenderConfig
}

func NewView() *View {
	return &View{
		RenderArgs:   RenderArgs{},
		RenderConfig: NewRenderConfig(),
	}
}

func (v *View) Render() string {
	return v.Renderer.Render(v.RenderConfig, v.RenderArgs)
}
