package repl

import (
	"bufio"
	"fmt"
	"io"

	// "github.com/pavank1301/Prism/token"
	"github.com/pavank1301/Prism/evaluator"
	"github.com/pavank1301/Prism/lexer"
	"github.com/pavank1301/Prism/object"
	"github.com/pavank1301/Prism/parser"
)

	const PROMPT = ">> "
	const PRISM = `                  (
 AAAAAaaaaaaaa!  Error!!!
              (   ()   )
    ) ________    //  )
 ()  |\       \  //
( \\__ \ ______\//
   \__) |       |
     |  |       |
      \ |       |
       \|_______|
       //    \\
      ((     ||
       \\    ||
     ( ()    ||
      (      () ) )`

func Start (in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for{
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned{
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0{
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated!=nil{
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string){
	io.WriteString(out, PRISM)
	io.WriteString(out, "\noops!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors{
		io.WriteString(out, "\t"+msg+"\t")
	}
}
	