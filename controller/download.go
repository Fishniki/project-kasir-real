package controller

import (
    "fmt"
    "net/http"
    "os"
    "github.com/jung-kurt/gofpdf"
    "github.com/PuerkitoBio/goquery"
)

func CetakPDF(w http.ResponseWriter, r *http.Request) {
    // Buka file HTML
    file, err := os.Open("view/index.html")
    if err != nil {
        http.Error(w, fmt.Sprintf("Could not open HTML file: %v", err), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Parse HTML file
    doc, err := goquery.NewDocumentFromReader(file)
    if err != nil {
        http.Error(w, fmt.Sprintf("Could not parse HTML file: %v", err), http.StatusInternalServerError)
        return
    }

    // Buat dokumen PDF
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "", 12)

    // Extract data from HTML and add to PDF
    doc.Find("body").Each(func(i int, s *goquery.Selection) {
        text := s.Text()
        pdf.MultiCell(0, 10, text, "", "", false)
    })

    // Output PDF
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", "attachment; filename=struk.pdf")
    err = pdf.Output(w)
    if err != nil {
        http.Error(w, fmt.Sprintf("Could not generate PDF: %v", err), http.StatusInternalServerError)
        return
    }
}
