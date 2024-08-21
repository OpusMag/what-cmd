package git

// Git represents a git command and its description
type Git struct {
    Name        string
    Description string
}

// List of git commands with descriptions
var Words = []Git{
    {"git init", "Create an empty Git repository or reinitialize an existing one"},
    {"git clone", "Clone a repository into a new directory"},
    {"git add", "Add file contents to the index"},
    {"git commit", "Record changes to the repository"},
    {"git status", "Show the working tree status"},
    {"git push", "Update remote refs along with associated objects"},
    {"git pull", "Fetch from and integrate with another repository or a local branch"},
    {"git fetch", "Download objects and refs from another repository"},
    {"git merge", "Join two or more development histories together"},
    {"git branch", "List, create, or delete branches"},
    {"git checkout", "Switch branches or restore working tree files"},
    {"git log", "Show commit logs"},
    {"git diff", "Show changes between commits, commit and working tree, etc"},
    {"git reset", "Reset current HEAD to the specified state"},
    {"git rm", "Remove files from the working tree and from the index"},
    {"git mv", "Move or rename a file, a directory, or a symlink"},
    {"git tag", "Create, list, delete or verify a tag object signed with GPG"},
    {"git stash", "Stash the changes in a dirty working directory away"},
    {"git rebase", "Reapply commits on top of another base tip"},
    {"git remote", "Manage set of tracked repositories"},
    {"git show", "Show various types of objects"},
    {"git archive", "Create an archive of files from a named tree"},
    {"git bisect", "Use binary search to find the commit that introduced a bug"},
    {"git blame", "Show what revision and author last modified each line of a file"},
    {"git cherry-pick", "Apply the changes introduced by some existing commits"},
    {"git grep", "Print lines matching a pattern"},
    {"git reflog", "Manage reflog information"},
    {"git submodule", "Initialize, update or inspect submodules"},
    {"git gc", "Cleanup unnecessary files and optimize the local repository"},
    {"git fsck", "Verifies the connectivity and validity of the objects in the database"},
    {"git clean", "Remove untracked files from the working tree"},
    {"git config", "Get and set repository or global options"},
    {"git describe", "Give an object a human-readable name based on an available ref"},
    {"git shortlog", "Summarize git log output"},
    {"git rev-parse", "Pick out and massage parameters"},
    {"git ls-tree", "List the contents of a tree object"},
    {"git update-index", "Register file contents in the working directory to the index"},
    {"git commit-tree", "Create a new commit object"},
    {"git symbolic-ref", "Read, modify and delete symbolic refs"},
    {"git worktree", "Manage multiple working trees"},
    {"git notes", "Add or inspect object notes"},
    {"git rerere", "Reuse recorded resolution of conflicted merges"},
    {"git range-diff", "Compare two commit ranges"},
    {"git replace", "Create, list, delete refs to replace objects"},
    {"git verify-commit", "Check the GPG signature of commits"},
    {"git verify-tag", "Check the GPG signature of tags"},
    {"git whatchanged", "Show logs with difference each commit introduces"},
    {"git instaweb", "Instantly browse your working repository in gitweb"},
    {"git daemon", "A really simple server for Git repositories"},
    {"git bundle", "Move objects and refs by archive"},
    {"git filter-branch", "Rewrite branches"},
    {"git mergetool", "Run merge conflict resolution tools to resolve merge conflicts"},
    {"git citool", "Graphical alternative to git-commit"},
    {"git gui", "A portable graphical interface to Git"},
    {"gitk", "The Git repository browser"},
    {"git instaweb", "Instantly browse your working repository in gitweb"},
    {"git archive", "Create an archive of files from a named tree"},
    {"git apply", "Apply a patch to files and/or to the index"},
    {"git am", "Apply a series of patches from a mailbox"},
    {"git cherry", "Find commits yet to be applied to upstream"},
    {"git format-patch", "Prepare patches for e-mail submission"},
    {"git send-email", "Send a collection of patches as emails"},
    {"git request-pull", "Generate a summary of pending changes"},
    {"git submodule", "Initialize, update or inspect submodules"},
    {"git svn", "Bidirectional operation between a Subversion repository and Git"},
    {"git fast-import", "Backend for fast Git data importers"},
    {"git fast-export", "Git data exporter"},
    {"git daemon", "A really simple server for Git repositories"},
    {"git instaweb", "Instantly browse your working repository in gitweb"},
    {"git instaweb", "Instantly browse your working repository in gitweb"},
}