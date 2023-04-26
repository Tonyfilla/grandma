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

type Dictionary struct {
	weightList map[int][]string
	keys       []int
}

func CreateDict(weightList map[int][]string, keys []int) *Dictionary {
	return &Dictionary{
		weightList: weightList,
		keys:       keys,
	}
}

func (d *Dictionary) First() (string, int, error) {
	if len(d.keys) > 0 {
		list := d.weightList[d.keys[0]]
		if len(list) < 1 {
			fmt.Print(list, d.keys[0], d)
		}
		return list[0], d.keys[0], nil
	}
	return "", 0, fmt.Errorf("empty")
}

func (d *Dictionary) Next(word string, key int) (string, int, error) {
	list := d.weightList[key]
	for i, val := range list {
		if val == word {
			if i < len(list)-1 {
				return list[i+1], key, nil
			}
			for ik, kk := range d.keys {
				if kk == key {
					if ik == len(d.keys)-1 {
						return "", 0, fmt.Errorf("the end")
					}
					list = d.weightList[d.keys[ik+1]]
					return list[0], d.keys[ik+1], nil
				}
			}
		}
	}
	return "", 0, fmt.Errorf("the end")
}

func (d *Dictionary) Del(word string, key int) {
	list := d.weightList[key]
	for i, val := range list {
		if val == word {
			if len(list) == 1 {
				delete(d.weightList, key)
				for ki, kk := range d.keys {
					if kk == key {
						keys := d.keys
						d.keys = keys[:ki]
						d.keys = append(d.keys, keys[ki+1:]...)
						break
					}
				}
			} else {
				list[i] = list[len(list)-1]
				d.weightList[key] = list[:len(list)-1]
			}
			break
		}
	}
}

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
	dict := CreateDict(weightList, keys)
	k := keys[0]
	sourceWord := weightList[k][0]
	researchWord, researchKey, _ := dict.Next(sourceWord, keys[0])
	dict.Del(sourceWord, k)
	var pass []string
	pass = append(pass, sourceWord)
	for len(pass) < 4 {
		word, key, path := dict.findNeighbor(10000, researchKey, 0, k+12, "", sourceWord, researchWord)
		println(word, key, path)
		pass = append(pass, word)
		dict.Del(word, key)
		researchWord, researchKey, _ = dict.First()
		sourceWord = word
	}
	return strings.Join(pass, " ")
}

// ищем лучшую пару
// удалеем из словаря пару
// ищем соседа
// ищем соседа

// смотрим на минимум
// смотримм на последнюю букву
// определям самый дальний угол клавы
// определяем глубину поиска по группам
// запускаем поиск лучшего соседа на глубину

func (d *Dictionary) findNeighbor(minPath, researchKey, keyGroup, deep int, bestNeighbor, sourceWord, researchWord string) (string, int, int) {

	// поиск лучшего соседа
	// запоминаем минимальный результат  и слово и группу
	// берем следующее слово - проверяем лучше минимума или нет
	// если да  записываем результат и берем следующее слово
	if deep < researchKey {
		return bestNeighbor, keyGroup, minPath
	}
	if researchWord != sourceWord {
		counter := countPath(sourceWord + researchWord)
		if counter < minPath {
			minPath = counter
			keyGroup = researchKey
			bestNeighbor = researchWord
		}
	}
	researchWord, researchKey, err := d.Next(researchWord, researchKey)
	if err != nil {
		return bestNeighbor, keyGroup, minPath
	}
	// next for neighbor group key
	return d.findNeighbor(minPath, researchKey, keyGroup, deep, bestNeighbor, sourceWord, researchWord)
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
