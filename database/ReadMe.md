# Su dụng
### . Cấu trúc lại thư mục
```
.
├── config  <-- Đọc cấu hình sử dụng Viper
│   └── config.go
├── controller <-- Các controller xử lý request đến
│   ├── ManufactureController.go
│   └── ProductController.go
├── model <-- Định nghĩa các model
│   ├── Category.go
│   ├── Country.go
│   ├── Image.go
│   └── Product.go
├── public <-- lưu trữ ảnh
│   └── img
├── repo <-- Lưu các phương thức xử lý dữ liệu chuyên cho từng Model
│   ├── CategoryRepo.go
│   ├── CountryRepo.go
│   ├── ProductRepo.go
│   └── Repo.go  <-- Lưu biến toàn cục Database Connection `var Db *gorm.DB`
├── routes <-- Cấu hình định tuyến các request đến ứng với phương thức của Controller`
│   └── ConfigRouter.go
├── sql <-- Lưu các file SQL script tạo bảng và xoá bảng
│   ├── DropTable.sql  <-- Tạo cấu trúc bảng
│   └── OnlineShop.sql <-- Drop các bảng
├── app.go <-- File chạy chính
├── dev.yml <-- File cấu hình ở môi trường development
├── go.mod
├── go.sum
├── GORM.md
```
## B1. 
Tạo database và kết nối database
## B2.
cấu hình định dạng YAML [dev.yml](dev.yml)
```yaml
db:
  host: localhost
  port: 3306
  user: yourUserName
  pass: yourPassword
  database: yourDatabase
```
### B3.
Chạy file OnlineShop.sql trong Thư mục sql

### B4.
Chạy file app.go để kết nối database



# API:  .
```
+get all product: http://127.0.0.1:3000/api/product
+get product by Id: http://127.0.0.1:3000//api/product/id

# get product params: 
lấy 10 sản phẩm bắt đầu từ sản phẩm thứ 3
http://127.0.0.1:3000/api/product/search/search_items?limit=10&newest=3
  limit: số page trên 1 trang
  newest: vị trí product bắt đầu lấy

lấy 10 sản phẩm theo categorys
1:"Giày thể thao nam",
2:"Giày thể thao nữ",
3:"Giày da nam",
4:"Giày da nữ",
http://127.0.0.1:3000/api/product/search/search_items?limit=10&categorys=1
  limit: số page trên 1 trang
  categorys: danh mục sản phẩm

sắp xếp sản phẩm theo giá 
order (desc, asc)
http://127.0.0.1:3000/api/product/search/search_items?limit=10&order=desc
  limit: số page trên 1 trang
  order: sắp xếp sản phẩm


lấy sản phẩm theo danh mục và sắp xếp 
order (desc, asc)
http://127.0.0.1:3000/api/product/search/search_items?limit=10&newest=3&categorys=1&order=desc
  limit: số page trên 1 trang
  newest: vị trí product bắt đầu lấy
  categorys: danh mục sản phẩm
  order: sắp xếp sản phẩm
```



