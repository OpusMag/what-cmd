package flags

// Flag represents a command and its flags
type Flag struct {
    Name        string
    Description string
}

// List of commands with flags and their descriptions
var Words = map[string]Flag{
    "sudo": {
        Name:        "sudo",
        Description: "Flags: -u (Run the command as a specified user), -s (Run the command in a shell), -b (Run the command in the background)",
    },
    "cat": {
        Name:        "cat",
        Description: "Flags: -n (Number all output lines), -b (Number non-blank output lines), -s (Suppress repeated empty output lines)",
    },
    "ls": {
        Name:        "ls",
        Description: "Flags: -l (Use a long listing format), -a (Include directory entries whose names begin with a dot (.) ), -h (With -l, print sizes in human-readable format (e.g., 1K 234M 2G))",
    },
    "cd": {
        Name:        "cd",
        Description: "Flags: No flags available",
    },
    "pwd": {
        Name:        "pwd",
        Description: "Flags: -L (Print the value of $PWD if it names the current working directory), -P (Print the physical directory, without any symbolic links)",
    },
    "touch": {
        Name:        "touch",
        Description: "Flags: -a (Change only the access time), -m (Change only the modification time), -c (Do not create any files)",
    },
    "mkdir": {
        Name:        "mkdir",
        Description: "Flags: -p (Make parent directories as needed), -v (Print a message for each created directory), -m (Set file mode (as in chmod), not a=rwx - umask)",
    },
    "rm": {
        Name:        "rm",
        Description: "Flags: -r (Remove directories and their contents recursively), -f (Ignore nonexistent files and arguments, never prompt), -v (Explain what is being done)",
    },
    "cp": {
        Name:        "cp",
        Description: "Flags: -r (Copy directories recursively), -i (Prompt before overwrite), -v (Explain what is being done)",
    },
    "mv": {
        Name:        "mv",
        Description: "Flags: -i (Prompt before overwrite), -f (Do not prompt before overwriting), -v (Explain what is being done)",
    },
    "clear": {
        Name:        "clear",
        Description: "Flags: No flags available",
    },
    "exit": {
        Name:        "exit",
        Description: "Flags: No flags available",
    },
    "echo": {
        Name:        "echo",
        Description: "Flags: -n (Do not output the trailing newline), -e (Enable interpretation of backslash escapes)",
    },
    "man": {
        Name:        "man",
        Description: "Flags: -k (Search the short descriptions and manual page names for the keyword), -f (Display only the first line of each entry), -a (Display all matching manual pages)",
    },
    "grep": {
        Name:        "grep",
        Description: "Flags: -i (Ignore case distinctions), -r (Read all files under each directory, recursively), -v (Invert the sense of matching, to select non-matching lines)",
    },
    "find": {
        Name:        "find",
        Description: "Flags: -name (Base of file name (the path with the leading directories removed) matches shell pattern), -type (File is of type (e.g., f for regular file, d for directory)), -exec (Execute command; true if 0 status is returned)",
    },
    "chmod": {
        Name:        "chmod",
        Description: "Flags: -R (Change files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made)",
    },
    "chown": {
        Name:        "chown",
        Description: "Flags: -R (Operate on files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made)",
    },
    "df": {
        Name:        "df",
        Description: "Flags: -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -T (Print file system type), -i (List inode information instead of block usage)",
    },
    "du": {
        Name:        "du",
        Description: "Flags: -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -s (Display only a total for each argument), -a (Write counts for all files, not just directories)",
    },
    "head": {
        Name:        "head",
        Description: "Flags: -n (Print the first NUM lines instead of the first 10), -c (Print the first NUM bytes of each file), -q (Never print headers giving file names)",
    },
    "tail": {
        Name:        "tail",
        Description: "Flags: -n (Output the last NUM lines, instead of the last 10), -f (Output appended data as the file grows), -c (Output the last NUM bytes)",
    },
    "tar": {
        Name:        "tar",
        Description: "Flags: -c (Create a new archive), -x (Extract files from an archive), -v (Verbally list files processed)",
    },
    "wget": {
        Name:        "wget",
        Description: "Flags: -O (Write documents to FILE), -q (Turn off Wget's output), -r (Specify recursive download)",
    },
    "curl": {
        Name:        "curl",
        Description: "Flags: -o (Write output to FILE instead of stdout), -I (Fetch the HTTP-header only), -L (Follow redirects)",
    },
    "ps": {
        Name:        "ps",
        Description: "Flags: -e (Select all processes), -f (Do full-format listing), -u (Display processes for a specific user)",
    },
    "kill": {
        Name:        "kill",
        Description: "Flags: -s (Specify the signal to send), -l (List all signal names), -p (Specify the process ID)",
    },
    "top": {
        Name:        "top",
        Description: "Flags: -b (Batch mode operation), -n (Number of iterations), -u (Display processes for a specific user)",
    },
    "htop": {
        Name:        "htop",
        Description: "Flags: -d (Delay between updates, in tenths of seconds), -u (Display processes for a specific user), -p (Show only the given PIDs)",
    },
    "ssh": {
        Name:        "ssh",
        Description: "Flags: -l (Specify the user to log in as), -p (Port to connect to on the remote host), -i (File from which the identity (private key) is read)",
    },
    "scp": {
        Name:        "scp",
        Description: "Flags: -r (Recursively copy entire directories), -P (Specify the port to connect to on the remote host), -i (File from which the identity (private key) is read)",
    },
    "nano": {
        Name:        "nano",
        Description: "Flags: -B (Make a backup of the file), -m (Enable mouse support), -i (Automatically indent new lines to the same position as the previous line)",
    },
    "vim": {
        Name:        "vim",
        Description: "Flags: -u (Use the specified vimrc file), -N (No compatible mode), -e (Start in Ex mode)",
    },
    "apt-get": {
        Name:        "apt-get",
        Description: "Flags: -y (Assume yes to all prompts and run non-interactively), -q (Quiet; produces output suitable for logging), -d (Download only; package files are only retrieved, not unpacked or installed)",
    },
    "yum": {
        Name:        "yum",
        Description: "Flags: -y (Assume yes to all prompts), -q (Quiet mode), -v (Verbose mode)",
    },
    "mount": {
        Name:        "mount",
        Description: "Flags: -t (Indicate the filesystem type), -o (Specify mount options), -v (Verbose mode)",
    },
    "umount": {
        Name:        "umount",
        Description: "Flags: -f (Force unmount), -l (Lazy unmount), -v (Verbose mode)",
    },
    "uname": {
        Name:        "uname",
        Description: "Flags: -a (Print all information), -r (Print the kernel release), -v (Print the kernel version)",
    },
    "uptime": {
        Name:        "uptime",
        Description: "Flags: No flags available",
    },
    "whoami": {
        Name:        "whoami",
        Description: "Flags: No flags available",
    },
    "hostname": {
        Name:        "hostname",
        Description: "Flags: -i (Display the IP address), -f (Display the FQDN (Fully Qualified Domain Name)), -d (Display the domain name)",
    },
    "history": {
        Name:        "history",
        Description: "Flags: -c (Clear the history list), -d (Delete the history entry at position offset), -a (Append the new history lines to the history file)",
    },
    "alias": {
        Name:        "alias",
        Description: "Flags: No flags available",
    },
    "unalias": {
        Name:        "unalias",
        Description: "Flags: -a (Remove all alias definitions)",
    },
    "env": {
        Name:        "env",
        Description: "Flags: No flags available",
    },
    "export": {
        Name:        "export",
        Description: "Flags: No flags available",
    },
        "source": {
    Name:        "source",
    Description: "Flags: No flags available",
	},
	"su": {
		Name:        "su",
		Description: "Flags: -c (Pass command to the shell with -c), -l (Make the shell a login shell), -s (Run the specified shell instead of the default)",
	},
	"passwd": {
		Name:        "passwd",
		Description: "Flags: -d (Delete a user's password), -e (Expire a user's password), -l (Lock a user's password)",
	},
	"useradd": {
		Name:        "useradd",
		Description: "Flags: -m (Create the user's home directory), -s (Specify the user's login shell), -G (Specify supplementary groups for the user)",
	},
	"usermod": {
		Name:        "usermod",
		Description: "Flags: -l (Change the user's login name), -G (Specify supplementary groups for the user), -s (Specify the user's login shell)",
	},
	"userdel": {
		Name:        "userdel",
		Description: "Flags: -r (Remove the user's home directory and mail spool), -f (Force removal of the user)",
	},
	"groupadd": {
		Name:        "groupadd",
		Description: "Flags: -g (Specify the group ID), -r (Create a system group)",
	},
	"groupmod": {
		Name:        "groupmod",
		Description: "Flags: -n (Change the name of the group), -g (Change the group ID)",
	},
	"groupdel": {
		Name:        "groupdel",
		Description: "Flags: No flags available",
	},
	"crontab": {
		Name:        "crontab",
		Description: "Flags: -e (Edit the current crontab), -l (List the current crontab), -r (Remove the current crontab)",
	},
	"at": {
		Name:        "at",
		Description: "Flags: -f (Read the job from the specified file), -m (Send mail to the user when the job has completed), -v (Verbose mode)",
	},
	"jobs": {
		Name:        "jobs",
		Description: "Flags: -l (List process IDs in addition to the normal information), -p (List only the process IDs), -n (Display only jobs that have changed status since the last notification)",
	},
	"bg": {
		Name:        "bg",
		Description: "Flags: No flags available",
	},
	"fg": {
		Name:        "fg",
		Description: "Flags: No flags available",
	},
	"killall": {
		Name:        "killall",
		Description: "Flags: -i (Ask for confirmation before killing each process), -v (Report if the signal was successfully sent), -w (Wait for all processes to die)",
	},
	"xargs": {
		Name:        "xargs",
		Description: "Flags: -n (Use at most max-args arguments per command line), -P (Run up to max-procs processes at a time), -I (Replace occurrences of replace-str in the initial arguments with names read from standard input)",
	},
	"tee": {
		Name:        "tee",
		Description: "Flags: -a (Append to the given files, do not overwrite), -i (Ignore interrupts)",
	},
	"diff": {
		Name:        "diff",
		Description: "Flags: -u (Use the unified output format), -c (Use the context output format), -i (Ignore case differences in file contents)",
	},
	"patch": {
		Name:        "patch",
		Description: "Flags: -p (Strip the smallest prefix containing num leading slashes from each file name found in the patch file), -R (Assume the patch was already applied and attempt to un-apply it), -N (Ignore patches that seem to be already applied)",
	},
	"sed": {
		Name:        "sed",
		Description: "Flags: -n (Suppress automatic printing of pattern space), -e (Add the script to the commands to be executed), -i (Edit files in place (makes backup if suffix supplied))",
	},
	"awk": {
		Name:        "awk",
		Description: "Flags: -F (Use the specified field separator), -v (Assign the specified variable), -f (Read the awk program source from the specified file)",
	},
	"tr": {
		Name:        "tr",
		Description: "Flags: -d (Delete characters in the first set from the input), -s (Replace each input sequence of a repeated character that is listed in the last specified set with a single occurrence of that character), -c (Complement the set of characters in string1)",
	},
	"cut": {
		Name:        "cut",
		Description: "Flags: -f (Select only these fields), -d (Use the specified delimiter), -c (Select only these characters)",
	},
	"sort": {
		Name:        "sort",
		Description: "Flags: -r (Reverse the result of comparisons), -n (Compare according to string numerical value), -k (Sort via a key)",
	},
	"uniq": {
		Name:        "uniq",
		Description: "Flags: -c (Prefix lines by the number of occurrences), -d (Only print duplicate lines), -u (Only print unique lines)",
	},
	"wc": {
		Name:        "wc",
		Description: "Flags: -l (Print the newline counts), -w (Print the word counts), -c (Print the byte counts)",
	},
	"basename": {
		Name:        "basename",
		Description: "Flags: -s (Remove a trailing suffix)",
	},
	"dirname": {
		Name:        "dirname",
		Description: "Flags: No flags available",
	},
	"readlink": {
		Name:        "readlink",
		Description: "Flags: -f (Canonicalize by following every symlink in every component of the given name recursively), -e (Same as -f, but fail if any component is missing or not a directory), -m (Canonicalize by following every symlink in every component of the given name recursively, but do not fail if any component is missing or not a directory)",
	},
	"ln": {
		Name:        "ln",
		Description: "Flags: -s (Make symbolic links instead of hard links), -f (Remove existing destination files), -v (Print name of each linked file)",
	},
	"stat": {
		Name:        "stat",
		Description: "Flags: -f (Display information about the filesystem instead of the file), -t (Print the information in terse form), -c (Use the specified format instead of the default)",
	},
	"file": {
		Name:        "file",
		Description: "Flags: -b (Do not prepend filenames to output lines), -i (Output MIME type strings), -z (Try to look inside compressed files)",
	},
	"strings": {
		Name:        "strings",
		Description: "Flags: -a (Scan the whole file, not just the data segment), -n (Print sequences of at least n characters (default is 4)), -t (Print the location of the string in the file)",
	},
	"md5sum": {
		Name:        "md5sum",
		Description: "Flags: -c (Read MD5 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode)",
	},
	"sha256sum": {
		Name:        "sha256sum",
		Description: "Flags: -c (Read SHA256 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode)",
	},
	"gzip": {
		Name:        "gzip",
		Description: "Flags: -d (Decompress), -k (Keep (don't delete) input files), -r (Operate recursively on directories)",
	},
	"gunzip": {
		Name:        "gunzip",
		Description: "Flags: -k (Keep (don't delete) input files), -r (Operate recursively on directories)",
	},
	"bzip2": {
		Name:        "bzip2",
		Description: "Flags: -d (Decompress), -k (Keep (don't delete) input files), -v (Verbose mode)",
	},
	"bunzip2": {
		Name:        "bunzip2",
		Description: "Flags: -k (keep the original file after decompression), -v (verbose mode, show more details during operation)",
	},
	"zip": {
		Name:        "zip",
		Description: "Flags: -r (recursively include files in subdirectories), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation)",
	},
	"unzip": {
		Name:        "unzip",
		Description: "Flags: -l (list the contents of the archive), -t (test the integrity of the archive), -d (specify the directory to extract files to)",
	},
	"rar": {
		Name:        "rar",
		Description: "Flags: -a (add files to the archive), -x (extract files from the archive), -v (create volumes with a specified size)",
	},
	"unrar": {
		Name:        "unrar",
		Description: "Flags: -x (extract files from the archive), -l (list the contents of the archive), -v (verbose mode, show more details during operation)",
	},
	"7z": {
		Name:        "7z",
		Description: "Flags: -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"7za": {
		Name:        "7za",
		Description: "Flags: -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"7zr": {
		Name:        "7zr",
		Description: "Flags: -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"free": {
		Name:        "free",
		Description: "Flags: -h (display sizes in human-readable format), -m (display memory in megabytes), -g (display memory in gigabytes)",
	},
	"who": {
		Name:        "who",
		Description: "Flags: -a (display all information), -b (show the last system boot time), -q (display the number of logged-in users)",
	},
	"last": {
		Name:        "last",
		Description: "Flags: -n (specify the number of lines to display), -R (do not display the hostname), -x (display system shutdown and reboot entries)",
	},
	"dmesg": {
		Name:        "dmesg",
		Description: "Flags: -C (clear the ring buffer), -c (read and clear the ring buffer), -T (display human-readable timestamps)",
	},
	"lsblk": {
		Name:        "lsblk",
		Description: "Flags: -f (display file system information), -o (specify output columns), -d (do not display child devices)",
	},
	"blkid": {
		Name:        "blkid",
		Description: "Flags: -p (probe for partitions), -s (show specific information), -o (specify output format)",
	},
	"fdisk": {
		Name:        "fdisk",
		Description: "Flags: -l (list partition tables), -u (specify units), -c (compatibility mode)",
	},
	"mkfs": {
		Name:        "mkfs",
		Description: "Flags: -t (specify the file system type), -c (check the device for bad blocks), -v (verbose mode, show more details during operation)",
	},
	"fsck": {
		Name:        "fsck",
		Description: "Flags: -a (automatically repair the file system), -r (prompt interactively before making repairs), -y (assume 'yes' to all prompts)",
	},
	"parted": {
		Name:        "parted",
		Description: "Flags: -l (list partition tables), -s (script mode, do not prompt for user input), -a (specify alignment type)",
	},
	"dd": {
		Name:        "dd",
		Description: "Flags: if (specify input file), of (specify output file), bs (specify block size)",
	},
	"sync": {
		Name:        "sync",
		Description: "Flags: No flags available",
	},
	"shutdown": {
		Name:        "shutdown",
		Description: "Flags: -h (halt the system), -r (reboot the system), -c (cancel a scheduled shutdown)",
	},
	"reboot": {
		Name:        "reboot",
		Description: "Flags: No flags available",
	},
	"systemctl": {
		Name:        "systemctl",
		Description: "Flags: -t (specify unit type), -a (show all units), -l (show full output)",
	},
	"journalctl": {
		Name:        "journalctl",
		Description: "Flags: -f (follow new log entries), -u (show logs for a specific unit), -p (specify priority level)",
	},
	"timedatectl": {
		Name:        "timedatectl",
		Description: "Flags: -s (set the system time), -p (print properties), -l (list time zones)",
	},
	"hostnamectl": {
		Name:        "hostnamectl",
		Description: "Flags: -s (set the system hostname), -p (print properties), -l (list hostnames)",
	},
	"localectl": {
		Name:        "localectl",
		Description: "Flags: -s (set locale settings), -p (print properties), -l (list locales)",
	},
	"loginctl": {
		Name:        "loginctl",
		Description: "Flags: -s (set login settings), -p (print properties), -l (list sessions)",
	},
	"nmcli": {
		Name:        "nmcli",
		Description: "Flags: -t (terse output), -p (pretty output), -m (specify mode)",
	},
	"iwconfig": {
		Name:        "iwconfig",
		Description: "Flags: -a (display all information), -d (enable debug mode), -s (display status)",
	},
	"iwlist": {
		Name:        "iwlist",
		Description: "Flags: -a (display all information), -d (enable debug mode), -s (display status)",
	},
	"ufw": {
		Name:        "ufw",
		Description: "Flags: -a (allow traffic), -d (deny traffic), -r (reject traffic)",
	},
	"fail2ban": {
		Name:        "fail2ban",
		Description: "Flags: -d (debug mode), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation)",
	},
	"logrotate": {
		Name:        "logrotate",
		Description: "Flags: -d (debug mode), -f (force mode), -v (verbose mode, show more details during operation)",
	},
	"cron": {
		Name:        "cron",
		Description: "Flags: No flags available",
	},
	"anacron": {
		Name:        "anacron",
		Description: "Flags: -d (debug mode), -s (safe mode), -u (update timestamps)",
	},
	"systemd-analyze": {
		Name:        "systemd-analyze",
		Description: "Flags: -p (specify property), -t (specify type), -f (specify filter)",
	},
	"pkill": {
		Name:        "pkill",
		Description: "Flags: -u (specify user), -t (specify terminal), -x (exact match)",
	},
	"pgrep": {
		Name:        "pgrep",
		Description: "Flags: -u (specify user), -t (specify terminal), -x (exact match)",
	},
	"nice": {
		Name:        "nice",
		Description: "Flags: -n (specify adjustment value)",
	},
	"renice": {
		Name:        "renice",
		Description: "Flags: -n (specify adjustment value), -p (specify process ID), -u (specify user)",
	},
	"ionice": {
		Name:        "ionice",
		Description: "Flags: -c (specify class), -n (specify priority), -p (specify process ID)",
	},
	"watch": {
		Name:        "watch",
		Description: "Flags: -n (specify interval), -d (highlight differences), -t (turn off header)",
	},
	"screen": {
		Name:        "screen",
		Description: "Flags: -d (detach session), -r (reattach session), -S (specify session name)",
	},
	"tmux": {
		Name:        "tmux",
		Description: "Flags: -d (detach session), -r (reattach session), -S (specify session name)",
	},
	"nohup": {
		Name:        "nohup",
		Description: "Flags: No flags available",
	},
	"disown": {
		Name:        "disown",
		Description: "Flags: -h (do not remove job from table)",
	},
    "git init": {
        Name:        "git init",
        Description: "-q, (Quiet mode (suppress output))",
    },
    "git clone": {
        Name:        "git clone",
        Description: "-b, (Specify branch to clone)",
    },
    "git add": {
        Name:        "git add",
        Description: "-A, (Add all changes to the staging area)",
    },
    "git commit": {
        Name:        "git commit",
        Description: "-m, (Specify commit message)",
    },
    "git status": {
        Name:        "git status",
        Description: "-s, (Display status in short format)",
    },
    "git push": {
        Name:        "git push",
        Description: "-u, (Set upstream for the branch)",
    },
    "git pull": {
        Name:        "git pull",
        Description: "--rebase, (Rebase instead of merge)",
    },
    "git fetch": {
        Name:        "git fetch",
        Description: "--all, (Fetch all remotes)",
    },
    "git merge": {
        Name:        "git merge",
        Description: "--no-ff, (Create a merge commit)",
    },
    "git branch": {
        Name:        "git branch",
        Description: "-d, (Delete branch)",
    },
    "git checkout": {
        Name:        "git checkout",
        Description: "-b, (Create and switch to new branch)",
    },
    "git log": {
        Name:        "git log",
        Description: "--oneline, (Display log in one line per commit format)",
    },
    "git diff": {
        Name:        "git diff",
        Description: "--staged, (Show staged changes)",
    },
    "git reset": {
        Name:        "git reset",
        Description: "--hard, (Reset working directory to last commit)",
    },
    "git rm": {
        Name:        "git rm",
        Description: "-r, (Remove directories recursively)",
    },
    "git mv": {
        Name:        "git mv",
        Description: "-f, (Force move or rename)",
    },
    "git tag": {
        Name:        "git tag",
        Description: "-a, (Create annotated tag)",
    },
    "git stash": {
        Name:        "git stash",
        Description: "pop, (Apply and remove stash)",
    },
    "git rebase": {
        Name:        "git rebase",
        Description: "-i, (Interactive rebase)",
    },
    "git remote": {
        Name:        "git remote",
        Description: "add, (Add a new remote)",
    },
    "git show": {
        Name:        "git show",
        Description: "--stat, (Show stats of changes)",
    },
    "git archive": {
        Name:        "git archive",
        Description: "--format, (Specify archive format), --prefix, (Add prefix to files)",
    },
    "git bisect": {
        Name:        "git bisect",
        Description: "start, (Start bisecting)",
    },
    "git blame": {
        Name:        "git blame",
        Description: "-L, (Annotate only the given line range)",
    },
    "git cherry-pick": {
        Name:        "git cherry-pick",
        Description: "-n, (Apply changes without committing)",
    },
    "git grep": {
        Name:        "git grep",
        Description: "-n, (Show line numbers)",
    },
    "git reflog": {
        Name:        "git reflog",
        Description: "expire, (Expire old reflog entries)",
    },
    "git submodule": {
        Name:        "git submodule",
        Description: "update, (Update submodules), foreach, (Run command in each submodule)",
    },
    "git gc": {
        Name:        "git gc",
        Description: "--aggressive, (Aggressive garbage collection)",
    },
    "git fsck": {
        Name:        "git fsck",
        Description: "--full, (Check full object connectivity)",
    },
    "git clean": {
        Name:        "git clean",
        Description: "-f, (Force clean)",
    },
    "git config": {
        Name:        "git config",
        Description: "--global, (Set global configuration)",
    },
    "git describe": {
        Name:        "git describe",
        Description: "--tags, (Describe with tags)",
    },
    "git shortlog": {
        Name:        "git shortlog",
        Description: "-s, (Summarize by commit count)",
    },
    "git rev-parse": {
        Name:        "git rev-parse",
        Description: "--abbrev-ref, (Show abbreviated ref)",
    },
    "git ls-tree": {
        Name:        "git ls-tree",
        Description: "-r, (List recursively)",
    },
    "git update-index": {
        Name:        "git update-index",
        Description: "--assume-unchanged, (Mark files as unchanged)",
    },
    "git symbolic-ref": {
        Name:        "git symbolic-ref",
        Description: "HEAD, (Show current branch)",
    },
    "git worktree": {
        Name:        "git worktree",
        Description: "add, (Add a new working tree)",
    },
    "git notes": {
        Name:        "git notes",
        Description: "add, (Add a note)",
    },
    "git rerere": {
        Name:        "git rerere",
        Description: "clear, (Clear rerere cache)",
    },
    "git range-diff": {
        Name:        "git range-diff",
        Description: "--creation-factor, (Set creation factor)",
    },
    "git replace": {
        Name:        "git replace",
        Description: "--graft, (Graft a commit)",
    },
    "git verify-commit": {
        Name:        "git verify-commit",
        Description: "--verbose, (Show detailed information)",
    },
    "git verify-tag": {
        Name:        "git verify-tag",
        Description: "--verbose, (Show detailed information)",
    },
    "git whatchanged": {
        Name:        "git whatchanged",
        Description: "-p, (Show patch format)",
    },
    "git instaweb": {
        Name:        "git instaweb",
        Description: "--start, (Start web server), --stop, (Stop web server), --httpd, (Specify HTTP daemon)",
    },
    "git daemon": {
        Name:        "git daemon",
        Description: "--export-all, (Export all repositories), --base-path, (Specify base path)",
    },
    "git bundle": {
        Name:        "git bundle",
        Description: "create, (Create a bundle)",
    },
    "git filter-branch": {
        Name:        "git filter-branch",
        Description: "--tree-filter, (Filter tree)",
    },
    "git mergetool": {
        Name:        "git mergetool",
        Description: "--tool, (Specify merge tool)",
    },
    "git citool": {
        Name:        "git citool",
        Description: "--no-wait, (Do not wait for user input)",
    },
    "git gui": {
        Name:        "git gui",
        Description: "blame, (Blame in GUI)",
    },
    "gitk": {
        Name:        "gitk",
        Description: "--all, (Show all refs)",
    },
    "git apply": {
        Name:        "git apply",
        Description: "--check, (Check if patch can be applied)",
    },
    "git am": {
        Name:        "git am",
        Description: "--abort, (Abort applying patches)",
    },
    "git cherry": {
        Name:        "git cherry",
        Description: "-v, (Verbose output)",
    },
    "git format-patch": {
        Name:        "git format-patch",
        Description: "-n, (Specify number of patches)",
    },
    "git send-email": {
        Name:        "git send-email",
        Description: "--to, (Specify recipient)",
    },
    "git request-pull": {
        Name:        "git request-pull",
        Description: "--summary, (Show summary)",
    },
    "git svn": {
        Name:        "git svn",
        Description: "clone, (Clone SVN repository)",
    },
    "git fast-import": {
        Name:        "git fast-import",
        Description: "--date-format, (Specify date format)",
    },
    "git fast-export": {
        Name:        "git fast-export",
        Description: "--signed-tags, (Export signed tags)",
    },
	"ping": {
		Name:        "ping",
		Description: "Flags: -c (Specify the number of packets to send), -i (Specify the interval between packets)",
	},
	"ifconfig": {
		Name:        "ifconfig",
		Description: "Flags: -a (Display all interfaces, including those that are down), -s (Display a short list of interfaces)",
	},
	"ip": {
		Name:        "ip",
		Description: "Flags: addr (Show or manipulate IP addresses), link (Show or manipulate network devices)",
	},
	"netstat": {
		Name:        "netstat",
		Description: "Flags: -a (Show all connections and listening ports), -r (Show the routing table)",
	},
	"ss": {
		Name:        "ss",
		Description: "Flags: -t (Show TCP sockets), -u (Show UDP sockets)",
	},
	"traceroute": {
		Name:        "traceroute",
		Description: "Flags: -m (Set the maximum number of hops), -p (Set the destination port to use)",
	},
	"mtr": {
		Name:        "mtr",
		Description: "Flags: -r (Generate a report in text format), -c (Set the number of pings to send)",
	},
	"dig": {
		Name:        "dig",
		Description: "Flags: +short (Provide a short output), +trace (Trace the path to the name server)",
	},
	"nslookup": {
		Name:        "nslookup",
		Description: "Flags: -type (Specify the type of DNS query (e.g., A, MX, etc.)), -timeout (Set the timeout period for the query)",
	},
	"host": {
		Name:        "host",
		Description: "Flags: -t (Specify the type of DNS query), -a (Display all available information)",
	},
	"rsync": {
		Name:        "rsync",
		Description: "Flags: -a (Enable archive mode (preserves permissions, times, symbolic links, etc.)), -v (Enable verbose output)",
	},
	"ftp": {
		Name:        "ftp",
		Description: "Flags: -n (Disable auto-login upon initial connection), -v (Enable verbose mode)",
	},
	"sftp": {
		Name:        "sftp",
		Description: "Flags: -P (Specify the port to connect to on the remote host), -r (Recursively copy entire directories)",
	},
	"telnet": {
		Name:        "telnet",
		Description: "Flags: -8 (Use an 8-bit data path), -E (Stop any escape character from being recognized)",
	},
	"nc": {
		Name:        "nc",
		Description: "Flags: -l (Enable listen mode for inbound connections), -u (Use UDP mode instead of the default TCP)",
	},
	"nmap": {
		Name:        "nmap",
		Description: "Flags: -sP (Perform a ping scan to determine which hosts are up), -sT (Perform a TCP connect scan)",
	},
	"tcpdump": {
		Name:        "tcpdump",
		Description: "Flags: -i (Specify the network interface to listen on), -w (Write the raw packets to a file)",
	},
	"wireshark": {
		Name:        "wireshark",
		Description: "Flags: -i (Specify the network interface to capture on), -k (Start capturing packets immediately)",
	},
	"iptables": {
		Name:        "iptables",
		Description: "Flags: -A (Append a rule to a chain), -D (Delete a rule from a chain)",
	},
	"firewalld": {
		Name:        "firewalld",
		Description: "Flags: --add-service (Add a service to the firewall), --remove-service (Remove a service from the firewall)",
	},
	"ethtool": {
		Name:        "ethtool",
		Description: "Flags: -s (Change network device settings), -i (Show driver information for a network device)",
	},
	"route": {
		Name:        "route",
		Description: "Flags: add (Add a new route), del (Delete an existing route)",
	},
	"arp": {
		Name:        "arp",
		Description: "Flags: -a (Display the ARP table), -d (Delete an entry from the ARP table)",
	},
    "whois": {
        Name:        "whois",
        Description: "Flags: -h (Specify the whois server to query), -p (Specify the port to connect to on the whois server)",
	},
	"ipcalc": {
		Name:        "ipcalc",
		Description: "Flags: -n (Show the network address), -b (Show the broadcast address)",
	},
	"netcat": {
		Name:        "netcat",
		Description: "Flags: -l (Enable listen mode for inbound connections), -u (Use UDP mode instead of the default TCP)",
	},
	"arping": {
		Name:        "arping",
		Description: "Flags: -c (Specify the count of ARP requests to send), -I (Specify the network interface to use)",
	},
	"iw": {
		Name:        "iw",
		Description: "Flags: dev (Show information about wireless devices), link (Show information about wireless links)",
	},
	"bridge": {
		Name:        "bridge",
		Description: "Flags: fdb (Manage the forwarding database), vlan (Manage VLANs)",
	},
	"tc": {
		Name:        "tc",
		Description: "Flags: qdisc (Manage queueing disciplines), filter (Manage traffic filters)",
	},
	"ipset": {
		Name:        "ipset",
		Description: "Flags: create (Create a new IP set), add (Add an entry to an IP set)",
	},
	"ip rule": {
		Name:        "ip rule",
		Description: "Flags: add (Add a new routing rule), del (Delete an existing routing rule)",
	},
	"ip addr": {
		Name:        "ip addr",
		Description: "Flags: add (Add a new IP address), del (Delete an existing IP address)",
	},
	"ip link": {
		Name:        "ip link",
		Description: "Flags: set (Change device attributes), show (Display device attributes)",
	},
	"ip route": {
		Name:        "ip route",
		Description: "Flags: add (Add a new route), del (Delete an existing route)",
	},
	"ip neigh": {
		Name:        "ip neigh",
		Description: "Flags: add (Add a new neighbor entry), del (Delete an existing neighbor entry)",
	},
	"ip maddr": {
		Name:        "ip maddr",
		Description: "Flags: add (Add a new multicast address), del (Delete an existing multicast address)",
	},
	"ip monitor": {
		Name:        "ip monitor",
		Description: "Flags: all (Monitor all events), link (Monitor link events)",
	},
	"ip tunnel": {
		Name:        "ip tunnel",
		Description: "Flags: add (Add a new tunnel), del (Delete an existing tunnel)",
	},
	"ip xfrm": {
		Name:        "ip xfrm",
		Description: "Flags: state (Manage IPsec state), policy (Manage IPsec policy)",
	},
	"ip netns": {
		Name:        "ip netns",
		Description: "Flags: add (Add a new network namespace), del (Delete an existing network namespace)",
	},
	"ip vrf": {
		Name:        "ip vrf",
		Description: "Flags: add (Add a new VRF (Virtual Routing and Forwarding)), del (Delete an existing VRF)",
	},
	"ip link set": {
		Name:        "ip link set",
		Description: "Flags: up (Enable a network device), down (Disable a network device), mtu (Change the Maximum Transmission Unit (MTU) of a device), address (Change the MAC address of a device), master (Set the master device), nomaster (Unset the master device), type (Change the type of a device), netns (Move a device to a different network namespace), alias (Set an alias for a device), group (Set the group of a device), txqueuelen (Change the transmit queue length of a device), vf (Change the attributes of a virtual function), xdp (Attach or detach an XDP (eXpress Data Path) program)",
	},
	"ip link add": {
		Name:        "ip link add",
		Description: "Flags: type (Specify the type of device to add (e.g., bridge, bond, dummy, etc.)), name (Specify the name of the device to add)",
	},
	"ip link delete": {
		Name:        "ip link delete",
		Description: "Flags: dev (Specify the name of the device to delete)",
	},
	"ip link show": {
		Name:        "ip link show",
		Description: "Flags: dev (Specify the name of the device to show)",
	},
	"ip link set dev": {
		Name:        "ip link set dev",
		Description: "Flags: mtu (Change the Maximum Transmission Unit (MTU) of the specified device), address (Change the MAC address of the specified device)",
	},
	"ip link set up": {
		Name:        "ip link set up",
		Description: "Flags: (Enable the specified device)",
	},
	"ip link set down": {
		Name:        "ip link set down",
		Description: "Flags: (Disable the specified device)",
	},
	"ip link set mtu": {
		Name:        "ip link set mtu",
		Description: "Flags: (Change the Maximum Transmission Unit (MTU) of the specified device)",
	},
	"ip link set address": {
		Name:        "ip link set address",
		Description: "Flags: (Change the MAC address of the specified device)",
	},
	"ip link set master": {
		Name:        "ip link set master",
		Description: "Flags: (Set the master device for the specified device)",
	},
	"ip link set nomaster": {
		Name:        "ip link set nomaster",
		Description: "Flags: (Unset the master device for the specified device)",
	},
	"ip link set type": {
		Name:        "ip link set type",
		Description: "Flags: (Change the type of the specified device)",
	},
	"ip link set netns": {
		Name:        "ip link set netns",
		Description: "Flags: (Move the specified device to a different network namespace)",
	},
	"ip link set alias": {
		Name:        "ip link set alias",
		Description: "Flags: (Set an alias for the specified device)",
	},
	"ip link set group": {
		Name:        "ip link set group",
		Description: "Flags: (Set the group for the specified device)",
	},
	"ip link set txqueuelen": {
		Name:        "ip link set txqueuelen",
		Description: "Flags: (Change the transmit queue length of the specified device)",
	},
	"ip link set vf": {
		Name:        "ip link set vf",
		Description: "Flags: (Change the attributes of a virtual function for the specified device)",
	},
	"ip link set xdp": {
		Name:        "ip link set xdp",
		Description: "Flags: (Attach or detach an XDP (eXpress Data Path) program to/from the specified device)",
	},
}