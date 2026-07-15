# Parking Zone API 🚗⚡

A Smart Parking & EV Charging Reservation System built with Go, designed for airports, malls, and other high-traffic facilities. The platform allows users to reserve parking spots while ensuring EV charging zones never exceed capacity, even under concurrent requests.

## Features

### Authentication
- User Registration
- User Login
- JWT-based Authentication
- Role-based Authorization (Driver & Admin)

### Parking Zones
- View all parking zones
- View zone availability in real-time
- Create, update, and manage parking zones (Admin)

### Reservations
- Reserve parking spots
- View personal reservations
- Cancel reservations
- View all reservations (Admin)

### Concurrency Protection
The reservation system prevents overbooking using:
- Database Transactions
- Row-Level Locking (`FOR UPDATE`)
- Atomic capacity validation

This ensures that when multiple users attempt to reserve the last available spot simultaneously, only one reservation succeeds.

---

## Tech Stack

- Go 1.22+
- Echo Framework
- GORM
- PostgreSQL
- JWT Authentication
- bcrypt Password Hashing
- go-playground/validator

---

## Architecture

This project follows **Clean Architecture** principles:

```
Handler → Service → Repository → Database
```

### Layers

- **DTO**: Request and response structures
- **Handler**: HTTP request handling
- **Service**: Business logic
- **Repository**: Database operations
- **Models**: Database entities

Dependency Injection is used to wire all layers together.

---

## User Roles

### Driver
- Register & Login
- View parking zones
- Reserve parking spots
- Manage own reservations

### Admin
- All Driver permissions
- Create parking zones
- Update pricing
- View all reservations

---

## Main API Routes

### Authentication
- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`

### Parking Zones
- `GET /api/v1/zones`
- `GET /api/v1/zones/:id`
- `POST /api/v1/zones` (Admin)

### Reservations
- `POST /api/v1/reservations`
- `GET /api/v1/reservations/my-reservations`
- `DELETE /api/v1/reservations/:id`
- `GET /api/v1/reservations` (Admin)

---

## Security

- Passwords are hashed using bcrypt
- JWT authentication for protected routes
- Role-based access control
- Input validation using validator package

---

## Future Improvements

- Reservation time slots
- Payment integration
- QR code check-in
- Email notifications
- Analytics dashboard

---

## Author

Built as a backend system design project using Go, PostgreSQL, Echo, and GORM while following Clean Architecture and industry-standard authentication practices.