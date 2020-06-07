package archdiag

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

type ArchDiag struct {
	G      Graph
	canvas *svg.SVG
	root   *box
}

func (t *ArchDiag) Generate(w io.Writer) error {
	if err := t.process(); err != nil {
		return err
	}

	if err := t.Draw(w); err != nil {
		return err
	}

	return nil
}

func (t *ArchDiag) process() error {
	t.root = &box{
		id:        root_id,
		direction: t.G.Direction,
		dimensions: dimensions{
			W: t.G.Width,
			H: t.G.Height,
		},
		Center: point{
			X: t.G.Width / 2,
			Y: t.G.Height / 2,
		},
	}

	for id, n := range t.G.Nodes {
		if err := t.processNode(t.root, id, n); err != nil {
			return err
		}
	}

	return nil
}

func (t *ArchDiag) processNode(p *box, id string, n Node) error {
	b := &box{
		id:     id,
		parent: p,
		label:  getStr(id, n.Label),
		style:  calcBoxStyle(n.boxStyle),
	}

	p.Append(b)

	if n.Nodes != nil {
		for cid, cn := range n.Nodes {
			if err := t.processNode(b, cid, cn); err != nil {
				return nil
			}
		}
	}

	return nil
}

func (t *ArchDiag) Draw(w io.Writer) error {
	attrs := calcGraphAttributes(t.G.GraphAttributes)
	t.canvas = svg.New(w)
	t.canvas.Start(attrs.Height, attrs.Width)
	t.root.Draw(t.canvas)
	t.canvas.End()
	return nil
}
