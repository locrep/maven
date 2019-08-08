package main

import (
    "fmt"
    "hash/crc32"
    "os"
    "io"
)

func getHash(filename string) (uint32, error) {
    // dosyayı aç
    f, err := os.Open(filename)
    if err != nil {
        return 0, err
    }
    // açılan dosyları mutlaka kapat!
    defer f.Close()

    // hash yapıcı oluştur
    h := crc32.NewIEEE()
    // dosyayı hash yapıcıya kopyala
    // - copy (dst, src) ve geriye (bytesWritten, error) döner
    if _, err := io.Copy(h, f); err != nil {
        return 0, err
    }
    // kaç byte döndüğü bizim için önemli değil, sadece hatayı
    // hatayı yakala
    return h.Sum32(), nil
}

func main() {
    h1, err := getHash("/tmp/test1.txt")
    if err != nil {
        return
    }
    h2, err := getHash("/tmp/test2.txt")
    if err != nil {
        return
    }
    fmt.Println(h1, h2, h1 == h2) // 1276159447 1707422633 false
}