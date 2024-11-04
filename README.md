# 📋 Project: Rumue - Insurance Price Check & Purchase Platform

This application, Rumue, is designed to provide users with an easy way to check insurance and vehicle registration prices, purchase policies, and monitor the status of their requests. It also includes features for notifications and receipt generation for user convenience.

---

## 🚀 Features

### 1. Price Checking for Insurance and Registration
- **Insurance Types**: Users can select insurance types such as First Class, Second Class, or Third Class.
- **Input Details**: Users enter car details like brand, model, year, and mileage.
- **Accurate Pricing**: The system calculates and displays a price based on car condition and insurance type, allowing users to quickly view relevant options.

### 2. Insurance or Registration Purchase
- **Seamless Purchase**: After checking the price, users can choose to proceed with the purchase.
- **Additional Car Information**: Users complete additional details and proceed with payment to finalize the purchase.

### 3. Request Status Tracking
- **Status Overview**: Users can view the current status of their insurance or registration requests.
- **Status Types**:
  - **Pending**: The purchase is under review, awaiting employee verification.
  - **Approved**: All details have been verified, and the process is completed successfully.
  - **Rejected**: The request has been denied due to incorrect information.

### 4. Expiry Notification
- **Renewal Reminders**: If a user has previously purchased insurance or registration with Rumue, an email reminder will be sent close to the expiry date to prompt renewal.

### 5. Receipt Generation
- **Automated Receipts**: Once an insurance or registration purchase is confirmed by an employee, a receipt is generated and sent to the user via email.
- **Employee Access**: Employees can download receipts if needed for record-keeping or customer inquiries.

---

## 🛠️ Getting Started

### Prerequisites
- Ensure Docker is installed and running on your system.
- Clone the repository and set up the environment variables.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/ThanawatPtd/SA-SEProject.git
   ```
2. Configure environment variables:
   - Refer to example.env and add your configurations.
3. Launch the application: 
    ```
    docker compose up -d
    ```
---

## ⚙️ Make Commands for Database Migrations
The following commands are available to manage database migrations:
Run migrations up:
```
make docker-migrate-up
```
Run migrations down:
```
make docker-migrate-down
```
Reset migrations:
```
make docker-migrate-reset  
```

---
## 👥 Developers

- ชุติพงษ์ ไตรยะสิทธิ์ 6510450291
- ศุภณัฐ สร้อยเพชร 6510450976
- ธนวัฒน์ โพธิเดช 6510450445
- เลิศพิพัฒน์ กาญจนเรืองโรจน์ 6510450917

