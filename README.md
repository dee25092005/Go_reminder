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

##  ໂຄງສ້າງໂປຣເຈັກ

```
go-modular-onboarding/
│
├── go.mod                 # ໄຟລ໌ root ຂອງ module
├── main.go                # entry point ຂອງໂປຣແກຣມ
│
├── storage/               # logic ກ່ຽວກັບ user
│   └── memory.go
│
└── calculator/            # logic ກ່ຽວກັບການຄຳນວນ
    └── math.go
```

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
