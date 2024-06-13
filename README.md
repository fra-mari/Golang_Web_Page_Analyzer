# Home24
This repository contains **the solution to the Home24 coding challenge**

## How to Run Using Docker

1. Clone the repository:

    ```sh
    git clone https://github.com/fra-mari/home24
    cd home24
    ```

2. Build the Docker image:

    ```sh
    docker build -t analyzer .
    ```

3. Run the Docker container:

    ```sh
    docker run -p 8080:8080 analyzer
    ```

The application will be accessible at `http://localhost:8080`. To gracefully shut it down, you may press Ctrl+C.
