package main

import "testing"

type testpair struct {
	stream []byte
	count  int
}

var testsGroupScore = []testpair{
	{[]byte("{}"), 1},
	{[]byte("{{{}}}"), 6},
	{[]byte("{{},{}}"), 5},
	{[]byte("{{{},{},{{}}}}"), 16},
	{[]byte("{<a>,<a>,<a>,<a>}"), 1},
	{[]byte("{{<ab>},{<ab>},{<ab>},{<ab>}}"), 9},
	{[]byte("{{<!!>},{<!!>},{<!!>},{<!!>}}"), 9},
	{[]byte("{{<a!>},{<a!>},{<a!>},{<ab>}}"), 3},
}

func TestCalculateTotalGroupScore(t *testing.T) {
	for _, pair := range testsGroupScore {
		totalScore, _ := CalculateStreamScores(pair.stream)
		if totalScore != pair.count {
			t.Error(
				"For", pair.stream,
				"expected", pair.count,
				"got", totalScore,
			)
		}
	}
}

var testsGarbageCount = []testpair{
	{[]byte("<>"), 0},
	{[]byte("<random characters>"), 17},
	{[]byte("<<<<>"), 3},
	{[]byte("<{!>}>"), 2},
	{[]byte("<!!>"), 0},
	{[]byte("<!!!>>"), 0},
	{[]byte("<{o\"i!a,<{i<a>"), 10},
}

func TestGarbageScore(t *testing.T) {
	for _, pair := range testsGarbageCount {
		_, garbageScore := CalculateStreamScores(pair.stream)
		if garbageScore != pair.count {
			t.Error(
				"For", pair.stream,
				"expected", pair.count,
				"got", garbageScore,
			)
		}
	}
}
