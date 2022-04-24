package assembler

type TextElement struct {
	Type  uint8
	Size  uint8
	Value []byte
}

const ( // node type mapping
	instruction = iota //
	register
	effective
	label
	immediate
)

type Program struct {
	ExternalLabels []uint32          //useful for linking programs
	DataSegment    []byte            //.data
	SymbolTable    map[string]uint32 //labels within program
	TextSegment    []byte            //.text
}

type Parser struct {
	lex *Lexer //Lexer struct reference
}

func (parser *Parser) Parse() {

}