#!/usr/bin/env bash

1. cpu.idle  接口  cpu
2.
3. https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=cpu.idle&endts=0&style=small&startts=-3600
4. Request Method:
GET
  参数
1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
cpu.idle
4. endts:
0
5. style:
small
6. startts:
-3600
返回数据格式
1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
    2. series:{cpu.idle: []}
    3. title:""
2. msg:””


  mem.memused.percent 接口 内存空闲比例
1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=mem.memused.percent&endts=0&style=small&startts=-3600
2. Request Method:
GET
  参数

1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
mem.memused.percent
4. endts:
0
5. style:
small
6. startts:
-3600
返回数据格式
1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
    2. series:{mem.memused.percent: []}
    3. title:""
2. msg:""





load.1min 接口  load.1min
1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=load.1min&endts=0&style=small&startts=-3600
2. Request Method:
GET

1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
load.1min
4. endts:
0
5. style:
small
6. startts:
-3600
1.
data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
    2. series:{load.1min: []}
    3. title:""
2. msg:""


   load.5min 接口  load.5min
1. https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=load.5min&endts=0&style=small&startts=-3600
2. Request Method:
GET
1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
load.5min
4. endts:
0
5. style:
small
6. startts:
-3600
1. 返回数据接口
data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
    2. series:{load.5min: []}
    3. title:""
2. msg:””


 load.15min  接口   load.15min

1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=load.15min&endts=0&style=small&startts=-3600
2. Request Method:
GET
1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
load.15min
4. endts:
0
5. style:
small
6. startts:
-3600

1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
    2. series:{load.15min: []}
    3. title:""
2. msg:""




1. df.statistics.used.percent 接口 机器整体磁盘使用比例
1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=df.statistics.used.percent&endts=0&style=small&startts=-3600
2. Request Method:
GET

1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
df.statistics.used.percent
4. endts:
0
5. style:
small
6. startts:
-3600


1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
        1. cf:"AVERAGE"
        2. end:0
        3. end_ts:1507876149
        4. graph:null
        5. id:0
        6. oid:1
        7. pos:0
        8. space_id:1
        9. span:10
        10. start:0
        11. start_ts:1507872549
        12. unit:"minute"
    2. series:{df.statistics.used.percent: []}
        1. df.statistics.used.percent:[]
    3. title:""
2. msg:”


  df.max.used.percent接口  空间使用率最大的磁盘的使用率

1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=df.max.used.percent&endts=0&style=small&startts=-3600
2. Request Method:
GET
1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
df.max.used.percent
4. endts:
0
5. style:
small
6. startts:
-3600

1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
        1. cf:"AVERAGE"
        2. end:0
        3. end_ts:1507876149
        4. graph:null
        5. id:0
        6. oid:1
        7. pos:0
        8. space_id:1
        9. span:10
        10. start:0
        11. start_ts:1507872549
        12. unit:"minute"
    2. series:{df.max.used.percent: []}
        1. df.max.used.percent:[]
    3. title:""
2. msg:""


disk.io.max.util 接口  io使用率最大的盘的io util

1. Request URL:
https://192.168.113.107:3000/apm/chart/1/counter/json?nodeid=0&endpoints=054776d53554dff73b80d91f2c6f55b5&counter=disk.io.max.util&endts=0&style=small&startts=-3600
2. Request Method:
GET

1. nodeid:
0
2. endpoints:
054776d53554dff73b80d91f2c6f55b5
3. counter:
disk.io.max.util
4. endts:
0
5. style:
small
6. startts:
-3600

1. data:{meta: {id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…},…}
    1. meta:{id: 0, oid: 1, space_id: 1, start: 0, end: 0, cf: "AVERAGE", pos: 0, start_ts: 1507872549,…}
        1. cf:"AVERAGE"
        2. end:0
        3. end_ts:1507876149
        4. graph:null
        5. id:0
        6. oid:1
        7. pos:0
        8. space_id:1
        9. span:10
        10. start:0
        11. start_ts:1507872549
        12. unit:"minute"
    2. series:{disk.io.max.util: []}
        1. disk.io.max.util:[]
    3. title:""
2. msg:""


//创建任务
 curl -XPOST http://^Cw.opdeck.com/probe/task/1/http -d '{"Url":"http://www.baidu.com", "Method":1,"PeriodSec":120}' --cookieie rywww=MTUwOTA5MDA5MHxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx4YcSYcXa6iuZ1ozI5p_N4Zf86WF_Il_KP54WNE7VWpbA==ZiMnRwWlVSaGRHRUJfNElBQVFJQk
{

	PeriodSec  int
	Url       string
	Method    int  1:GET 2:POST 3:HEAD
	Header    {}
	Cookies   string
	BasicAuth {
	    User   string
	    Passwd string
	}
	ServerIp  string
	Matcher  {
	    Target   int   1:BODY 2:HEAD
	    Method   int   1:INCLUD 2:EXCLUDE
	    Content  string
	    StatusCode int
     }
}



