# Docker
A Dockerfile has been provided to run this application.  The default port exposed is 8080.

# Environment Variables
The following environment variables are needed by creating a `.env` file and load it up with the environment variables.
|Variable|Purpose|example|
|---|---|---|
|`MONGODB_URI`|Address to mongo server|`mongodb://servername:27017` or `mongodb://username:password@hostname:port` or `mongodb+srv://` schema|
|`SECRET_KEY`|Secret key for JWT tokens|`secret123`|

Alternatively

It can be hardcoded at the below locations:
1. `MongoURI` inside database.go . delete or comment out the funtion `os.Getenv()` and the `os` lib import
```bash
MongoDbURI := "mongodb://myUserAdmin:abc123@10.0.0.2:27017"
```
2. `SECRET_KEY` inside auth.go . delete or comment out the funtion `os.Getenv()` and the `os` lib import
```bash
var SECRET_KEY string = "secret123"
```

# Running with Go

Clone the repository into a directory of your choice Run the command `go mod tidy` to download the necessary packages.

You'll need to add a .env file and add a MongoDB connection string with the name `MONGODB_URI` to access your collection for task and user storage.
You'll also need to add `SECRET_KEY` to the .env file for JWT Authentication.

Run the command `go run main.go` and the project should run on `locahost:8080`

# Git 
Working in the app folder

1. initialise local directory as a Git repository. By default, the initial branch is called main. you can set the name of the default branch using -b
```bash
git init
```
2. set user name 
```bash
git config --global user.name kkap-one
```
3. set user email
```bash
git config --global user.email kkap-one
```
4. Add the files in your new local repository. This stages them for the commit.
```bash
git add .
```
5. Commit the files that you've staged in your local repository. `-m` used to provide message describing the commit.
```bash
git commit -m "initial commit"
```
6. Add the URL for the remote repository where your local repository will be pushed
```bash
git remote add origin https://github.com/kkap-one/app-tasky-wiz.git
```

7. If an another origin exists, change origin to target remote repo by
```bash
git remote set-url origin https://github.com/kkap-one/app-tasky-wiz.git
```

8. Verfiy remote repo set using
```bash
git remote -v
```

8. To push the changes in your local repository to GitHub. -f option on git push forces the push
```bash
git push -f origin main
```

9. for update/recommit etc
```bash
git add .
git commit -m "comment"
git push
```

# License

This project is licensed under the terms of the MIT license.

Original project: https://github.com/dogukanozdemir/golang-todo-mongodb