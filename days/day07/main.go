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
	return fs.Root.GetSize()
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

func (i *FsItem) GetName() string {
	return i.Name
}

func (i *FsItem) GetSize() int {
	return i.Size
}

func (i *FsItem) GetType() int {
	return i.Type
}

func (i *FsItem) GetTypeStr() string {
	return TypeToStr(i.GetType())
}

func (i *FsItem) Repr(indent string) {
	fmt.Printf("%s- %s (%s, size=%d)\n", indent, i.GetName(), i.GetTypeStr(), i.GetSize())
}

func (i *FsItem) Print(indent string) {
	i.Repr(indent)
	for _, child := range i.Children {
		child.Print(indent + "  ")
	}
}

func (i *FsItem) UpdateSize(size int) {
	i.Size += size
	if i.Parent != nil && i.Parent.GetType() == TypeDir {
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
	for i, line := range lines {
		switch parts := strings.Split(line, " "); parts[0] {
		case "$":
			if parts[1] == "cd" {
				dirName := parts[2]
				if dirName == ".." {
					current = current.Parent
				} else if dirName != fs.Root.GetName() {
					next, exists := current.Children[dirName]
					if !exists {
						log.Fatal("[ERROR] Line ", i+1, ": No such dir: ", dirName)
					}
					current = next
				}

			} else { // `ls`, do nothing
				{
				}
			}
		case TypeToStr(TypeDir):
			dir := NewFsItem(parts[1], TypeDir, 0, current)
			current.Children[dir.GetName()] = dir
		default:
			size := StrToInt(parts[0])
			file := NewFsItem(parts[1], TypeFile, size, current)
			current.Children[file.GetName()] = file
			current.UpdateSize(file.GetSize())
		}
	}
	return fs
}

func sumDirSizesBelow(i *FsItem, maxSize int) (ans int) {
	if i.GetSize() <= maxSize {
		ans += i.GetSize()
	}
	for _, child := range i.Children {
		if child.GetType() == TypeDir {
			ans += sumDirSizesBelow(child, maxSize)
		}
	}
	return
}

func getClosestAbove(i *FsItem, minSize, currentSize int) int {
	if i.GetSize() < minSize || i.GetSize() > currentSize {
		return currentSize
	}
	curMin := i.GetSize()
	for _, d := range i.Children {
		if d.GetType() == TypeDir {
			curMin = utils.Min(curMin, getClosestAbove(d, minSize, i.GetSize()))
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
