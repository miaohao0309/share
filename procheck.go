package main

import "log"
import "time"
import "flag"
import "fmt"
import "os"
import "strings"
import "os/exec"
import "io/ioutil"
import "encoding/json"
import "net/http"
import "html/template"

func main() {
    var configPath,processStat string

    flag.StringVar(&configPath, "config", "", "Config file path for procheck.")
    flag.Parse()
       
    if configPath == "" {
        fmt.Println("Please supply a config file for procheck,such as '--config=procheck.json'")
        return
    }
 
    data, err := ioutil.ReadFile(configPath)
    if err != nil { 
        fmt.Println(err.Error())
        return
    }

    var res map[string]interface{}
    if err := json.Unmarshal(data, &res); err != nil {
        fmt.Println("Please check the config file!!")
        os.Exit(1)
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(`<html>
        <head><title>Other Exporter</title></head>
        <body>
        <h1>Other Exporter</h1>
        <p><a href='/metrics'>Metrics</a></p>
        </body>
        </html>`))
    })
    
    map1 := make(map[string]string)
    
    go func() {
        for {
            time.Sleep(3*time.Second)
            for k, v := range res {
                if k == "process" {
                    for name, port := range v.(map[string]interface{}) {
                        processStat = checkProcess(name, port.(string))
                        map1[name] = processStat
                    }      
                }
            }
        }
    }()
        
    tmpl := template.Must(template.New("index").Parse(`
<html>
<body>
<pre style="word-wrap: break-word; white-space: pre-wrap;">
# HELP process_status two status of the process, up or down.
# TYPE process_status gauge
{{- range $key, $value := . }}
process_status{name="{{- $key -}}"} {{ $value -}}
{{ end -}}
</body>
</html>
    `))

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, map1 )
    })
    log.Fatal(http.ListenAndServe(":9101", nil))
}

func checkProcess(name string, port string) string {
    var tmp1 []byte
    var f1,f2,tmp2 string
    var err error
    var cmd *exec.Cmd
    
    cmd = exec.Command("systemctl", "is-active", name)
    tmp1,_ = cmd.Output()
    if strings.Trim(string(tmp1), "\n") == "active" {
        f1 = "1"
    } else {
        f1 = "0"
    }
    
    tmp2 += "netstat -tlpn |grep " + port
    cmd = exec.Command("/bin/sh", "-c", tmp2)
    if _, err = cmd.Output();err != nil {
        f2 = "0"
    }else {
        f2 = "1"
    }

    if f1 == "0" && f2 == "0" {
        return "down"
    } else {
        return "up"
    }
}
