package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.Remove(9)
	fmt.Println(x.String())
	fmt.Println(x.Len())
	x.Clear()
	fmt.Println(x.String())
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	t := x.Copy()
	fmt.Println(t.String())
	t.Remove(144)
	fmt.Println(t.String())
	fmt.Println(x.String())
}

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, x%64
	return word < len(s.words) && s.words[word]&(1 << uint(bit)) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, x%64
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << uint(bit)
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1 << uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, x%64
	if word < len(s.words) {
		s.words[word] &^= (1 << uint(bit))
	}
}


func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var words []uint64
	for _, word := range s.words {
		words = append(words, word)
	}
	return &IntSet{words}
}
