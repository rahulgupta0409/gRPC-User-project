# gRPC Go Application README
This Go application is built on the power of gRPC to enable a client-server architecture. It offers two essential methods: **GetUserById** and **GetUserListByIds** for efficient retrieval of user data. What makes this application unique is the utilization of bidirectional streaming, allowing real-time communication between clients and the server.



## Prerequisites
* Go (1.20 or later)
* Docker (for containerization)
* Postman or any gRPC client tool for testing

## Quick Start

You can get this application up and running with just one command, assuming you have Docker installed.
```bash
docker run -it -p 3001:3001 rahul0409/totality-assignment
```

## Getting Started
Follow these steps to run and test the gRPC application.

### 1. Clone the Repository
Clone the repository using the command
~~~ 
git clone https://github.com/yourusername/your-repo.git 
~~~
```
cd totality-assignment
```
### 2. Build the Docker Image
Build the Docker image for the application.

```bash
docker build -t totality-assignment .
```
### 3. Run the Docker Container
Run a Docker container based on the image, mapping the gRPC service port (e.g., 3001) to a port on your host machine (e.g., 3001).

```bash
docker run -p 3001:3001 totality-assignment 
```
### 4. Test the gRPC Service
Use a gRPC client (e.g., BloomRPC, grpcurl) or Postman to connect to the gRPC service.
* Method 1: GetUserById
Request: Send a gRPC request to localhost:3001 with the GetUserById method. Provide a User_Id to retrieve a specific user.

#### Sample Request
```
{
    "User_Id" : 1
}
```

* Method 2: GetUserListByIds
Request: Send a gRPC request to localhost:3001 with the GetUserListByIds method. Provide an array of User_Id to retrieve multiple users.

#### Sample Request
```
{
   "UserRequestList": [{"UserId":1}, {"UserId":3}, {"UserId":2},  {"UserId":18}]
}
```

## License
This project is licensed under the MIT License.

## Acknowledgments
Special thanks to the gRPC community for their support and resources.