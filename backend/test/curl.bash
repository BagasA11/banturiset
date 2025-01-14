for ($var = 1; $var -le 100; $var++) {
    Invoke-WebRequest "http://localhost:8080/api/project/opendonasi?page=1" 
}

