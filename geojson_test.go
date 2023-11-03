package gcfgeobackend 

import (
	"context"
	"fmt"
	pasproj "github.com/HRMonitorr/PasetoprojectBackend"
	"github.com/whatsauth/watoken"
	"testing"
)

//func TestGCHandlerFunc(t *testing.T) {
//	data := GCHandlerFunc("string", "GIS", "geogis")
//
//	fmt.Printf("%+v", data)
//}

func TestGeneratePaseto(t *testing.T) {
	//privateKey, publicKey := watoken.GenerateKey()
	//fmt.Println(privateKey)
	//fmt.Println(publicKey)
	hasil, err := watoken.Encode("fahad", "PrivateKey")
	fmt.Println(hasil, err)
}

func TestUpdateData(t *testing.T) {
	data := LonLatProperties{
		Type:   "Polygon",
		Name:   "muda",
		Volume: "1",
	}
	up := UpdateNameGeo("MONGOSTRING", "GismongoDB", context.Background(), data)
	fmt.Println(up)
}

func TestDeleteDataGeo(t *testing.T) {
	data := LonLatProperties{
		Type:   "Point",
		Name:   "asum",
		Volume: "1",
	}
	up := DeleteDataGeo("MONGOSTRING", "GismongoDB", context.Background(), data)
	fmt.Println(up)
}

func TestInsertUser(t *testing.T) {
	conn := GetConnectionMongo("MONGOSTRING", "GismongoDB")
	pass, _ := pasproj.HashPass("fahad")
	data := RegisterStruct{
		Username: "fahad",
		Password: pass,
	}
	ins := InsertUserdata(conn, data.Username, data.Password)
	fmt.Println(ins)
}

func TestIsexist(t *testing.T) {
	data := IsExist("token", "PublicKey")
	fmt.Println(data)
}
