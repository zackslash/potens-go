package i18n

type Translatable struct {
	defaultLanguage string
	translations    Translations
}

type Translations map[string]string

// NewTranslatable Create a new translatable object
func NewTranslatable(translations Translations) *Translatable {
	t := &Translatable{
		translations:    translations,
		defaultLanguage: "en",
	}
	return t
}

func (t *Translatable) SetDefaultLanguage(language string) string {
	t.defaultLanguage = language
	return language
}

func (t *Translatable) Clear() {
	t.translations = make(Translations)
}
func (t *Translatable) Set(language string, translation string) {
	if t.translations == nil {
		t.Clear()
	}
	t.translations[language] = translation
}

func (t *Translatable) Get(language string) string {
	if val, ok := t.translations[language]; ok {
		return val
	}
	if val, ok := t.translations[language[:2]]; ok {
		return val
	}
	if len(t.defaultLanguage) < 2 {
		t.defaultLanguage = "en"
	}
	if val, ok := t.translations[t.defaultLanguage]; ok {
		return val
	}
	return ""
}
