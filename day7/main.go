package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	a(input)
	b(input)
}

type dir struct {
	name  string
	files []file
	dirs  []dir
}

type file struct {
	name string
	size int
}

func findDirIdx(dirName string, dir *dir) int {
	for idx, d := range dir.dirs {
		if strings.HasPrefix(dirName, d.name) {
			return idx
		}
	}
	return -1
}

func getCurrentDirPointer(idx int, dirIdx []int, dirPtr *dir) *dir {
	if idx < len(dirIdx) {
		return getCurrentDirPointer(idx+1, dirIdx, &dirPtr.dirs[dirIdx[idx]])
	}

	return dirPtr
}

func parseDir(input string) dir {
	lines := strings.Split(input, "\r\n")
	var myDir dir
	dirIdx := []int{}
	r, _ := regexp.Compile(`^\$ cd ([a-z]+)$`)
	for _, line := range lines {
		if strings.HasPrefix(line, "$ cd /") {
			myDir = dir{
				name:  "/",
				files: []file{},
				dirs:  []dir{},
			}
			continue
		} else if strings.HasPrefix(line, "$ cd ..") {
			dirIdx = dirIdx[:len(dirIdx)-1]
			continue
		}

		cmd := strings.Split(line, " ")
		currentDir := getCurrentDirPointer(0, dirIdx, &myDir)

		if r.MatchString(line) {
			idx := findDirIdx(cmd[2], currentDir)
			if idx == -1 {
				panic("invalid directory")
			}
			dirIdx = append(dirIdx, idx)
		} else if strings.HasPrefix(line, "dir ") {
			currentDir.dirs = append(currentDir.dirs, dir{
				name:  cmd[1],
				dirs:  []dir{},
				files: []file{},
			})
		} else if v, err := strconv.Atoi(cmd[0]); err == nil {
			currentDir.files = append(currentDir.files, file{
				name: cmd[1],
				size: v,
			})
		}
	}
	return myDir
}

func sumNestedDir(dir *dir, sizes map[string]int, isLimit100k bool) (int, map[string]int) {
	fileSize := 0
	for _, f := range dir.files {
		fileSize += f.size
	}
	for _, d := range dir.dirs {
		size, _ := sumNestedDir(&d, sizes, isLimit100k)
		fileSize += size
	}
	if isLimit100k {
		if fileSize <= 100000 {
			sizes[dir.name] += fileSize
		}
	} else {
		sizes[dir.name] += fileSize
	}

	return fileSize, sizes
}

func a(input string) {
	myDir := parseDir(input)
	sizes := make(map[string]int)
	sumNestedDir(&myDir, sizes, true)
	ans := 0
	for _, v := range sizes {
		ans += v
	}
	fmt.Println(ans)
}

func b(input string) {
	const (
		fileSystemSize = 70000000
		updateSize     = 30000000
	)
	myDir := parseDir(input)
	sizes := make(map[string]int)
	sumNestedDir(&myDir, sizes, false)

	neededSize := updateSize - (fileSystemSize - sizes["/"])
	minDirSize := sizes["/"]
	for _, v := range sizes {
		if v >= neededSize && v <= minDirSize {
			minDirSize = v
		}
	}
	fmt.Println(minDirSize)
}
