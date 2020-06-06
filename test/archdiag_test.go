package test

import (
	"fmt"
	"testing"

	"github.com/bounded-infinity/archdiag"
	"github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBooks(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "archdiag suite")
}

func assertGraph(actual, expected archdiag.Graph) {
	p := "archdiag.Graph.Nodes"
	Expect(actual.Name).Should(Equal(expected.Name), "archdiag.Graph.Name")
	Expect(actual.Direction).Should(Equal(expected.Direction), "archdiag.Graph.Direction")
	assertNodes(p, actual.Nodes, expected.Nodes)
}

func assertNodes(p string, actualMap, expectedMap map[string]archdiag.Node) {
	if expectedMap == nil {
		Expect(actualMap).Should(BeNil(), "%v == nil", p)
		return
	}

	Expect(len(actualMap)).Should(Equal(len(expectedMap)), "len(%v)", p)

	for id, expected := range expectedMap {
		np := fmt.Sprintf(`%v["%v"]`, p, id)
		actual, aok := actualMap[id]
		Expect(aok).Should(BeTrue(), np)
		assertNode(np, actual, expected)
	}
}

func assertNode(p string, a, e archdiag.Node) {
	Expect(a.Label).Should(Equal(e.Label), "%v.Label", p)
	assertNodes(fmt.Sprintf("%v.Nodes", p), a.Nodes, e.Nodes)
}

func readFile(d *archdiag.ArchDiag, g archdiag.Graph) error {
	fn := fmt.Sprintf("./%v", g.Name)
	return d.ReadFromFile(fn)
}

var _ = ginkgo.Describe("g1", func() {
	d := &archdiag.ArchDiag{}
	expected := archdiag.Graph{
		Nodes: map[string]archdiag.Node{
			"server-1": {},
		},
	}

	ginkgo.It("read json", func() {
		expected.Name = "g1.json"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})

	ginkgo.It("read yaml", func() {
		expected.Name = "g1.yaml"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})
})

var _ = ginkgo.Describe("g2", func() {
	d := &archdiag.ArchDiag{}
	expected := archdiag.Graph{
		Nodes: map[string]archdiag.Node{
			"server-1": {},
			"server-2": {},
		},
	}

	ginkgo.It("read json", func() {
		expected.Name = "g2.json"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})

	ginkgo.It("read yaml", func() {
		expected.Name = "g2.yaml"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})
})

var _ = ginkgo.Describe("g4", func() {
	d := &archdiag.ArchDiag{}
	expected := archdiag.Graph{
		GraphAttributes: archdiag.GraphAttributes{
			Direction: archdiag.DirectionVertical,
		},
		Nodes: map[string]archdiag.Node{
			"dc-1": {
				Nodes: map[string]archdiag.Node{
					"server-1": {},
					"server-2": {},
				},
			},
		},
	}

	ginkgo.It("read json", func() {
		expected.Name = "g4.json"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})

	ginkgo.It("read yaml", func() {
		expected.Name = "g4.yaml"
		Expect(readFile(d, expected)).Should(Succeed())
		assertGraph(d.G, expected)
	})
})
