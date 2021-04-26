package constants

import "time"

const (
	AppNameTestUnits = "TestUnits"

	ServiceNameBMAdmin = "benchmark.AdminServer.AdminObj"
	ServiceNameCpp     = "TestUnits.CppTars.testObj"
	ServiceNameGolang  = "TestUnits.GolangTars.testObj"
	ServiceNameJava    = "TestUnits.JavaTarstestObj"
	ServiceNameNodeJs  = "TestUnits.NodejsTars.testObj"
	ServiceNamePhp     = "TestUnits.PhpTars.testObj"

	LangCpp    = "cpp"
	LangGolang = "golang"
	LangJava   = "java"
	LangNodejs = "nodejs"
	LangPHP    = "php"

	CfgReloadDur = 5 * time.Second

	ServiceNameResFetcher = "TarsTestToolKit.ResFetcher.fetcherObj"

	LocatorKeyLocal            = "local"
	ServiceNameServiceRegistry = "tars.tarsregistry.QueryObj"

	TaskStatusRunning = 1
	TaskStatusSucc    = 2
	TaskStatusFailed  = 3
)

var LangMap = map[string]string{
	LangCpp:    ServiceNameCpp,
	LangGolang: ServiceNameGolang,
	LangJava:   ServiceNameJava,
	LangNodejs: ServiceNameNodeJs,
	LangPHP:    ServiceNamePhp,
}
