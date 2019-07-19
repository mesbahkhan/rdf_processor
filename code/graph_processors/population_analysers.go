package graph_processors

import (
	"fmt"
	"github.com/wallix/triplestore"
)

func Analyse_graphs_population(
	triple_input_graphs []triplestore.RDFGraph) {

	for index, graph := range triple_input_graphs {

		Analyse_graph_population(graph, index)

	}

}

func Analyse_graph_population(
	graph triplestore.RDFGraph,
	index int) {

	number_of_triples := len(graph.Triples())
	triples := graph.Triples()
	count_of_objects := count_objects(triples)
	count_of_subjects := count_subjects(triples)
	count_of_predicates := count_predicates(triples)

	fmt.Printf(
		"graph number: %v\nnumber of triples: %v\nnumber of subjects :%v\nnumber of predicates :%v\nnumber of objects :%v\n",
		index,
		number_of_triples,
		count_of_subjects,
		count_of_predicates,
		count_of_objects)

}

func count_objects(
	triples []triplestore.Triple) int {

	unique_objects_map := map[triplestore.Object]bool{}

	var unique_objects_list []triplestore.Object

	for _, triple := range triples {

		if !unique_objects_map[triple.Object()] {
			unique_objects_map[triple.Object()] = true
			unique_objects_list = append(unique_objects_list, triple.Object())
		}

	}
	return len(unique_objects_list)
}

func count_subjects(
	triples []triplestore.Triple) int {

	unique_objects_map := map[string]bool{}

	var unique_objects_list []string

	for _, triple := range triples {

		if !unique_objects_map[triple.Subject()] {
			unique_objects_map[triple.Subject()] = true
			unique_objects_list = append(unique_objects_list, triple.Subject())
		}

	}
	return len(unique_objects_list)
}

func count_predicates(
	triples []triplestore.Triple) int {

	unique_objects_map := map[string]bool{}

	var unique_objects_list []string

	for _, triple := range triples {

		if !unique_objects_map[triple.Predicate()] {
			unique_objects_map[triple.Predicate()] = true
			unique_objects_list = append(unique_objects_list, triple.Predicate())
		}

	}
	return len(unique_objects_list)
}
