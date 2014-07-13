package main

import (
    "fmt"
    "net/http"
    _ "io/ioutil"
    "io"
    "encoding/csv"
    "strconv"
    matrix "github.com/skelterjohn/go.matrix"
)

func main() {
    m1 := matrix.MakeDenseMatrix([]float64{1,2,3,4,5,6}, 3, 2)
    fmt.Println(m1)
    fmt.Println(matrix.Transpose(m1))
    res, _ := http.Get("http://archive.ics.uci.edu/ml/machine-learning-databases/wine-quality/winequality-red.csv")
    defer res.Body.Close()

    reader := csv.NewReader(res.Body)
    reader.Comma = ';'
    //contents, _ := ioutil.ReadAll(res.Body)
    for {
        col, err := reader.Read()
        if err == io.EOF { break }
        //strings.Split(col, ";")
        x := make([]float64, len(col))
        for i := range col { x[i], _ = strconv.ParseFloat(col[i], 64) }
        m1 := matrix.MakeDenseMatrix(x, len(x), 1)
        fmt.Println(m1)
    }

}
