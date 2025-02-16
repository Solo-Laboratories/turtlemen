group "default" {
    targets = ["auth-mock", "server", "web-app"]
}

target "auth-mock"{
    context = "./backend/auth-mock"
    dockerfile = "auth-mock.Dockerfile"
    args = {}
    tags = ["turtlemen/auth-mock:latest"]
}

target "web-app"{
    context = "./frontend/web"
    dockerfile = "web-app.Dockerfile"
    args = {}
    tags = ["turtlemen/web-app:latest"]
}

target "server"{
    context = "./backend/server"
    dockerfile = "backend.Dockerfile"
    args = {}
    tags = ["turtlemen/backend:latest"]
}