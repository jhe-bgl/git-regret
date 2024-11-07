This is a Go program that can reverse a mistaken commit for each file, restoring them to their last committed state. It identifies all the files involved in the erroneous commit, then for each file, it locates the commit prior to the mistake and creates a corresponding git restore command. This approach is very safe, as it doesn't make any actual changes to the Git repository. It only generates git restore commands.

How to use:

Modify the repository path and commit ID in main.go to suit your needs, then simply run the program.
