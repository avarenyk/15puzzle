## 15 puzzle game

This is implementation  of [15 puzzle](https://en.wikipedia.org/wiki/15_puzzle) game with Go language using Terminal as GUI

## Prerequisites

Before running the code, ensure you have the following installed on your machine:

- [Go](https://golang.org/dl/) (version 1.22.7 or later)
- Git (to clone the repository)

### Installing Go

1. **Download and Install Go:**

    - Visit the [Go Downloads page](https://golang.org/dl/).
    - Download the installer suitable for your operating system.
    - Follow the installation instructions provided on the page.

2. **Verify the Installation:**

   Open a terminal and run the following command to ensure Go is installed correctly:

   ```sh
   go version
   ```
   You should see output indicating the installed Go version.

## Install && Run 

1. **Clone Repository**
   ```sh
   git clone git@github.com:avarenyk/15puzzle.git
   ```
2. **Navigate to repo**
   ```sh
   cd 15puzzle
   ```
3. **Install dependencies** 
   ```sh
   go mod download
   ```
4. **Testing** 
   ```sh
   go test ./...
   ```
5. **Running application** 
   ```sh
   go run main.go
   ```

### TODO

1. We could leverage PuzzleGamer interface to use Proxy pattern to add logging, metrics, etc
2. Add web server and web interface to a game
3. Implement possibility to host multiple games for a different players on web