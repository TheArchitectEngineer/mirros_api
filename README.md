# mirros_api — Go port (work in progress)

This branch contains an initial Go scaffold for mirros_api.

Quickstart (local using docker-compose)
1. Copy `.env.example` to `.env` and edit if needed.
2. Start services:
   docker-compose up --build
3. The API will be available at http://localhost:8080
   - Health: GET /health
   - Users: GET /api/users, POST /api/users

What this scaffold includes
- Gin HTTP router
- GORM + Postgres database connection
- Example User model + handlers
- Dockerfile and docker-compose for local development

Planned next steps
- Translate all endpoints from the existing repository to Go (controllers, services, middleware).
- Add authentication (JWT), migrations, and tests.
- Open a draft PR from mirr_os_go with the initial port, then iterate.

If you want any changes to the layout (e.g., prefer chi or sqlx instead of Gin/GORM), tell me and I’ll adapt.