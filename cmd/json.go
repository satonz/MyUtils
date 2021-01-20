package cmd

import (
	"MyUtils/internal/json2struct"
	"log"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json转换和处理",
	Long:  "json转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var json2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "json转换",
	Long:  "json转换",
	Run: func(cmd *cobra.Command, args []string) {
		parser, err := json2struct.NewParser(str)
		if err != nil {
			log.Fatalf("json2struct.NewParser err: %v", err)
		}
		content := parser.Json2Struct()
		log.Printf("输出结果: %s", content)
	},
}

func init() {
	jsonCmd.AddCommand(json2structCmd)
	// 输入结构应为{\"name\":1}  需要转义符号
	json2structCmd.Flags().StringVarP(&str, "str", "s", "", "请输入json字符串")
}
