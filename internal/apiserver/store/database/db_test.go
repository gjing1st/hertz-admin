// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 19:05$

package database

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	statistic2 "information-publish/internal/pkg/model/statistic"
	"information-publish/internal/pkg/service/statistic"
	"testing"
)

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}
type LogRecord struct {
	JobName string     `bson:"jobName"` //任务名
	Command string     `bson:"command"` //shell命令
	Err     string     `bson:"err"`     //脚本错误
	Content string     `bson:"content"` //脚本输出
	Tp      TimePorint //执行时间
}

var (
	err        error
	collection *mongo.Collection
	lr         *LogRecord
	iResult    *mongo.InsertOneResult
	id         primitive.ObjectID
	cursor     *mongo.Cursor
)

type Tes struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestGetMgoCli(t *testing.T) {
	cli := GetMgoCli()
	//fmt.Println("cli", cli)
	mgo := cli.Database("test")

	collection = mgo.Collection("test_table111")
	//插入某一条数据
	//tt := struct {
	//	Id   int
	//	Name string
	//}{1, "guojing"}
	var ss statistic.StatisticService
	//statistic.Init()
	res, _ := ss.RealTimeQuery("mmyypt_app_events", "csmp", 0)
	var r statistic2.RealTimeResponse
	json.Unmarshal(res, &r)
	//if iResult, err = collection.InsertOne(context.TODO(), r); err != nil {
	//	fmt.Print("err ", err)
	//	return
	//}
	//_id:默认生成一个全局唯一ID
	//id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())

	cond := Tes{
		Id:   1,
		Name: "guojing",
	}
	if cursor, err = collection.Find(context.TODO(), cond); err != nil {
		fmt.Println("==========err===", err)
		return
	}
	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("=====err", err)
	}
	fmt.Println("len", len(results))
	for _, result := range results {
		fmt.Println("~~~~~~~~~~", result)
	}
}
