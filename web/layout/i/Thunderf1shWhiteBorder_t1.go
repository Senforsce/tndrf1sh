// Code generated by Tndr(t1) - DO NOT EDIT.

// Tndr(t1): version: v0.0.3

package i

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/senforsce/tndr"
import "context"
import "io"
import "bytes"

func Thunderf1shWhiteBorder() t1.Component {
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
		_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteString("<svg fill=\"#000000\" height=\"800px\" width=\"800px\" version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" viewBox=\"0 0 511.998 511.998\" xml:space=\"preserve\"><g><g><g><path d=\"M304.356,255.566c-4.355,0-7.885,3.53-7.885,7.885c0,6.955-5.657,12.613-12.612,12.613s-12.613-5.658-12.613-12.613
\n				c0-4.355-3.53-7.885-7.885-7.885s-7.885,3.53-7.885,7.885c0,6.955-5.658,12.613-12.613,12.613s-12.612-5.658-12.612-12.613
\n				c0-4.355-3.53-7.885-7.885-7.885s-7.885,3.53-7.885,7.885c0,15.65,12.733,28.383,28.382,28.383c8.05,0,15.327-3.368,20.498-8.769
\n				c5.17,5.402,12.448,8.769,20.498,8.769c15.65,0,28.382-12.733,28.382-28.383C312.241,259.097,308.71,255.566,304.356,255.566z\"></path> <path d=\"M183.11,230.953c-4.355,0-7.885,3.53-7.885,7.885v8.706c0,4.355,3.53,7.885,7.885,7.885s7.885-3.53,7.885-7.885v-8.706
\n				C190.995,234.483,187.465,230.953,183.11,230.953z\"></path> <path d=\"M422.926,203.377c-3.236-6.992-10.032-11.336-17.737-11.336h-76.972l32.138-172.285c1.405-7.533-2.3-14.806-9.219-18.099
\n				c-6.919-3.292-14.898-1.58-19.859,4.259L91.913,287.763c-4.988,5.873-6.077,13.865-2.841,20.857
\n				c3.235,6.992,10.032,11.336,17.737,11.336h76.976l-32.14,172.284c-1.406,7.533,2.298,14.806,9.217,18.099
\n				c2.347,1.118,4.814,1.658,7.246,1.658c4.738,0,9.336-2.056,12.615-5.917l239.364-281.846
\n				C425.072,218.362,426.162,210.37,422.926,203.377z M408.065,214.027L168.701,495.874c-0.208,0.245-0.443,0.52-1.061,0.227
\n				c-0.618-0.294-0.551-0.652-0.493-0.967l33.882-181.616c0.43-2.305-0.188-4.682-1.686-6.485c-1.498-1.803-3.721-2.846-6.065-2.846
\n				h-86.469c-2.176,0-3.12-1.531-3.424-2.189c-0.305-0.658-0.86-2.369,0.549-4.028L343.298,16.123
\n				c0.144-0.169,0.301-0.354,0.59-0.354c0.128,0,0.282,0.037,0.472,0.127c0.618,0.294,0.552,0.652,0.493,0.967L310.974,198.48
\n				c-0.43,2.305,0.188,4.682,1.686,6.485c1.498,1.803,3.721,2.847,6.065,2.847h86.463c2.176,0,3.12,1.531,3.424,2.189
\n				C408.918,210.658,409.475,212.367,408.065,214.027z\"></path> <path d=\"M343.613,230.953c-4.355,0-7.885,3.53-7.885,7.885v8.706c0,4.355,3.53,7.885,7.885,7.885c4.355,0,7.885-3.53,7.885-7.885
\n				v-8.706C351.498,234.483,347.968,230.953,343.613,230.953z\"></path></g></g></g></svg>")
		if t1_7745c5c3_Err != nil {
			return t1_7745c5c3_Err
		}
		if !t1_7745c5c3_IsBuffer {
			_, t1_7745c5c3_Err = t1_7745c5c3_Buffer.WriteTo(t1_7745c5c3_W)
		}
		return t1_7745c5c3_Err
	})
}