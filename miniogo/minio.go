package miniogo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go-skeleton-manager-rabbitmq/system"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var client *minio.Client

func Connect() error {
	endpoint := os.Getenv("MIN_HOST")
	acceskey := os.Getenv("MIN_ACCESS")
	secretkey := os.Getenv("MIN_SECRET")
	log.Println(endpoint, acceskey, secretkey)
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(acceskey, secretkey, ""),
		Secure: false,
	})
	log.Println(err)
	if err != nil {
		return err
	}

	client = minioClient
	log.Println("Successfully connected to Minio")

	return nil
}

func Disconnect() {
	client = nil
	log.Println("Successfully disconnect Minio")
}

// function to upload file to minio bucket http://localhost:9000/image/edit.png
func GetUrlMinio(path, nmfile string) (url string) {
	host := os.Getenv("MIN_HOST")

	if nmfile != "" {
		url = "http://" + host + "/" + path + "/" + nmfile
	}
	return url
}

//func PutObject(ctx context.Context, file *multipart.FileHeader, bucketName, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (info minio.UploadInfo, err error) {
//	if client == nil {
//		err := Connect()
//		if err != nil {
//			//return nil, err
//		}
//	}
//
//	src, err := file.Open()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer src.Close()
//
//	fileStat := file.Size
//
//	uploadInfo, err := client.PutObject(ctx, bucketName, objectName, src, fileStat, minio.PutObjectOptions{ContentType: "application/octet-stream"})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
//}

func MakeBucket(ctx context.Context, bucket string) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			//return nil, err
		}
	}

	// get file from minio
	err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: false})
	if err != nil {
		log.Println(err)
		//return nil, err
	}

	log.Printf("Successfully Create Bucket %s\n", bucket)
	//return nil, nil
}

func CheckBucket(ctx context.Context, bucket string) (res bool, err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return false, err
		}
	}

	// get file from minio
	found, err := client.BucketExists(ctx, bucket)
	if err != nil {
		log.Println(err)
		return false, err
	}

	//log.Printf("Successfully Create Bucket %s\n", bucket)
	return found, nil
}

func UploadWithName(file *multipart.FileHeader, ctx context.Context, bucketName string, filename string) (minio.UploadInfo, error) {

	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return minio.UploadInfo{}, err
		}
	}

	found, _ := CheckBucket(ctx, bucketName)
	if !found {
		fmt.Println(found, "--------------------  GA NEMU")
		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`
		MakeBucket(ctx, bucketName)
		SetBucketPolicy(ctx, bucketName, policy)
	}

	log.Println("Starting process upload")
	// convert file to io.Reader
	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return minio.UploadInfo{}, err
	}
	defer src.Close()

	// // Cek File Size
	// sizefile := file.Size
	// if sizefile > 150000 {
	// 	return minio.UploadInfo{}, errors.New("file terlalu besar")
	// }

	// Ini Function Cek File Extention
	typefile := file.Header.Get("Content-Type")
	if system.IsImageFile(typefile) {
		fmt.Println("ini Image")
		// Function Compress Ditaruh Sini
	}

	// upload file to minio
	rs, err := client.PutObject(
		ctx, bucketName, filename, src,
		file.Size, minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		log.Println(err)
		return minio.UploadInfo{}, err
	}

	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, file.Size)
	return rs, nil
}

func UploadPath(ctx context.Context, bucketName string, file *multipart.FileHeader, path string) (minio.UploadInfo, error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return minio.UploadInfo{}, err
		}
	}

	found, _ := CheckBucket(ctx, bucketName)
	if !found {
		fmt.Println(found, "-------------------- GA NEMU")
		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`
		MakeBucket(ctx, bucketName)
		SetBucketPolicy(ctx, bucketName, policy)
	}

	_, err := file.Open()
	if os.IsNotExist(err) {
		// create error file not exist
		return minio.UploadInfo{}, errors.New("file not exist")
	}

	if file.Size > 150000 {
		return minio.UploadInfo{}, errors.New("file terlalu besar")
	}

	// upload file to minio
	rs, err := client.FPutObject(
		ctx, bucketName, file.Filename, path,
		minio.PutObjectOptions{},
	)

	if err != nil {
		log.Println(err)
		return minio.UploadInfo{}, err
	}

	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, rs.Size)
	return rs, nil
}

// get file miniobucket
func GetObject(ctx context.Context, bucketName string, fileName string) (res *minio.Object, err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return nil, err
		}
	}

	// get file from minio
	fileRes, err := client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Successfully get file %s\n", fileName)
	return fileRes, nil
}

func GetBucketPolicy(ctx context.Context, bucketName string) (policy string, err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return "", err
		}
	}

	// get file from minio
	policy, err = client.GetBucketPolicy(ctx, bucketName)
	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Printf("Policy : %s\n", policy, "--------------------")
	return policy, nil
}

func SetBucketPolicy(ctx context.Context, bucketName string, policy string) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			//return "", err
		}
	}

	// get file from minio
	err := client.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		log.Println(err)
		//return "", err
	}

	log.Printf("Policy : %s\n", policy, "--------------------")
	//return policy, nil
}

// delete object minio
func RemoveObject(ctx context.Context, bucketName string, fileName string) (err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return err
		}
	}

	err2 := client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err2 != nil {
		log.Println(err2)
		return err2
	}

	log.Printf("Successfully delete file %s\n", fileName)
	return nil
}

