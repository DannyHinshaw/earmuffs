package schemas

type WordsPayload struct {
	ProfanitiesRegex string
	ExcludeWords     []string
}

type LineMatch struct {
	Profanities []string
	Line        int
}

type ResultsMessage struct {
	Matches []LineMatch
	Path    string
}
