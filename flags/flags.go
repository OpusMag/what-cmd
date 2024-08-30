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
        Description: ": -u (Run the command as a specified user), -s (Run the command in a shell), -b (Run the command in the background)",
    },
    "cat": {
        Name:        "cat",
        Description: ": -n (Number all output lines), -b (Number non-blank output lines), -s (Suppress repeated empty output lines)",
    },
    "ls": {
        Name:        "ls",
        Description: ": -l (Use a long listing format), -a (Include directory entries whose names begin with a dot (.) ), -h (With -l, print sizes in human-readable format (e.g., 1K 234M 2G))",
    },
    "cd": {
        Name:        "cd",
        Description: ": No flags available",
    },
    "pwd": {
        Name:        "pwd",
        Description: ": -L (Print the value of $PWD if it names the current working directory), -P (Print the physical directory, without any symbolic links)",
    },
    "touch": {
        Name:        "touch",
        Description: ": -a (Change only the access time), -m (Change only the modification time), -c (Do not create any files)",
    },
    "mkdir": {
        Name:        "mkdir",
        Description: ": -p (Make parent directories as needed), -v (Print a message for each created directory), -m (Set file mode (as in chmod), not a=rwx - umask)",
    },
    "rm": {
        Name:        "rm",
        Description: ": -r (Remove directories and their contents recursively), -f (Ignore nonexistent files and arguments, never prompt), -v (Explain what is being done)",
    },
    "cp": {
        Name:        "cp",
        Description: ": -r (Copy directories recursively), -i (Prompt before overwrite), -v (Explain what is being done)",
    },
    "mv": {
        Name:        "mv",
        Description: ": -i (Prompt before overwrite), -f (Do not prompt before overwriting), -v (Explain what is being done)",
    },
    "clear": {
        Name:        "clear",
        Description: ": No flags available",
    },
    "exit": {
        Name:        "exit",
        Description: ": No flags available",
    },
    "echo": {
        Name:        "echo",
        Description: ": -n (Do not output the trailing newline), -e (Enable interpretation of backslash escapes)",
    },
    "man": {
        Name:        "man",
        Description: ": -k (Search the short descriptions and manual page names for the keyword), -f (Display only the first line of each entry), -a (Display all matching manual pages)",
    },
    "grep": {
        Name:        "grep",
        Description: ": -i (Ignore case distinctions), -r (Read all files under each directory, recursively), -v (Invert the sense of matching, to select non-matching lines)",
    },
    "find": {
        Name:        "find",
        Description: ": -name (Base of file name (the path with the leading directories removed) matches shell pattern), -type (File is of type (e.g., f for regular file, d for directory)), -exec (Execute command; true if 0 status is returned)",
    },
    "chmod": {
        Name:        "chmod",
        Description: ": -R (Change files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made)",
    },
    "chown": {
        Name:        "chown",
        Description: ": -R (Operate on files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made)",
    },
    "df": {
        Name:        "df",
        Description: ": -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -T (Print file system type), -i (List inode information instead of block usage)",
    },
    "du": {
        Name:        "du",
        Description: ": -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -s (Display only a total for each argument), -a (Write counts for all files, not just directories)",
    },
    "head": {
        Name:        "head",
        Description: ": -n (Print the first NUM lines instead of the first 10), -c (Print the first NUM bytes of each file), -q (Never print headers giving file names)",
    },
    "tail": {
        Name:        "tail",
        Description: ": -n (Output the last NUM lines, instead of the last 10), -f (Output appended data as the file grows), -c (Output the last NUM bytes)",
    },
    "tar": {
        Name:        "tar",
        Description: ": -c (Create a new archive), -x (Extract files from an archive), -v (Verbally list files processed)",
    },
    "wget": {
        Name:        "wget",
        Description: ": -O (Write documents to FILE), -q (Turn off Wget's output), -r (Specify recursive download)",
    },
    "curl": {
		Name:        "curl",
		Description: ": -X GET (Perform a GET request), -X POST (Perform a POST request), -X PUT (Perform a PUT request), -X DELETE (Perform a DELETE request), -I (Fetch the HTTP headers only), -d (Send data in a POST request), -F (Submit form data), -H (Pass custom header(s) to the server), -u (Server user and password), -o (Write output to a file), -O (Save the file with the same name as the remote file), -L (Follow redirects), -k (Allow insecure server connections when using SSL), -s (Silent mode (don't show progress meter or error messages)), -v (Verbose mode (show detailed information about the request and response)), -C - (Resume a previous file transfer), -b (Send cookies from a file), -c (Save cookies to a file), -T (Upload a file), --proxy (Use a proxy for the request), --limit-rate (Limit the transfer rate), --max-time (Maximum time allowed for the transfer), --compressed (Request a compressed response), --data-urlencode (URL encode the data before sending it in a POST request), --http2 (Use HTTP/2 protocol), --http3 (Use HTTP/3 protocol)",
	},
    "ps": {
        Name:        "ps",
        Description: ": -e (Select all processes), -f (Do full-format listing), -u (Display processes for a specific user)",
    },
    "kill": {
        Name:        "kill",
        Description: ": -s (Specify the signal to send), -l (List all signal names), -p (Specify the process ID)",
    },
    "top": {
        Name:        "top",
        Description: ": -b (Batch mode operation), -n (Number of iterations), -u (Display processes for a specific user)",
    },
    "htop": {
        Name:        "htop",
        Description: ": -d (Delay between updates, in tenths of seconds), -u (Display processes for a specific user), -p (Show only the given PIDs)",
    },
    "ssh": {
        Name:        "ssh",
        Description: ": -l (Specify the user to log in as), -p (Port to connect to on the remote host), -i (File from which the identity (private key) is read)",
    },
    "scp": {
        Name:        "scp",
        Description: ": -r (Recursively copy entire directories), -P (Specify the port to connect to on the remote host), -i (File from which the identity (private key) is read)",
    },
    "nano": {
        Name:        "nano",
        Description: ": -B (Make a backup of the file), -m (Enable mouse support), -i (Automatically indent new lines to the same position as the previous line)",
    },
    "vim": {
        Name:        "vim",
        Description: ": -u (Use the specified vimrc file), -N (No compatible mode), -e (Start in Ex mode)",
    },
    "apt-get": {
        Name:        "apt-get",
        Description: ": -y (Assume yes to all prompts and run non-interactively), -q (Quiet; produces output suitable for logging), -d (Download only; package files are only retrieved, not unpacked or installed)",
    },
    "yum": {
        Name:        "yum",
        Description: ": -y (Assume yes to all prompts), -q (Quiet mode), -v (Verbose mode)",
    },
    "mount": {
        Name:        "mount",
        Description: ": -t (Indicate the filesystem type), -o (Specify mount options), -v (Verbose mode)",
    },
    "umount": {
        Name:        "umount",
        Description: ": -f (Force unmount), -l (Lazy unmount), -v (Verbose mode)",
    },
    "uname": {
        Name:        "uname",
        Description: ": -a (Print all information), -r (Print the kernel release), -v (Print the kernel version)",
    },
    "uptime": {
        Name:        "uptime",
        Description: ": No flags available",
    },
    "whoami": {
        Name:        "whoami",
        Description: ": No flags available",
    },
    "hostname": {
        Name:        "hostname",
        Description: ": -i (Display the IP address), -f (Display the FQDN (Fully Qualified Domain Name)), -d (Display the domain name)",
    },
    "history": {
        Name:        "history",
        Description: ": -c (Clear the history list), -d (Delete the history entry at position offset), -a (Append the new history lines to the history file)",
    },
    "alias": {
        Name:        "alias",
        Description: ": No flags available",
    },
    "unalias": {
        Name:        "unalias",
        Description: ": -a (Remove all alias definitions)",
    },
    "env": {
        Name:        "env",
        Description: ": No flags available",
    },
    "export": {
        Name:        "export",
        Description: ": No flags available",
    },
    "source": {
		Name:        "source",
		Description: ": No flags available",
	},
	"su": {
		Name:        "su",
		Description: ": -c (Pass command to the shell with -c), -l (Make the shell a login shell), -s (Run the specified shell instead of the default)",
	},
	"passwd": {
		Name:        "passwd",
		Description: ": -d (Delete a user's password), -e (Expire a user's password), -l (Lock a user's password)",
	},
	"useradd": {
		Name:        "useradd",
		Description: ": -m (Create the user's home directory), -s (Specify the user's login shell), -G (Specify supplementary groups for the user)",
	},
	"usermod": {
		Name:        "usermod",
		Description: ": -l (Change the user's login name), -G (Specify supplementary groups for the user), -s (Specify the user's login shell)",
	},
	"userdel": {
		Name:        "userdel",
		Description: ": -r (Remove the user's home directory and mail spool), -f (Force removal of the user)",
	},
	"groupadd": {
		Name:        "groupadd",
		Description: ": -g (Specify the group ID), -r (Create a system group)",
	},
	"groupmod": {
		Name:        "groupmod",
		Description: ": -n (Change the name of the group), -g (Change the group ID)",
	},
	"groupdel": {
		Name:        "groupdel",
		Description: ": No flags available",
	},
	"crontab": {
		Name:        "crontab",
		Description: ": -e (Edit the current crontab), -l (List the current crontab), -r (Remove the current crontab)",
	},
	"at": {
		Name:        "at",
		Description: ": -f (Read the job from the specified file), -m (Send mail to the user when the job has completed), -v (Verbose mode)",
	},
	"jobs": {
		Name:        "jobs",
		Description: ": -l (List process IDs in addition to the normal information), -p (List only the process IDs), -n (Display only jobs that have changed status since the last notification)",
	},
	"bg": {
		Name:        "bg",
		Description: ": No flags available",
	},
	"fg": {
		Name:        "fg",
		Description: ": No flags available",
	},
	"killall": {
		Name:        "killall",
		Description: ": -i (Ask for confirmation before killing each process), -v (Report if the signal was successfully sent), -w (Wait for all processes to die)",
	},
	"xargs": {
		Name:        "xargs",
		Description: ": -n (Use at most max-args arguments per command line), -P (Run up to max-procs processes at a time), -I (Replace occurrences of replace-str in the initial arguments with names read from standard input)",
	},
	"tee": {
		Name:        "tee",
		Description: ": -a (Append to the given files, do not overwrite), -i (Ignore interrupts)",
	},
	"diff": {
		Name:        "diff",
		Description: ": -u (Use the unified output format), -c (Use the context output format), -i (Ignore case differences in file contents)",
	},
	"patch": {
		Name:        "patch",
		Description: ": -p (Strip the smallest prefix containing num leading slashes from each file name found in the patch file), -R (Assume the patch was already applied and attempt to un-apply it), -N (Ignore patches that seem to be already applied)",
	},
	"sed": {
		Name:        "sed",
		Description: ": -n (Suppress automatic printing of pattern space), -e (Add the script to the commands to be executed), -i (Edit files in place (makes backup if suffix supplied))",
	},
	"awk": {
		Name:        "awk",
		Description: ": -F (Use the specified field separator), -v (Assign the specified variable), -f (Read the awk program source from the specified file)",
	},
	"tr": {
		Name:        "tr",
		Description: ": -d (Delete characters in the first set from the input), -s (Replace each input sequence of a repeated character that is listed in the last specified set with a single occurrence of that character), -c (Complement the set of characters in string1)",
	},
	"cut": {
		Name:        "cut",
		Description: ": -f (Select only these fields), -d (Use the specified delimiter), -c (Select only these characters)",
	},
	"sort": {
		Name:        "sort",
		Description: ": -r (Reverse the result of comparisons), -n (Compare according to string numerical value), -k (Sort via a key)",
	},
	"uniq": {
		Name:        "uniq",
		Description: ": -c (Prefix lines by the number of occurrences), -d (Only print duplicate lines), -u (Only print unique lines)",
	},
	"wc": {
		Name:        "wc",
		Description: ": -l (Print the newline counts), -w (Print the word counts), -c (Print the byte counts)",
	},
	"basename": {
		Name:        "basename",
		Description: ": -s (Remove a trailing suffix)",
	},
	"dirname": {
		Name:        "dirname",
		Description: ": No flags available",
	},
	"readlink": {
		Name:        "readlink",
		Description: ": -f (Canonicalize by following every symlink in every component of the given name recursively), -e (Same as -f, but fail if any component is missing or not a directory), -m (Canonicalize by following every symlink in every component of the given name recursively, but do not fail if any component is missing or not a directory)",
	},
	"ln": {
		Name:        "ln",
		Description: ": -s (Make symbolic links instead of hard links), -f (Remove existing destination files), -v (Print name of each linked file)",
	},
	"stat": {
		Name:        "stat",
		Description: ": -f (Display information about the filesystem instead of the file), -t (Print the information in terse form), -c (Use the specified format instead of the default)",
	},
	"file": {
		Name:        "file",
		Description: ": -b (Do not prepend filenames to output lines), -i (Output MIME type strings), -z (Try to look inside compressed files)",
	},
	"strings": {
		Name:        "strings",
		Description: ": -a (Scan the whole file, not just the data segment), -n (Print sequences of at least n characters (default is 4)), -t (Print the location of the string in the file)",
	},
	"md5sum": {
		Name:        "md5sum",
		Description: ": -c (Read MD5 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode)",
	},
	"sha256sum": {
		Name:        "sha256sum",
		Description: ": -c (Read SHA256 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode)",
	},
	"gzip": {
		Name:        "gzip",
		Description: ": -d (Decompress), -k (Keep (don't delete) input files), -r (Operate recursively on directories)",
	},
	"gunzip": {
		Name:        "gunzip",
		Description: ": -k (Keep (don't delete) input files), -r (Operate recursively on directories)",
	},
	"bzip2": {
		Name:        "bzip2",
		Description: ": -d (Decompress), -k (Keep (don't delete) input files), -v (Verbose mode)",
	},
	"bunzip2": {
		Name:        "bunzip2",
		Description: ": -k (keep the original file after decompression), -v (verbose mode, show more details during operation)",
	},
	"zip": {
		Name:        "zip",
		Description: ": -r (recursively include files in subdirectories), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation)",
	},
	"unzip": {
		Name:        "unzip",
		Description: ": -l (list the contents of the archive), -t (test the integrity of the archive), -d (specify the directory to extract files to)",
	},
	"rar": {
		Name:        "rar",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -v (create volumes with a specified size)",
	},
	"unrar": {
		Name:        "unrar",
		Description: ": -x (extract files from the archive), -l (list the contents of the archive), -v (verbose mode, show more details during operation)",
	},
	"7z": {
		Name:        "7z",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"7za": {
		Name:        "7za",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"7zr": {
		Name:        "7zr",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive)",
	},
	"free": {
		Name:        "free",
		Description: ": -h (display sizes in human-readable format), -m (display memory in megabytes), -g (display memory in gigabytes)",
	},
	"who": {
		Name:        "who",
		Description: ": -a (display all information), -b (show the last system boot time), -q (display the number of logged-in users)",
	},
	"last": {
		Name:        "last",
		Description: ": -n (specify the number of lines to display), -R (do not display the hostname), -x (display system shutdown and reboot entries)",
	},
	"dmesg": {
		Name:        "dmesg",
		Description: ": -C (clear the ring buffer), -c (read and clear the ring buffer), -T (display human-readable timestamps)",
	},
	"lsblk": {
		Name:        "lsblk",
		Description: ": -f (display file system information), -o (specify output columns), -d (do not display child devices)",
	},
	"blkid": {
		Name:        "blkid",
		Description: ": -p (probe for partitions), -s (show specific information), -o (specify output format)",
	},
	"fdisk": {
		Name:        "fdisk",
		Description: ": -l (list partition tables), -u (specify units), -c (compatibility mode)",
	},
	"mkfs": {
		Name:        "mkfs",
		Description: ": -t (specify the file system type), -c (check the device for bad blocks), -v (verbose mode, show more details during operation)",
	},
	"fsck": {
		Name:        "fsck",
		Description: ": -a (automatically repair the file system), -r (prompt interactively before making repairs), -y (assume 'yes' to all prompts)",
	},
	"parted": {
		Name:        "parted",
		Description: ": -l (list partition tables), -s (script mode, do not prompt for user input), -a (specify alignment type)",
	},
	"dd": {
		Name:        "dd",
		Description: ": if (specify input file), of (specify output file), bs (specify block size)",
	},
	"sync": {
		Name:        "sync",
		Description: ": No flags available",
	},
	"shutdown": {
		Name:        "shutdown",
		Description: ": -h (halt the system), -r (reboot the system), -c (cancel a scheduled shutdown)",
	},
	"reboot": {
		Name:        "reboot",
		Description: ": No flags available",
	},
	"systemctl": {
		Name:        "systemctl",
		Description: ": -t (specify unit type), -a (show all units), -l (show full output)",
	},
	"journalctl": {
		Name:        "journalctl",
		Description: ": -f (follow new log entries), -u (show logs for a specific unit), -p (specify priority level)",
	},
	"timedatectl": {
		Name:        "timedatectl",
		Description: ": -s (set the system time), -p (print properties), -l (list time zones)",
	},
	"hostnamectl": {
		Name:        "hostnamectl",
		Description: ": -s (set the system hostname), -p (print properties), -l (list hostnames)",
	},
	"localectl": {
		Name:        "localectl",
		Description: ": -s (set locale settings), -p (print properties), -l (list locales)",
	},
	"loginctl": {
		Name:        "loginctl",
		Description: ": -s (set login settings), -p (print properties), -l (list sessions)",
	},
	"nmcli": {
		Name:        "nmcli",
		Description: ": -t (terse output), -p (pretty output), -m (specify mode)",
	},
	"iwconfig": {
		Name:        "iwconfig",
		Description: ": -a (display all information), -d (enable debug mode), -s (display status)",
	},
	"iwconfig essid": {
		Name:        "iwconfig essid",
		Description: ": None",
	},
	"iwconfig mode": {
		Name:        "iwconfig mode",
		Description: ": None",
	},
	"iwconfig freq": {
		Name:        "iwconfig freq",
		Description: ": None",
	},
	"iwconfig ap": {
		Name:        "iwconfig ap",
		Description: ": None",
	},
	"iwconfig rate": {
		Name:        "iwconfig rate",
		Description: ": None",
	},
	"iwconfig txpower": {
		Name:        "iwconfig txpower",
		Description: ": None",
	},
	"iwconfig retry": {
		Name:        "iwconfig retry",
		Description: ": None",
	},
	"iwconfig rts": {
		Name:        "iwconfig rts",
		Description: ": None",
	},
	"iwconfig frag": {
		Name:        "iwconfig frag",
		Description: ": None",
	},
	"iwconfig key": {
		Name:        "iwconfig key",
		Description: ": None",
	},
	"iwconfig power": {
		Name:        "iwconfig power",
		Description: ": None",
	},
	"iwlist": {
		Name:        "iwlist",
		Description: ": -a (display all information), -d (enable debug mode), -s (display status)",
	},
	"iwlist scan": {
		Name:        "iwlist scan",
		Description: ": None",
	},
	"iwlist frequency": {
		Name:        "iwlist frequency",
		Description: ": None",
	},
	"iwlist bitrate": {
		Name:        "iwlist bitrate",
		Description: ": None",
	},
	"iwlist encryption": {
		Name:        "iwlist encryption",
		Description: ": None",
	},
	"iwlist power": {
		Name:        "iwlist power",
		Description: ": None",
	},
	"iwlist txpower": {
		Name:        "iwlist txpower",
		Description: ": None",
	},
	"iwlist retry": {
		Name:        "iwlist retry",
		Description: ": None",
	},
	"ufw": {
		Name:        "ufw",
		Description: ": -a (allow traffic), -d (deny traffic), -r (reject traffic)",
	},
	"fail2ban": {
		Name:        "fail2ban",
		Description: ": -d (debug mode), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation)",
	},
	"logrotate": {
		Name:        "logrotate",
		Description: ": -d (debug mode), -f (force mode), -v (verbose mode, show more details during operation)",
	},
	"cron": {
		Name:        "cron",
		Description: ": No flags available",
	},
	"anacron": {
		Name:        "anacron",
		Description: ": -d (debug mode), -s (safe mode), -u (update timestamps)",
	},
	"systemd-analyze": {
		Name:        "systemd-analyze",
		Description: ": -p (specify property), -t (specify type), -f (specify filter)",
	},
	"pkill": {
		Name:        "pkill",
		Description: ": -u (specify user), -t (specify terminal), -x (exact match)",
	},
	"pgrep": {
		Name:        "pgrep",
		Description: ": -u (specify user), -t (specify terminal), -x (exact match)",
	},
	"nice": {
		Name:        "nice",
		Description: ": -n (specify adjustment value)",
	},
	"renice": {
		Name:        "renice",
		Description: ": -n (specify adjustment value), -p (specify process ID), -u (specify user)",
	},
	"ionice": {
		Name:        "ionice",
		Description: ": -c (specify class), -n (specify priority), -p (specify process ID)",
	},
	"watch": {
		Name:        "watch",
		Description: ": -n (specify interval), -d (highlight differences), -t (turn off header)",
	},
	"screen": {
		Name:        "screen",
		Description: ": -d (detach session), -r (reattach session), -S (specify session name)",
	},
	"tmux": {
		Name:        "tmux",
		Description: ": -d (detach session), -r (reattach session), -S (specify session name)",
	},
	"nohup": {
		Name:        "nohup",
		Description: ": No flags available",
	},
	"disown": {
		Name:        "disown",
		Description: ": -h (do not remove job from table)",
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
		Description: ": -c (Specify the number of packets to send), -i (Specify the interval between packets)",
	},
	"ifconfig": {
		Name:        "ifconfig",
		Description: ": -a (Display all interfaces, including those that are down), -s (Display a short list of interfaces)",
	},
	"ip": {
		Name:        "ip",
		Description: ": addr (Show or manipulate IP addresses), link (Show or manipulate network devices)",
	},
	"netstat": {
		Name:        "netstat",
		Description: ": -a (Show all connections and listening ports), -r (Show the routing table)",
	},
	"ss": {
		Name:        "ss",
		Description: ": -t (Show TCP sockets), -u (Show UDP sockets)",
	},
	"traceroute": {
		Name:        "traceroute",
		Description: ": -m (Set the maximum number of hops), -p (Set the destination port to use), -4 (Use IPv4), -6 (Use IPv6), -I (Use ICMP ECHO for probes), -T (Use TCP SYN for probes), -U (Use UDP datagrams for probes), -p (Set the base UDP port number used in probes), -q (Set the number of probe queries per hop), -w (Set the time (in seconds) to wait for a response to a probe), -m (Set the max number of hops), -f (Set the first TTL (time-to-live) value)",
	},
	"mtr": {
		Name:        "mtr",
		Description: ": --report (Output using report mode), --report-cycles (Set the number of pings sent), --interval (Set the interval between ICMP ECHO requests), --timeout (Set the timeout for each probe), --tcp (Use TCP SYN packets instead of ICMP ECHO), --udp (Use UDP datagrams instead of ICMP ECHO), --port (Set the target port for TCP/UDP), --max-ttl (Set the max number of hops), --first-ttl (Set the first TTL (time-to-live) value), --show-ips (Show IP numbers instead of hostnames)",
	},
	"dig": {
		Name:        "dig",
		Description: ": +short (Provide a short output), +trace (Trace the path to the name server), @server (Specify the DNS server to query), -b (Set the source IP address of the query), -c (Specify the query class (IN, CH, HS)), -f (Perform batch processing of multiple queries), -k (Specify a TSIG key file), -p (Specify the port number to query on the server), -q (Specify the query name), -t (Specify the query type (A, AAAA, MX, etc.)), -x (Perform a reverse lookup), -y (Specify a TSIG key), +short (Provide a short answer), +trace (Trace the delegation path from the root name servers), +stats (Print query statistics), +noall (Set or clear all display flags), +answer (Display the answer section of a reply), +authority (Display the authority section of a reply), +additional (Display the additional section of a reply), +question (Display the question section of a reply)",
	},
	"nslookup": {
		Name:        "nslookup",
		Description: ": -type (Specify the type of DNS query (e.g., A, MX, etc.)), -timeout (Set the timeout period for the query)",
	},
	"host": {
		Name:        "host",
		Description: ": -t (Specify the type of DNS query), -a (Display all available information)",
	},
	"rsync": {
		Name:        "rsync",
		Description: ": -a (Archive mode), -v (Verbose mode), -z (Compress file data during the transfer), -r (Recurse into directories), -u (Skip files that are newer on the receiver), -l (Copy symlinks as symlinks), -p (Preserve permissions), -t (Preserve modification times), -g (Preserve group), -o (Preserve owner (super-user only)), -D (Preserve device and special files), -e (Specify the remote shell to use), --delete (Delete extraneous files from destination directories), --exclude (Exclude files matching PATTERN), --include (Include files matching PATTERN), --progress (Show progress during transfer), --stats (Give some file-transfer stats)",
	},
	"ftp": {
		Name:        "ftp",
		Description: ": -n (Disable auto-login upon initial connection), -v (Enable verbose mode)",
	},
	"ftp open": {
		Name:        "ftp open",
		Description: ": None",
	},
	"ftp close": {
		Name:        "ftp close",
		Description: ": None",
	},
	"ftp put": {
		Name:        "ftp put",
		Description: ": None",
	},
	"ftp mput": {
		Name:        "ftp mput",
		Description: ": None",
	},
	"ftp get": {
		Name:        "ftp get",
		Description: ": None",
	},
	"ftp mget": {
		Name:        "ftp mget",
		Description: ": None",
	},
	"ftp delete": {
		Name:        "ftp delete",
		Description: ": None",
	},
	"ftp mdelete": {
		Name:        "ftp mdelete",
		Description: ": None",
	},
	"ftp rename": {
		Name:        "ftp rename",
		Description: ": None",
	},
	"ftp mkdir": {
		Name:        "ftp mkdir",
		Description: ": None",
	},
	"ftp rmdir": {
		Name:        "ftp rmdir",
		Description: ": None",
	},
	"ftp ls": {
		Name:        "ftp ls",
		Description: ": None",
	},
	"ftp lcd": {
		Name:        "ftp lcd",
		Description: ": None",
	},
	"ftp cd": {
		Name:        "ftp cd",
		Description: ": None",
	},
	"ftp ascii": {
		Name:        "ftp ascii",
		Description: ": None",
	},
	"ftp binary": {
		Name:        "ftp binary",
		Description: ": None",
	},
	"ftp passive": {
		Name:        "ftp passive",
		Description: ": None",
	},
	"ftp active": {
		Name:        "ftp active",
		Description: ": None",
	},
	"sftp": {
		Name:        "sftp",
		Description: ": -P (Specify the port to connect to on the remote host), -r (Recursively copy entire directories)",
	},
	"sftp get": {
		Name:        "sftp get",
		Description: ": None",
	},
	"sftp put": {
		Name:        "sftp put",
		Description: ": None",
	},
	"sftp ls": {
		Name:        "sftp ls",
		Description: ": None",
	},
	"sftp cd": {
		Name:        "sftp cd",
		Description: ": None",
	},
	"sftp lcd": {
		Name:        "sftp lcd",
		Description: ": None",
	},
	"sftp mkdir": {
		Name:        "sftp mkdir",
		Description: ": None",
	},
	"sftp rmdir": {
		Name:        "sftp rmdir",
		Description: ": None",
	},
	"sftp rm": {
		Name:        "sftp rm",
		Description: ": None",
	},
	"sftp rename": {
		Name:        "sftp rename",
		Description: ": None",
	},
	"sftp chmod": {
		Name:        "sftp chmod",
		Description: ": None",
	},
	"sftp chown": {
		Name:        "sftp chown",
		Description: ": None",
	},
	"sftp chgrp": {
		Name:        "sftp chgrp",
		Description: ": None",
	},
	"sftp lls": {
		Name:        "sftp lls",
		Description: ": None",
	},
	"sftp lmkdir": {
		Name:        "sftp lmkdir",
		Description: ": None",
	},
	"sftp lrm": {
		Name:        "sftp lrm",
		Description: ": None",
	},
	"sftp lrename": {
		Name:        "sftp lrename",
		Description: ": None",
	},
	"sftp lchmod": {
		Name:        "sftp lchmod",
		Description: ": None",
	},
	"sftp lchown": {
		Name:        "sftp lchown",
		Description: ": None",
	},
	"sftp lchgrp": {
		Name:        "sftp lchgrp",
		Description: ": None",
	},
	"telnet": {
		Name:        "telnet",
		Description: ": -8 (Use an 8-bit data path), -E (Stop any escape character from being recognized)",
	},
	"telnet open": {
    Name:        "telnet open",
    Description: ": None",
	},
	"telnet close": {
		Name:        "telnet close",
		Description: ": None",
	},
	"telnet quit": {
		Name:        "telnet quit",
		Description: ": None",
	},
	"telnet send": {
		Name:        "telnet send",
		Description: ": None",
	},
	"telnet status": {
		Name:        "telnet status",
		Description: ": None",
	},
	"telnet set": {
		Name:        "telnet set",
		Description: ": None",
	},
	"telnet unset": {
		Name:        "telnet unset",
		Description: ": None",
	},
	"telnet toggle": {
		Name:        "telnet toggle",
		Description: ": None",
	},
	"telnet mode": {
		Name:        "telnet mode",
		Description: ": None",
	},
	"telnet display": {
		Name:        "telnet display",
		Description: ": None",
	},
	"telnet environ": {
		Name:        "telnet environ",
		Description: ": None",
	},
	"nc": {
		Name:        "nc",
		Description: ": -l (Listen mode, for inbound connects), -p (Local port number), -e (Program to execute after connect), -n (Numeric-only IP addresses, no DNS), -u (UDP mode), -v (Verbose mode), -z (Zero-I/O mode (used for scanning)), -w (Timeout for connects and final net reads), -i (Delay interval for lines sent), -k (Keep inbound sockets open for multiple connects), -s (Local source address), -o (Hex dump of traffic), -r (Randomize remote ports), -q (Quit after EOF on stdin and delay of secs)",
	},
	"nmap": {
		Name:        "nmap",
		Description: ": -sP (Perform a ping scan to determine which hosts are up), -sT (Perform a TCP connect scan), -sS (TCP SYN scan), -sU (UDP scan), -sV (Version detection), -O (OS detection), -A (Aggressive scan options), -p (Specify ports to scan), -T (Set timing template), -oN (Normal output to file), -oX (XML output to file), -oG (Grepable output to file), -oA (Output in the three major formats at once), -iL (Input from list of hosts/networks), -iR (Scan random hosts), -Pn (Treat all hosts as online -- skip host discovery), -6 (Enable IPv6 scanning), -sC (Run default scripts), -script (Specify custom scripts to run)",
	},
	"tcpdump": {
		Name:        "tcpdump",
		Description: ": -i (Specify the interface to listen on), -w (Write the raw packets to a file), -r (Read packets from a file), -c (Exit after receiving a specified number of packets), -s (Set the snapshot length), -v (Verbose output), -vv (More verbose output), -vvv (Even more verbose output), -e (Print link-level header on each dump line), -q (Quick (quiet?) output), -X (Print packet data in both hex and ASCII), -XX (Print packet data in hex and ASCII, including link-level header), -A (Print each packet (minus its link level header) in ASCII), -D (Print the list of available interfaces), -l (Make stdout line buffered), -tt (Print an unformatted timestamp on each dump line), -ttt (Print a delta (micro-second resolution) between current and previous line on each dump line), -tttt (Print a timestamp in default format proceeded by date on each dump line), -ttttt (Print a delta (micro-second resolution) between current and first line on each dump line), -C (Before writing a raw packet to a file, check whether the file is larger than file_size and, if so, close the current file and open a new one), -G (Rotate dump files every specified number of seconds), -W (Used in conjunction with the -C or -G options, limits the number of files created to specified number), -Z (Drop privileges to user and group after opening the capture file), -K (Do not attempt to verify IP, TCP, or UDP checksums), -E (Decrypt IPsec ESP packets using provided key), -M (Use the specified secret for IPsec AH authentication)",
	},
	"wireshark": {
		Name:        "wireshark",
		Description: ": -i (Specify the interface to listen on), -k (Start capturing immediately), -w (Write the raw packets to a file), -r (Read packets from a file), -c (Exit after receiving a specified number of packets), -f (Set the capture filter expression), -s (Set the snapshot length), -v (Verbose output), -h (Display help information), -b (Set ring buffer options), -t (Set the time stamp format), -n (Disable network object name resolution), -N (Set name resolution flags), -S (Update packet list in real-time), -T (Set the format of text output), -X (Specify an option to pass to the plugin), -Y (Set the display filter expression), -z (Specify statistics to calculate and display)",
	},
	"iptables": {
		Name:        "iptables",
		Description: ": -A (Append a rule to a chain), -D (Delete a rule from a chain)",
	},
	"firewalld": {
		Name:        "firewalld",
		Description: ": --add-service (Add a service to the firewall), --remove-service (Remove a service from the firewall)  --state, Check the current state of firewalld, --reload, Reload firewalld configuration, --add-port, Add a port to the firewall, --remove-port, Remove a port from the firewall, --list-all, List all firewall rules, --zone, Manage firewall zones,--permanent, Make changes permanent",
	},
	"ethtool": {
		Name:        "ethtool",
		Description: ": -s (Change network device settings), -i (Show driver information for a network device)",
	},
	"route": {
		Name:        "route",
		Description: ": add (Add a new route), del (Delete an existing route)",
	},
	"arp": {
		Name:        "arp",
		Description: ": -a (Display the ARP table), -d (Delete an entry from the ARP table)",
	},
    "whois": {
        Name:        "whois",
        Description: ": -h (Specify the whois server to query), -p (Specify the port to connect to on the whois server)",
	},
	"ipcalc": {
		Name:        "ipcalc",
		Description: ": -n (Show the network address), -b (Show the broadcast address)",
	},
	"netcat": {
		Name:        "netcat",
		Description: ": -l (Enable listen mode for inbound connections), -u (Use UDP mode instead of the default TCP)",
	},
	"arping": {
		Name:        "arping",
		Description: ": -c (Specify the count of ARP requests to send), -I (Specify the network interface to use)",
	},
	"iw": {
		Name:        "iw",
		Description: ": dev (Show information about wireless devices), link (Show information about wireless links)",
	},
	"bridge": {
		Name:        "bridge",
		Description: ": fdb (Manage the forwarding database), vlan (Manage VLANs)",
	},
	"tc": {
		Name:        "tc",
		Description: ": qdisc (Manage queueing disciplines), filter (Manage traffic filters)",
	},
	"ipset": {
		Name:        "ipset",
		Description: ": create (Create a new IP set), add (Add an entry to an IP set)",
	},
	"ip rule": {
		Name:        "ip rule",
		Description: ": add (Add a new routing rule), del (Delete an existing routing rule)",
	},
	"ip addr": {
		Name:        "ip addr",
		Description: ": add (Add a new IP address), del (Delete an existing IP address)",
	},
	"ip link": {
		Name:        "ip link",
		Description: ": set (Change device attributes), show (Display device attributes)",
	},
	"ip route": {
		Name:        "ip route",
		Description: ": add (Add a new route), del (Delete an existing route)",
	},
	"ip neigh": {
		Name:        "ip neigh",
		Description: ": add (Add a new neighbor entry), del (Delete an existing neighbor entry)",
	},
	"ip maddr": {
		Name:        "ip maddr",
		Description: ": add (Add a new multicast address), del (Delete an existing multicast address)",
	},
	"ip monitor": {
		Name:        "ip monitor",
		Description: ": all (Monitor all events), link (Monitor link events)",
	},
	"ip tunnel": {
		Name:        "ip tunnel",
		Description: ": add (Add a new tunnel), del (Delete an existing tunnel)",
	},
	"ip xfrm": {
		Name:        "ip xfrm",
		Description: ": state (Manage IPsec state), policy (Manage IPsec policy)",
	},
	"ip netns": {
		Name:        "ip netns",
		Description: ": add (Add a new network namespace), del (Delete an existing network namespace)",
	},
	"ip vrf": {
		Name:        "ip vrf",
		Description: ": add (Add a new VRF (Virtual Routing and Forwarding)), del (Delete an existing VRF)",
	},
	"ip link set": {
		Name:        "ip link set",
		Description: ": up (Enable a network device), down (Disable a network device), mtu (Change the Maximum Transmission Unit (MTU) of a device), address (Change the MAC address of a device), master (Set the master device), nomaster (Unset the master device), type (Change the type of a device), netns (Move a device to a different network namespace), alias (Set an alias for a device), group (Set the group of a device), txqueuelen (Change the transmit queue length of a device), vf (Change the attributes of a virtual function), xdp (Attach or detach an XDP (eXpress Data Path) program)",
	},
	"ip link add": {
		Name:        "ip link add",
		Description: ": type (Specify the type of device to add (e.g., bridge, bond, dummy, etc.)), name (Specify the name of the device to add)",
	},
	"ip link delete": {
		Name:        "ip link delete",
		Description: ": dev (Specify the name of the device to delete)",
	},
	"ip link show": {
		Name:        "ip link show",
		Description: ": dev (Specify the name of the device to show)",
	},
	"ip link set dev": {
		Name:        "ip link set dev",
		Description: ": mtu (Change the Maximum Transmission Unit (MTU) of the specified device), address (Change the MAC address of the specified device)",
	},
	"ip link set up": {
		Name:        "ip link set up",
		Description: ": (Enable the specified device)",
	},
	"ip link set down": {
		Name:        "ip link set down",
		Description: ": (Disable the specified device)",
	},
	"ip link set mtu": {
		Name:        "ip link set mtu",
		Description: ": (Change the Maximum Transmission Unit (MTU) of the specified device)",
	},
	"ip link set address": {
		Name:        "ip link set address",
		Description: ": (Change the MAC address of the specified device)",
	},
	"ip link set master": {
		Name:        "ip link set master",
		Description: ": (Set the master device for the specified device)",
	},
	"ip link set nomaster": {
		Name:        "ip link set nomaster",
		Description: ": (Unset the master device for the specified device)",
	},
	"ip link set type": {
		Name:        "ip link set type",
		Description: ": (Change the type of the specified device)",
	},
	"ip link set netns": {
		Name:        "ip link set netns",
		Description: ": (Move the specified device to a different network namespace)",
	},
	"ip link set alias": {
		Name:        "ip link set alias",
		Description: ": (Set an alias for the specified device)",
	},
	"ip link set group": {
		Name:        "ip link set group",
		Description: ": (Set the group for the specified device)",
	},
	"ip link set txqueuelen": {
		Name:        "ip link set txqueuelen",
		Description: ": (Change the transmit queue length of the specified device)",
	},
	"ip link set vf": {
		Name:        "ip link set vf",
		Description: ": (Change the attributes of a virtual function for the specified device)",
	},
	"ip link set xdp": {
		Name:        "ip link set xdp",
		Description: ": (Attach or detach an XDP (eXpress Data Path) program to/from the specified device)",
	},
	"docker run": {
		Name:        "docker run",
		Description: ": -d (Run container in background and print container ID), -p (Publish a container's port to the host), -v (Bind mount a volume)",
	},
	"docker start": {
		Name:        "docker start",
		Description: ": -a (Attach STDOUT/STDERR and forward signals), -i (Attach container's STDIN)",
	},
	"docker stop": {
		Name:        "docker stop",
		Description: ": -t (Seconds to wait for stop before killing it)",
	},
	"docker restart": {
		Name:        "docker restart",
		Description: ": -t (Seconds to wait for stop before killing it)",
	},
	"docker rm": {
		Name:        "docker rm",
		Description: ": -f (Force the removal of a running container), -v (Remove the volumes associated with the container)",
	},
	"docker rmi": {
		Name:        "docker rmi",
		Description: ": -f (Force removal of the image)",
	},
	"docker images": {
		Name:        "docker images",
		Description: ": -a (Show all images), -q (Only show numeric IDs)",
	},
	"docker ps": {
		Name:        "docker ps",
		Description: ": -a (Show all containers), -q (Only display numeric IDs)",
	},
	"docker exec": {
		Name:        "docker exec",
		Description: ": -d (Detached mode: run command in the background), -i (Keep STDIN open even if not attached), -t (Allocate a pseudo-TTY)",
	},
	"docker build": {
		Name:        "docker build",
		Description: ": -t (Name and optionally a tag in the 'name:tag' format), -f (Name of the Dockerfile)",
	},
	"docker pull": {
		Name:        "docker pull",
		Description: ": -a (Download all tagged images in the repository)",
	},
	"docker push": {
		Name:        "docker push",
		Description: ": -a (Push all tagged images in the repository)",
	},
	"docker login": {
		Name:        "docker login",
		Description: ": -u (Username), -p (Password)",
	},
	"docker logout": {
		Name:        "docker logout",
		Description: ": None",
	},
	"docker tag": {
		Name:        "docker tag",
		Description: ": None",
	},
	"docker inspect": {
		Name:        "docker inspect",
		Description: ": -f (Format the output using a Go template)",
	},
	"docker logs": {
		Name:        "docker logs",
		Description: ": -f (Follow log output), -t (Show timestamps)",
	},
	"docker commit": {
		Name:        "docker commit",
		Description: ": -m (Commit message), -a (Author)",
	},
	"docker network": {
		Name:        "docker network",
		Description: ": None",
	},
	"docker volume": {
		Name:        "docker volume",
		Description: ": None",
	},
	"docker info": {
		Name:        "docker info",
		Description: ": None",
	},
	"docker version": {
		Name:        "docker version",
		Description: ": None",
	},
	"docker swarm": {
		Name:        "docker swarm",
		Description: ": None",
	},
	"docker node": {
		Name:        "docker node",
		Description: ": None",
	},
	"docker service": {
		Name:        "docker service",
		Description: ": None",
	},
	"docker stack": {
		Name:        "docker stack",
		Description: ": None",
	},
	"docker secret": {
		Name:        "docker secret",
		Description: ": None",
	},
	"docker config": {
		Name:        "docker config",
		Description: ": None",
	},
	"docker plugin": {
		Name:        "docker plugin",
		Description: ": None",
	},
	"docker container": {
		Name:        "docker container",
		Description: ": None",
	},
	"docker image": {
		Name:        "docker image",
		Description: ": None",
	},
	"docker system": {
		Name:        "docker system",
		Description: ": None",
	},
	"docker context": {
		Name:        "docker context",
		Description: ": None",
	},
	"docker builder": {
		Name:        "docker builder",
		Description: ": None",
	},
	"docker checkpoint": {
		Name:        "docker checkpoint",
		Description: ": None",
	},
	"docker trust": {
		Name:        "docker trust",
		Description: ": None",
	},
	"systemctl start": {
		Name:        "systemctl start",
		Description: ": --no-block (Do not wait until operation finished), --quiet (Suppress output)",
	},
	"systemctl stop": {
		Name:        "systemctl stop",
		Description: ": --no-block (Do not wait until operation finished), --quiet (Suppress output)",
	},
	"systemctl restart": {
		Name:        "systemctl restart",
		Description: ": --no-block (Do not wait until operation finished), --quiet (Suppress output)",
	},
	"systemctl reload": {
		Name:        "systemctl reload",
		Description: ": --no-block (Do not wait until operation finished), --quiet (Suppress output)",
	},
	"systemctl enable": {
		Name:        "systemctl enable",
		Description: ": --now (Start the unit after enabling it)",
	},
	"systemctl disable": {
		Name:        "systemctl disable",
		Description: ": --now (Stop the unit after disabling it)",
	},
	"systemctl status": {
		Name:        "systemctl status",
		Description: ": --no-pager (Do not pipe output into a pager), --full (Do not ellipsize unit names)",
	},
	"systemctl is-active": {
		Name:        "systemctl is-active",
		Description: ": --quiet (Suppress output)",
	},
	"systemctl is-enabled": {
		Name:        "systemctl is-enabled",
		Description: ": --quiet (Suppress output)",
	},
	"systemctl list-units": {
		Name:        "systemctl list-units",
		Description: ": --all (Show all loaded units, regardless of their state), --type (List only units of a particular type)",
	},
	"systemctl list-unit-files": {
		Name:        "systemctl list-unit-files",
		Description: ": --state (Show only unit files in the specified state)",
	},
	"systemctl daemon-reload": {
		Name:        "systemctl daemon-reload",
		Description: ": None",
	},
	"systemctl mask": {
		Name:        "systemctl mask",
		Description: ": --runtime (Mask only temporarily until the next reboot)",
	},
	"systemctl unmask": {
		Name:        "systemctl unmask",
		Description: ": --runtime (Unmask only temporarily until the next reboot)",
	},
	"systemctl isolate": {
		Name:        "systemctl isolate",
		Description: ": --no-block (Do not wait until operation finished), --quiet (Suppress output)",
	},
	"systemctl kill": {
		Name:        "systemctl kill",
		Description: ": --signal (Specify the signal to send)",
	},
	"systemctl show": {
		Name:        "systemctl show",
		Description: ": --property (Show only properties specified)",
	},
	"systemctl cat": {
		Name:        "systemctl cat",
		Description: ": None",
	},
	"systemctl edit": {
		Name:        "systemctl edit",
		Description: ": --full (Edit the full unit file instead of creating a drop-in snippet)",
	},
	"systemctl set-property": {
		Name:        "systemctl set-property",
		Description: ": None",
	},
	"systemctl help": {
		Name:        "systemctl help",
		Description: ": None",
	},
	"journalctl --boot": {
		Name:        "journalctl --boot",
		Description: ": -k (Show only kernel messages), -p (Show messages with the specified priority)",
	},
	"journalctl --list-boots": {
		Name:        "journalctl --list-boots",
		Description: ": None",
	},
	"journalctl --unit": {
		Name:        "journalctl --unit",
		Description: ": -f (Follow new messages), -n (Show the specified number of most recent messages)",
	},
	"journalctl --since": {
		Name:        "journalctl --since",
		Description: ": None",
	},
	"journalctl --until": {
		Name:        "journalctl --until",
		Description: ": None",
	},
	"journalctl --follow": {
		Name:        "journalctl --follow",
		Description: ": None",
	},
	"journalctl --output": {
		Name:        "journalctl --output",
		Description: ": -o (Specify the output format)",
	},
	"journalctl --priority": {
		Name:        "journalctl --priority",
		Description: ": None",
	},
	"journalctl --grep": {
		Name:        "journalctl --grep",
		Description: ": None",
	},
	"systemd-analyze blame": {
		Name:        "systemd-analyze blame",
		Description: ": --no-pager (Do not pipe output into a pager)",
	},
	"systemd-analyze critical-chain": {
		Name:        "systemd-analyze critical-chain",
		Description: ": --fuzz (Specify the fuzz factor for the critical chain)",
	},
	"systemd-analyze plot": {
		Name:        "systemd-analyze plot",
		Description: ": --file (Specify the output file for the SVG plot)",
	},
	"systemd-analyze dump": {
		Name:        "systemd-analyze dump",
		Description: ": None",
	},
	"systemd-analyze verify": {
		Name:        "systemd-analyze verify",
		Description: ": --man (Show the man page for the unit file)",
	},
	"systemd-analyze security": {
		Name:        "systemd-analyze security",
		Description: ": --no-pager (Do not pipe output into a pager)",
	},
	"systemd-analyze time": {
		Name:        "systemd-analyze time",
		Description: ": None",
	},
	"cron start": {
		Name:        "cron start",
		Description: ": None",
	},
	"cron stop": {
		Name:        "cron stop",
		Description: ": None",
	},
	"cron restart": {
		Name:        "cron restart",
		Description: ": None",
	},
	"cron reload": {
		Name:        "cron reload",
		Description: ": None",
	},
	"cron status": {
		Name:        "cron status",
		Description: ": None",
	},
	"cron list": {
		Name:        "cron list",
		Description: ": None",
	},
	"cron add": {
		Name:        "cron add",
		Description: ": None",
	},
	"cron remove": {
		Name:        "cron remove",
		Description: ": None",
	},
	"cron edit": {
		Name:        "cron edit",
		Description: ": None",
	},
	"ufw enable": {
		Name:        "ufw enable",
		Description: ": None",
	},
	"ufw disable": {
		Name:        "ufw disable",
		Description: ": None",
	},
	"ufw status": {
		Name:        "ufw status",
		Description: ": None",
	},
	"ufw allow": {
		Name:        "ufw allow",
		Description: ": None",
	},
	"ufw deny": {
		Name:        "ufw deny",
		Description: ": None",
	},
	"ufw reject": {
		Name:        "ufw reject",
		Description: ": None",
	},
	"ufw limit": {
		Name:        "ufw limit",
		Description: ": None",
	},
	"ufw delete": {
		Name:        "ufw delete",
		Description: ": None",
	},
	"ufw reset": {
		Name:        "ufw reset",
		Description: ": None",
	},
	"ufw reload": {
		Name:        "ufw reload",
		Description: ": None",
	},
	"ufw logging": {
		Name:        "ufw logging",
		Description: ": None",
	},
	"nmcli general": {
		Name:        "nmcli general",
		Description: ": None",
	},
	"nmcli networking": {
		Name:        "nmcli networking",
		Description: ": None",
	},
	"nmcli radio": {
		Name:        "nmcli radio",
		Description: ": None",
	},
	"nmcli connection": {
		Name:        "nmcli connection",
		Description: ": None",
	},
	"nmcli device": {
		Name:        "nmcli device",
		Description: ": None",
	},
	"nmcli agent": {
		Name:        "nmcli agent",
		Description: ": None",
	},
	"nmcli monitor": {
		Name:        "nmcli monitor",
		Description: ": None",
	},
	"loginctl list-sessions": {
		Name:        "loginctl list-sessions",
		Description: ": None",
	},
	"loginctl session-status": {
		Name:        "loginctl session-status",
		Description: ": None",
	},
	"loginctl terminate-session": {
		Name:        "loginctl terminate-session",
		Description: ": None",
	},
	"loginctl list-users": {
		Name:        "loginctl list-users",
		Description: ": None",
	},
	"loginctl user-status": {
		Name:        "loginctl user-status",
		Description: ": None",
	},
	"loginctl enable-linger": {
		Name:        "loginctl enable-linger",
		Description: ": None",
	},
	"loginctl disable-linger": {
		Name:        "loginctl disable-linger",
		Description: ": None",
	},
	"loginctl lock-session": {
		Name:        "loginctl lock-session",
		Description: ": None",
	},
	"loginctl unlock-session": {
		Name:        "loginctl unlock-session",
		Description: ": None",
	},
	"loginctl lock-sessions": {
		Name:        "loginctl lock-sessions",
		Description: ": None",
	},
	"loginctl unlock-sessions": {
		Name:        "loginctl unlock-sessions",
		Description: ": None",
	},
	"localectl status": {
		Name:        "localectl status",
		Description: ": None",
	},
	"localectl set-locale": {
		Name:        "localectl set-locale",
		Description: ": None",
	},
	"localectl list-locales": {
		Name:        "localectl list-locales",
		Description: ": None",
	},
	"localectl set-keymap": {
		Name:        "localectl set-keymap",
		Description: ": None",
	},
	"localectl list-keymaps": {
		Name:        "localectl list-keymaps",
		Description: ": None",
	},
	"hostnamectl status": {
		Name:        "hostnamectl status",
		Description: ": None",
	},
	"hostnamectl set-hostname": {
		Name:        "hostnamectl set-hostname",
		Description: ": None",
	},
	"hostnamectl set-icon-name": {
		Name:        "hostnamectl set-icon-name",
		Description: ": None",
	},
	"hostnamectl set-chassis": {
		Name:        "hostnamectl set-chassis",
		Description: ": None",
	},
	"hostnamectl set-deployment": {
		Name:        "hostnamectl set-deployment",
		Description: ": None",
	},
	"hostnamectl set-location": {
		Name:        "hostnamectl set-location",
		Description: ": None",
	},
	"timedatectl status": {
		Name:        "timedatectl status",
		Description: ": None",
	},
	"timedatectl set-time": {
		Name:        "timedatectl set-time",
		Description: ": None",
	},
	"timedatectl set-timezone": {
		Name:        "timedatectl set-timezone",
		Description: ": None",
	},
	"timedatectl list-timezones": {
		Name:        "timedatectl list-timezones",
		Description: ": None",
	},
	"timedatectl set-ntp": {
		Name:        "timedatectl set-ntp",
		Description: ": None",
	},
}