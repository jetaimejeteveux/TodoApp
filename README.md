# TodoApp
Aplikasi Back-End untuk melakukan aktivitas ToDo

# List Library
- "fmt"
- "net/http"
- "encoding/json"
- "database/sql"
- "github.com/go-sql-driver/mysql"
- "github.com/gorilla/mux"
- "github.com/gorilla/schema"
- "github.com/swaggo/swag/cmd/swag"
- "github.com/swaggo/http-swaggo"
- "github.com/alecthomas/template"

# How to Run Application

-> Database MySQL
   pada folder "database/db.go"
   
   Ubah Parameter sesuai dengan database local yang dipakai
   exec file SQL Query
   
-> Buka Command Prompt dan arahkan terminal sesuai dengan lokasi file
   masukan perintah "go run main.go"
-> Aplikasi siap digunakan


# How to Use Application
* Menampilkan semua \
Alamat    : http://localhost:8080/todos \
Method    : GET \
Parameter : - 
* Menampilkan ToDo dengan ID \
   Alamat    : http://localhost:8080/todos/1 \
   Method    : GET \
   Parameter : - \
   *note "/1" sebagai inputan ToDo ID ke - n
* Membuat ToDo \
   Alamat    : http://localhost:8080/todos \
   Method    : POST \
   Parameter 
   >"name" : nama todo 
* Edit Todo \
   Alamat    : http://localhost:8080/todos/1 \
   Method    : PUT \
   Parameter : 
   >"name" : nama todo (string) \
   >"complete : true (boolean) \
*note "/1" sebagai inputan ToDo ID ke - n
* Delete Todo \
   Alamat    : http://localhost:8080/todos/1 \
   Method    : DELETE \
   Parameter : - \
   *note "/1" sebagai inputan ToDo ID ke - n
* Menampilkan Dokumentasi \
   Alamat : http://localhost:8080/swagger/
   
