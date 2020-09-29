# realworld-exemple-go-restapi

# :+1: Go REST-API [BETA]
GO API template as REST.
## :rocket: Installation
>**launching**
```
go run main.go
```
## :round_pushpin: EndPoints
### **[GET]** /api/foo
_Foo Search services with filters_
>**INPUT [Query]**
>Exemple of an input URL
```
http://localhost:8080/api/foo?name=lise
```
>**OUTPUT [JSON]**
>Output structure example
```
{
    "Name": "List",
    "Age": 10,
    "City": "Lyon"
}
```

### **[GET]** /api/foo/{id}
_Foo Find services with ID_
>**INPUT [Query]**
>Exemple of an input URL
```
http://localhost:8080/api/foo/ID
```
>**OUTPUT [JSON]**
>Output structure example
```
{
    "Name": "List",
    "Age": 10,
    "City": "Lyon"
}
```

### **[POST]** /api/foo
_Foo Register services_
>**INPUT [JSON]**
>Exemple of an input URL
```
{
    "Name": "List",
    "Age": 10,
    "City": "Lyon"
}
```
>**OUTPUT [STRING]**
>Output structure example
```
"ID"
```

### **[DELETE]** /api/foo/{id}
_Foo Delete services_
>**OUTPUT [STRING]**
>Output structure example
```
"Ok"
```
