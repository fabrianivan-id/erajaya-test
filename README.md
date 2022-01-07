# Erajaya-Test
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)
[![Go Reference](https://img.shields.io/badge/gmaps-reference-blue?logo=GMaps&logoColor=white)](https://github.com/googlemaps/google-maps-services-go)

# Table of Content
- [Description](#description)
- [How to Use](#how-to-use)
- [Cloud Architecture](#cloud)
- [Feature](#feature)
- [Endpoints](#endpoints)
- [Credits](#credits)

# Description
Erajaya test merupakan projek berbasis Restful API yang dibuat dengan struktur MVC(Media-View-Controller) yang lebih mudah untuk dilakukan perbaikan error dan bug.

# Cloud Architecture
![Cloud](https://github.com/fabrianivan21/erajaya-test/blob/main/documentation/AWS.jpg)
- Saya menggunakan arsitektur cloud seperti ini adalah untuk mencegah kebocoran data produk yang belum akan dirilis sehingga hanya produk sudah dirilis saja yang dapat diakses

# Feature
List of overall feature in this Project (To get more details see the API Documentation below)
| No.| Feature        | Role                     | Keterangan                                                                              |
| :- | :------------- | :----------------------- | :-------------------------------------------------------------------------------------- |
| 1. | Product        | Admin                    | Add Product, Show product sorted by latest, A to Z, Z to A, highest price, lowest price |                                     

# How to Use
- Install Go and Database MySQL/XAMPP
- Clone this repository in your $PATH:
```
$ git clone https://github.com/fabrianivan21/erajaya-test.git
```
- Create file .env based on this project 
``
sample-env
``
- Don't forget to create database name as you want in your MySQL
- Run program with command
```
go run main.go
```
# Endpoints
Read the API documentation here [API Endpoint Documentation]() (Swagger)

# Credits
- [Fabrian Ivan Prasetya](https://github.com/fabrianivan21) (Author)
