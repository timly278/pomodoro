## Must be implemented:
- implement refresh token and verified email
- How to achieve https or SSL 

## Functioning at db layer:
- how many pomodoro per task
- list tasks by status: done, in-progress, ...


## APIs

- /api/users/
    - PUT("/api/users)
        update user's setting

- /api/pomodoros
    - POST("/api/pomodoros")
        create new Pomodoro
            
-  /api/tasks/
    - POST("/api/tasks")
        createNewTask

    - GET("/api/tasks")
        list all done tasks or in-progess
        
    - PUT("/api/tasks/:id")
        updateTask
        
    - DELETE("/api/tasks/:id")
        remove a task
            
-  /api/report
    - GET("/api/report")
        general statistic numbers: hours focused, days accessed, ...

    - GET("/api/report/week/:id")
            list pomo by week counted from here and nowv
        
    - GET("/api/report/month/:id")
        list pomo by month
            
- /api/types/   - newType, updateType, deleteType
        GET  - list all available types
        POST - create a new type
        PUT  - update type
            /api/types/:id
        
- /api/goals/


## Not yet and need to be tested:
- UpdateUserSetting
- CreateNewPomoType
- ListPomoType
- UpdateType