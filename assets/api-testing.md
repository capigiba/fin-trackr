
# API Testing Commands

## 1. Register User
```bash
curl -X POST http://localhost:8080/users/register \
-H "Content-Type: application/json" \
-d '{
    "user_name": "newuser",
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "password": "password123"
}'
```

## 2. Login User
```bash
curl -X POST http://localhost:8080/users/login \
-H "Content-Type: application/json" \
-d '{
    "user_name": "newuser",
    "password": "password123"
}'
```

## 3. Get User Information
Replace `TOKEN_HERE` with the token obtained from the login request and `USER_ID` with the user's ID.
```bash
curl -X GET http://localhost:8080/users/USER_ID \
-H "Authorization: Bearer TOKEN_HERE"
```

## 4. Update User Information
Replace `USER_ID` and `TOKEN_HERE` as needed.
```bash
curl -X PUT http://localhost:8080/users/USER_ID \
-H "Authorization: Bearer TOKEN_HERE" \
-H "Content-Type: application/json" \
-d '{
    "user_name": "newuser",
    "first_name": "Jane",
    "last_name": "Doe",
    "email": "jane.doe@example.com",
    "password": "newpassword123",
    "status": "active"
}'
```

## 5. Delete User
```bash
curl -X DELETE http://localhost:8080/users/USER_ID \
-H "Authorization: Bearer TOKEN_HERE"
```

## 6. Post Transaction
Replace `TOKEN_HERE` with the appropriate token.
```bash
curl -X POST http://localhost:8080/transaction \
-H "Authorization: Bearer TOKEN_HERE" \
-H "Content-Type: application/json" \
-d '{
    "user_id": 1,
    "year": 2024,
    "month": 9,
    "timestamp": "2024-09-20T12:00:00Z",
    "description": "Weekly groceries",
    "amount": "100.00"
}'
```

## 7. Get Transactions for a Specific Month and Year
```bash
curl -X GET http://localhost:8080/transaction/2024/09 \
-H "Authorization: Bearer TOKEN_HERE"
```