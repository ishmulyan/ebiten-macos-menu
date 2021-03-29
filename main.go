package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/objc"
)

type Game struct {
	sceneName string
}

func (g *Game) Update() error { return nil }

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, g.sceneName)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) New() {
	g.sceneName = "New game"
}

func (g *Game) About() {
	g.sceneName = "About"
}

func main() {
	g := &Game{}
	setupMacOSMenu(g)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func setupMacOSMenu(g *Game) {
	// menu handler
	menuHandlerClass := objc.NewClass("MenuHandler", "NSObject")
	menuHandlerClass.AddMethod("about:", func(_ objc.Object) {
		g.About()
	})
	menuHandlerClass.AddMethod("newGame:", func(_ objc.Object) {
		g.New()
	})
	menuHandler := menuHandlerClass.Alloc()

	// main submenu
	mainItem := cocoa.NSMenuItem_New()
	mainSubmenu := cocoa.NSMenu_New()
	mainItem.SetSubmenu(mainSubmenu)

	aboutItem := cocoa.NSMenuItem_Init("About", objc.Sel("about:"), "")
	aboutItem.SetTarget(menuHandler)
	mainSubmenu.AddItem(aboutItem)

	mainSubmenu.AddItem(cocoa.NSMenuItem_Separator())

	quitItem := cocoa.NSMenuItem_Init("Quit", objc.Sel("terminate:"), "q")
	mainSubmenu.AddItem(quitItem)

	// game submenu
	gameItem := cocoa.NSMenuItem_New()
	gameSubmenu := cocoa.NSMenu_New()
	gameSubmenu.SetTitle("Game")
	gameItem.SetSubmenu(gameSubmenu)

	newGameItem := cocoa.NSMenuItem_Init("New", objc.Sel("newGame:"), "n")
	newGameItem.SetTarget(menuHandler)
	gameSubmenu.AddItem(newGameItem)

	app := objc.Get("NSApplication").Get("sharedApplication")
	menu := cocoa.NSMenu_New()
	menu.AddItem(mainItem)
	menu.AddItem(gameItem)
	app.Set("mainMenu:", menu)
}
