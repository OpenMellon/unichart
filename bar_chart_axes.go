package unichart

import (
	"github.com/OpenMellon/unichart/dataset/sequence"
	"github.com/OpenMellon/unichart/render"
)

func (bc *BarChart) drawXAxis(r render.Renderer, canvasBox render.Box) {
	if !bc.XAxis.Hidden {
		axisStyle := bc.XAxis.InheritFrom(bc.styleDefaultsAxes())
		axisStyle.WriteToRenderer(r)

		width, spacing, _ := bc.calculateScaledTotalSize(canvasBox)

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.Stroke()

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Left, canvasBox.Bottom+defaultVerticalTickHeight)
		r.Stroke()

		cursor := canvasBox.Left
		for index, bar := range bc.Bars {
			barLabelBox := render.Box{
				Top:    canvasBox.Bottom + defaultXAxisMargin,
				Left:   cursor,
				Right:  cursor + width + spacing,
				Bottom: bc.Height(),
			}

			if len(bar.Label) > 0 {
				render.Text.DrawWithin(r, bar.Label, barLabelBox, axisStyle)
			}

			axisStyle.WriteToRenderer(r)
			if index < len(bc.Bars)-1 {
				r.MoveTo(barLabelBox.Right, canvasBox.Bottom)
				r.LineTo(barLabelBox.Right, canvasBox.Bottom+defaultVerticalTickHeight)
				r.Stroke()
			}
			cursor += width + spacing
		}
	}
}

func (bc *BarChart) drawHorizontalXAxis(r render.Renderer, canvasBox render.Box, yr sequence.Range, ticks []Tick) {
	if !bc.YAxis.Style.Hidden {
		axisStyle := bc.YAxis.Style.InheritFrom(bc.styleDefaultsAxes())
		axisStyle.WriteToRenderer(r)

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.Stroke()

		r.MoveTo(canvasBox.Left, canvasBox.Bottom)
		r.LineTo(canvasBox.Left-defaultHorizontalTickWidth, canvasBox.Bottom)
		r.Stroke()

		var tx int
		var tb render.Box
		for _, t := range ticks {
			tx = canvasBox.Left + yr.Translate(t.Value)

			axisStyle.GetStrokeOptions().WriteToRenderer(r)
			r.MoveTo(tx, canvasBox.Bottom)
			r.LineTo(tx, canvasBox.Bottom+defaultHorizontalTickWidth)
			r.Stroke()

			axisStyle.GetTextOptions().WriteToRenderer(r)
			tb = r.MeasureText(t.Label)
			render.Text.Draw(r, t.Label, tx-(tb.Width()>>1), canvasBox.Bottom+defaultXAxisMargin+5, axisStyle)
		}
	}
}

func (bc *BarChart) drawYAxis(r render.Renderer, canvasBox render.Box, yr sequence.Range, ticks []Tick) {
	if !bc.YAxis.Style.Hidden {
		axisStyle := bc.YAxis.Style.InheritFrom(bc.styleDefaultsAxes())
		axisStyle.WriteToRenderer(r)

		r.MoveTo(canvasBox.Right, canvasBox.Top)
		r.LineTo(canvasBox.Right, canvasBox.Bottom)
		r.Stroke()

		r.MoveTo(canvasBox.Right, canvasBox.Bottom)
		r.LineTo(canvasBox.Right+defaultHorizontalTickWidth, canvasBox.Bottom)
		r.Stroke()

		var ty int
		var tb render.Box
		for _, t := range ticks {
			ty = canvasBox.Bottom - yr.Translate(t.Value)

			axisStyle.GetStrokeOptions().WriteToRenderer(r)
			r.MoveTo(canvasBox.Right, ty)
			r.LineTo(canvasBox.Right+defaultHorizontalTickWidth, ty)
			r.Stroke()

			axisStyle.GetTextOptions().WriteToRenderer(r)
			tb = r.MeasureText(t.Label)
			render.Text.Draw(r, t.Label, canvasBox.Right+defaultYAxisMargin+5, ty+(tb.Height()>>1), axisStyle)
		}

	}
}

func (bc *BarChart) drawHorizontalYAxis(r render.Renderer, canvasBox render.Box) {
	if !bc.XAxis.Hidden {
		defaultStyle := bc.styleDefaultsAxes()
		defaultStyle.TextHorizontalAlign = render.TextHorizontalAlignRight
		defaultStyle.TextVerticalAlign = render.TextVerticalAlignMiddle

		axisStyle := bc.XAxis.InheritFrom(defaultStyle)

		axisStyle.WriteToRenderer(r)

		width, spacing, _ := bc.calculateScaledTotalSize(canvasBox)

		cursor := canvasBox.Bottom - len(bc.Bars)*(width+spacing)

		r.MoveTo(canvasBox.Left, cursor)
		r.LineTo(canvasBox.Left, canvasBox.Bottom)
		r.Stroke()

		for index, bar := range bc.Bars {
			tb := r.MeasureText(bar.Label)

			barLabelBox := render.Box{
				Top:    cursor + spacing,
				Left:   canvasBox.Left - tb.Width() - (2 * defaultYAxisMargin),
				Right:  canvasBox.Left - defaultYAxisMargin,
				Bottom: cursor + width,
			}

			if len(bar.Label) > 0 {
				render.Text.DrawWithin(r, bar.Label, barLabelBox, axisStyle)
			}

			axisStyle.WriteToRenderer(r)
			if index < len(bc.Bars) {
				r.MoveTo(canvasBox.Left, cursor)
				r.LineTo(canvasBox.Left-defaultHorizontalTickWidth, cursor)
				r.Stroke()
			}
			cursor += width + spacing
		}
	}
}
