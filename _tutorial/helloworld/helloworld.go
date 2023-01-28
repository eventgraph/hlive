package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SamHennessy/hlive"
	l "github.com/SamHennessy/hlive"
	. "github.com/SamHennessy/hlive/domhelpers"
	"github.com/rs/zerolog/log"
)

// Step 1
func homeStep1() *l.Page {
	page := l.NewPage()
	page.DOM().Body().Add("Hello, world!")

	// Ignore this, well explain it later
	addStageButtons(page, 1)

	return page
}

// Step 2
func homeStep2() *l.Page {
	var message string

	input := l.NewComponent("input")
	input.Add(l.Attrs{"type": "text"})
	input.Add(l.On("keyup", func(ctx context.Context, e l.Event) {
		message = e.Value
	}))

	page := l.NewPage()
	page.DOM().Body().Add(l.NewTag("div", input))
	page.DOM().Body().Add(l.T("hr"))
	page.DOM().Body().Add("Hello, ", &message)

	// Ignore this, well explain it later
	addStageButtons(page, 2)

	return page
}

// Step 2.1
func homeStep3() *l.Page {
	var message string

	page := l.NewPage()
	page.DOM().Body().Add(
		l.T(
			"div",
			l.C("input",
				l.Attrs{"type": "text"},
				l.On("keyup", func(_ context.Context, e l.Event) {
					message = e.Value
				}),
			)),
		l.T("hr"),
		"Hello!!, ", &message,
	)

	// Ignore this, well explain it later
	addStageButtons(page, 3)

	return page
}

// Use the DSL
func homeStep4() *l.Page {
	var message string

	page := l.NewPage()
	page.DOM().Body().Add(
		Division(
			InputComponent(
				InputAttributes(
					AttrType, "text",
				),
				OnKeyDown(func(ctx context.Context, e l.Event) {
					message = e.Value
				}),
			),
		),
		HorizontalRule(),
		"Hello, ", &message,
	)

	// Ignore this, well explain it later
	addStageButtons(page, 4)

	return page
}

func main() {
	http.Handle("/step1", l.NewPageServer(homeStep1))
	http.Handle("/step2", l.NewPageServer(homeStep2))
	http.Handle("/step3", l.NewPageServer(homeStep3))
	http.Handle("/step4", l.NewPageServer(homeStep4))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/step1", http.StatusFound)
	}))

	log.Info().Msg("Listing on :3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Info().Err(err)
	}
}

func addStageButtons(page *hlive.Page, stage int) {
	body := page.DOM().Body()

	body.Add(HorizontalRule())

	if stage > 1 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					"onclick", fmt.Sprintf("window.location.href = '/step%d'", stage-1),
				),
				"Previous",
			),
		)
	}

	if stage < 4 {
		body.Add(
			ButtonComponent(
				ButtonAttributes(
					"onclick", fmt.Sprintf("window.location.href = '/step%d'", stage+1),
				),
				"Next",
			),
		)
	}
}
