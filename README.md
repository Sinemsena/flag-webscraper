# Flag ile Çalışan Web Scraper
Golang kullanarak web sitelerinden başlık, açıklama ve tarih bilgilerini çeken CLI web
scraper aracıdır.

## Kullanım

1. Yardım ve Seçim Ekranı
Herhangi bir parametre vermeden çalıştırdığınızda kullanılabilir flagleri ve site listesini görüntülenir:

```bash

go run main.go
```
2. Temel Tarama
Belirlenen sitenin tüm bilgilerini (Başlık, Açıklama, Tarih) getirir:

```bash

go run main.go -1
```

3. Filtreleme Seçenekleri

Tarih Filtreleme: Tarih bilgisini gizlemek için -date kullanılır.

```bash

go run main.go -1 -date
```
 * Açıklama Filtreleme: Açıklama bilgisini gizlemek için -description kullanılır.

```bash

go run main.go -1 -description
```
 * Tam Filtreleme: Sadece haber başlıklarını görmek için her iki parametreyi de ekleyerek filtreleme yapılır.

```bash

go run main.go -1 -date -description
```
## Kayıt Sistemi
Yapılan her tarama, projenin kök dizininde otomatik olarak oluşturulan /sonuc klasörüne, o anki saat damgasıyla .txt formatında kaydedilir.
