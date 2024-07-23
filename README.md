<div align="center">
  <br>
  <h1>Libry API</h1>
  <p>📖 A simple library REST API 📖</p>
  <br>
</div>

## Table of Contents

- [Description](#description)
- [Features](#features)
- [Entities](#entities)
- [Run Locally](#run-locally)

## Description

[`^ back to top ^`](#table-of-contents)

**Libry API** is a simple REST API that allows people to borrow books from a library. This project is made with Go. It is created as an exercise in the Go phase of the Backend Development Training.

## Features

[`^ back to top ^`](#table-of-contents)

- View list of users & their currently borrowed books.
- Add a new user.
- View list of books.
- Add a new book.
- View list of borrowed books.
- Borrow a book.
- Return a book.

## Entities

[`^ back to top ^`](#table-of-contents)

There are 3 entities: **Book**, **User**, & **Borrow**.

**Book**

- id: `string`
- title: `string`
- author: `string`

**User**

- id: `string`
- username: `string`
- books: `[]*Book`

**Borrow**

- id: `string`
- user_id: `string`
- book_id: `string`
- start_date: `timestamp`
- end_date: `timestamp`
- status: `string`

## Run Locally

[`^ back to top ^`](#table-of-contents)

- Make sure you have [Go](https://go.dev) installed on your computer.

- Clone the repo.

  ```bash
  git clone https://github.com/nadiannis/libry-api.git
  ```

  ```bash
  cd libry-api
  ```

- Install the dependencies.

  ```bash
  go mod tidy
  ```

- Run the server. The web server will run on port 8080.

  ```bash
  go run ./cmd
  ```
