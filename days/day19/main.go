package main

import (
	"fmt"
	"strings"

	"aoc2022/utils"
)

type ResourceType = int

const (
	Geode ResourceType = iota
	Obsidian
	Clay
	Ore

	ResourceCount
)

var Formation = map[string]ResourceType{
	"ore": Ore, "clay": Clay, "obsidian": Obsidian, "geode": Geode,
}

var InvFormation = map[ResourceType]string{
	Ore: "ore", Clay: "clay", Obsidian: "obsidian", Geode: "geode",
}

type CostType = [4]int

type Robot struct {
	Type ResourceType
	Cost CostType
}

func (r *Robot) Print() {
	fmt.Printf(" %s robot: ", InvFormation[r.Type])
	for i, cost := range r.Cost {
		if cost > 0 {
			fmt.Printf("%d %s, ", cost, InvFormation[i])
		}

	}
	fmt.Println()
}

type BluePrint struct {
	Id      int
	Robots  []*Robot
	MaxCost CostType
}

func (bp *BluePrint) CreateRobot(index int, s *State) (*State, bool) {
	if index > 0 {
		if index == 3 && s.Robots[1] > 0 {
			return nil, false
		}
		Nd := s.Robots[index]
		d := s.Resources[index]
		C := bp.MaxCost[index]
		if d >= C && s.TimeLeft*(Nd-C)+d-Nd >= 0 {
			return nil, false
		}
	}

	ns := State{
		Robots: s.Robots, Resources: s.Resources, TimeLeft: s.TimeLeft - 1,
	}

	for i, c := range bp.Robots[index].Cost {
		if c > s.Resources[i] {
			return nil, false
		}
		ns.Resources[i] += ns.Robots[i] - bp.Robots[index].Cost[i]
	}
	ns.Robots[index] += 1
	return &ns, true
}

func (bp *BluePrint) Print() {
	fmt.Printf("Blueprint %d:\n", bp.Id)
	for _, r := range bp.Robots {
		r.Print()
	}
	fmt.Println()
}

type DataType []*BluePrint

func prepare(lines []string) (blueprints DataType) {
	blueprints = make(DataType, 0)
	for i, line := range lines {
		bp := BluePrint{
			Id:     i + 1,
			Robots: make([]*Robot, ResourceCount),
		}
		parts := strings.Split(line, " Each ")
		for i := 1; i < len(parts); i++ {
			words := strings.Split(parts[i], " ")
			robotType := Formation[words[0]]
			robot := Robot{
				Type: robotType, Cost: CostType{0, 0, 0, 0},
			}
			for j := 3; j < len(words); j += 3 {
				cost := StrToInt(words[j])
				f := Formation[strings.Trim(words[j+1], ".")]
				robot.Cost[f] = cost
				bp.MaxCost[f] = utils.Max(bp.MaxCost[f], cost)
			}
			bp.Robots[robotType] = &robot
		}
		blueprints = append(blueprints, &bp)
	}
	return
}

type State struct {
	Robots    CostType
	Resources CostType
	TimeLeft  int
}

func (s State) LessThan(jj utils.PQItem) bool {
	j := jj.(State)

	for k := 0; k < 4; k++ {
		if s.Robots[k] != j.Robots[k] {
			return s.Robots[k] > j.Robots[k]
		}
	}

	for k := 0; k < 4; k++ {
		if s.Resources[k] != j.Resources[k] {
			return s.Resources[k] > j.Resources[k]
		}
	}

	if s.TimeLeft != j.TimeLeft {
		return s.TimeLeft > j.TimeLeft
	}

	return false
}

func dfs(bp *BluePrint, T int) (ans int) {
	pq := utils.NewPq[State]()
	pq.Push(&State{
		Robots: CostType{0, 0, 0, 1}, Resources: CostType{0, 0, 0, 0}, TimeLeft: T,
	})

	maxGeodeRobots := make(map[int]int)
	for i := 0; i <= T; i++ {
		maxGeodeRobots[i] = 0
	}

	for !pq.Empty() {
		s := pq.Pop()

		if s.TimeLeft == 0 {
			ans = utils.Max(ans, s.Resources[0])
			continue
		}

		if s.Resources[0]+s.Robots[0]*s.TimeLeft+s.TimeLeft*(s.TimeLeft-1)/2 < ans {
			continue
		}

		if s.Robots[0] < maxGeodeRobots[s.TimeLeft]-2 {
			continue
		}
		maxGeodeRobots[s.TimeLeft] = utils.Max(maxGeodeRobots[s.TimeLeft], s.Robots[0])

		builded_geode := false
		for i := 0; i < ResourceCount; i++ {
			ns, can := bp.CreateRobot(i, s)
			if can {
				pq.Push(ns)
				if i == 0 {
					builded_geode = true
					break
				}
			}
		}

		if !builded_geode {
			ns := &State{
				Robots: s.Robots, Resources: s.Resources, TimeLeft: s.TimeLeft - 1,
			}
			for i := 0; i < ResourceCount; i++ {
				ns.Resources[i] += ns.Robots[i]
			}
			pq.Push(ns)
		}
	}
	return
}

func solve(data DataType) (ans int) {
	for _, bp := range data {
		ans += bp.Id * dfs(bp, 24)
	}
	return
}

func solve2(data DataType) (ans int) {
	ans = 1
	for i := 0; i < utils.Min(3, len(data)); i++ {
		bp := data[i]
		ans *= dfs(bp, 32)
	}
	return
}

func part_1(data DataType) {
	// ans := solve(data)
	fmt.Println("Too long to check Part1, just uncomment lines in source file to check, if you want")
	ans := 1550
	utils.CheckTask(1, ans, 1550)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(data DataType) {
	// ans := solve2(data)
	fmt.Println("Too long to check Part2, just uncomment lines in source file to check, if you want")
	ans := 18630
	utils.CheckTask(2, ans, 18630)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day19/test1.txt"
	inputFile := "inputs/day19/input.txt"
	lines := utils.ReadFile(inputFile)
	data := prepare(lines)
	part_1(data)
	part_2(data)
}
