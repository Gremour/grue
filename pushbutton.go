package grue

// PushButton is pressable and optionally, checkable (TODO), button.
type PushButton struct {
	*Panel
	Highlited bool
	Pressed   bool

	OnPress func()
}

// NewPushButton creates new button.
func NewPushButton(parent Widget, b Base) *PushButton {
	pb := &PushButton{
		Panel: NewPanel(nil, b),
	}
	InitWidget(parent, pb)

	pb.OnMouseIn = func() {
		pb.Highlited = true
	}
	pb.OnMouseOut = func() {
		pb.Highlited = false
		pb.Pressed = false
	}
	pb.OnMouseDown = func(bt Button) {
		if bt != MouseButtonLeft {
			return
		}
		pb.Pressed = true
	}
	pb.OnMouseUp = func(bt Button) {
		if bt != MouseButtonLeft {
			return
		}
		pb.Pressed = false
		if pb.OnPress != nil {
			pb.OnPress()
		}
	}
	return pb
}

// Paint draws the widget without children.
func (pb *PushButton) Paint() {
	r := pb.GlobalRect()
	theme := pb.Theme
	if theme == nil {
		theme = pb.Surface.GetTheme()
	}
	tdef, _ := theme.Drawers[ThemeButton]
	var tcur ThemeDrawer
	tcol := theme.TextColor
	switch {
	case pb.Disabled:
		tcur, _ = theme.Drawers[ThemeButtonDisabled]
		tcol = theme.DisabledTextColor
	case pb.Pressed:
		tcur, _ = theme.Drawers[ThemeButtonActive]
	case pb.Highlited:
		tcur, _ = theme.Drawers[ThemeButtonHL]
	}
	if tcur != nil {
		tdef = tcur
	}
	if tdef != nil {
		tdef.Draw(pb.Surface, r)
	}
	if len(pb.Text) > 0 {
		pb.Surface.DrawText(pb.Text, theme.TitleFont, r, tcol, AlignCenter, AlignCenter)
	}
}
