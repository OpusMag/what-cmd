package models

type Item struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        ItemType `json:"type"`
	Source      string   `json:"source,omitempty"`   // File path or source location
	Tags        []string `json:"tags,omitempty"`     // Additional metadata
	Priority    int      `json:"priority,omitempty"` // Search ranking priority
}

type ItemType int

const (
	Command ItemType = iota
	Flag
	Hotkey
	SystemCommand  // Discovered from system PATH
	CustomCommand  // From user-specified paths
	UserAlias      // From shell configuration files
	UserFunction   // Shell functions
	PackageManager // Commands from package managers
)

func (t ItemType) String() string {
	switch t {
	case Command:
		return "builtin-command"
	case Flag:
		return "flag"
	case Hotkey:
		return "hotkey"
	case SystemCommand:
		return "system-command"
	case CustomCommand:
		return "custom-command"
	case UserAlias:
		return "user-alias"
	case UserFunction:
		return "user-function"
	case PackageManager:
		return "package-command"
	default:
		return "unknown"
	}
}

func (i *Item) IsUserDefined() bool {
	return i.Type == UserAlias || i.Type == UserFunction
}

func (i *Item) IsSystemDiscovered() bool {
	return i.Type == SystemCommand || i.Type == CustomCommand
}
