# CORS Testing:


- Make sure to export env variables into your local system env variables:
```sh
export ALLOWED_ORIGINS=http://localhost:3000,https://myfrontend.com
go run main.go
```

## Test 1:

- Make a GET Request on endpoint with Header:
- GET Request: `http://localhost:8080/student`

- Header:
```plain
Accept: */*
User-Agent: Sahil Sharma
Origin: http://localhost:3000
Content-Type: application/json
Authorization: Sahil sahilwep
```
- You will get an empty JSON [], Everything is okey

## Test 2:

- Try using this request, this will block as it's origin is changed.:
```plain
Accept: */*
User-Agent: Sahil Sharma
Origin: http://evilHost:3000
Content-Type: application/json
Authorization: Sahil sahilwep
```

- You will get response:
```plain
{
  "error": "CORS policy: Origin not allowed"
}
```