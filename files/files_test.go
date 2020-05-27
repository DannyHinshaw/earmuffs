package files

import (
	"strings"
	"testing"
)

func TestGetWordsDirPath(t *testing.T) {
	path := GetWordsDirPath()
	if !strings.HasSuffix(path, ProjectName+"/data/") {
		t.Fatal("excludes is an empty array ")
	}
}

func TestLoadExcludes(t *testing.T) {
	path := GetWordsDirPath()
	excludes := LoadExcludes(path + "/words/excludes.txt")
	if len(excludes) < 1 {
		t.Fatal("excludes file is empty")
	}
}

func TestLoadProfanities(t *testing.T) {
	path := GetWordsDirPath()
	excludes := LoadExcludes(path + "/words/bad-words.txt")
	if len(excludes) < 1 {
		t.Fatal("profanities is an empty array ")
	}
}

func TestGetWordsPayload(t *testing.T) {
	payload := GetWordsPayload()
	lenProfanity := len(payload.ProfanitiesRegex)
	if lenProfanity < 1 {
		t.Fatal("profanities is an empty array ")
	}

	lenExcludes := len(payload.ExcludeWords)
	if lenExcludes < 1 {
		t.Fatal("excludes file is empty")
	}
}
