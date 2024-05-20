package luckymanygrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"spider/luckymanygrpc/grpc_server/pb"
)

type server struct {
	pb.DispenseServer
}
type Recharge struct {
	ActivityAddGold   int    `json:"activity_add_gold"`
	AdgroupName       string `json:"adgroup_name"`
	Amount            int    `json:"amount"`
	BankCode          string `json:"bank_code"`
	CampaignName      string `json:"campaign_name"`
	Channel           string `json:"channel"`
	Coin              int    `json:"coin"`
	CreateDayFirstPay bool   `json:"create_day_first_pay"`
	CreativeName      string `json:"creative_name"`
	Currency          string `json:"currency"`
	Event             string `json:"event"`
	FirstCharge       bool   `json:"first_charge"`
	IdrAmount         int    `json:"idr_amount"`
	NetworkName       string `json:"network_name"`
	NumberOneRecharge int    `json:"number_one_recharge"`
	OrderID           string `json:"order_id"`
	PayChannel        int    `json:"pay_channel"`
	PayChn            string `json:"pay_chn"`
	PayType           int    `json:"pay_type"`
	Pkg               string `json:"pkg"`
	ProductID         string `json:"product_id"`
	RechargeCoin      int    `json:"recharge_coin"`
	StartTs           int    `json:"start_ts"`
	ThirdAmount       int    `json:"third_amount"`
	ThirdCurrency     string `json:"third_currency"`
	Ts                int    `json:"ts"`
	UserID            int    `json:"user_id"`
}

func (s *server) OnStreamMsg(ctx context.Context, req *pb.StreamMsg) (*pb.Null, error) {
	// 实现你的服务逻辑
	var decodedData Recharge
	err := json.Unmarshal(req.Data, &decodedData)
	if err != nil {
		fmt.Println("Error decoding binary data:", err)
	}
	log.Printf("Received user_id: %v", decodedData.UserID)
	return &pb.Null{}, nil
}

func StartServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Start server: %v\n", lis.Addr())
	s := grpc.NewServer()
	pb.RegisterDispenseServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
