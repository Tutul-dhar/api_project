
# Docker

## 🐳 What is Docker?

Docker is an open-source platform designed to simplify application development, deployment, and operation using containers. Containers enable fast, consistent, and portable application delivery.

## 🚀 Why Docker?

- Isolated environments: Apps run in containers independently from host system.

- Lightweight & fast: Containers use fewer resources than traditional VMs.

- Portable & consistent: Run anywhere — local, on-prem, or cloud.

- CI/CD friendly: Ideal for modern DevOps workflows.

- Simplified scaling: Start, stop, and scale containers easily.

## 🧱 Docker Key Components

### 🔧 Docker Daemon (dockerd)
- Runs in background.

- Manages images, containers, volumes, and networks.

### 💬 Docker Client (docker)
- CLI tool to interact with the daemon:

        docker build ...

        docker run ...

        docker push ...

### 🧊 Docker Images
- Blueprint for containers.

- Created using Dockerfile.

- Read-only + layered.

#### 📁 Example: A Python image + Flask app configuration

### 📦 Docker Containers
- Running instance of an image.

- Isolated, but configurable (network, storage, etc).

        docker run -it ubuntu /bin/bash

### 📚 Docker Registries
- Store and share Docker images.

- Default: Docker Hub

        docker pull ubuntu
        docker push your-image-name

### 🧰 Example Docker Command Explained

    docker run -it ubuntu /bin/bash

### Breakdown:
- docker run: Start a container.

- -it: Interactive terminal mode.

- ubuntu: Image to use (pulled if not found locally).

- /bin/bash: Command to run inside container.

### What happens:
- Pulls image (if needed).

- Creates container.

- Allocates filesystem + networking.

- Starts container with terminal.

- After `exit`, container stops (but not deleted).

## 🖥️ Docker Desktop: Your All-in-One Container Platform

Docker Desktop is an easy-to-use application that lets you build, ship, and run containers right from your desktop — available for macOS, Windows, and Linux.

### 🧪 Run Your First Docker Container
Open your terminal and run:

    docker run -d -p 8080:80 docker/welcome-to-docker

- -d: Run in detached (background) mode.

- -p 8080:80: Map host port 8080 to container port 80.

- docker/welcome-to-docker: Sample image with a web frontend.

### 🧰 Benefits of Docker Desktop
- Simplifies container lifecycle management.

- Works across Mac, Windows, and Linux.

- Solves "it works on my machine" issues.

- Built-in tools like Docker Compose, Kubernetes, and image signing.

## 📤 Docker: Build and Push Image to Docker Hub
This section explains how to build a container image for your application and share it publicly (or privately) using Docker Hub.

### 📦 What Is a Container Image?
A container image is a lightweight, standalone package that includes:

- App code

- Runtime (e.g., Node.js)

- Libraries & environment

- Configuration & dependencies

Any machine with Docker can run it — no additional setup needed.

### 🏷️ Docker Hub
Docker Hub is the default container registry. You can:

- Pull trusted base images (node, mysql, nginx, etc.)

- Push your own images

- Make images public or private

### 🔐 Step 1: Sign In to Docker Hub
Open Docker Desktop and click "Sign in" at the top-right.

    Create an account if you don't have one yet.

### 📚 Step 2: Create a Repository
- Go to hub.docker.com

- Click Create Repository

- Fill in:

    - Name: getting-started-todo-app

    - Description: (optional)

    - Visibility: Public or Private

### 🛠️ Step 3: Build and Push the Image

#### 📁 Clone the Project (if not already)

    git clone https://github.com/docker/getting-started-todo-app
    cd getting-started-todo-app

#### 🏗️ Build the Image
Replace DOCKER_USERNAME with your actual Docker Hub username:

    docker build -t DOCKER_USERNAME/getting-started-todo-app .

#### 📋 Verify the Image Was Built

    docker image ls

Output sample:

    REPOSITORY                          TAG       IMAGE ID       CREATED          SIZE
    tutul/getting-started-todo-app     latest    abc1234efg56    2 minutes ago    1.12GB

#### 📤 Push the Image to Docker Hub

    docker push DOCKER_USERNAME/getting-started-todo-app

## 💡 Why Use Docker Compose?
When your application needs multiple services (e.g. backend, database, cache), managing each with docker run gets complex. Docker Compose simplifies this:

- Define all services in one file: compose.yaml

- Run with a single command

- Network, volume, image build — all handled automatically

### Best Practice:
Each container should do one thing well. Avoid bloated containers with multiple responsibilities.

### ⚙️ Key Difference

- Dockerfile	- Builds an image (what to package)
- compose.yaml	- Defines services/containers (what to run)

A compose.yaml file can reference Dockerfiles to build images on the fly.

### ▶️ 1. Start the App
    docker compose up -d --build

#### What this command does:

- Builds app image (if needed)

- Pulls MySQL image

- Creates network

- Creates persistent volume for DB

- Starts both containers

