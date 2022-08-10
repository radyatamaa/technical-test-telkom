# technical-test-telkom

This project has  4 Domain layer :

 * Models Layer
 * Repository Layer
 * Usecase Layer  
 * Delivery Layer

#### The diagram:

![golang clean architecture](https://github.com/bxcodec/go-clean-arch/raw/master/clean-arch.png)

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How To Run This Project

```bash
#move to directory
cd $GOPATH/src/github.com/bxcodec

# Clone into YOUR $GOPATH/src
git clone https://ilyas_bashirah@bitbucket.org/bridce/technical-test-telkom.git

#move to project
cd technical-test-telkom

# Test the code
make test

# Run Project
go run main.go


```

Or with `docker-compose`

```bash
#move to directory
cd $GOPATH/src/github.com/bxcodec

# Clone into YOUR $GOPATH/src
git clone https://ilyas_bashirah@bitbucket.org/bridce/technical-test-telkom.git

#move to project
cd technical-test-telkom

# Build the docker image first
make docker

# Run the application
make run

# check if the containers are running
docker ps

# Execute the call
curl localhost:5000/api/sample-module

# Stop
make stop
```
### Validation Unit Test Commit
if you want to apply unit test validation on commit copy file "pre-commit" to folder ".git/hooks/"      

### List Command
                    
Function  | Command
------------- | -------------
Update Swagger  | swag init -g main.go --output swagger
Run All Unit Test  | make test 
Docker Build  | make docker
Docker Run  | make run 
Docker Stop  | make stop
Run All Unit Test  | make test 
Build  | make engine

### Standard Status Code 
                    
Description  | Code
------------- | -------------
Success  | 200
Not Found  | 404 
Conflict  | 409
UnAuthorize  | 401 
Bad Param Input  | 400
Other  | 500 
