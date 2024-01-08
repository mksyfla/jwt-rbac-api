# Jobs API Spec

## Create Job API

Endpoint : POST /api/v1/jobs

Headers:
- Authorization: token
- Role: UMKM

Body Request :
```
title: string, reqiured
description: string, required
deadline: time/date?, required
reward: string
tags: array?
draft: boolean, required (false)
```

Response :
```
HTTP Response 201
{
  "message": "job created"
  "data": {
    "id": "job_id" 
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

## Draft Job API

Endpoint : POST /api/v1/jobs

Headers:
- Authorization: token
- Role: UMKM

Body Request :
```
title: string
description: string
deadline: time/date?
reward: string
tags: array?
draft: boolean (true)
```

Response :
```
HTTP Response 201
{
  "message": "drafted"
  "data": {
    "id": "job_id"
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
