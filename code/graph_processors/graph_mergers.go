package graph_processors

import (
	"github.com/wallix/triplestore"
)

func Merge_all(
	triple_input_graphs []triplestore.RDFGraph,
	mapping_graph triplestore.RDFGraph) {
	merged_graph_excluding_mapping :=
		Merge_graphs(
			triple_input_graphs)
	var triple_input_and_mapping_graphs []triplestore.RDFGraph
	triple_input_and_mapping_graphs =
		append(
			triple_input_and_mapping_graphs,
			merged_graph_excluding_mapping,
			mapping_graph)
	merged_graph :=
		Merge_graphs(
			triple_input_and_mapping_graphs)
	Draw_graph(
		merged_graph,
		"merged_graph_including_mapping.dot")

	Analyse_graph_population(
		merged_graph,
		1)

}

func Merge_mapping_to_difference_graph(
	merged_difference_graph triplestore.RDFGraph,
	mapping_graph triplestore.RDFGraph) {
	var graph_set_for_merging_mapping []triplestore.RDFGraph
	graph_set_for_merging_mapping = append(
		graph_set_for_merging_mapping,
		merged_difference_graph,
		mapping_graph)
	merged_difference_graph_with_mapping :=
		Merge_graphs(
			graph_set_for_merging_mapping)
	Draw_graph(
		merged_difference_graph_with_mapping,
		"merged_differences_with_mapping_graph.dot")
}

func Merge_graphs(
	triple_graphs []triplestore.RDFGraph) triplestore.RDFGraph {

	triple_graph_2_triples := triple_graphs[1].Triples()
	triple_graph_1_triples := triple_graphs[0].Triples()

	merged_tipple_store := triplestore.NewSource()

	for _, triple_store_1_triple := range triple_graph_1_triples {
		merged_tipple_store.Add(triple_store_1_triple)
	}

	for _, triple_store_2_triple := range triple_graph_2_triples {
		merged_tipple_store.Add(triple_store_2_triple)
	}

	merged_triple_graph := merged_tipple_store.Snapshot()

	return merged_triple_graph
}
