# Users API Spec

## register User API

Endpoint : POST /api/v1/users

Body Request :
```
nama: string, required
email: string, required
password: string, required
category: string, required (UMKM | MAHASISWA)
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
Endpoint : POST /api/v1/login

Body Request :
```
email: string, required
password: string, required
```

Reponse :
```
HTTP Response 201
{
  "data": {
    "token": token
  }
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
Endpoint: GET /api/v1/users/{id}

Response :
- UMKM :
```
HTTP Response 200
{
  nama: string
  email: string
  verified: boolean
  banner: string
  profile: string
  jobs: [
    {
      title: string
      description: string
      image: string
    },
    {
      title: string
      description: string
      image: string
    }
  ]
}
```
- MAHASISWA :
```
HTTP Response 200
{
  nama: string
  email: string
  badge: boolean
  banner: string
  profile: string
  jobs_reuslt: [
    {
      title: string
      description: string
      image: string
    },
    {
      title: string
      description: string
      image: string
    }
  ]
}
```

## Update User API

Endpoint : PUT /api/v1/users/{id}

Headers:
- Authorization: token

Body Request :
```
nama: string, required
email: string, required
password: string, required
posisi: string, required
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