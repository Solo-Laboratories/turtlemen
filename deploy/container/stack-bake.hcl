group "default" {
    targets = ["authentication", "backend", "frontend"]
}

target "authentication"{
    context = "./server/authentication"
    dockerfile = "authenitcation.Dockerfile"
    args = {}
    tags = ["turtlemen/authenitcation:latest"]
}

target "frontend"{
    context = "./app/web"
    dockerfile = "frontend.Dockerfile"
    args = {}
    tags = ["turtlemen/frontend:latest"]
}

target "backend"{
    context = "./server/turtlemen-server"
    dockerfile = "backend.Dockerfile"
    args = {}
    tags = ["turtlemen/backend:latest"]
}