package main

import (
	"log"

	"github.com/aarzilli/nucular"
	"github.com/raedahgroup/dcrseedgen/helper"
)

type App struct {
	currentPage  string
	pageChanged  bool
	masterWindow nucular.MasterWindow
	pages        map[string]page
}

const (
	appName  = "DCR Seed Generator"
	homePage = "seed"

	navPaneWidth            = 220
	contentPaneXOffset      = 25
	contentPaneWidthPadding = 55
)

func main() {
	app := &App{
		pageChanged: true,
		currentPage: homePage,
	}

	// register pages
	pages := getPages()
	app.pages = make(map[string]page, len(pages))
	for _, page := range pages {
		app.pages[page.name] = page
	}

	// load logo once
	err := helper.LoadLogo()
	if err != nil {
		log.Fatal(err)
	}

	window := nucular.NewMasterWindow(nucular.WindowNoScrollbar, appName, app.render)
	if err := helper.InitStyle(window); err != nil {
		log.Fatal(err)
	}

	app.masterWindow = window
	window.Main()
}

func (app *App) changePage(page string) {
	app.currentPage = page
	app.pageChanged = true
	app.masterWindow.Changed()
}

func (app *App) render(window *nucular.Window) {
	currentPage := app.pages[app.currentPage]

	if app.pageChanged {
		currentPage.handler.BeforeRender()
		app.pageChanged = false
	}

	helper.DrawPageHeader(window)
	window.Row(38).Dynamic(3)

	window.Label("", "LC")
	if window.ButtonText("Generate Seed") && app.currentPage != "seed" {
		app.changePage("seed")
	}

	if window.ButtonText("Generate Address") && app.currentPage != "address" {
		app.changePage("address")
	}
	currentPage.handler.Render(window)
}
