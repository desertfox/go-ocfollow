package list

import "fmt"

func (m Model) View() string {
	var output string

	for i, v := range m.pods {
		if i == m.cursor {
			output += fmt.Sprintf("-> %v\n", v)
		} else {
			output += fmt.Sprintf("%v\n", v)
		}
	}

	return output
}
