package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"syscall/js"
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
	c := make(chan struct{}, 0)

	li := map[int][]Word{} // 隣接リスト
	for i := 0; i < len(words); i++ {
		w := words[i]
		for _, e := range w.End {
			for _, v := range startMap[e] {
				li[w.No] = append(li[w.No], *v)
			}
		}
	}

	getShortestChain := func(start int) []int {
		if _, ok := noMap[start]; !ok {
			return nil
		}
		if noMap[start].IsLast {
			return []int{start}
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

	js.Global().Set("goGetShortestChain", js.FuncOf(func(_ js.Value, args []js.Value) any {
		start := args[0].Int()
		p := getShortestChain(start)
		arr := js.Global().Get("Array").New(len(p))
		for i, v := range p {
			arr.SetIndex(i, js.ValueOf(v))
		}
		return arr
	}))

	fmt.Println("Wasm Go Initialized")
	<-c
}
