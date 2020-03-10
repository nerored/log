package log

import "testing"

func TestPrintInfo(t *testing.T) {
	Info("hero world this is %v -----l",
		NewCombo("combo with red", FGC_RED))

	Debu("hello world this is %v",
		NewCombo("combo with blue", FGC_BLUE))

	Trac("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))

	Fata("hello world this is %v and %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD),
		NewCombo("combo with2 3", FMT_UNDERLINED))

	Erro("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))

	Warn("hello world this is %v and %v",
		NewCombo("combo with 1", FGC_GREEN),
		NewCombo("combo with2 2", FMT_BOLD))
}