// delete object batch minio
func RemoveObjects(bucketName string, fileName ...string) (err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return err
		}
	}

	for _, file := range fileName {
		err2 := client.RemoveObject(context.Background(), bucketName, file, minio.RemoveObjectOptions{})
		if err2 != nil {
			log.Println(err2)
			return err2
		}
	}

	log.Printf("Successfully delete %d file\n", len(fileName))
	return nil
}

// metadata object minio
func StatObject(ctx context.Context, bucketName string, fileName string) (res minio.ObjectInfo, err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return minio.ObjectInfo{}, err
		}
	}

	data, err2 := client.StatObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err2 != nil {
		log.Println(err2)
		return minio.ObjectInfo{}, err2
	}

	return data, nil
}

// get batch object metadata minio
func StatObjects(bucketName string, fileName ...string) (res []minio.ObjectInfo, err error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return []minio.ObjectInfo{}, err
		}
	}

	res = []minio.ObjectInfo{}
	for _, file := range fileName {
		data, err2 := client.StatObject(context.Background(), bucketName, file, minio.GetObjectOptions{})
		if err2 != nil {
			log.Println(err2)
			return []minio.ObjectInfo{}, err2
		}
		res = append(res, data)
	}

	return res, nil
}

//  ============================= Minio API LOYALTO ==========================
var clientApi *minio.Client

func ConnectApi() error {
	endpoint := os.Getenv("MINIO_SERVER")
	acceskey := os.Getenv("MINIO_ACCESS_KEY")
	secretkey := os.Getenv("MINIO_SECRET_KEY")

	// Initialize minio client object.
	minioClientApi, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(acceskey, secretkey, ""),
		Secure: false,
	})

	if err != nil {
		return err
	}

	clientApi = minioClientApi
	log.Println("Successfully connected to Minio")

	return nil
}

type UploadFileInfo struct {
	Key      string
	RspMinio map[string]interface{}
}

func UploadMinioDev(fileUpload *multipart.FileHeader) (response map[string]interface{}) {
	bucketName := os.Getenv("MINIO_BUCKET")
	path := os.Getenv("MINIO_PATH")
	Type := os.Getenv("MINIO_TYPE")
	exp := os.Getenv("MINIO_EXPIRED")
	api := os.Getenv("MINIO_API")
	url := api + "/api/v1/upload"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := fileUpload.Open()
	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("object", fileUpload.Filename)
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
	}
	_ = writer.WriteField("bucket", bucketName)
	_ = writer.WriteField("path", path)
	_ = writer.WriteField("type", Type)
	_ = writer.WriteField("expired", exp)
	err2 := writer.Close()
	if err2 != nil {
		fmt.Println(err2)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	var json_map map[string]interface{}
	err3 := json.Unmarshal(body, &json_map)
	if err3 != nil {
		fmt.Print(err3)
	}

	return json_map
}

func Upload(file *multipart.FileHeader) (rsp *UploadFileInfo, err error) {
	bucketName := os.Getenv("MINIO_BUCKET")
	api := os.Getenv("MINIO_API")

	var filekey string
	ctx := context.Background()
	var res map[string]interface{}
	if api == "http://s3-minio:9193" {
		res = UploadMinioDev(file)
		response := res["data"].(map[string]interface{})
		filekey = response["name"].(string)
	} else {
		info, _ := UploadFile(file, ctx, bucketName)
		filekey = info.Key
	}

	filenameKey := &UploadFileInfo{
		Key:      filekey,
		RspMinio: res,
	}

	return filenameKey, nil
}

func UploadFile(file *multipart.FileHeader, ctx context.Context, bucketName string) (minio.UploadInfo, error) {
	// check if client variable empty
	if client == nil {
		err := Connect()
		if err != nil {
			return minio.UploadInfo{}, err
		}
	}

	found, _ := CheckBucket(ctx, bucketName)
	if !found {
		policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": 
					"Allow","Principal": {"AWS": ["*"]},"Resource": 
					["arn:aws:s3:::` + bucketName + `/*"],"Sid": ""}]}`
		MakeBucket(ctx, bucketName)
		SetBucketPolicy(ctx, bucketName, policy)
	}

	//timeunix := time.Now().Format("060102150405")
	timeunix := strconv.Itoa(int(time.Now().UnixNano()))
	extention := strings.ReplaceAll(filepath.Ext(file.Filename), ".", "")
	filename := timeunix + "." + extention

	log.Println("Starting process upload")
	// convert file to io.Reader
	src, err := file.Open()
	if err != nil {
		log.Println(err)
		return minio.UploadInfo{}, err
	}
	defer src.Close()

	// // Cek File Size
	// sizefile := file.Size
	// if sizefile > 150000 {
	// 	return minio.UploadInfo{}, errors.New("file terlalu besar")
	// }

	// Ini Function Cek File Extention
	typefile := file.Header.Get("Content-Type")
	if system.IsImageFile(typefile) {
		fmt.Println("ini Image")
		// Function Compress Ditaruh Sini
	}

	// upload file to minio
	rs, err := client.PutObject(
		ctx, bucketName, filename, src,
		file.Size, minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)

	if err != nil {
		log.Println(err)
		return minio.UploadInfo{}, err
	}

	log.Printf("Successfully uploaded %s of size %d\n", file.Filename, file.Size)
	return rs, nil
}
