package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ocfollow "github.com/desertfox/go-ocfollow/pkg"
	"github.com/desertfox/go-ocfollow/pkg/list"
)

func main() {
	var opts []tea.ProgramOption
	opts = append(opts, tea.WithAltScreen())

	//kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	m := ocfollow.NewModel("", "", list.NewModel([]string{"test", "test2", "test3"}, 0))
	p := tea.NewProgram(m, opts...)

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := p.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
