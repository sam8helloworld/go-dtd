package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// TODO: DTDファイルを読み取る
	f, err := os.Open("examples/example01.dtd")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	// TODO: DTDファイルを字句解析する
	// lexer := lexer.NewLexer(data)
	// tokens, err := lexer.Execute()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// TODO: 字句解析したトークン群を構文解析してDTDの構造体
	// parser := parser.NewParser(tokens)
	// dtd, err := parser.Execute()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// TODO: DTDの構造体からGoのxmlに準拠したUnmarshal用の構造体ファイルを出力する
}
