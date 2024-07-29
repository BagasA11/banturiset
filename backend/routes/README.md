<!-- Daftar enpoint -->
# Daftar enpoint api

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
- `POST $host/api/file/upload?folder=proposal`
- `POST $host/api/file/upload?folder=klirens`
- `POST $host/api/file/upload?folder=laporan`
- `POST $host/api/file/upload?folder=panduan`
- `POST $host/api/file/upload?folder=gambar`
- `POST $host/api/file/upload?folder=profil`
- `POST $host/api/file/upload?folder=icon`
- `GET $host/api/file/download?fileurl=<fileurl>`

5. Endpoint Pengajuan Skema Penelitian
- `GET $host/api/pengajuan/create`
- `POST $host/api/pengajuan/open`

6. Endpoint DetailBudgets