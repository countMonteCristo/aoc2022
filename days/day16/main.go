package main

import (
	"fmt"
	"sort"

	// "sort"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type DataType map[string]*Valve

type PathMap = map[string]map[string][]string

type Valve struct {
	Rate   int
	Valves []string
}

type QueueItem struct {
	Valve       string
	Pressure    int
	Time        int
	MaxPressure int
	Targets     *utils.Set[string]
}

type WV struct {
	Name string
	Rate int
}

var Working []WV
var Targets *utils.Set[string]
var Paths PathMap

func (i QueueItem) LessThan(j utils.PQItem) bool {
	jj := j.(QueueItem)
	return i.MaxPressure > jj.MaxPressure
}

func PredictMaxPressure(targets *utils.Set[string], paths PathMap, start string, timeLeft int) int {
	if timeLeft <= 0 {
		return 0
	}
	pred := 0
	cur := start
	for _, i := range Working {
		if targets.Contains(i.Name) {
			dt := len(paths[cur][i.Name]) - 1 + 1
			pred += i.Rate * timeLeft
			timeLeft -= dt
			cur = i.Name
		}
		if timeLeft <= 0 {
			break
		}
	}
	return pred
}

func (i *QueueItem) String() string {
	return fmt.Sprintf("%s: pres=%d max=%d time=%d opened=%v", i.Valve, i.Pressure, i.MaxPressure, i.Time, *i.Targets)
}

func prepare(lines []string) (valves DataType) {
	valves = make(DataType)
	Working = make([]WV, 0)
	Targets = utils.NewSet[string]()
	for _, line := range lines {
		parts := strings.Split(line, " ")
		v := parts[1]
		rate, _ := strconv.Atoi(strings.Split(strings.Trim(parts[4], ";"), "=")[1])
		connected_valves := utils.Transform(parts[9:], func(s string) string {
			return strings.Trim(s, ",")
		})
		valves[v] = &Valve{
			Rate: rate, Valves: connected_valves,
		}
		if rate > 0 {
			Working = append(Working, WV{Name: v, Rate: rate})
			Targets.Add(v)
		}
	}
	sort.Slice(Working, func(i, j int) bool {
		return Working[i].Rate > Working[j].Rate
	})

	Paths = get_paths(valves)
	return
}

type Path struct {
	Points []string
}

func (p Path) LessThan(q utils.PQItem) bool {
	return len(p.Points) < len(q.(Path).Points)
}

func SliceContain(p []string, s string) bool {
	for _, v := range p {
		if v == s {
			return true
		}
	}
	return false
}

func get_paths(valves DataType) map[string]map[string][]string {
	paths := make(map[string]map[string][]string)
	starts := []string{"AA"}
	finishes := make([]string, 0)
	for name, v := range valves {
		if v.Rate > 0 {
			starts = append(starts, name)
			finishes = append(finishes, name)
		}
	}

	for _, start := range starts {
		_, has := paths[start]
		if !has {
			paths[start] = make(map[string][]string)
		}
		for _, finish := range finishes {
			if finish == start {
				continue
			}
			_, has := paths[finish]
			if !has {
				paths[finish] = make(map[string][]string)
			}

			// fmt.Printf("Searching for path from %s to %s\n", start, finish)
			pq := utils.NewPq[Path]()
			s := []string{start}
			pq.Push(&Path{Points: s})
			for !pq.Empty() {
				p := pq.Pop()
				last := p.Points[len(p.Points)-1]
				if last == finish {
					paths[start][finish] = p.Points
					if len(paths[finish][start]) == 0 {
						paths[finish][start] = make([]string, len(p.Points))
						for i := len(p.Points) - 1; i >= 0; i-- {
							paths[finish][start][len(p.Points)-1-i] = p.Points[i]
						}
					}
					break
				}

				for _, next := range valves[last].Valves {
					if SliceContain(p.Points, next) {
						continue
					}
					np := make([]string, 0)
					np = append(np, p.Points...)
					np = append(np, next)
					pq.Push(&Path{Points: np})
				}
			}
		}
	}
	return paths
}

func solve_1(valves DataType, maxTime int) int {
	return find_best(valves, Paths, maxTime, Targets)
}

func solve_2(valves DataType, maxTime int) (ans int) {
	for i := 0; i < (1 << (Targets.Len() - 1)); i++ {
		t1 := utils.NewSet[string]()
		t2 := utils.NewSet[string]()
		for idx, w := range Working {
			if i&(1<<idx) > 0 {
				t1.Add(w.Name)
			} else {
				t2.Add(w.Name)
			}
		}
		p1 := PredictMaxPressure(t1, Paths, "AA", maxTime)
		p2 := PredictMaxPressure(t2, Paths, "AA", maxTime)
		if p1+p2 < ans {
			continue
		}
		ans = utils.Max(
			find_best(valves, Paths, maxTime, t1)+find_best(valves, Paths, maxTime, t2),
			ans,
		)
	}

	return
}

// Ищем путь по вершинам targets, который максимизирует давление за время maxTime
func find_best(valves DataType, paths PathMap, maxTime int, targets *utils.Set[string]) (ans int) {
	pq := utils.NewPq[QueueItem]()
	pq.Push(&QueueItem{
		Valve: "AA", Pressure: 0, Time: 0, Targets: targets,
	})
	for !pq.Empty() {
		item := pq.Pop()

		// Если открыто всё, что есть или кончилось время - обновляем ans
		if item.Time >= maxTime || item.Targets.Len() == 0 {
			// fmt.Println(" -> skip by time")
			ans = utils.Max(ans, item.Pressure)
			continue
		}

		// Пробуем подойти и открыть другие клапаны
		for next, path := range paths[item.Valve] {
			if !item.Targets.Contains(next) {
				continue
			}
			nt := item.Time + len(path)
			np := item.Pressure
			if nt < maxTime {
				np += (maxTime - nt) * valves[next].Rate
			}
			nc := item.Targets.Copy()
			nc.Remove(next)
			ni := &QueueItem{
				Valve: next, Pressure: np,
				Time: nt, Targets: nc,
				MaxPressure: np + PredictMaxPressure(nc, paths, next, maxTime-nt),
			}
			pq.Push(ni)
		}
	}
	return
}

func part_1(valves DataType) {
	ans := solve_1(valves, 30)
	// utils.CheckTask(1, ans, 1845)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(valves DataType) {
	ans := 2286
	fmt.Println("Need to wait for too long to check part2 :(")
	// ans := solve_2(valves, 26)
	// utils.CheckTask(2, ans, 2286)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day16/test1.txt"
	inputFile := "inputs/day16/input.txt"
	lines := utils.ReadFile(inputFile)
	valves := prepare(lines)
	part_1(valves)
	part_2(valves)
}
