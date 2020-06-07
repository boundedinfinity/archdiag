package archdiag

type TextAnchor string

const (
	TextAnchorStart  TextAnchor = "start"
	TextAnchorMiddle TextAnchor = "middle"
	TextAnchorEnd    TextAnchor = "end"
)

type Direction string

const (
	DirectionVertical   Direction = "vertical"
	DirectionHorizontal Direction = "horizontal"
)

type Format string

const (
	FormatYaml Format = "yaml"
	FormatJson Format = "json"
)

type Graph struct {
	Name        string                `json:"name" yaml:"name"`
	Label       string                `json:"label" yaml:"label"`
	Nodes       map[string]Node       `json:"nodes" yaml:"nodes"`
	Connections map[string]Connection `json:"connections" yaml:"connections"`
	GraphAttributes
	boxStyle
}

func (t Graph) CenterPoint() point {
	return point{
		X: t.Width / 2,
		Y: t.Height / 2,
	}
}

var DEFAULT_GRAPH_ATTRIBUTES = GraphAttributes{
	Height:    250,
	Width:     250,
	Direction: DirectionHorizontal,
}

type GraphAttributes struct {
	Height    int       `json:"height" yaml:"height"`
	Width     int       `json:"width" yaml:"width"`
	Direction Direction `json:"direction" yaml:"direction"`
}

func calcGraphAttributes(l0 GraphAttributes) GraphAttributes {
	return GraphAttributes{
		Height:    getInt(DEFAULT_GRAPH_ATTRIBUTES.Height, l0.Height),
		Width:     getInt(DEFAULT_GRAPH_ATTRIBUTES.Width, l0.Width),
		Direction: Direction(getStr(string(DEFAULT_GRAPH_ATTRIBUTES.Direction), string(l0.Direction))),
	}
}

type Node struct {
	Label string          `json:"label" yaml:"label"`
	Nodes map[string]Node `json:"nodes" yaml:"nodes"`
	boxStyle
}

type Connection struct {
	Label string `json:"label" yaml:"label"`
	From  string `json:"from" yaml:"from"`
	To    string `json:"to" yaml:"to"`
}

type point struct {
	X int
	Y int
}

type dimensions struct {
	H int
	W int
}
