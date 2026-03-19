# Synconomics Backend API Documentation

Base URL: `http://localhost:8080` (or your deployed server URL)

Semua endpoint API diawali dengan `/api`.

---

## 1. Authentication (Manual)

### Register User
Membuat akun baru dengan nama, email, dan password.

- **URL:** `/api/register`
- **Method:** `POST`
- **Headers:** `Content-Type: application/json`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "secretpassword"
}
```

**Success Response (201 Created):**
```json
{
  "token": "eyJhbGciOiJIUzI...",
  "user": {
    "ID": 1,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "provider": "manual"
  }
}
```

**Error Response (400 Bad Request):**
```json
{
  "error": "invalid body request" // atau pesan error spesifik lainnya
}
```

---

### Login User
Login menggunakan email dan password untuk mendapatkan token JWT.

- **URL:** `/api/login`
- **Method:** `POST`
- **Headers:** `Content-Type: application/json`

**Request Body:**
```json
{
  "email": "johndoe@example.com",
  "password": "secretpassword"
}
```

**Success Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI...",
  "user": {
    "ID": 1,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "provider": "manual"
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "error": "invalid email or password"
}
```

---

## 2. Authentication (Google OAuth)

### Initiate Google Login
Buka URL ini di browser untuk memulai proses login via akun Google.

- **URL:** `/api/auth/google`
- **Method:** `GET`
- **Action:** Akan me-redirect user ke halaman persetujuan Google.

---

### Google OAuth Callback (Internal)
Endpoint ini secara otomatis diakses oleh Google setelah user memberikan persetujuan login. Backend akan mendaftarkan/login user dan me-redirect kembali ke frontend.

- **URL:** `/api/auth/google/callback`
- **Method:** `GET`
- **Action:** Redirect to Frontend URL.  
  Contoh: `http://localhost:3000/auth/callback?token=eyJhbGci...`

---

## 3. Protected Routes

Semua request di bawah ini memerlukan header **Authorization** dengan format Bearer Token yang didapat saat Login atau Register.

**Header Format:**
```
Authorization: Bearer <your_jwt_token_here>
```

### Get Profile
Mengambil data profil pengguna yang sedang login berdasarkan token JWT.

- **URL:** `/api/profile`
- **Method:** `GET`
- **Headers:** `Authorization: Bearer <token>`

**Success Response (200 OK):**
```json
{
  "token": "eyJhbGciOiJIUzI...",
  "user": {
    "ID": 1,
    "name": "John Doe",
    "email": "johndoe@example.com",
    "provider": "google",
    "google_id": "10423...",
    "avatar": "https://lh3.googleusercontent.com/a/..."
  }
}
```

**Error Response (401 Unauthorized):**
```json
{
  "error": "unauthorized" // Token tidak valid atau tidak disediakan
}
```
**Error Response (404 Not Found):**
```json
{
  "error": "user not found"
}
```
