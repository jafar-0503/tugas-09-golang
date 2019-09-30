package main

import "fmt"
import "strconv"

func main(){
  defer refresh()
  var nilai string
  fmt.Print("Masukkan Nilai Angka:")
  fmt.Scanln(&nilai)

  var number int
  var err error
  number, err = strconv.Atoi(nilai)

  if err == nil{
    fmt.Println(number, "Ini Merupakan Nilai Angka")
  }else{
    fmt.Println(nilai, "Ini Bukan Nilai Angka")
    panic(err.Error())
    fmt.Println("Nilai Angka Saya Ditampilkan")
  }
}

func refresh(){
  if r:=recover(); r != nil{
    fmt.Println("Menampilkan Yang Bukan Nilai Angka ")
  }else{
    fmt.Println("Nilai Angka Lancar-lancar Saja")
  }
}
