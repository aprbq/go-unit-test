.PHONY: run test test-unit test-integration test-all bench cover tidy

# รันโปรแกรมหลัก (ตัวอย่าง mock ใน main.go)
run:
	go run .

# รัน unit test ทั้งหมด (ไฟล์ที่ไม่มี build tag)
test:
	go test ./... -v

# รันเฉพาะ unit test ของ handlers
test-unit:
	go test gotest/handlers -v

# รันเฉพาะ integration test (ไฟล์ที่มี //go:build integration)
test-integration:
	go test gotest/handlers -v -tags=integration

# รันเทสทุกชนิดรวม integration
test-all:
	go test ./... -v -tags=integration

# รัน benchmark ทั้งหมด
bench:
	go test ./... -bench=. -benchmem

# รันเทสพร้อมสรุป coverage
cover:
	go test ./... -tags=integration -coverprofile=coverage.out
	go tool cover -func=coverage.out

# จัดระเบียบ dependency
tidy:
	go mod tidy
