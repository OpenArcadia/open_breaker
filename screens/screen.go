package screens

type Screen interface {
	Create()
	Render()
	Dispose()
}

var current Screen
var needInit = true

func ChangeScreen(s Screen) {
	if current != nil {
		current.Dispose()
	}
	current = s
	needInit = true
}

func Update() {
	if current == nil {
		return
	}

	if needInit {
		current.Create()
		needInit = false
	}

	current.Render()
}
