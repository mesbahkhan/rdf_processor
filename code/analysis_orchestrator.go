package code

import (
	"rdf_processor/code/graph_processors"
	"rdf_processor/code/triples_analysers"
	"rdf_processor/code/triples_processors"
)

func Orchestrate_analysis() {

	//Analyse

	triple_input_graphs,
		merged_difference_graph :=
		triples_analysers.
			Analyse_triples()

	//Prepare Mapping Inputs

	mapping_graph :=
		triples_processors.
			Prepare_mapping_graph()

	//Merge Mapping with difference graph

	graph_processors.
		Merge_mapping_to_difference_graph(
			merged_difference_graph,
			mapping_graph)

	//Merge mapping and
	graph_processors.
		Merge_all(
			triple_input_graphs,
			mapping_graph)
}
