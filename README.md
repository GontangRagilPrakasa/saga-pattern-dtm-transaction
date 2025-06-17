# 🌀 DTM Saga Pattern with Golang (HTTP + Docker)

Implementasi distributed transaction menggunakan [DTM (Distributed Transaction Manager)](https://github.com/dtm-labs/dtm) dengan pola **SAGA**, HTTP protocol, dan service A & B di dalam container Docker.

---

## 🧱 Arsitektur

Client
|
v
+-------------------+
| DTM HTTP Server |
| Port: 36789 |
+-------------------+
| |
v v
Service A Service B
:8081 :8082


---

## 🚀 Getting Started

### 1. 🔧 Requirements
- Docker & Docker Compose
- Golang ≥ 1.18
- Port: `8081`, `8082`, `36789`

### 2. 🛠️ Build & Run

```bash
# Build semua service
docker compose build

# Jalankan seluruh stack (DTM + services)
docker compose up
📦 Service Detail

🅰️ Service A
/try: Dummy sukses
/cancel: Logging rollback
http.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
	log.Println("[Service A] HIT /try")
	w.WriteHeader(http.StatusOK)
})

http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
	log.Println("[Service A] HIT /cancel (compensate)")
	w.WriteHeader(http.StatusOK)
})
🅱️ Service B (Simulasi Error)
http.HandleFunc("/try", func(w http.ResponseWriter, r *http.Request) {
	log.Println("[Service B] HIT /try")

	var payload struct {
		Amount int `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Failed to decode JSON:", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Received amount: %d\n", payload.Amount)

	// Trigger error jika amount > 500
	if payload.Amount > 500 {
		log.Println("[Service B] HIT /try FAILED")
		http.Error(w, "simulated failure", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
})

http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
	log.Println("[Service B] HIT /cancel (compensate)")
	w.WriteHeader(http.StatusOK)
})
🧪 Submit Saga Request

sagaReq := SagaRequest{
	GID:       "saga-golang-demo-001",
	TransType: "saga",
	Steps: []SagaStep{
		{Action: "http://service-a:8081/try", Compensate: "http://service-a:8081/cancel"},
		{Action: "http://service-b:8082/try", Compensate: "http://service-b:8082/cancel"},
	},
	Payloads: []interface{}{
		map[string]interface{}{"amount": 100},
		map[string]interface{}{"amount": 600}, // Trigger failure here
	},
}

data, _ := json.Marshal(sagaReq)
http.Post("http://localhost:36789/api/dtmsvr/submit", "application/json", bytes.NewReader(data))
🐞 Troubleshooting

Masalah	Solusi
Cancel tidak dipanggil	Pastikan branch gagal memberikan HTTP 5xx, bukan 200
Status stuck di submitted	Tunggu cron retry dari DTM atau trigger manual
Cron timezone	DTM pakai UTC, konversi waktu manual ke Asia/Jakarta
Trigger manual	Gunakan API /api/dtmsvr/cron atau restart container DTM
Error: current status 'succeed'	Jangan submit dua kali dengan GID sama
🔐 HTTPS Support (Optional)

Jika ingin menggunakan HTTPS:

Buat reverse proxy (Nginx, Caddy) untuk DTM dan service
Pastikan semua endpoint (action, compensate) pakai https://...
📂 .gitignore (rekomendasi)

# Golang
*.exe
*.test
*.out
bin/
build/

# OS / IDE
.DS_Store
.idea/
.vscode/

# Testing
*.cover
*.log

# Env
.env
🧾 Referensi

DTM Docs
DTM HTTP Saga Tutorial
SAGA Pattern Explained
👨‍💻 Author

Made with ❤️ by Gontang Prakasa


---

Kalau kamu ingin saya bantu generate struktur folder dan `docker-compose.yml` lengkap juga, tinggal bilang saja ya.