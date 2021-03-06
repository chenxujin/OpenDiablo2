package d2ui

import (
	"github.com/OpenDiablo2/D2Shared/d2common/d2enum"
	"github.com/OpenDiablo2/D2Shared/d2common/d2interface"
	"github.com/OpenDiablo2/D2Shared/d2common/d2resource"
	"github.com/OpenDiablo2/D2Shared/d2data/d2datadict"
	"github.com/OpenDiablo2/OpenDiablo2/d2render"
	"github.com/hajimehoshi/ebiten"
)

type Scrollbar struct {
	x, y, height    int
	visible         bool
	enabled         bool
	currentOffset   int
	maxOffset       int
	lastDirChange   int
	onActivate      func()
	scrollbarSprite d2render.Sprite
}

func CreateScrollbar(fileProvider d2interface.FileProvider, x, y, height int) Scrollbar {
	result := Scrollbar{
		visible:         true,
		enabled:         true,
		x:               x,
		y:               y,
		height:          height,
		scrollbarSprite: d2render.CreateSprite(fileProvider.LoadFile(d2resource.Scrollbar), d2datadict.Palettes[d2enum.Sky]),
	}
	return result
}

func (v Scrollbar) GetEnabled() bool {
	return v.enabled
}

func (v *Scrollbar) SetEnabled(enabled bool) {
	v.enabled = enabled
}

func (v *Scrollbar) SetPressed(pressed bool) {}
func (v *Scrollbar) GetPressed() bool        { return false }

func (v *Scrollbar) OnActivated(callback func()) {
	v.onActivate = callback
}

func (v Scrollbar) getBarPosition() int {
	return int((float32(v.currentOffset) / float32(v.maxOffset)) * float32(v.height-30))
}

func (v *Scrollbar) Activate() {
	_, my := ebiten.CursorPosition()
	barPosition := v.getBarPosition()
	if my <= v.y+barPosition+15 {
		if v.currentOffset > 0 {
			v.currentOffset--
			v.lastDirChange = -1
		}
	} else {
		if v.currentOffset < v.maxOffset {
			v.currentOffset++
			v.lastDirChange = 1
		}
	}
	if v.onActivate != nil {
		v.onActivate()
	}
}

func (v Scrollbar) GetLastDirChange() int {
	return v.lastDirChange
}

func (v *Scrollbar) Draw(target *ebiten.Image) {
	if !v.visible || v.maxOffset == 0 {
		return
	}
	offset := 0
	if !v.enabled {
		offset = 2
	}
	v.scrollbarSprite.MoveTo(v.x, v.y)
	v.scrollbarSprite.DrawSegments(target, 1, 1, 0+offset)
	v.scrollbarSprite.MoveTo(v.x, v.y+v.height-10)
	v.scrollbarSprite.DrawSegments(target, 1, 1, 1+offset)
	if v.maxOffset == 0 || v.currentOffset < 0 || v.currentOffset > v.maxOffset {
		return
	}
	v.scrollbarSprite.MoveTo(v.x, v.y+10+v.getBarPosition())
	offset = 0
	if !v.enabled {
		offset = 1
	}
	v.scrollbarSprite.DrawSegments(target, 1, 1, 4+offset)
}

func (v *Scrollbar) GetSize() (width, height uint32) {
	return 10, uint32(v.height)
}

func (v *Scrollbar) MoveTo(x, y int) {
	v.x = x
	v.y = y
}

func (v *Scrollbar) GetLocation() (x, y int) {
	return v.x, v.y
}

func (v *Scrollbar) GetVisible() bool {
	return v.visible
}

func (v *Scrollbar) SetVisible(visible bool) {
	v.visible = visible
}

func (v *Scrollbar) SetMaxOffset(maxOffset int) {
	v.maxOffset = maxOffset
	if v.maxOffset < 0 {
		v.maxOffset = 0
	}
	if v.currentOffset > v.maxOffset {
		v.currentOffset = v.maxOffset
	}
	if v.maxOffset == 0 {
		v.currentOffset = 0
	}
}

func (v *Scrollbar) SetCurrentOffset(currentOffset int) {
	v.currentOffset = currentOffset
}

func (v Scrollbar) GetMaxOffset() int {
	return v.maxOffset
}

func (v Scrollbar) GetCurrentOffset() int {
	return v.currentOffset
}