### 🛑 2. Stop the App

    docker compose down

Removes:

- Containers

- Network
### 🗑️ To remove volumes (DB data):
    docker compose down --volumes

## 🔍 What Are Image Layers?
Docker images are built in layers, each representing a set of filesystem changes (files added, modified, or deleted).

Typical example:

- Base OS (e.g., Ubuntu)

- Runtime (e.g., Python or Node.js)

- Application dependencies (e.g., requirements.txt)

- App source code

#### ✅ Benefits:

- Layer reuse between images speeds up builds and saves storage.

- Each layer is immutable, ensuring consistency.

- Uses a union filesystem to stack layers into a complete view for the container.

## 📄 What is a Dockerfile?
A Dockerfile is a text-based script that defines how to build a Docker image. It contains instructions to:

- Set base image

- Install dependencies

- Copy files

- Configure the container (e.g., default command)

        `STAGE 1: Build the Go binary`

        FROM golang:1.22 AS builder

        WORKDIR /app

        COPY go.mod go.sum ./

        RUN go mod download

        COPY . .

        RUN go build -o main .

        `STAGE 2: Run the binary in minimal image`

        FROM debian:bookworm-slim

        WORKDIR /app

        COPY --from=builder /app/main .

        EXPOSE 8080

        CMD ["./main"]

### 🧱 Common Dockerfile Instructions


| Instruction         | Description                                                |
| ------------------- | ---------------------------------------------------------- |
| `FROM <image>`      | Sets the **base image**                                    |
| `WORKDIR <path>`    | Sets the **working directory** in the image                |
| `COPY <src> <dest>` | **Copies files** from local host to image                  |
| `RUN <command>`     | **Runs a command** while building the image                |
| `ENV <key> <val>`   | Sets an **environment variable**                           |
| `EXPOSE <port>`     | Informs Docker that container will **listen on this port** |
| `USER <user>`       | Sets the **user** the container will run as                |
| `CMD ["command"]`   | Sets the **default command** to run on container start     |


### 🏗 Build the image

    docker build -t my-node-app .

### 🚀 Run the container

    docker run -p 3000:3000 my-node-app

## 🏗️ Build the Docker Image
Build an image from the Dockerfile in the current directory:

    docker build -t YOUR_DOCKER_USERNAME/image-name .

## 🏷️ Tag the Image (Optional)
If you want to add a version tag (like v1):

    docker image tag YOUR_DOCKER_USERNAME/image-name YOUR_DOCKER_USERNAME/image-name:v1


## 📦 List and Inspect Your Image
View your built image:

    docker image ls

Check its layer history:


    docker image history YOUR_DOCKER_USERNAME/image-name

## ☁️ Push the Image to Docker Hub

### 1. Login to Docker Hub

    docker login

Enter your Docker Hub credentials when prompted.

### 2. Push the Image

    docker push YOUR_DOCKER_USERNAME/image-name

Push with version tag (if tagged):

    docker push YOUR_DOCKER_USERNAME/image-name:v1

## 🧪 Run the Image

To verify it works:

    docker run YOUR_DOCKER_USERNAME/image-name

Or with tag:

    docker run YOUR_DOCKER_USERNAME/image-name:v1

## 🧹 Cleanup (Optional)
Remove local images to free space:

    docker image rm YOUR_DOCKER_USERNAME/image-name

## 🐳 Dockerfile (Multi-Stage)

        `STAGE 1: Build the Go binary`

        FROM golang:1.22 AS builder

        WORKDIR /app

        COPY go.mod go.sum ./

        RUN go mod download

        COPY . .

        RUN go build -o main .

        `STAGE 2: Run the binary in minimal image`

        FROM debian:bookworm-slim

        WORKDIR /app

        COPY --from=builder /app/main .

        EXPOSE 8080

        CMD ["./main"]

### Build the Docker Image

    docker build -t go-multistage-app .

### Run the Docker Container

    docker run -p 8080:8080 go-multistage-app

### Test It (in a new terminal)

    curl http://localhost:8080

Make sure your Go program listens on port 8080, e.g.:

    http.ListenAndServe(":8080", nil)

### View Image Info

    docker images

### Clean Up

    docker ps         # get CONTAINER ID
    docker rm -f <container_id>

Remove image:

    docker rmi go-multistage-app

## 🔄 What is Port Publishing?
Port publishing is the process of forwarding traffic from your host machine to a container’s internal port.

    docker run -d -p HOST_PORT:CONTAINER_PORT image-name

Example:

    docker run -d -p 8080:80 nginx

- 8080: host machine port

- 80: container port

### 📦 Use Case: Welcome to Docker App

▶️ Run Using Docker CLI

    docker run -d -p 8080:80 docker/welcome-to-docker

View Running Container

    docker ps

Example output:

    CONTAINER ID   IMAGE                   PORTS                    NAMES
    a527355c9c53   docker/welcome-to-docker 0.0.0.0:8080->80/tcp     clever_morse

