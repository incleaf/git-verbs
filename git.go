package main

import (
	"io/ioutil"
	"log"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

// cloneRepository is a function that returns a repository of given Git url
func cloneRepository(url string) *git.Repository {
	// repository, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
	// 	URL:   url,
	// })

	// Tempdir to clone the repository
	dir, err := ioutil.TempDir("", "clone-example")
	if err != nil {
		log.Fatal(err)
	}

	// defer os.RemoveAll(dir) // clean up

	// Clones the repository into the given dir, just as a normal git clone does
	repository, err := git.PlainClone(dir, true, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		log.Fatal(err)
	}

	CheckIfError(err)
	return repository
}

// retrieveGitHistory is a function that returns entire git log of a given repository
func retrieveGitHistory(repository *git.Repository, ref *plumbing.Reference) []*object.Commit {
	cIter, err := repository.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	var history []*object.Commit
	err = cIter.ForEach(func(commit *object.Commit) error {
		history = append(history, commit)
		return nil
	})

	CheckIfError(err)

	return history
}

// CloneAndRetrieveGitHistory is a function that returns entire Git history of a given repository_url
func CloneAndRetrieveGitHistory(repositoryURL string) []*object.Commit {
	repository := cloneRepository(repositoryURL)
	ref, err := repository.Head()
	CheckIfError(err)

	history := retrieveGitHistory(repository, ref)
	return history
}
