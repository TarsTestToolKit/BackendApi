# noinspection SqlNoDataSourceInspectionForFile

delete
from `t_adapter_conf`
where `application` = 'TestUnits'
   or `application` = 'TarsTestToolKit';
delete
from `t_server_conf`
where `application` = 'TestUnits'
   or `application` = 'TarsTestToolKit';

### TestUnits.cpp
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TestUnits', 'cpp', 'localip.tars.com', 'TestUnits.cpp.testObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22001', 2000, '', 'TestUnits.cpp.testObj', 10000, 60000, now(), 'admin',
        'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TestUnits', 'cpp', '', 'localip.tars.com', now(), '', '', 'tars.cpp.default', 0, 'active', 'active', 0,
        '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_cpp', '');

### TestUnits.java
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TestUnits', 'java', 'localip.tars.com', 'TestUnits.java.testObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22002', 2000, '', 'TestUnits.java.testObj', 10000, 60000, now(), 'admin',
        'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TestUnits', 'java', '', 'localip.tars.com', now(), '', '', 'tars.springboot', 0, 'active', 'active', 0,
        '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_java', '');


### TestUnits.nodejs
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TestUnits', 'nodejs', 'localip.tars.com', 'TestUnits.nodejs.testObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22003', 2000, '', 'TestUnits.nodejs.testObj', 10000, 60000, now(), 'admin',
        'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TestUnits', 'nodejs', '', 'localip.tars.com', now(), '', '', 'tars.nodejs.default', 0, 'active', 'active', 0,
        '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_nodejs', '');


### TestUnits.golang
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TestUnits', 'golang', 'localip.tars.com', 'TestUnits.golang.testObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22004', 2000, '', 'TestUnits.golang.testObj', 10000, 60000, now(), 'admin',
        'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TestUnits', 'golang', '', 'localip.tars.com', now(), '', '', 'tars.default', 0, 'active', 'active', 0, '2.1.0',
        now(), '', '2.1.0', now(), 'admin', 'tars_go', '');


### TestUnits.php
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TestUnits', 'php', 'localip.tars.com', 'TestUnits.php.testObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22005', 2000, '', 'TestUnits.php.testObj', 10000, 60000, now(), 'admin',
        'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TestUnits', 'php', '', 'localip.tars.com', now(), '', '', 'tars.tarsphp.default', 0, 'active', 'active', 0,
        '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_php', '');

###ResFetcher
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TarsTestToolKit', 'ResFetcher', 'localip.tars.com', 'TarsTestToolKit.ResFetcher.fetcherObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 22006', 2000, '', 'TarsTestToolKit.ResFetcher.fetcherObj', 10000, 60000,
        now(), 'admin', 'tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TarsTestToolKit', 'ResFetcher', '', 'localip.tars.com', now(), '', '', 'tars.default', 0, 'active', 'active',
        0, '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_go', '');

###BackendApi
replace into `t_adapter_conf` (`application`, `server_name`, `node_name`, `adapter_name`, `registry_timestamp`,
                               `thread_num`, `endpoint`, `max_connections`, `allow_ip`, `servant`, `queuecap`,
                               `queuetimeout`, `posttime`, `lastuser`, `protocol`, `handlegroup`)
VALUES ('TarsTestToolKit', 'BackendApi', 'localip.tars.com', 'TarsTestToolKit.BackendApi.apiObjAdapter', now(), 5,
        'tcp -h localip.tars.com -t 60000 -p 9001', 2000, '', 'TarsTestToolKit.BackendApi.apiObj', 10000, 60000, now(),
        'admin', 'not_tars', '');
replace into `t_server_conf` (`application`, `server_name`, `node_group`, `node_name`, `registry_timestamp`,
                              `base_path`, `exe_path`, `template_name`, `bak_flag`, `setting_state`, `present_state`,
                              `process_id`, `patch_version`, `patch_time`, `patch_user`, `tars_version`, `posttime`,
                              `lastuser`, `server_type`, `profile`)
VALUES ('TarsTestToolKit', 'BackendApi', '', 'localip.tars.com', now(), '', '', 'tars.default', 0, 'active', 'active',
        0, '2.1.0', now(), '', '2.1.0', now(), 'admin', 'tars_go', '');