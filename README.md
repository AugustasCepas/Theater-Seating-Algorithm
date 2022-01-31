# Theater Seating Algorithm / API

# Admin

## To install & run:
``` 
1) Update Theater-Seating-Algorithm/Docker/docker-compose.yml: git_username (can be found through cmd with 'git config --list' call - user.name) and github_personal_token (can be found at C:\Users\*user*\.ssh id_ed25519 or similar named file) - this step is required to fetch go modules via git
2) Open cmd
3) call 'cd Docker'
4) call 'docker-compose up'
```

## To access DB:
```
psql --host=database --username=admin --dbname=theater_seating_database

password: pass1234
```

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
0 - means that seat is not reserved

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
