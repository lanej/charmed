package story_test

import tea "charm.land/bubbletea/v2"

// Native commands precede the scheduled frame command in the runtime batch.
func applicationMessage(command tea.Cmd) tea.Msg {
	message := command()
	if batch, ok := message.(tea.BatchMsg); ok {
		for _, next := range batch {
			if next != nil {
				return applicationMessage(next)
			}
		}
	}

	return message
}
