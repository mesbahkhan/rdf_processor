package triples_analysers

import (
	"github.com/wallix/triplestore"
	"rdf_processor/code/graph_processors"
	"rdf_processor/code/triples_processors"
)

func Analyse_triples() (
	[]triplestore.RDFGraph,
	triplestore.RDFGraph) {

	triple_input_graphs :=
		triples_processors.
			Get_triples_graphs()

	graph_processors.
		Analyse_graphs_population(
			triple_input_graphs)

	merged_difference_graph :=
		report_differences_between_graphs(
			triple_input_graphs)

	return triple_input_graphs, merged_difference_graph
}

func report_differences_between_graphs(
	triple_input_graphs []triplestore.RDFGraph) triplestore.RDFGraph {

	merged_difference_graph :=
		compare_triples_graphs(
			triple_input_graphs)

	graph_processors.Draw_graph(
		merged_difference_graph,
		"merged_difference_graph.dot")

	return merged_difference_graph
}

func compare_triples_graphs(
	triple_input_graphs []triplestore.RDFGraph) triplestore.RDFGraph {

	difference_graphs :=
		graph_processors.
			Get_differences_between_graphs(
				triple_input_graphs)

	merged_difference_graph :=
		graph_processors.
			Merge_graphs(
				difference_graphs)

	return merged_difference_graph
}
