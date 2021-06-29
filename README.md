# inventory-api

![image](https://user-images.githubusercontent.com/18530378/123744251-1f00b480-d8cc-11eb-9a30-d8c44760f9dc.png)

**Prerequisites:**

golang env
POSTMAN
mySQL Workbench
(You will have to change username and password in main.go file)

**How to run?**

go build (it willl create one executable file)
.\inventory-api.exe

**POSTMAN guide:**


In file function initializeRoutes helps to route request. you can get request string from

for e.g shop.Router.HandleFunc("/Categorys", shop.getCategorys).Methods("GET")

the request in postman would be http://localhost:8080/Categorys type of request would be GET (common prefix : http://localhost:8080)