### 🌐 Publishing to Ephemeral Ports

Let Docker assign a random port on the host:

    docker run -p 80 nginx

Check assigned port:

    docker ps

### 🔄 Publish All EXPOSED Ports

If the image has EXPOSE 80, you can publish all exposed ports:

    docker run -P nginx

Docker maps container ports to random host ports.

### ⚙️ Run Using Docker Compose

1. Create a directory and add this compose.yaml file:

        services:
            app:
                image: docker/welcome-to-docker
                ports:
                    - 8080:80

2. Start the application

    docker compose up

3. Open in browser

    http://localhost:8080


## 🐳 Docker: Overriding Container Defaults

When a Docker container starts, it runs an application or command based on its image configuration. You can override defaults like ports, environment variables, resource limits, networks, and startup commands using docker run flags or Docker Compose settings.

### 📦 Run Multiple Instances of Postgres

#### Instance 1 (Port 5432):

    docker run -d -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres

#### Instance 2 (Port 5433):

    docker run -d -e POSTGRES_PASSWORD=secret -p 5433:5432 postgres

Each container listens on port 5432 inside the container, but maps to different ports on the host.

### 🌍 Connect to a Custom Network

#### 1. Create custom network:
    docker network create mynetwork

#### 2. Run container in that network:

    docker run -d -e POSTGRES_PASSWORD=secret -p 5434:5432 --network mynetwork postgres

#### 3. Inspect network:

    docker network inspect mynetwork

### 🔑 Key Differences:

| Feature             | Default Bridge     | Custom Network            |
| ------------------- | ------------------ | ------------------------- |
| DNS Name Resolution | ❌ No               | ✅ Yes (by container name) |
| Isolation           | ❌ All share bridge | ✅ Scoped to network       |

### ⚙️ Set Environment Variables

Using -e:

    docker run -e foo=bar postgres env
Output:

    foo=bar

#### Using .env file:
Create .env:

    POSTGRES_PASSWORD=secret

Run:

    docker run --env-file .env postgres env

### 🔒 Manage Resources
Limit container's CPU and memory:

    docker run -d -e POSTGRES_PASSWORD=secret --memory="512m" --cpus=".5" postgres

Monitor usage:

    docker stats

### 🛠️ Override CMD and ENTRYPOINT
With Docker CLI:


    docker run -e POSTGRES_PASSWORD=secret postgres docker-entrypoint.sh -h localhost -p 5432

### 🐙 Docker Compose Overrides
Create a compose.yml:


    services:
        postgres:
            image: postgres
            entrypoint: ["docker-entrypoint.sh", "postgres"]
            command: ["-h", "localhost", "-p", "5432"]
            environment:
                POSTGRES_PASSWORD: secret

Bring up service:

    docker compose up -d

Enter container shell via Docker Dashboard or:

    docker exec -it <container_id_or_name> bash

Connect to Postgres inside:

    psql -U postgres

## 🗃️ Persisting Container Data with Docker Volumes

By default, data in a container is ephemeral—deleted once the container is removed. To persist data (like a database), use Docker volumes.


### 📦 What Are Volumes?
- Volumes store data outside of the container.

- You can:

    - Reuse data even after deleting containers.

    - Share data between containers.

    - Inspect and manage volume contents.

### 📌 Create and Mount a Volume

#### 1. Create a volume manually (optional):

    docker volume create log-data

#### 2. Start a container with the volume:

    docker run -d -p 80:80 -v log-data:/logs docker/welcome-to-docker

This mounts the volume log-data into the container's /logs directory.

### 🐘 Persisting Data in Postgres
#### 1. Run Postgres with a volume:

    docker run --name=db -e POSTGRES_PASSWORD=secret -d -v postgres_data:/var/lib/postgresql/data postgres

#### 2. Connect to Postgres:

    docker exec -ti db psql -U postgres

#### 3. Inside PostgreSQL shell:

    CREATE TABLE tasks (
        id SERIAL PRIMARY KEY,
        description VARCHAR(100)
    );
    INSERT INTO tasks (description) VALUES ('Finish work'), ('Have fun');
    SELECT * FROM tasks;

#### 4. Output:

    id | description
    ----+-------------
    1 | Finish work
    2 | Have fun
    (2 rows)

#### 5. Exit shell:

    \q

### 🔁 Restart Container, Keep Data

#### 1. Stop and remove the container:

    docker stop db
    docker rm db

#### 2. Start new container using same volume:

    docker run --name=new-db -d -v postgres_data:/var/lib/postgresql/data postgres

#### 3. Verify data persists:

    docker exec -ti new-db psql -U postgres -c "SELECT * FROM tasks"

### 🧹 Manage Volumes
#### List all volumes:

    docker volume ls

#### Remove specific volume:

    docker rm -f new-db   # Make sure no container is using the volume
    docker volume rm postgres_data

#### Remove all unused volumes:

    docker volume prune






