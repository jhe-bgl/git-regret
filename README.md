
## Introduction
This is a Go program that can reverse a mistaken commit for each file, restoring them to their committed state before mistaken one. It identifies all the files involved in the erroneous commit, then for each file, it locates the commit prior to the mistake and creates a corresponding git restore command. This approach is very safe, as it doesn't make any actual changes to the Git repository. It only generates git restore commands. It can be used on mistake commits that happen before other other collaborators push their code without affecting the whole commit history. 

The difference of it compare to git revert is it will discard files changes after the mistaken commit.

## How to use:

Modify the repository path and commit ID in main.go to suit your needs, then simply run it with
```sh
go run .
```

## License

MIT