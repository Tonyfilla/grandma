package main

import (
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/tjarratt/babble"
	"os"
	"sort"
	"strings"
)

// построить граф алфавита
// написать тесты
// подготовить тестовы сет

func generatePassword(dictionary []string) string {
	weightList := make(map[int][]string)
	// построить мапу с весом для слов(вес -> связный список)
	for _, word := range dictionary {
		word = strings.ToUpper(word)
		// отбросить ненужные слова (меньше 5 букв)
		if len(word) < 5 || len(word) > 10 {
			continue
		}
		fmt.Println(word)
		count := countPath(word)
		val, ok := weightList[count]
		if !ok {
			weightList[count] = []string{word}
		}
		val = append(val, word)
		weightList[count] = val
	}
	fmt.Println(weightList)
	keys := make([]int, 0, len(weightList))
	for k := range weightList {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// TODO: дописать поиск лучшего соседа,походу рекрурсия нужна
	var password []string
	lesspath := 1000
	bestNeibh := ""
	for index, val := range keys {
		set := weightList[val]
		if len(set) > 1 {
			for _, samesetVal := range set {
				path := countPath(set[0] + samesetVal)
				if path < lesspath {
					lesspath = path
					bestNeibh = samesetVal
				}
			}
		}
		for _, neibSetKey := range keys[index:] {
			if neibSetKey+val > lesspath {
				password = append(password, bestNeibh)
				if len(password) > 3 {
					return fmt.Sprint(password)
				}
				lesspath = 1000
				bestNeibh = ""
				break
			}
			fmt.Println(index, neibSetKey)
			neibSet := weightList[neibSetKey]
			for _, v := range neibSet {
				path := countPath(set[0] + v)
				if path < lesspath {
					lesspath = path
					bestNeibh = v
				}
			}
		}
	}
	return fmt.Sprint(password)
}

func countPath(word string) int {
	runes := []rune(word)
	counter := 0
	for index, _ := range runes {
		if index < len(runes)-1 {
			path, _ := graph.ShortestPath(g, int(runes[index])-65, int(runes[index+1])-65)
			counter = counter + len(path)
		}
	}
	return counter
}

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	//				   0    1    2    3    4    5    6    7    8   9    10   11   12
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//  13   14   15   16   17   18   19   20   21   22   23   24    25

var g = graph.New(graph.IntHash)

func makeGraph() {
	for index, _ := range arr {
		_ = g.AddVertex(index)
	}
	//Q
	_ = g.AddEdge(16, 22)
	//W
	_ = g.AddEdge(22, 4)
	//E
	_ = g.AddEdge(4, 17)
	//R
	_ = g.AddEdge(17, 19)
	//T
	_ = g.AddEdge(19, 24)
	//Y
	_ = g.AddEdge(24, 20)
	//U
	_ = g.AddEdge(20, 8)
	//I
	_ = g.AddEdge(8, 14)
	//O
	_ = g.AddEdge(14, 15)

	//A
	_ = g.AddEdge(0, 18)
	_ = g.AddEdge(0, 16)
	//S
	_ = g.AddEdge(18, 3)
	_ = g.AddEdge(18, 22)
	//D
	_ = g.AddEdge(3, 4)
	_ = g.AddEdge(3, 5)
	//F
	_ = g.AddEdge(5, 17)
	_ = g.AddEdge(5, 6)
	//G
	_ = g.AddEdge(6, 19)
	_ = g.AddEdge(6, 7)
	//H
	_ = g.AddEdge(7, 24)
	_ = g.AddEdge(7, 9)
	//J
	_ = g.AddEdge(9, 20)
	_ = g.AddEdge(9, 10)
	//K
	_ = g.AddEdge(10, 8)
	_ = g.AddEdge(10, 11)
	//L
	_ = g.AddEdge(11, 14)

	//Z
	_ = g.AddEdge(25, 23)
	_ = g.AddEdge(25, 0)
	//X
	_ = g.AddEdge(23, 2)
	_ = g.AddEdge(23, 18)
	//C
	_ = g.AddEdge(2, 3)
	_ = g.AddEdge(2, 21)
	//V
	_ = g.AddEdge(21, 1)
	_ = g.AddEdge(21, 5)
	//B
	_ = g.AddEdge(23, 2)
	_ = g.AddEdge(23, 18)
	//N
	_ = g.AddEdge(13, 7)
	_ = g.AddEdge(13, 12)
	//M
	_ = g.AddEdge(12, 9)
	//_ = g.AddEdge(23, 18)
	//красивое
	file, _ := os.Create("./simple1.gv")
	_ = draw.DOT(g, file)

}

func main() {
	// составить граф клавиатуры
	makeGraph()
	//сгенерить входные данные
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	babbler.Count = 200
	pass := generatePassword(strings.Split(babbler.Babble(), " "))
	println(pass)
}
