// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.857
package timepicker

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/axzilla/templui/component/button"
	"github.com/axzilla/templui/component/selectbox"
	"github.com/axzilla/templui/icon"
	"github.com/axzilla/templui/util"
	"strconv"
	"time"
)

type Props struct {
	ID          string
	Class       string
	Attributes  templ.Attributes
	Name        string
	Value       time.Time
	Use12Hours  bool
	AMLabel     string
	PMLabel     string
	Placeholder string
	Required    bool
	Disabled    bool
	HasError    bool
}

type SelectOption struct {
	Label string
	Value string
}

func TimePicker(props ...Props) templ.Component {
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
		var p Props
		if len(props) > 0 {
			p = props[0]
		}
		if p.ID == "" {
			p.ID = util.RandomID()
		}
		if p.Placeholder == "" {
			p.Placeholder = "Select time"
		}
		if p.AMLabel == "" {
			p.AMLabel = "AM"
		}
		if p.PMLabel == "" {
			p.PMLabel = "PM"
		}
		var templ_7745c5c3_Var2 = []any{util.TwMerge("relative", p.Class)}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Value != (time.Time{}) {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, " data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format("15:04"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 53, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, " data-use12hours=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%t", p.Use12Hours))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 55, Col: 51}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" x-data=\"timePicker\" data-input-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(p.ID)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 57, Col: 22}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\"><div class=\"relative\"><input type=\"hidden\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(p.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 62, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Required {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, " required")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, ">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			var templ_7745c5c3_Var9 = []any{
				util.TwMerge(
					"timepicker-display truncate",
					util.If(p.Value == (time.Time{}), "text-muted-foreground"),
				),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var9...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<span class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var9).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if p.Value != (time.Time{}) {
				if p.Use12Hours {
					var templ_7745c5c3_Var11 string
					templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format("03:04 PM"))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 91, Col: 35}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					var templ_7745c5c3_Var12 string
					templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(p.Value.Format("15:04"))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 93, Col: 32}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				var templ_7745c5c3_Var13 string
				templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(p.Placeholder)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 96, Col: 21}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = icon.Clock(icon.Props{
				Class: "ml-2 h-4 w-4 text-muted-foreground",
			}).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = button.Button(button.Props{
			ID:      p.ID,
			Type:    "button",
			Variant: button.VariantOutline,
			Class: util.TwMerge(
				"w-full flex justify-between items-center",
				util.If(p.HasError, "border-destructive ring-destructive"),
			),
			Disabled: p.Disabled,
			Attributes: util.MergeAttributes(
				templ.Attributes{
					"@click": "toggleTimePicker",
				},
				p.Attributes,
			),
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "</div><div x-show=\"open\" x-ref=\"timePickerPopup\" @click.away=\"closeTimePicker\" x-transition.opacity class=\"absolute left-0 z-50 w-72 p-4 rounded-lg border bg-popover shadow-md\" :class=\"positionClass\" @resize.window=\"updatePosition\" style=\"display: none;\"><div class=\"grid grid-cols-2 gap-2 flex-1\"><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var14 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Var15 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
				templ_7745c5c3_Var16 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
					if p.Value != (time.Time{}) {
						if p.Use12Hours {
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, " ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							hour := p.Value.Hour() % 12

							if hour == 0 {
								hour = 12
							}
							var templ_7745c5c3_Var17 string
							templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", hour))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 135, Col: 37}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						} else {
							var templ_7745c5c3_Var18 string
							templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", p.Value.Hour()))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 137, Col: 47}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
					}
					return nil
				})
				templ_7745c5c3_Err = selectbox.Value(selectbox.ValueProps{
					Placeholder: "Hour",
				}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var16), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = selectbox.Trigger(selectbox.TriggerProps{
				Name: "hour",
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var15), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var19 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
				templ_7745c5c3_Var20 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
					for _, option := range hourOptions(p.Use12Hours) {
						templ_7745c5c3_Var21 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
							templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
							var templ_7745c5c3_Var22 string
							templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(option.Label)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 153, Col: 24}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return nil
						})
						templ_7745c5c3_Err = selectbox.Item(selectbox.ItemProps{
							Value: option.Value,
							Selected: p.Value != (time.Time{}) &&
								((p.Use12Hours &&
									((p.Value.Hour()%12 == 0 && option.Value == "12") ||
										(p.Value.Hour()%12 == parseInt(option.Value)))) ||
									(!p.Use12Hours && p.Value.Hour() == parseInt(option.Value))),
						}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var21), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = selectbox.Group().Render(templ.WithChildren(ctx, templ_7745c5c3_Var20), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = selectbox.Content().Render(templ.WithChildren(ctx, templ_7745c5c3_Var19), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = selectbox.SelectBox(selectbox.Props{
			ID:    "hour-select-" + p.ID,
			Class: "w-full",
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var14), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "</div><div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var23 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Var24 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
				templ_7745c5c3_Var25 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
					if p.Value != (time.Time{}) {
						var templ_7745c5c3_Var26 string
						templ_7745c5c3_Var26, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%02d", p.Value.Minute()))
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 172, Col: 48}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var26))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = selectbox.Value(selectbox.ValueProps{
					Placeholder: "Minute",
				}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var25), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = selectbox.Trigger(selectbox.TriggerProps{
				Name: "minute",
			}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var24), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Var27 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
				templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
				templ_7745c5c3_Var28 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
					templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
					for _, option := range minuteOptions() {
						templ_7745c5c3_Var29 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
							templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
							var templ_7745c5c3_Var30 string
							templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(option.Label)
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 183, Col: 24}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							return nil
						})
						templ_7745c5c3_Err = selectbox.Item(selectbox.ItemProps{
							Value:    option.Value,
							Selected: p.Value != (time.Time{}) && p.Value.Minute() == parseInt(option.Value),
						}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var29), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					return nil
				})
				templ_7745c5c3_Err = selectbox.Group().Render(templ.WithChildren(ctx, templ_7745c5c3_Var28), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				return nil
			})
			templ_7745c5c3_Err = selectbox.Content().Render(templ.WithChildren(ctx, templ_7745c5c3_Var27), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = selectbox.SelectBox(selectbox.Props{
			ID:    "minute-select-" + p.ID,
			Class: "w-full",
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var23), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</div></div><div class=\"flex justify-between mt-4 gap-2\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if p.Use12Hours {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, "<div class=\"flex justify-center gap-2\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var31 = []any{
				util.TwMerge(
					"px-3 py-1 text-sm rounded-md hover:bg-muted",
					util.If(p.Value != (time.Time{}) && p.Value.Hour() < 12, "bg-primary text-primary-foreground"),
				),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var31...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, "<button type=\"button\" data-period=\"AM\" @click=\"selectPeriod\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var32 string
			templ_7745c5c3_Var32, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var31).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var32))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var33 string
			templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs(p.AMLabel)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 205, Col: 18}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, "</button> ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var34 = []any{
				util.TwMerge(
					"px-3 py-1 text-sm rounded-md hover:bg-muted",
					util.If(p.Value != (time.Time{}) && p.Value.Hour() >= 12, "bg-primary text-primary-foreground"),
				),
			}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var34...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "<button type=\"button\" data-period=\"PM\" @click=\"selectPeriod\" class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var35 string
			templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var34).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var36 string
			templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(p.PMLabel)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 218, Col: 18}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, "</button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, "<div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Var37 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "Done")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = button.Button(button.Props{
			Type:    "button",
			Variant: button.VariantSecondary,
			Attributes: templ.Attributes{
				"@click": "closeTimePicker",
			},
		}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var37), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func parseInt(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return val
}