#根据task id(2) 查询 orgId = 1 下面的http task详细信息
curl  http://www.opdeck.com/probe/task/org/1/http/task/2  --cookie rywww=MTUwOTA5MDA5MHxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx4cSYcXa6iuZ1ozI5p_N4Zf86WF_Il_KP54WNE7VWpbA==

#查询orgId = 1  下面的http task详细信息
curl  http://www.opdeck.com/probe/task/org/1/http  --cookie rywww=MTUwOTA5MDA5MHxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx4cSYcXa6iuZ1ozI5p_N4Zf86WF_Il_KP54WNE7VWpbA==

#将orgID=1下的task id = 2 的task 绑定 rule=7的报警规则
curl -XPOST http://www.opdeck.com/probe/task/org/1/http/task/2/bind/rule/7   --cookie rywww=MTUwOTA5MDA5MHxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx4cSYcXa6iuZ1ozI5p_N4Zf86WF_Il_KP54WNE7VWpbA==

##将orgID=1下的task id = 2 的task 解绑定 rule=7的报警规则
curl -XPOST http://www.opdeck.com/probe/task/org/1/http/task/2/unbind/rule/7   --cookie rywww=MTUwOTA5MDA5MHxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx4cSYcXa6iuZ1ozI5p_N4Zf86WF_Il_KP54WNE7VWpbA==

curl -XPUT http://www.opdeck.com/probe/worker/2 -d '{"Id":0,"Password":"","Status":"","StartTimestamp":0,"UpdateTimestamp":0,"Country":"中国","Province":"","City":"深圳","Operator":"联通","Label":{"sys":null,"user":null,"other":{"Location":[114.07,22.62]}}}' --cookie rywww=MTUxMDAyMzk5NXxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXx-VgnClo6DFJwNiczfRqfXRu3P4Da8znrvOfFWwugJ6g==


#	Matcher   *HttpSpecMatcher   `protobuf:"bytes,9,opt,name=Matcher,json=matcher" json:"Matcher" xorm:"json"`
#	WebImage  string             `protobuf:"bytes,10,opt,name=WebImage,json=webImage" json:"WebImage"`




curl -XPOST http://www.opdeck.com/probe/task/org/1/http -d '{
    "TaskObj": {
        "NodeId": 88,  //可选
        "Type": 1,    //必填1
        "Name": "test007",
        "PeriodSec": 300, //必填监控的时间间隔，单位s
        "Url": "http://www.baidu.com",
        "Method": 1, //http 请求方法， 1：GET 2：POST 3：HEAD
        "Stop" : false, //是否暂停， 默认不填是开启
        "Header": http header， 这是一个object(key-vaule)
        "Cookies": "sfdasdfasdf",
        "BasicAuth": {
            "User": "username",
            "Passwd": "password"
        }
        "Body" "Post body",
        "Matcher": {}
    },
    "Rules": [
        {
            "MaxStep": 1,
            "Metric": "url.http.delay",
            "Op": ">",
            "RightValue": 100000,
            RunBegin:"",
	        RunEnd: ""
        }
    ],
    "TeamIds": [
        10
    ]
}' --cookie rywww=MTUxMDkzMTY4OXxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXyqY2atn57lPd2U3aXd_CdH_q04DET0JYHnkQw-zeJGMA==

curl -XPOST http://www.opdeck.com/probe/task/org/1/http -d '{
    "TaskObj": {
        "Type": 1,
        "Name": "test007",
        "PeriodSec": 300,
        "Url": "http://www.baidu.com",
        "Method": 1,
        "Rules":[100]
    },
    "Rules": [
        {
            "MaxStep": 1,
            "Metric": "url.http.delay",
            "Op": ">",
            "RightValue": 100000,
            "RunBegin": "",
            "RunEnd": ""
        }
    ],
    "TeamIds": [
        10
    ]
}' --cookie rywww=MTUxMDkzMTY4OXxNUC1CQXdFQkNrTnZiMnRwWlVSaGRHRUJfNElBQVFJQkJsVnpaWEpKWkFFRUFBRUlWWE5sY201aGJXVUJEQUFBQUJIX2dnRUNBUW94UUhSbGMzUXVZMjl0QUE9PXyqY2atn57lPd2U3aXd_CdH_q04DET0JYHnkQw-zeJGMA==
