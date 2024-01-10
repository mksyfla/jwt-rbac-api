# Jobs API Spec

## Create Job API (Done)

Endpoint : POST /api/v1/jobs

Headers:
- Authorization: token
- Role: UMKM

Body Request :
```
title: string
description: string
deadline: int (epoch)
reward: string, optional
tag: string
image: []string (base64, jpg)
```

Response :
```
HTTP Response 201
{
  "message": "job created"
  "data": {
    "id": string 
  }
}
```

Response Body Error :
```
HTTP Response 400
{
  "message": "field is empty"
}
```

## Draft Create Job API (Done)

Endpoint : POST /api/v1/draft

Headers:
- Authorization: token
- Role: UMKM

Body Request :
```
title: string, optional
description: string, optional
deadline: int (epoch), optional
reward: string, optional
tag: string, optional
image: []string (base64, jpg), optional
```

Response :
```
HTTP Response 201
{
  "message": "drafted"
  "data": {
    "id": string
  }
}
```

## Update Job API

Endpoint : PUT /api/v1/jobs/:job_id

Headers:
- Authorization: token
- Role: UMKM

Body Request :
```
title: string
description: string
deadline: int (epoch)
reward: string, optional
tag: string
image: []string (base64, jpg)
```

Response :
```
HTTP Response 200
{
  "message": "job updated"
}
```

Response Body Error :
```
HTTP Response 400
{
  "message": "field is empty"
}
```

Response Body Error :
```
HTTP Response 404
{
  "message": "job not found"
}
```

## Delete Job API

Endpoint : Delete /api/v1/jobs/:job_id

Headers:
- Authorization: token
- Role: UMKM

Response :
```
HTTP Response 200
{
  "message": "deleted"
}
```

Response Body Error :
```
HTTP Response 404
{
  "message": "job not found"
}
```

## Create Comment Jobs API

Endpoint : POST /api/v1/jobs/:job_id/comments/

Headers :
- Authorization

Body Request : 
```
message: string
```

Response :
```
HTTP Response 201
{
  "message": "comment added"
  "data": {
    "id": "comment_id" 
  }
}
```

Response Body Error :
```
HTTP Response 400
{
  "message": "no comment"
}
```

## Delete Comment Jobs API

Endpoint : DELETE /api/v1/jobs/:job_id/comments/:comment_id

Headers :
- Authorization

Body Request : 
```
message: string
```

Response :
```
HTTP Response 200
{
  "message": "comment deleted"
}
```

Response Body Error :
```
HTTP Response 404
{
  "message": "comment not found"
}
```

## POST Reply Jobs API

Endpoint : POST /api/v1/jobs/:job_id/comments/:comment_id/replies

Headers :
- Authorization

Body Request : 
```
message: string
```

Response :
```
HTTP Response 201
{
  "message": "reply added"
  "data": {
    "id": string
  }
}
```

Response Body Error :
```
HTTP Response 400
{
  "message": "no reply"
}
```

## Delete Reply Jobs API

Endpoint : DELETE /api/v1/jobs/:job_id/comments/:comment_id/replies/:reply_id

Headers :
- Authorization

Body Request : 
```
message: string
```

Response :
```
HTTP Response 200
{
  "message": "reply added"
}
```

Response Body Error :
```
HTTP Response 404
{
  "message": "reply not found"
}
```

## Get Jobs API

Endpoint : Get /api/v1/jobs

Response Body :
```
HTTP Response 200
{
  "data": [
    {
      "job_id": string
      "job_title": string
      "job_description": string
      "job_picture": string (image url)
      "user_name": string
    },
    {
      "job_id": string
      "job_title": string
      "job_description": string
      "job_picture": string (image url)
      "user_name": string
    }
  ],
}
``` 

## Get Detail Job API

Endpoint : Get /api/v1/jobs/:job_id

Response Body :
```
HTTP Response 200
{
  "data": {
    "id_job": string
    "job_title": string
    "job_description": string
    "job_picture": string (image url)
    "user_name": string
    "result": [
      {
        "result_id": string
        "result_picture": string (image url)
      },
      {
        "result_id": string
        "result_picture": string (image url)
      }
    ],
    "comments": [
      {
        "comment_id": string
        "user_name": string
        "user_picture": string (image url)
        "comment_message": string
        "replies": [
          {
            "reply_id": string
            "user_name": string
            "user_picture": string (image url)
            "reply_message": string
          }
        ]
      }
    ]
  }
}
``` 