func hourOptions(use12Hours bool) []SelectOption {
	options := make([]SelectOption, 0)
	if use12Hours {
		options = append(options, SelectOption{
			Label: "12",
			Value: "12",
		})
		for i := 1; i <= 11; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	} else {
		for i := 0; i <= 23; i++ {
			options = append(options, SelectOption{
				Label: fmt.Sprintf("%02d", i),
				Value: fmt.Sprintf("%d", i),
			})
		}
	}
	return options
}

func minuteOptions() []SelectOption {
	options := make([]SelectOption, 60)
	for i := 0; i < 60; i++ {
		options[i] = SelectOption{
			Label: fmt.Sprintf("%02d", i),
			Value: fmt.Sprintf("%d", i),
		}
	}
	return options
}

func Script() templ.Component {
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
		templ_7745c5c3_Var38 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var38 == nil {
			templ_7745c5c3_Var38 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		handle := templ.NewOnceHandle()
		templ_7745c5c3_Var39 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "<script defer nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var40 string
			templ_7745c5c3_Var40, templ_7745c5c3_Err = templ.JoinStringErrs(templ.GetNonce(ctx))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `component/timepicker/time_picker.templ`, Line: 284, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var40))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "\">\n\t\t\tdocument.addEventListener('alpine:init', () => {\n\t\t\t\tAlpine.data('timePicker', () => ({\n\t\t\t\t\topen: false,\n\t\t\t\t\tformValue: null,\n\t\t\t\t\tdisplayValue: null,\n\t\t\t\t\tselectedHour: 0,\n\t\t\t\t\tselectedMinute: 0,\n\t\t\t\t\tisPM: false,\n\t\t\t\t\thours: [],\n\t\t\t\t\tminutes: [],\n\t\t\t\t\tuse12Hours: false,\n\t\t\t\t\tposition: 'bottom',\n\n\t\t\t\t\tinit() {\n\t\t\t\t\t\tthis.use12Hours = this.$el.dataset.use12hours === 'true';\n\t\t\t\t\t\tif (this.$el.dataset?.value) {\n\t\t\t\t\t\t\tconst [hours, minutes] = this.$el.dataset.value.split(':').map(Number);\n\t\t\t\t\t\t\tthis.selectedMinute = minutes;\n\t\t\t\t\t\t\t\n\t\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\t\tthis.isPM = hours >= 12;\n\t\t\t\t\t\t\t\tthis.selectedHour = hours > 12 ? hours - 12 : (hours === 0 ? 12 : hours);\n\t\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\t\tthis.selectedHour = hours;\n\t\t\t\t\t\t\t}\n\t\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t\t}\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Set up listeners for the selects\n\t\t\t\t\t\tthis.$nextTick(() => {\n                            // We handle selection by listening to click events on the select items\n                            // This is handled by the SelectScript but we need to sync the values\n                            \n                            const hourSelectId = 'hour-select-' + this.$el.dataset.inputId;\n                            const minuteSelectId = 'minute-select-' + this.$el.dataset.inputId;\n                            \n                            // Listen for changes on the hidden inputs to sync our values\n                            const hourInput = document.querySelector(`[data-select-id=\"${hourSelectId}\"] input[type=\"hidden\"]`);\n                            const minuteInput = document.querySelector(`[data-select-id=\"${minuteSelectId}\"] input[type=\"hidden\"]`);\n                            \n                            if (hourInput) {\n                                hourInput.addEventListener('change', () => {\n                                    this.selectedHour = parseInt(hourInput.value);\n                                    this.updateValues();\n                                });\n                            }\n                            \n                            if (minuteInput) {\n                                minuteInput.addEventListener('change', () => {\n                                    this.selectedMinute = parseInt(minuteInput.value);\n                                    this.updateValues();\n                                });\n                            }\n                        });\n\t\t\t\t\t},\n\n\t\t\t\t\ttoggleTimePicker() {\n\t\t\t\t\t\tthis.open = !this.open;\n\t\t\t\t\t\tif (this.open) {\n\t\t\t\t\t\t\tthis.$nextTick(() => this.updatePosition());\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tupdatePosition() {\n\t\t\t\t\t\tconst inputId = this.$root.dataset.inputId;\n\t\t\t\t\t\tconst trigger = document.getElementById(inputId);\n\t\t\t\t\t\tconst popup = this.$refs.timePickerPopup;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (!trigger || !popup) return;\n\t\t\t\t\t\t\n\t\t\t\t\t\tconst rect = trigger.getBoundingClientRect();\n\t\t\t\t\t\tconst popupRect = popup.getBoundingClientRect();\n\t\t\t\t\t\tconst viewportHeight = window.innerHeight;\n\t\t\t\t\t\t\n\t\t\t\t\t\tif (rect.bottom + popupRect.height > viewportHeight && rect.top > popupRect.height) {\n\t\t\t\t\t\t\tthis.position = 'top';\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.position = 'bottom';\n\t\t\t\t\t\t}\n\t\t\t\t\t},\n\n\t\t\t\t\tpositionClass() {\n\t\t\t\t\t\treturn this.position === 'bottom' ? 'top-full mt-1' : 'bottom-full mb-1';\n\t\t\t\t\t},\n\n\t\t\t\t\tcloseTimePicker() {\n\t\t\t\t\t\tthis.open = false;\n\t\t\t\t\t},\n\n\t\t\t\t\tupdateValues() {\n\t\t\t\t\t\tlet hour24 = this.selectedHour;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tif (this.isPM && hour24 !== 12) hour24 += 12;\n\t\t\t\t\t\t\tif (!this.isPM && hour24 === 12) hour24 = 0;\n\t\t\t\t\t\t}\n\t\t\t\t\t\tthis.formValue = `${String(hour24).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')}`;\n\t\t\t\t\t\tif (this.use12Hours) {\n\t\t\t\t\t\t\tthis.displayValue = `${String(this.selectedHour).padStart(2, '0')}:${String(this.selectedMinute).padStart(2, '0')} ${this.isPM ? 'PM' : 'AM'}`;\n\t\t\t\t\t\t} else {\n\t\t\t\t\t\t\tthis.displayValue = this.formValue;\n\t\t\t\t\t\t}\n                        \n                        // Update the display element with the new value\n                        const displayEl = this.$el.querySelector('.timepicker-display');\n                        if (displayEl) {\n                            displayEl.textContent = this.displayValue;\n                            displayEl.classList.remove('text-muted-foreground');\n                        }\n                        \n                        // Update the hidden form input\n                        const hiddenInput = this.$el.querySelector('input[type=\"hidden\"][name]');\n                        if (hiddenInput) {\n                            hiddenInput.value = this.formValue;\n                        }\n\t\t\t\t\t},\n\n\t\t\t\t\tselectPeriod(event) {\n\t\t\t\t\t\tconst periodButton = event.currentTarget;\n\t\t\t\t\t\tconst period = periodButton.getAttribute('data-period');\n\t\t\t\t\t\tthis.isPM = period === 'PM';\n\t\t\t\t\t\t\n\t\t\t\t\t\t// Update UI for AM/PM buttons\n\t\t\t\t\t\tconst periodButtons = this.$el.querySelectorAll('[data-period]');\n\t\t\t\t\t\tperiodButtons.forEach(btn => {\n\t\t\t\t\t\t\tbtn.classList.remove('bg-primary', 'text-primary-foreground');\n\t\t\t\t\t\t});\n\t\t\t\t\t\tperiodButton.classList.add('bg-primary', 'text-primary-foreground');\n\t\t\t\t\t\t\n\t\t\t\t\t\tthis.updateValues();\n\t\t\t\t\t}\n\t\t\t\t}));\n\t\t\t});\n\t\t</script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = handle.Once().Render(templ.WithChildren(ctx, templ_7745c5c3_Var39), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
