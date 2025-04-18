package main

import (
	"fmt"
	"log"

	"github.com/Time-Capsule/Api-Gateway
/api"
	"github.com/Time-Capsule/Api-Gateway
/api/handler"
	_ "github.com/Time-Capsule/Api-Gateway
/docs"
	pb "github.com/Time-Capsule/Api-Gateway
/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	MemoryConn, err := grpc.NewClient(fmt.Sprintf("memory%s", ":7777"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer MemoryConn.Close()
	TimeLineConn, err := grpc.NewClient(fmt.Sprintf("timeline%s", ":8888"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer TimeLineConn.Close()

	comment := pb.NewCommentServiceClient(MemoryConn)
	media := pb.NewMediaServiceClient(MemoryConn)
	memory := pb.NewMemoryServiceClient(MemoryConn)
	history := pb.NewHistoricalServiceClient(TimeLineConn)
	mile := pb.NewMilestoneServiceClient(TimeLineConn)
	time := pb.NewTimeLineServiceClient(TimeLineConn)
	cus := pb.NewCustomEventsServiceClient(TimeLineConn)
	h := handler.NewHandler(comment, media, memory, cus, history, mile, time)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
