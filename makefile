export MONGO_URI=mongodb://root:pass@localhost:27017/?authSource=admin&w=majority
export MONGO_DB_NAME = user


run-debug:
	go run main.go