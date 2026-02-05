# Booking System API

High-performance booking system written in Go.

## Status
🚧 Work in Progress

## Problems
This section is designed to describe the significant challenges I face during the project.
### Race condition (demo)

`CreateBooking` does not check the availability of the resource and is not protected from competition. 

In this regard, I decided to conduct a test and wrote a load test script where `50 goroutines` all send `POST /api/bookings` with the same `user_id=7`, `resource_id=1`. The result was as follows:
```RaceCondition.go
Test/Race Condition: 50 goroutines and 1 resource.
User 1: done. resource 1 booked
User 2: done. resource 1 booked
User 3: done. resource 1 booked
User 4: done. resource 1 booked
User 5: done. resource 1 booked
User 6: done. resource 1 booked
User 7: done. resource 1 booked
User 8: fail 400 ({"error":"ОШИБКА: INSERT или UPDATE в таблице \"bookings\" нарушает ограничение внешнего ключа \"bookings_user_id_fkey\" (SQLSTATE 23503)"}
)
...
User 50: fail 400 ({"error":"ОШИБКА: INSERT или UPDATE в таблице \"bookings\" нарушает ограничение внешнего ключа \"bookings_user_id_fkey\" SQLSTA внешнего ключа \"bookings_user_id_fkey\" (SQLSTATE 23503)"}
)

Result RaceCondition/Test:
Total attempts: 50
Done bookings: 7 - bug <- one resource has been booked 7 times
Race Condition: 7 users have received one resource
```

This demonstrates a `race condition`: multiple parallel requests manage to create reservations for the same resource.

`To fix this, I'm thinking about digging into transactions and locks at the database level.`