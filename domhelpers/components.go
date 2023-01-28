package domhelpers

import "github.com/SamHennessy/hlive"

// <a> tag defines a hyperlink, which is used to link from one page to another.
func HyperlinkComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("a", elements...)
}

// <abbr> tag defines an abbreviation or an acronym, like "Mr.", "Dec.", "ASAP", "ATM".
func AbbreviationComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("abbr", elements...)
}

// <address> tag defines the contact information for its nearest article or body element ancestor.
func AddressComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("address", elements...)
}

// <area> tag defines a hot-spot region on an image, and optionally associates it with a hypertext link.
func AreaComponent(elements ...any) *hlive.Component { return hlive.NewComponent("area", elements...) }

// <article> element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable.
func ArticleComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("article", elements...)
}

// <aside> element represents a section of the web page that encloses content which is tangentially related to the content around it.
func AsideComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("aside", elements...)
}

// <audio> element is used to embed audio content in an HTML document without requiring any additional plug-in like Flash player.
func AudioComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("audio", elements...)
}

// <b> (short for bold) tag displays text in a bold style. This element typically renders the text it encloses in a bold typeface without conveying any extra importance.
func BoldComponent(elements ...any) *hlive.Component { return hlive.NewComponent("b", elements...) }

// <base> tag defines the base URL and a common target to for all relative URLs contained within a document. There must be no more than one <base> tag per document.
func BaseComponent(elements ...any) *hlive.Component { return hlive.NewComponent("base", elements...) }

// <bdi> (stands for bi-directional isolation) element represents a span of text that is to be isolated from its surroundings for the purposes of bidirectional text formatting.
func BidirectionalIsolationComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("bdi", elements...)
}

// <bdo> (stands for bi-directional override) element represents the override of the current directionality.
func BidirectionalOverrideComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("bdo", elements...)
}

// <blockquote> tag defines a section that is quoted from another source.
func BlockQuoteComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("blockquote", elements...)
}

// <br> tag produces a line break in text.
func LineBreakComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("br", elements...)
}

// <button> tag defines a clickable button.
func ButtonComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("button", elements...)
}

// <canvas> tag is used to draw graphics, on the fly, via scripting (usually JavaScript).
func CanvasComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("canvas", elements...)
}

// <caption> tag defines a table caption.
func CaptionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("caption", elements...)
}

// <cite> tag defines the title of a work (e.g. a book, a song, a movie, a painting, a sculpture, etc.).
func CitationComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("cite", elements...)
}

// <code> tag defines a piece of computer code.
func CodeComponent(elements ...any) *hlive.Component { return hlive.NewComponent("code", elements...) }

// <col> tag defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
func ColumnComponent(elements ...any) *hlive.Component { return hlive.NewComponent("col", elements...) }

// <colgroup> tag defines a group of columns within a table.
func ColumnGroupComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("colgroup", elements...)
}

// <data> tag links a given content with a machine-readable translation. If the content is time- or date-related, the <time> element must be used.
func DataComponent(elements ...any) *hlive.Component { return hlive.NewComponent("data", elements...) }

// <datalist> tag specifies a list of pre-defined options for input controls.
func DataListComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("datalist", elements...)
}

// <dd> tag defines a description/value of a term in a description list.
func DefinitionDescriptionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("dd", elements...)
}

// <del> tag defines text that has been deleted from a document.
func DeletedComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("del", elements...)
}

// <details> tag specifies additional details that the user can view or hide on demand.
func DetailsComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("details", elements...)
}

// <dfn> (short for definition) tag represents the defining instance of a term.
func DefinitionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("dfn", elements...)
}

// <dialog> tag defines a dialog box or window.
func DialogComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("dialog", elements...)
}

// <div> tag defines a section in a document.
func DivisionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("div", elements...)
}

// <dl> (short for definition list) tag defines a description list.
func DefinitionListComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("dl", elements...)
}

// <dt> (short for definition term) tag represent a term or an item in a definition list.
func DefinitionTermComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("dt", elements...)
}

// <em> (short for emphasis) tag is used to emphasize the text content.
func EmphasisComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("em", elements...)
}

