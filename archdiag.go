package archdiag

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

type ArchDiag struct {
	G       Graph
	p       point
	canvas  *svg.SVG
	drawers map[string]box
}

func (t *ArchDiag) Generate(w io.Writer) error {
	if err := t.Process(); err != nil {
		return err
	}

	if err := t.Draw(w); err != nil {
		return err
	}

	return nil
}

func (t *ArchDiag) Process() error {
	t.drawers = make(map[string]box)

	for id, n := range t.G.Nodes {
		dc := box{
			id:    id,
			label: getStr(id, n.Label),
			style: calcBoxStyle(n.boxStyle),
		}

		dc.CalcDimensions()
		dc.CalcTopLeft(t.p)
		t.drawers[id] = dc
	}

	return nil
}

func (t *ArchDiag) Draw(w io.Writer) error {
	attrs := calcGraphAttributes(t.G.GraphAttributes)
	t.canvas = svg.New(w)
	t.canvas.Start(attrs.Height, attrs.Width)

	for _, dc := range t.drawers {
		dc.Draw(t.canvas)
	}

	t.canvas.End()

	return nil
}
