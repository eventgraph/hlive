package domhelpers

import (
	"github.com/SamHennessy/hlive"
)

func OnClick(onClick hlive.EventHandler) *hlive.EventBinding {
	return hlive.On("click", onClick)
}

func OnKeyDown(onKeyDown hlive.EventHandler) *hlive.EventBinding {
	return hlive.On("keydown", onKeyDown)
}

func OnKeyUp(onKeyUp hlive.EventHandler) *hlive.EventBinding {
	return hlive.On("keyup", onKeyUp)
}
