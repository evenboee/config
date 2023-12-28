package config

type config struct {
	filenames []string

	envPrefix string
	varPrefix string

	omitDefaults bool
	omitEnvVars  bool

	ignoreMissingFiles bool

	tagName string

	defaultOverrides map[string]string
}

func defaultOptions() *config {
	return &config{
		filenames:          []string{},
		tagName:            "config",
		ignoreMissingFiles: true,
		defaultOverrides:   make(map[string]string),
	}
}

type Option func(*config)

func WithFilename(filename string) Option {
	return func(c *config) {
		c.filenames = append(c.filenames, filename)
	}
}

func WithFilenames(filenames ...string) Option {
	return func(c *config) {
		c.filenames = append(c.filenames, filenames...)
	}
}

func WithEnvPrefix(prefix string) Option {
	return func(c *config) {
		c.envPrefix = prefix + "_"
	}
}

func WithVarPrefix(prefix string) Option {
	return func(c *config) {
		c.varPrefix = prefix + "_"
	}
}

func WithPrefix(prefix string) Option {
	return func(c *config) {
		WithEnvPrefix(prefix)(c)
		WithVarPrefix(prefix)(c)
	}
}

func OmitDefaults(c *config) {
	c.omitDefaults = true
}

func OmitEnvVars(c *config) {
	c.omitEnvVars = true
}

func WithTagName(name string) Option {
	return func(c *config) {
		c.tagName = name
	}
}

func WithDefaultOverride(key string, value string) Option {
	return func(c *config) {
		c.defaultOverrides[key] = value
	}
}

func WithIgnoreMissingFiles(shouldIgnore bool) Option {
	return func(c *config) {
		c.ignoreMissingFiles = shouldIgnore
	}
}
