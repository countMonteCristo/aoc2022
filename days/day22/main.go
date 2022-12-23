package main

import (
	"fmt"
	"strconv"

	"aoc2022/utils"
)

type IntPoint = utils.Point2d[int]

var Zero = IntPoint{X: 0, Y: 0}
var Zero3d = Point3d{X: 0, Y: 0, Z: 0}

var DD = []IntPoint{
	{X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 0, Y: -1},
}

var Rotations = map[byte]int{
	'R': 1, 'L': -1,
}

type FacingType = int

const (
	Right FacingType = iota
	Down
	Left
	Up

	FacingCount
)

type Tunnel struct {
	Point  IntPoint
	Facing FacingType
}

// ---------------------------------------------------------------------------------------------------------------------

type CommandType = int

const (
	Go CommandType = iota
	Turn
)

type Command struct {
	Type     CommandType
	Steps    int
	Rotation int
}

func (c *Command) String() string {
	return fmt.Sprintf("(%d %d %d)", c.Type, c.Steps, c.Rotation)
}

// ---------------------------------------------------------------------------------------------------------------------

type Board struct {
	W, H                int
	Cells               [][]byte
	Tunnels             map[Tunnel]IntPoint
	Start               IntPoint
	CubeSize            int
	PointToTunnelFacing map[IntPoint]FacingType

	Commands []Command
}

func (b *Board) Wrap(p IntPoint) IntPoint {
	return IntPoint{
		X: (p.X + b.W) % b.W,
		Y: (p.Y + b.H) % b.H,
	}
}

func (b *Board) Contains(p *IntPoint) bool {
	if p.X < 0 || p.X >= b.W || p.Y < 0 || p.Y >= b.H {
		return false
	}
	return b.Cells[p.Y][p.X] != 0
}

func (b *Board) GetWrapped(p, dp *IntPoint) IntPoint {
	np := b.Wrap(*p)
	for !b.Contains(&np) {
		np = b.Wrap(np.Plus(dp))
	}
	return np
}

func (b *Board) Connected(p, q IntPoint, size int) (y0, a Point3d, connected bool) {
	d := p.Minus(&q)
	if utils.Manhattan(d, Zero) == size {
		connected = true
		y0 = Point3d{X: (p.X + q.X) / 2, Y: (p.Y + q.Y) / 2, Z: 0}
		a = Point3d{X: utils.Sign(-d.Y), Y: utils.Sign(d.X), Z: 0}
	}
	return
}

func (b *Board) GetFaces() []*CubeFace {
	faces := make([]*CubeFace, 0, 6)

	ids := make([][]int, b.H)
	for i := 0; i < b.H; i++ {
		ids[i] = make([]int, b.W)
	}

	face_id := 1
	for i := 0; i < b.H; i++ {
		for j := 0; j < b.W; j++ {
			if b.Cells[i][j] != 0 && ids[i][j] == 0 {
				for di := i; di < i+b.CubeSize; di++ {
					for dj := j; dj < j+b.CubeSize; dj++ {
						ids[di][dj] = face_id
					}
				}

				face := &CubeFace{
					Center2d:   IntPoint{X: j + b.CubeSize/2, Y: i + b.CubeSize/2},
					Corner:     make(map[IntPoint]Point3d),
					RealCorner: make(map[IntPoint]IntPoint),
				}
				face.Center = FromIntPoint(face.Center2d)

				p1 := IntPoint{X: j, Y: i}
				p2 := p1
				for i := 0; i < FacingCount; i++ {
					d := DD[i]
					face.Corner[p1] = FromIntPoint(p1)
					face.RealCorner[p1] = p2
					d1, d2 := d.Prod(b.CubeSize), d.Prod(b.CubeSize-1)
					p1.Add(&d1)
					p2.Add(&d2)
				}

				faces = append(faces, face)
				face_id++
			}
		}
	}

	return faces
}

func (b *Board) GetTunnelsForEdge(p1, p2 IntPoint) []Tunnel {
	tunnels := make([]Tunnel, 0)
	dp := IntPoint{
		X: utils.Sign(p2.X - p1.X), Y: utils.Sign(p2.Y - p1.Y),
	}
	facings := make(map[FacingType]int)
	p := p1
	for p != p2 {
		for f, dd := range DD {
			q := p.Plus(&dd)
			if !b.Contains(&q) {
				facings[f]++
			}
		}
		p.Add(&dp)
	}

	maxv := 0
	maxf := Up
	for f, v := range facings {
		if v > maxv {
			maxv = v
			maxf = f
		}
	}

	p = p1
	for p != p2 {
		t := Tunnel{Point: p, Facing: maxf}
		tunnels = append(tunnels, t)
		p.Add(&dp)
	}
	tunnels = append(tunnels, Tunnel{Point: p2, Facing: maxf})

	return tunnels
}

