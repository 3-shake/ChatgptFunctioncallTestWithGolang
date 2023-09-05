package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//json to go

type Arguments struct {
	DocumentURL []string `json:"document_URL"`
	Tag         []string `json:"tag"`
}

func main() {
	functionCalls := make(map[string]interface{})
	message := os.Args[1]

	i := 0
	for {

		fmt.Printf("---- %d ----\n", i)

		fmt.Printf("Input -> %s\n", message)

		BuiltMessage := BuildMessages(message)

		resp, err := ChatGPT(BuiltMessage)
		if err != nil {
			fmt.Printf("Error executing chat GPT: %v\n", err)
			return
		}

		// FunctionCallが生成されてなければここでループを抜ける
		if resp.Choices[0].Message.FunctionCall == nil {
			fmt.Printf("Generated message -> %s\n", resp.Choices[0].Message)
			break
		}

		fmt.Printf("Generated function -> %s\n", resp.Choices[0].Message.FunctionCall.Name)
		// fmt.Println(resp.Choices[0].Message.FunctionCall.Arguments)

		functionName := resp.Choices[0].Message.FunctionCall.Name
		arguments := resp.Choices[0].Message.FunctionCall.Arguments

		// 新しいFunction Callの関数名を辞書に追加
		functionCalls[functionName] = arguments

		// 辞書全体の表示
		functionMapBytes, err := json.MarshalIndent(functionCalls, "", "  ")
		if err != nil {
			fmt.Printf("Error marshalling functionMap: %v\n", err)
			return
		}

		// 取り残したFunction Callを生成させるためのメッセージ
		message = fmt.Sprintf(`
%s
The above is a generated Function Call.
It is generated based on the following text.
---
%s
---
Are there any other Function Calls?
`, string(functionMapBytes), os.Args[1])

		i++

	}

	fmt.Println("---- Done ----")

	// for ループを抜けた後、functionCalls の内容を表示
	functionMapBytes, err := json.MarshalIndent(functionCalls, "", "  ")
	fmt.Println("results")
	if err != nil {
		fmt.Printf("Error marshalling functionCalls: %v\n", err)
		return
	}
	fmt.Println(string(functionMapBytes))

	// fmt.Println("resp.Choices[0].Message ->", resp.Choices[0].Message)

	// // FunctionCallがあるか調べる。
	// if resp.Choices[0].Message.FunctionCall != nil {
	// 	// fmt.Println(resp.Choices[0].Message)
	// 	fmt.Println(resp.Choices[0].Message.FunctionCall.Name)
	// 	fmt.Println(resp.Choices[0].Message.FunctionCall.Arguments)

	// 	args := &Arguments{}
	// 	err := json.Unmarshal([]byte(resp.Choices[0].Message.FunctionCall.Arguments), args)
	// 	if err != nil {
	// 		fmt.Printf("Error decoding JSON: %v\n", err)
	// 		return
	// 	}

	// 	if resp.Choices[0].Message.FunctionCall.Name == "get_official_documents" {
	// 		for i, url := range args.DocumentURL {
	// 			fmt.Printf("URL %d: %s\n", i, url)

	// 			fmt.Printf("%s\n", GetRequest(url))
	// 		}
	// 	}
	// }
	// GetRequest("https://www.yahoo.co.jp/")
}
