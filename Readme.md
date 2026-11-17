# Personal Portfolio Website

A personal portfolio website built with **Golang, Gin Framework, Bootstrap 5, and Clean Architecture**. This application is designed to showcase professional profiles, work experience, technical skills, projects, and contact information through a responsive and modern web interface.

The application uses **JSON files as a lightweight data storage solution**, without requiring any relational or NoSQL database.

---

## 1. Project Overview

The Personal Portfolio Website serves as a centralized platform to present professional identity, technical expertise, career experience, and project achievements.

### Key Features

- **Professional Profile** – Display personal information, professional summary, and career objectives.
- **Work Experience** – Showcase professional experience, internships, and responsibilities.
- **Education** – Display academic background and qualifications.
- **Technical Skills** – Present technical competencies categorized by expertise.
- **Portfolio Projects** – Showcase selected projects with descriptions, technologies, and responsibilities.
- **Certifications & Achievements** – Display relevant certifications and accomplishments.
- **Contact Information** – Provide email, LinkedIn, and other professional communication links.
- **Responsive Design** – Adapt the interface to desktop, tablet, and mobile devices.

---

## 2. Technology Stack

| Component | Technology |
|---|---|
| Backend | Golang |
| Web Framework | Gin Framework |
| Frontend | HTML5, CSS3, JavaScript |
| UI Framework | Bootstrap 5 |
| Icons | Bootstrap Icons |
| Data Storage | JSON Files |
| Architecture | Clean Architecture |
| Template Engine | Go HTML Template |
| Version Control | Git & GitHub |

---

## 3. System Architecture

The application implements Clean Architecture to separate business logic, data access, and presentation concerns.

Each layer has a clearly defined responsibility to improve maintainability, scalability, and testability.

### Architecture Layers

**1. Domain Layer**
- Defines core entities and business models.
- Contains repository interfaces.
- Independent of external frameworks and infrastructure.

**2. Use Case Layer**
- Handles application business logic.
- Defines operations for retrieving profile, experiences, skills, and projects.
- Coordinates interactions between domain and repository implementations.

**3. Infrastructure Layer**
- Implements JSON-based data repositories.
- Handles reading and parsing JSON files.
- Provides data persistence and retrieval mechanisms.

**4. Delivery Layer**
- Handles HTTP requests and responses.
- Implements Gin routes and controllers.
- Renders HTML templates using Go HTML Template.

### Architecture Flow

```text
        USER / BROWSER
               |
               v
        GIN HTTP ROUTER
               |
               v
       DELIVERY / HANDLER
               |
               v
          USE CASE
               |
               v
       DOMAIN / ENTITIES
               |
               v
      REPOSITORY INTERFACE
               |
               v
    INFRASTRUCTURE / JSON
               |
               v
         JSON STORAGE
```

---

## 4. Project Structure

