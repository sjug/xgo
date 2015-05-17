package house

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}
func Verse(subject string, relPhrases []string, nounPhrase string) string {

	if len(relPhrases) == 0 {
		return Embed(subject, nounPhrase)
	}
	return Embed(subject, Verse(relPhrases[0], relPhrases[1:], nounPhrase))

}

func Song() (output string) {
	subject := "This is"
	nounPhrase := "the house that Jack built."
	relPhrases := []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	output += Embed(subject, nounPhrase)
	for n := len(relPhrases) - 1; n >= 0; n-- {
		output += "\n\n" + Verse(subject, relPhrases[n:], nounPhrase)
	}
	return
}
