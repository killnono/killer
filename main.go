package main

import (
    "github.com/gin-gonic/gin"
    // "github.com/gorilla/websocket"
    "net/http"
    "fmt"
    // "time"
    // "encoding/json"
    // "github.com/satori/go.uuid"
)

type Room struct {
    RoomId string `json:"roomId"`
    RoomName string    `json:"roomName"`
    RoomSize int     `json:"roomSize"`
    RoomInfo string `json:"roomInfo"`
    RoomState bool   `json:"roomState"`
} 

var roomDB = make([]Room, 0, 0) // Room 数据集合

// //模拟房间信息录入数据库
// var roomAlive = make([]Room, 0, 0)
// //模拟房间信息录入数据库
// var roomDB = make([]Room, 0, 0)


func main(){
    initRooms()

    router := gin.Default()

    // index api
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg":"welcome to watchdog",
        })
    })

     // index api
     router.POST("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg":"welcome to watchdog",
        })
    })

    //获取所有房间信息api
    router.GET("/connect", func(c *gin.Context) {
        roomConnect(c)  
       })

    router.GET("/disconnect", func(c *gin.Context) {
        roomConnect(c)  
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

    // go manager.start()
    // go manager.checkAlive()
   
    router.Run(":12345")
    
}

//房间接入连接
func roomConnect(c *gin.Context){
    roomId := c.DefaultQuery("roomId","")
    var msg string = "该房间不存在,请先通过管理后台配置录入"

    for j := 0; j < len(roomDB); j++ {
        room := roomDB[j];
        rId := room.RoomId
        if (rId == roomId) {
            msg = "房间接入成功"
            break;
        }
    }

    // for j := 0; j < len(roomDB); j++ {
    //     room := roomDB[j];
    //     rId := room.RoomId
    //     if (rId == roomId) {
    //         roomAlive = append(roomAlive,room)    
    //         msg = "房间接入成功"
    //         break;
    //     }
    // }
    c.JSON(200, gin.H{
        "msg" : msg,
    })
    // fmt.Println("已接入房间设备：",roomAlive)
}

//房间接入连接
// func roomDisconnect(c *gin.Context){
//     roomId := c.DefaultQuery("roomId","")
//     var msg string = "该房间不存在,请先通过管理后台配置录入"

//     for j := 0; j < len(rooms); j++ {
//         room := rooms[j];
//         rId := room.RoomId
//         if (rId == roomId) {
//             msg = "房间接入成功"
//             break;
//         }
//     }

//     for j := 0; j < len(roomDB); j++ {
//         room := roomDB[j];
//         rId := room.RoomId
//         if (rId == roomId) {
//             roomAlive = append(roomAlive,room)    
//             msg = "房间接入成功"
//             break;
//         }
//     }
//     c.JSON(200, gin.H{
//         "msg" : msg,
//     })
//     fmt.Println("已接入房间设备：",roomAlive)
// }

//获取所有房间信息和状态服务
func getRooms(c *gin.Context){
    c.JSON(200, gin.H{
        "data" : roomDB,
    })
}

//获取所有房间信息和状态服务
func getRoomByRoomId(c *gin.Context){
    succeed := false;
    roomId := c.DefaultQuery("roomId","");
    var r Room
    for j := 0; j < len(roomDB); j++ {
        room := roomDB[j];
        rId := room.RoomId
        if (rId == roomId) {
           r = room
           succeed = true
           break;
        }
    }
    if(succeed){
        c.JSON(200, r)
    }else {
        c.JSON(400, gin.H{
            "msg" : "房间不存在,查询失败",
        })
    }  
}

// 更新接入的房间状态信息
func updateRoom(c *gin.Context){
    succeed := false
    var r Room;
    err := c.Bind(&r)
    if err != nil {
        fmt.Println(err)
    }
    for j := 0; j < len(roomDB); j++ {
        room := roomDB[j]
        rId := room.RoomId
        if (rId == r.RoomId) {
            room.RoomState =  r.RoomState 
            roomDB[j] = room
            succeed = true
            break;
        }
    }
    if(succeed){
        c.JSON(http.StatusOK, gin.H{
            "roomId":r.RoomId,
            "roomState":r.RoomState,
        })
    }else{
        c.JSON(http.StatusBadRequest, gin.H{
            "msg":"房间更新失败",
        }) 
    }
}
// 房间信息数据初始化，模拟数据库
func initRooms(){
   room1 := Room{
       RoomId : "58f98fcf9d570c8583074629",
       RoomName : "NASA",
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

//    rooms[0] = room1;
//    rooms[1] = room2;

   roomDB = append(roomDB, room1)
   roomDB = append(roomDB, room2)

   fmt.Println("rooms1",roomDB)
}


// // 实现websocket
// type ClientManager struct {
//     clients    map[*Client]bool
//     broadcast  chan []byte
//     register   chan *Client
//     unregister chan *Client
// }

// type Client struct {
//     id     string
//     roomId string
//     socket *websocket.Conn
//     send   chan []byte
// }

// type Message struct {
//     Sender    string `json:"sender,omitempty"`
//     Recipient string `json:"recipient,omitempty"`
//     Content   string `json:"content,omitempty"`
// }

// var manager = ClientManager{
//     broadcast:  make(chan []byte),
//     register:   make(chan *Client),
//     unregister: make(chan *Client),
//     clients:    make(map[*Client]bool),
// }

// func (manager *ClientManager) start() {
//     fmt.Println("manager start")

//     for {
//         select {
//         case conn := <-manager.register:
//             manager.clients[conn] = true
//             jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected."})
//             manager.send(jsonMessage, conn)// 通知其他接入设备，有新的设备接入，忽略自己本身
//             fmt.Println("目前接入的设备有：")
//             for conn := range manager.clients {
//                 fmt.Println(conn.id)
//             }
//         case conn := <-manager.unregister:
//             if _, ok := manager.clients[conn]; ok {
//                 close(conn.send)
//                 delete(manager.clients, conn)
//                 jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
//                 manager.send(jsonMessage, conn)//通知其他接入设备，设备断开，忽略自己本身
//             }
//         case message := <-manager.broadcast:
//             for conn := range manager.clients {
//                 select {
//                 case conn.send <- message:
//                 default:
//                     close(conn.send)
//                     delete(manager.clients, conn)
//                 }
//             }
//         }
//     }
// }

// func (manager *ClientManager) send(message []byte, ignore *Client) {
//     for conn := range manager.clients {
//         if conn != ignore {
//             conn.send <- message
//         }
//     }
// }

// func (c *Client) read() {

//     defer func() {
//         manager.unregister <- c
//         c.socket.Close()
//         fmt.Println("unregister","id:",c.id,"--roomId:",c.roomId)

//     }()

//     for {
//         _, message, err := c.socket.ReadMessage()
//         if err != nil {
//             manager.unregister <- c
//             c.socket.Close()
//             break
//         }
//         fmt.Println("读取socket",message)
//         data,err:= json.Marshal(roomDB)//读取全新的房间信息
//         if err != nil {
//             fmt.Println(err)

//         }
//         // jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
        
//         manager.broadcast <- data
//     }
// }

// func (c *Client) write() {

//     defer func() {
//         c.socket.Close()
//     }()

//     for {
//         select {
//         case message, ok := <-c.send:
//             if !ok {
//                 c.socket.WriteMessage(websocket.CloseMessage, []byte{})
//                 return
//             }

//             c.socket.WriteMessage(websocket.TextMessage, message)
//         }
//     }
// }

// func ws(res http.ResponseWriter, req *http.Request) {
//     conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
//     if error != nil {
//         http.NotFound(res, req)
//         return
//     }
//     uid,_ := uuid.NewV4()
//     rId := "";
//     dType:= req.Header.Get("deviceType")
//     if(dType == "room"){
//         rId = req.Header.Get("roomId") 
//     }

//     client := &Client{id: uid.String() ,roomId:rId,socket: conn, send: make(chan []byte)}
//     manager.register <- client

//     go client.read()
//     go client.write()
// }


// func (manager *ClientManager) checkAlive() {
//     fmt.Println("manager checkAlive2")
//     ticker := time.NewTicker(10 * time.Second) 
// 	defer ticker.Stop()
// 	for {
// 		select {
//         case <-ticker.C:
//             fmt.Println("for checkAlive")
//             fmt.Println("alive now socket client")
//             for conn := range manager.clients {
//                 // fmt.Println("alive:",conn.id,"房间id:",conn.roomId)
//                 conn.socket.SetWriteDeadline(time.Now().Add(5 * time.Second))
//                 if err := conn.socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
//                 fmt.Println("房间：",conn.roomId,"--断开了")
//                 return
//                  }

// 		    }
// 	    }

//     }
// }
// func (manager *ClientManager) checkAlive2() {
//     fmt.Println("manager checkAlive")
//     ticker := time.NewTicker(10 * time.Second) 
// 	defer ticker.Stop()
// 	for {
// 		select {
//         case <-ticker.C:
//             fmt.Println("for checkAlive")
//             fmt.Println("alive now socket client")
//             for conn := range manager.clients {
//                 fmt.Println("alive:",conn.id,"房间id:",conn.roomId)
//             }

//             for j := 0; j < len(roomDB); j++ {
//                 room := roomDB[j];
//                 rId := room.RoomId
//                 alive := false
//                 for conn := range manager.clients {
//                         if(rId == conn.roomId){
//                             alive = true
//                             break
//                         }else{
//                             continue
//                         } 
//                      }
//                      roomDB[j].RoomState=alive;
   
//             }

// 		}
// 	}

// }