// <embed> tag defines a container for an external application or interactive content (a plug-in).
func EmbedComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("embed", elements...)
}

// <fieldset> tag is used to group several controls as well as labels (<label>) within a web form.
func FieldSetComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("fieldset", elements...)
}

// <figcaption> tag represents a caption or legend describing the rest of the contents of its parent <figure> element.
func FigureCaptionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("figcaption", elements...)
}

// <figure> tag specifies self-contained content, potentially with an optional caption, which is specified using the <figcaption> element.
func FigureComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("figure", elements...)
}

// <footer> tag represents a footer for its nearest sectioning content or sectioning root element.
func FooterComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("footer", elements...)
}

// <form> tag represents a document section containing interactive controls for submitting information to a web server.
func FormComponent(elements ...any) *hlive.Component { return hlive.NewComponent("form", elements...) }

// <hgroup> tag represents a multi-level heading for a section of a document.
func HeadingGroupComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("hgroup", elements...)
}

// <h1> is the highest section level
func Heading1Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h1", elements...)
}

// <h2> is the second highest section level
func Heading2Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h2", elements...)
}

// <h3> is the third highest section level
func Heading3Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h3", elements...)
}

// <h4> is the fourth highest section level
func Heading4Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h4", elements...)
}

// <h5> is the fifth highest section level
func Heading5Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h5", elements...)
}

// <h6> is the sixth highest section level
func Heading6Component(elements ...any) *hlive.Component {
	return hlive.NewComponent("h6", elements...)
}

// <hr> tag defines a thematic break in an HTML page (e.g. a shift of topic).
func HorizontalRuleComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("hr", elements...)
}

// <i> tag defines a part of text in an alternate voice or mood. The content inside is typically displayed in italic type.
func ItalicComponent(elements ...any) *hlive.Component { return hlive.NewComponent("i", elements...) }

// <iframe> tag defines an inline frame used to embed another document within the current HTML document.
func InlineFrameComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("iframe", elements...)
}

// <img> tag defines an image in an HTML page.
func ImageComponent(elements ...any) *hlive.Component { return hlive.NewComponent("img", elements...) }

// <input> tag specifies an input field where the user can enter data.
func InputComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("input", elements...)
}

// <ins> (short for insert) tag defines a range of text that has been added to a document.
func InsertedComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("ins", elements...)
}

// <kbd> tag defines keyboard text.
func KeyboardComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("kbd", elements...)
}

// <keygen> tag defines a key-pair generator field (for forms).
func KeygenComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("keygen", elements...)
}

// <label> tag defines a label for an <input> element.
func LabelComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("label", elements...)
}

// <legend> tag defines a caption for the <fieldset> element.
func LegendComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("legend", elements...)
}

// <li> tag defines a list item.
func ListItemComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("li", elements...)
}

// <main> tag specifies the main content of a document.
func MainComponent(elements ...any) *hlive.Component { return hlive.NewComponent("main", elements...) }

// <map> tag is used with <area> elements to define an image map (a clickable link area).
func MapComponent(elements ...any) *hlive.Component { return hlive.NewComponent("map", elements...) }

// <mark> tag defines marked/highlighted text.
func MarkedComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("mark", elements...)
}

// <menu> tag defines a list/menu of commands.
func MenuComponent(elements ...any) *hlive.Component { return hlive.NewComponent("menu", elements...) }

// <meta> tag provides metadata about the HTML document. Metadata will not be displayed on the page, but will be machine parsable.
func MetaComponent(elements ...any) *hlive.Component { return hlive.NewComponent("meta", elements...) }

// <meter> tag defines a scalar measurement within a known range (a gauge).
func MeterComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("meter", elements...)
}

// <nav> tag defines a set of navigation links.
func NavigationComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("nav", elements...)
}

// <noscript> tag defines an alternate content for users that do not support client-side scripts.
func NoScriptComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("noscript", elements...)
}

// <object> tag defines an embedded object within an HTML document.
func ObjectComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("object", elements...)
}

