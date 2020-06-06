package archdiag

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
)

type box struct {
	id         string
	label      string
	children   map[string]box
	direction  Direction
	dimensions dimensions
	topLeft    point
	style      boxStyle
}

var DEFAULT_BOX_STYLE = boxStyle{
	Padding:         5,
	BorderColor:     "#7B8894",
	BorderSize:      1,
	BorderRadius:    5,
	BackgroundColor: "#E5F5FD",
	LabelSize:       15,
	LabelColor:      "#696969",
}

type boxStyle struct {
	Padding         int    `json:"padding" yaml:"padding"`
	BorderColor     string `json:"borderColor" yaml:"borderColor"`
	BorderSize      int    `json:"borderSize" yaml:"borderSize"`
	BorderRadius    int    `json:"borderRadius" yaml:"borderRadius"`
	BackgroundColor string `json:"backgroundColor" yaml:"backgroundColor"`
	LabelSize       int    `json:"labelSize" yaml:"labelSize"`
	LabelColor      string `json:"labelColor" yaml:"labelColor"`
}

func calcBoxStyle(l1 boxStyle) boxStyle {
	d := DEFAULT_BOX_STYLE
	return boxStyle{
		BorderRadius:    getInt(d.BorderRadius, l1.BorderRadius),
		BackgroundColor: getStr(d.BackgroundColor, l1.BackgroundColor),
		BorderColor:     getStr(d.BorderColor, l1.BorderColor),
		BorderSize:      getInt(d.BorderSize, l1.BorderSize),
		LabelColor:      getStr(d.LabelColor, l1.LabelColor),
		LabelSize:       getInt(d.LabelSize, l1.LabelSize),
		Padding:         getInt(d.Padding, l1.Padding),
	}
}

func (t *box) CalcDimensions() {
	t.dimensions.H = t.style.LabelSize + 2*t.style.Padding
	t.dimensions.W = len(t.label)*t.style.LabelSize/2 + 2*t.style.Padding

	if t.children != nil {
		for _, cn := range t.children {
			cn.CalcDimensions()

			switch t.direction {
			case DirectionHorizontal:
				t.dimensions.W += cn.dimensions.W + 2*t.style.Padding
			case DirectionVertical:
				t.dimensions.H += cn.dimensions.H + 2*t.style.Padding
			}
		}
	}
}

func (t *box) CalcTopLeft(topLeft point) {
	t.topLeft.X = getInt(t.style.Padding, topLeft.X)
	t.topLeft.Y = getInt(t.style.Padding, topLeft.Y)

	if t.children != nil {
		lastTopLeft := point{
			X: t.topLeft.X,
			Y: t.topLeft.Y,
		}

		for _, cn := range t.children {
			switch t.direction {
			case DirectionHorizontal:
				lastTopLeft.X = t.TopLeft().X + t.dimensions.W + t.style.Padding
				lastTopLeft.Y = t.topLeft.Y
			case DirectionVertical:
				lastTopLeft.X = t.TopLeft().X
				lastTopLeft.Y = t.TopLeft().Y + t.dimensions.H + t.style.Padding
			}

			cn.CalcTopLeft(lastTopLeft)
		}
	}
}

func (t box) Draw(canvas *svg.SVG) {
	canvas.Roundrect(
		t.TopLeft().X, t.TopLeft().Y,
		t.dimensions.W, t.dimensions.H,
		t.style.BorderRadius, t.style.BorderRadius,
		t.rectAttrs(),
	)

	textBottomLeft := point{
		X: t.TopLeft().X + t.style.Padding,
		Y: t.TopLeft().Y + t.style.Padding + t.style.LabelSize,
	}

	canvas.Text(textBottomLeft.X, textBottomLeft.Y, t.label, t.textAttrs())

	if t.children != nil {
		for _, sc := range t.children {
			sc.Draw(canvas)
		}
	}
}

func (t box) TopLeft() point {
	return point{
		X: t.topLeft.X,
		Y: t.topLeft.Y,
	}
}

func (t box) TopCenter() point {
	return point{
		X: t.topLeft.X + t.dimensions.W/2,
		Y: t.topLeft.Y,
	}
}

func (t box) TopRight() point {
	return point{
		X: t.topLeft.X + t.dimensions.W,
		Y: t.topLeft.Y,
	}
}

func (t box) RightCenter() point {
	return point{
		X: t.topLeft.X + t.dimensions.W,
		Y: t.topLeft.Y * t.dimensions.H / 2,
	}
}

func (t box) LeftCenter() point {
	return point{
		X: t.topLeft.X,
		Y: t.topLeft.Y * t.dimensions.H / 2,
	}
}

func (t box) BottomLeft() point {
	return point{
		X: t.topLeft.X,
		Y: t.topLeft.Y * t.dimensions.H,
	}
}

func (t box) BottomCenter() point {
	return point{
		X: t.topLeft.X * t.dimensions.W / 2,
		Y: t.topLeft.Y * t.dimensions.H,
	}
}

func (t box) BottomRight() point {
	return point{
		X: t.topLeft.X * t.dimensions.W,
		Y: t.topLeft.Y * t.dimensions.H,
	}
}

func (t box) textAttrs() string {
	attrs := make([]string, 0)

	attrs = append(attrs, fmt.Sprintf("text-anchor:%v", TextAnchorStart))

	if t.style.LabelColor != "" {
		attrs = append(attrs, fmt.Sprintf("fill:%v", t.style.LabelColor))
	}

	if t.style.LabelSize != 0 {
		attrs = append(attrs, fmt.Sprintf("font-size:%vpx", t.style.LabelSize))
	}

	return strings.Join(attrs, ";")
}

func (t box) rectAttrs() string {
	attrs := make([]string, 0)

	if t.style.BackgroundColor != "" {
		attrs = append(attrs, fmt.Sprintf("fill:%v", t.style.BackgroundColor))
	}

	if t.style.BorderColor != "" {
		attrs = append(attrs, fmt.Sprintf("stroke:%v", t.style.BorderColor))
	}

	if t.style.BorderSize != 0 {
		attrs = append(attrs, fmt.Sprintf("stroke-width:%vpx", t.style.BorderSize))
	}

	return strings.Join(attrs, ";")
}
