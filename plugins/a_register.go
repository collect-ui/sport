package plugins

import (
	templateService "github.com/collect-ui/collect/src/collect/service_imp"
)

func GetRegisterList() []templateService.ModuleResult {
	l := make([]templateService.ModuleResult, 0)
	l = append(l, &AnalysisIp{})
	l = append(l, &Shell{})
	l = append(l, &ShellTerm{})
	l = append(l, &ParamKey2Arr{})
	l = append(l, &RenameField{})
	l = append(l, &MultiArr{})           // 数组成数组
	l = append(l, &HandlerPassword{})    // 数组成数组
	l = append(l, &ValueTransfer{})      // 值转换
	l = append(l, &ReadFile{})           // 值转换
	l = append(l, &Sftp{})               // 值转换
	l = append(l, &AnalysisAttendance{}) // 值转换
	l = append(l, &ToLocalFile{})
	l = append(l, &GenSportLevel{})
	l = append(l, &Xml2Json{})
	// 执行shell
	l = append(l, &Ssh{})
	return l
}
