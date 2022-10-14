package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"

	"github.com/exercism/go-representer"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide the path to the folder with the examples as argument")
	}

	goExe, err := exec.LookPath("go")
	if err != nil {
		log.Fatalf("failed to find go: %v", err)
	}

	examplesPath := os.Args[1]

	representationToSolutions := map[string][]int{}
	solutionsWithCompilerOrTestFailure := 0

	for i := 0; i < 500; i++ {
		// Write a log every 10 solutions so we can see some progress in the terminal
		// and notice if the processing is stuck.
		if i > 0 && i%10 == 0 {
			fmt.Printf("%d solutions processed\n", i)
		}

		pathToSolutionFolder := fmt.Sprintf("%s/%d", examplesPath, i)
		testCmd := &exec.Cmd{
			Dir:  pathToSolutionFolder,
			Path: goExe,
			Args: []string{"go", "test", "-timeout", "10s"},
		}
		err = testCmd.Run()
		if err != nil {
			solutionsWithCompilerOrTestFailure++
			continue
		}

		reprBytes, _, err := representer.Extract(pathToSolutionFolder)
		if err != nil {
			log.Fatalf("representer failed: %v", err)
		}

		err = os.WriteFile(pathToSolutionFolder+"/.meta/representation.txt", reprBytes, 0644)
		if err != nil {
			fmt.Printf("Failed to write representation for solution %d.\n", i)
		}

		representation := string(reprBytes)
		representationToSolutions[representation] = append(representationToSolutions[representation], i)
	}

	fmt.Printf("\nAll solutions processed!\n\n")

	type group struct {
		count     int
		solutions []int
	}

	groups := []group{}
	uniqueSolutions := []int{}

	for _, solutions := range representationToSolutions {
		if len(solutions) == 1 {
			uniqueSolutions = append(uniqueSolutions, solutions[0])
			continue
		}

		groups = append(groups, group{
			count:     len(solutions),
			solutions: solutions,
		})
	}

	sort.Slice(groups, func(i, j int) bool { return groups[i].count > groups[j].count })
	sort.Ints(uniqueSolutions)

	for _, group := range groups {
		fmt.Printf("%d solutions have the same representation as solution %d.\n", group.count, group.solutions[0])
	}
	fmt.Printf("\n%d unique solutions: %v\n\n", len(uniqueSolutions), uniqueSolutions)
	fmt.Println("Number of solutions that did not compile or failed the tests: ", solutionsWithCompilerOrTestFailure)
}
