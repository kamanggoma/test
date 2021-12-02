package main

import (
    "time"
	"sync"
    "strconv"
    "github.com/tarm/serial"
)
/*
func main() {
    go func() {
        c := &serial.Config{Name: "COM3", Baud: 921600, Size: 8, StopBits: serial.Stop1}
        Port, _ := serial.OpenPort(c)
        Port.Flush()
        i := 0
        for {
            i = i + 1
            s := strconv.Itoa(i)
            Port.Write([]byte(s + "..."))
            println("write:"+s)
            time.Sleep(time.Duration(1) * time.Millisecond) // 버퍼링
            // time.Sleep(time.Duration(1000) * time.Millisecond) // 블로킹
        }
    }()
    go func() {
        c := &serial.Config{Name: "COM3", Baud: 921600, Size: 8, StopBits: serial.Stop1}
        Port, _ := serial.OpenPort(c)
        Port.Flush()
        for {
            buf := make([]byte, 10)
            n, _ := Port.Read(buf)
            println("                     read:" + string(buf[:n]))
            time.Sleep(time.Duration(10) * time.Millisecond) // 버퍼링
            // time.Sleep(time.Duration(1) * time.Millisecond) // 블로킹
        }
    }()
    for {
        time.Sleep(time.Duration(1000) * time.Millisecond)
    }
}
*/