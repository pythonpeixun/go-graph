package main

import (
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	tests := []struct {
		start string
		adj   map[string][]string
	}{
		{
			start: "s",
			adj:   getGraph(),
		},
	}

	for _, test := range tests {
		bfs(test.start, test.adj)
	}
}

func TestRelax(t *testing.T) {
	tests := []struct {
		start string
		adj   map[string][]string
		want  []string
	}{
		{
			start: "s",
			adj:   getGraph(),
			want:  []string{"a", "x"},
		},
	}

	for _, test := range tests {
		candidate := relax(test.start, test.adj)
		if !reflect.DeepEqual(candidate, test.want) {
			t.Errorf("want: %v, got: %v\n", test.want, candidate)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		candidate  []string
		visited    map[string]struct{}
		want       []string
		newVisited map[string]struct{}
	}{
		{
			candidate: []string{"a", "x"},
			visited:   map[string]struct{}{},
			want:      []string{"a", "x"},
			newVisited: map[string]struct{}{
				"a": struct{}{},
				"x": struct{}{},
			},
		},
		{
			candidate: []string{"a", "x"},
			visited:   map[string]struct{}{"a": struct{}{}},
			want:      []string{"x"},
			newVisited: map[string]struct{}{
				"a": struct{}{},
				"x": struct{}{},
			},
		},
		{
			candidate: []string{"a", "x"},
			visited: map[string]struct{}{
				"a": struct{}{},
				"x": struct{}{},
			},
			newVisited: map[string]struct{}{
				"a": struct{}{},
				"x": struct{}{},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		next := filter(test.candidate, test.visited)
		if !reflect.DeepEqual(next, test.want) {
			t.Errorf("want: %v, got: %v\n", test.want, next)
		}

		if !reflect.DeepEqual(test.visited, test.newVisited) {
			t.Errorf("[malformed visited] want: %v, got: %v\n", test.newVisited, test.visited)
		}
	}

}

func getGraph() map[string][]string {
	return map[string][]string{
		"s": []string{"a", "x"},
		"a": []string{"s", "z"},
		"x": []string{"s", "d", "c"},

		"z": []string{"a"},
		"d": []string{"x", "f"},
		"c": []string{"v"},

		"f": nil,
		"v": nil,
	}
}
