package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PartTwo() {

	var sum int = 0

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var min string
		var max string

		line := scanner.Text()

		array := []byte(line)
		trieForward := Constructor([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"})
		trieReverse := ConstructorReverse([]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"})

		for _, c := range array {

			if c <= 'a' {
				min = string(c)
				break
			}
			min = trieForward.Query(byte(c))

			if min != "" {
				break
			}

		}

		for i := len(array) - 1; i >= 0; i-- {
			if array[i] >= '0' && array[i] <= '9' {
				max = string(array[i])
				break
			}

			max = trieReverse.Query(byte(array[i]))

			if max != "" {
				break
			}
		}

		lineInt, _ := strconv.Atoi(doubleIfSingleDigit(min + max))

		sum += lineInt

	}

	fmt.Println(sum)
}

func doubleIfSingleDigit(n string) string {
	if len(n) == 1 {
		return n + n
	}
	return n
}

type TrieNode struct {
	children [26]*TrieNode
	word     bool
	value    string
}

type StreamChecker struct {
	trie   *TrieNode
	stream []byte
}

func Constructor(words []string) StreamChecker {

	var value string

	sc := StreamChecker{
		trie:   &TrieNode{},
		stream: []byte{},
	}

	// populate trie
	for _, word := range words {

		switch word {
		case "one":
			value = "1"
		case "two":
			value = "2"
		case "three":
			value = "3"
		case "four":
			value = "4"
		case "five":
			value = "5"
		case "six":
			value = "6"
		case "seven":
			value = "7"
		case "eight":
			value = "8"
		case "nine":
			value = "9"
		default:
			value = ""
		}

		node := sc.trie
		reverserdWord := reverse(word)

		for _, c := range reverserdWord {

			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}

			node = node.children[idx]
		}
		node.word = true
		node.value = value
	}
	return sc
}

func ConstructorReverse(words []string) StreamChecker {

	var value string

	sc := StreamChecker{
		trie:   &TrieNode{},
		stream: []byte{},
	}

	// populate trie
	for _, word := range words {

		switch word {
		case "one":
			value = "1"
		case "two":
			value = "2"
		case "three":
			value = "3"
		case "four":
			value = "4"
		case "five":
			value = "5"
		case "six":
			value = "6"
		case "seven":
			value = "7"
		case "eight":
			value = "8"
		case "nine":
			value = "9"
		default:
			value = ""
		}

		node := sc.trie
		for _, c := range word {

			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
			}

			node = node.children[idx]
		}
		node.word = true
		node.value = value
	}
	return sc
}

func (this *StreamChecker) Query(letter byte) string {

	this.stream = append([]byte{letter}, this.stream...)
	node := this.trie
	for _, c := range this.stream {

		idx := c - 'a'
		if node.word == true {
			return node.value
		}

		if node.children[idx] == nil {
			return ""
		}
		node = node.children[idx]
	}
	return node.value
}

func reverse(word string) []byte {
	res := []byte(word)

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {

		res[i], res[j] = res[j], res[i]

	}

	return res

}
