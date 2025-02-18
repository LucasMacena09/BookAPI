# BookAPI

BookAPI is an improved version of my previous API, **GoLibrary**. This version introduces significant enhancements, including switching from the Gin framework to the **Fiber** framework for better performance and simplicity. Additionally, unlike the previous version that used a Go slice (`[]book`) to store data, **BookAPI** now utilizes a **PostgreSQL database** for persistent storage.

## Features
- Retrieve all books
- Retrieve a book by its ID
- Add new books
- Delete books

## Project Structure
```
BookAPI/
├── main.go
├── models/
│   ├── books.go
├── storage/
│   ├── postgres.go
├── handlers/
│   ├── book_handler.go
├── routes/
│   ├── routes.go
```

## Running the API
### Prerequisites:
- Go installed (version 1.23.4 or later)
- Fiber framework
- PostgreSQL database configured

### Steps:
```sh
git clone https://github.com/LucasMacena09/BookAPI.git
cd BookAPI
go run main.go
```

## Endpoints
| Method | Endpoint         | Description           |
|--------|-----------------|-----------------------|
| GET    | /books          | Get all books        |
| GET    | /books/:id      | Get a book by ID     |
| POST   | /create_books   | Add a new book       |
| DELETE | /delete_book/:id | Delete a book       |

## Future Improvements
- Implement user authentication
- Add book update functionality
- Implement unit tests for better reliability

This project is a step forward in my Go backend development journey!

