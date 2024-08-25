package commands

// Command represents a command and its description
type Command struct {
    Name        string
    Description string
}

// List of commands with descriptions
var Words = []Command{
    {"sudo", "Super User Do. Allows a user to run commands as an administrator"},
    {"cat", "Concatenate and display the content of files"},
    {"ls", "List directory contents"},
    {"cd", "Change the current directory"},
    {"pwd", "Print the current working directory"},
    {"touch", "Create an empty file"},
    {"mkdir", "Create a new directory"},
    {"rm", "Remove files or directories"},
    {"cp", "Copy files or directories"},
    {"mv", "Move or rename files or directories"},
    {"clear", "Clear the terminal screen"},
    {"exit", "Exit the shell"},
    {"echo", "Display a line of text"},
    {"man", "Display the manual for a command"},
    {"grep", "Search text using patterns"},
    {"find", "Search for files in a directory hierarchy"},
    {"chmod", "Change file modes or Access Control Lists"},
    {"chown", "Change file owner and group"},
    {"df", "Report file system disk space usage"},
    {"du", "Estimate file space usage"},
    {"head", "Output the first part of files"},
    {"tail", "Output the last part of files"},
    {"tar", "Store, list or extract files in an archive"},
    {"wget", "Retrieve files from the web"},
    {"curl", "Transfer data from or to a server"},
    {"ps", "Report a snapshot of current processes"},
    {"kill", "Send a signal to a process"},
    {"top", "Display Linux tasks"},
    {"htop", "Interactive process viewer"},
    {"ssh", "OpenSSH SSH client (remote login program)"},
    {"scp", "Secure copy (remote file copy program)"},
    {"nano", "Simple text editor"},
    {"vim", "Vi IMproved, a programmer's text editor"},
    {"apt-get", "A package handling utility for Debian-based distributions"},
    {"yum", "Package manager for RPM-based distributions"},
    {"ping", "Send ICMP ECHO_REQUEST to network hosts"},
    {"ifconfig", "Configure a network interface"},
    {"ip", "Show/manipulate routing, devices, policy routing and tunnels"},
    {"netstat", "Print network connections, routing tables, interface statistics, masquerade connections, and multicast memberships"},
    {"ss", "Another utility to investigate sockets"},
    {"mount", "Mount a file system"},
    {"umount", "Unmount file systems"},
    {"uname", "Print system information"},
    {"uptime", "Tell how long the system has been running"},
    {"whoami", "Print the current user id and name"},
    {"hostname", "Show or set the system's host name"},
    {"history", "Show the command history"},
    {"alias", "Create an alias for a command"},
    {"unalias", "Remove an alias"},
    {"env", "Show the environment variables"},
    {"export", "Set an environment variable"},
    {"source", "Execute commands from a file in the current shell"},
    {"su", "Switch user"},
    {"passwd", "Change user password"},
    {"useradd", "Create a new user"},
    {"usermod", "Modify a user account"},
    {"userdel", "Delete a user account"},
    {"groupadd", "Create a new group"},
    {"groupmod", "Modify a group"},
    {"groupdel", "Delete a group"},
    {"crontab", "Schedule periodic background jobs"},
    {"at", "Schedule a command to run once at a later time"},
    {"jobs", "List active jobs"},
    {"bg", "Place a job in the background"},
    {"fg", "Place a job in the foreground"},
    {"killall", "Kill processes by name"},
    {"xargs", "Build and execute command lines from standard input"},
    {"tee", "Read from standard input and write to standard output and files"},
    {"diff", "Compare files line by line"},
    {"patch", "Apply a diff file to an original"},
    {"sed", "Stream editor for filtering and transforming text"},
    {"awk", "Pattern scanning and processing language"},
    {"tr", "Translate or delete characters"},
    {"cut", "Remove sections from each line of files"},
    {"sort", "Sort lines of text files"},
    {"uniq", "Report or omit repeated lines"},
    {"wc", "Print newline, word, and byte counts for each file"},
    {"basename", "Strip directory and suffix from filenames"},
    {"dirname", "Strip last component from file name"},
    {"readlink", "Print value of a symbolic link or canonical file name"},
    {"ln", "Make links between files"},
    {"stat", "Display file or file system status"},
    {"file", "Determine file type"},
    {"strings", "Print the strings of printable characters in files"},
    {"md5sum", "Compute and check MD5 message digest"},
    {"sha256sum", "Compute and check SHA256 message digest"},
    {"gzip", "Compress files"},
    {"gunzip", "Decompress files"},
    {"bzip2", "Compress files with bzip2"},
    {"bunzip2", "Decompress files with bzip2"},
    {"zip", "Package and compress files"},
    {"unzip", "Extract compressed files"},
    {"rar", "Archive manager for .rar files"},
    {"unrar", "Extract .rar files"},
    {"7z", "File archiver with high compression ratio"},
    {"7za", "Standalone version of 7z"},
    {"7zr", "Reduced version of 7z"},
    {"free", "Display memory usage"},
    {"who", "Show who is logged on"},
    {"last", "Show listing of last logged in users"},
    {"dmesg", "Print or control the kernel ring buffer"},
    {"lsblk", "List information about block devices"},
    {"blkid", "Locate/print block device attributes"},
    {"fdisk", "Partition table manipulator for Linux"},
    {"mkfs", "Build a Linux file system"},
    {"fsck", "File system consistency check and repair"},
    {"parted", "A partition manipulation program"},
    {"dd", "Convert and copy a file"},
    {"sync", "Flush file system buffers"},
    {"shutdown", "Halt, power-off or reboot the machine"},
    {"reboot", "Reboot the system"},
    {"systemctl", "Examine and control the systemd system and service manager"},
    {"journalctl", "Query and display messages from the journal"},
    {"timedatectl", "Control the system time and date"},
    {"hostnamectl", "Control the system hostname"},
    {"localectl", "Control the system locale and keyboard layout settings"},
    {"loginctl", "Control the systemd login manager"},
    {"nmcli", "Command-line client for NetworkManager"},
    {"iwconfig", "Configure wireless network interfaces"},
    {"iwlist", "Get more detailed wireless information"},
    {"traceroute", "Print the route packets take to the network host"},
    {"mtr", "Network diagnostic tool"},
    {"dig", "DNS lookup"},
    {"nslookup", "Query Internet name servers interactively"},
    {"host", "DNS lookup utility"},
    {"rsync", "Remote file and directory synchronization"},
    {"ftp", "File Transfer Protocol client"},
    {"sftp", "Secure File Transfer Protocol"},
    {"telnet", "User interface to the TELNET protocol"},
    {"nc", "Netcat, a versatile networking tool"},
    {"nmap", "Network exploration tool and security/port scanner"},
    {"tcpdump", "Command-line packet analyzer"},
    {"wireshark", "Network protocol analyzer"},
    {"iptables", "Administration tool for IPv4 packet filtering and NAT"},
    {"firewalld", "A firewall management tool"},
    {"ufw", "Uncomplicated Firewall, a frontend for iptables"},
    {"fail2ban", "Ban IPs that show malicious signs"},
    {"logrotate", "Rotates, compresses, and mails system logs"},
    {"cron", "Daemon to execute scheduled commands"},
    {"anacron", "Run commands periodically"},
    {"systemd-analyze", "Analyze system boot-up performance"},
    {"pkill", "Send signal to processes based on name and other attributes"},
    {"pgrep", "Look up or signal processes based on name and other attributes"},
    {"nice", "Run a program with modified scheduling priority"},
    {"renice", "Alter priority of running processes"},
    {"ionice", "Set or get I/O scheduling class and priority"},
    {"watch", "Execute a program periodically, showing output fullscreen"},
    {"screen", "Terminal multiplexer"},
    {"tmux", "Terminal multiplexer"},
    {"nohup", "Run a command immune to hangups, with output to a non-tty"},
    {"disown", "Remove jobs from current shell"},
    {"ethtool", "Display or change Ethernet device settings"},
    {"arp", "Manipulate the system ARP cache"},
    {"whois", "Client for the whois directory service"},
    {"ipcalc", "Calculate IP information for a given IP address"},
    {"netcat", "Utility for reading from and writing to network connections using TCP or UDP"},
    {"arping", "Send ARP REQUEST to a neighbor host"},
    {"iw", "Show/manipulate wireless devices and their configuration"},
    {"bridge", "Show/manipulate bridge addresses and devices"},
    {"tc", "Show/manipulate traffic control settings"},
    {"ipset", "Administration tool for IP sets"},
    {"ip rule", "Routing policy database management"},
    {"ip addr", "Show/manipulate IP addresses"},
    {"ip link", "Show/manipulate network devices"},
    {"ip route", "Show/manipulate routing tables"},
    {"ip neigh", "Show/manipulate neighbor objects"},
    {"ip maddr", "Show/manipulate multicast addresses"},
    {"ip monitor", "Watch for netlink messages"},
    {"ip tunnel", "Show/manipulate tunnel devices"},
    {"ip xfrm", "Show/manipulate IPSec policies"},
    {"ip netns", "Show/manipulate network namespaces"},
    {"ip vrf", "Show/manipulate VRF devices"},
    {"ip link set", "Change device attributes"},
    {"ip link add", "Add a new device"},
    {"ip link delete", "Delete a device"},
    {"ip link show", "Display device attributes"},
    {"ip link set dev", "Change device attributes"},
    {"ip link set up", "Enable a device"},
    {"ip link set down", "Disable a device"},
    {"ip link set mtu", "Change device MTU"},
    {"ip link set address", "Change device MAC address"},
    {"ip link set master", "Set master device"},
    {"ip link set nomaster", "Unset master device"},
    {"ip link set type", "Change device type"},
    {"ip link set netns", "Move device to network namespace"},
    {"ip link set alias", "Set device alias"},
    {"ip link set group", "Set device group"},
    {"ip link set txqueuelen", "Change device transmit queue length"},
    {"ip link set vf", "Change virtual function attributes"},
    {"ip link set xdp", "Attach/detach XDP program"},
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