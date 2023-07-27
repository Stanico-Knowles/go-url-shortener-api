<a name="readme-top"></a>

# URL Shortener API - Go Gin

## About The Project

This is a URL shortening API built with the Go Gin web framework and Gorm ORM. I decided to a modular approach and implement a [controller-service-repository](https://tom-collings.medium.com/controller-service-repository-16e29a4684e5) pattern. Although it is not conventional in Go, it came to down to preference. This project uses a MySQL database, however you can vist the [gorm documentation](https://gorm.io/docs/connecting_to_the_database.html) to see how to implement the database driver of your choice.

### Built With

<a href="https://gin-gonic.com/">
    <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="200" width="150"/>
</a>

## Getting Started

To run the project locally, follow these simple steps:

### Installation

1. Get the Docker Engine at [https://docs.docker.com/engine/](https://docs.docker.com/engine/) or Docker Desktop (recommended) for an easy-to-use GUI at [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/)
2. Clone the repo
   ```sh
   git clone https://github.com/Stanico-Knowles/go-url-shortener-api.git
   ```
3. Install Go packages
   ```sh
   go mod download
   ```
4. Build docker containers and run project
   ```js
   docker compose -f docker-compose.yml up -d
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

Project Link: [https://github.com/Stanico-Knowles/go-url-shortener-api](https://github.com/Stanico-Knowles/go-url-shortener-api)
