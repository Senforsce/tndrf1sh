// Code generated by Tndr(t1) - DO NOT EDIT.

// Tndr(t1): version: v0.0.3

package componentpreview

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

import "github.com/senforsce/tndr0cean/router"
import "github.com/senforsce/tndrf1sh/web/layout/v1/meta"

func wrapChildren(c *router.Context) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var1 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var1 == nil {
			t1_7745c5c3_Var1 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		var t1_7745c5c3_Var2 = []any{c.Get(" cls:componentPreview ").(string)}
		t1_7745c5c3_Err = t1.RenderCSSItems(ctx, t1_7745c5c3_Buffer, t1_7745c5c3_Var2...)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<div class=\"")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString(t1.EscapeString(t1.CSSClasses(t1_7745c5c3_Var2).String()))
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("\">")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = t1_7745c5c3_Var1.Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("</div>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}

func HTMX(c *router.Context, contents t1.Component, cfg meta.ComponentConfig) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
		t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
		if !t1_7745c5c3_IsBuffer {
			t1_7745c5c3_Buffer = t1.GetBuffer()
			defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
		}
		ctx = t1.InitializeContext(ctx)
		t1_7745c5c3_Var3 := t1.GetChildren(ctx)
		if t1_7745c5c3_Var3 == nil {
			t1_7745c5c3_Var3 = t1.NopComponent
		}
		ctx = t1.ClearChildren(ctx)
		t1_7745c5c3_Var4 := t1.ComponentFunc(func(ctx context.Context, t1_7745c5c3_W io.Writer) (t1_7745c5c3_Err error) {
			t1_7745c5c3_Buffer, t1_7745c5c3_IsBuffer := t1_7745c5c3_W.(*bytes.Buffer)
			if !t1_7745c5c3_IsBuffer {
				t1_7745c5c3_Buffer = t1.GetBuffer()
				defer t1.ReleaseBuffer(t1_7745c5c3_Buffer)
			}
			t1_7745c5c3_Err = contents.Render(ctx, t1_7745c5c3_Buffer)
			if t1_7745c5c3_Err != nil {
				return t1_7745c5c3_Err
			}
			if !t1_7745c5c3_IsBuffer {
				_, t1_7745c5c3_Err = io.Copy(t1_7745c5c3_W, t1_7745c5c3_Buffer)
			}
			return t1_7745c5c3_Err
		})
		t1_7745c5c3_Err = wrapChildren(c).Render(t1.WithChildren(ctx, t1_7745c5c3_Var4), t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		t1_7745c5c3_Err = meta.O8(cfg.MetaConfig).Render(ctx, t1_7745c5c3_Buffer)
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}