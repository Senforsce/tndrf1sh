// Code generated by Tndr(t1) - DO NOT EDIT.

// Tndr(t1): version: v0.0.3

package i

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

func Thunderf1shYellowHover() t1.Component {
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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<svg height=\"800px\" width=\"800px\" version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 511.999 511.999\" xml:space=\"preserve\"><path id=\"SVGCleanerId_0\" style=\"fill:#F7D64C;\" d=\"M318.726,199.927h86.463c9.936,0,15.318,11.632,8.886,19.206L174.712,500.979
\n	c-5.824,6.858-16.965,1.554-15.316-7.29l33.882-181.615H106.81c-9.936,0-15.318-11.632-8.886-19.206L337.288,11.02
\n	c5.823-6.858,16.965-1.554,15.316,7.29L318.726,199.927z\"></path> <g><path id=\"SVGCleanerId_0_1_\" style=\"fill:#F7D64C;\" d=\"M318.726,199.927h86.463c9.936,0,15.318,11.632,8.886,19.206
\n		L174.712,500.979c-5.824,6.858-16.965,1.554-15.316-7.29l33.882-181.615H106.81c-9.936,0-15.318-11.632-8.886-19.206L337.288,11.02
\n		c5.823-6.858,16.965-1.554,15.316,7.29L318.726,199.927z\"></path></g> <path style=\"opacity:0.1;enable-background:new    ;\" d=\"M171.906,484.026l33.726-180.778c0.417-2.235-1.298-4.3-3.571-4.3h-83.744
\n	c-9.628,0-15.052-10.716-9.927-18.405l-10.466,12.324c-6.432,7.573-1.049,19.206,8.887,19.206h86.469l-33.882,181.615
\n	c-1.65,8.844,9.491,14.148,15.316,7.29l5.405-6.365C175.249,494.236,170.81,489.899,171.906,484.026z\"></path> <path style=\"opacity:0.3;fill:#FFFFFF;enable-background:new    ;\" d=\"M304.356,155.55c0,11.056,8.005,20.222,18.53,22.075
\n	l8.223-44.086c-1.398-0.272-2.84-0.425-4.317-0.425C314.402,133.113,304.356,143.158,304.356,155.55z M405.189,199.927H357.58
\n	c5.363,17.548,21.681,30.312,40.984,30.312c2.208,0,4.377-0.168,6.495-0.49l9.016-10.616
\n	C420.507,211.559,415.125,199.927,405.189,199.927z M337.288,11.02l-45.067,53.066c1.324,25.435,18.453,46.678,41.75,54.114
\n	l18.633-99.891C354.253,9.466,343.112,4.163,337.288,11.02z\"></path> <path d=\"M183.111,230.954c-4.355,0-7.885,3.53-7.885,7.885v8.706c0,4.355,3.53,7.885,7.885,7.885s7.885-3.53,7.885-7.885v-8.706
\n	C190.996,234.484,187.465,230.954,183.111,230.954z M304.356,255.567c-4.355,0-7.885,3.53-7.885,7.885
\n	c0,6.955-5.657,12.612-12.612,12.612s-12.613-5.658-12.613-12.612c0-4.355-3.53-7.885-7.885-7.885c-4.355,0-7.885,3.53-7.885,7.885
\n	c0,6.955-5.658,12.612-12.613,12.612c-6.955,0-12.612-5.658-12.612-12.612c0-4.355-3.53-7.885-7.885-7.885
\n	c-4.355,0-7.885,3.53-7.885,7.885c0,15.65,12.733,28.382,28.382,28.382c8.05,0,15.327-3.368,20.498-8.769
\n	c5.17,5.402,12.448,8.769,20.498,8.769c15.65,0,28.382-12.733,28.382-28.382C312.241,259.098,308.711,255.567,304.356,255.567z
\n	 M343.614,230.954c-4.355,0-7.885,3.53-7.885,7.885v8.706c0,4.355,3.53,7.885,7.885,7.885c4.355,0,7.885-3.53,7.885-7.885v-8.706
\n	C351.499,234.484,347.969,230.954,343.614,230.954z M422.926,203.378c-3.236-6.992-10.032-11.336-17.737-11.336h-76.972
\n	l32.138-172.286c1.405-7.533-2.3-14.806-9.219-18.099c-6.919-3.292-14.898-1.58-19.859,4.259L91.913,287.764
\n	c-4.988,5.873-6.077,13.865-2.841,20.857c3.235,6.992,10.032,11.336,17.737,11.336h76.976l-32.14,172.284
\n	c-1.406,7.533,2.298,14.806,9.217,18.099c2.347,1.118,4.814,1.658,7.246,1.658c4.738,0,9.336-2.056,12.615-5.917l239.364-281.846
\n	C425.073,218.363,426.162,210.371,422.926,203.378z M408.066,214.028L168.701,495.875c-0.208,0.244-0.443,0.521-1.061,0.227
\n	s-0.551-0.652-0.493-0.967l33.882-181.616c0.43-2.305-0.188-4.682-1.686-6.485c-1.498-1.803-3.721-2.846-6.065-2.846H106.81
\n	c-2.176,0-3.12-1.531-3.424-2.189c-0.305-0.658-0.86-2.368,0.549-4.028L343.299,16.124c0.144-0.169,0.301-0.354,0.59-0.354
\n	c0.128,0,0.282,0.037,0.472,0.127c0.618,0.294,0.552,0.652,0.493,0.967L310.974,198.48c-0.43,2.305,0.188,4.682,1.686,6.485
\n	c1.498,1.803,3.721,2.847,6.065,2.847h86.463c2.176,0,3.12,1.531,3.424,2.189C408.918,210.659,409.475,212.368,408.066,214.028z\"></path></svg>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}