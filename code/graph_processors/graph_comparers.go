package graph_processors

import "github.com/wallix/triplestore"

func Get_differences_between_graphs(
	graphs []triplestore.RDFGraph) []triplestore.RDFGraph {

	graph_1 := graphs[0]
	graph_2 := graphs[1]

	extra_graph_1_triples_graph :=
		compare_graph_pair(
			graph_1,
			graph_2)

	extra_graph_2_triples_graph :=
		compare_graph_pair(
			graph_2,
			graph_1)

	Draw_graph(
		extra_graph_1_triples_graph,
		"extra_in_1.dot")

	Draw_graph(
		extra_graph_2_triples_graph,
		"extra_in_2.dot")

	var merged_difference_graphs []triplestore.RDFGraph

	merged_difference_graphs =
		append(
			merged_difference_graphs,
			extra_graph_1_triples_graph,
			extra_graph_2_triples_graph)

	return merged_difference_graphs

}

func compare_graph_pair(
	graph_1 triplestore.RDFGraph,
	graph_2 triplestore.RDFGraph) triplestore.RDFGraph {

	graph_1_triples :=
		graph_1.
			Triples()

	graph_1_triples_not_in_graph_2 :=
		triplestore.
			NewSource()

	for _, graph_1_triple := range graph_1_triples {

		if !graph_2.Contains(graph_1_triple) {
			graph_1_triples_not_in_graph_2.Add(graph_1_triple)
		}
	}
	graph_1_triples_not_in_graph_2_graph := graph_1_triples_not_in_graph_2.Snapshot()

	return graph_1_triples_not_in_graph_2_graph
}
