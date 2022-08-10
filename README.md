# technical-test-telkom

SOAL 

1. Di Indonesia, ada pecahan mata uang rupiah, yaitu :
* 100.000,-
* 50.000,-
* 20.000,-
* 10.000,-
* 5.000,-
* 2.000,-
* 1.000,-
* 500,-
* 200,-
* 100,-
Buatlah sebuah fungsi untuk menghitung berapa lembar pecahan yang harus dikeluarkan dari
input harga (dengan pembulatan ke atas jita punya harga pecahan antara 1 sampai 99)
Input : 145.000
Output:
{
“Rp. 100.000”:1,
“Rp. 20.000”:2,
“Rp. 5.000”:1,
}
Input: 2050
Ouput:
{
“Rp. 2.000”:1,
“Rp. 100”:1,
}

2. Anda diminta untuk membuat sebuah function dimana function tersebut berfungsi untuk
menentukan apakah dari dua data string yang diberikan membutuhkan sekali proses edit atau
lebih. Jika lebih dari sekali proses edit berarti function tersebut akan mengembalikan response
False, sedangkan jika hanya sekali proses edit maka function tersebut akan mengembalikan
response True. Proses edit di sini dapat berarti melakukan insert sebuah character, remove
sebuah character, atau replace sebuah character.
Contoh
GIVEN INPUT 1 → telkom
GIVEN INPUT 2 → telecom
RESULT → False
GIVEN INPUT 1 → telkom
GIVEN INPUT 2 → tlkom
RESULT → True

3. Apakah ada kesalahan dari script di bawah ini? Jika ada tolong jelaskan dimana letak
kesalahannya dan bagaimana anda memperbaikinya. Jika tidak ada, tolong jelaskan untuk apa
script di bawah ini.
FROM golang
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go get github.com/tools/godep
RUN godep restore
RUN go install github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
LISTEN 80

4. Menurut anda apakah tujuan penggunaan microservices?

5. Bagaimana cara index bekerja pada sebuah database?

6. Buat sebuah sebuah backend service Shopping Cart yang harus memilik api:
● Api tambahProduk
- Attribute kodeProduk, kuantitas
- Menambahkan produk dengan kuantitas yang ditentukan.
- Apabila produk sudah ada di dalam Cart, tambahkan kuantitasnya.
● Api hapusProduk(string kodeProduk)
- Menghapus produk dari Cart.
● Api tampilkanCart
- Menampilkan isi Cart dengan format {kodeProduk}- ({kuantitas })
Buatlah backend service cart berikut feature code dan unit testnya.