## Must be implemented:
- implement refresh token and verified email
- How to achieve https or SSL 

## Functioning at db layer:
- how many pomodoro per task
- list tasks by status: done, in-progress, ...

![alt text](./pomodoro.jpg)

## APIs

<h3>/api/users/</h3>

- **PUT("/api/users")**:
    Update user's setting with the json parameter as below:
    | Name   | Type | Description |
    | - | :- | :-- |
    | username| string | `required`   |
    | alarm_sound  | string  | `required`|
    | repeat_alarm   | int32   |the number of sounds that the alarm will ring `required`|

<h3> /api/pomodoros </h3>

- **POST("/api/pomodoros")**:
    Create new Pomodoro with json parameters:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |type_id|int64|id of pomodoro type, `required`|
    |task_id"|int64|id of the relevant task|
    |focused_degree|int32|min = 1, max = 5, `required`|

<h3> /api/tasks/ </h3>

- POST("/api/tasks"):
    Create new task

- GET("/api/tasks"):
    List all done tasks or in-progess
   
- PUT("/api/tasks/:id"):
    updateTask

- **DELETE**("/api/tasks/:id"):
    remove a task
            
<h3> /api/report </h3>

- **GET("/api/report")**:
    General statistic numbers: hours focused, days accessed, ...
    <br>Query parameters:

    | Name   | Type | Description |
    | - | :-: | :-- |
    |year|int32|`required`|
    |month|int32|`required`|
    
    <br>Response fields:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |days_accessed|int64|Number of accessed days in the specified month `required`|
    |minutes_focused|int64|Number of focused minutes in the specified month `required`|

- **GET("/api/report/month/")**:
    List pomodoros in the specified month
    <br>Query parameters:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |year|int32|`required`|
    |month|int32|`required`|

    <br>Response fields is a 2D array described as [day_i][pomodoro_j]:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |type_id|int64|Id of pomodoro type `required`|
    |focused_degree|int32|range[1,5] `required`|

- **GET("/api/report/date")**: 
    List pomo by date: one specific day
    <br>Query parameters:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |date_time|rfc3339|the day need to aquire pomodoros `required`|
    |page_size|int32|number of pomodoro want to get`required`|
    |page_id|int32|page number to get pomodoro `required`|

    <br>Response fields is an array:
    | Name   | Type | Description |
    | - | :-: | :-- |
    |type_id|int64|Id of pomodoro type `required`|
    |focused_degree|int32|range[1,5] `required`|

<h3> /api/types/ </h3>

- **GET("/api/types"):**
    List all available types
- **POST("/api/types"):**
    Create a new type
- **PUT("/api/types/:id"):**
    Update type
        


## Not yet and need to be tested:
- UpdateUserSetting
- CreateNewPomoType
- ListPomoType
- UpdateType