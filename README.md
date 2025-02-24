# Cattle Farm Management API

A robust REST API for managing cattle farms, built with Go and Fiber framework. This system helps farmers track their livestock, monitor health records, manage breeding cycles, and optimize milk production.

## Features

### Cattle Management

- Track individual cattle with unique identification
- Monitor pregnancy status and breeding cycles
- Record birth events and maintain genealogy
- Track weight records
- Manage death records

### Milk Production

- Record daily milk production
- Track milking periods
- Monitor milk production efficiency
- Identify milkable cattle automatically

### Health Management

- Record and track illnesses
- Monitor antibiotic usage
- Track treatment periods
- Manage milking restrictions during treatment

### Breeding Management

- Track insemination records
- Monitor pregnancy status
- Manage breeding cycles
- Record birth events

## Tech Stack

- **Go** - Programming Language
- **Fiber** - Web Framework
- **GORM** - ORM Library
- **SQLite** - Database
- **JWT** - Authentication
- **Redis** - Caching (optional)

## Installation

1. Clone the repository

```bash
git clone https://github.com/zxselimcan/cattle-farm-api
cd cattle-farm-api
```

2. Install dependencies

```bash
go mod download
```

3. Create .env file

```env
JWT_SECRET=your_jwt_secret
SMTP_HOST=your_smtp_host
SMTP_PORT=your_smtp_port
SMTP_MAIL=your_smtp_mail
SMTP_PASSWORD=your_smtp_password
```

4. Run the application

```bash
GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 go run main.go
```

## API Documentation

### Authentication Endpoints

- `POST /auth/register` - Register new user
- `POST /auth/login` - Login user
- `POST /auth/email-verification` - Verify email

### Cattle Endpoints

- `GET /api/cattle` - Get all cattle
- `GET /api/cattle/milkable` - Get milkable cattle
- `GET /api/cattle/pregnant` - Get pregnant cattle
- `POST /api/cattle` - Add new cattle

### Health Record Endpoints

- `POST /api/cattle/:cattle_uuid/illness-records` - Add illness record
- `GET /api/cattle/:cattle_uuid/illness-records` - Get illness records

### Milk Production Endpoints

- `POST /api/cattle/:cattle_uuid/milking-records` - Add milking record
- `GET /api/cattle/:cattle_uuid/milking-records` - Get milking records

### Weight Record Endpoints

- `POST /api/cattle/:cattle_uuid/weight-records` - Add weight record
- `GET /api/cattle/:cattle_uuid/weight-records` - Get weight records

## Data Models

### Cattle Status Types

- `MILKABLE` - Ready for milking
- `DRY_1` - First dry period
- `DRY_2` - Second dry period
- `DRY_3` - Third dry period
- `PREGNANT` - Currently pregnant
- `NOT_READY` - Not ready for milking
- `DEAD` - Deceased
- `MALE` - Male cattle

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Acknowledgments

- Special thanks to veterinarians and farmers who provided domain expertise:
  - Collaborated with veterinarians to understand cattle health management, breeding cycles, and medical record-keeping requirements
  - Conducted on-site visits to dairy farms and milking facilities to study operational workflows
  - Gathered detailed feedback from dairy farm operators about their daily challenges and needs
  - Analyzed milk production processes and facility management systems through first-hand observations
  - Integrated real-world farming practices and industry standards into the system design

## Support

For support, email zxselimcan@icloud.com or open an issue in the repository.
