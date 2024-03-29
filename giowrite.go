package main

// A simple Gio program. See https://gioui.org for more information.

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"

	//"gioui.org/text/opentype"
	"gioui.org/widget/material"
	//"golang.org/x/image/font/gofont/goitalic"
	//"golang.org/x/image/font/gofont/goregular"

	//	"gioui.org/text/shape"
	"gioui.org/unit"
)

func main() {
	go func() {
		s, sep := "", " "
		for _, arg := range os.Args[1:] {
			s += sep + arg
		}
		w := app.NewWindow()
		if err := loop(w, s); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window, s string) error {
	gtx := &layout.Context{
		Queue: w.Queue(),
	}
	/*
		regular, err := sfnt.Parse(goregular.TTF)
		if err != nil {
			panic("failed to load font")
		}
		family := &shape.Family{
			Regular: regular,
		}
	*/
	/*
		shaper := new(text.Shaper)
		shaper.Register(text.Font{}, opentype.Must(
			opentype.Parse(goregular.TTF),
		))
		shaper.Register(text.Font{Style: text.Italic}, opentype.Must(
			opentype.Parse(goitalic.TTF),
		))
		th := material.NewTheme(shaper)
	*/
	th := material.NewTheme()
	//var cfg app.Config
	//var faces shape.Faces
	maroon := color.RGBA{127, 0, 0, 255}
	//face := faces.For(regular, unit.Sp(72))
	message := "Hello, Gio"
	if s != "" {
		message = s
	}

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			//cfg = e.Config
			//faces.Reset(&cfg)
			//cs := layout.RigidConstraints(e.Size)
			gtx.Reset(e.Config, e.Size)
			gtx.Ops.Reset()
			var material op.MacroOp
			material.Record(gtx.Ops)
			paint.ColorOp{Color: maroon}.Add(gtx.Ops)
			material.Stop()

			gtx.Constraints.Height.Min = 0
			//	text.Label{Material: material, Size: unit.Sp(72), Alignment: text.Middle, Text: message}.Layout(gtx, family)
			lbl := th.H1(message)
			lbl.Color = maroon
			lbl.Alignment = text.Middle
			lbl.Layout(gtx)
			dims := gtx.Dimensions
			//log.Println(dims)
			op.TransformOp{}.Offset(f32.Point{Y: float32(dims.Size.Y)}).Add(gtx.Ops)
			//message += " 2"
			//text.Label{Material: material, Size: unit.Sp(72), Alignment: text.Middle, Text: message}.Layout(gtx, family)
			//material.Add(gtx.Ops)
			lbl = th.Label(unit.Sp(72), message)
			lbl.Color = maroon
			lbl.Alignment = text.Middle
			lbl.Layout(gtx)
			paint.ColorOp{Color: maroon}.Add(gtx.Ops)
			for i := 1; overlap(100*i, e.Size.X-(100*i)); i++ {
				paint.PaintOp{Rect: f32.Rectangle{
					Min: f32.Point{X: float32(100 * i), Y: float32(100 * i)},
					Max: f32.Point{X: float32(100*i) + 100, Y: float32(100*i + 100)},
				}}.Add(gtx.Ops)
				paint.PaintOp{Rect: f32.Rectangle{
					Min: f32.Point{X: float32(e.Size.X - (100*i + 100)), Y: float32(100 * i)},
					Max: f32.Point{X: float32(e.Size.X - 100*i), Y: float32(100*i) + 100},
				}}.Add(gtx.Ops)
			}
			//w.Update(gtx.Ops)
			e.Frame(gtx.Ops)
		}
	}
}

func overlap(left, right int) bool {
	log.Println(left, right)
	if left < right {
		return true
	}
	return false
}
