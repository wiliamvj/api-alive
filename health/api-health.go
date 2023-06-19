package health

import (
  "bytes"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"

  "github.com/robfig/cron"
)

type MessageDiscordPayload struct {
  Content string `json:"content"`
}

func ApiHealthCheck() {
  c := cron.New()

  c.AddFunc("*/30 * * * * *", func() {
    res, err := http.Get(os.Getenv("HEALTH_ENDPOINT"))
    if err != nil || res.StatusCode != http.StatusAccepted {
      content := fmt.Sprintf("O código de status HTTP é %d", res.StatusCode)
      payload := MessageDiscordPayload{
        Content: content,
      }

      jsonPayload, err := json.Marshal(payload)
      if err != nil {
        log.Println("Error encoding JSON payload:", err)
        return
      }

      _, err = http.Post(os.Getenv("DISCORD_WEBHOOK"), "application/json", bytes.NewBuffer(jsonPayload))
      if err != nil {
        log.Fatal("Error to send message for api health check:", err)
      }
      log.Println("Message sent to the webhook successfully!")
    }
  })

}
