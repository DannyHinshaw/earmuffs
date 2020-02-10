package files

import (
	"testing"
)

func TestCheckLineProfanity(t *testing.T) {
	profane := "fuck this shit"
	match1 := CheckLineProfanity(profane, 1)
	if len(match1.Profanities) < 2 {
		t.Fatal("did not detect profane words")
	}

	clean := "testing this sentence for parser"
	match2 := CheckLineProfanity(clean, 1)
	if len(match2.Profanities) > 0 {
		t.Fatal("falsely detected profane words")
	}
}

func TestListFiles(t *testing.T) {
	res, err := ListFiles("")
	if len(res) > 0 || err != nil {
		t.Fail()
	}
}
