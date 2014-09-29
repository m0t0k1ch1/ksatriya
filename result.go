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

type Result struct {
	Renderer     Renderer
	RenderArgs   RenderArgs
	RenderConfig *RenderConfig
}

func NewResult() *Result {
	return &Result{
		RenderArgs:   RenderArgs{},
		RenderConfig: NewRenderConfig(),
	}
}

func (r *Result) Render() string {
	return r.Renderer.Render(r.RenderConfig, r.RenderArgs)
}
