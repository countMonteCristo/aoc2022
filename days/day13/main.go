package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"aoc2022/utils"
)

type ItemType int

const (
	TypeOpen ItemType = iota
	TypeClosed
	TypeValue
)

type Item struct {
	Type  ItemType
	Value int
}

func (i *Item) String() (s string) {
	switch i.Type {
	case TypeOpen:
		s += "["
	case TypeClosed:
		s += "]"
	case TypeValue:
		s += strconv.Itoa(i.Value)
	}
	return
}

type Packet []*Item
type PacketPairs [][]Packet

func (p Packet) String() (s string) {
	for index, i := range p {
		s += fmt.Sprintf("%v", i)
		if i.Type != TypeOpen {
			s += ","
		}
		if i.Type == TypeClosed && p[index-1].Type != TypeOpen {
			s = s[:len(s)-3] + "],"
		}
	}
	return s[:len(s)-1]
}

func parsePacket(line string) Packet {
	array := make(Packet, 0)
	num_str := ""
	for _, c := range line {
		switch {
		case c == '[':
			array = append(array, &Item{Type: TypeOpen})
		case c == ']' || c == ',':
			if len(num_str) > 0 {
				v, _ := strconv.Atoi(num_str)
				array = append(array, &Item{Type: TypeValue, Value: v})
				num_str = ""
			}
			if c == ']' {
				array = append(array, &Item{Type: TypeClosed})
			}
		default:
			num_str += string(c)
		}
	}
	return array
}

func prepare(lines []string) (groups PacketPairs) {
	packets := make([]Packet, 0)
	n := 0
	for _, line := range lines {
		if len(line) == 0 {
			groups = append(groups, packets)
			packets = make([]Packet, 0)
			n = 0
		} else {
			packets = append(packets, parsePacket(line))
			n++
		}
	}
	if n > 0 {
		groups = append(groups, packets)
	}
	return
}

type IntPair utils.Pair[int]

func findRange(packet Packet, start int) IntPair {
	if packet[start].Type == TypeValue {
		return IntPair{First: start, Second: start}
	}
	count := 1
	for i := start + 1; i < len(packet); i++ {
		if packet[i].Type == TypeOpen {
			count++
		}
		if packet[i].Type == TypeClosed {
			count--
		}
		if count == 0 {
			return IntPair{First: start + 1, Second: i - 1}
		}
	}
	panic("Could not find range")
}

func updateIndex(index int, typ ItemType) int {
	index += 1
	if typ != TypeValue {
		index += 1
	}
	return index
}

func compare(p1, p2 Packet, f, s IntPair) int {
	index_1, index_2 := f.First, s.First
	for index_1 <= f.Second && index_2 <= s.Second {
		i1, i2 := p1[index_1], p2[index_2]
		r1, r2 := findRange(p1, index_1), findRange(p2, index_2)

		v := utils.Sign(i1.Value - i2.Value)
		if i1.Type != TypeValue || i2.Type != TypeValue {
			v = compare(p1, p2, r1, r2)
		}
		if v != 0 {
			return v
		}

		index_1, index_2 = updateIndex(r1.Second, i1.Type), updateIndex(r2.Second, i2.Type)
	}
	if index_1 > f.Second && index_2 <= s.Second {
		return -1
	}
	if index_2 > s.Second && index_1 <= f.Second {
		return 1
	}
	return 0
}

func comparePair(p1, p2 Packet) int {
	r1 := IntPair{First: 1, Second: len(p1) - 2}
	r2 := IntPair{First: 1, Second: len(p2) - 2}
	return compare(p1, p2, r1, r2)
}

func solve_1(groups PacketPairs) (ans int) {
	for i, packets := range groups {
		if v := comparePair(packets[0], packets[1]); v == -1 {
			ans += (i + 1)
		}
	}
	return
}

func solve_2(groups PacketPairs) (ans int) {
	packets := make([]Packet, 0)
	for _, group := range groups {
		packets = append(packets, group...)
	}
	p1, p2 := parsePacket("[[2]]"), parsePacket("[[6]]")
	packets = append(packets, p1, p2)

	sort.Slice(packets, func(i, j int) bool {
		return comparePair(packets[i], packets[j]) == -1
	})

	ans = 1
	for i, p := range packets {
		if reflect.DeepEqual(p, p1) || reflect.DeepEqual(p, p2) {
			ans *= (i + 1)
		}
	}
	return
}

func part_1(pairs PacketPairs) {
	ans := solve_1(pairs)
	utils.CheckTask(1, ans, 6046)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(pairs PacketPairs) {
	ans := solve_2(pairs)
	utils.CheckTask(2, ans, 21423)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day13/test1.txt"
	inputFile := "inputs/day13/input.txt"
	lines := utils.ReadFile(inputFile)
	pairs := prepare(lines)
	part_1(pairs)
	part_2(pairs)
}
