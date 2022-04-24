
type Program struct { 
	StartAddress uint32 // entry point of program, specified by __start__ lable in text section
	DataStart uint32 // start of data segment
	Size uint32 // size of entire program 
	Data []byte //data segment
	SymbolTable map[string]uint32 //table to map labels to addresses
}

type Parser struct {
	lex *Lexer //Lexer struct reference
}


func (parser *Parser) Parse() *Program { 
	//returns a pointer to a newly assembled program 
}	