package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/quillee/monkey/src/lexer"
	"github.com/quillee/monkey/src/token"
)

const PROMPT = ">> "

func Start (in io.Reader, out io.Writer) {
    scanner := bufio.NewScanner(in)
    if scanner == nil {
        fmt.Fprintf(out, "Error loading REPL")
        return
    }

    for {
        fmt.Fprint(out, PROMPT)
        scanned := scanner.Scan()
        if !scanned {
            return
        }

        l := lexer.New(scanner.Text())
        for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
            fmt.Fprintf(out, "%+v\n", tok)
        }
    }
}

