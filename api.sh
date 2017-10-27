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


