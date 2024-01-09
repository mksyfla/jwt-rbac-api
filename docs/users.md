# Users API Spec

## register User API

done

Endpoint : POST /api/v1/users

Body Request :
```
name: string, required
email: string, required
password: string, required
category: string, required (UMKM | MAHASISWA | LAINNYA)
```

Reponse :
```
HTTP Response 201
{
  "message": "user created"
  "data": {
    "id": "user_id" 
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

## Login User API

done

Endpoint : POST /api/v1/login

Body Request :
```
email: string, required
password: string, required
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

## Get User API
don\'t know how to serving the image 

Endpoint: GET /api/v1/users/{id}

Response :
- UMKM :
```
HTTP Response 200
{
  "name": string
  "email": string
  "verified": boolean
  "banner": string
  "profile": string
  "jobs": [
    {
      "title": string
      "description": string
      "image": string
    },
    {
      "title": string
      "description": string
      "image": string
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
  "banner": string
  "profile": string
  "jobs_result": [
    {
      "title": string
      "description": string
      "image": string
    },
    {
      "title": string
      "description": string
      "image": string
    }
  ]
}
```

## Update User API

not yet

Endpoint : PUT /api/v1/users/{id}

Headers:
- Authorization: token

Body Request :
```
name: string, required
email: string, required
password: string, required
skills: string, required
profile: file, required
banner: file, required
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
      "profile": string
    },
    {
      "name": string
      "category": string
      "profile": string
    }
  ]
}
```

## Post a Request to get Verified

Endpoint: POST /api/v1/users/{id}/verified

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

Endpoint: POST /api/v1/users/{id}/expert

Headers:
- Authorization: token

Response :
```
HTTP Response 200
{
  "message": "request uploaded"
}
```