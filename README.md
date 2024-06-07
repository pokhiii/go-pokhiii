# go-pokhiii

go-pokhiii is an experimental eCommerce API built with Go and the Gin framework.

## Pre-requisites

Make sure you have the following installed on your machine:

1. **Go**: You can download and install Go from [the official Go website](https://golang.org/dl/).
2. **Git**: Ensure you have Git installed for version control.

## Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/your-username/go-pokhiii.git
    cd go-pokhiii
    ```

2. **Initialize Go Modules**:
    Ensure you are in the project directory and initialize the Go module:
    ```sh
    go mod tidy
    ```

## Running the Project

To run the project on your local machine, follow these steps:

1. **Run the Application**:
    ```sh
    go run main.go
    ```

2. The server will start on port `8080`. You can access the API endpoints using a tool like Postman or curl.

## API Endpoints

### List Products

- **URL**: `/products`
- **Method**: `GET`
- **Description**: Returns a list of all products.
- **Sample Request**:
    ```sh
    curl -X GET http://localhost:8080/products
    ```
- **Sample Response**:
    ```json
    [
        {
            "id": 1,
            "name": "Hoodie with Pocket",
            "description": "Pellentesque habitant morbi tristique...",
            "price": 45,
            "slug": "hoodie-with-pocket"
        },
        {
            "id": 2,
            "name": "Hoodie with Zipper",
            "description": "Pellentesque habitant morbi tristique...",
            "price": 45,
            "slug": "hoodie-with-zipper"
        },
        {
            "id": 3,
            "name": "Long Sleeve Tee",
            "description": "Pellentesque habitant morbi tristique...",
            "price": 25,
            "slug": "long-sleeve-tee"
        }
    ]
    ```

### Get Product Details

- **URL**: `/product/:product-slug`
- **Method**: `GET`
- **Description**: Returns details of a specific product by its slug.
- **Sample Request**:
    ```sh
    curl -X GET http://localhost:8080/product/hoodie-with-zipper
    ```
- **Sample Response**:
    ```json
    {
        "id": 2,
        "name": "Hoodie with Zipper",
        "description": "Pellentesque habitant morbi tristique...",
        "price": 45,
        "slug": "hoodie-with-zipper"
    }
    ```

## Custom 404 Handler

If a route is not found, the server will return a JSON response with a 404 status code and a message.

- **Sample 404 Response**:
    ```json
    {
        "message": "Page not found"
    }
    ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gin Web Framework](https://github.com/gin-gonic/gin)
