package main

import (
	"bufio"
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"os"
	"sort"
	"strings"
)

func main() {
	makeGraph()
	// make a map with weight and list of words
	file, err := os.Open("../../../Downloads/words_alpha.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	weightList := make(map[int][]string)
	for scanner.Scan() {
		word := strings.ToUpper(scanner.Text())
		if len(word) < 4 {
			continue
		}
		count := countPath(word)
		val, ok := weightList[count]
		if !ok {
			weightList[count] = []string{word}
		}
		val = append(val, word)
		weightList[count] = val
	}
	keys := make([]int, 0, len(weightList))
	for k := range weightList {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	password := ""
	startWordIndex := 0
	startWord := weightList[keys[startWordIndex]][0]
	for len(password) < 20 {
		minPathWeight := 1000
		nextword := ""
		nextwordIndex := 0
		password += startWord
		for indexKey := range keys {
			// exit if current weight more than minPath
			if keys[indexKey] >= minPathWeight {
				fmt.Println("the deep of search - ", keys[indexKey])
				// delete word from list, words should be different
				list := weightList[keys[startWordIndex]]
				if len(list) == 1 {
					delete(weightList, startWordIndex)
				} else {
					for in := range list {
						if list[in] == startWord {
							list[in] = list[len(list)-1]
							weightList[keys[startWordIndex]] = list[:len(list)-2]
							break
						}
					}
				}
				break
			}
			for _, wordI := range weightList[keys[indexKey]] {
				if wordI == startWord {
					continue
				}
				Sword := []rune(wordI)
				Fword := []rune(startWord)
				between := string(Fword[len(Fword)-1]) + string(Sword[0])
				path := countPath(between)
				if minPathWeight > path+keys[indexKey] {
					minPathWeight = path + keys[indexKey]
					nextword = wordI
					nextwordIndex = indexKey
				}
			}
		}
		fmt.Println(startWord)
		fmt.Println(minPathWeight)
		fmt.Println(nextword)
		fmt.Println("--------------------------")

		startWord = nextword
		startWordIndex = nextwordIndex
	}
	fmt.Println(password)
}

var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	//				   0    1    2    3    4    5    6    7    8   9    10   11   12
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//  13   14   15   16   17   18   19   20   21   22   23   24    25

var g = graph.New(graph.IntHash)

func makeGraph() {
	for index := range arr {
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

var (
	hashPathAlphabet = make(map[string]int, 52)
)

func countPath(word string) int {
	runes := []rune(word)
	counter := 0
	for index := range runes {
		if index < len(runes)-1 {
			if runes[index] == runes[index+1] {
				continue
			}
			res, ok := hashPathAlphabet[string(runes[index])+string(runes[index+1])]
			if !ok {
				path, _ := graph.ShortestPath(g, int(runes[index])-65, int(runes[index+1])-65)
				hashPathAlphabet[string(runes[index])+string(runes[index+1])] = len(path)
				hashPathAlphabet[string(runes[index+1])+string(runes[index])] = len(path)
				counter += len(path)
			}
			counter += res
		}
	}
	return counter
}
