package main

import (
	"fmt"
	"strings"

	"aoc2022/utils"
)

type MaskType = [][]uint8

type Figure struct {
	Width, Height int64
	Mask          MaskType
	Rocks         I64pSet
}

func (f *Figure) Init() {
	f.Height, f.Width = int64(len(f.Mask)), int64(len(f.Mask[0]))
	f.Rocks = NewI64pSet()
	for x := int64(0); x < f.Width; x++ {
		for y := int64(0); y < f.Height; y++ {
			if f.Mask[y][x] == 1 {
				f.Rocks.Add(I64Point{X: x, Y: -y})
			}
		}
	}
}

func (f *Figure) CanMove(topLeft, d I64Point, room *Room, fail func(I64Point, int64) bool) bool {
	for r := range f.Rocks.Iter() {
		np := topLeft.Plus(r)
		np.Add(d)
		if fail(np, room.Width) || room.Rocks.Contains(np) {
			return false
		}
	}
	return true
}

func (f *Figure) CanShift(topLeft, d I64Point, room *Room) bool {
	return f.CanMove(topLeft, d, room, func(np I64Point, width int64) bool {
		return np.X < 0 || np.X >= width
	})
}

func (f *Figure) CanFall(topLeft, d I64Point, room *Room) bool {
	return f.CanMove(topLeft, d, room, func(np I64Point, width int64) bool {
		return np.Y < 0
	})
}

type CacheKey struct {
	FigureId int
	CmdIndex int
}

type Room struct {
	Rocks    I64pSet
	Width    int64
	Commands []int
	Figures  []*Figure
	Cache    map[CacheKey]int64
	Heights  map[int64]int64
}

func (room *Room) AddFigure(f *Figure, topLeft I64Point) {
	for r := range f.Rocks.Iter() {
		np := topLeft.Plus(r)
		room.Rocks.Add(np)
	}
}

func (room *Room) Clear() {
	room.Rocks.Clear()
	for k := range room.Cache {
		delete(room.Cache, k)
	}
	for k := range room.Heights {
		delete(room.Heights, k)
	}
}

func prepare(lines []string) (room *Room) {
	room = &Room{
		Rocks: NewI64pSet(),
		Commands: utils.Transform(strings.Split(lines[0], ""), func(c string) (idx int) {
			if c == ">" {
				idx = 1
			}
			return
		}),
		Figures: []*Figure{
			{Mask: MaskType{{1, 1, 1, 1}}},
			{Mask: MaskType{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}},
			{Mask: MaskType{{0, 0, 1}, {0, 0, 1}, {1, 1, 1}}},
			{Mask: MaskType{{1}, {1}, {1}, {1}}},
			{Mask: MaskType{{1, 1}, {1, 1}}},
		},
		Cache:   make(map[CacheKey]int64),
		Heights: make(map[int64]int64),
		Width:   int64(7),
	}

	for _, fig := range room.Figures {
		fig.Init()
	}
	return
}

func solve(room *Room, N int64) (lastHeight int64) {
	room.Clear()

	v_space, h_space := int64(3), int64(2)
	cmd_idx, fig_idx := 0, 0
	dp := []I64Point{
		{X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: -1},
	}

	for i := int64(0); i < N; i++ {
		key := CacheKey{FigureId: fig_idx, CmdIndex: cmd_idx}
		n_figures_prev, exists := room.Cache[key]

		if exists && i > 3000 {
			countLeft := N - n_figures_prev
			period := i - n_figures_prev
			count, rem := countLeft/period, countLeft%period
			lastHeight = room.Heights[n_figures_prev+rem] + (lastHeight-room.Heights[n_figures_prev])*count
			return
		} else {
			fig := room.Figures[fig_idx]
			topLeft := I64Point{X: h_space, Y: int64(lastHeight) - 1 + v_space + fig.Height}

			for {
				d := dp[room.Commands[cmd_idx]]
				if fig.CanShift(topLeft, d, room) {
					topLeft.Add(d)
				}
				cmd_idx = (cmd_idx + 1) % len(room.Commands)

				if fig.CanFall(topLeft, dp[2], room) {
					topLeft.Add(dp[2])
				} else {
					room.AddFigure(fig, topLeft)
					break
				}
			}
			fig_idx = (fig_idx + 1) % len(room.Figures)

			lastHeight = utils.Max(lastHeight, topLeft.Y+1)
			room.Heights[i+1] = lastHeight

			room.Cache[key] = i
		}
	}
	return
}

func part_1(room *Room) {
	ans := solve(room, 2022)
	utils.CheckTask(1, ans, 3157)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(room *Room) {
	ans := solve(room, 1000000000000)
	utils.CheckTask(2, ans, 1581449275319)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day17/test1.txt"
	inputFile := "inputs/day17/input.txt"
	lines := utils.ReadFile(inputFile)
	room := prepare(lines)
	part_1(room)
	part_2(room)
}
