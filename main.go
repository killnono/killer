package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
    "fmt"
)

type Room struct {
    
    RoomId string `json:"roomId"`
    RoomName string    `json:"roomName"`
    RoomSize int     `json:"roomSize"`
    RoomInfo string `json:"roomInfo"`
    RoomState bool   `json:"roomState"`
} 

var rooms [2]Room // Room 数据集

func main(){
    initRooms()

    router := gin.Default()

    // index api
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg":"welcome to watchdog",
        })
    })

    //获取所有房间信息api
    router.GET("/rooms", func(c *gin.Context) {
     getRooms(c)  
    })

    //根据房间ID获取房间信息
    router.GET("/room", func(c *gin.Context) {
        getRoomByRoomId(c)  
       })

    //更新房间状态api
    router.POST("/currentState", func(c *gin.Context) {
        updateRoom(c)
    })

    router.Run(":12345")
    startTimer()
}

//获取所有房间信息和状态服务
func getRooms(c *gin.Context){
    c.JSON(200, gin.H{
        "data" : rooms,
    })
}

//获取所有房间信息和状态服务
func getRoomByRoomId(c *gin.Context){
    roomId := c.DefaultQuery("roomId","");
    var r Room;
    for j := 0; j < len(rooms); j++ {
        room := rooms[j];
        rId := room.RoomId
        if (rId == roomId) {
           r = room;
        }
    }
    c.JSON(200, gin.H{
        "data" : r,
    })
}

// 更新接入的房间状态信息
func updateRoom(c *gin.Context){
    var r Room;
    err := c.Bind(&r)
    if err != nil {
        fmt.Println(err)
    }
    for j := 0; j < len(rooms); j++ {
        room := rooms[j];
        rId := room.RoomId
        if (rId == r.RoomId) {
            room.RoomState =  r.RoomState 
            rooms[j] = room;
        }
    }
    c.JSON(http.StatusOK, gin.H{
        "roomId":r.RoomId,
        "state":r.RoomState,
    })
}
// 房间信息数据初始化，模拟数据库
func initRooms(){
   room1 := Room{
       RoomId : "58f98fcf9d570c8583074629",
       RoomName : "Nasa",
       RoomInfo : "投影,音响",
       RoomSize : 20 ,
       RoomState : false,
   } 

   room2 := Room{
    RoomId : "58f98fcf9d570c8583074620",
    RoomName : "地球",
    RoomInfo : "投影",
    RoomSize : 10 ,
    RoomState : false,
   }

   rooms[0] = room1;
   rooms[1] = room2;

}


// test: 启动定时器
func startTimer(){
    fmt.Println("startTimer func ")
    t2 := time.NewTimer(time.Second * 10)
    for {
        select {
        // case <-t1.C:
        //     println("5s timer")
        //     t1.Reset(time.Second * 5)

        case <-t2.C:
            fmt.Println("10s timer")
            t2.Reset(time.Second * 10)
        }
    }

}




