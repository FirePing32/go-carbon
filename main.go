package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FirePing32/go-carbon/utils"
    "fmt"
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
        utils.GetJson(APIUrl, content)
        filename := utils.GetFileName(content.Files)
        fileContent := content.Files[filename].(map[string]interface{})["content"]
        return c.JSON(fileContent)
    })

    app.Listen(":3000")
}

