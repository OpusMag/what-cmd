package models

type Item struct {
	Name        string
	Description string
	Type        ItemType
}

type ItemType int

const (
	Command ItemType = iota
	Flag
	Hotkey
)

func (t ItemType) String() string {
	switch t {
	case Command:
		return "command"
	case Flag:
		return "flag"
	case Hotkey:
		return "hotkey"
	default:
		return "unknown"
	}
}
