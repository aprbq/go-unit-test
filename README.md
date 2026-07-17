# Go Testing 

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v3-00ADD8?logo=go&logoColor=white)
![testify](https://img.shields.io/badge/testify-1.11-00ADD8?logo=go&logoColor=white)

## โครงสร้างโปรเจกต์
- `repositories/` — ชั้นติดต่อข้อมูล (`PromotionRepository`) พร้อม mock
- `services/` — business logic (`CalculateDiscount`) รับ repository เข้ามาแบบ dependency injection
- `handlers/` — HTTP handler ด้วย Fiber v3 รับ service เข้ามา

## Build tags แยกชนิดเทส
ใช้ build tag บนบรรทัดแรกของไฟล์เพื่อแยกว่าเป็น unit หรือ integration test
- `//go:build unit` — unit test ที่ mock service/repository ทั้งหมด 
- `//go:build integration` — integration test ที่ต่อ service จริงเข้ากับ repository mock

## คำสั่งรันเทส
```bash
# รันเฉพาะ unit test 
go test gotest/handlers -v 

# รันเฉพาะ unit test ถ้าไฟล์เทสมี //go:build unit อยู่บนสุด
go test gotest/handlers -v -tags=unit

# รันเฉพาะ integration test
go test gotest/handlers -v -tags=integration

# รันเฉพาะเจาะจงเทสเคส
go test gotest/services -run="TestcheckGrade/success grade a"

```

## คำสั่งผ่าน Makefile
| คำสั่ง | ทำอะไร |
|---|---|
| `make run` | รันโปรแกรมหลัก (`go run .`) |
| `make test` | รัน unit test ทั้งหมด |
| `make test-unit` | รันเฉพาะ unit test ของ handlers |
| `make test-integration` | รันเฉพาะ integration test (`-tags=integration`) |
| `make test-all` | รันเทสทุกชนิดรวม integration |
| `make bench` | รัน benchmark พร้อม memory stats |
| `make cover` | รันเทสพร้อมสรุป coverage รายฟังก์ชัน |
| `make tidy` | จัดระเบียบ dependency (`go mod tidy`) |

## เทคนิคที่ใช้
- **Mock** — ใช้ `stretchr/testify/mock` จำลอง dependency แล้วตั้งค่าด้วย `.On(...).Return(...)` และตรวจด้วย `AssertNotCalled`
- **Table-driven test** — วนเคสหลายกรณีใน `CheckGrade` เพื่อลดโค้ดซ้ำ
- **Benchmark & Example test** — วัดประสิทธิภาพและเป็นเอกสารตัวอย่างการใช้งานไปในตัว

---

## ขยายความแนวคิดพื้นฐาน

### 1. Whitebox vs Blackbox testing
ดูจาก package ของไฟล์เทสว่าอยู่ที่ไหน
- **Whitebox** — ไฟล์เทสอยู่ใน package เดียวกับโค้ดที่เทส (เช่น `package services`) จึงมองเห็นและเรียกใช้ตัวแปร/ฟังก์ชันที่เป็น private (ตัวเล็ก) ได้ เหมาะกับการเทส logic ภายใน
- **Blackbox** — ไฟล์เทสอยู่คนละ package โดยเติม `_test` ต่อท้าย (เช่น `package handlers_test`) มองเห็นเฉพาะสิ่งที่ export (ตัวใหญ่) เท่านั้น ทำให้เทสผ่านมุมมองของผู้ใช้งานจริงเหมือนเรียกจากภายนอก
  > ในโปรเจกต์นี้ไฟล์ใน `handlers/` ใช้ `package handlers_test` จึงเป็น blackbox

### 2. ไฟล์ test ไม่มีผลต่อการรันปกติ
ไฟล์ที่ลงท้ายด้วย `_test.go` จะถูกคอมไพล์ **เฉพาะตอนรัน `go test` เท่านั้น** ไม่ถูกรวมเข้าไปตอน `go build` หรือ `go run` ปกติ ดังนั้นต่อให้เทสยังแดงอยู่ ก็ยังสามารถ build และรัน `go run main` ได้ตามปกติ เทสกับโค้ดโปรดักชันจึงแยกขาดจากกัน

### 3. Integration test
คือการเทสที่ให้หลายส่วนทำงาน **ต่อกันจริง** เช่น service เรียก repository ที่ต่อกับ database จริง เพื่อยืนยันว่าเมื่อประกอบร่างกันแล้วทำงานถูกต้อง
- ข้อดี: ใกล้เคียงการใช้งานจริง จับบั๊กที่เกิดจากการเชื่อมต่อระหว่างชั้นได้
- ข้อเสีย: ช้ากว่า ต้องเตรียม dependency ภายนอก และเปราะกว่า

### 4. Unit test ต้อง mock
unit test มุ่งทดสอบ **ฟังก์ชันเดียว/หน่วยเดียว** ให้แยกออกจาก dependency ภายนอก (database, API, service อื่น) เพื่อให้
- รันเร็วและได้ผลเหมือนเดิมทุกครั้ง (deterministic) ไม่ขึ้นกับสถานะภายนอก
- โฟกัสว่า logic ของหน่วยนั้นทำงานถูกต้องไหม

จึงใช้ **mock** สร้างของปลอมมาแทน dependency แล้วกำหนดพฤติกรรมเอง เพื่อบังคับผลลัพธ์ให้ตรงตามเคสที่ต้องการทดสอบ
