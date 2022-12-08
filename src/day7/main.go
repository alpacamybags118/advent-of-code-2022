package main

import (
	helpers "advent-of-code-2022/src/helpers"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Directory struct {
	name            string
	ParentDirectory *Directory
	Subdirectories  []*Directory
	Files           []File
}

type Command struct {
	name string
	args string
}

func main() {
	useSample := flag.Bool("use-sample", false, "Provide flag to run solution with sample data rather than input data")
	dayToRun := flag.String("part", "1", "Provide which part to solve. Defaults to 1")

	flag.Parse()

	if *dayToRun == "1" {
		Part1(*useSample)
	} else {
		Part2(*useSample)
	}
}

func Part1(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileSystem, _ := CreateFileSystem(input)
	//PrintFileSystem(fileSystem, 0)
	sum := GetSumOfEligableDirs(fileSystem)
	fmt.Println(sum)
}

func Part2(useSample bool) {
	input, err := helpers.ReadInputAsStringArray(useSample)

	if err != nil {
		fmt.Println(err)
		return
	}

	fileSystem, _ := CreateFileSystem(input)
	//PrintFileSystem(fileSystem, 0)
	sum := FindSmallestSum(fileSystem)

	fmt.Println(sum)
}

func CreateFileSystem(input []string) (Directory, error) {
	var root Directory = Directory{
		name:            "/",
		ParentDirectory: nil,
		Subdirectories:  make([]*Directory, 0),
		Files:           make([]File, 0),
	}
	currentWorkingDir := &root

	for i := 0; i < len(input); i++ {
		firstChar := strings.Trim(string(input[i][0]), " ")
		if firstChar == "$" {
			command := GetCommand(input[i])

			switch command.name {
			case "cd":

				//fmt.Printf("cding into %s\n", command.args)
				//fmt.Println(currentWorkingDir)
				currentWorkingDir = RunChangeDirectoryCommand(currentWorkingDir, command.args, currentWorkingDir)
				if currentWorkingDir == nil {
					return root, errors.New(fmt.Sprintf("Could not find directory %s \n", command.args))
				}
				break
			case "ls":
				start := i + 1
				endIndex := ScanForNextCommand(input, start)
				currentWorkingDir = RunListCommand(currentWorkingDir, input[start:endIndex])
				break
			}
		}
	}

	return root, nil
}

func GetCommand(input string) Command {
	var name string
	var args string = ""

	parsedCommand := strings.Split(strings.TrimLeft(strings.TrimPrefix(input, "$"), " "), " ")
	name = parsedCommand[0]

	if (len(parsedCommand)) > 1 {
		args = parsedCommand[1]
	}

	return Command{
		name: name,
		args: args,
	}
}

func RunChangeDirectoryCommand(currentDirectory *Directory, args string, parentDir *Directory) *Directory {
	if args == "/" {
		//fmt.Println("returning root dir")
		return currentDirectory
	}

	fmt.Printf("Going to %s\n", args)
	if args == ".." {
		fmt.Printf("Changed directory up! Going from %s to %s\n", currentDirectory.name, currentDirectory.ParentDirectory.name)
		return currentDirectory.ParentDirectory
	}

	for _, child := range currentDirectory.Subdirectories {
		if child.name == args {
			return child
		}
	}

	return nil
}

func RunListCommand(workingDir *Directory, output []string) *Directory {
	for _, line := range output {
		lineContents := strings.Split(line, " ")
		if strings.Contains(lineContents[0], "dir") {
			newDir := &Directory{
				name:            lineContents[1],
				ParentDirectory: workingDir,
				Subdirectories:  make([]*Directory, 0),
				Files:           make([]File, 0),
			}

			workingDir.Subdirectories = append(workingDir.Subdirectories, newDir)

			//fmt.Printf("new dir: %s parent dir: %s\n", newDir.name, newDir.ParentDirectory.name)
		} else {
			size, _ := strconv.Atoi(lineContents[0])
			file := File{
				name: lineContents[1],
				size: size,
			}

			// var parentDir string = workingDir.name

			// if workingDir.ParentDirectory != nil {
			// 	parentDir = workingDir.ParentDirectory.name
			// }
			//fmt.Printf("apending %s to dir %s, parent dir %s\n", file.name, workingDir.name, parentDir)
			workingDir.Files = append(workingDir.Files, file)
		}
	}

	return workingDir
}

func ScanForNextCommand(input []string, startingIndex int) int {
	for i := startingIndex; i < len(input); i++ {
		firstChar := strings.Trim(string(input[i][0]), " ")
		if firstChar == "$" {
			return i
		}
	}

	return len(input)
}

func PrintFileSystem(dir Directory, depth int) {
	var spacing string = ""
	for i := 0; i < depth; i++ {
		spacing = fmt.Sprintf(" %s", spacing)
	}
	fmt.Printf("%s%s\n", spacing, dir.name)
	for _, files := range dir.Files {
		fmt.Printf(" %s- %s\n", spacing, files.name)
	}
	for _, subDirs := range dir.Subdirectories {
		if subDirs.Subdirectories != nil && len(subDirs.Subdirectories) > 0 {
			//fmt.Printf("dir %s subdirs are: ", subDirs.name)

			for _, subdir := range subDirs.Subdirectories {
				fmt.Printf("%s ", subdir.name)
			}
			fmt.Println()
		}
		newDepth := depth + 1
		PrintFileSystem(*subDirs, newDepth)
	}

	return
}

func GetSumOfEligableDirs(dir Directory) int {
	var sum int = 0
	dirList := GetAllDirectories(dir)
	for _, dirs := range dirList {
		if dirs.name == "/" {
			continue
		}

		dirSum := GetSumOfDirectory(*dirs)
		fmt.Printf("sum of %s is %v\n", dirs.name, dirSum)
		if dirSum <= 100000 {
			sum += dirSum
		}
	}

	return sum
}

func FindSmallestSum(dir Directory) int {
	var requiredFreeSpace int = 30000000
	var space int = 0

	dirList := GetAllDirectories(dir)

	totalSizeUsed := GetSumOfDirectory(dir)

	freeSpace := 70000000 - totalSizeUsed
	minSpaceNeeded := requiredFreeSpace - freeSpace

	for _, dirs := range dirList {
		if dirs.name == "/" {
			continue
		}

		dirSum := GetSumOfDirectory(*dirs)

		if dirSum >= minSpaceNeeded {
			if space == 0 {
				space = dirSum
			} else if dirSum < space {
				space = dirSum
			}
		}
	}

	return space
}

func GetAllDirectories(dir Directory) []*Directory {
	var dirList []*Directory = make([]*Directory, 0)

	dirList = append(dirList, &dir)

	if len(dir.Subdirectories) > 0 {
		for _, subDir := range dir.Subdirectories {
			dirList = append(dirList, GetAllDirectories(*subDir)...)
		}
	}

	return dirList
}

func GetSumOfDirectory(dir Directory) int {
	var sum int = 0
	for _, file := range dir.Files {
		sum += file.size
	}

	if len(dir.Subdirectories) > 0 {
		for _, subDir := range dir.Subdirectories {
			sum += GetSumOfDirectory(*subDir)
		}
	}

	return sum
}
