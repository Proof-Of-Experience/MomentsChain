package main

import (
    "os"
    "io"
    "io/ioutil"
    "runtime"
    "strings"
    "path"
    "fmt"

    "moments.life/chain"
    "moments.life/http"
    "github.com/rifflock/lfshook"
    "github.com/gin-gonic/gin"
    log "github.com/sirupsen/logrus"
)

func init() {
    // Initialize logger
    log.SetFormatter(&log.TextFormatter{
        CallerPrettyfier: func(f *runtime.Frame) (string, string) {
            s := strings.Split(f.Function, ".")
            funcname := s[len(s) - 1]
            _, filename := path.Split(f.File)
            printd := fmt.Sprintf("%s:%d", filename, f.Line)
            return funcname, printd
        },
        ForceColors: true,
    })

    var filename string = "../log/moments.log"
    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

    if err == nil {
        fw := io.Writer(f)
        log.AddHook(lfshook.NewHook(fw, &log.JSONFormatter{
            CallerPrettyfier: func(f *runtime.Frame) (string, string) {
                s := strings.Split(f.Function, ".")
                funcname := s[len(s)-1]
                _, filename := path.Split(f.File)
                printd := fmt.Sprintf("%s:%d", filename, f.Line)
                return funcname, printd
            },
            PrettyPrint: false,
        }))
    }

    // Only log the info severity or above
    log.SetLevel(log.InfoLevel)
    val, _ := os.LookupEnv("LOGLEVEL")
    if strings.ToLower(val) == "debug" {
        log.SetLevel(log.DebugLevel)
    } else if strings.ToLower(val) == "trace" {
        log.SetLevel(log.TraceLevel)
    }

    log.SetReportCaller(true)

    var filename1 string = "../log/gin.log" // TODO - make it configurable /env
     f1, err := os.OpenFile(filename1, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
     val1, _ := os.LookupEnv("DISPLAYSTDOUT")
     if strings.ToLower(val1) == "true" {
         gin.DefaultWriter = io.MultiWriter(f1, os.Stdout)
         log.SetOutput(os.Stdout)
     } else {
         gin.DefaultWriter = io.Writer(f1)
         log.SetOutput(ioutil.Discard)
     }
}

func main() {
    os.Setenv("LOGLEVEL", "debug")

    chain.RunChain()
    http.StartHTTP()
}
