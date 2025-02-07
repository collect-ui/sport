package devops

func GetTable() (map[string]interface{}, map[string][]string) {
	modelMap := make(map[string]interface{})
	primaryKeyMap := make(map[string][]string)

	webshellLog := WebshellLog{}
	modelMap["webshell_log"] = webshellLog
	primaryKeyMap["webshell_log"] = webshellLog.PrimaryKey()

	webshellToken := WebshellToken{}
	modelMap["webshell_token"] = webshellToken
	primaryKeyMap["webshell_token"] = webshellToken.PrimaryKey()

	tableServerEnv := ServerEnv{}
	modelMap["server_env"] = tableServerEnv
	primaryKeyMap["server_env"] = tableServerEnv.PrimaryKey()
	//
	tableServerInstallSoft := ServerInstallSoft{}
	modelMap["server_install_soft"] = tableServerInstallSoft
	primaryKeyMap["server_install_soft"] = tableServerInstallSoft.PrimaryKey()

	serverInstance := ServerInstance{}
	modelMap["server_instance"] = serverInstance
	primaryKeyMap["server_instance"] = serverInstance.PrimaryKey()

	serverOsUsers := ServerOsUsers{}
	modelMap["server_os_users"] = serverOsUsers
	primaryKeyMap["server_os_users"] = serverOsUsers.PrimaryKey()

	return modelMap, primaryKeyMap
}
