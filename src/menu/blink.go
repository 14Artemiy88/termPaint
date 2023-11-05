package menu

var (
	BlinkCursor string
	BlinkPhase  bool
	BlinkTime   = DefBlinkTime
	blink       = map[bool]string{
		true:  "|",
		false: " ",
	}
)

const DefBlinkTime = 50

func Blink() {
	BlinkCursor = blink[BlinkPhase]
	BlinkTime--
	if BlinkTime == 0 {
		BlinkPhase = !BlinkPhase
		BlinkTime = DefBlinkTime
	}
}
