//go:build windows

package menu

import "github.com/secoba/wails/v2/internal/platform/win32"

type Menu struct {
	menu win32.HMENU
}
