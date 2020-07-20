export DEV_PORT=8080
export DB_URI=mongodb://localhost:27017
export DB_NAME=sirius
ACCESS_TOKEN_KEY=yoursecret

dev:
	go run main.go