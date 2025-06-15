package search

import (
	"what-cmd/commands"
	"what-cmd/flags"
	"what-cmd/hotkeys"
	"what-cmd/internal/models"
)

func ConvertCommands(cmds []commands.Command) []models.Item {
	items := make([]models.Item, 0, len(cmds))
	for _, cmd := range cmds {
		items = append(items, models.Item{
			Name:        cmd.Name,
			Description: cmd.Description,
			Type:        models.Command,
		})
	}
	return items
}

func ConvertFlags(flagMap map[string]flags.Flag) []models.Item {
	items := make([]models.Item, 0, len(flagMap))
	for name, flag := range flagMap {
		items = append(items, models.Item{
			Name:        name,
			Description: flag.Description,
			Type:        models.Flag,
		})
	}
	return items
}

func ConvertHotkeys(hotkeys []hotkeys.Hotkey) []models.Item {
	items := make([]models.Item, 0, len(hotkeys))
	for _, hk := range hotkeys {
		items = append(items, models.Item{
			Name:        hk.Name,
			Description: hk.Description,
			Type:        models.Hotkey,
		})
	}
	return items
}
