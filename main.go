// Group application errors into concise operator-facing summaries.
package main

import ("bytes"; "encoding/json"; "fmt"; "io"; "net/http"; "os"; "strings"; "time")
type Config struct { BaseURL string `json:"base_url"`; APIKey string `json:"api_key"`; Model string `json:"model"` }
func main() {
  var cfg Config; raw, _ := os.ReadFile("config/development.json"); if json.Unmarshal(raw, &cfg) != nil { panic("invalid config") }
  input, err := os.ReadFile("examples/input.txt"); if err != nil { panic(err) }
  payload := map[string]any{"model": cfg.Model, "messages": []map[string]string{
    {"role":"system", "content":"Perform error log summarization. Return concise JSON for human review."},
    {"role":"user", "content":string(input)}}, "max_tokens":256, "temperature":0}
  body, _ := json.Marshal(payload); req, _ := http.NewRequest("POST", strings.TrimRight(cfg.BaseURL, "/")+"/chat/completions", bytes.NewReader(body))
  req.Header.Set("Authorization", "Bearer "+cfg.APIKey); req.Header.Set("Content-Type", "application/json")
  client := &http.Client{Timeout:30*time.Second}; resp, err := client.Do(req); if err != nil { panic(err) }; defer resp.Body.Close()
  out, _ := io.ReadAll(resp.Body); if resp.StatusCode >= 300 { panic(fmt.Sprintf("gateway returned %d: %s", resp.StatusCode, out)) }; fmt.Println(string(out))
}
