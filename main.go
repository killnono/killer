package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "github.com/gorilla/websocket"
// "io"
)

func main(){

    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World------gin web http server")
    })

    // r.GET("/ping", ping)  

    router.Run(":12345")
  
    // r.Run(bindAddress)  

}



// var upGrader = websocket.Upgrader{  
//     CheckOrigin: func (r *http.Request) bool {  
//        return true  
//     },  
//  }  
 
//  //webSocket请求ping 返回pong  
//  func ping(c *gin.Context) {  
//     //升级get请求为webSocket协议
//     ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)  
//     if err != nil {  
//        return  
//     }  
//     defer ws.Close()  
//     for {
//        //读取ws中的数据  
//        mt, message, err := ws.ReadMessage()  
//        if err != nil {  
//           break  
//        }  
//        if string(message) == "ping" {  
//           message = []byte("pong")  
//        }  
//        //写入ws数据
//        err = ws.WriteMessage(mt, message)  
//        if err != nil {  
//           break  
//        }  
//     }  
//  }  
 




