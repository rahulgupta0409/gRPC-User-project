------gRPC Go Application README----
This is a gRPC-based Go application for a client-server architecture. It provides two methods: GetUserById and GetUserListByIds for retrieving user data.

Prerequisites
Go (1.17 or later)
Docker (for containerization)
Postman or any gRPC client tool for testing
Getting Started
Follow these steps to run and test the gRPC application.

1. Clone the Repository
bash
Copy code
git clone https://github.com/yourusername/your-repo.git
cd your-repo
2. Build the Docker Image
Build the Docker image for the application.

bash
Copy code
docker build -t your-image-name .
3. Run the Docker Container
Run a Docker container based on the image, mapping the gRPC service port (e.g., 50051) to a port on your host machine (e.g., 50051). Replace your-image-name with the image name.

bash
Copy code
docker run -p 50051:50051 your-image-name
4. Test the gRPC Service
Use a gRPC client (e.g., BloomRPC, grpcurl) or Postman to connect to the gRPC service.
Method 1: GetUserById
Request: Send a gRPC request to localhost:50051 with the GetUserById method. Provide a user_id to retrieve a specific user.

Sample Request (gRPCurl):

bash
Copy code
grpcurl -plaintext -d '{"user_id": 1}' -proto user.proto localhost:50051 userservice.UserService/GetUserById
Method 2: GetUsersByIds
Request: Send a gRPC request to localhost:50051 with the GetUsersByIds method. Provide an array of user_ids to retrieve multiple users.

Sample Request (gRPCurl):