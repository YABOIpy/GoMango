# GoMango
A WebApplication for simple Project Managment that runs on a Locally Hosted Server

<p align="center" style="text-align: center">  
  <img src="https://user-images.githubusercontent.com/110062350/231760953-8558058f-167c-49d7-93b5-3ad9663c1a64.png">
</p>



![image](https://github.com/YABOIpy/GoMango/assets/110062350/d8cdf99a-3cf9-4598-a1a3-2eb28de11038)


# Setup
![image](https://github.com/YABOIpy/GoMango/assets/110062350/eb41ee64-ebde-4ff9-86b7-6b0c9e25bf6e)
```scss
1. install go at https://go.dev
2. install xampp
3. download GoMango and unzip it
4. open cmd from the directory that gomango is in and type: "go run ."
5. after installing xampp, open the htdocs folder. open the gomango download and inside you will find the gomango folder. copy and paste it into htdocs
6. open your browser and go to: http://127.0.0.1:8080
7. after that open phpmyadmin from xampp by going to: http://127.0.0.1/phpmyadmin
8. create a database called gomango and create a table with "description"	"sdate" "edate"	"name" .OR import the sql databse from the gomango file
9. to find all the projects click the view projects tab
```

# Usage
```scss
1. to create a project Open: http://127.0.0.1:8080
```
![image](https://github.com/YABOIpy/GoMango/assets/110062350/dd703265-fe60-4a03-8ac4-d521cb3c7d04)

```scss
2. to see the projects click the view projects tab
```
![image](https://github.com/YABOIpy/GoMango/assets/110062350/bec894c5-4c9e-4a06-8f8d-a1df1a8936f0)

# Configuration

```hcl
io_mode = "async"

service "http" "web_proxy" {
  listen_addr = "0.0.0.0:8080"

  process "main" {
    command = "server"
  }
}
```


```json
{
  "DataBase": {
    "dbname": "gomango",
    "dbtype": "mysql",
    "port": 3306
  },
  "DbAuth": {
    "ServerIPv4": ["127.0.0.1", "8080"],
    "username": "root",
    "password": ""
  }
}

```

# How
![Diagram](https://github.com/YABOIpy/GoMango/assets/110062350/ac85db67-533b-4c19-9aa8-7f281b6a5033)

