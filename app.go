package main

import (
	"changeme/pkg/applications"
	"code.rocketnine.space/tslocum/desktop"
	"context"
	"fmt"
	"github.com/robotn/gohook"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/execabs"
)

// App struct
type App struct {
	ctx   context.Context
	show  bool
	focus bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.show = true

	go func() {
		hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
			if a.show {
				a.Hide()
			} else {
				a.Show()
			}
		})
		//hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		//	if a.show {
		//		runtime.WindowHide(a.ctx)
		//		a.show = false
		//	}
		//})
		s := hook.Start()
		<-hook.Process(s)
	}()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Hide() {
	runtime.WindowHide(a.ctx)
	a.show = false
	a.ToBlur()
}

func (a *App) Show() {
	runtime.WindowShow(a.ctx)
	runtime.WindowSetSize(a.ctx, 800, 600)
	runtime.WindowCenter(a.ctx)
	a.show = true
	runtime.EventsEmit(a.ctx, "show")
}

func (a *App) ToFocus() {
	a.focus = true
}

func (a *App) ToBlur() {
	a.focus = false
}

func (a *App) ListApplications() ([]*desktop.Entry, error) {
	return applications.List()
}

func (a *App) RunApplication(cmd string) {
	go execabs.Command("sh", "-c", cmd).Run()
}

func (a *App) domReady(ctx context.Context) {
	a.Show()
}
