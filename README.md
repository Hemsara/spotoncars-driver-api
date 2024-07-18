# SpotOnCars Server

This repository contains the server-side code for the SpotOnCars admin tracking application. It is built using Go and the Gin web framework.

## Features

### Authentication

- **Admin Login**
  - Endpoint: `/auth/login`
  - Controller: `authController.LoginAdmin`

### Driver Application

- **Log Driver Activity**
  - Endpoint: `/driverApp/drivers/log`
  - Middleware: `middleware.DriverAuthenticationGuard`
  - Controller: `dvrController.LogDriver`

### Booking Management

- **Get Active Bookings**
  - Endpoint: `/bookings/active`
  - Middleware: `middleware.AuthenticationGuard`
  - Controller: `bookingController.GetActiveBookings`
  
- **Get Booking History**
  - Endpoint: `/bookings/history`
  - Middleware: `middleware.AuthenticationGuard`
  - Controller: `bookingController.GetBookingsHistory`

### Driver Management

- **Get All Drivers**
  - Endpoint: `/drivers`
  - Middleware: `middleware.AuthenticationGuard`
  - Controller: `dvrController.GetAllDrivers`

### Notifications

- **Send Notification**
  - Endpoint: `/alerts`
  - Middleware: `middleware.AuthenticationGuard`
  - Controller: `alertController.SendNotification`


