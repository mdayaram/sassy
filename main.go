package main

import (
	"fmt"
	"gosass"
	"io/ioutil"
	"os"
)

func CompilerError(errors string) {
	println("Errors while compiling sass...")
	println(errors)
	println("\nExiting...")
}

func main() {
	opts := gosass.Options{
		OutputStyle:  gosass.NESTED_STYLE,
		IncludePaths: make([]string, 0),
	}

	if len(os.Args) > 1 {
		ctx := gosass.FileContext{
			Options:      opts,
			InputPath:    os.Args[1],
			OutputString: "",
			ErrorStatus:  0,
			ErrorMessage: "",
		}

		gosass.CompileFile(&ctx)
		if ctx.ErrorStatus != 0 {
			CompilerError(ctx.ErrorMessage)
			os.Exit(1)
		}

		fmt.Print(ctx.OutputString)
	} else {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			CompilerError(err.Error())
			os.Exit(1)
		}

		ctx := gosass.Context{
			Options:      opts,
			SourceString: string(input),
			OutputString: "",
			ErrorStatus:  0,
			ErrorMessage: "",
		}

		gosass.Compile(&ctx)
		if ctx.ErrorStatus != 0 {
			CompilerError(ctx.ErrorMessage)
			os.Exit(1)
		}

		fmt.Print(ctx.OutputString)
	}
}
