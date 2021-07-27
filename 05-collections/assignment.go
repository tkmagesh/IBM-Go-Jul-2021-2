/*
Given the following string:

Proident occaecat ex esse do duis mollit sunt id ea ut non consectetur incididunt Consectetur culpa non amet ipsum et ex duis Minim reprehenderit est officia est quis laborum laboris aute laboris laboris ut labore Lorem Reprehenderit culpa Lorem non duis nisi mollit cillum eiusmod aliquip est laborum mollit duis Aliqua eu tempor ad aliquip esse nostrud laborum excepteur ipsum excepteur eiusmod Sint nisi duis esse occaecat officia ut amet anim ex veniam est officia labore Elit esse exercitation officia quis ipsum adipisicing fugiat ex irure Culpa dolore laborum amet deserunt aliqua esse duis pariatur cillum est adipisicing Lorem eu Elit anim aute eiusmod consequat exercitation anim laborum consectetur excepteur adipisicing cillum Nostrud proident consequat cupidatat velit fugiat duis Ex tempor eu qui elit nostrud amet Nostrud dolore anim ex aliqua reprehenderit dolor velit adipisicing Consectetur eiusmod incididunt qui proident quis id reprehenderit Sit anim ea eu magna mollit laborum amet voluptate Dolore sunt laboris consectetur tempor nisi ut duis veniam cupidatat do Qui culpa ipsum incididunt sunt velit ipsum Est elit duis aliqua qui dolor aliqua enim laborum reprehenderit magna fugiat consectetur esse Quis nulla fugiat incididunt eu


Find out the size of the word that occurs the most by size

ex:
10 - 2 letters
7 - 3 letters
14 - 4 letters

Output:
	"The size of the word that occurs the most(14 times) is 4"

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Proident occaecat ex esse do duis mollit sunt id ea ut non consectetur incididunt Consectetur culpa non amet ipsum et ex duis Minim reprehenderit est officia est quis laborum laboris aute laboris laboris ut labore Lorem Reprehenderit culpa Lorem non duis nisi mollit cillum eiusmod aliquip est laborum mollit duis Aliqua eu tempor ad aliquip esse nostrud laborum excepteur ipsum excepteur eiusmod Sint nisi duis esse occaecat officia ut amet anim ex veniam est officia labore Elit esse exercitation officia quis ipsum adipisicing fugiat ex irure Culpa dolore laborum amet deserunt aliqua esse duis pariatur cillum est adipisicing Lorem eu Elit anim aute eiusmod consequat exercitation anim laborum consectetur excepteur adipisicing cillum Nostrud proident consequat cupidatat velit fugiat duis Ex tempor eu qui elit nostrud amet Nostrud dolore anim ex aliqua reprehenderit dolor velit adipisicing Consectetur eiusmod incididunt qui proident quis id reprehenderit Sit anim ea eu magna mollit laborum amet voluptate Dolore sunt laboris consectetur tempor nisi ut duis veniam cupidatat do Qui culpa ipsum incididunt sunt velit ipsum Est elit duis aliqua qui dolor aliqua enim laborum reprehenderit magna fugiat consectetur esse Quis nulla fugiat incididunt eu"

	words := strings.Split(str, " ")
	wordCountBySize := getWordsCountBySize(words)
	wordSize, count := getMaxWordsCountBySize(wordCountBySize)
	fmt.Printf("The size of the word that occurs the most(%d times) is %d\n", count, wordSize)

}

func getWordsCountBySize(words []string) *map[int]int {
	wordCountBySize := map[int]int{}
	for _, word := range words {
		wordCountBySize[len(word)] += 1
	}
	return &wordCountBySize
}

func getMaxWordsCountBySize(wordCountBySize *map[int]int) (wordSize, count int) {
	for size, cnt := range *wordCountBySize {
		if count < cnt {
			count = cnt
			wordSize = size
		}
	}
	return
}
