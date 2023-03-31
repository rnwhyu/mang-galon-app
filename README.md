# mang-galon-app

base url: https://mang-galon-app-production.up.railway.app

## Routing :
test pada POSTMAN

### Register
<hr>

#### "POST" - `/register`
untuk register user
- Body:
  ```
  {
    "full_name": "string",
    "email": "string",
    "username": "string",
    "password": "string",
    "role_id": 0  // untuk mempermudah testing, sebaiknya di handle server dan otomatis jadi 'buyer'
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```

### Login
<hr>

#### "POST" - `/login`
untuk login user
- Body:
  ```
  {
    "username": "string",
    "password": "string"
  }
  ```
- Response:
  ```
  {
    "data": "string",
    "message": "string"
  }
  ```

### User
<hr>

#### "GET" - `/user`
untuk mendapatkan info user sekarang
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "data": {
      "user_id": 0,
      "full_name": "string",
      "email": "string",
      "username": "string",
      "role_id": 0,
      "created_at": "string"
    },
    "message": "string"
  }
  ```

### Galon
<hr>

#### "GET" - `/galon`
untuk menampilkan seluruh data galon
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "data": [
      {
        "galon_id": 0,
        "brand_name": "string",
        "stock": 0,
        "updatestok_at": "string",
        "created_at": "string"
      }
    ]
    "message": "string"
  }
  ```

#### "POST" - `/galon` - **SELLER ONLY**
untuk mendaftarkan galon baru
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Body:
  ```
  {
    "brand_name": "string",
    "stock": 0
  }
  ```
- Response:
  ```
  {
    "data": {
      "galon_id": 0,
      "brand_name": "string",
      "stock": 0,
      "updatestok_at": "string",
      "created_at": "string"
    },
    "message": "string"
  }
  ```

#### "PUT" - `/galon/:id` - **SELLER ONLY**
untuk mengupdate stok galon
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Body:
  ```
  {
    "stock": 0
  }
  ```
- Response:
  ```
  {
    "data": {
      "galon_id": 0,
      "brand_name": "string",
      "stock": 0,
      "updatestok_at": "string",
      "created_at": "string"
    },
    "message": "string"
  }
  ```

#### "DELETE" - `/galon/:id` - **SELLER ONLY**
untuk menghapus galon
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```


### Order
<hr>

#### "GET" - `/order` - **SELLER ONLY**
untuk menampilkan seluruh order
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "data": [
      {
        "id": 0,
        "user_id": 0,
        "galon_id": 0,
        "total_order": 0,
        "status": "string",
        "updated_at": "string",
        "created_at": "string"
      }
    ]
    "message": "string"
  }
  ```

#### "GET" - `/order/user` - **BUYER ONLY**
untuk menampilkan order milik user
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "data": [
      {
        "id": 0,
        "user_id": 0,
        "galon_id": 0,
        "total_order": 0,
        "status": "string",
        "updated_at": "string",
        "created_at": "string"
      }
    ]
    "message": "string"
  }
  ```

#### "POST" - `/order` - **BUYER ONLY**
untuk membuat order baru
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Body:
  ```
  {
    "galon_id": 0,
    "total_order": 0
  } 
  ```
- Response:
  ```
  {
    "data": {
      "id": 0,
      "user_id": 0,
      "galon_id": 0,
      "total_order": 0,
      "status": "string",
      "updated_at": "string",
      "created_at": "string"
    }
    "message": "string"
  }
  ```

#### "PUT" - `/order/:id/processing` - **SELLER ONLY**
untuk mengupdate status order menjadi `Processing`
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```

#### "PUT" - `/order/:id/on-delivery` - **SELLER ONLY**
untuk mengupdate status order menjadi `On Delivery`
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```

#### "PUT" - `/order/:id/delivered` - **SELLER ONLY**
untuk mengupdate status order menjadi `Delivered`
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```

#### "PUT" - `/order/:id/completed` - **SELLER ONLY**
untuk mengupdate status order menjadi `Completed`
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```

#### "DELETE" - `/order/:id/cancel` - **BUYER ONLY**
untuk membatalkan order dan mengubah status order menjadi `Canceled`
- Header:
  ```
  {
    "Authorization": "Bearer <token>"
  }
  ```
- Response:
  ```
  {
    "message": "string"
  }
  ```
