<!-- Daftar enpoint -->
# Dokumentasi Url Enpoint api

## Glosarium url
1. Kamus
    - `<name>`: nama endpoint (opsional). Nama suatu endpoint di dokumen `.MD`
    - `<method>`: method yang digunakan untuk request. `{GET, POST, PUT, DELETE}`
    - `<url>`: alamat url endpoint 
    - `$host`: host webservice + port layanan berjalan. (port default = 8080). `Contoh: localhost:8080`
2.  format penyajian: `<name> <method> <url>`
    Contoh: 
        - Upload proposal `POST $host/api/file/upload?folder=proposal` 
        - `                POST $host/api/user/donatur/create/$id`

## Daftar URL
1. Endpoint auth
- `login: POST $host/api/login`
- `refresh token: POST $host/api/refresh`

2. Endpoint Users
- `registrasi: POST $host/api/register`
- `POST $host/api/user/donatur/create/$id`
- `POST $host/api/user/peneliti/create/:id`
- `GET $host/api/user/verify`
- `GET $host/api/user/review/$id`
- `GET $host/api/user/profile`
- `PUT $host/api/verifikasi/$id`
- `PUT $host/api/user/complete-payment`

3. Endpoint Project
- `GET $host/api/project/opendonasi?page=<page>`
- `GET $host/api/project/diverifikasi?page=<page>`
- `GET $host/api/project/ongoing?page=<page>`
- `GET $host/api/project/$id/detail`
- `GET $host/api/project/myproject?limit=<limit>`
- `GET $host/api/project/revisi`
- `GET $host/api/project/:id/preview`
- `GET $host/api/project/hassubmit?page=<page>`
- `GET $host/api/project/$id/review`
- `POST $host/api/project/create`
- `PUT $host/api/project/$id/edit`
- `PUT $host/api/project/$id/upload/proposal`
- `PUT $host/api/project/$id/upload/klirens`
- `PUT $host/api/project/$id/submit`
- `PUT $host/api/project/$id/reject`
- `PUT $host/api/project/$id/verifikasi`
- `PUT $host/api/project/my_contrib?page=<page>`

4. Endpoint Upload and Donwload File
- Upload proposal `POST $host/api/file/upload?folder=proposal`
- Upload klirens `POST $host/api/file/upload?folder=klirens`
- Upload laporan `POST $host/api/file/upload?folder=laporan`
- Upload panduan skema penelitian `POST $host/api/file/upload?folder=panduan`
- Upload gambar `POST $host/api/file/upload?folder=gambar`
- Upload foto profil`POST $host/api/file/upload?folder=profil`
- Upload icon skema penelitian `POST $host/api/file/upload?folder=icon`
- Download file `GET $host/api/file/download?fileurl=<fileurl>`

5. Endpoint Pengajuan Skema Penelitian
- `GET $host/api/pengajuan/create`
- `POST $host/api/pengajuan/open`

6. Endpoint DetailBudgets
- Create budget: `POST $host/api/project/:id/budget/create`
- Update budget `PUT $host/api/project/:id/budget/:budgetid/update`
- Delete budget `DELETE $host/api/project/:id/budget/:budgetid/delete`

7. Endpoint Tahapan Penelitian
- Create `POST $host/api/project/:id/tahap/create`
- List `GET $host/api/project/:id/tahap/list`
- Update `PUT $host/api/project/:id/tahap/update/:tahapid`
- Delete `DELETE $host/api/project/:id/tahap/delete/:tahapid`

8. Endpoint Donasi Penelitian
- Create Donasi `POST $host/api/project/:id/donasi/create`
- Request webhook dari payment gateway `POST $host/api/donasi/notif`
- Get Detail Donasi `GET $host/api/donasi/:id`
- Melihat History Donasi pribadi `GET $host/api/donasi/history`
- Melihat History Donasi suatu proyek `GET $host/api/project/:id/donasi/histori`
- Melihat History Donatur suatu proyek `GET $host/api/project/:id/donasi/contributor`

9. 