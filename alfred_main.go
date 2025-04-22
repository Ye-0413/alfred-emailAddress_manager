package main

import (
	"os"
	"strings"

	"github.com/crispgm/alfred-latex-table/internal"

	aw "github.com/deanishe/awgo"
)

var (
	wf *aw.Workflow
)

func run() {
	query := ""
	if len(os.Args) > 1 {
		query = strings.TrimSpace(os.Args[1])
	}

	// Get all email addresses
	emails := internal.GetEmailAddresses()

	// Filter emails based on query
	for _, email := range emails {
		if query == "" || strings.Contains(strings.ToLower(email.Name), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(email.Email), strings.ToLower(query)) {
			wf.NewItem(email.Name).
				Subtitle(email.Email).
				Arg(email.Email).
				Valid(true)
		}
	}

	wf.SendFeedback()
	return
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
