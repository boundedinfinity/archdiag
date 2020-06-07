package archdiag

import (
	"fmt"
	"strings"

	svg "github.com/ajstarks/svgo"
)

const (
	root_id = "root"
)

type box struct {
	id         string
	label      string
	parent     box
	bmap       map[string]*box
	blist      []*box
	direction  Direction
	dimensions dimensions
	Center     point
	style      boxStyle
}

func (t *box) Append(b *box) error {
	if t.bmap == nil {
		t.bmap = make(map[string]*box)
	}

	if t.blist == nil {
		t.blist = make([]*box, 0)
	}

	if _, ok := t.bmap[b.id]; ok {
		return fmt.Errorf("id %v exists", b.id)
	}

	t.bmap[b.id] = b
	t.blist = append(t.blist, b)

	return nil
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

func (t *box) calcDimensions() {
	if t.label == "" {
		t.dimensions.H = 2 * t.style.Padding
		t.dimensions.W = 2 * t.style.Padding
	} else {
		t.dimensions.H = t.style.LabelSize + 2*t.style.Padding
		t.dimensions.W = len(t.label)*t.style.LabelSize/2 + 2*t.style.Padding
	}

	if t.bmap != nil {
		for _, cn := range t.bmap {
			cn.calcDimensions()

			switch t.direction {
			case DirectionHorizontal:
				t.dimensions.W += cn.dimensions.W + 2*t.style.Padding
			case DirectionVertical:
				t.dimensions.H += cn.dimensions.H + 2*t.style.Padding
			}
		}
	}
}

func (t box) Draw(canvas *svg.SVG) {
	t.calcDimensions()

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

	if t.bmap != nil {
		for _, sc := range t.bmap {
			sc.Draw(canvas)
		}
	}
}

func (t box) TopLeft() point {
	return point{
		X: t.Center.X - t.dimensions.W/2,
		Y: t.Center.Y - t.dimensions.H/2,
	}
}

func (t box) TopCenter() point {
	return point{
		X: t.Center.X,
		Y: t.Center.Y - t.dimensions.H/2,
	}
}

func (t box) TopRight() point {
	return point{
		X: t.Center.X + t.dimensions.W/2,
		Y: t.Center.Y - t.dimensions.H/2,
	}
}

func (t box) RightCenter() point {
	return point{
		X: t.Center.X + t.dimensions.W/2,
		Y: t.Center.Y,
	}
}

func (t box) LeftCenter() point {
	return point{
		X: t.Center.X - t.dimensions.W/2,
		Y: t.Center.Y,
	}
}

func (t box) BottomLeft() point {
	return point{
		X: t.Center.X - t.dimensions.W/2,
		Y: t.Center.Y + t.dimensions.H/2,
	}
}

func (t box) BottomCenter() point {
	return point{
		X: t.Center.X,
		Y: t.Center.Y + t.dimensions.H/2,
	}
}

func (t box) BottomRight() point {
	return point{
		X: t.Center.X + t.dimensions.W/2,
		Y: t.Center.Y + t.dimensions.H/2,
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
