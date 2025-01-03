// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package shared

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

func Button(label, class string, attr templ.Attributes) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var2 = []any{fmt.Sprintf("enabled:hover:bg-crust %v %v", DefaultBtnClasses, class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attr)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><span>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 10, Col: 15}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

// w-full h-full inline-flex justify-center items-center rounded-md bg-base px-3 py-2 text-sm font-semibold  hover:bg-crust m-1 transition-colors duration-200
func AddButton(attr templ.Attributes, class, label string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var6 = []any{fmt.Sprintf("enabled:hover:bg-green/25 enabled:hover:text-green %v %v", DefaultBtnClasses, class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var6...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var6).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attr)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if label != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 21, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs("add")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 23, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func RemoveButton(attr templ.Attributes, class, label string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var10 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var10 == nil {
			templ_7745c5c3_Var10 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var11 = []any{fmt.Sprintf("enabled:hover:bg-red/25 enabled:hover:text-red %v %v", DefaultBtnClasses, class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var11...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var12 string
		templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var11).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attr)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if label != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(label)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 34, Col: 33}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var14 string
			templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs("remove")
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 36, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func IconButton(iconName, class string, attr templ.Attributes) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var15 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var15 == nil {
			templ_7745c5c3_Var15 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var16 = []any{fmt.Sprintf("w-full h-full fill-text inline-flex justify-center items-center rounded-md bg-base px-1 py-1 text-sm font-semibold  enabled:hover:bg-crust m-1 transition-colors duration-200 enabled:hover:ring-2 enabled:hover:ring-surface0 %v", class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var16...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var17 string
		templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var16).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/shared/button.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attr)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><svg viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\" class=\"fill-text w-2/3 h-2/3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LabelToIcon(iconName).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</svg></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func placeholderIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var18 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var18 == nil {
			templ_7745c5c3_Var18 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M10.3523 16.1576C10.3523 14.554 10.4958 13.4007 10.7829 12.6979C11.0699 11.9951 11.6787 11.2279 12.6092 10.3964C13.421 9.68362 14.0396 9.06493 14.4653 8.54028C14.891 8.01564 15.1038 7.41674 15.1038 6.74361C15.1038 5.93189 14.8316 5.25875 14.2871 4.7242C13.7427 4.18966 12.9854 3.92238 12.0153 3.92238C11.0056 3.92238 10.2384 4.22925 9.71376 4.84299C9.18911 5.45673 8.8179 6.08037 8.60012 6.71391C8.60012 6.71391 8.09497 7.87014 6.65769 7.45738C5.22042 7.04461 5.54132 5.40724 5.54132 5.40724C5.95708 4.14016 6.7193 3.04137 7.82799 2.11086C8.93669 1.18035 10.3325 0.715092 12.0153 0.715092C14.0941 0.715092 15.6928 1.29419 16.8114 2.45237C17.93 3.61056 18.4893 5.00138 18.4893 6.62482C18.4893 7.61472 18.2764 8.46109 17.8508 9.16392C17.4251 9.86676 16.7569 10.6636 15.8462 11.5545C14.8761 12.4851 14.2871 13.1928 14.0792 13.6779C13.8714 14.1629 13.7674 14.9895 13.7674 16.1576C13.7674 16.1576 13.6801 17.5472 12.0718 17.5472C10.4636 17.5472 10.3523 16.1576 10.3523 16.1576ZM12.0153 23.2849C11.362 23.2849 10.8027 23.0523 10.3374 22.587C9.87215 22.1218 9.63952 21.5625 9.63952 20.9091C9.63952 20.2558 9.87215 19.6965 10.3374 19.2313C10.8027 18.766 11.362 18.5334 12.0153 18.5334C12.6686 18.5334 13.2279 18.766 13.6932 19.2313C14.1584 19.6965 14.3911 20.2558 14.3911 20.9091C14.3911 21.5625 14.1584 22.1218 13.6932 22.587C13.2279 23.0523 12.6686 23.2849 12.0153 23.2849Z\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func memberAddIcon(toogle bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var19 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var19 == nil {
			templ_7745c5c3_Var19 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if toogle {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M19.3914 12.6606C18.8082 12.6606 18.3355 12.1879 18.3355 11.6047L18.3355 9.34254L16.2236 9.34254C15.6405 9.34254 15.1677 8.86979 15.1677 8.28663L15.1677 8.18637C15.1677 7.60321 15.6405 7.13047 16.2236 7.13047L18.3355 7.13047L18.3355 4.86826C18.3355 4.2851 18.8082 3.81236 19.3914 3.81236L19.3914 3.81236C19.9745 3.81236 20.4473 4.2851 20.4473 4.86826L20.4473 7.13047L22.5591 7.13047C23.1423 7.13047 23.615 7.60321 23.615 8.18637L23.615 8.28663C23.615 8.86979 23.1423 9.34254 22.5591 9.34254L20.4473 9.34254L20.4473 11.6047C20.4473 12.1879 19.9745 12.6606 19.3914 12.6606L19.3914 12.6606ZM8.83227 10.4486C7.67077 10.4486 6.67646 10.0154 5.84933 9.14898C5.0222 8.28259 4.60864 7.24107 4.60864 6.02443C4.60864 4.80779 5.0222 3.76627 5.84933 2.89988C6.67646 2.03348 7.67077 1.60028 8.83227 1.60028C9.99377 1.60028 10.9881 2.03348 11.8152 2.89988C12.6423 3.76627 13.0559 4.80779 13.0559 6.02443C13.0559 7.24107 12.6423 8.28259 11.8152 9.14898C10.9881 10.0154 9.99377 10.4486 8.83227 10.4486ZM1.44091 22.3997C0.857747 22.3997 0.385 21.927 0.385 21.3438L0.385 16.2C0.385 15.5732 0.538987 14.9972 0.84696 14.4718C1.15493 13.9464 1.5641 13.5455 2.07445 13.269C3.16556 12.6975 4.27427 12.2689 5.40057 11.9832C6.52687 11.6975 7.67077 11.5546 8.83227 11.5546C9.99377 11.5546 11.1377 11.6975 12.264 11.9832C13.3903 12.2689 14.499 12.6975 15.5901 13.269C16.1004 13.5455 16.5096 13.9464 16.8176 14.4718C17.1256 14.9972 17.2795 15.5732 17.2795 16.2L17.2795 21.3438C17.2795 21.927 16.8068 22.3997 16.2236 22.3997L1.44091 22.3997ZM2.49682 20.1876L15.1677 20.1876L15.1677 16.2C15.1677 15.9972 15.1193 15.8129 15.0225 15.6469C14.9257 15.481 14.7982 15.352 14.6398 15.2598C13.6895 14.7621 12.7303 14.3888 11.7624 14.14C10.7945 13.8911 9.81779 13.7667 8.83227 13.7667C7.84676 13.7667 6.87004 13.8911 5.90213 14.14C4.93421 14.3888 3.97509 14.7621 3.02477 15.2598C2.86639 15.352 2.7388 15.481 2.64201 15.6469C2.54521 15.8129 2.49682 15.9972 2.49682 16.2L2.49682 20.1876ZM8.83227 8.2365C9.41302 8.2365 9.91018 8.0199 10.3237 7.58671C10.7373 7.15351 10.9441 6.63275 10.9441 6.02443C10.9441 5.41611 10.7373 4.89535 10.3237 4.46215C9.91018 4.02895 9.41302 3.81236 8.83227 3.81236C8.25152 3.81236 7.75437 4.02895 7.3408 4.46215C6.92724 4.89535 6.72045 5.41611 6.72045 6.02443C6.72045 6.63275 6.92724 7.15351 7.3408 7.58671C7.75437 8.0199 8.25152 8.2365 8.83227 8.2365Z\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M18.3355 9.34254L16.2738 9.34254C15.6629 9.34254 15.1677 8.84735 15.1677 8.2365L15.1677 8.2365C15.1677 7.62565 15.6629 7.13047 16.2738 7.13047L18.3355 7.13047L20.4473 7.13047L22.509 7.13047C23.1198 7.13047 23.615 7.62565 23.615 8.2365L23.615 8.2365C23.615 8.84735 23.1198 9.34254 22.509 9.34254L20.4473 9.34254L18.3355 9.34254ZM8.83227 10.4486C7.67077 10.4486 6.67646 10.0154 5.84933 9.14898C5.0222 8.28259 4.60864 7.24107 4.60864 6.02443C4.60864 4.80779 5.0222 3.76627 5.84933 2.89988C6.67646 2.03348 7.67077 1.60028 8.83227 1.60028C9.99377 1.60028 10.9881 2.03348 11.8152 2.89988C12.6423 3.76627 13.0559 4.80779 13.0559 6.02443C13.0559 7.24107 12.6423 8.28259 11.8152 9.14898C10.9881 10.0154 9.99377 10.4486 8.83227 10.4486ZM1.49104 22.3997C0.880189 22.3997 0.385 21.9045 0.385 21.2937L0.385 16.2C0.385 15.5732 0.538987 14.9972 0.84696 14.4718C1.15493 13.9464 1.5641 13.5455 2.07445 13.269C3.16556 12.6975 4.27427 12.2689 5.40057 11.9832C6.52687 11.6975 7.67077 11.5546 8.83227 11.5546C9.99377 11.5546 11.1377 11.6975 12.264 11.9832C13.3903 12.2689 14.499 12.6975 15.5901 13.269C16.1004 13.5455 16.5096 13.9464 16.8176 14.4718C17.1256 14.9972 17.2795 15.5732 17.2795 16.2L17.2795 21.2937C17.2795 21.9045 16.7844 22.3997 16.1735 22.3997L1.49104 22.3997ZM2.49682 20.1876L15.1677 20.1876L15.1677 16.2C15.1677 15.9972 15.1193 15.8129 15.0225 15.6469C14.9257 15.481 14.7982 15.352 14.6398 15.2598C13.6895 14.7621 12.7303 14.3888 11.7624 14.14C10.7945 13.8911 9.81779 13.7667 8.83227 13.7667C7.84676 13.7667 6.87004 13.8911 5.90213 14.14C4.93421 14.3888 3.97509 14.7621 3.02477 15.2598C2.86639 15.352 2.7388 15.481 2.64201 15.6469C2.54521 15.8129 2.49682 15.9972 2.49682 16.2L2.49682 20.1876ZM8.83227 8.2365C9.41302 8.2365 9.91018 8.0199 10.3237 7.58671C10.7373 7.15351 10.9441 6.63275 10.9441 6.02443C10.9441 5.41611 10.7373 4.89535 10.3237 4.46215C9.91018 4.02895 9.41302 3.81236 8.83227 3.81236C8.25152 3.81236 7.75437 4.02895 7.3408 4.46215C6.92724 4.89535 6.72045 5.41611 6.72045 6.02443C6.72045 6.63275 6.92724 7.15351 7.3408 7.58671C7.75437 8.0199 8.25152 8.2365 8.83227 8.2365Z\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func skillAddIcon(toogle bool) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var20 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var20 == nil {
			templ_7745c5c3_Var20 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if toogle {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M17 10C13.141 10 10 13.141 10 17C10 20.859 13.141 24 17 24C20.859 24 24 20.859 24 17C24 13.141 20.859 10 17 10ZM17 22C14.243 22 12 19.757 12 17C12 14.243 14.243 12 17 12C19.757 12 22 14.243 22 17C22 19.757 19.757 22 17 22ZM20.5 17C20.5 17.553 20.053 18 19.5 18L18 18L18 19.5C18 20.053 17.553 20.5 17 20.5C16.447 20.5 16 20.053 16 19.5L16 18L14.5 18C13.947 18 13.5 17.553 13.5 17C13.5 16.447 13.947 16 14.5 16L16 16L16 14.5C16 13.947 16.447 13.5 17 13.5C17.553 13.5 18 13.947 18 14.5L18 16L19.5 16C20.053 16 20.5 16.447 20.5 17ZM9 22L4 22C2.897 22 2 21.103 2 20C2 18.897 2.897 18 4 18L7 18C7.553 18 8 17.553 8 17C8 16.447 7.553 16 7 16L6 16L6 2L15 2C16.654 2 18 3.346 18 5L18 7C18 7.553 18.447 8 19 8C19.553 8 20 7.553 20 7L20 5C20 2.243 17.757 0 15 0L5 0C2.243 0-2.66454e-15 2.243-2.66454e-15 5L-2.66454e-15 20C-2.66454e-15 22.206 1.794 24 4 24L9 24C9.553 24 10 23.553 10 23C10 22.447 9.553 22 9 22ZM2 5C2 3.698 2.839 2.598 4 2.184L4 16C3.268 16 2.591 16.212 2 16.556L2 5Z\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M17 10C13.141 10 10 13.141 10 17C10 20.859 13.141 24 17 24C20.859 24 24 20.859 24 17C24 13.141 20.859 10 17 10ZM17 22C14.243 22 12 19.757 12 17C12 14.243 14.243 12 17 12C19.757 12 22 14.243 22 17C22 19.757 19.757 22 17 22ZM20.5 17C20.5 17.553 20.053 18 19.5 18L18 18L16 18L14.5 18C13.947 18 13.5 17.553 13.5 17C13.5 16.447 13.947 16 14.5 16L16 16L18 16L19.5 16C20.053 16 20.5 16.447 20.5 17ZM9 22L4 22C2.897 22 2 21.103 2 20C2 18.897 2.897 18 4 18L7 18C7.553 18 8 17.553 8 17C8 16.447 7.553 16 7 16L6 16L6 2L15 2C16.654 2 18 3.346 18 5L18 7C18 7.553 18.447 8 19 8C19.553 8 20 7.553 20 7L20 5C20 2.243 17.757 0 15 0L5 0C2.243 0 0 2.243 0 5L0 20C0 22.206 1.794 24 4 24L9 24C9.553 24 10 23.553 10 23C10 22.447 9.553 22 9 22ZM2 5C2 3.698 2.839 2.598 4 2.184L4 16C3.268 16 2.591 16.212 2 16.556L2 5Z\"></path>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func memberIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M5.5 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0m3-5a5 5 0 1 0 0 10a5 5 0 0 0 0-10m7 0h-1v2h1a3 3 0 1 1 0 6h-1v2h1a5 5 0 0 0 0-10M0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0zm24 0a5 5 0 0 0-5-5h-1v2h1a3 3 0 0 1 3 3v2h2z\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func teamIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var22 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var22 == nil {
			templ_7745c5c3_Var22 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M12 11a5 5 0 0 1 5 5v6h-2v-6a3 3 0 0 0-2.824-2.995L12 13a3 3 0 0 0-2.995 2.824L9 16v6H7v-6a5 5 0 0 1 5-5m-6.5 3q.42.001.81.094a6 6 0 0 0-.301 1.575L6 16v.086a1.5 1.5 0 0 0-.356-.08L5.5 16a1.5 1.5 0 0 0-1.493 1.355L4 17.5V22H2v-4.5A3.5 3.5 0 0 1 5.5 14m13 0a3.5 3.5 0 0 1 3.5 3.5V22h-2v-4.5a1.5 1.5 0 0 0-1.355-1.493L18.5 16q-.264.001-.5.085V16c0-.666-.108-1.306-.308-1.904c.258-.063.53-.096.808-.096m-13-6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m13 0a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5m-13 2a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m13 0a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1M12 2a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 2a2 2 0 1 0 0 4a2 2 0 0 0 0-4\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func skillIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var23 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var23 == nil {
			templ_7745c5c3_Var23 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"M7.616 21q-1.085 0-1.85-.766Q5 19.47 5 18.385V6q0-1.258.871-2.129T8 3h11v13.77q-.663 0-1.14.475t-.475 1.14t.476 1.139T19 20v1zM6 16.363q.33-.29.735-.442q.403-.152.88-.152h.77V4H8q-.817 0-1.409.591Q6 5.183 6 6zm3.385-.594H18V4H9.385zM6 16.363V4.385zM7.616 20h9.363q-.285-.33-.44-.732q-.155-.4-.155-.884q0-.457.152-.87t.443-.745H7.616q-.689 0-1.152.476T6 18.385q0 .688.464 1.151T7.616 20\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func settingsIcon() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var24 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var24 == nil {
			templ_7745c5c3_Var24 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<path d=\"m9.25 22l-.4-3.2q-.325-.125-.612-.3t-.563-.375L4.7 19.375l-2.75-4.75l2.575-1.95Q4.5 12.5 4.5 12.338v-.675q0-.163.025-.338L1.95 9.375l2.75-4.75l2.975 1.25q.275-.2.575-.375t.6-.3l.4-3.2h5.5l.4 3.2q.325.125.613.3t.562.375l2.975-1.25l2.75 4.75l-2.575 1.95q.025.175.025.338v.674q0 .163-.05.338l2.575 1.95l-2.75 4.75l-2.95-1.25q-.275.2-.575.375t-.6.3l-.4 3.2zM11 20h1.975l.35-2.65q.775-.2 1.438-.587t1.212-.938l2.475 1.025l.975-1.7l-2.15-1.625q.125-.35.175-.737T17.5 12t-.05-.787t-.175-.738l2.15-1.625l-.975-1.7l-2.475 1.05q-.55-.575-1.212-.962t-1.438-.588L13 4h-1.975l-.35 2.65q-.775.2-1.437.588t-1.213.937L5.55 7.15l-.975 1.7l2.15 1.6q-.125.375-.175.75t-.05.8q0 .4.05.775t.175.75l-2.15 1.625l.975 1.7l2.475-1.05q.55.575 1.213.963t1.437.587zm1.05-4.5q1.45 0 2.475-1.025T15.55 12t-1.025-2.475T12.05 8.5q-1.475 0-2.488 1.025T8.55 12t1.013 2.475T12.05 15.5M12 12\"></path>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
