# Personal Notes App

A simple Personal Notes App built with Go (backend), Vite (frontend), and PostgreSQL (database). This project demonstrates a full-stack web application with user authentication, CRUD operations for notes, and responsive design.

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Setup and Installation](#setup-and-installation)
  - [Development Environment](#development-environment)
  - [Production Environment](#production-environment)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- User Authentication (Login, Signup)
- Create, Read, Update, and Delete (CRUD) operations for personal notes
- Responsive design with Vite (frontend)
- Backend API built with Go
- PostgreSQL as the database
- Hot reloading for both frontend and backend during development
- Dockerized for easy setup and deployment

## Tech Stack

- **Frontend**: Vite (with React or Vue)
- **Backend**: Go (with Fiber)
- **Database**: PostgreSQL
- **Containerization**: Docker and Docker Compose

## Setup and Installation

### Development Environment

To set up the development environment with hot reloading for both the frontend and backend:

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/personal-notes-app.git
   cd personal-notes-app
   ```

2. **Create `.env` File**:

   Create a `.env` file in the root directory with the following content:

   ```bash
   POSTGRES_USER=your_user
   POSTGRES_PASSWORD=your_password
   POSTGRES_DB=your_db
   DB_HOST=db
   DB_PORT=5432
   ```

3. **Build and Start the Development Environment**:

   ```bash
   docker-compose up backend-dev frontend-dev db
   ```

   - **Backend**: Runs on [http://localhost:8080](http://localhost:8080) with hot reloading using `air`.
   - **Frontend**: Runs on [http://localhost:5173](http://localhost:5173) with Vite's hot reloading.
   - **Database**: Runs on `localhost:5432`.

4. **Access the Application**:

   - Frontend: [http://localhost:5173](http://localhost:5173)
   - Backend API: [http://localhost:8080](http://localhost:8080)

### Production Environment

To set up the production environment:

1. **Build and Start the Production Environment**:

   ```bash
   docker-compose up backend frontend db
   ```

   - **Frontend**: Accessible at [http://localhost](http://localhost).
   - **Backend**: Runs on [http://localhost:8080](http://localhost:8080).
   - **Database**: Runs on `localhost:5432`.

2. **Access the Application**:

   - Frontend: [http://localhost](http://localhost)
   - Backend API: [http://localhost:8080](http://localhost:8080)

### Stopping Services

To stop all services, run:

```bash
docker-compose down
```

## Usage

- **Frontend Development**: Edit the code in the `frontend/src` directory. Changes will automatically reflect due to Vite's hot reloading.
- **Backend Development**: Edit the code in the `backend` directory. Changes will automatically reflect due to `air` hot reloading.
- **Database Management**: Modify the `db/init.sql` file to change the initial state of the database.