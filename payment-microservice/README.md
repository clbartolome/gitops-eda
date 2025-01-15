
# Payment Microservice

This project is a simple Go-based microservice that simulates a payment process. It features a modern, PayPal-inspired UI and includes loading and status effects. The web interface is hosted via Go's built-in HTTP server and supports simulated payment requests through an API.

---

## Project Structure

```
payment_microservice/
├── main.go               # Go code to serve the application and handle API
├── templates/
│   └── index.html        # Frontend HTML for the web page
├── static/
│   └── styles.css        # CSS styles for the frontend
├── go.mod                # Go module file
├── go.sum                # Dependencies
└── README.md             # This README file
```

---

## Prerequisites

- **Go**: Version 1.16+ installed on your machine.
- **Browser**: A modern browser to test the web application.

---

## Getting Started

1. **Clone the Repository**
   ```bash
   git clone https://github.com/your-repo/payment_microservice.git
   cd payment_microservice
   ```

2. **Initialize the Module**
   ```bash
   go mod init payment_microservice
   ```

3. **Run the Application**
   ```bash
   go run main.go
   ```

4. **Access the Web Interface**
   Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

---

## API Endpoints

### `POST /api/pay`
Simulates payment processing.

- **Request Body** (JSON):
  ```json
  {
      "amount": 100,
      "currency": "USD"
  }
  ```

- **Response** (JSON):
  - On Success:
    ```json
    {
        "status": "success",
        "message": "Payment processed successfully!",
        "amount": 100,
        "currency": "USD"
    }
    ```
  - On Error:
    ```json
    {
        "error": "Invalid amount format"
    }
    ```

---

## Create Image

```bash
# Create image
podman build -t payment-microservice:1.0.0 . 

# Run Image
podman run -p 8080:8080 payment-microservice:1.0.0 . 

# Test payment
curl -X POST http://localhost:8080/api/pay -H "Content-Type: application/json" -d '{"amount": 100, "currency": "USD"}'

# Tag and push into quay
podman tag payment-microservice:1.0.0 quay.io/calopezb/payment-microservice:1.0.0
podman push quay.io/calopezb/payment-microservice:1.0.0
```

---