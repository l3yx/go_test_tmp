package main

import (
	"archive/zip"
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("Hello World")
}

func handler1(req *http.Request) {
	cmdName := req.URL.Query()["cmd"][0]
	cmd := exec.Command(cmdName)
	cmd.Run()
}

func handler(db *sql.DB, req *http.Request) {
	q := fmt.Sprintf("SELECT ITEM,PRICE FROM PRODUCT WHERE ITEM_CATEGORY='%s' ORDER BY PRICE",
		req.URL.Query()["category"])
	db.Query(q)
}

func unzip(f string) {
	r, _ := zip.OpenReader(f)
	for _, f := range r.File {
		p, _ := filepath.Abs(f.Name)
		// BAD: This could overwrite any file on the file system
		ioutil.WriteFile(p, []byte("present"), 0666)
	}
}
