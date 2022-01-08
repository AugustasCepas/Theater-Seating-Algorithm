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
1) cd Theater-Seating-Algorithm/Docker
2) Enter git credentials: git username and git private SSH key in docker-compose.yml (required to fetch go modules via git)
3) docker-compose up
```

## To access DB:
```
psql --host=database --username=admin --dbname=theater_seating_database

pass1234
```