# Theater Seating Algorithm / API

# Endpoints:
## Reserve Seats
```
POST: http://localhost:8000/api/v1/layout/1/section/1/rank/1/reserve
BODY:
{
   "reserve": "3"
}

RESULT:
{
    "result": "3 Seat(s) reserved"
}
```

## Get Layout Seats
```
GET: http://localhost:8000/api/v1/layout/1/section/1/seats
RESULT:
[
    "12234455",
    "00000666",
    "00000000"
]
```

## Get User Reserved Seats
```
GET: http://localhost:8000/api/v1/seats/2
RESULT:
[
    {
        "Number": 3,
        "RowId": 1
    },
    {
        "Number": 5,
        "RowId": 1
    }
]
```


# Admin

## To install & run:
```
1) cd Theater-Seating-Algorithm/src
2) go get .
3) go run *.go
```

## To run DB:
```
docker-compose up

psql --host=database --username=unicorn_user --dbname=theater_seating_database

magical_password
```