func (b *Board) Fold(faces []*CubeFace) {
	connections := make(map[int]*utils.Set[int])
	rotated := make(map[int]*utils.Set[int])

	for i := range faces {
		rotated[i] = utils.NewSet[int]()
		connections[i] = utils.NewSet[int]()
	}

	for i, f1 := range faces {
		for j := i + 1; j < len(faces); j++ {
			f2 := faces[j]
			_, _, connected := b.Connected(f1.Center2d, f2.Center2d, b.CubeSize)
			if connected {
				connections[i].Add(j)
				connections[j].Add(i)
			}
		}
	}

	for {
		n_rots := 0
		for fid, nbrs := range connections {
			if nbrs.Len() == 1 {
				nxt := nbrs.Pop()
				delete(connections, fid)
				y0, a, _ := b.Connected(faces[fid].Center2d, faces[nxt].Center2d, b.CubeSize)
				rotated[nxt].Add(fid)
				rotated[nxt].Update(rotated[fid])
				connections[nxt].Remove(fid)

				faces[fid].Rotate(y0, a)
				for face_id := range rotated[fid].Iter() {
					faces[face_id].Rotate(y0, a)
				}
				n_rots++
			}
		}
		if n_rots == 0 {
			break
		}
	}
}

func (b *Board) CreateTunnelsFromFolded(faces []*CubeFace) {
	for i, face := range faces {
		for j := i + 1; j < len(faces); j++ {
			nface := faces[j]

			edge := face.GetCommonEdge(nface)
			if len(edge) == 2 {
				pair1, pair2 := edge[0], edge[1]
				if pair1[0] == pair1[1] && pair2[0] == pair2[1] { // inner edge
					continue
				}
				e1_tunnels := b.GetTunnelsForEdge(face.RealCorner[pair1[0]], face.RealCorner[pair2[0]])
				e2_tunnels := b.GetTunnelsForEdge(nface.RealCorner[pair1[1]], nface.RealCorner[pair2[1]])
				for i := 0; i < len(e1_tunnels); i++ {
					t1, t2 := e1_tunnels[i], e2_tunnels[i]
					b.Tunnels[t1] = t2.Point
					b.Tunnels[t2] = t1.Point
					b.PointToTunnelFacing[t1.Point] = t1.Facing
					b.PointToTunnelFacing[t2.Point] = t2.Facing
				}
			}
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

type Point3d struct {
	X, Y, Z int
}

func FromIntPoint(p IntPoint) Point3d {
	return Point3d{X: p.X, Y: p.Y, Z: 0}
}

func (p *Point3d) Plus(q Point3d) Point3d {
	return Point3d{
		X: p.X + q.X, Y: p.Y + q.Y, Z: p.Z + q.Z,
	}
}
func (p *Point3d) Minus(q Point3d) Point3d {
	return Point3d{
		X: p.X - q.X, Y: p.Y - q.Y, Z: p.Z - q.Z,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

type CubeFace struct {
	Center     Point3d
	Center2d   IntPoint
	Corner     map[IntPoint]Point3d
	RealCorner map[IntPoint]IntPoint
}

func (f *CubeFace) GetCommonEdge(o *CubeFace) (edge [][2]IntPoint) {
	edge = make([][2]IntPoint, 0)
	for c1 := range f.Corner {
		for c2 := range o.Corner {
			if f.Corner[c1] == o.Corner[c2] {
				edge = append(edge, [2]IntPoint{c1, c2})
			}
		}
	}
	return
}

func (c *CubeFace) Rotate(y0, a Point3d) {
	m := GetMatrix(a, 1)
	nc := y0.Plus(m.Mul(c.Center.Minus(y0)))
	if nc.Z < 0 {
		m = GetMatrix(a, -1)
		nc = y0.Plus(m.Mul(c.Center.Minus(y0)))
	}

	c.Center = nc

	new_corners := make(map[IntPoint]Point3d)
	for p2, p3 := range c.Corner {
		np3 := y0.Plus(m.Mul(p3.Minus(y0)))
		new_corners[p2] = np3
	}
	c.Corner = new_corners
}

// ---------------------------------------------------------------------------------------------------------------------

type Matrix [3][3]int

func GetMatrix(a Point3d, s int) Matrix {
	x, y, z := a.X, a.Y, a.Z
	return Matrix{
		{x * x, x*y - s*z, x*z + s*y},
		{y*x + s*z, y * y, y*z - s*x},
		{z*x - s*y, z*y + s*x, z * z},
	}
}

func (m Matrix) Mul(p Point3d) Point3d {
	return Point3d{
		X: m[0][0]*p.X + m[0][1]*p.Y + m[0][2]*p.Z,
		Y: m[1][0]*p.X + m[1][1]*p.Y + m[1][2]*p.Z,
		Z: m[2][0]*p.X + m[2][1]*p.Y + m[2][2]*p.Z,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func prepare(lines []string, part1 bool) (board Board) {
	board.H = len(lines) - 2
	for i := 0; i < board.H; i++ {
		board.W = utils.Max(board.W, len(lines[i]))
	}

	board.Cells = make([][]byte, board.H)
	for i := 0; i < board.H; i++ {
		board.Cells[i] = make([]byte, board.W)
		for j := 0; j < board.W; j++ {
			if j >= len(lines[i]) {
				continue
			}
			if lines[i][j] == '.' || lines[i][j] == '#' {
				board.Cells[i][j] = lines[i][j]
			}
		}
	}

	for j, c := range board.Cells[0] {
		if c == '.' {
			board.Start = IntPoint{X: j, Y: 0}
			break
		}
	}

	board.Tunnels = make(map[Tunnel]IntPoint)
	if part1 {
		for i := 0; i < board.H; i++ {
			for j := 0; j < board.W; j++ {
				if board.Cells[i][j] == 0 {
					continue
				}
				p := IntPoint{X: j, Y: i}
				for facing, dp := range DD {
					np := p.Plus(&dp)
					if board.Contains(&np) {
						continue
					}
					tp := board.GetWrapped(&np, &dp)
					board.Tunnels[Tunnel{Point: p, Facing: facing}] = tp
				}
			}
		}
	} else {
		board.CubeSize = board.W
		for _, line := range board.Cells {
			begin, end := 0, 0
			for x := 0; x < len(line); x++ {
				if line[x] != 0 {
					begin = x
					break
				}
			}
			for x := len(line) - 1; x >= 0; x-- {
				if line[x] != 0 {
					end = x
					break
				}
			}
			board.CubeSize = utils.Min(board.CubeSize, end-begin+1)
		}

		board.PointToTunnelFacing = make(map[IntPoint]int)
		faces := board.GetFaces()
		board.Fold(faces)
		board.CreateTunnelsFromFolded(faces)
	}

	board.Commands = make([]Command, 0)
	path := lines[len(lines)-1]
	start_idx := 0
	for i := 0; i < len(path); i++ {
		if path[i] == 'R' || path[i] == 'L' {
			if i > start_idx {
				n, err := strconv.Atoi(path[start_idx:i])
				if err != nil {
					panic(fmt.Sprintf("Cannot convert %s to int", path[start_idx:i]))
				}
				board.Commands = append(board.Commands, Command{Type: Go, Steps: n})
			}
			board.Commands = append(board.Commands, Command{Type: Turn, Rotation: Rotations[path[i]]})
			start_idx = i + 1
		}
	}
	if start_idx < len(path) {
		n, _ := strconv.Atoi(path[start_idx:])
		board.Commands = append(board.Commands, Command{Type: Go, Steps: n})
	}

	return
}

// ---------------------------------------------------------------------------------------------------------------------

func solve(board Board, part2 bool) (ans int) {
	pos := board.Start
	face := Right
	for _, cmd := range board.Commands {
		switch cmd.Type {
		case Go:
			for i := 0; i < cmd.Steps; i++ {
				dp := DD[face]
				np := pos.Plus(&dp)
				if board.Contains(&np) {
					if board.Cells[np.Y][np.X] == '#' {
						break
					}
				} else {
					tunnel := Tunnel{Point: pos, Facing: face}
					tp := board.Tunnels[tunnel]
					if board.Cells[tp.Y][tp.X] == '#' {
						break
					}
					np = tp
					if part2 {
						face = (board.PointToTunnelFacing[np] + FacingCount/2) % FacingCount
					}
				}
				pos = np
			}
		case Turn:
			face = (face + cmd.Rotation + FacingCount) % FacingCount
		default:
			panic(fmt.Sprintf("Unknown command type: %d", cmd.Type))
		}
	}
	return (pos.X+1)*4 + (pos.Y+1)*1000 + face
}

// ---------------------------------------------------------------------------------------------------------------------

func part_1(lines []string) {
	board := prepare(lines, true)
	ans := solve(board, false)
	utils.CheckTask(1, ans, 13566)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(lines []string) {
	board := prepare(lines, false)
	ans := solve(board, true)
	utils.CheckTask(2, ans, 11451)
	fmt.Println("[Part 2] Answer:", ans)
}

func main() {
	// inputFile := "inputs/day22/test1.txt"
	inputFile := "inputs/day22/input.txt"
	lines := utils.ReadFile(inputFile)
	part_1(lines)
	part_2(lines)
}
