package pot

import "fmt"

type Reference struct {
	Base string
	Line int
}

type Entry struct {
	Reference Reference
	Singular  []rune
	Plural    []rune
	Context   []rune
}

func (entry *Entry) PushSingular(character rune) {
	entry.Singular = append(entry.Singular, character)
}

func (entry *Entry) PushPlural(character rune) {
	entry.Plural = append(entry.Plural, character)
}

func (entry *Entry) PushContext(character rune) {
	entry.Singular = append(entry.Context, character)
}

func (entry *Entry) ReferenceAsString() string {
	return fmt.Sprintf("# %s:%d\n", trimPath(entry.Reference.Base), entry.Reference.Line)
}

func (entry *Entry) AsString() string {
	var potEntry string

	if len(entry.Context) > 0 {
		potEntry += fmt.Sprintf("msgctxt \"%s\"\n", string(entry.Context))
	}

	if len(entry.Plural) > 0 {
		potEntry += fmt.Sprintf("msgid \"%s\"\nmsgid_plural \"%s\"\nmsgstr[0] \"\"\nmsgstr[1] \"\"\n\n", string(entry.Singular), string(entry.Plural))
	} else {
		potEntry += fmt.Sprintf("msgid \"%s\"\nmsgstr \"\"\n\n", string(entry.Singular))
	}

	return potEntry
}
