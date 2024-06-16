# Web Page Analyzer for Home24
_This repository contains **the solution to the Home24 coding challenge**_

![made-with-go](https://img.shields.io/badge/Made_with-Go-blue) ![MIT license](https://img.shields.io/badge/License-MIT-orange.svg) ![Maintenance](https://img.shields.io/badge/Maintained%5F-yes-green.svg)

### Description

|                  ![gif](img/wpanalyzer.gif)                  |
| :----------------------------------------------------------: |
| <span style="color:grey"> <i><b>Fig. 1</b>: The Web Page Analyzer in action</i></span> |

---
### Implementation

#### Tech Stack
<p>
<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" height="24" />
<img src="https://img.shields.io/badge/html5-%23E34F26.svg?&style=for-the-badge&logo=html5&logoColor=white" height="24" />
<img src="https://img.shields.io/badge/css3-%231572B6.svg?&style=for-the-badge&logo=css3&logoColor=white" height="24"/>
<img src="https://img.shields.io/badge/docker-%232496ED.svg?&style=for-the-badge&logo=docker&logoColor=white" height="24"/>
</p>


---
### How To Use This Code
#### 🔔&nbsp; Prerequisites

- [**Git**](https://www.git-scm.com/downloads)
- [**Go**](https://go.dev/dl/) (version 1.22 or higher)
- _Optional (but recommended)_: [**Docker**](https://www.docker.com/products/docker-desktop/)

First of all, **clone this repository** and **navigate inside the folder**:

```sh
git clone https://github.com/fra-mari/home24
cd home24
```    
 
Then, use the following instructions to build and start the application, either directly or, if you do not wish to install `Go`, using `Docker`. 
 
#### 📌&nbsp; Direct build, on UNIX systems

⚠️ **N.B.**: If you use a Windows system, or you prefer to use `Docker`, please follow the instructions [in the following paragraph](https://github.com/fra-mari/home24?tab=readme-ov-file#%EF%B8%8F--on-windows).

1. Ensure you have Go installed on your system. You can download it from the official [Go website](https://go.dev/dl/).

2. In the project directory, download the dependencies:
    
    ```sh
    go mod tidy
    ```
    
3. Build the application:     

    ```sh
    go build -o analyzer_build
    ```

	_Note_: The `-o` flag specifies the output file name. In this example, the compiled binary will be 	named `analyzer_build` and placed in the current directory.

4. Set the server to release mode:    
	
	```sh
    GIN_MODE=release
    ```
    
5. Start the application:
	
	```sh
    ./analyzer_build
    ```    
The application will be accessible at `http://localhost:8080`. To gracefully shut it down, you may press `Ctrl+C`.

#### 📌&nbsp; Docker build, on Windows (and on UNIX, too, if you wish)

1. Clone the repository:



2. Build the Docker image:

    ```sh
    docker build -t analyzer .
    ```

3. Run the Docker container:

    ```sh
    docker run -p 8080:8080 analyzer
    ```

The application will be accessible at `http://localhost:8080`. To gracefully shut it down, you may press `Ctrl+C`.

---
### Possible Improvements and To Dos
- [ ] Provide customers with a budget to spend into the supermarket.
- [ ] Add a tool for the user to visualise the movements of each customer on the supermarket map.
- [ ] Tests.
