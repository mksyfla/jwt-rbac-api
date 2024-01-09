# Users API Spec

## register User API (Done)

Endpoint : POST /api/v1/users

Body Request :
```
name: string
email: string
password: string
category: string (UMKM | MAHASISWA | LAINNYA)
```

Reponse :
```
HTTP Response 201
{
  "message": "user created"
  "data": {
    "user_id": string
  }
}
```

Response Body Error :
- body request not complete
```
HTTP Response 400
{
  "message": "field is empty"
}
```
- email is used
```
HTTP Response 400
{
  "message": "email is used"
}
```

## Update User API

not yet

Endpoint : PUT /api/v1/users/:job_id

Headers:
- Authorization: token

Body Request :
```
name: string
email: string
password: string
skills: string
profile: file (base64, jpg)
banner: file (base64, jpg)
```

Reponse :
```
HTTP Response 200
{
  "message": "user updated"
}
```

Response Body Error :
- id not found
```
HTTP Response 404
{
  "message": "id not found"
}
```
- data tidak valid
```
HTTP Response 400
{
  "message": "field is empty"
}
```

## Login User API (Done)

done

Endpoint : POST /api/v1/login

Body Request :
```
email: string
password: string
```

Reponse :
Cookie Authorization
```
HTTP Response 201
{
  "message": "login success"
}
```

Response Body Error :
```
HTTP Response 400
{
  "message": "email or password wrong"
}
```

## Get Users API

done

Endpoint: GET /api/v1/users/

Response :
```
HTTP Response 200
{
  "data": [
    {
      "name": string
      "category": string
      "profile": string (image url)
    },
    {
      "name": string
      "category": string
      "profile": string (image url)
    }
  ]
}
```

## Get User Details API
don\'t know how to serving the image 

Endpoint: GET /api/v1/users/:job_id

Response :
- UMKM :
```
HTTP Response 200
{
  "user_id": string
  "user_name": string
  "user_email": string
  "user_verified": boolean
  "user_banner": string (image url)
  "user_profile": string (image url)
  "user_jobs": [
    {
      "job_id": string
      "job_title": string
      "job_description": string
      "job_image": string (image url)
    },
    {
      "job_id": string
      "job_title": string
      "job_description": string
      "job_image": string (image url)
    }
  ]
}
```
- MAHASISWA :
```
HTTP Response 200
{
  "name": string
  "email": string
  "badge": boolean
  "banner": string (image url)
  "profile": string (image url)
  "result": [
    {
      "title": string
      "description": string
      "image": string (image url)
    },
    {
      "title": string
      "description": string
      "image": string (image url)
    }
  ]
}
```


## Post a Request to get Verified

Endpoint: POST /api/v1/users/:job_id/verified

Headers:
- Authorization: token

Response :
```
HTTP Response 200
{
  "message": "request uploaded"
}
```

## Post a Request to get Expert Badge

Endpoint: POST /api/v1/users/:job_id/expert

Headers:
- Authorization: token

Response :
```
HTTP Response 200
{
  "message": "request uploaded"
}
```