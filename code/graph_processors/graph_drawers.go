package graph_processors

import (
	"fmt"
	"github.com/emicklei/dot"
	"github.com/wallix/triplestore"
	"io/ioutil"
	"rdf_processor/code/helpers"
)

func Generate_input_triples_graphs(
	triple_stores []triplestore.Source) []triplestore.RDFGraph {

	var triple_graphs []triplestore.RDFGraph

	graph_1 :=
		triple_stores[0].
			Snapshot()
	graph_2 :=
		triple_stores[1].
			Snapshot()

	Draw_graph(
		graph_1,
		"graph_1.dot")
	Draw_graph(
		graph_2,
		"graph_2.dot")

	triple_graphs = append(triple_graphs, graph_1, graph_2)

	return triple_graphs
}

func Draw_graph(
	graph triplestore.RDFGraph,
	filename string) {

	g := dot.NewGraph(dot.Directed)

	triples := graph.Triples()

	for _, triple := range triples {

		subject_node := g.Node(triple.Subject())
		resource_string, _ := triple.Object().Resource()
		object_node := g.Node(resource_string)

		g.Edge(subject_node, object_node, triple.Predicate())

	}

	fmt.Println(g.String())

	err := ioutil.WriteFile("./outputs/"+filename, []byte(g.String()), 0644)
	helpers.Check_error(err)

}
