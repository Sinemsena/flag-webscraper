package main

import (
        "bufio"
        "flag"
        "fmt"
        "os"
        "strings"
        "time"

        "github.com/fatih/color"
        "github.com/gocolly/colly/v2"
        "github.com/mbndr/figlet4go"
)

func main() {
    ascii, _ := figlet4go.NewAsciiRender().Render("Yavuzlar-2")
    color.Cyan(ascii)

    site1 := flag.Bool("1", false, "displays the first news site")
    site2 := flag.Bool("2", false, "displays the second news site")
    site3 := flag.Bool("3", false, "displays the third news site")
    tarihfiltreli := flag.Bool("date", false, "filters the date part")
    aciklamafiltreli := flag.Bool("description", false, "filters the description part")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        fmt.Fprintf(os.Stderr, " -1\tdisplays the first news site\n -2\tdisplays the second news site\n -3\tdisplays the third news site\n -date\n\tfilters the date part\n -description\n\tfilters the description part\n")
    }

    flag.Parse()

    // Hiçbir seçm yapılmadıys seçim ekranını göster
    if !*site1 && !*site2 && !*site3 {
        flag.Usage()
        return
    }

    url, siteAdi := "", ""
    if *site1 {
        url, siteAdi = "https://thehackernews.com/", "HackerNews"
    } else if *site2 {
        url, siteAdi = "https://www.securityweek.com/", "SecurityWeek"
    } else if *site3 {
        url, siteAdi = "https://www.hackread.com/", "HackRead"
    }

    scrapeAndShow(url, siteAdi, *tarihfiltreli, *aciklamafiltreli)
}

func scrapeAndShow(url string, siteAdi string, tarihFiltreleme bool, aciklamaFiltereleme bool) {
	klasörAdi :="sonuc"
	//klasr yoksa oluştır
	if _, err := os.Stat(klasörAdi); os.IsNotExist(err) {
		err := os.Mkdir(klasörAdi, 0755)
			if err != nil {
				color.Red("Klasör oluşturulamadı", err)
				return
			}
		}

    // Her tarama içni ayrı dosya 
    dosyaAdi := fmt.Sprintf("%s/%s_%s", klasörAdi, siteAdi, time.Now().Format("150405")) 
    if tarihFiltreleme { dosyaAdi += "_tarihfiltreli" }
    if aciklamaFiltereleme { dosyaAdi += "_aciklamafiltreli" }
    dosyaAdi += ".txt"
    dosya, _ := os.Create(dosyaAdi)
    defer dosya.Close()
    writer := bufio.NewWriter(dosya)

    c := colly.NewCollector()
    sayac := 1

    c.OnHTML("article, .body-post, .cat-post-item", func(e *colly.HTMLElement) {
            title := strings.TrimSpace(e.ChildText("h2, .home-title, .cs-entry__title"))
            if title == "" { return }
            color.Magenta("%d. Haber:", sayac)
            fmt.Println(title)
            writer.WriteString(fmt.Sprintf("%d. Haber: %s\n", sayac, title))

            if !aciklamaFiltereleme {
                    aciklama := strings.TrimSpace(e.ChildText("p, .home-desc, .cs-entry__excerpt"))
                    if aciklama != "" {
                        color.Blue("Açıklama:")
                        fmt.Println(aciklama)
                        writer.WriteString("Açıklama: " + aciklama + "\n")
                    }
            }

            if !tarihFiltreleme {
                    tarih := strings.TrimSpace(e.ChildText("time, .h-datetime, .cs-meta-date"))
                    if tarih != "" {
                        color.Red("Tarih:")
                        fmt.Println(tarih)
                        writer.WriteString("Tarih: " + tarih + "\n")
                    }
            }
            fmt.Println("")
            writer.WriteString("\n")
            sayac++
    })
    c.Visit(url)
    writer.Flush()
    color.Green("\n %d haber '%s' dosyasına kaydedildi.", sayac-1, dosyaAdi)
}