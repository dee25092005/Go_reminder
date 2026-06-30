# Go Modular Onboarding

ບັນທຶກການຮຽນພື້ນຖານພາສາ Go ຂອງຕົວເອງ ເຜື່ອຍ້້ຳຄວາມຈຳ.

---

## 1. Go ແມ່ນຫຍັງ?

Go ແມ່ນພາສາໂປຣແກຣມທີ່ **compiled** (ແປເປັນພາສາຄອມພິວເຕີໂດຍກົງ) ພັດທະນາໂດຍ **Google**.

- ເຮັດວຽກໄວເພາະ compile ເປັນ binary ໂດຍກົງ
- syntax ອ່ານງ່າຍ ທຽບເທົ່າກັບ Python
- ເໝາະສຳລັບ backend, CLI tools, ແລະ system programming

---

## 2. ຄວາມແຕກຕ່າງ (ບໍ່ມີ OOP)

Go **ບໍ່ມີ class ແລະ ບໍ່ມີ inheritance** ແບບພາສາ OOP ທົ່ວໄປ. ແທນທີ່ນັ້ນ Go ໃຊ້:

- **ຂໍ້ມູນ** → ຫຸ້ມຫໍ່ດ້ວຍ `struct`
- **ພຶດຕິກຳ (Behavior)** → ຜູກມັດກັບ `struct` ໂດຍໃຊ້ `func` ທີ່ມີ receiver (ເອີ້ນວ່າ **method**)

```go
type User struct {
    Name string
    Age  int
}

func (u User) Greet() {
    fmt.Println("Hello,", u.Name)
}
```

---

## 3. Public vs Private

Go **ບໍ່ໃຊ້** keyword `public` ຫຼື `private`. ມັນໃຊ້**ຕົວອັກສອນຕົວທຳອິດຂອງຊື່** ເພື່ອກຳນົດການເຂົ້າເຖິງ:

| ການຂຽນ | ຄວາມໝາຍ | ຕົວຢ່າງ |
|--------|---------|---------|
| ຕົວອັກສອນໃຫຍ່ | **Exported** (ໃຊ້ນອກ package ໄດ້) | `type User struct` |
| ຕົວອັກສອນນ້ອຍ | **Unexported** (ໃຊ້ໄດ້ສະເພາະໃນ package) | `type user struct` |

```go
type User struct {        // ສົ່ງອອກໄດ້
    Name  string          // field ສົ່ງອອກໄດ້
    email string          // field ພາຍໃນເທົ່ານັ້ນ
}
```

---

## 4. Value Receiver (ຕົວຮັບແບບ Value)

ເມື່ອເອີ້ນໃຊ້ method, Go ຈະ **copy** ຂໍ້ມູນທັງໝົດຂອງ struct ມາໃຊ້.

-  **ບໍ່ສາມາດ**ປ່ຽນແປງຂໍ້ມູນຕົ້ນສະບັບ (ເຮັດວຽກກັບ copy ເທົ່ານັ້ນ)
-  ປອດໄພຈາກການແກ້ໄຂໂດຍບໍ່ຕັ້ງໃຈ (Read-only)
-  ໃຊ້ memory ເພີ່ມ ຖ້າ struct ໃຫຍ່

```go
type User struct {
    Name string
}

// Value receiver (ບໍ່ມີເຄື່ອງໝາຍ *)
func (u User) DisplayName() {
    fmt.Println(u.Name)
}
```

---

## 5. Pointer Receiver (ຕົວຮັບແບບ Pointer)

ເມື່ອເອີ້ນໃຊ້ method, Go ຈະສົ່ງ **ທີ່ຢູ່ memory (address)** ຂອງ struct ຕົ້ນສະບັບໂດຍກົງ.

-  **ສາມາດ**ປ່ຽນແປງຂໍ້ມູນຕົ້ນສະບັບໄດ້ທັນທີ
-  ປະຢັດ memory ແລະ ໄວ ກໍລະນີ struct ໃຫຍ່
-  ຕ້ອງລະວັງ ເພາະຂໍ້ມູນຖືກແກ້ໄຂໄດ້ໂດຍກົງ

```go
// Pointer receiver (ມີເຄື່ອງໝາຍ * ໜ້າ type)
func (u *User) ChangeName(newName string) {
    u.Name = newName  // ຂໍ້ມູນຕົ້ນສະບັບຖືກປ່ຽນທັນທີ
}
```

### ສະຫຼຸບ: ເລືອກໃຊ້ແບບໃດ?

| ສະຖານະການ | ໃຊ້ |
|-----------|-----|
| ຕ້ອງການແກ້ໄຂ struct | **Pointer** (`*T`) |
| struct ມີຂະໜາດໃຫຍ່ | **Pointer** (`*T`) |
| struct ນ້ອຍ ແລະ ບໍ່ຕ້ອງການແກ້ໄຂ | **Value** (`T`) |

---

## 6. ຊັ້ນ HTTP Request (Headers & Content Types)

ການເຊື່ອມຕໍ່ທຸກຄັ້ງເລີ່ມຈາກຝັ່ງ client (ເຊັ່ນ: Postman ຫຼື mobile app). client ຕ້ອງບອກ server ວ່າຂໍ້ມູນທີ່ສົ່ງມາແມ່ນຫຍັງ ແລະ ໃຜເປັນຄົນສົ່ງ — ຜ່ານ **headers**.

- **`Content-Type: application/json`** → ບອກ Go server ວ່າ "payload ໃນ request body ແມ່ນ JSON string, ບໍ່ແມ່ນ plain text"
- **`Authorization: Bearer <JWT>`** → header ປະເພດ passport. ແທນທີ່ຈະສົ່ງ password ໄປຊໍ້າໆ, client ສົ່ງ string ນີ້ເພື່ອພິສູດວ່າເຄີຍ login ແລ້ວ

