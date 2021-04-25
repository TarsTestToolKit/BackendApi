#!/bin/bash
trap 'exit' SIGTERM SIGINT

#公共函数
function LOG_ERROR()
{
	if (( $# < 1 ))
	then
		echo -e "\033[33m usesage: LOG_ERROR msg \033[0m";
	fi

	local msg=$(date +%Y-%m-%d" "%H:%M:%S);

    msg="${msg} $@";

	echo -e "\033[31m $msg \033[0m";
}

function LOG_WARNING()
{
	if (( $# < 1 ))
	then
		echo -e "\033[33m usesage: LOG_WARNING msg \033[0m";
	fi

	local msg=$(date +%Y-%m-%d" "%H:%M:%S);

    msg="${msg} $@";

	echo -e "\033[33m $msg \033[0m";
}

function LOG_DEBUG()
{
	if (( $# < 1 ))
	then
		LOG_WARNING "Usage: LOG_DEBUG logmsg";
	fi

	local msg=$(date +%Y-%m-%d" "%H:%M:%S);

    msg="${msg} $@";

 	echo -e "\033[40;37m $msg \033[0m";
}

function LOG_INFO()
{
	if (( $# < 1 ))
	then
		LOG_WARNING "Usage: LOG_INFO logmsg";
	fi

	local msg=$(date +%Y-%m-%d" "%H:%M:%S);

	for p in $@
	do
		msg=${msg}" "${p};
	done

	echo -e "\033[32m $msg \033[0m"
}

#exec sql
function exec_mysql_sql()
{
    LOG_DEBUG "mysql -h${MYSQL_HOST} -u${MYSQL_USER} -p${MYSQL_PASS} -P${MYSQL_PORT} --default-character-set=utf8 -D$1 < $2"
    mysql -h${MYSQL_HOST} -u${MYSQL_USER} -p${MYSQL_PASS} -P${MYSQL_PORT} --default-character-set=utf8 -D$1 < $2

    ret=$?
    LOG_DEBUG "mysql return:${ret}"

    return $ret
}
MYSQL_PASS=$1
exec_mysql_sql db_tars sql/services.sql
sleep 5s

## TestUnits.golang
LOG_INFO "upload_and_publish TestUnits.golang"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/golang.tgz" \
  --form "application=TestUnits" \
  --form "module_name=golang" \
  --form "comment=developer-auto-upload"

## TestUnits.php
LOG_INFO "upload_and_publish TestUnits.php"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/php.tgz" \
  --form "application=TestUnits" \
  --form "module_name=php" \
  --form "comment=developer-auto-upload"

## TestUnits.java
LOG_INFO "upload_and_publish TestUnits.java"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/java.tgz" \
  --form "application=TestUnits" \
  --form "module_name=java" \
  --form "comment=developer-auto-upload"

## TestUnits.cpp
LOG_INFO "upload_and_publish TestUnits.cpp"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/cpp.tgz" \
  --form "application=TestUnits" \
  --form "module_name=cpp" \
  --form "comment=developer-auto-upload"

## TestUnits.nodejs
LOG_INFO "upload_and_publish TestUnits.nodejs"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/nodejs.tgz" \
  --form "application=TestUnits" \
  --form "module_name=nodejs" \
  --form "comment=developer-auto-upload"

## TarsTestToolKit.ResFetcher
LOG_INFO "upload_and_publish TarsTestToolKit.ResFetcher"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/ResFetcher.tgz" \
  --form "application=TarsTestToolKit" \
  --form "module_name=ResFetcher" \
  --form "comment=developer-auto-upload"

## TarsTestToolKit.BackendApi
LOG_INFO "upload_and_publish TarsTestToolKit.BackendApi"
curl -i --request POST "http://127.0.0.1:3000/api/upload_and_publish" \
  --form "suse=@./patches/BackendApi.tgz" \
  --form "application=TarsTestToolKit" \
  --form "module_name=BackendApi" \
  --form "comment=developer-auto-upload"
