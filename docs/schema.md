# DB Table Schema

## questions

|Name|Type|Null|Key|Default|説明|
|-|-|-|-|-|-|
|id|VARCHAR(36)|NO|PRI|||
|question|LONGTEXT|NO|||質問文|
|answer|LONGTEXT||||回答文|
|answerer|VARCHAR(32)||||回答者のid|
|created_at|datetime|||current_timestamp|質問された日|
|updated_at|datetime|||current_timestamp|回答が更新された日|

