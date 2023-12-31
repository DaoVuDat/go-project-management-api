# Project Management API

Welcome to the Project Management API! This API is built using Go, PostgreSQL, sqlc, Echo Framework, Air, go-migration, and more. Below is a checklist to guide you through setting up and implementing various features.

## Project Setup Checklist

- [x] **Set up PostgreSQL and sqlc**
  - Initialize your PostgreSQL database.
  - Install and configure sqlc for generating Go code from SQL.

- [x] **Integrate Echo Framework**
  - Implement routes and handlers using the Echo router for a clean and efficient API structure.

- [x] **Add Air for Hot Reloading**
  - Install Air to enable automatic reloading during development.

- [x] **Utilize go-migration**
  - Implement database migrations using go-migration to manage schema changes.

- [x] **Integrate ozzo-validation**
  - Validate incoming requests
  
- [ ] **Other Libraries (Future)**


## Project Features Checklist

- [x] **Simple API - user_account model**
  - Implement CRUD operations for the user_account model.

- [x] **Simple API - user_profile model**
  - Develop API endpoints to manage user profiles.

- [x] **Simple API - project model**
  - Create API routes for handling project-related operations.

- [x] **Authentication via email and password**
  - Set up user authentication using email and password.

- [x] **Authorization**
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