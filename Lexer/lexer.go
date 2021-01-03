package Lexer

type Lexer struct {
	input        string // Texto a ser recebido e interpretado
	position     int    // Posição atual
	readPosition int    // Posição atual + 1
	ch           byte   // Caractere atual
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // Começa um novo lexer com o "texto" passado

	return l
}

func (l *Lexer) readChar() {
	// Verificamos se a proxima posição é um EOF, se for indicamos isso com um 0
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Se não, lemos o caractere que a posição representa
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition // Sempre guardamos a ultima posição no position

	l.readPosition += 1
}
