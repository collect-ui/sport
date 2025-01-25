package model

func addTable() {

	//31 collect_doc
	tableCollectDoc := CollectDoc{}
	modelMap["collect_doc"] = tableCollectDoc
	primaryKeyMap["collect_doc"] = tableCollectDoc.PrimaryKey()
	//32 collect_doc_demo
	tableCollectDocDemo := CollectDocDemo{}
	modelMap["collect_doc_demo"] = tableCollectDocDemo
	primaryKeyMap["collect_doc_demo"] = tableCollectDocDemo.PrimaryKey()
	//33 collect_doc_important
	tableCollectDocImportant := CollectDocImportant{}
	modelMap["collect_doc_important"] = tableCollectDocImportant
	primaryKeyMap["collect_doc_important"] = tableCollectDocImportant.PrimaryKey()
	//34 collect_doc_params
	tableCollectDocParams := CollectDocParams{}
	modelMap["collect_doc_params"] = tableCollectDocParams
	primaryKeyMap["collect_doc_params"] = tableCollectDocParams.PrimaryKey()
	//35 collect_doc_result
	tableCollectDocResult := CollectDocResult{}
	modelMap["collect_doc_result"] = tableCollectDocResult
	primaryKeyMap["collect_doc_result"] = tableCollectDocResult.PrimaryKey()
	//36 config_detail
	tableConfigDetail := ConfigDetail{}
	modelMap["config_detail"] = tableConfigDetail
	primaryKeyMap["config_detail"] = tableConfigDetail.PrimaryKey()
	//37 config_group
	tableConfigGroup := ConfigGroup{}
	modelMap["config_group"] = tableConfigGroup
	primaryKeyMap["config_group"] = tableConfigGroup.PrimaryKey()
	//89 doc_group
	tableDocGroup := DocGroup{}
	modelMap["doc_group"] = tableDocGroup
	primaryKeyMap["doc_group"] = tableDocGroup.PrimaryKey()
	//242 role
	tableRole := Role{}
	modelMap["role"] = tableRole
	primaryKeyMap["role"] = tableRole.PrimaryKey()

	//242 role
	userRoleIdList := UserRoleIDList{}
	modelMap["user_role_id_list"] = userRoleIdList
	primaryKeyMap["user_role_id_list"] = userRoleIdList.PrimaryKey()
	//293 sys_code
	tableSysCode := SysCode{}
	modelMap["sys_code"] = tableSysCode
	primaryKeyMap["sys_code"] = tableSysCode.PrimaryKey()

	//353 user_role_id_list
	userAccount := UserAccount{}
	modelMap["user_account"] = userAccount
	primaryKeyMap["user_account"] = userAccount.PrimaryKey()

	trainingProject := TrainingProject{}
	modelMap["training_project"] = trainingProject
	primaryKeyMap["training_project"] = trainingProject.PrimaryKey()

	course := Course{}
	modelMap["course"] = course
	primaryKeyMap["course"] = course.PrimaryKey()
	courseVideo := CourseVideo{}
	modelMap["course_video"] = courseVideo
	primaryKeyMap["course_video"] = courseVideo.PrimaryKey()
	// 成绩
	courseGrade := CourseGrade{}
	modelMap["course_grade"] = courseGrade
	primaryKeyMap["course_grade"] = courseGrade.PrimaryKey()
	// 签到
	courseAttendance := CourseAttendance{}
	modelMap["course_attendance"] = courseAttendance
	primaryKeyMap["course_attendance"] = courseAttendance.PrimaryKey() // 签到
	attachment := Attachment{}
	modelMap["attachment"] = attachment
	primaryKeyMap["attachment"] = attachment.PrimaryKey()
	priceDocTree := PriceDocTree{}
	modelMap["price_doc_tree"] = priceDocTree
	primaryKeyMap["price_doc_tree"] = priceDocTree.PrimaryKey()
	priceDocProject := PriceDocProject{}
	modelMap["price_doc_project"] = priceDocProject
	primaryKeyMap["price_doc_project"] = priceDocProject.PrimaryKey()
	priceDoc := PriceDoc{}
	modelMap["price_doc"] = priceDoc
	primaryKeyMap["price_doc"] = priceDoc.PrimaryKey()
	// 技术文档
	technicalDoc := TechnicalDoc{}
	modelMap["technical_doc"] = technicalDoc
	primaryKeyMap["technical_doc"] = technicalDoc.PrimaryKey()
	// 菜单
	sysMenu := SysMenu{}
	modelMap["sys_menu"] = sysMenu
	primaryKeyMap["sys_menu"] = sysMenu.PrimaryKey()

	priceDocQingProject := PriceDocQingProject{}
	modelMap["price_doc_qing_project"] = priceDocQingProject
	primaryKeyMap["price_doc_qing_project"] = priceDocQingProject.PrimaryKey()

	priceDocQingTree := PriceDocQingTree{}
	modelMap["price_doc_qing_tree"] = priceDocQingTree
	primaryKeyMap["price_doc_qing_tree"] = priceDocQingTree.PrimaryKey()

	priceDocQing := PriceDocQing{}
	modelMap["price_doc_qing"] = priceDocQing
	primaryKeyMap["price_doc_qing"] = priceDocQing.PrimaryKey()

	roleMenu := RoleMenu{}
	modelMap["role_menu"] = roleMenu
	primaryKeyMap["role_menu"] = roleMenu.PrimaryKey()

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

}
