package triples_processors

import (
	"github.com/wallix/triplestore"
	"rdf_processor/code/triples_processors/internal"
)

func get_triples_datasets() []triplestore.Source {

	var source_triples []triplestore.Source

	triples_dataset_1 :=
		internal.
			Read_csv_data("./input/1.tsv", "tab")
	triples_dataset_2 :=
		internal.
			Read_csv_data("./input/2.tsv", "tab")

	triple_store_1 :=
		get_triple_dataset(
			triples_dataset_1)

	triple_store_2 :=
		get_triple_dataset(
			triples_dataset_2)

	source_triples = append(source_triples, triple_store_1, triple_store_2)

	return source_triples
}

func get_triple_dataset(
	triples_dataset [][]string) triplestore.Source {

	triple_store :=
		triplestore.
			NewSource()

	for _, triple_data := range triples_dataset {

		walli_triple :=
			triplestore.
				SubjPred(
					triple_data[0],
					triple_data[1]).
				Resource(
					triple_data[2])

		triple_store.
			Add(
				walli_triple)

	}

	return triple_store
}
