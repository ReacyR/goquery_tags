package main

import (
    "net/http"
    "fmt"
    "encoding/csv"
    "os"

    "github.com/PuerkitoBio/goquery"
)


func errorCheck(err error) {
        if err != nil {
            fmt.Println(err)
        }
}


func main() {

    url := "https://stackoverflow.com/questions/tagged/goquery?tab=newest&pagesize=50"

    response, err := http.Get(url)
    defer response.Body.Close()
    errorCheck(err)

    if response.StatusCode > 400 {
        fmt.Println("Status Code: ", response.StatusCode)
    }

    doc, err := goquery.NewDocumentFromReader(response.Body)
    errorCheck(err)

    file , err := os.Create("goquery_tags.csv")
    errorCheck(err)
    writer := csv.NewWriter(file)

    doc.Find("div.bt.bc-black-100.mln24.pl24").Find("div.mln24").Each(func(i int, s *goquery.Selection) {
    
        title := s.Find("a").Text()
        user := s.Find("div.user-details").Find("a").Text()
        userProfile, _ := s.Find("div.user-details").Find("a").Attr("href")
        datePosted := s.Find("span.relativetime").Text()
        views := s.Find("div.views").Text()
        answers := s.Find("div.status.answered").Find("strong").Text()
        votes := s.Find("span.vote-count-post").Text()
        tags := s.Find("div.tags").Text()
        link, _ := s.Find("a").Attr("href")
       
        goqueryQuestions := []string{title, user, userProfile, datePosted, views, answers,votes, tags, link}

        writer.Write(goqueryQuestions)

    })
    
    writer.Flush()

}
