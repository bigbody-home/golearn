package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	Host string

	Port int

	Username string

	Secret string

	Table string

	CsvPath string

	Db string

	rootCmd = &cobra.Command{
		Use:        "light",
		SuggestFor: []string{"-l$HOSTNAME -P$PORT -u$YOURUSER -s$YOURPWD -d$YOURDB -t$YOURTABLE -c$YOURCSVPATH"},
		Short:      "一个并发解析文件内容到数据库的工具",
		Long:       "一个并发解析文件内容到数据库的工具。只需要指定要入库的表和csv文件路径，表字段和单个文件的csv文件要一一对应",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("主机名：%v,用户：%v，数据库：%v，表：%v，文件目录：%v", Host, Username, Db, Table, CsvPath)
			ParseCsvToDb(CsvPath, Table)
		},
	}
)

func Execute() {
	rootCmd.Execute()
}
func init() {
	rootCmd.Flags().StringVarP(&Host, "host", "l", "localhost", "指定数据库的链接地址")
	rootCmd.Flags().StringVarP(&Username, "username", "u", "light_demo", "指定数据库的用户名")
	rootCmd.Flags().StringVarP(&Secret, "secret", "s", "light_demo", "指定数据库的账户密码")
	rootCmd.Flags().StringVarP(&CsvPath, "csvpath", "c", "/test", "存放csv文件的目录")
	rootCmd.Flags().StringVarP(&Db, "database", "d", "test", "指定db的名字")
	rootCmd.Flags().StringVarP(&Table, "table", "t", "test", "指定表的名字")
	rootCmd.Flags().IntVarP(&Port, "port", "P", 3306, "指定数据库端口")
}
