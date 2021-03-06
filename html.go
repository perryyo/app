package app

// HTMLConfig represents a configuration for an html document.
type HTMLConfig struct {
	// The title.
	Title string

	// The meta data.
	Metas []Meta

	// The css file paths to include.
	// Inludes all .css files in resources/css if nil no set.
	CSS []string

	// The javascript file paths to include.
	// Inludes all .js files in resources/js if not set.
	Javascripts []string
}

// Meta represents a page metadata.
type Meta struct {
	Name      MetaName
	Content   string
	HTTPEquiv MetaHTTPEquiv
}

// MetaName represents a metadata name value.
type MetaName string

// Constants that define metadata name values.
const (
	ApplicationNameMeta MetaName = "application-name"
	AuthorMeta          MetaName = "author"
	DescriptionMeta     MetaName = "description"
	GeneratorMeta       MetaName = "generator"
	KeywordsMeta        MetaName = "keywords"
	ViewportMeta        MetaName = "viewport"
)

// MetaHTTPEquiv represents a metadata http-equiv value.
type MetaHTTPEquiv string

// Constants that define metadata http-equiv values.
const (
	ContentTypeMeta  MetaHTTPEquiv = "content-type"
	DefaultStyleMeta MetaHTTPEquiv = "default-style"
	RefreshMeta      MetaHTTPEquiv = "refresh"
)
