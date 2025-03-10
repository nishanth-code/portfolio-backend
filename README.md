# 🚀 Portfolio Backend – Go & Gin API  

This is the backend service for my personal portfolio, built using **Go** and the **Gin web framework**. It provides APIs for managing projects, blog posts, authentication, and contact form submissions.  

## 🌟 Features  

✅ **Project Management** – Fetch & add portfolio projects dynamically  
✅ **JWT Authentication** – Secure login & token-based access  
✅ **Contact Form API** – Handle user inquiries  
✅ **Optimized Performance** – Fast & lightweight backend with Gin  
✅ **Database Support** – Uses **MongoDB** for data storage  
✅ **Dockerized Deployment** – Easily deploy on cloud platforms  

## 🛠️ Tech Stack  

- **Language:** Go  
- **Framework:** Gin  
- **Database:**  MongoDB  
- **Authentication:** JWT  
 

## 🚀 Installation & Setup  

### 1️⃣ **Clone the Repository**  
```sh
git clone https://github.com/nishanth-code/portfolio-backend.git
cd portfolio-backend

### installation of dependency

go mod tidy


3️⃣ Set Up Environment Variables

PORT=8080
SECRET_KEY="your-secret-key"
DB_URL="your-database-url"


4️⃣ Run the Server

go run main.go
