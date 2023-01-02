package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FirePing32/go-carbon/utils"
    "fmt"
    "log"
    "strconv"
    "encoding/base64"
)

func main() {

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        resp := c.SendString("API is UP!")
        return resp
    })

    app.Get("/api", func(c *fiber.Ctx) error {
        gistId := c.Query("gistid")
        fSize := c.Query("fsize")
        fColor := c.Query("fcolor")
        bgColor := c.Query("bgcolor")
        if gistId == "" || fSize == ""|| fColor == "" || bgColor == "" {
            queryErr := map[string]interface{}{
                "statusCode": 400,
                "error": "Missing required fields. See https://github.com/FirePing32/go-carbon for details",
            }
            return c.JSON(queryErr)
        }
        APIUrl := fmt.Sprintf("https://api.github.com/gists/%s", gistId)
        var content = new(utils.Response)
        code, err := utils.GetJson(APIUrl, content)
        if code == 200 && err == nil {
            filename := utils.GetFileName(content.Files)
            fileContent := content.Files[filename].(map[string]interface{})["content"]

            fSize, e := strconv.Atoi(fSize)
            if e != nil {
                panic(e)
            }
            b, err := utils.GenerateImage(fileContent.(string), fColor, bgColor, float64(fSize))
                if err != nil {
                    log.Println(err)
                    return err
                }

            imgData := base64.StdEncoding.EncodeToString(b)
            imgMap := map[string]interface{}{
                "base64Data": imgData,
            }

            return c.JSON(imgMap)
        } else {
            resp := map[string]interface{}{
                "statusCode": code,
                "error": err,
            }
            return c.JSON(resp)
        }
    })

    app.Listen(":3000")
}

