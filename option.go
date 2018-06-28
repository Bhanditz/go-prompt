package prompt

// Option is the type to replace default parameters.
// prompt.New accepts any number of options (this is functional option pattern).
type Option func(prompt *Prompt) error

// OptionParser to set a custom ConsoleParser object. An argument should implement ConsoleParser interface.
func OptionParser(x ConsoleParser) Option {
	return func(p *Prompt) error {
		p.in = x
		return nil
	}
}

// OptionWriter to set a custom ConsoleWriter object. An argument should implement ConsoleWriter interface.
func OptionWriter(x ConsoleWriter) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.out = x
		})
		return nil
	}
}

// OptionTitle to set title displayed at the header bar of terminal.
func OptionTitle(x string) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.title = x
		})
		return nil
	}
}

// OptionPrefix to set prefix string.
func OptionPrefix(x string) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.prefix = x
		})
		return nil
	}
}

// OptionCompletionWordSeparator to set word separators. Enable only ' ' if empty.
func OptionCompletionWordSeparator(x string) Option {
	return func(p *Prompt) error {
		p.completion.wordSeparator = x
		return nil
	}
}

// OptionLivePrefix to change the prefix dynamically by callback function
func OptionLivePrefix(f func() (prefix string, useLivePrefix bool)) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.livePrefixCallback = f
		})
		return nil
	}
}

// OptionPrefixTextColor change a text color of prefix string
func OptionPrefixTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.prefixTextColor = x
		})
		return nil
	}
}

// OptionPrefixBackgroundColor to change a background color of prefix string
func OptionPrefixBackgroundColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.prefixBGColor = x
		})
		return nil
	}
}

// OptionInputTextColor to change a color of text which is input by user
func OptionInputTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.inputTextColor = x
		})
		return nil
	}
}

// OptionInputBGColor to change a color of background which is input by user
func OptionInputBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.inputBGColor = x
		})
		return nil
	}
}

// OptionPreviewSuggestionTextColor to change a text color which is completed
func OptionPreviewSuggestionTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.previewSuggestionTextColor = x
		})
		return nil
	}
}

// OptionPreviewSuggestionBGColor to change a background color which is completed
func OptionPreviewSuggestionBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.previewSuggestionBGColor = x
		})
		return nil
	}
}

// OptionSuggestionTextColor to change a text color in drop down suggestions.
func OptionSuggestionTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.suggestionTextColor = x
		})
		return nil
	}
}

// OptionSuggestionBGColor change a background color in drop down suggestions.
func OptionSuggestionBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.suggestionBGColor = x
		})
		return nil
	}
}

// OptionSelectedSuggestionTextColor to change a text color for completed text which is selected inside suggestions drop down box.
func OptionSelectedSuggestionTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedSuggestionTextColor = x
		})
		return nil
	}
}

// OptionSelectedSuggestionBGColor to change a background color for completed text which is selected inside suggestions drop down box.
func OptionSelectedSuggestionBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedSuggestionBGColor = x
		})
		return nil
	}
}

// OptionDescriptionTextColor to change a background color of description text in drop down suggestions.
func OptionDescriptionTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedDescriptionTextColor = x
		})
		return nil
	}
}

// OptionDescriptionBGColor to change a background color of description text in drop down suggestions.
func OptionDescriptionBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedDescriptionBGColor = x
		})
		return nil
	}
}

// OptionSelectedDescriptionTextColor to change a text color of description which is selected inside suggestions drop down box.
func OptionSelectedDescriptionTextColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedDescriptionTextColor = x
		})
		return nil
	}
}

// OptionSelectedDescriptionBGColor to change a background color of description which is selected inside suggestions drop down box.
func OptionSelectedDescriptionBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.selectedDescriptionBGColor = x
		})
		return nil
	}
}

// OptionScrollbarThumbColor to change a thumb color on scrollbar.
func OptionScrollbarThumbColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.scrollbarThumbColor = x
		})
		return nil
	}
}

// OptionScrollbarBGColor to change a background color of scrollbar.
func OptionScrollbarBGColor(x Color) Option {
	return func(p *Prompt) error {
		p.rendererOptions = append(p.rendererOptions, func(r *Renderer) {
			r.scrollbarBGColor = x
		})
		return nil
	}
}

// OptionMaxSuggestion specify the max number of displayed suggestions.
func OptionMaxSuggestion(x uint16) Option {
	return func(p *Prompt) error {
		p.completion.max = x
		return nil
	}
}

// OptionHistory to set history expressed by string array.
func OptionHistory(x []string) Option {
	return func(p *Prompt) error {
		p.history.histories = x
		p.history.Clear()
		return nil
	}
}

// OptionSwitchKeyBindMode set a key bind mode.
func OptionSwitchKeyBindMode(m KeyBindMode) Option {
	return func(p *Prompt) error {
		p.keyBindMode = m
		return nil
	}
}

// OptionAddKeyBind to set a custom key bind.
func OptionAddKeyBind(b ...KeyBind) Option {
	return func(p *Prompt) error {
		p.keyBindings = append(p.keyBindings, b...)
		return nil
	}
}

// OptionAddASCIICodeBind to set a custom key bind.
func OptionAddASCIICodeBind(b ...ASCIICodeBind) Option {
	return func(p *Prompt) error {
		p.ASCIICodeBindings = append(p.ASCIICodeBindings, b...)
		return nil
	}
}

// New returns a Prompt with powerful auto-completion.
func New(executor Executor, completer Completer, opts ...Option) *Prompt {
	pt := &Prompt{
		in:              NewStandardInputParser(),
		rendererOptions: make([]RendererOption, 0, 12),
		buf:             NewBuffer(),
		executor:        executor,
		history:         NewHistory(),
		completion:      NewCompletionManager(completer, 6),
		keyBindMode:     EmacsKeyBind, // All the above assume that bash is running in the default Emacs setting
	}

	for _, opt := range opts {
		if err := opt(pt); err != nil {
			panic(err)
		}
	}
	return pt
}
