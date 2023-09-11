package cmd

import (
	"github.com/solodba/mcube/apps"
	"github.com/solodba/mcube/logger"
	"github.com/solodba/mcube/version"
	_ "github.com/solodba/ms_repair/apps/all"
	"github.com/solodba/ms_repair/cmd/start"
	"github.com/solodba/ms_repair/conf"
	"github.com/spf13/cobra"
)

// 全局参数
var (
	showVersion bool
	Username    string
	Password    string
	Host        string
	Port        int32
)

// 根命令
var RootCmd = &cobra.Command{
	Use:     "ms_repair [init|handle]",
	Short:   "ms_repair service",
	Long:    "ms_repair service",
	Example: "ms_repair -v",
	RunE: func(cmd *cobra.Command, args []string) error {
		if showVersion {
			logger.L().Info().Msgf(version.ShortVersion())
			return nil
		}
		return cmd.Help()
	},
}

// 加载全局配置
func LoadConfigFromCmd() {
	conf.Conf = conf.NewDefaultConfig()
	conf.Conf.MySQL.Username = Username
	conf.Conf.MySQL.Password = Password
	conf.Conf.MySQL.Host = Host
	conf.Conf.MySQL.Port = Port
	conf.Conf.MySQL.DB = "mysql"
	conf.Conf.MySQL.MaxOpenConn = 50
	conf.Conf.MySQL.MaxIdleConn = 10
	conf.Conf.MySQL.MaxLifeTime = 600
	conf.Conf.MySQL.MaxIdleTime = 300
	conf.Conf.CmdConf.Username = Username
	conf.Conf.CmdConf.Password = Password
	conf.Conf.CmdConf.Host = Host
	conf.Conf.CmdConf.Port = Port
}

// 初始化函数
func Initial() {
	LoadConfigFromCmd()
	err := apps.InitInternalApps()
	cobra.CheckErr(err)
}

// 执行函数
func Execute() {
	cobra.OnInitialize(Initial)
	RootCmd.AddCommand(start.Cmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

// 初始化函数
func init() {
	RootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show project binlog parse version")
	RootCmd.PersistentFlags().StringVarP(&Username, "username", "u", "test", "connect mysql username")
	RootCmd.PersistentFlags().StringVarP(&Password, "password", "p", "test", "connect mysql password")
	RootCmd.PersistentFlags().StringVarP(&Host, "host", "m", "127.0.0.1", "mysql host ip")
	RootCmd.PersistentFlags().Int32VarP(&Port, "port", "P", 3306, "mysql port")
}
