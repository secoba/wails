package app

import (
	"context"

	"github.com/secoba/wails/v2/internal/frontend"
	"github.com/secoba/wails/v2/internal/logger"
	"github.com/secoba/wails/v2/internal/menumanager"
	"github.com/secoba/wails/v2/pkg/menu"
	"github.com/secoba/wails/v2/pkg/options"
)

// App defines a Wails application structure
type App struct {
	frontend frontend.Frontend
	logger   *logger.Logger
	options  *options.App

	menuManager *menumanager.Manager

	// Indicates if the app is in debug mode
	debug bool

	// Indicates if the devtools is enabled
	devtoolsEnabled bool

	// OnStartup/OnShutdown
	startupCallback  func(ctx context.Context)
	shutdownCallback func(ctx context.Context)
	ctx              context.Context
}

// Shutdown the application
func (a *App) Shutdown() {
	if a.frontend != nil {
		a.frontend.Quit()
	}
}

// SetApplicationMenu sets the application menu
func (a *App) SetApplicationMenu(menu *menu.Menu) {
	if a.frontend != nil {
		a.frontend.MenuSetApplicationMenu(menu)
	}
}
