# Building Web-RTC Project: 



## Architecture: 

- Routes -> Fiber -> Rooms
- WebSockets -> WebSockets/Fiber -> Chat
- Stream -> WebRTC -> Stream

```plain
                             |----> Chat
    Rooms ---> WebSocket ----
                 ^          |----> Viewer
                 |
              Stream
```

- Fiber => similar to express
