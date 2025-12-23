package xmljsontransformation

type Options struct {
	RootName    string
	PrettyPrint bool

	// XML → JSON options
	AttrPrefix string
	TextKey    string
}

type Option func(*Options)

func defaultOptions() *Options {
	return &Options{
		RootName:    "root",
		PrettyPrint: false,
		AttrPrefix: "@",
		TextKey:    "#text",
	}
}

func WithRoot(name string) Option {
	return func(o *Options) {
		o.RootName = name
	}
}

func WithPrettyPrint(v bool) Option {
	return func(o *Options) {
		o.PrettyPrint = v
	}
}

func WithAttrPrefix(p string) Option {
	return func(o *Options) {
		o.AttrPrefix = p
	}
}

func WithTextKey(k string) Option {
	return func(o *Options) {
		o.TextKey = k
	}
}
