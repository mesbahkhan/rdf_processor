package triples_processors

import (
	"github.com/wallix/triplestore"
	"rdf_processor/code/graph_processors"
	"rdf_processor/code/triples_processors/internal"
)

func Get_triples_graphs() []triplestore.RDFGraph {

	triple_stores :=
		get_triples_datasets()

	triple_input_graphs :=
		graph_processors.Generate_input_triples_graphs(
			triple_stores)

	return triple_input_graphs
}

func Prepare_mapping_graph() triplestore.RDFGraph {

	mapping_dataset :=
		internal.
			Read_csv_data("./input/mapping.tsv", "tab")

	mapping_store :=
		get_triple_dataset(
			mapping_dataset)

	mapping_graph :=
		mapping_store.
			Snapshot()

	return mapping_graph
}
