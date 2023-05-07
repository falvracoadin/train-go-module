package system

import (
	"bytes"
	"crypto/md5"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"github.com/nfnt/resize"
	excel "github.com/xuri/excelize/v2"
	"gopkg.in/gomail.v2"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CurlGet(url string) (res map[string]interface{}, err error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if body != nil {
		json.Unmarshal(body, &res)
	}
	return res, err
}

func CurlPost(url string, contentType string, form map[string]interface{}) (res map[string]interface{}, err error) {

	json_data, _ := json.Marshal(form)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(json_data))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if body != nil {
		json.Unmarshal(body, &res)
	}
	return res, err
}

func CurlPut(url string, form map[string]interface{}) (res map[string]interface{}, err error) {

	json_data, _ := json.Marshal(form)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if body != nil {
		json.Unmarshal(body, &res)
	}
	return res, err
}

func CurlDelete(url string, form map[string]interface{}) (res map[string]interface{}, err error) {

	json_data, _ := json.Marshal(form)
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(json_data))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	if body != nil {
		json.Unmarshal(body, &res)
	}
	return res, err
}

func getExtention(base64 string) string {
	png := strings.Contains(base64, "image/png")
	jpeg := strings.Contains(base64, "image/jpeg")
	gif := strings.Contains(base64, "image/gif")

	ex := ".jpg"

	if png {
		ex = ".png"
	} else if jpeg {
		ex = ".jpeg"
	} else if gif {
		ex = ".gif"
	}

	return ex
}

func Compress(nmfile string, height string, width string) error {
	file, err := os.Open(nmfile)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	hg, err := strconv.Atoi(height)
	uhg := uint(hg)

	wd, err := strconv.Atoi(width)
	uwd := uint(wd)

	// (width, height, input file, kernel sampling)
	m := resize.Resize(uwd, uhg, img, resize.Lanczos3)

	out, err := os.Create(nmfile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	fmt.Println("Ini Masuk Compress")
	return jpeg.Encode(out, m, nil)
}

func CheckFormValidation(r map[string]interface{}) (res bool) {
	var v = 0
	for _, val := range r {
		if val == nil || val == "" {
			v += 1
		}
	}
	return v > 0
}

// param width 700 height 700
func UploadBase64ToImg(param string, path string, name string) (data interface{}, err error) {
	// Read form fields
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)

		if err != nil {
			fmt.Println("tidak bisa create")
		}
	}

	if param != "" {
		b64data := param[strings.IndexByte(param, ',')+1:]
		dec, err := b64.StdEncoding.DecodeString(b64data)
		if err != nil {
			return data, err
		}

		time := time.Now().Unix()
		//fmt.Println(nm, "alex")s
		nmfile := name + strconv.Itoa(int(time)) + getExtention(param)
		f, err := os.Create(path + nmfile)
		if err != nil {
			return data, err
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			return data, err
		}
		if err := f.Sync(); err != nil {
			return data, err
		}
		fmt.Println(nmfile)
		Compress(path+nmfile, "0", "700")
		fmt.Println("Berhasil Compress")
		return nmfile, err
	}
	fmt.Println("alex2")
	return data, err
}

// Ini Function Cek IsImage
func IsImageFile(typefile string) (valid bool) {
	return strings.Contains(typefile, "image")
}

func IntToCharStr(i int) string {
	return string('A' - 1 + i)
}
func RequestGetParams(e echo.Context) (res map[string]interface{}) {
	e.MultipartForm()
	fromParam := make(map[string]interface{})
	fromJson := make(map[string]interface{})

	r := e.Request()
	if err := r.ParseForm(); err != nil {
	}
	jsn, _ := json.Marshal(r.Form)
	if err := json.Unmarshal(jsn, &fromParam); err != nil {
	}

	if err := json.NewDecoder(e.Request().Body).Decode(&fromJson); err != nil {
	}

	if len(fromJson) > 0 {
		res = fromJson
	} else if len(fromParam) > 0 {
		var data = make(map[string]interface{})
		for k, val := range fromParam {
			var afterDec []interface{}
			v, _ := json.Marshal(val)
			if err := json.Unmarshal(v, &afterDec); err != nil {
			}
			data[k] = afterDec[0]
		}
		res = data
	}

	if res == nil {
		res = map[string]interface{}{}
	}

	return res
}

