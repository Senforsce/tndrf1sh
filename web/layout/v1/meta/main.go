package meta

import (
	"net/url"
	"strings"

	"github.com/senforsce/tndr0cean/router"
)

type TextEditorConfig struct {
	RelativeFilePath string
	PayloadSelector  string
	Payload          string
	Contents         string
	LibraryCss       string
}

type SceneConfig struct {
	Mode            string
	PayloadSelector string
	TargetSelector  string
}

type ComponentConfig struct {
	Ctx              *router.Context
	MetaConfig       Config
	TextEditorConfig TextEditorConfig
	SceneConfig      SceneConfig
}

func NewComponentConfig(ctx *router.Context) ComponentConfig {
	return ComponentConfig{
		Ctx: ctx,
	}
}

func (cc *ComponentConfig) WithSceneConfig(sc SceneConfig) ComponentConfig {
	cc.SceneConfig = sc
	return *cc
}

func (cc *ComponentConfig) WithMetaConfig(mc Config) ComponentConfig {
	cc.MetaConfig = mc
	return *cc
}

func (cc *ComponentConfig) WithTextEditorConfig(tec TextEditorConfig) ComponentConfig {
	cc.TextEditorConfig = tec
	return *cc
}

func ShouldInspect(url *url.URL) bool {
	u, err := url.Parse(url.String())
	if err != nil {
		panic(err)
	}

	var shouldInspect bool = false

	if strings.Contains(u.Path, "/preview") {
		shouldInspect = true
	}

	return shouldInspect
}

type Config struct {
	TargetSelector string
	Payload        map[string]string
	Inspect        bool
}