```text
portfolio/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   │   ├── profile.go
│   │   │   ├── experience.go
│   │   │   ├── education.go
│   │   │   ├── skill.go
│   │   │   └── project.go
│   │   │
│   │   └── repository/
│   │       ├── profile_repository.go
│   │       ├── experience_repository.go
│   │       ├── education_repository.go
│   │       ├── skill_repository.go
│   │       └── project_repository.go
│   │
│   ├── usecase/
│   │   ├── profile_usecase.go
│   │   ├── experience_usecase.go
│   │   ├── education_usecase.go
│   │   ├── skill_usecase.go
│   │   └── project_usecase.go
│   │
│   ├── infrastructure/
│   │   └── persistence/
│   │       └── json/
│   │           ├── profile_repository.go
│   │           ├── experience_repository.go
│   │           ├── education_repository.go
│   │           ├── skill_repository.go
│   │           └── project_repository.go
│   │
│   └── delivery/
│       └── http/
│           ├── handler/
│           │   └── portfolio_handler.go
│           └── routes/
│               └── routes.go
│
├── data/
│   ├── profile.json
│   ├── experiences.json
│   ├── education.json
│   ├── skills.json
│   └── projects.json
│
├── web/
│   ├── templates/
│   │   ├── layouts/
│   │   │   └── base.html
│   │   ├── pages/
│   │   │   ├── index.html
│   │   │   ├── about.html
│   │   │   ├── experience.html
│   │   │   └── projects.html
│   │   └── components/
│   │       ├── navbar.html
│   │       ├── footer.html
│   │       └── project_card.html
│   │
│   └── static/
│       ├── css/
│       │   └── style.css
│       ├── js/
│       │   └── main.js
│       ├── images/
│       │   ├── profile.jpg
│       │   └── projects/
│       └── assets/
│
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

---

## 5. JSON Data Storage

This application does not use a database. All portfolio information is stored in JSON files inside the `data/` directory.

### Example: profile.json

```json
{
  "name": "Eka Yulia Sari",
  "title": "IT Project Manager | System Analyst | Full-Stack Developer",
  "email": "your-email@example.com",
  "location": "Yogyakarta, Indonesia",
  "linkedin": "https://www.linkedin.com/",
  "github": "https://github.com/",
  "summary": "IT professional with experience in project management, system analysis, and full-stack software development."
}
```

### Example: experiences.json

```json
[
  {
    "company": "PT Global Data Inspirasi",
    "position": "System Analyst",
    "year": "2026",
    "description": "Conducted business process analysis, requirements gathering, SRS documentation, and system modeling."
  },
  {
    "company": "BLKPP DIY",
    "position": "Golang Backend Engineer",
    "year": "2026",
    "description": "Developed backend services using Golang and Gin Framework with Clean Architecture."
  }
]
```

### Example: skills.json

```json
[
  {
    "category": "Backend Development",
    "items": [
      "Golang",
      "Gin Framework",
      "PHP",
      "CodeIgniter",
      "RESTful API"
    ]
  },
  {
    "category": "Frontend Development",
    "items": [
      "HTML5",
      "CSS3",
      "JavaScript",
      "Bootstrap",
      "AJAX"
    ]
  },
  {
    "category": "Project Management",
    "items": [
      "Agile",
      "Scrum",
      "Kanban",
      "Requirements Engineering",
      "System Analysis"
    ]
  }
]
```

### Example: projects.json

```json
[
  {
    "id": 1,
    "name": "Intelligent Inventory Management System",
    "category": "Web Development",
    "description": "A web-based inventory management platform supporting asset tracking, lending, complaints, and maintenance.",
    "technologies": [
      "Golang",
      "Gin",
      "Bootstrap",
      "PostgreSQL"
    ],
    "image": "/static/images/projects/inventory.png"
  },
  {
    "id": 2,
    "name": "Stock Price Forecasting Using Machine Learning",
    "category": "Data Science",
    "description": "A comparative stock price forecasting study using statistical and machine learning models.",
    "technologies": [
      "Python",
      "ARIMA",
      "ARIMAX",
      "Random Forest",
      "XGBoost",
      "LSTM",
      "GRU"
    ],
    "image": "/static/images/projects/forecasting.png"
  }
]
```

---

## 6. Application Pages

| Page / Section | Description |
|---|---|
| Home | Hero section, introduction, professional title, and call-to-action |
| About | Professional summary and personal profile |
| Experience | Work experience and internship history |
| Education | Academic background |
| Skills | Technical and professional competencies |
| Projects | Portfolio projects with descriptions and technologies |
| Contact | Email, LinkedIn, GitHub, and contact information |

---

## 7. Installation & Setup

### Prerequisites

Make sure the following tools are installed:

- Go 1.22 or later
- Git
- Visual Studio Code
- Web Browser

### Step 1: Clone Repository

```bash
git clone https://github.com/username/portfolio.git

cd portfolio
```

### Step 2: Initialize Go Module

```bash
go mod init portfolio
```

### Step 3: Install Dependencies

```bash
go get github.com/gin-gonic/gin

go mod tidy
```

### Step 4: Configure JSON Data

Update the JSON files inside the `data/` directory with your personal information, experience, skills, education, and projects.

### Step 5: Run Application

```bash
go run cmd/server/main.go
```

### Step 6: Access Website

Open your browser and navigate to:

```text
http://localhost:8080
```

---

## 8. Development Guidelines

To maintain Clean Architecture principles, follow these guidelines:

1. Keep business logic inside the Use Case layer.
2. Keep domain entities independent of Gin and JSON implementation details.
3. Define repository interfaces in the Domain layer.
4. Implement JSON file access in the Infrastructure layer.
5. Handle HTTP requests and HTML rendering in the Delivery layer.
6. Keep HTML templates separate from backend logic.
7. Store portfolio content in JSON files rather than hardcoding it in HTML.
8. Avoid direct file access from HTTP handlers.

**Note:** JSON storage is suitable for a personal portfolio with relatively static content. For this implementation, JSON files are primarily read-only data sources. If write operations are introduced, appropriate file locking and validation should be implemented.

---

## 9. Future Improvements

Potential enhancements include:

- Add animations and interactive UI components.
- Add project detail pages.
- Implement project filtering by technology or category.
- Add downloadable CV functionality.
- Add SEO optimization and metadata.
- Add dark mode.
- Integrate a contact form with email delivery.
- Deploy the application to a cloud hosting platform.

---

## 10. License

This project is developed for personal portfolio and professional presentation purposes.

---

**Built with Golang, Gin Framework, Bootstrap 5, and Clean Architecture.**