func UploadFile(c echo.Context, param string, path string) (data interface{}, err error) {
	// Read form fields
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)

		if err != nil {
			fmt.Println("tidak bisa create")
		}
	}
	file, err := c.FormFile(param)
	if err != nil {
		return nil, err
	}
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(path + file.Filename)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}

	return file.Filename, err
}

func ReqTrimSpace(req map[string]interface{}) (res map[string]interface{}) {
	// for k, val := range req {
	// 	if reflect.TypeOf(val).String() == "string" {
	// 		req[k] = strings.TrimSpace(val.(string))
	// 		fmt.Println(req[k], "----- alex")
	// 	}
	// 	// valstr := val.(string)
	// 	// val = strings.TrimSpace(valstr)
	// 	// req[k] = val
	// }
	return req
}

func MultiReqTrimSpace(req []interface{}) (res []interface{}) {
	// for _, val := range req {
	// 	data_req := val.(map[string]interface{})
	// 	for k, v := range data_req {
	// 		if reflect.TypeOf(v).String() == "string" {
	// 			data_req[k] = strings.TrimSpace(v.(string))
	// 		}
	// 	}
	// }
	return req
}

func SendEmail(email string, subject string, msg string) (data string, err error) {

	Email := os.Getenv("SMTP_EMAIL")
	Password := os.Getenv("SMTP_PASSWORD")
	Host := "smtp.gmail.com"
	Port := 587

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", Email)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", msg)
	//mailer.Attach("../img/barang-1638329840.jpg")

	dialer := gomail.NewDialer(
		Host,
		Port,
		Email,
		Password,
	)

	abc := dialer.DialAndSend(mailer)
	if abc != nil {
		log.Fatal(abc.Error())
	}

	log.Println("Mail sent!")
	return data, err
}

func ToCharStr(i int) string {
	return string('A' - 1 + i)
}

func ReadExcel(file interface{}) (res []map[string]interface{}) {
	// Membaca file excel
	// Deklarasi 2 kali karena menggunakan library yang berbeda
	excel, _ := excel.OpenFile("./" + file.(string))
	f, _ := excelize.OpenFile("./" + file.(string))

	// Deklarasi sheet yang akan dibaca
	sheetName := "Sheet One"

	// Get Total Column
	var datacol []string
	cols, _ := excel.GetCols(sheetName)

	for _, col := range cols {
		for k, rowCell := range col {
			// Deklarasi Kondisi Column yang mau dibaca dengan kondisi cell tidak kososng dan index k merupakan index 0
			if rowCell != "" && k == 0 {
				datacol = append(datacol, rowCell)
			}
		}
	}

	// Deklarasi Total Rows and Cell
	CountRows := len(f.GetRows(sheetName))
	totalCols := len(datacol)
	var totalRows int

	for i := 2; i <= CountRows; i++ {
		if f.GetCellValue(sheetName, fmt.Sprintf("A%d", i)) != "" {
			totalRows++
		}
	}

	index := 2

	var r = make([]map[string]interface{}, totalRows-1)
	for i := 2; i <= totalRows; i++ {
		r[i-2] = make(map[string]interface{})
		if f.GetCellValue(sheetName, fmt.Sprintf("A%d", i)) != "" {
			for j := 1; j <= totalCols; j++ {
				key := strings.ReplaceAll(strings.ToLower(
					strings.TrimSpace(
						f.GetCellValue(
							sheetName, fmt.Sprintf(
								ToCharStr(j)+"%d", 1),
						),
					),
				), " ", "_")
				value := f.GetCellValue(sheetName, fmt.Sprintf(ToCharStr(j)+"%d", index))
				r[i-2][key] = value
			}
		}
		index++
	}

	return r
}

func ExportExcel(e echo.Context) error {

	xlsx := excelize.NewFile()

	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)

	// Deklarasi Header Bahasa
	xlsx.SetCellValue(sheet1Name, "A1", "Nama Lokasi")
	xlsx.SetCellValue(sheet1Name, "E1", "Deskripsi")
	xlsx.SetCellValue(sheet1Name, "F1", "Alamat")

	style, err := xlsx.NewStyle(`{"number_format": 49}`)
	if err != nil {
		fmt.Println(err)
	}
	xlsx.SetCellStyle(sheet1Name, "C1", "C1000", style)
	xlsx.SetCellStyle(sheet1Name, "D1", "D1000", style)

	var d []map[string]interface{}
	for k, v := range d {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("J%d", k+2), v["id"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("K%d", k+2), v["name"])
	}

	xlsx.SaveAs("./Template.xlsx")

	return e.File("Template.xlsx")
}
