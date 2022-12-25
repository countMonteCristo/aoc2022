package main

import (
	"fmt"
	"log"
	"strings"

	"aoc2022/utils"
)

// ------------------------------------------------------------------------------------------------

type FS struct {
	Root *FsItem
}

func NewFS() *FS {
	return &FS{
		Root: NewFsItem("/", TypeDir, 0, nil),
	}
}

func (fs *FS) GetSize() int {
	return fs.Root.Size
}

const (
	TypeDir = iota
	TypeFile
)

func TypeToStr(typ int) string {
	switch typ {
	case TypeDir:
		return "dir"
	case TypeFile:
		return "file"
	default:
		log.Fatal("Unknown FsItem type: ", typ)
	}
	return "UNKNOWN_FS_TYPE"
}

type FsItem struct {
	Name     string
	Type     int
	Children map[string]*FsItem
	Parent   *FsItem
	Size     int
}

func NewFsItem(name string, typ, size int, parent *FsItem) *FsItem {
	return &FsItem{
		Name: name, Parent: parent, Type: typ, Size: size,
		Children: make(map[string]*FsItem),
	}
}

func (i *FsItem) GetTypeStr() string {
	return TypeToStr(i.Type)
}

func (i *FsItem) Repr(indent string) {
	fmt.Printf("%s- %s (%s, size=%d)\n", indent, i.Name, i.GetTypeStr(), i.Size)
}

func (i *FsItem) Print(indent string) {
	i.Repr(indent)
	for _, child := range i.Children {
		child.Print(indent + "  ")
	}
}

func (i *FsItem) UpdateSize(size int) {
	i.Size += size
	if i.Parent != nil && i.Parent.Type == TypeDir {
		i.Parent.UpdateSize(size)
	}
}

// ------------------------------------------------------------------------------------------------

func Tree(i *FsItem) {
	i.Print("")
}

// ------------------------------------------------------------------------------------------------

func PrepareFS(lines []string) *FS {
	fs := NewFS()
	current := fs.Root
	for _, line := range lines {
		switch parts := strings.Split(line, " "); parts[0] {
		case "$":
			if parts[1] == "cd" {
				dirName := parts[2]
				if dirName == ".." {
					current = current.Parent
				} else if dirName != fs.Root.Name {
					current = current.Children[dirName]
				}
			}	// if `ls` - do nothing
		case TypeToStr(TypeDir):
			dir := NewFsItem(parts[1], TypeDir, 0, current)
			current.Children[dir.Name] = dir
		default:
			size := StrToInt(parts[0])
			file := NewFsItem(parts[1], TypeFile, size, current)
			current.Children[file.Name] = file
			current.UpdateSize(file.Size)
		}
	}
	return fs
}

func sumDirSizesBelow(i *FsItem, maxSize int) (ans int) {
	if i.Size <= maxSize {
		ans += i.Size
	}
	for _, child := range i.Children {
		if child.Type == TypeDir {
			ans += sumDirSizesBelow(child, maxSize)
		}
	}
	return
}

func getClosestAbove(i *FsItem, minSize, currentSize int) int {
	if i.Size < minSize || i.Size > currentSize {
		return currentSize
	}
	curMin := i.Size
	for _, d := range i.Children {
		if d.Type == TypeDir {
			curMin = utils.Min(curMin, getClosestAbove(d, minSize, i.Size))
		}
	}
	return curMin
}

// ------------------------------------------------------------------------------------------------

func part_1(fs *FS) {
	ans := sumDirSizesBelow(fs.Root, 100000)
	utils.CheckTask(1, ans, 1325919)
	fmt.Println("[Part 1] Answer:", ans)
}

func part_2(fs *FS) {
	updateSize := 30000000
	fsSize := 70000000
	ans := getClosestAbove(fs.Root, updateSize-(fsSize-fs.GetSize()), fs.GetSize())
	utils.CheckTask(2, ans, 2050735)
	fmt.Println("[Part 2] Answer:", ans)
}

// ------------------------------------------------------------------------------------------------

func main() {
	// inputFile := "inputs/day07/test1.txt"
	inputFile := "inputs/day07/input.txt"
	lines := utils.ReadFile(inputFile)
	fs := PrepareFS(lines)
	// Tree(fs.Root)
	part_1(fs)
	part_2(fs)
}
