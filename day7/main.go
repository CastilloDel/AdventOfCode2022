package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("The result for the first part is:", part1())
	fmt.Println("The result for the second part is:", part2())
}

func part1() int {
	content, err := ioutil.ReadFile("day7/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	fileSystem := parseFileSystem(string(content))
	directorySizes := getDirectorySizes(fileSystem)

	total := 0
	for _, size := range directorySizes {
		if size < 100000 {
			total += size
		}
	}

	return total
}

func part2() int {
	content, err := ioutil.ReadFile("day7/input")
	if err != nil {
		panic("Couldn't read the input")
	}

	fileSystem := parseFileSystem(string(content))
	directorySizes := getDirectorySizes(fileSystem)

	sizeAvalaible := 70000000
	sizeNeeded := 30000000
	sizeUsed := directorySizes["/"]
	idealSizeToDelete := sizeNeeded - (sizeAvalaible - sizeUsed)
	sizeToDelete := sizeUsed
	for _, size := range directorySizes {
		if size >= idealSizeToDelete && size < sizeToDelete {
			sizeToDelete = size
		}
	}

	return sizeToDelete
}

type FileSystem = map[string]FileSystemElement

type FileSystemElement struct {
	name string
	path string
	// 0 if the element is a directory
	size int
	// empty if the element is a file
	childrenPaths []string
}

type File struct {
	name string
	size int
}

type LS struct {
	elements []FileSystemElement
}

type CD struct {
	destination string
}

func parseFileSystem(s string) FileSystem {
	unparsedCommands := strings.Split(s[2:], "\n$ ")

	commands := make([]interface{}, 0)
	for _, unparsedCommand := range unparsedCommands {
		commands = append(commands, parseCommand(unparsedCommand))
	}

	fileSystem := recreateFileSystem(commands)
	return fileSystem
}

func parseCommand(s string) interface{} {
	if s[:2] == "cd" {
		return parseCD(s)
	}
	return parseLS(s)
}

func parseCD(s string) CD {
	regex := regexp.MustCompile("cd (.+)")
	matches := regex.FindStringSubmatch(s)
	return CD{destination: matches[1]}
}

func parseLS(s string) LS {
	elements := make([]FileSystemElement, 0)
	for _, line := range strings.Split(s, "\n")[1:] {
		elements = append(elements, parseFileSystemElement(line))
	}
	return LS{elements}
}

func parseFileSystemElement(s string) FileSystemElement {
	regex := regexp.MustCompile("(\\d+) ([\\w\\.]+)|dir ([\\w\\.]+)")
	matches := regex.FindStringSubmatch(s)
	if matches[3] != "" {
		return FileSystemElement{name: matches[3], size: 0}
	}
	size, err := strconv.Atoi(matches[1])
	if err != nil {
		panic("Couldn't parse the size of a file")
	}
	return FileSystemElement{name: matches[2], size: size}
}

func recreateFileSystem(commands []interface{}) FileSystem {
	path := "/"
	fileSystem := make(FileSystem)
	for _, command := range commands {
		switch value := command.(type) {
		case CD:
			dest := value.destination
			if dest == "/" {
				path = dest
			} else if dest == ".." {
				pathParts := strings.Split(path, "/")
				path = strings.Join(pathParts[:len(pathParts)-1], "/")
			} else {
				if path == "/" {
					path += dest
				} else {
					path += "/" + dest
				}
			}
		case LS:
			elements := []string{}
			for _, element := range value.elements {
				if path == "/" {
					element.path = path + element.name
				} else {
					element.path = path + "/" + element.name
				}
				fileSystem[element.path] = element
				elements = append(elements, element.path)
			}
			fileSystem[path] = FileSystemElement{name: fileSystem[path].name, path: path, size: 0, childrenPaths: elements}
		}
	}
	return fileSystem
}

func getDirectorySizes(fileSystem FileSystem) map[string]int {
	sizes := map[string]int{}
	getDirectorySize(fileSystem, "/", sizes)
	return sizes
}

func getDirectorySize(fileSystem FileSystem, path string, sizes map[string]int) {
	size := 0
	for _, childrenPath := range fileSystem[path].childrenPaths {
		element := fileSystem[childrenPath]
		if element.size != 0 {
			size += element.size
		} else {
			getDirectorySize(fileSystem, childrenPath, sizes)
			size += sizes[childrenPath]
		}
	}
	sizes[path] = size
}
