package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	common "github.com/KiraCore/sekai/INTERX/common"
	interx "github.com/KiraCore/sekai/INTERX/config"
)

// CacheDataCheck is a function to check cache data if it's expired.
func CacheDataCheck(rpcAddr string, isLog bool) {
	for {
		err := filepath.Walk(interx.GetResponseCacheDir(),
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				delete := false

				if !info.IsDir() && info.Size() != 0 {
					// check cache json data

					common.Mutex.Lock()
					data, _ := ioutil.ReadFile(path)
					common.Mutex.Unlock()

					result := common.InterxResponse{}
					err := json.Unmarshal([]byte(data), &result)

					if err == nil && result.ExpireAt.Before(time.Now()) && result.Response.Block != NodeStatus.Block {
						delete = true
					}
				}

				if path != interx.GetResponseCacheDir() && delete {
					if isLog {
						fmt.Println("deleting file ... ", path)
					}

					common.Mutex.Lock()
					err := os.Remove(path)
					common.Mutex.Unlock()

					if err != nil {
						if isLog {
							fmt.Println("Error deleting file: ", err)
						}
						return err
					}

					return nil
				}

				return nil
			})

		if err != nil {
			log.Println(err)
		}
	}
}
