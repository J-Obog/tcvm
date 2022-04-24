package assembler

type Instruction struct {
	Opcode    uint8
	Type      uint8
	Operands  []uint32
	SymbolRef string
}

type SymbolTable map[string]struct {
	Section uint8
	Offset  uint32
}

type Program struct {
	StartAddress uint32        // entry point of program, specified by __start__ lable in text section
	DataStart    uint32        // start of data segment
	Size         uint32        // size of entire program
	Data         []byte        //data segment
	Tbl          SymbolTable   //table to map labels to addresses
	Code         []Instruction //list of instructions
}

type Parser struct {
	lex *Lexer //Lexer struct reference
}

func (parser *Parser) Parse() *Program {
	//returns a pointer to a newly assembled program
	return nil
}