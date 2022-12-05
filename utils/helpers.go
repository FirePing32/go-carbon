package utils

import (
	"encoding/json"
	"net/http"
)

func GetJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

// func AssertType(files interface{}) interface{} {
//     m, ok := files.(map[string]interface{})
//     if !ok {
//         return fmt.Errorf("want type map[string]interface{};")
//     }
//     for k, v := range m {
//         fmt.Println(k, "=>", v)
//     }
//     return m
// }

func GetFileName(files map[string]interface{}) string {
    keys := make([]string, 0, len(files))
    for k := range files {
        keys = append(keys, k)
    }
    var filename string
    for i := 0; i<len(keys); i++ {
        if i>=0 && i<len(keys) {
            filename = keys[i]
        }
    }
    return filename
}

func GetCode() {

}
