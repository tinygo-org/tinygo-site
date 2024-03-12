# Generate documentation for the machine package

To re-generate the complete documentation, tinygo needs to be [installed](https://tinygo.org/getting-started/install/), than simply run make:

`make`

For more specific document updates, run:

`make machine`, `make interfaces`, `make pins`, `make matrix`

or for help:

`make help`

## Development

### Markdown Parser

Existing documents can be parsed "manually" by finding the section with "strings.Index", finding the block to replace
by text search and replacing the block with a new content by string concatenation. A more powerful way is to use AST
based parsing and rendering of the markdown document.

Some existing parsers, see also [commonmark implementations](https://github.com/commonmark/commonmark-spec/wiki/List-of-CommonMark-Implementations):

* <https://github.com/russross/blackfriday>, since v2 there is a separate Parse() method, but maybe dead
* <https://gitlab.com/golang-commonmark/markdown>, seems small and basic, separate parse and render methods
* <https://github.com/gomarkdown/markdown>, one successor of "blackfriday", separate parser, nice [description](https://blog.kowalczyk.info/article/cxn3/advanced-markdown-processing-in-go.html) good examples to modify AST
* <https://github.com/yuin/goldmark>, focused on HTML render, most powerful, could not get it to work yet for modifying AST content

Some existing renderer:

* <https://github.com/go-spectest/markdown>, fluent builder
* <https://github.com/gomarkdown/markdown/md>, for gomarkdown AST, <https://github.com/gomarkdown/markdown/issues/285>
* <https://github.com/Kunde21/markdownfmt/v3/markdown>, for goldmark AST, with table support
* <https://github.com/teekennedy/goldmark-markdown>, for goldmark AST, contains nice examples for transform

Some tools to manipulate known AST formats:

* non found yet written in go, so <https://github.com/go-spectest/markdown> fluent builder can be a starting point

We use [goldmark-markdown](https://pkg.go.dev/github.com/teekennedy/goldmark-markdown) for creation of hardware-matrix page. Unfortunately "KindTable" is currently not supported by the [renderer](https://github.com/teekennedy/goldmark-markdown/issues/19), so the site will be created by "text/template".

### Debugging

Each target can be processed separately for the update of the specific document part, e.g.:

`go run . --pins arduino-mega1280`