// <ol> tag defines an ordered list.
func OrderedListComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("ol", elements...)
}

// <optgroup> tag defines a group of related options in a drop-down list.
func OptionGroupComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("optgroup", elements...)
}

// <option> tag defines an option in a drop-down list.
func OptionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("option", elements...)
}

// <output> tag represents the result of a calculation.
func OutputComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("output", elements...)
}

// <p> tag defines a paragraph.
func ParagraphComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("p", elements...)
}

// <param> tag defines a parameter for an object.
func ParameterComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("param", elements...)
}

// <picture> tag specifies multiple source for img element
func PictureComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("picture", elements...)
}

// <pre> tag defines preformatted text.
func PreformattedComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("pre", elements...)
}

// <progress> tag represents the progress of a task.
func ProgressComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("progress", elements...)
}

// <q> tag defines a short quotation.
func QuotationComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("q", elements...)
}

// <rp> tag defines what to show in browsers that do not support ruby annotations.
func RubyParenthesisComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("rp", elements...)
}

// <rt> tag defines an explanation/pronunciation of characters (for East Asian typography).
func RubyTextComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("rt", elements...)
}

// <ruby> tag defines a ruby annotation (for East Asian typography).
func RubyComponent(elements ...any) *hlive.Component { return hlive.NewComponent("ruby", elements...) }

// <samp> tag defines sample output from a computer program.
func SampleComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("samp", elements...)
}

// <script> tag is used to define a client-side script (JavaScript).
func ScriptComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("script", elements...)
}

// <section> tag defines sections in a document, such as chapters, headers, footers, or any other sections of the document.
func SectionComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("section", elements...)
}

// <select> tag defines a drop-down list.
func SelectComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("select", elements...)
}

// <small> tag defines smaller text.
func SmallComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("small", elements...)
}

// <source> tag defines multiple media resources for media elements (<video> and <audio>).
func SourceComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("source", elements...)
}

// <span> tag is used to group inline-elements in a document.
func SpanComponent(elements ...any) *hlive.Component { return hlive.NewComponent("span", elements...) }

// <strong> tag defines important text
func StrongComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("strong", elements...)
}

// <style> tag defines style information for a document
func StyleComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("style", elements...)
}

// <sub> tag defines subscripted text
func SubscriptComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("sub", elements...)
}

// <summary> tag defines a visible heading for a <details> element
func SummaryComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("summary", elements...)
}

// <sup> tag defines superscripted text
func SuperscriptComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("sup", elements...)
}

// <table> tag defines an HTML table
func TableComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("table", elements...)
}

// <tbody> tag groups the body content in a table
func TableBodyComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("tbody", elements...)
}

// <td> tag defines a standard cell in an HTML table
func TableDataComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("td", elements...)
}

// <textarea> tag defines a multi-line text input control
func TextAreaComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("textarea", elements...)
}

// <tfoot> tag groups the footer content in a table
func TableFooterComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("tfoot", elements...)
}

// <th> tag defines a header cell in an HTML table
func TableHeadComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("th", elements...)
}

// <thead> tag groups the header content in a table
func TableHeaderComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("thead", elements...)
}

// <time> tag defines a date/time
func TimeComponent(elements ...any) *hlive.Component { return hlive.NewComponent("time", elements...) }

// <title> tag defines a title for the document
func TitleComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("title", elements...)
}

// <tr> tag defines a row in an HTML table
func TableRowComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("tr", elements...)
}

// <track> tag defines text tracks for media elements (<video> and <audio>)
func TrackComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("track", elements...)
}

// <u> tag defines text that should be stylistically different from normal text
func UnderlineComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("u", elements...)
}

// <ul> tag defines an unordered list
func UnorderedListComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("ul", elements...)
}

// <var> tag defines a variable
func VariableComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("var", elements...)
}

// <video> tag defines a video or movie
func VideoComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("video", elements...)
}

// <wbr> tag defines a possible line-break
func WordBreakOpportunityComponent(elements ...any) *hlive.Component {
	return hlive.NewComponent("wbr", elements...)
}
