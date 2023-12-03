# Project Management API

Welcome to the Project Management API! This API is built using Go, PostgreSQL, sqlc, go-chi router, Air, go-migration, and more. Below is a checklist to guide you through setting up and implementing various features.

## Project Setup Checklist

- [x] **Set up PostgreSQL and sqlc**
  - Initialize your PostgreSQL database.
  - Install and configure sqlc for generating Go code from SQL.

- [x] **Integrate Go-chi Router**
  - Implement routes and handlers using the go-chi router for a clean and efficient API structure.

- [x] **Add Air for Hot Reloading**
  - Install Air to enable automatic reloading during development.

- [x] **Utilize go-migration**
  - Implement database migrations using go-migration to manage schema changes.

- [ ] **Other Libraries (Future)**


## Project Features Checklist

- [ ] **Simple API - user_account model**
  - Implement CRUD operations for the user_account model.

- [ ] **Simple API - user_profile model**
  - Develop API endpoints to manage user profiles.

- [ ] **Simple API - project model**
  - Create API routes for handling project-related operations.

- [ ] **Authentication via email and password**
  - Set up user authentication using email and password.

- [ ] **Authorization**
  - Implement authorization logic to control access to various API endpoints.

- [ ] **Social Authentication**
  - Enable social authentication options for users.

## Getting Started

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/DaoVuDat/go-project-management-api.git
   cd project-management-api

2. **Install Dependencies:**
   ```bash
   make install

3. **Database Setup:**
   ```bash
   ... 
   
4. **Run the Application:**
    ```bash
   make dev