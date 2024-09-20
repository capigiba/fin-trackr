# TODO

* [Vietnamese version](assets/vi-README.md)

## 1. Expense Tracker -----> [In process]
Description: Build a simple application to track personal expenses. Users can add, edit, and delete transactions, categorize them, and view basic reports (e.g., total expenses per month).

Key Concepts:
Working with CRUD operations and databases.
Understanding financial transactions and categorization.
Basic data visualization (optional).

Key Learning Areas:

Golang database handling (e.g., using SQLite or PostgreSQL).
Implementing REST APIs for financial transactions.
Security in handling personal financial data.

## 2. Currency Converter API
Description: Create a currency converter API that fetches real-time exchange rates from an external API and performs currency conversions for users.

Key Concepts:
Working with external financial APIs (e.g., OpenExchangeRates).
Data processing and calculations (handling floats/decimals).

Key Learning Areas:

Golang’s HTTP client for calling external APIs.
Implementing rate-limiting for APIs.
Converting currencies and managing currency exchange data.

## 3. Simple Payment Gateway
Description: Build a simple payment gateway that allows users to initiate payments (mock credit card details) and processes transactions, generating payment receipts.

Key Concepts:
Payment validation and processing.
Mock transactions (no real banking involved) to simulate how a gateway works.

Key Learning Areas:

Using JSON and HTTP for requests.
Understanding how real-world payment gateways operate.
Secure handling of payment data.

## 4. Budgeting Application
Description: Create a budgeting tool where users can set budgets for different categories (e.g., groceries, entertainment) and track their spending against those budgets.

Key Concepts:
Budget creation and tracking.
Handling different financial categories.

Key Learning Areas:

Implementing business logic around budgets.
Working with time-bound financial data (e.g., monthly budgets).
Golang’s struct and data validation features.

## 5. Microloan Simulation
Description: Build a basic loan simulation tool that calculates loan repayment schedules (including interest rates) for users based on loan amount, interest rate, and loan duration.

Key Concepts:
Understanding loan repayment schedules.
Interest rate calculations (compound interest, amortization schedules).

Key Learning Areas:

Handling financial calculations in Golang.
Creating and managing loan data (e.g., using SQL databases).
Simulating and visualizing loan payments over time.

## 6. Stock Portfolio Tracker
Description: A basic application that allows users to track the value of their stock portfolio, including real-time stock price updates from a public API.

Key Concepts:
Working with real-time financial data.
Portfolio management and calculating portfolio value.

Key Learning Areas:

Fetching stock prices from an external API (e.g., Alpha Vantage).
Working with time series data in Golang.
Real-time data handling and updates.

## 7. KYC (Know Your Customer) Validation System
Description: Build a simple KYC system where users can submit documents (e.g., ID, passport), and the system validates and stores their information.

Key Concepts:
Document validation and management.
Handling sensitive customer data securely.

Key Learning Areas:

File uploads and secure storage in Golang.
Data validation and authentication mechanisms.
Understanding KYC requirements for financial systems.

## 8. Cryptocurrency Price Tracker
Description: Create an application that fetches and displays cryptocurrency prices (e.g., Bitcoin, Ethereum) from a public API and shows trends over time.

Key Concepts:
Fetching real-time financial data.
Data visualization for price trends.

Key Learning Areas:

Golang’s HTTP package for API requests.
Understanding how cryptocurrency APIs work.
Time-series data processing and visualization.

## 9. Basic Invoice Generation System
Description: Develop a tool that generates invoices for services/products, including customer information, itemized costs, and total amount due.

Key Concepts:
Invoice creation and PDF generation.
Managing customer and product/service data.

Key Learning Areas:

PDF generation in Golang (e.g., using gofpdf).
Handling business logic around invoices and payment terms.
Database operations for storing invoice and customer data.

## 10. Transaction Reconciliation Tool
Description: Build a tool that compares bank statements with internal transaction records to identify discrepancies.

Key Concepts:
Bank statement parsing and comparison.
Reconciliation algorithms for matching transactions.

Key Learning Areas:

File parsing (CSV, XLSX, etc.) in Golang.
Algorithms for reconciling large datasets.
Automating financial checks and balances.