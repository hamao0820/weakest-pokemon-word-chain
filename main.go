package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed data/words.json
var wordsJSON []byte

type Pokemon struct {
	No   int    `json:"no"`
	Name string `json:"name"`
}

type Word struct {
	Pokemon
	Ruby   string   `json:"reading"`
	Start  string   `json:"start"`
	End    []string `json:"end"`
	IsLast bool     `json:"is_last"`
}

var words []*Word
var noMap map[int]*Word
var startMap map[string][]*Word

func init() {
	json.Unmarshal(wordsJSON, &words)

	noMap = make(map[int]*Word)
	startMap = make(map[string][]*Word)
	for _, w := range words {
		noMap[w.No] = w
		startMap[w.Start] = append(startMap[w.Start], w)
	}
}

func main() {
	li := map[int][]Word{} // 隣接リスト
	for i := 0; i < len(words); i++ {
		w := words[i]
		for _, e := range w.End {
			for _, v := range startMap[e] {
				li[w.No] = append(li[w.No], *v)
			}
		}
	}

	bfs := func(start int) []int {
		if _, ok := noMap[start]; !ok {
			return nil
		}
		visited := map[int]bool{}
		queue := [][]int{{start}}
		for len(queue) > 0 {
			path := queue[0]
			queue = queue[1:]
			node := path[len(path)-1]
			if visited[node] {
				continue
			}
			visited[node] = true
			for _, v := range li[node] {
				newPath := append(path, v.No)
				queue = append(queue, newPath)
				if v.IsLast {
					return newPath
				}
			}
		}
		return nil
	}

	for k := range noMap {
		path := bfs(k)
		for _, p := range path {
			fmt.Printf("%s ", noMap[p].Name)
		}
	}
}
