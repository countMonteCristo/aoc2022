package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/utils"
)

type IntPoint = utils.Point2d[int64]
type Surface struct {
	Beacons    *utils.Set[IntPoint]
	SensorDist map[IntPoint]int64
}

func parseFromStr(s, trim string) (res int64) {
	res, _ = strconv.ParseInt(strings.Split(strings.Trim(s, trim), "=")[1], 10, 64)
	return
}

func prepare(lines []string) (surf *Surface) {
	surf = &Surface{
		Beacons:    utils.NewSet[IntPoint](),
		SensorDist: make(map[IntPoint]int64),
	}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		sensor := IntPoint{
			X: parseFromStr(parts[2], ","), Y: parseFromStr(parts[3], ":"),
		}
		beacon := IntPoint{
			X: parseFromStr(parts[8], ","), Y: parseFromStr(parts[9], ":"),
		}
		surf.Beacons.Add(beacon)
		surf.SensorDist[sensor] = utils.Manhattan(sensor, beacon)
	}

	return
}

func solve_1(surf *Surface, y0 int64) (ans int) {
	beaconsAtBusy := 0
	for b := range surf.Beacons.Iter() {
		if b.Y == y0 {
			beaconsAtBusy++
		}
	}

	busyX := utils.NewSet[int64]()
	for sensor, sendsor_dist := range surf.SensorDist {
		// s.X - (sendsor_dist - |s.Y- y0|) <= x <= s.X + (sendsor_dist - |s.Y - y0|)
		delta := sendsor_dist - utils.Abs(sensor.Y-y0)
		xl, xr := sensor.X-delta, sensor.X+delta
		if xr >= xl {
			for x := xl; x <= xr; x++ {
				busyX.Add(x)
			}
		}
	}
	ans = busyX.Len() - beaconsAtBusy
	return
}

func solve_2(surf *Surface, vmin, vmax int64) (ans int64) {
	sensors := make([]IntPoint, 0)
	for s := range surf.SensorDist {
		sensors = append(sensors, s)
	}
	sort.Slice(sensors, func(i, j int) bool {
		return surf.SensorDist[sensors[i]] < surf.SensorDist[sensors[j]]
	})

	dd := []IntPoint{
		{X: 1, Y: 1}, {X: -1, Y: 1}, {X: -1, Y: -1}, {X: 1, Y: -1},
	}

	for _, sen := range sensors {
		rad := surf.SensorDist[sen] + 1
		cur := IntPoint{X: sen.X, Y: sen.Y - rad}
		found := false
		for i := int64(0); i < 4*rad; i++ {
			dp := dd[i/rad]
			if cur.X < vmin || cur.X > vmax || cur.Y < vmin || cur.Y > vmax {
				cur.Add(dp)
				continue
			}
			ok := true
			for other_sensor, other_dist := range surf.SensorDist {
				if utils.Manhattan(cur, other_sensor) <= other_dist {
					ok = false
					break
				}
			}
			if ok {
				found = true
				break
			}
			cur.Add(dp)
		}
		if found {
			ans = 4000000*cur.X + cur.Y
			break
		}
	}
	return
}

func part_1(surf *Surface) {
	// ans := solve(surf, 10)	// test example
	ans := solve_1(surf, 2000000)
	utils.CheckTask(1, ans, 5083287)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(surf *Surface) {
	// ans := solve2(surf, 0, 20)	// test example
	ans := solve_2(surf, 0, 4000000)
	utils.CheckTask(2, ans, 13134039205729)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day15/test1.txt"
	inputFile := "inputs/day15/input.txt"
	lines := utils.ReadFile(inputFile)
	surf := prepare(lines)
	part_1(surf)
	part_2(surf)
}
