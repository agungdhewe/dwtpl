# DwLogger
Suatu modul lightweight logger untuk development applikasi golang, serta idependent dari 3rd party modul lain.
Modul ini hanya menambahkan beberapa fungsi ke logger standard golang, dengan fungsi `Info()`, `Log()`, `Warning()` dan `Error()`.
Fungsi-fungsi tersebut digunakan lebih untuk keperluan logging di screen, dengan menambahkan beberapa visual sehingga diharapkan lebih memudahkan proses debugging. 

### Contoh Penggunaan
    dwlog.New()
    dwlog.Info("ini info")
    dwlog.Log("coba screen logging script")
    dwlog.Warning("ini adalah warning")
    dwlog.Error("ini log untuk error")
    
### Contoh Tampilan saat running
![contoh tampilan](https://github.com/agungdhewe/dwlog/blob/main/ss.png)
