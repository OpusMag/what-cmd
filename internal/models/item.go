package models

type Item struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        ItemType `json:"type"`
	Source      string   `json:"source,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Priority    int      `json:"priority,omitempty"`
}

type ItemType int

const (
	Command ItemType = iota
	Flag
	Hotkey
	SystemCommand
	CustomCommand
	UserAlias
	UserFunction 
	PackageManager
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
