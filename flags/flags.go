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
		Description: ": -n (Do not output the trailing newline), -e (Enable interpretation of backslash escapes), -E (Disable interpretation of backslash escapes), --help (Display help information), --version (Output version information)",
	},
    "man": {
		Name:        "man",
		Description: ": -k (Search the short descriptions and manual page names for the keyword), -f (Display only the first line of each entry), -a (Display all matching manual pages), -w (List the location of the manual page files that would be displayed), -P (Specify which pager to use), -M (Specify the list of directories to search for manual pages), -L (Specify the locale for the manual pages), -C (Specify the configuration file to use), -7 (Display manual pages in ASCII format), -B (Specify the web browser to use for displaying HTML manual pages), -H (Display manual pages in HTML format), -X (Display manual pages in X Window System format)",
	},
    "grep": {
		Name:        "grep",
		Description: ": -i (Ignore case), -r, -R (Recurse directories), -v (Invert match), -E, -F, -G, -P (Pattern types), -w (Whole words), -x (Whole lines), -c (Count matches), -l, -L (List files), -n (Line numbers), -H, -h (File names), -o (Matched parts), -q (Quiet), -s (Suppress errors), -b (Byte offset), -d (Handle directories), -a, -I (Binary files), -z (Zero byte lines)",
	},
    "find": {
		Name:        "find",
		Description: ": -name (Match shell pattern), -type (File type), -exec (Execute command), -iname (Case-insensitive name), -path (Match shell pattern path), -ipath (Case-insensitive path), -regex (Match regex), -iregex (Case-insensitive regex), -size (File size), -perm (Permission bits), -user (File owner), -group (File group), -mtime (Modified n*24 hours ago), -atime (Accessed n*24 hours ago), -ctime (Status changed n*24 hours ago), -newer (Modified more recently than reference), -delete (Delete files), -print (Print file name), -ls (List in ls format)",
	},
    "chmod": {
		Name:        "chmod",
		Description: ": -R (Change files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made), --reference=RFILE (Use RFILE's mode instead of MODE values), --preserve-root (Fail to operate recursively on /), --no-preserve-root (Do not treat / specially), --help (Display help information), --version (Output version information)",
	},
    "chown": {
		Name:        "chown",
		Description: ": -R (Operate on files and directories recursively), -v (Output a diagnostic for every file processed), -c (Like verbose but report only when a change is made), --dereference (Affect the referent of each symbolic link), -h (Affect symbolic links instead of any referenced file), --from=CURRENT_OWNER:CURRENT_GROUP (Change the owner and/or group only if its current owner and/or group match those specified), --preserve-root (Fail to operate recursively on /), --no-preserve-root (Do not treat / specially), --reference=RFILE (Use RFILE's owner and group), --help (Display help information), --version (Output version information)",
	},
    "df": {
		Name:        "df",
		Description: ": -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -T (Print file system type), -i (List inode information instead of block usage), -a (Include dummy file systems), -B (Scale sizes by SIZE before printing them), --total (Produce a grand total)",
	},
	"du": {
		Name:        "du",
		Description: ": -h (Print sizes in human-readable format (e.g., 1K 234M 2G)), -s (Display only a total for each argument), -a (Write counts for all files, not just directories), -c (Produce a grand total), -d (Print the total for a directory (or file) only if it is N or fewer levels below the command line argument), -L (Dereference all symbolic links), -m (Print sizes in megabytes)",
	},
	"head": {
		Name:        "head",
		Description: ": -n (Print the first NUM lines instead of the first 10), -c (Print the first NUM bytes of each file), -q (Never print headers giving file names), -v (Always print headers giving file names), -z (Line delimiter is NUL instead of newline)",
	},
	"tail": {
		Name:        "tail",
		Description: ": -n (Output the last NUM lines, instead of the last 10), -f (Output appended data as the file grows), -c (Output the last NUM bytes), -q (Never output headers giving file names), -v (Always output headers giving file names), --pid=PID (Terminate after process ID, PID dies), -s (Sleep for N seconds between iterations)",
	},
	"tar": {
		Name:        "tar",
		Description: ": -c (Create a new archive), -x (Extract files from an archive), -v (Verbally list files processed), -f (Use archive file or device ARCHIVE), -z (Filter the archive through gzip), -j (Filter the archive through bzip2), -J (Filter the archive through xz), -C (Change to directory DIR)",
	},
	"wget": {
		Name:        "wget",
		Description: ": -O (Write documents to FILE), -q (Turn off Wget's output), -r (Specify recursive download), -N (Turn on timestamping), -P (Directory prefix where files will be saved), -A (Specify comma-separated list of accepted extensions), -R (Specify comma-separated list of rejected extensions), -l (Specify recursion level), --limit-rate (Limit download speed), --no-check-certificate (Don't validate the server's certificate)",
	},
	"curl": {
		Name:        "curl",
		Description: ": -X (HTTP method), -I (Headers only), -d (POST data), -F (Form data), -H (Custom headers), -u (User:password), -o (Output file), -O (Save as remote file), -L (Follow redirects), -k (Insecure SSL), -s (Silent mode), -v (Verbose), -C - (Resume transfer), -b (Send cookies), -c (Save cookies), -T (Upload file), --proxy (Proxy), --limit-rate (Limit rate), --max-time (Max time), --compressed (Compressed response), --data-urlencode (URL encode data), --http2 (HTTP/2), --http3 (HTTP/3), --cert (Client cert), --cacert (CA cert), --capath (CA cert dir)",
	},
	"ps": {
		Name:        "ps",
		Description: ": -e (Select all processes), -f (Do full-format listing), -u (Display processes for a specific user), -a (Select all processes except session leaders and processes not associated with a terminal), -x (Select processes without controlling ttys), -o (User-defined format), -p (Select by PID), --sort (Specify sorting order)",
	},
	"kill": {
		Name:        "kill",
		Description: ": -s (Specify the signal to send), -l (List all signal names), -p (Specify the process ID), -L (List signal names and numbers), -n (Specify signal number)",
	},
	"top": {
		Name:        "top",
		Description: ": -b (Batch mode operation), -n (Number of iterations), -u (Display processes for a specific user), -d (Delay between updates), -p (Monitor specific PIDs), -c (Show command line instead of command name), -H (Show threads)",
	},
	"htop": {
		Name:        "htop",
		Description: ": -d (Delay between updates, in tenths of seconds), -u (Display processes for a specific user), -p (Show only the given PIDs), -s (Sort by a specified column), -C (Disable color), -t (Tree view)",
	},
	"ssh": {
		Name:        "ssh",
		Description: ": -l (Specify the user to log in as), -p (Port to connect to on the remote host), -i (File from which the identity (private key) is read), -o (Specify options), -v (Verbose mode), -q (Quiet mode), -C (Enable compression), -N (Do not execute a remote command), -T (Disable pseudo-terminal allocation)",
	},
	"scp": {
		Name:        "scp",
		Description: ": -r (Recursively copy entire directories), -P (Specify the port to connect to on the remote host), -i (File from which the identity (private key) is read), -v (Verbose mode), -q (Quiet mode), -C (Enable compression), -l (Limit the bandwidth, specified in Kbit/s)",
	},
	"nano": {
		Name:        "nano",
		Description: ": -B (Make a backup of the file), -m (Enable mouse support), -i (Automatically indent new lines to the same position as the previous line), -c (Constantly show the cursor position), -l (Display line numbers), -E (Convert typed tabs to spaces), -S (Smooth scrolling)",
	},
	"vim": {
		Name:        "vim",
		Description: ": -u (Use the specified vimrc file), -N (No compatible mode), -e (Start in Ex mode), -s (Silent mode), -R (Read-only mode), -Z (Restricted mode), -y (Start in easy mode), -d (Start in diff mode)",
	},
	"apt-get": {
		Name:        "apt-get",
		Description: ": -y (Assume yes to all prompts and run non-interactively), -q (Quiet; produces output suitable for logging), -d (Download only; package files are only retrieved, not unpacked or installed), -s (No-act; perform a simulation of events that would occur but do not actually change the system), -f (Attempt to correct a system with broken dependencies), --purge (Remove packages and their configuration files), --reinstall (Reinstall the specified packages), --ignore-missing (Ignore missing packages; do not fail)",
	},
	"yum": {
		Name:        "yum",
		Description: ": -y (Assume yes to all prompts), -q (Quiet mode), -v (Verbose mode), -C (Run entirely from cache; do not update the cache), --enablerepo (Enable additional repositories), --disablerepo (Disable repositories), --nogpgcheck (Disable GPG signature checking), --noplugins (Disable all plugins)",
	},
	"mount": {
		Name:        "mount",
		Description: ": -t (Indicate the filesystem type), -o (Specify mount options), -v (Verbose mode), -a (Mount all filesystems mentioned in fstab), -r (Mount the filesystem read-only), -w (Mount the filesystem read-write), -L (Mount by label), -U (Mount by UUID)",
	},
	"umount": {
		Name:        "umount",
		Description: ": -f (Force unmount), -l (Lazy unmount), -v (Verbose mode), -a (Unmount all filesystems mentioned in fstab), -r (Remount the filesystem read-only), -d (Detach the loop device associated with the mount)",
	},
    "uname": {
		Name:        "uname",
		Description: ": -a (Print all information), -r (Print the kernel release), -v (Print the kernel version), -s (Print the kernel name), -n (Print the network node hostname), -m (Print the machine hardware name), -p (Print the processor type), -i (Print the hardware platform), -o (Print the operating system)",
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
		Description: ": -i (Display the IP address), -f (Display the FQDN (Fully Qualified Domain Name)), -d (Display the domain name), -s (Display the short hostname), -A (Display all FQDNs of the machine), -I (Display all network addresses of the host), -b (Display the hostname as a boot parameter)",
	},
	"history": {
		Name:        "history",
		Description: ": -c (Clear the history list), -d (Delete the history entry at position offset), -a (Append the new history lines to the history file), -r (Read the history file and append its contents to the history list), -w (Write the current history to the history file), -n (Read the history file and append only the lines not already read), -p (Perform history expansion on the following args and display the result without storing it in the history list), -s (Append the arguments to the history list as a single entry)",
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
		Description: ": -u (Remove variable from the environment)",
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
		Description: ": -c (Pass command to the shell with -c), -l (Make the shell a login shell), -s (Run the specified shell instead of the default), -m (Do not reset environment variables), -P (Do not reset environment variables), -g (Specify the primary group), -G (Specify supplementary groups), - (Make the shell a login shell)",
	},
	"passwd": {
		Name:        "passwd",
		Description: ": -d (Delete a user's password), -e (Expire a user's password), -l (Lock a user's password), -u (Unlock a user's password), -n (Set the minimum number of days between password changes), -x (Set the maximum number of days between password changes), -w (Set the number of days of warning before a password change is required), -i (Set the number of days after a password expires until the account is disabled)",
	},
	"useradd": {
		Name:        "useradd",
		Description: ": -m (Create the user's home directory), -s (Specify the user's login shell), -G (Specify supplementary groups for the user), -u (Specify the user ID), -d (Specify the user's home directory), -c (Add a comment for the user), -e (Specify the account expiration date), -f (Specify the number of days after a password expires until the account is permanently disabled)",
	},
	"usermod": {
		Name:        "usermod",
		Description: ": -l (Change the user's login name), -G (Specify supplementary groups for the user), -s (Specify the user's login shell), -u (Change the user ID), -d (Change the user's home directory), -c (Change the comment for the user), -e (Change the account expiration date), -f (Change the number of days after a password expires until the account is permanently disabled)",
	},
	"userdel": {
		Name:        "userdel",
		Description: ": -r (Remove the user's home directory and mail spool), -f (Force removal of the user), -Z (Remove any SELinux user mapping for the user)",
	},
	"groupadd": {
		Name:        "groupadd",
		Description: ": -g (Specify the group ID), -r (Create a system group), -f (Exit with success status if the specified group already exists)",
	},
	"groupmod": {
		Name:        "groupmod",
		Description: ": -n (Change the name of the group), -g (Change the group ID), -o (Allow using duplicate (non-unique) GID)",
	},
	"groupdel": {
		Name:        "groupdel",
		Description: ": No flags available",
	},
	"crontab": {
		Name:        "crontab",
		Description: ": -e (Edit the current crontab), -l (List the current crontab), -r (Remove the current crontab), -u (Specify the user whose crontab is to be manipulated)",
	},
	"at": {
		Name:        "at",
		Description: ": -f (Read the job from the specified file), -m (Send mail to the user when the job has completed), -v (Verbose mode), -q (Specify the queue to use), -t (Specify the time to run the job)",
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
		Description: ": -i (Ask for confirmation before killing each process), -v (Report if the signal was successfully sent), -w (Wait for all processes to die), -q (Do not complain if no processes were killed), -e (Require an exact match for very long names), -u (Kill only processes the specified user owns)",
	},
	"xargs": {
		Name:        "xargs",
		Description: ": -n (Use at most max-args arguments per command line), -P (Run up to max-procs processes at a time), -I (Replace occurrences of replace-str in the initial arguments with names read from standard input), -t (Print the command line on the standard error output before executing it), -r (Do not run the command if the standard input is empty), -0 (Input items are terminated by a null character instead of by whitespace)",
	},
	"tee": {
		Name:        "tee",
		Description: ": -a (Append to the given files, do not overwrite), -i (Ignore interrupts), -p (Write to the output file as well as the standard output)",
	},
	"diff": {
		Name:        "diff",
		Description: ": -u (Use the unified output format), -c (Use the context output format), -i (Ignore case differences in file contents), -r (Recursively compare any subdirectories found), -N (Treat absent files as empty), -q (Output only whether files differ)",
	},
	"patch": {
		Name:        "patch",
		Description: ": -p (Strip the smallest prefix containing num leading slashes from each file name found in the patch file), -R (Assume the patch was already applied and attempt to un-apply it), -N (Ignore patches that seem to be already applied), -b (Make a backup before applying the patch), -d (Change to directory DIR before applying the patch), -i (Read the patch from the specified file)",
	},
	"sed": {
		Name:        "sed",
		Description: ": -n (Suppress automatic printing of pattern space), -e (Add the script to the commands to be executed), -i (Edit files in place (makes backup if suffix supplied)), -r (Use extended regular expressions in the script), -f (Add the contents of script-file to the commands to be executed)",
	},
	"awk": {
		Name:        "awk",
		Description: ": -F (Use the specified field separator), -v (Assign the specified variable), -f (Read the awk program source from the specified file), -W (Enable compatibility mode), -b (Use binary mode for all I/O operations)",
	},
	"tr": {
		Name:        "tr",
		Description: ": -d (Delete characters in the first set from the input), -s (Replace each input sequence of a repeated character that is listed in the last specified set with a single occurrence of that character), -c (Complement the set of characters in string1), -t (Truncate the first set to the length of the second set)",
	},
	"cut": {
		Name:        "cut",
		Description: ": -f (Select only these fields), -d (Use the specified delimiter), -c (Select only these characters), -s (Suppress lines with no delimiter character)",
	},
	"sort": {
		Name:        "sort",
		Description: ": -r (Reverse the result of comparisons), -n (Compare according to string numerical value), -k (Sort via a key), -t (Use the specified character as the field separator), -u (Output only the first of an equal run), -o (Write result to the specified file)",
	},
	"uniq": {
		Name:        "uniq",
		Description: ": -c (Prefix lines by the number of occurrences), -d (Only print duplicate lines), -u (Only print unique lines), -i (Ignore differences in case when comparing), -f (Skip fields when comparing), -s (Skip characters when comparing), -w (Compare no more than n characters in lines)",
	},
	"wc": {
		Name:        "wc",
		Description: ": -l (Print the newline counts), -w (Print the word counts), -c (Print the byte counts), -m (Print the character counts), -L (Print the length of the longest line)",
	},
	"basename": {
		Name:        "basename",
		Description: ": -s (Remove a trailing suffix), -a (Support multiple arguments and treat each as a NAME), -z (End each output line with NUL, not newline)",
	},
	"dirname": {
		Name:        "dirname",
		Description: ": -z (End each output line with NUL, not newline)",
	},
	"readlink": {
		Name:        "readlink",
		Description: ": -f (Canonicalize by following every symlink in every component of the given name recursively), -e (Same as -f, but fail if any component is missing or not a directory), -m (Canonicalize by following every symlink in every component of the given name recursively, but do not fail if any component is missing or not a directory), -n (Do not output the trailing newline), -s (Suppress error messages about nonexistent or unreadable files)",
	},
	"ln": {
		Name:        "ln",
		Description: ": -s (Make symbolic links instead of hard links), -f (Remove existing destination files), -v (Print name of each linked file), -n (Do not dereference symlinks), -b (Make a backup of each existing destination file), -T (Treat LINK_NAME as a normal file always), -L (Dereference TARGETs that are symbolic links), -P (Make hard links directly to symbolic links)",
	},
	"stat": {
		Name:        "stat",
		Description: ": -f (Display information about the filesystem instead of the file), -t (Print the information in terse form), -c (Use the specified format instead of the default), -L (Dereference symbolic links), -h (Print sizes in human readable format), --printf (Print the specified format string)",
	},
	"file": {
		Name:        "file",
		Description: ": -b (Do not prepend filenames to output lines), -i (Output MIME type strings), -z (Try to look inside compressed files), -L (Follow symbolic links), -s (Read block or character special files), -k (Keep going after the first match), -r (Read the magic file from the specified directory)",
	},
	"strings": {
		Name:        "strings",
		Description: ": -a (Scan the whole file, not just the data segment), -n (Print sequences of at least n characters (default is 4)), -t (Print the location of the string in the file), -e (Specify the character encoding), -o (Print the offset of each string in octal), -T (Print the length of each string)",
	},
	"md5sum": {
		Name:        "md5sum",
		Description: ": -c (Read MD5 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode), -w (Warn about improperly formatted checksum lines), --tag (Create a BSD-style checksum), --strict (Exit non-zero for improperly formatted checksum lines)",
	},
	"sha256sum": {
		Name:        "sha256sum",
		Description: ": -c (Read SHA256 sums from the FILEs and check them), -b (Read in binary mode), -t (Read in text mode), -w (Warn about improperly formatted checksum lines), --tag (Create a BSD-style checksum), --strict (Exit non-zero for improperly formatted checksum lines)",
	},
	"gzip": {
		Name:        "gzip",
		Description: ": -d (Decompress), -k (Keep (don't delete) input files), -r (Operate recursively on directories), -c (Write on standard output, keep original files unchanged), -f (Force compression or decompression), -l (List compressed file contents), -n (Do not save the original file name and time stamp), -q (Suppress all warnings), -v (Verbose mode), -1 to -9 (Set the compression level)",
	},
	"gunzip": {
		Name:        "gunzip",
		Description: ": -k (Keep (don't delete) input files), -r (Operate recursively on directories), -c (Write on standard output, keep original files unchanged), -f (Force decompression), -l (List compressed file contents), -n (Do not save the original file name and time stamp), -q (Suppress all warnings), -v (Verbose mode)",
	},
	"bzip2": {
		Name:        "bzip2",
		Description: ": -d (Decompress), -k (Keep (don't delete) input files), -v (Verbose mode), -z (Compress), -c (Write on standard output, keep original files unchanged), -f (Force compression or decompression), -q (Suppress all warnings), -1 to -9 (Set the block size to 100k to 900k)",
	},
	"bunzip2": {
		Name:        "bunzip2",
		Description: ": -k (keep the original file after decompression), -v (verbose mode, show more details during operation), -c (Write on standard output, keep original files unchanged), -f (Force decompression), -q (Suppress all warnings)",
	},
	"zip": {
		Name:        "zip",
		Description: ": -r (recursively include files in subdirectories), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation), -m (move files into the zip archive), -j (junk (don't record) directory names), -0 to -9 (Set the compression level), -x (exclude the following files), -i (include only the following files)",
	},
	"unzip": {
		Name:        "unzip",
		Description: ": -l (list the contents of the archive), -t (test the integrity of the archive), -d (specify the directory to extract files to), -n (never overwrite existing files), -o (overwrite existing files without prompting), -q (quiet mode), -v (verbose mode), -x (exclude the following files), -i (include only the following files)",
	},
	"rar": {
		Name:        "rar",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -v (create volumes with a specified size), -r (recurse into directories), -m (set the compression level), -p (set the password), -k (lock the archive), -s (convert paths to short format)",
	},
	"unrar": {
		Name:        "unrar",
		Description: ": -x (extract files from the archive), -l (list the contents of the archive), -v (verbose mode, show more details during operation), -p (set the password), -y (assume Yes on all queries), -o+ (overwrite files without prompting), -o- (do not overwrite files)",
	},
	"7z": {
		Name:        "7z",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive), -r (recurse into directories), -m (set the compression level), -p (set the password), -y (assume Yes on all queries), -o (specify the output directory)",
	},
	"7za": {
		Name:        "7za",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive), -r (recurse into directories), -m (set the compression level), -p (set the password), -y (assume Yes on all queries), -o (specify the output directory)",
	},
	"7zr": {
		Name:        "7zr",
		Description: ": -a (add files to the archive), -x (extract files from the archive), -t (specify the type of archive), -r (recurse into directories), -m (set the compression level), -p (set the password), -y (assume Yes on all queries), -o (specify the output directory)",
	},
	"free": {
		Name:        "free",
		Description: ": -h (display sizes in human-readable format), -m (display memory in megabytes), -g (display memory in gigabytes), -b (display memory in bytes), -k (display memory in kilobytes), -t (display total memory), -s (display memory usage periodically)",
	},
	"who": {
		Name:        "who",
		Description: ": -a (display all information), -b (show the last system boot time), -q (display the number of logged-in users), -m (show only the current terminal), -r (show current runlevel), -T (display terminal line status)",
	},
	"last": {
		Name:        "last",
		Description: ": -n (specify the number of lines to display), -R (do not display the hostname), -x (display system shutdown and reboot entries), -a (display the hostname in the last column), -d (display the IP address instead of the hostname)",
	},
	"dmesg": {
		Name:        "dmesg",
		Description: ": -C (clear the ring buffer), -c (read and clear the ring buffer), -T (display human-readable timestamps), -f (restrict output to defined facilities), -l (restrict output to defined levels), -n (set the level at which logging of messages is done to the console), -r (print raw message buffer), -w (wait for new messages)",
	},
	"lsblk": {
		Name:        "lsblk",
		Description: ": -f (display file system information), -o (specify output columns), -d (do not display child devices), -a (display all devices), -b (display sizes in bytes), -n (do not print the header), -p (display full device path), -t (display topology information)",
	},
	"blkid": {
		Name:        "blkid",
		Description: ": -p (probe for partitions), -s (show specific information), -o (specify output format), -L (list all labels), -U (list all UUIDs), -c (use the specified cache file), -g (garbage collect the cache file)",
	},
	"fdisk": {
		Name:        "fdisk",
		Description: ": -l (list partition tables), -u (specify units), -c (compatibility mode), -b (specify the sector size), -s (list the size of a partition), -t (list the partition table type)",
	},
	"mkfs": {
		Name:        "mkfs",
		Description: ": -t (specify the file system type), -c (check the device for bad blocks), -v (verbose mode, show more details during operation), -L (set the volume label), -n (do not actually create the file system), -q (quiet mode), -f (force creation even if the file system already exists)",
	},
	"fsck": {
		Name:        "fsck",
		Description: ": -a (automatically repair the file system), -r (prompt interactively before making repairs), -y (assume 'yes' to all prompts), -n (assume 'no' to all prompts), -C (display completion/progress bars), -V (verbose mode), -t (specify the file system type), -A (check all file systems listed in /etc/fstab)",
	},
	"parted": {
		Name:        "parted",
		Description: ": -l (list partition tables), -s (script mode, do not prompt for user input), -a (specify alignment type), -m (machine-readable output), -v (verbose mode), -h (display help message), -V (display version information)",
	},
	"dd": {
		Name:        "dd",
		Description: ": if (specify input file), of (specify output file), bs (specify block size), count (specify number of blocks to copy), skip (skip blocks at start of input), seek (skip blocks at start of output), conv (convert the file as specified), status (specify the level of information to print)",
	},
	"sync": {
		Name:        "sync",
		Description: ": No flags available",
	},
	"shutdown": {
		Name:        "shutdown",
		Description: ": -h (halt the system), -r (reboot the system), -c (cancel a scheduled shutdown), -k (send a warning message, but do not actually shut down), -P (power off the machine), -H (halt the machine), -f (reboot fast, without fsck), -F (force fsck on reboot)",
	},
	"reboot": {
		Name:        "reboot",
		Description: ": No flags available",
	},
	"systemctl": {
		Name:        "systemctl",
		Description: ": -t (specify unit type), -a (show all units), -l (show full output), -q (suppress output), -r (reload the systemd manager configuration), -n (limit the number of journal entries shown), --no-pager (do not pipe output into a pager), --no-ask-password (do not ask for system passwords)",
	},
	"journalctl": {
		Name:        "journalctl",
		Description: ": -f (follow new log entries), -u (show logs for a specific unit), -p (specify priority level), -n (number of lines to show), -r (reverse output), -o (output format), -x (show explanatory messages), --no-pager (do not pipe output into a pager), --since (show logs since a specific date/time), --until (show logs until a specific date/time)",
	},
	"timedatectl": {
		Name:        "timedatectl",
		Description: ": -s (set the system time), -p (print properties), -l (list time zones), -n (do not query the server), -H (operate on a remote host), --adjust-system-clock (adjust the system clock when changing the local RTC time)",
	},
	"hostnamectl": {
		Name:        "hostnamectl",
		Description: ": -s (set the system hostname), -p (print properties), -l (list hostnames), -H (operate on a remote host), --static (set the static hostname), --transient (set the transient hostname), --pretty (set the pretty hostname)",
	},
	"localectl": {
		Name:        "localectl",
		Description: ": -s (set locale settings), -p (print properties), -l (list locales), -H (operate on a remote host), --no-convert (do not convert the specified locale settings)",
	},
	"loginctl": {
		Name:        "loginctl",
		Description: ": -s (set login settings), -p (print properties), -l (list sessions), -H (operate on a remote host), --kill-session (kill a specific session), --kill-user (kill all sessions of a specific user), --no-pager (do not pipe output into a pager)",
	},
	"nmcli": {
		Name:        "nmcli",
		Description: ": -t (terse output), -p (pretty output), -m (specify mode), -f (specify fields to display), -g (specify groups to display), -a (ask for missing parameters), -w (specify timeout for operations), --colors (enable or disable color output)",
	},
	"iwconfig": {
		Name:        "iwconfig",
		Description: ": -a (display all information), -d (enable debug mode), -s (display status), essid (Set the ESSID), mode (Set the mode), freq (Set the frequency), ap (Set the access point), rate (Set the bit rate), txpower (Set the transmit power), retry (Set the retry limit), rts (Set the RTS threshold), frag (Set the fragmentation threshold), key (Set the encryption key), power (Set the power management mode)",
	},
	"iwconfig essid": {
		Name:        "iwconfig essid",
		Description: ": Set the ESSID (network name)",
	},
	"iwconfig mode": {
		Name:        "iwconfig mode",
		Description: ": Set the mode (Managed, Ad-Hoc, Master, Repeater, Secondary, Monitor, Auto)",
	},
	"iwconfig freq": {
		Name:        "iwconfig freq",
		Description: ": Set the frequency or channel",
	},
	"iwconfig ap": {
		Name:        "iwconfig ap",
		Description: ": Set the access point address",
	},
	"iwconfig rate": {
		Name:        "iwconfig rate",
		Description: ": Set the bit rate",
	},
	"iwconfig txpower": {
		Name:        "iwconfig txpower",
		Description: ": Set the transmit power",
	},
	"iwconfig retry": {
		Name:        "iwconfig retry",
		Description: ": Set the retry limit",
	},
	"iwconfig rts": {
		Name:        "iwconfig rts",
		Description: ": Set the RTS (Request to Send) threshold",
	},
	"iwconfig frag": {
		Name:        "iwconfig frag",
		Description: ": Set the fragmentation threshold",
	},
	"iwconfig key": {
		Name:        "iwconfig key",
		Description: ": Set the encryption key",
	},
	"iwconfig power": {
		Name:        "iwconfig power",
		Description: ": Set the power management mode",
	},
	"iwlist": {
		Name:        "iwlist",
		Description: ": -a (display all information), -d (enable debug mode), -s (display status), scan (Scan for available networks), frequency (Show available frequencies), bitrate (Show available bit rates), encryption (Show encryption information), power (Show power management information), txpower (Show transmit power information), retry (Show retry limit information)",
	},
	"iwlist scan": {
		Name:        "iwlist scan",
		Description: ": Scan for available networks",
	},
	"iwlist frequency": {
		Name:        "iwlist frequency",
		Description: ": Show available frequencies",
	},
	"iwlist bitrate": {
		Name:        "iwlist bitrate",
		Description: ": Show available bit rates",
	},
	"iwlist encryption": {
		Name:        "iwlist encryption",
		Description: ": Show encryption information",
	},
	"iwlist power": {
		Name:        "iwlist power",
		Description: ": Show power management information",
	},
	"iwlist txpower": {
		Name:        "iwlist txpower",
		Description: ": Show transmit power information",
	},
	"iwlist retry": {
		Name:        "iwlist retry",
		Description: ": Show retry limit information",
	},
	"ufw": {
		Name:        "ufw",
		Description: ": -a (allow traffic), -d (deny traffic), -r (reject traffic), enable (Enable the firewall), disable (Disable the firewall), status (Show the firewall status), default (Set the default policy), logging (Set the logging level), reset (Reset the firewall to default settings)",
	},
	"fail2ban": {
		Name:        "fail2ban",
		Description: ": -d (debug mode), -q (quiet mode, suppress output), -v (verbose mode, show more details during operation), start (Start the fail2ban service), stop (Stop the fail2ban service), restart (Restart the fail2ban service), status (Show the status of the fail2ban service), reload (Reload the fail2ban configuration)",
	},
	"logrotate": {
		Name:        "logrotate",
		Description: ": -d (debug mode), -f (force mode), -v (verbose mode, show more details during operation), -s (specify state file), -m (specify mail command), -l (specify log file), -p (specify PID file)",
	},
	"cron": {
		Name:        "cron",
		Description: ": No flags available",
	},
	"anacron": {
		Name:        "anacron",
		Description: ": -d (debug mode), -s (safe mode), -u (update timestamps), -t (specify timestamp file), -n (run jobs now), -q (quiet mode)",
	},
	"systemd-analyze": {
		Name:        "systemd-analyze",
		Description: ": -p (specify property), -t (specify type), -f (specify filter), blame (Show time taken to initialize each service), critical-chain (Show the critical chain of services), plot (Generate a graphical plot of boot process), dot (Generate a dependency graph in dot format)",
	},
	"pkill": {
		Name:        "pkill",
		Description: ": -u (specify user), -t (specify terminal), -x (exact match), -f (match against the full command line), -g (specify process group), -P (specify parent process ID), -n (kill the newest process), -o (kill the oldest process), -l (list signal names), -s (specify session ID)",
	},
	"pgrep": {
		Name:        "pgrep",
		Description: ": -u (specify user), -t (specify terminal), -x (exact match), -l (list PID and name), -f (match against full argument lists), -n (select most recently started), -o (select oldest process), -v (negate the matching), -c (count of matching processes)",
	},
	"nice": {
		Name:        "nice",
		Description: ": -n (specify adjustment value), -10 to 19 (specify the priority increment)",
	},
	"renice": {
		Name:        "renice",
		Description: ": -n (specify adjustment value), -p (specify process ID), -u (specify user), -g (specify process group ID)",
	},
	"ionice": {
		Name:        "ionice",
		Description: ": -c (specify class), -n (specify priority), -p (specify process ID), -t (set the idle I/O scheduling class)",
	},
	"watch": {
		Name:        "watch",
		Description: ": -n (specify interval), -d (highlight differences), -t (turn off header), -p (precise mode), -e (exit when command produces output), -g (exit when command exits with a non-zero status), -x (do not exec, but interpret the command as a shell command)",
	},
	"screen": {
		Name:        "screen",
		Description: ": -d (detach session), -r (reattach session), -S (specify session name), -ls (list sessions), -X (send a command to a session), -RR (reattach if possible, otherwise start a new session), -dmS (start a session in detached mode with a name)",
	},
	"tmux": {
		Name:        "tmux",
		Description: ": -d (detach session), -r (reattach session), -S (specify session name), -L (specify socket name), -f (specify configuration file), -2 (force 256-color mode), -u (force UTF-8 mode), -v (verbose mode), -V (print version information)",
	},
	"nohup": {
		Name:        "nohup",
		Description: ": No flags available",
	},
	"disown": {
		Name:        "disown",
		Description: ": -h (do not remove job from table), -a (remove all jobs), -r (remove running jobs)",
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
		Description: ": +short (Short output), +trace (Trace path), @server (DNS server), -b (Source IP), -c (Query class), -f (Batch queries), -k (TSIG key file), -p (Port), -q (Query name), -t (Query type), -x (Reverse lookup), -y (TSIG key), +stats (Query stats), +noall (Clear display flags), +answer (Answer section), +authority (Authority section), +additional (Additional section), +question (Question section)",
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
		Description: ": none",
	},
	"ftp close": {
		Name:        "ftp close",
		Description: ": none",
	},
	"ftp put": {
		Name:        "ftp put",
		Description: ": none",
	},
	"ftp mput": {
		Name:        "ftp mput",
		Description: ": none",
	},
	"ftp get": {
		Name:        "ftp get",
		Description: ": none",
	},
	"ftp mget": {
		Name:        "ftp mget",
		Description: ": none",
	},
	"ftp delete": {
		Name:        "ftp delete",
		Description: ": none",
	},
	"ftp mdelete": {
		Name:        "ftp mdelete",
		Description: ": none",
	},
	"ftp rename": {
		Name:        "ftp rename",
		Description: ": none",
	},
	"ftp mkdir": {
		Name:        "ftp mkdir",
		Description: ": none",
	},
	"ftp rmdir": {
		Name:        "ftp rmdir",
		Description: ": none",
	},
	"ftp ls": {
		Name:        "ftp ls",
		Description: ": none",
	},
	"ftp lcd": {
		Name:        "ftp lcd",
		Description: ": none",
	},
	"ftp cd": {
		Name:        "ftp cd",
		Description: ": none",
	},
	"ftp ascii": {
		Name:        "ftp ascii",
		Description: ": none",
	},
	"ftp binary": {
		Name:        "ftp binary",
		Description: ": none",
	},
	"ftp passive": {
		Name:        "ftp passive",
		Description: ": none",
	},
	"ftp active": {
		Name:        "ftp active",
		Description: ": none",
	},
	"sftp": {
		Name:        "sftp",
		Description: ": -P (Specify the port to connect to on the remote host), -r (Recursively copy entire directories)",
	},
	"sftp get": {
		Name:        "sftp get",
		Description: ": none",
	},
	"sftp put": {
		Name:        "sftp put",
		Description: ": none",
	},
	"sftp ls": {
		Name:        "sftp ls",
		Description: ": none",
	},
	"sftp cd": {
		Name:        "sftp cd",
		Description: ": none",
	},
	"sftp lcd": {
		Name:        "sftp lcd",
		Description: ": none",
	},
	"sftp mkdir": {
		Name:        "sftp mkdir",
		Description: ": none",
	},
	"sftp rmdir": {
		Name:        "sftp rmdir",
		Description: ": none",
	},
	"sftp rm": {
		Name:        "sftp rm",
		Description: ": none",
	},
	"sftp rename": {
		Name:        "sftp rename",
		Description: ": none",
	},
	"sftp chmod": {
		Name:        "sftp chmod",
		Description: ": none",
	},
	"sftp chown": {
		Name:        "sftp chown",
		Description: ": none",
	},
	"sftp chgrp": {
		Name:        "sftp chgrp",
		Description: ": none",
	},
	"sftp lls": {
		Name:        "sftp lls",
		Description: ": none",
	},
	"sftp lmkdir": {
		Name:        "sftp lmkdir",
		Description: ": none",
	},
	"sftp lrm": {
		Name:        "sftp lrm",
		Description: ": none",
	},
	"sftp lrename": {
		Name:        "sftp lrename",
		Description: ": none",
	},
	"sftp lchmod": {
		Name:        "sftp lchmod",
		Description: ": none",
	},
	"sftp lchown": {
		Name:        "sftp lchown",
		Description: ": none",
	},
	"sftp lchgrp": {
		Name:        "sftp lchgrp",
		Description: ": none",
	},
	"telnet": {
		Name:        "telnet",
		Description: ": -8 (Use an 8-bit data path), -E (Stop any escape character from being recognized)",
	},
	"telnet open": {
		Name:        "telnet open",
		Description: ": none",
	},
	"telnet close": {
		Name:        "telnet close",
		Description: ": none",
	},
	"telnet quit": {
		Name:        "telnet quit",
		Description: ": none",
	},
	"telnet send": {
		Name:        "telnet send",
		Description: ": none",
	},
	"telnet status": {
		Name:        "telnet status",
		Description: ": none",
	},
	"telnet set": {
		Name:        "telnet set",
		Description: ": none",
	},
	"telnet unset": {
		Name:        "telnet unset",
		Description: ": none",
	},
	"telnet toggle": {
		Name:        "telnet toggle",
		Description: ": none",
	},
	"telnet mode": {
		Name:        "telnet mode",
		Description: ": none",
	},
	"telnet display": {
		Name:        "telnet display",
		Description: ": none",
	},
	"telnet environ": {
		Name:        "telnet environ",
		Description: ": none",
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
		Description: ": -i (Interface), -w (Write file), -r (Read file), -c (Packet count), -s (Snapshot length), -v, -vv, -vvv (Verbose), -e (Link-level header), -q (Quiet), -X (Hex and ASCII), -XX (Hex, ASCII, link-level), -A (ASCII), -D (List interfaces), -l (Line buffered), -tt, -ttt, -tttt, -ttttt (Timestamps), -C (File size limit), -G (Rotate files), -W (Limit files), -Z (Drop privileges), -K (No checksum verify), -E (Decrypt IPsec), -M (IPsec AH secret)",
	},
	"wireshark": {
		Name:        "wireshark",
		Description: ": -i (Interface), -k (Immediate start), -w (Write to file), -r (Read from file), -c (Packet count), -f (Capture filter), -s (Snapshot length), -v (Verbose), -h (Help), -b (Ring buffer), -t (Timestamp format), -n (No name resolution), -N (Name resolution flags), -S (Real-time update), -T (Text output format), -X (Plugin option), -Y (Display filter), -z (Statistics)",
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
		Description: ": none",
	},
	"docker tag": {
		Name:        "docker tag",
		Description: ": none",
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
		Description: ": none",
	},
	"docker volume": {
		Name:        "docker volume",
		Description: ": none",
	},
	"docker info": {
		Name:        "docker info",
		Description: ": none",
	},
	"docker version": {
		Name:        "docker version",
		Description: ": none",
	},
	"docker swarm": {
		Name:        "docker swarm",
		Description: ": none",
	},
	"docker node": {
		Name:        "docker node",
		Description: ": none",
	},
	"docker service": {
		Name:        "docker service",
		Description: ": none",
	},
	"docker stack": {
		Name:        "docker stack",
		Description: ": none",
	},
	"docker secret": {
		Name:        "docker secret",
		Description: ": none",
	},
	"docker config": {
		Name:        "docker config",
		Description: ": none",
	},
	"docker plugin": {
		Name:        "docker plugin",
		Description: ": none",
	},
	"docker container": {
		Name:        "docker container",
		Description: ": none",
	},
	"docker image": {
		Name:        "docker image",
		Description: ": none",
	},
	"docker system": {
		Name:        "docker system",
		Description: ": none",
	},
	"docker context": {
		Name:        "docker context",
		Description: ": none",
	},
	"docker builder": {
		Name:        "docker builder",
		Description: ": none",
	},
	"docker checkpoint": {
		Name:        "docker checkpoint",
		Description: ": none",
	},
	"docker trust": {
    Name:        "docker trust",
    Description: ": key generate (Generate and load a signing key-pair), key load (Load a private key for signing tags), signer add (Add a signer to a repository), signer remove (Remove a signer from a repository), inspect (Return low-level information about keys and signatures), revoke (Revoke trust for an image), sign (Sign an image), view (View signed tags in a repository)",
	},
	"docker compose": {
		Name:        "docker compose",
		Description: ": up (Create and start containers), down (Stop and remove containers, networks, images, and volumes), build (Build or rebuild services), ps (List containers), exec (Execute a command in a running container), logs (View output from containers), config (Validate and view the Compose file), pull (Pull service images), push (Push service images), restart (Restart services), rm (Remove stopped containers), run (Run a one-off command), start (Start services), stop (Stop services), top (Display the running processes), version (Show the Docker Compose version information)",
	},
	"docker save": {
		Name:        "docker save",
		Description: ": -o (Write to a file instead of STDOUT)",
	},
	"docker load": {
		Name:        "docker load",
		Description: ": -i (Read from a tar archive file), -q (Suppress the load output)",
	},
	"docker history": {
		Name:        "docker history",
		Description: ": -H (Print sizes and dates in human readable format), --no-trunc (Don’t truncate output), -q (Only show numeric IDs)",
	},
	"docker prune": {
		Name:        "docker prune",
		Description: ": -a (Remove all unused images, not just dangling ones), -f (Do not prompt for confirmation), --volumes (Prune volumes)",
	},
	"docker attach": {
		Name:        "docker attach",
		Description: ": --no-stdin (Do not attach STDIN), --sig-proxy (Proxy all received signals to the process)",
	},
	"docker diff": {
		Name:        "docker diff",
		Description: ": No flags available",
	},
	"docker events": {
		Name:        "docker events",
		Description: ": --since (Show all events created since timestamp), --until (Stream events until this timestamp), -f (Filter output based on conditions provided), --format (Format the output using the given Go template)",
	},
	"docker export": {
		Name:        "docker export",
		Description: ": -o (Write to a file instead of STDOUT)",
	},
	"docker import": {
		Name:        "docker import",
		Description: ": -c (Apply Dockerfile instruction to the created image), -m (Set commit message for imported image), --change (Apply Dockerfile instruction to the created image)",
	},
	"docker pause": {
		Name:        "docker pause",
		Description: ": No flags available",
	},
	"docker unpause": {
		Name:        "docker unpause",
		Description: ": No flags available",
	},
	"docker update": {
		Name:        "docker update",
		Description: ": --cpu-shares (Change CPU shares), --memory (Change memory limit), --restart (Change restart policy), --cpus (Change number of CPUs)",
	},
	"docker top": {
		Name:        "docker top",
		Description: ": No flags available",
	},
	"docker stats": {
		Name:        "docker stats",
		Description: ": --all (Show all containers), --no-stream (Disable streaming stats and only pull the first result), --format (Pretty-print images using a Go template)",
	},
	"docker wait": {
		Name:        "docker wait",
		Description: ": No flags available",
	},
	"docker rename": {
		Name:        "docker rename",
		Description: ": No flags available",
	},
	"docker cp": {
		Name:        "docker cp",
		Description: ": -L (Follow symbolic links in the source path)",
	},
	"docker login": {
		Name:        "docker login",
		Description: ": -u (Username), -p (Password), --password-stdin (Take the password from stdin)",
	},
	"docker logout": {
		Name:        "docker logout",
		Description: ": No flags available",
	},
	"docker search": {
		Name:        "docker search",
		Description: ": --filter (Filter output based on conditions provided), --format (Pretty-print search using a Go template), --limit (Max number of search results), --no-trunc (Don’t truncate output)",
	},
	"docker volume create": {
		Name:        "docker volume create",
		Description: ": --driver (Specify volume driver name), --label (Set metadata for a volume), --opt (Set driver specific options)",
	},
	"docker volume inspect": {
		Name:        "docker volume inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker volume ls": {
		Name:        "docker volume ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print volumes using a Go template), -q (Only display volume names)",
	},
	"docker volume prune": {
		Name:        "docker volume prune",
		Description: ": -f (Do not prompt for confirmation)",
	},
	"docker volume rm": {
		Name:        "docker volume rm",
		Description: ": No flags available",
	},
	"docker network create": {
		Name:        "docker network create",
		Description: ": --driver (Driver to manage the Network), --subnet (Subnet in CIDR format), --gateway (IPv4 or IPv6 Gateway for the master subnet), --ip-range (Allocate container IP from a sub-range), --label (Set metadata for a network), --opt (Set driver specific options)",
	},
	"docker network inspect": {
		Name:        "docker network inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker network ls": {
		Name:        "docker network ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print networks using a Go template), -q (Only display network IDs)",
	},
	"docker network prune": {
		Name:        "docker network prune",
		Description: ": -f (Do not prompt for confirmation)",
	},
	"docker network rm": {
		Name:        "docker network rm",
		Description: ": No flags available",
	},
	"docker plugin create": {
		Name:        "docker plugin create",
		Description: ": No flags available",
	},
	"docker plugin disable": {
		Name:        "docker plugin disable",
		Description: ": -f (Force disable)",
	},
	"docker plugin enable": {
		Name:        "docker plugin enable",
		Description: ": No flags available",
	},
	"docker plugin inspect": {
		Name:        "docker plugin inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker plugin install": {
		Name:        "docker plugin install",
		Description: ": --alias (Local name for plugin), --disable (Do not enable the plugin on install), --grant-all-permissions (Grant all permissions necessary to run the plugin)",
	},
	"docker plugin ls": {
		Name:        "docker plugin ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print plugins using a Go template), -q (Only display plugin IDs)",
	},
	"docker plugin push": {
		Name:        "docker plugin push",
		Description: ": No flags available",
	},
	"docker plugin rm": {
		Name:        "docker plugin rm",
		Description: ": -f (Force remove)",
	},
	"docker plugin set": {
		Name:        "docker plugin set",
		Description: ": No flags available",
	},
	"docker plugin upgrade": {
		Name:        "docker plugin upgrade",
		Description: ": --skip-remote-check (Do not check if specified plugin matches existing plugin image)",
	},
	"docker secret create": {
		Name:        "docker secret create",
		Description: ": --label (Set metadata for a secret)",
	},
	"docker secret inspect": {
		Name:        "docker secret inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker secret ls": {
		Name:        "docker secret ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print secrets using a Go template), -q (Only display secret IDs)",
	},
	"docker secret rm": {
		Name:        "docker secret rm",
		Description: ": No flags available",
	},
	"docker config create": {
		Name:        "docker config create",
		Description: ": --label (Set metadata for a config)",
	},
	"docker config inspect": {
		Name:        "docker config inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker config ls": {
		Name:        "docker config ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print configs using a Go template), -q (Only display config IDs)",
	},
	"docker config rm": {
		Name:        "docker config rm",
		Description: ": No flags available",
	},
	"docker swarm init": {
		Name:        "docker swarm init",
		Description: ": --advertise-addr (Advertise address), --autolock (Enable manager autolocking), --cert-expiry (Validity period for node certificates), --dispatcher-heartbeat (Dispatcher heartbeat period), --external-ca (Specifications of one or more certificate signing endpoints), --listen-addr (Listen address), --max-snapshots (Number of additional Raft snapshots to retain), --snapshot-interval (Number of log entries between Raft snapshots)",
	},
	"docker swarm join": {
		Name:        "docker swarm join",
		Description: ": --advertise-addr (Advertise address), --availability (Availability of the node), --data-path-addr (Address or interface to use for data path traffic), --listen-addr (Listen address), --token (Swarm join token)",
	},
	"docker swarm leave": {
		Name:        "docker swarm leave",
		Description: ": -f (Force leave)",
	},
	"docker swarm update": {
		Name:        "docker swarm update",
		Description: ": --autolock (Enable or disable manager autolocking), --cert-expiry (Validity period for node certificates), --dispatcher-heartbeat (Dispatcher heartbeat period), --external-ca (Specifications of one or more certificate signing endpoints), --max-snapshots (Number of additional Raft snapshots to retain), --snapshot-interval (Number of log entries between Raft snapshots)",
	},
	"docker node promote": {
		Name:        "docker node promote",
		Description: ": No flags available",
	},
	"docker node demote": {
		Name:        "docker node demote",
		Description: ": No flags available",
	},
	"docker node inspect": {
		Name:        "docker node inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker node ls": {
		Name:        "docker node ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print nodes using a Go template), -q (Only display node IDs)",
	},
	"docker node rm": {
		Name:        "docker node rm",
		Description: ": No flags available",
	},
	"docker service create": {
		Name:        "docker service create",
		Description: ": --name (Service name), --replicas (Number of tasks), --constraint (Placement constraints), --label (Service labels), --mode (Service mode), --publish (Publish a port), --network (Network attachments), --env (Set environment variables), --mount (Attach a filesystem mount to the service), --secret (Specify secrets to expose to the service), --config (Specify configs to expose to the service)",
	},
	"docker service inspect": {
		Name:        "docker service inspect",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker service ls": {
		Name:        "docker service ls",
		Description: ": -f (Filter output based on conditions provided), --format (Pretty-print services using a Go template), -q (Only display service IDs)",
	},
	"docker service rm": {
		Name:        "docker service rm",
		Description: ": No flags available",
	},
	"docker service scale": {
		Name:        "docker service scale",
		Description: ": No flags available",
	},
	"docker service update": {
		Name:        "docker service update",
		Description: ": --image (Service image), --replicas (Number of tasks), --constraint-add (Add placement constraints), --constraint-rm (Remove placement constraints), --label-add (Add service labels), --label-rm (Remove service labels), --publish-add (Add port to be published), --publish-rm (Remove port to be published), --env-add (Add environment variables), --env-rm (Remove environment variables), --mount-add (Add a filesystem mount to the service), --mount-rm (Remove a filesystem mount from the service), --secret-add (Add secrets to the service), --secret-rm (Remove secrets from the service), --config-add (Add configs to the service), --config-rm (Remove configs from the service)",
	},
	"docker stack deploy": {
		Name:        "docker stack deploy",
		Description: ": --compose-file (Path to a Compose file), --prune (Prune services that are no longer referenced), --with-registry-auth (Send registry authentication details to Swarm agents)",
	},
	"docker stack rm": {
		Name:        "docker stack rm",
		Description: ": No flags available",
	},
	"docker stack services": {
		Name:        "docker stack services",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker stack ps": {
		Name:        "docker stack ps",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker stack ls": {
		Name:        "docker stack ls",
		Description: ": --format (Format the output using the given Go template)",
	},
	"docker checkpoint create": {
		Name:        "docker checkpoint create",
		Description: ": --checkpoint-dir (Use a custom checkpoint storage directory), --leave-running (Leave the container running after checkpointing)",
	},
	"docker checkpoint ls": {
		Name:        "docker checkpoint ls",
		Description: ": --checkpoint-dir (Use a custom checkpoint storage directory)",
	},
	"docker checkpoint rm": {
		Name:        "docker checkpoint rm",
		Description: ": --checkpoint-dir (Use a custom checkpoint storage directory)",
	},
	"docker checkpoint restore": {
		Name:        "docker checkpoint restore",
		Description: ": --checkpoint-dir (Use a custom checkpoint storage directory), --leave-running (Leave the container running after restoring from checkpoint)",
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
		Description: ": none",
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
		Description: ": none",
	},
	"systemctl edit": {
		Name:        "systemctl edit",
		Description: ": --full (Edit the full unit file instead of creating a drop-in snippet)",
	},
	"systemctl set-property": {
		Name:        "systemctl set-property",
		Description: ": none",
	},
	"systemctl help": {
		Name:        "systemctl help",
		Description: ": none",
	},
	"journalctl --boot": {
		Name:        "journalctl --boot",
		Description: ": -k (Show only kernel messages), -p (Show messages with the specified priority)",
	},
	"journalctl --list-boots": {
		Name:        "journalctl --list-boots",
		Description: ": none",
	},
	"journalctl --unit": {
		Name:        "journalctl --unit",
		Description: ": -f (Follow new messages), -n (Show the specified number of most recent messages)",
	},
	"journalctl --since": {
		Name:        "journalctl --since",
		Description: ": none",
	},
	"journalctl --until": {
		Name:        "journalctl --until",
		Description: ": none",
	},
	"journalctl --follow": {
		Name:        "journalctl --follow",
		Description: ": none",
	},
	"journalctl --output": {
		Name:        "journalctl --output",
		Description: ": -o (Specify the output format)",
	},
	"journalctl --priority": {
		Name:        "journalctl --priority",
		Description: ": none",
	},
	"journalctl --grep": {
		Name:        "journalctl --grep",
		Description: ": none",
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
		Description: ": none",
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
		Description: ": none",
	},
	"cron start": {
		Name:        "cron start",
		Description: ": none",
	},
	"cron stop": {
		Name:        "cron stop",
		Description: ": none",
	},
	"cron restart": {
		Name:        "cron restart",
		Description: ": none",
	},
	"cron reload": {
		Name:        "cron reload",
		Description: ": none",
	},
	"cron status": {
		Name:        "cron status",
		Description: ": none",
	},
	"cron list": {
		Name:        "cron list",
		Description: ": none",
	},
	"cron add": {
		Name:        "cron add",
		Description: ": none",
	},
	"cron remove": {
		Name:        "cron remove",
		Description: ": none",
	},
	"cron edit": {
		Name:        "cron edit",
		Description: ": none",
	},
	"ufw enable": {
		Name:        "ufw enable",
		Description: ": none",
	},
	"ufw disable": {
		Name:        "ufw disable",
		Description: ": none",
	},
	"ufw status": {
		Name:        "ufw status",
		Description: ": none",
	},
	"ufw allow": {
		Name:        "ufw allow",
		Description: ": none",
	},
	"ufw deny": {
		Name:        "ufw deny",
		Description: ": none",
	},
	"ufw reject": {
		Name:        "ufw reject",
		Description: ": none",
	},
	"ufw limit": {
		Name:        "ufw limit",
		Description: ": none",
	},
	"ufw delete": {
		Name:        "ufw delete",
		Description: ": none",
	},
	"ufw reset": {
		Name:        "ufw reset",
		Description: ": none",
	},
	"ufw reload": {
		Name:        "ufw reload",
		Description: ": none",
	},
	"ufw logging": {
		Name:        "ufw logging",
		Description: ": none",
	},
	"nmcli general": {
		Name:        "nmcli general",
		Description: ": none",
	},
	"nmcli networking": {
		Name:        "nmcli networking",
		Description: ": none",
	},
	"nmcli radio": {
		Name:        "nmcli radio",
		Description: ": none",
	},
	"nmcli connection": {
		Name:        "nmcli connection",
		Description: ": none",
	},
	"nmcli device": {
		Name:        "nmcli device",
		Description: ": none",
	},
	"nmcli agent": {
		Name:        "nmcli agent",
		Description: ": none",
	},
	"nmcli monitor": {
		Name:        "nmcli monitor",
		Description: ": none",
	},
	"loginctl list-sessions": {
		Name:        "loginctl list-sessions",
		Description: ": none",
	},
	"loginctl session-status": {
		Name:        "loginctl session-status",
		Description: ": none",
	},
	"loginctl terminate-session": {
		Name:        "loginctl terminate-session",
		Description: ": none",
	},
	"loginctl list-users": {
		Name:        "loginctl list-users",
		Description: ": none",
	},
	"loginctl user-status": {
		Name:        "loginctl user-status",
		Description: ": none",
	},
	"loginctl enable-linger": {
		Name:        "loginctl enable-linger",
		Description: ": none",
	},
	"loginctl disable-linger": {
		Name:        "loginctl disable-linger",
		Description: ": none",
	},
	"loginctl lock-session": {
		Name:        "loginctl lock-session",
		Description: ": none",
	},
	"loginctl unlock-session": {
		Name:        "loginctl unlock-session",
		Description: ": none",
	},
	"loginctl lock-sessions": {
		Name:        "loginctl lock-sessions",
		Description: ": none",
	},
	"loginctl unlock-sessions": {
		Name:        "loginctl unlock-sessions",
		Description: ": none",
	},
	"localectl status": {
		Name:        "localectl status",
		Description: ": none",
	},
	"localectl set-locale": {
		Name:        "localectl set-locale",
		Description: ": none",
	},
	"localectl list-locales": {
		Name:        "localectl list-locales",
		Description: ": none",
	},
	"localectl set-keymap": {
		Name:        "localectl set-keymap",
		Description: ": none",
	},
	"localectl list-keymaps": {
		Name:        "localectl list-keymaps",
		Description: ": none",
	},
	"hostnamectl status": {
		Name:        "hostnamectl status",
		Description: ": none",
	},
	"hostnamectl set-hostname": {
		Name:        "hostnamectl set-hostname",
		Description: ": none",
	},
	"hostnamectl set-icon-name": {
		Name:        "hostnamectl set-icon-name",
		Description: ": none",
	},
	"hostnamectl set-chassis": {
		Name:        "hostnamectl set-chassis",
		Description: ": none",
	},
	"hostnamectl set-deployment": {
		Name:        "hostnamectl set-deployment",
		Description: ": none",
	},
	"hostnamectl set-location": {
		Name:        "hostnamectl set-location",
		Description: ": none",
	},
	"timedatectl status": {
		Name:        "timedatectl status",
		Description: ": none",
	},
	"timedatectl set-time": {
		Name:        "timedatectl set-time",
		Description: ": none",
	},
	"timedatectl set-timezone": {
		Name:        "timedatectl set-timezone",
		Description: ": none",
	},
	"timedatectl list-timezones": {
		Name:        "timedatectl list-timezones",
		Description: ": none",
	},
	"timedatectl set-ntp": {
		Name:        "timedatectl set-ntp",
		Description: ": none",
	},
	"lazygit": {
    	Name:        "lazygit",
    	Description:       ": -v (Show version), --version (Show version), -h (Show help), --help (Show help)",
	},
	"fzf": {
		Name:        "fzf",
		Description: "A command-line fuzzy finder",
		Description:       ": --version (Show version), --help (Show help), --height (Set height), --reverse (Reverse layout), --border (Show border), --prompt (Set prompt), --multi (Enable multi-select), --no-mouse (Disable mouse), --inline-info (Show inline info), --preview (Set preview command), --preview-window (Set preview window), --preview-wrap (Wrap preview), --preview-border (Set preview border), --preview-panes (Set preview panes), --preview-lines (Set preview lines), --preview-columns (Set preview columns), --preview-on (Set preview on), --bind (Set key bindings), --expect (Set expected keys), --header (Set header), --header-lines (Set header lines), --ansi (Enable ANSI color), --color (Set color), --no-bold (Disable bold), --tabstop (Set tab stop), --margin (Set margin), --min-height (Set minimum height), --min-width (Set minimum width), --prompt (Set prompt), --pointer (Set pointer), --marker (Set marker), --print-query (Print query), --query (Set query), --select-1 (Select first match), --exit-0 (Exit if no match), --tiebreak (Set tiebreak), --cycle (Enable cycling), --keep-right (Keep right), --keep-bottom (Keep bottom), --sync (Enable sync), --no-clear (Disable clear), --print0 (Print null), --delimiter (Set delimiter), --nth (Set nth), --with-nth (Set with nth), --layout (Set layout), --min (Set minimum), --max (Set maximum), --no-sort (Disable sort), --filter (Set filter)",
	},
	"rg": {
		Name:        "rg",
		Description: "ripgrep, a line-oriented search tool that recursively searches your current directory for a regex pattern",
		Description:       ": -i (Ignore case), --ignore-case (Ignore case), -v (Invert match), --invert-match (Invert match), -w (Word regexp), --word-regexp (Word regexp), -r (Replace), --replace (Replace)",
	},
	"postman": {
		Name:        "postman",
		Description: "A collaboration platform for API development",
		Description:       ": --help (Show help), --version (Show version)",
	},
	"http": {
		Name:        "http",
		Description: "HTTPie, a command line HTTP client",
		Description:       ": --json (Send JSON), --form (Send form), --pretty (Pretty print), --style (Set style), --verbose (Verbose output)",
	},
	"jq": {
		Name:        "jq",
		Description: "A lightweight and flexible command-line JSON processor",
		Description:       ": -c (Compact output), --compact-output (Compact output), -r (Raw output), --raw-output (Raw output), -s (Slurp), --slurp (Slurp), -R (Raw input), --raw-input (Raw input)",
	},
	"scp": {
		Name:        "scp",
		Description: "Secure copy (remote file copy program)",
		Description:       ": -r (Recursive), -p (Preserve attributes), -q (Quiet mode), -C (Compression), -i (Identity file)",
	},
	"hg": {
		Name:        "hg",
		Description: "Mercurial version control system",
		Description:       ": -v (Verbose), --verbose (Verbose), -q (Quiet), --quiet (Quiet), --debug (Debug), --traceback (Traceback)",
	},
	"kubernetes": {
		Name:        "kubernetes",
		Description:       ": --kubeconfig (Set kubeconfig), --context (Set context), --namespace (Set namespace), --user (Set user)",
	},
	"kubectx": {
		Name:        "kubectx",
		Description:       ": -h (Show help), --help (Show help)",
	},
	"kubens": {
		Name:        "kubens",
		Description:       ": -h (Show help), --help (Show help)",
	},
	"svn": {
		Name:        "svn",
		Description:       ": --version (Show version), --help (Show help), --quiet (Quiet mode), --verbose (Verbose mode)",
	},
	"svn checkout": {
    	Name:        "svn checkout",
   		Description: ": --revision (Specify revision), --quiet (Suppress output), --depth (Limit operation by depth), --force (Force operation), --ignore-externals (Ignore externals)",
	},
	"svn commit": {
		Name:        "svn commit",
		Description: ": --message (Log message), --file (Log message file), --quiet (Suppress output), --depth (Limit operation by depth), --targets (Targets file)",
	},
	"svn update": {
		Name:        "svn update",
		Description: ": --revision (Specify revision), --quiet (Suppress output), --depth (Limit operation by depth), --set-depth (Set new working copy depth), --force (Force operation)",
	},
	"svn add": {
		Name:        "svn add",
		Description: ": --quiet (Suppress output), --depth (Limit operation by depth), --force (Force operation), --no-ignore (Disregard default and svn:ignore property ignores)",
	},
	"svn delete": {
		Name:        "svn delete",
		Description: ": --message (Log message), --file (Log message file), --quiet (Suppress output), --force (Force operation), --keep-local (Keep local copy)",
	},
	"svn status": {
		Name:        "svn status",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --show-updates (Display update information), --depth (Limit operation by depth), --no-ignore (Disregard default and svn:ignore property ignores)",
	},
	"svn log": {
		Name:        "svn log",
		Description: ": --revision (Specify revision), --quiet (Suppress output), --verbose (Print extra information), --stop-on-copy (Stop on copy), --incremental (Output incrementally)",
	},
	"svn diff": {
		Name:        "svn diff",
		Description: ": --revision (Specify revision), --quiet (Suppress output), --depth (Limit operation by depth), --diff-cmd (Use external diff command), --extensions (Pass arguments to diff command)",
	},
	"svn merge": {
		Name:        "svn merge",
		Description: ": --revision (Specify revision), --quiet (Suppress output), --depth (Limit operation by depth), --force (Force operation), --record-only (Record merge information only)",
	},
	"svn revert": {
		Name:        "svn revert",
		Description: ": --quiet (Suppress output), --depth (Limit operation by depth), --targets (Targets file), --changelist (Operate only on members of changelist)",
	},
	"hg init": {
		Name:        "hg init",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --debug (Print debug information), --traceback (Print traceback on error), --config (Set/override config option)",
	},
	"hg clone": {
		Name:        "hg clone",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --debug (Print debug information), --traceback (Print traceback on error), --config (Set/override config option)",
	},
	"hg add": {
		Name:        "hg add",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --debug (Print debug information), --traceback (Print traceback on error), --config (Set/override config option)",
	},
	"hg commit": {
		Name:        "hg commit",
		Description: ": --message (Log message), --file (Log message file), --addremove (Mark new/missing files), --close-branch (Mark branch as closed), --date (Record datecode as commit date)",
	},
	"hg update": {
		Name:        "hg update",
		Description: ": --clean (Discard uncommitted changes), --check (Check for uncommitted changes), --date (Update to date), --rev (Update to revision), --quiet (Suppress output)",
	},
	"hg merge": {
		Name:        "hg merge",
		Description: ": --tool (Specify merge tool), --preview (Preview merge), --quiet (Suppress output), --verbose (Print extra information), --debug (Print debug information)",
	},
	"hg push": {
		Name:        "hg push",
		Description: ": --force (Force push), --new-branch (Allow push of new branch), --rev (Push only specified revision), --bookmark (Push only specified bookmark), --branch (Push only specified branch)",
	},
	"hg pull": {
		Name:        "hg pull",
		Description: ": --update (Update to new branch head), --force (Force pull), --rev (Pull only specified revision), --bookmark (Pull only specified bookmark), --branch (Pull only specified branch)",
	},
	"hg log": {
		Name:        "hg log",
		Description: ": --rev (Show specified revision), --branch (Show specified branch), --limit (Limit number of changesets), --date (Show changesets matching date), --keyword (Do case-insensitive search for keyword)",
	},
	"hg status": {
		Name:        "hg status",
		Description: ": --all (Show status of all files), --modified (Show only modified files), --added (Show only added files), --removed (Show only removed files), --deleted (Show only deleted files)",
	},
	"hg diff": {
		Name:        "hg diff",
		Description: ": --rev (Show changes relative to revision), --change (Show changes in revision), --stat (Show diffstat), --git (Use git extended diff format), --unified (Number of lines of context)",
	},
	"hg revert": {
		Name:        "hg revert",
		Description: ": --rev (Revert to revision), --date (Revert to date), --all (Revert all changes), --no-backup (Do not save backup), --interactive (Prompt for each change)",
	},
	"bzr init": {
		Name:        "bzr init",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --no-tree (Do not create working tree), --format (Specify repository format), --standalone (Create standalone branch)",
	},
	"bzr branch": {
		Name:        "bzr branch",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --stacked (Create stacked branch), --use-existing-dir (Use existing directory), --revision (Branch at revision)",
	},
	"bzr add": {
		Name:        "bzr add",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --dry-run (Show what would be done), --no-recurse (Do not recurse into subdirectories), --file (Add specific file)",
	},
	"bzr commit": {
		Name:        "bzr commit",
		Description: ": --message (Log message), --file (Log message file), --quiet (Suppress output), --verbose (Print extra information), --unchanged (Commit even if no changes)",
	},
	"bzr update": {
		Name:        "bzr update",
		Description: ": --revision (Update to revision), --quiet (Suppress output), --verbose (Print extra information), --force (Force update), --dry-run (Show what would be done)",
	},
	"bzr merge": {
		Name:        "bzr merge",
		Description: ": --revision (Merge revision), --quiet (Suppress output), --verbose (Print extra information), --force (Force merge), --dry-run (Show what would be done)",
	},
	"bzr push": {
		Name:        "bzr push",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --stacked (Create stacked branch), --use-existing-dir (Use existing directory), --revision (Push at revision)",
	},
	"bzr pull": {
		Name:        "bzr pull",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --overwrite (Overwrite local changes), --revision (Pull at revision), --dry-run (Show what would be done)",
	},
	"bzr log": {
		Name:        "bzr log",
		Description: ": --revision (Show revision), --quiet (Suppress output), --verbose (Print extra information), --short (Show short log), --line (Show log in one line)",
	},
	"bzr status": {
		Name:        "bzr status",
		Description: ": --quiet (Suppress output), --verbose (Print extra information), --short (Show short status), --revision (Show status at revision), --no-pending (Do not show pending merges)",
	},
	"bzr diff": {
		Name:        "bzr diff",
		Description: ": --revision (Show diff at revision), --quiet (Suppress output), --verbose (Print extra information), --stat (Show diffstat), --diff-options (Pass options to diff command)",
	},
	"bzr revert": {
		Name:        "bzr revert",
		Description: ": --revision (Revert to revision), --quiet (Suppress output), --verbose (Print extra information), --no-backup (Do not save backup), --dry-run (Show what would be done)",
	},
	"podman run": {
		Name:        "podman run",
		Description: ": --detach (Run container in background), --name (Assign name to container), --rm (Remove container after exit), --interactive (Keep STDIN open), --tty (Allocate pseudo-TTY)",
	},
	"podman build": {
		Name:        "podman build",
		Description: ": --tag (Name and optionally a tag in the 'name:tag' format), --file (Name of the Dockerfile), --quiet (Suppress output), --no-cache (Do not use cache), --rm (Remove intermediate containers)",
	},
	"podman pull": {
		Name:        "podman pull",
		Description: ": --quiet (Suppress output), --all-tags (Download all tagged images), --disable-content-trust (Skip image verification), --platform (Set platform), --tls-verify (Verify TLS)",
	},
	"podman push": {
		Name:        "podman push",
		Description: ": --quiet (Suppress output), --all-tags (Push all tagged images), --disable-content-trust (Skip image verification), --platform (Set platform), --tls-verify (Verify TLS)",
	},
	"podman images": {
		Name:        "podman images",
		Description: ": --quiet (Only show numeric IDs), --all (Show all images), --digests (Show digests), --no-trunc (Do not truncate output), --format (Pretty-print images using a Go template)",
	},
	"podman ps": {
		Name:        "podman ps",
		Description: ": --all (Show all containers), --quiet (Only display numeric IDs), --no-trunc (Do not truncate output), --size (Display total file sizes), --filter (Filter output based on conditions)",
	},
	"podman stop": {
		Name:        "podman stop",
		Description: ": --time (Seconds to wait for stop before killing it), --all (Stop all running containers), --ignore (Ignore errors), --latest (Stop the latest container), --cidfile (Read container ID from file)",
	},
	"podman start": {
		Name:        "podman start",
		Description: ": --attach (Attach container's STDOUT and STDERR), --interactive (Keep STDIN open), --latest (Start the latest container), --all (Start all containers), --detach-keys (Override the key sequence for detaching a container)",
	},
	"podman rm": {
		Name:        "podman rm",
		Description: ": --force (Force removal of running container), --volumes (Remove anonymous volumes), --all (Remove all containers), --latest (Remove the latest container), --cidfile (Read container ID from file)",
	},
	"podman rmi": {
		Name:        "podman rmi",
		Description: ": --force (Force removal of image), --all (Remove all images), --prune (Remove dangling images), --no-prune (Do not remove untagged parents), --filter (Filter output based on conditions)",
	},
	"rkt": {
    	Name:        "rkt",
    	Description: ": --insecure-options (Set insecure options), --dir (Set directory), --local-config (Set local config), --system-config (Set system config), --user-config (Set user config)",
	},
	"rkt run": {
		Name:        "rkt run",
		Description: ": --insecure-options (Set insecure options), --dir (Set directory), --local-config (Set local config), --system-config (Set system config), --user-config (Set user config)",
	},
	"rkt fetch": {
		Name:        "rkt fetch",
		Description: ": --insecure-options (Set insecure options), --dir (Set directory), --local-config (Set local config), --system-config (Set system config), --user-config (Set user config)",
	},
	"rkt list": {
		Name:        "rkt list",
		Description: ": --full (Show full output), --no-legend (Do not print legend), --format (Format output), --sort (Sort output), --filter (Filter output)",
	},
	"rkt rm": {
		Name:        "rkt rm",
		Description: ": --force (Force removal), --uuid-file (Read UUID from file), --wait (Wait for removal), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"rkt gc": {
		Name:        "rkt gc",
		Description: ": --grace-period (Set grace period), --mark-only (Mark for garbage collection only), --no-legend (Do not print legend), --quiet (Suppress output), --uuid-file (Read UUID from file)",
	},
	"rkt stop": {
		Name:        "rkt stop",
		Description: ": --force (Force stop), --uuid-file (Read UUID from file), --wait (Wait for stop), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"minikube": {
		Name:        "minikube",
		Description: ": --vm-driver (Set VM driver), --cpus (Set CPUs), --memory (Set memory), --disk-size (Set disk size), --kubernetes-version (Set Kubernetes version)",
	},
	"minikube start": {
		Name:        "minikube start",
		Description: ": --vm-driver (Set VM driver), --cpus (Set CPUs), --memory (Set memory), --disk-size (Set disk size), --kubernetes-version (Set Kubernetes version)",
	},
	"minikube stop": {
		Name:        "minikube stop",
		Description: ": --all (Stop all running clusters), --force (Force stop), --wait (Wait for stop), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"minikube delete": {
		Name:        "minikube delete",
		Description: ": --all (Delete all clusters), --force (Force delete), --wait (Wait for delete), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"minikube status": {
		Name:        "minikube status",
		Description: ": --format (Format output), --wait (Wait for status), --quiet (Suppress output), --no-legend (Do not print legend), --cluster (Specify cluster)",
	},
	"minikube dashboard": {
		Name:        "minikube dashboard",
		Description: ": --url (Display dashboard URL), --wait (Wait for dashboard), --quiet (Suppress output), --no-legend (Do not print legend), --cluster (Specify cluster)",
	},
	"minikube addons": {
		Name:        "minikube addons",
		Description: ": --enable (Enable addon), --disable (Disable addon), --list (List addons), --wait (Wait for addon), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"minikube service": {
		Name:        "minikube service",
		Description: ": --url (Display service URL), --wait (Wait for service), --quiet (Suppress output), --no-legend (Do not print legend), --namespace (Specify namespace)",
	},
	"helm": {
		Name:        "helm",
		Description: ": --debug (Enable debug), --home (Set home), --host (Set host), --kube-context (Set kube context), --tiller-namespace (Set tiller namespace)",
	},
	"helm repo": {
		Name:        "helm repo",
		Description: ": --add (Add repository), --remove (Remove repository), --update (Update repository), --list (List repositories), --wait (Wait for repository), --quiet (Suppress output)",
	},
	"helm search": {
		Name:        "helm search",
		Description: ": --repo (Specify repository), --version (Specify version), --regexp (Use regular expression), --devel (Include development versions), --wait (Wait for search), --quiet (Suppress output)",
	},
	"helm install": {
		Name:        "helm install",
		Description: ": --name (Specify release name), --namespace (Specify namespace), --version (Specify version), --values (Specify values file), --wait (Wait for install), --quiet (Suppress output)",
	},
	"helm upgrade": {
		Name:        "helm upgrade",
		Description: ": --name (Specify release name), --namespace (Specify namespace), --version (Specify version), --values (Specify values file), --wait (Wait for upgrade), --quiet (Suppress output)",
	},
	"helm uninstall": {
		Name:        "helm uninstall",
		Description: ": --name (Specify release name), --namespace (Specify namespace), --wait (Wait for uninstall), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"helm list": {
		Name:        "helm list",
		Description: ": --all (List all releases), --namespace (Specify namespace), --deployed (List deployed releases), --failed (List failed releases), --pending (List pending releases), --quiet (Suppress output)",
	},
	"helm status": {
		Name:        "helm status",
		Description: ": --name (Specify release name), --namespace (Specify namespace), --wait (Wait for status), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"helm rollback": {
		Name:        "helm rollback",
		Description: ": --name (Specify release name), --namespace (Specify namespace), --version (Specify version), --wait (Wait for rollback), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"kubeadm": {
		Name:        "kubeadm",
		Description: ": --config (Set config), --kubernetes-version (Set Kubernetes version), --pod-network-cidr (Set pod network CIDR), --service-cidr (Set service CIDR), --token (Set token)",
	},
	"kubeadm init": {
		Name:        "kubeadm init",
		Description: ": --config (Set config), --kubernetes-version (Set Kubernetes version), --pod-network-cidr (Set pod network CIDR), --service-cidr (Set service CIDR), --token (Set token)",
	},
	"kubeadm join": {
		Name:        "kubeadm join",
		Description: ": --token (Specify token), --discovery-token-ca-cert-hash (Specify discovery token CA cert hash), --control-plane (Join as control plane), --node-name (Specify node name), --quiet (Suppress output)",
	},
	"kubeadm reset": {
		Name:        "kubeadm reset",
		Description: ": --force (Force reset), --cert-dir (Specify certificate directory), --cri-socket (Specify CRI socket), --ignore-preflight-errors (Ignore preflight errors), --quiet (Suppress output)",
	},
	"kubeadm upgrade": {
		Name:        "kubeadm upgrade",
		Description: ": --config (Set config), --kubernetes-version (Set Kubernetes version), --etcd-upgrade (Upgrade etcd), --renew-certificates (Renew certificates), --quiet (Suppress output)",
	},
	"kubeadm token": {
		Name:        "kubeadm token",
		Description: ": --print-join-command (Print join command), --ttl (Set token TTL), --usages (Specify token usages), --groups (Specify token groups), --quiet (Suppress output)",
	},
	"kubeadm config": {
		Name:        "kubeadm config",
		Description: ": --print-default (Print default config), --kubeconfig (Specify kubeconfig), --component-config (Manage component config), --upload (Upload config), --quiet (Suppress output)",
	},
	"kops": {
		Name:        "kops",
		Description: ": --state (Set state), --name (Set name), --zones (Set zones), --node-count (Set node count), --node-size (Set node size)",
	},
	"kops create": {
		Name:        "kops create",
		Description: ": --state (Set state), --name (Set name), --zones (Set zones), --node-count (Set node count), --node-size (Set node size)",
	},
	"kops update": {
		Name:        "kops update",
		Description: ": --state (Set state), --name (Set name), --yes (Apply changes), --dry-run (Show what would be done), --quiet (Suppress output)",
	},
	"kops delete": {
		Name:        "kops delete",
		Description: ": --state (Set state), --name (Set name), --yes (Apply changes), --dry-run (Show what would be done), --quiet (Suppress output)",
	},
	"kops get": {
		Name:        "kops get",
		Description: ": --state (Set state), --name (Set name), --output (Specify output format), --fields (Specify fields), --quiet (Suppress output)",
	},
	"kops validate": {
		Name:        "kops validate",
		Description: ": --state (Set state), --name (Set name), --wait (Wait for validation), --quiet (Suppress output), --no-legend (Do not print legend)",
	},
	"kops rolling-update": {
		Name:        "kops rolling-update",
		Description: ": --state (Set state), --name (Set name), --yes (Apply changes), --dry-run (Show what would be done), --quiet (Suppress output)",
	},
	"kubectl": {
		Name:        "kubectl",
		Description: ": --kubeconfig (Set kubeconfig), --context (Set context), --namespace (Set namespace), --user (Set user)",
	},
	"kubectl get": {
		Name:        "kubectl get",
		Description: ": --output (Specify output format), --watch (Watch for changes), --all-namespaces (Show all namespaces), --selector (Specify label selector), --field-selector (Specify field selector)",
	},
	"kubectl describe": {
		Name:        "kubectl describe",
		Description: ": --namespace (Specify namespace), --selector (Specify label selector), --field-selector (Specify field selector), --show-events (Show events), --quiet (Suppress output)",
	},
	"kubectl create": {
		Name:        "kubectl create",
		Description: ": --filename (Specify filename), --namespace (Specify namespace), --output (Specify output format), --dry-run (Show what would be done), --quiet (Suppress output)",
	},
	"kubectl delete": {
		Name:        "kubectl delete",
		Description: ": --filename (Specify filename), --namespace (Specify namespace), --output (Specify output format), --force (Force delete), --grace-period (Set grace period)",
	},
	"kubectl apply": {
		Name:        "kubectl apply",
		Description: ": --filename (Specify filename), --namespace (Specify namespace), --output (Specify output format), --dry-run (Show what would be done), --force (Force apply)",
	},
	"kubectl logs": {
		Name:        "kubectl logs",
		Description: ": --follow (Follow logs), --previous (Show previous logs), --timestamps (Show timestamps), --tail (Show last N lines), --since (Show logs since duration)",
	},
	"kubectl exec": {
		Name:        "kubectl exec",
		Description: ": --stdin (Keep STDIN open), --tty (Allocate pseudo-TTY), --container (Specify container), --namespace (Specify namespace), --quiet (Suppress output)",
	},
	"kubectl scale": {
		Name:        "kubectl scale",
		Description: ": --replicas (Set number of replicas), --filename (Specify filename), --namespace (Specify namespace), --output (Specify output format), --quiet (Suppress output)",
	},
	"kubectl expose": {
		Name:        "kubectl expose",
		Description: ": --port (Set port), --target-port (Set target port), --type (Set service type), --name (Set service name), --namespace (Specify namespace)",
	},
	"vagrant": {
    	Name:        "vagrant",
    	Description: ": --version (Show version), --help (Show help), --provider (Specify provider), --debug (Enable debug), --no-color (Disable color output)",
	},
	"vagrant init": {
		Name:        "vagrant init",
		Description: ": --minimal (Create minimal Vagrantfile), --output (Specify output file), --force (Overwrite existing Vagrantfile), --box (Specify box), --box-version (Specify box version)",
	},
	"vagrant up": {
		Name:        "vagrant up",
		Description: ": --provision (Run provisioners), --no-provision (Do not run provisioners), --provider (Specify provider), --parallel (Enable parallel execution), --destroy-on-error (Destroy on error)",
	},
	"vagrant halt": {
		Name:        "vagrant halt",
		Description: ": --force (Force halt), --provider (Specify provider), --parallel (Enable parallel execution), --no-color (Disable color output), --debug (Enable debug)",
	},
	"vagrant destroy": {
		Name:        "vagrant destroy",
		Description: ": --force (Force destroy), --parallel (Enable parallel execution), --no-color (Disable color output), --debug (Enable debug), --graceful (Graceful shutdown)",
	},
	"vagrant ssh": {
		Name:        "vagrant ssh",
		Description: ": --command (Run command), --plain (Plain mode), --extra-args (Extra SSH arguments), --no-tty (Disable TTY), --no-color (Disable color output)",
	},
	"vagrant status": {
		Name:        "vagrant status",
		Description: ": --machine-readable (Machine-readable output), --debug (Enable debug), --no-color (Disable color output), --provider (Specify provider), --parallel (Enable parallel execution)",
	},
	"vagrant provision": {
		Name:        "vagrant provision",
		Description: ": --provision-with (Specify provisioners), --debug (Enable debug), --no-color (Disable color output), --parallel (Enable parallel execution), --force (Force provision)",
	},
	"vagrant reload": {
		Name:        "vagrant reload",
		Description: ": --provision (Run provisioners), --no-provision (Do not run provisioners), --provider (Specify provider), --parallel (Enable parallel execution), --force (Force reload)",
	},
	"packer": {
		Name:        "packer",
		Description: ": --version (Show version), --help (Show help), --debug (Enable debug), --color (Enable color output), --log (Enable logging)",
	},
	"packer build": {
		Name:        "packer build",
		Description: ": --only (Build only specified builders), --except (Exclude specified builders), --parallel-builds (Number of parallel builds), --timestamp-ui (Show timestamps), --var (Set variable)",
	},
	"packer validate": {
		Name:        "packer validate",
		Description: ": --syntax-only (Check syntax only), --var (Set variable), --var-file (Specify variable file), --except (Exclude specified builders), --only (Validate only specified builders)",
	},
	"packer inspect": {
		Name:        "packer inspect",
		Description: ": --except (Exclude specified builders), --only (Inspect only specified builders), --var (Set variable), --var-file (Specify variable file), --debug (Enable debug)",
	},
	"packer fix": {
		Name:        "packer fix",
		Description: ": --var (Set variable), --var-file (Specify variable file), --except (Exclude specified builders), --only (Fix only specified builders), --debug (Enable debug)",
	},
	"packer version": {
		Name:        "packer version",
		Description: ": --help (Show help), --debug (Enable debug), --color (Enable color output), --log (Enable logging), --timestamp-ui (Show timestamps)",
	},
	"packer fmt": {
		Name:        "packer fmt",
		Description: ": --check (Check format), --diff (Show differences), --list (List files), --write (Write result), --recursive (Recurse into directories)",
	},
	"terraform": {
		Name:        "terraform",
		Description: ": --version (Show version), --help (Show help), --debug (Enable debug), --no-color (Disable color output), --log-level (Set log level)",
	},
	"terraform init": {
		Name:        "terraform init",
		Description: ": --backend (Configure backend), --force-copy (Force copy), --get (Download modules), --input (Ask for input), --lock (Lock state)",
	},
	"terraform plan": {
		Name:        "terraform plan",
		Description: ": --destroy (Create destroy plan), --detailed-exitcode (Detailed exit code), --input (Ask for input), --lock (Lock state), --out (Save plan)",
	},
	"terraform apply": {
		Name:        "terraform apply",
		Description: ": --auto-approve (Skip approval), --input (Ask for input), --lock (Lock state), --parallelism (Set parallelism), --refresh-only (Refresh state only)",
	},
	"terraform destroy": {
		Name:        "terraform destroy",
		Description: ": --auto-approve (Skip approval), --input (Ask for input), --lock (Lock state), --parallelism (Set parallelism), --refresh-only (Refresh state only)",
	},
	"terraform validate": {
		Name:        "terraform validate",
		Description: ": --json (Output in JSON), --no-color (Disable color output), --check-variables (Check variables), --backend (Configure backend), --lock (Lock state)",
	},
	"terraform fmt": {
		Name:        "terraform fmt",
		Description: ": --check (Check format), --diff (Show differences), --list (List files), --write (Write result), --recursive (Recurse into directories)",
	},
	"terraform show": {
		Name:        "terraform show",
		Description: ": --json (Output in JSON), --no-color (Disable color output), --module-depth (Set module depth), --lock (Lock state), --input (Ask for input)",
	},
	"terraform state": {
		Name:        "terraform state",
		Description: ": --backup (Create backup), --lock (Lock state), --input (Ask for input), --no-color (Disable color output), --parallelism (Set parallelism)",
	},
	"terraform output": {
		Name:        "terraform output",
		Description: ": --json (Output in JSON), --no-color (Disable color output), --state (Specify state file), --module (Specify module), --lock (Lock state)",
	},
	"ansible": {
		Name:        "ansible",
		Description: ": --version (Show version), --help (Show help), --inventory (Specify inventory), --limit (Limit hosts), --module-name (Specify module)",
	},
	"ansible-playbook": {
		Name:        "ansible-playbook",
		Description: ": --inventory (Specify inventory), --limit (Limit hosts), --tags (Specify tags), --skip-tags (Skip tags), --check (Check mode)",
	},
	"ansible-galaxy": {
		Name:        "ansible-galaxy",
		Description: ": --role-file (Specify role file), --roles-path (Specify roles path), --server (Specify server), --ignore-errors (Ignore errors), --force (Force operation)",
	},
	"ansible-vault": {
		Name:        "ansible-vault",
		Description: ": --vault-id (Specify vault ID), --new-vault-id (Specify new vault ID), --ask-vault-pass (Ask for vault password), --vault-password-file (Specify vault password file), --output (Specify output file)",
	},
	"ansible-doc": {
		Name:        "ansible-doc",
		Description: ": --module-path (Specify module path), --type (Specify type), --list (List modules), --json (Output in JSON), --quiet (Suppress output)",
	},
	"ansible-config": {
		Name:        "ansible-config",
		Description: ": --list (List configuration), --dump (Dump configuration), --only-changed (Show only changed), --quiet (Suppress output), --json (Output in JSON)",
	},
	"ansible-inventory": {
		Name:        "ansible-inventory",
		Description: ": --list (List inventory), --graph (Show graph), --host (Show host), --yaml (Output in YAML), --json (Output in JSON)",
	},
	"chef": {
		Name:        "chef",
		Description: ": --version (Show version), --help (Show help), --config (Set config), --log-level (Set log level), --environment (Set environment)",
	},
	"chef-client": {
		Name:        "chef-client",
		Description: ": --config (Set config), --log-level (Set log level), --environment (Set environment), --override-runlist (Override runlist), --json-attributes (Specify JSON attributes)",
	},
	"chef-solo": {
		Name:        "chef-solo",
		Description: ": --config (Set config), --log-level (Set log level), --environment (Set environment), --override-runlist (Override runlist), --json-attributes (Specify JSON attributes)",
	},
	"knife": {
		Name:        "knife",
		Description: ": --config (Set config), --log-level (Set log level), --environment (Set environment), --server-url (Set server URL), --user (Specify user)",
	},
	"chef-server-ctl": {
		Name:        "chef-server-ctl",
		Description: ": --config (Set config), --log-level (Set log level), --environment (Set environment), --server-url (Set server URL), --user (Specify user)",
	},
	"chef-apply": {
		Name:        "chef-apply",
		Description: ": --config (Set config), --log-level (Set log level), --environment (Set environment), --override-runlist (Override runlist), --json-attributes (Specify JSON attributes)",
	},
	"chef-shell": {
    	Name:        "chef-shell",
    	Description: ": --client (Run as client), --solo (Run as solo), --config (Specify config file), --log-level (Set log level), --environment (Set environment)",
	},
	"chef-zero": {
		Name:        "chef-zero",
		Description: ": --port (Set port), --host (Set host), --log-level (Set log level), --daemonize (Run as daemon), --config (Specify config file)",
	},
	"puppet": {
		Name:        "puppet",
		Description: ": --version (Show version), --help (Show help), --config (Set config), --log-level (Set log level), --environment (Set environment)",
	},
	"puppet apply": {
		Name:        "puppet apply",
		Description: ": --noop (No operation), --verbose (Verbose output), --debug (Debug output), --modulepath (Set module path), --environment (Set environment)",
	},
	"puppet agent": {
		Name:        "puppet agent",
		Description: ": --test (Run in test mode), --onetime (Run once), --no-daemonize (Do not daemonize), --verbose (Verbose output), --environment (Set environment)",
	},
	"puppet cert": {
		Name:        "puppet cert",
		Description: ": --list (List certificates), --sign (Sign certificate), --revoke (Revoke certificate), --generate (Generate certificate), --clean (Clean certificate)",
	},
	"puppet config": {
		Name:        "puppet config",
		Description: ": --print (Print config), --set (Set config), --delete (Delete config), --environment (Set environment), --section (Specify section)",
	},
	"puppet describe": {
		Name:        "puppet describe",
		Description: ": --list (List resource types), --short (Short description), --providers (List providers), --meta (Show meta parameters), --environment (Set environment)",
	},
	"puppet device": {
		Name:        "puppet device",
		Description: ": --target (Specify target), --user (Specify user), --password (Specify password), --waitforcert (Wait for certificate), --environment (Set environment)",
	},
	"puppet doc": {
		Name:        "puppet doc",
		Description: ": --all (Show all documentation), --outputdir (Set output directory), --mode (Set mode), --environment (Set environment), --modulepath (Set module path)",
	},
	"puppet facts": {
		Name:        "puppet facts",
		Description: ": --terminus (Set terminus), --render-as (Set render format), --environment (Set environment), --modulepath (Set module path), --debug (Debug output)",
	},
	"puppet filebucket": {
		Name:        "puppet filebucket",
		Description: ": --local (Use local filebucket), --remote (Use remote filebucket), --bucket (Specify bucket), --debug (Debug output), --environment (Set environment)",
	},
	"puppet help": {
		Name:        "puppet help",
		Description: ": --all (Show all commands), --verbose (Verbose output), --debug (Debug output), --environment (Set environment), --modulepath (Set module path)",
	},
	"puppet lookup": {
		Name:        "puppet lookup",
		Description: ": --node (Specify node), --facts (Specify facts), --environment (Set environment), --modulepath (Set module path), --debug (Debug output)",
	},
	"puppet module": {
		Name:        "puppet module",
		Description: ": --install (Install module), --uninstall (Uninstall module), --list (List modules), --search (Search modules), --environment (Set environment)",
	},
	"puppet parser": {
		Name:        "puppet parser",
		Description: ": --validate (Validate manifests), --dump (Dump AST), --render-as (Set render format), --environment (Set environment), --modulepath (Set module path)",
	},
	"puppet resource": {
		Name:        "puppet resource",
		Description: ": --types (List resource types), --edit (Edit resource), --environment (Set environment), --modulepath (Set module path), --debug (Debug output)",
	},
	"puppet server": {
		Name:        "puppet server",
		Description: ": --config (Specify config file), --log-level (Set log level), --daemonize (Run as daemon), --environment (Set environment), --debug (Debug output)",
	},
	"puppet status": {
		Name:        "puppet status",
		Description: ": --all (Show all statuses), --verbose (Verbose output), --debug (Debug output), --environment (Set environment), --modulepath (Set module path)",
	},
	"puppet task": {
		Name:        "puppet task",
		Description: ": --run (Run task), --list (List tasks), --environment (Set environment), --modulepath (Set module path), --debug (Debug output)",
	},
	"puppet test": {
		Name:        "puppet test",
		Description: ": --all (Run all tests), --verbose (Verbose output), --debug (Debug output), --environment (Set environment), --modulepath (Set module path)",
	},
	"puppet type": {
		Name:        "puppet type",
		Description: ": --list (List types), --providers (List providers), --environment (Set environment), --modulepath (Set module path), --debug (Debug output)",
	},
}