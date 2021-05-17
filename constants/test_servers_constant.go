package constants

import "time"

const (
	AppNameTestUnits = "TestUnits"

	ServiceFullNameBMAdmin = "benchmark.AdminServer.AdminObj"
	ServiceFullNameCpp     = "TestUnits.CppTars.testObj"
	ServiceFullNameGolang  = "TestUnits.GolangTars.testObj"
	ServiceFullNameJava    = "TestUnits.JavaTars.testObj"
	ServiceFullNameNodeJs  = "TestUnits.NodejsTars.testObj"
	ServiceFullNamePhp     = "TestUnits.PhpTars.testObj"

	ServeNameCpp    = "CppTars"
	ServeNameGolang = "GolangTars"
	ServeNameJava   = "JavaTars"
	ServeNameNodeJs = "NodejsTars"
	ServeNamePhp    = "PhpTars"

	LangCpp    = "cpp"
	LangGolang = "golang"
	LangJava   = "java"
	LangNodejs = "nodejs"
	LangPHP    = "php"

	CfgReloadDur = 30 * time.Second

	ServiceNameResFetcher = "TarsTestToolKit.ResFetcher.fetcherObj"

	LocatorKeyLocal            = "local"
	ServiceNameServiceRegistry = "tars.tarsregistry.QueryObj"

	TaskStatusRunning = 1
	TaskStatusSucc    = 2
	TaskStatusFailed  = 3
)

var LangMap = map[string]string{
	LangCpp:    ServiceFullNameCpp,
	LangGolang: ServiceFullNameGolang,
	LangJava:   ServiceFullNameJava,
	LangNodejs: ServiceFullNameNodeJs,
	LangPHP:    ServiceFullNamePhp,
}

var ServNameMap = map[string]string{
	LangCpp:    ServeNameCpp,
	LangGolang: ServeNameGolang,
	LangJava:   ServeNameJava,
	LangNodejs: ServeNameNodeJs,
	LangPHP:    ServeNamePhp,
}
