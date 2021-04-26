package main

import(
    "fmt"
    "net"
    "os"
    "sync"
)

func port(ip string, p int, wg *sync.WaitGroup){
    address := fmt.Sprintf("%s:%d",ip,p)
    conn,err := net.Dial("tcp",address)
    if err != nil {
        wg.Done()
        return
    }
    conn.Close()
    fmt.Printf("[+] Puerto: %d - ABIERTO\n",p)
    wg.Done()
}

func main(){
    var wg sync.WaitGroup
    for i:=1; i<=65535; i++{
        wg.Add(1)
        go port(os.Args[1], i, &wg)
    }
    wg.Wait()
}
