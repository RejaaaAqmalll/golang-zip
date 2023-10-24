package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
)

// func get Image form URL
func getImageUrl(url string) (image.Image, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request failed with status: %s", response.Status)
	}

	imageData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(imageData))

	if err != nil {
		return nil, err
	}

	return img, nil
}

// func saveImage(img image.Image, filePath string) error {
// 	outPut, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer outPut.Close()

// 	// memilih format image jpg/png
// 	if err := png.Encode(outPut, img); err != nil {
// 		return err
// 	}
// 	return nil
// }

func generateZip(w http.ResponseWriter, r *http.Request, nama string) {
	imageUrl := "https://central-kube.smartlink.id/smartlink_qr/view?text=tesmesasge=&label=tes&kode=D001&tipe=dryer&channel=XENDIT"

	img, err := getImageUrl(imageUrl)

	if err != nil {
		fmt.Println("Error fetch image : ", err.Error())
		return
	}

	zipData := new([]byte)
	// arsip, err := os.Create(nama + ".zip")
	// if err != nil {
	// 	panic(err)
	// }
	// defer arsip.Close()

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	for i := 1; i < 8; i++ {
		namaGambar := fmt.Sprintf("gambar%d.png", i)

		// fileWriter, err := zipWriter.Create(namaGambar)
		// if err != nil {
		// 	panic(err)
		// }

		data := []byte(namaGambar)

		w1, err := zipWriter.Create("qr/" + namaGambar)
		if err != nil {
			panic(err)
		}
		png.Encode(w1, img)

		_, err = w1.Write(data)
		if err != nil {
			panic(err)
		}

	}
	zipWriter.Close()

	w.Write(*zipData)
}

// func download(w http.ResponseWriter, r *http.Request) {
// 	nama := r.FormValue("nama")
// 	if nama == "" {
// 		http.Error(w, "Error can't empety", 500)
// 		return
// 	}

// 	generateZip(w, r, nama)

// 	w.Header().Set("Content-Type", "application/zip")
// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", nama+".zip"))

// }

func main() {
	http.HandleFunc("/unduh", func(w http.ResponseWriter, r *http.Request) {
		nama := r.FormValue("nama")
		if nama == "" {
			http.Error(w, "Error can't empety", 500)
			return
		}

		generateZip(w, r, nama)

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", nama+".zip"))
	})
	fmt.Println("Server started at :8030")
	http.ListenAndServe(":8030", nil)
}

// func main() {
// 	// buat file zip
// 	arsip, err := os.Create("arsip.zip")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer arsip.Close()
// 	ZipWriter := zip.NewWriter(arsip)

// 	// open file hasil generteQR
// 	file1, err := os.Open("test.csv")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer file1.Close()

// 	// masukan hasil generate Qr sesuai folder
// 	w1, err := ZipWriter.Create("csv/test.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	if _, err := io.Copy(w1, file1); err != nil {
// 		panic(err)
// 	}

// 	ZipWriter.Close()
// }
