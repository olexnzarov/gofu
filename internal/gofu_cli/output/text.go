package output

type TextOutput struct {
	text string
}

func Text(value string) *TextOutput {
	return &TextOutput{text: value}
}

func (o *Output) Text(key string, value string) *Output {
	return o.Add(key, Text(value))
}

func (o *TextOutput) Text() string {
	return o.text
}

func (o *TextOutput) Object() interface{} {
	return o.text
}
