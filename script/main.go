package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ProjectInfo struct {
	ProjectID          string `json:"project_id"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_desc"`
	License            string `json:"project_lisc"`
}

func (s *ProjectInfo) String() string {
	obj, err := json.MarshalIndent(s, "", "	")
	if err != nil {
		return "err"
	}
	return string(obj)
}

func cleanup(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return text

}

func ReadPrompt(prompt string, reader *bufio.Reader) string {
	fmt.Print(fmt.Sprintf("%%(%s)%% -> ", prompt))
	text, _ := reader.ReadString('\n')
	return cleanup(text)
}

func main() {

	data := ProjectInfo{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to GeekBeacon")
	fmt.Println("---------------------")

	data.ProjectName = ReadPrompt("Project Name", reader)
	data.ProjectID = ReadPrompt("Project ID", reader)
	data.ProjectDescription = ReadPrompt("Description", reader)
	fmt.Println(data.String())

}
