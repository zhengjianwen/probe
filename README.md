# probe
distributed probe framework
The system is consist of master, scheduler, worker components

### Master
    
    1. Accept user request and store to DB
    2. Report user statics
    3. Report admin system status
    
### Scheduler 

    1. Do workers health check(this need to be move to single service)
    2. Schedule tasks to working worker
 
### Worker

    1. Accept Scheduler's tasks
    2. Respote task status and result
    3. Tasks are divided into long-term tasks and short-term tasks
         
    
### Use 

    1. probe worker start --master 127.0.0.1:9100    
    
    