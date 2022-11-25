package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/FirePing32/go-carbon/utils"
    "fmt"
    "log"
    "encoding/json"
    "net/http"
    "io"
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
        apiResp, err := http.Get(APIUrl)
        if err != nil {
            log.Fatalln(err)
            return c.SendString(err.Error())
        }

        contents, err := io.ReadAll(apiResp.Body)
        var result types.Response
        json.Unmarshal(contents, &result)
        resp, err := json.Marshal(result)
        return c.SendString(string(resp))
    })

    app.Listen(":3000")
}