```http
POST /users HTTP/1.1
Host: localhost:8080
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

{
  "name": "Dee",
  "email": "dee@example.com"
}
```

ໃນ Go ຝັ່ງ server, ດຶງ header ດ້ວຍ:

```go
contentType := r.Header.Get("Content-Type")
authHeader  := r.Header.Get("Authorization")
```

---

## 7. Middleware (ຕົວດັກຈັບ Request)

ກ່ອນ request ຈະໄປຮອດ business logic, ມັນຈະຜ່ານ **middleware chain** ກ່ອນ — ຄືກັບລົດທີ່ຜ່ານດ່ານກວດຄວາມປອດໄພທີ່ລະດ່ານ.

### Logger Middleware

ບັນທຶກເວລາທີ່ request ເຂົ້າມາ (`time.Now()`), ສົ່ງຕໍ່ໃຫ້ chain ຕໍ່ໄປ, ແລ້ວຄຳນວນເວລາທີ່ໃຊ້ (microsecond) ຫຼັງຈາກ response ສຳເລັດ.

```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)  // ສົ່ງຕໍ່ໃຫ້ chain ຕໍ່ໄປ
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    })
}
```

### RequireAuth Middleware (ຍາມເຝົ້າປະຕູ)

ຂັ້ນຕອນການເຮັດວຽກ:

1. ດຶງ `Authorization` header
2. ຕັດ prefix `"Bearer "` ອອກ
3. parse token string
4. ໃຊ້ secret key ຈາກ `.env` ກວດສອບ signature (cryptographic math)
5. ຖ້າ signature ຖືກປ່ຽນແປງ → ຢຸດທັນທີ (`return`) ແລະຕອບ `401 Unauthorized` **ໂດຍບໍ່ຕ້ອງເຂົ້າ DB ເລີຍ** 

```go
func RequireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return  // ຢຸດທີ່ນີ້ — ບໍ່ສົ່ງຕໍ່
        }

        next.ServeHTTP(w, r)
    })
}
```

---

## 8. Router (ServeMux Multiplexer)

ຫຼັງ request ຜ່ານ global middleware ໝົດແລ້ວ, ມັນຈະຮອດ `http.NewServeMux()`. ມັນເຮັດໜ້າທີ່ເປັນ **ຕຳຫຼວດຈະລາຈອນ** ໂດຍແບ່ງເສັ້ນທາງຕາມ URL path.

| Path | Handler | ການປ້ອງກັນ |
|------|---------|-----------|
| `/login` | `LoginHandler` | public (ໃຊ້ອອກ token) |
| `/users` | `UserHandler` | secured (ຕ້ອງມີ token) |

```go
mux := http.NewServeMux()

// public route — ໃຜກໍ່ເຂົ້າໄດ້
mux.HandleFunc("/login", LoginHandler)

// protected route — ຫຸ້ມດ້ວຍ RequireAuth ກ່ອນ
mux.Handle("/users", RequireAuth(http.HandlerFunc(UserHandler)))

// ຫຸ້ມທັງ mux ດ້ວຍ Logger ເພື່ອບັນທຶກທຸກ request
http.ListenAndServe(":8080", Logger(mux))
```

ໂຄງສ້າງລຳດັບການເຮັດວຽກ:

```
Client → Logger → RequireAuth → ServeMux → UserHandler → DB
```

---

## 9. Database & SQL Joins (Relational Engine)

ໃນ `SQLStore`, ກ້າວຂ້າມການ query ແບບ single-table ໄປສູ່ສະຖາປັດຕະຍະກຳແບບ relational ແທ້ໆ.

### Foreign Keys

ເຊື່ອມຕາຕະລາງ `profiles` ກັບ `users` ໂດຍໃຊ້ `REFERENCES users(id)`. ນີ້ສ້າງ "ສະພານ" ລະຫວ່າງສອງຕາຕະລາງ.

```sql
CREATE TABLE users (
    id    SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
);

CREATE TABLE profiles (
    id      SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),   -- ສະພານເຊື່ອມ
    bio     TEXT
);
```

### INNER JOIN

ຮວມຂໍ້ມູນຈາກສອງຕາຕະລາງຜ່ານ key. **ຕ້ອງມີ record ຢູ່ທັງສອງຝັ່ງ** ຈິ່ງຈະປະກົດໃນຜົນລັບ — ຖ້າຝັ່ງໃດຝັ່ງໜຶ່ງບໍ່ມີ, record ນັ້ນຈະຖືກຕັດອອກຈາກ `UserProfileDTO`.

```go
query := `
    SELECT u.id, u.email, p.bio
    FROM users u
    INNER JOIN profiles p ON u.id = p.user_id
    WHERE u.id = $1
`

var dto UserProfileDTO
err := db.QueryRow(query, userID).Scan(&dto.ID, &dto.Email, &dto.Bio)
```

### ປະເພດ JOIN ທີ່ຄວນຮູ້

| JOIN | ພຶດຕິກຳ |
|------|--------|
| `INNER JOIN` | ເອົາສະເພາະ record ທີ່ມີຢູ່ທັງສອງຝັ່ງ |
| `LEFT JOIN` | ເອົາທຸກ record ຈາກຝັ່ງຊ້າຍ, NULL ຖ້າຝັ່ງຂວາບໍ່ມີ |
| `RIGHT JOIN` | ກົງກັນຂ້າມກັບ LEFT JOIN |


---

##  ການເລີ່ມຕົ້ນໃຊ້ງານ

```bash
# clone repo
git clone <repo-url>
cd go-modular-onboarding

# run
go run main.go

# build
go build -o app
./app
```
