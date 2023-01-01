package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FirePing32/go-carbon/utils"
    "fmt"
    "log"
    "encoding/base64"
)

func main() {

    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        resp := c.SendString("API is UP!")
        return resp
    })

    app.Get("/:gistid", func(c *fiber.Ctx) error {
        gistId := c.Params("gistid")
        APIUrl := fmt.Sprintf("https://api.github.com/gists/%s", gistId)
        var content = new(utils.Response)
        code, err := utils.GetJson(APIUrl, content)
        if code == 200 && err == nil {
            filename := utils.GetFileName(content.Files)
            fileContent := content.Files[filename].(map[string]interface{})["content"]

            b, err := utils.GenerateImage(fileContent.(string), "#ffffff", "#300a24", 32)
                if err != nil {
                    log.Println(err)
                    return err
                }

            imgData := base64.StdEncoding.EncodeToString(b)
            imgMap := map[string]interface{}{
                "data": imgData,
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

