package main

import "fmt"

func main() {
	// fmt.Println("Fetching most starred repositories on GitHub...")
	// mostStarredRepositories := FetchMostStarredRepositories()
	// fmt.Printf("Toatal Count: %v\n", mostStarredRepositories.TotalCount)

	history := CloneAndRetrieveGitHistory("https://github.com/facebook/react")
	// history := CloneAndRetrieveGitHistory("https://github.com/src-d/go-git")
	fmt.Println(len(history))
}
