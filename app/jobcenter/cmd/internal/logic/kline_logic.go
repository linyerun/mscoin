package logic

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
	"mscoin/app/jobcenter/cmd/internal/svc"
	"mscoin/app/jobcenter/model"
	"mscoin/common/constant"
	"mscoin/common/tool"
	"sync"
	"time"
)

type KLineLogic struct {
	wg sync.WaitGroup

	svc *svc.ServiceContext
	ctx context.Context
	logx.Logger
}

func NewKLineLogic(ctx context.Context, svc *svc.ServiceContext) *KLineLogic {
	return &KLineLogic{
		ctx:    ctx,
		svc:    svc,
		Logger: logx.WithContext(ctx),
	}
}

func (kl *KLineLogic) Do(period string) {
	defer kl.wg.Wait()
	go kl.getKlineData("BTC-USDT", "BTC/USDT", period)
	go kl.getKlineData("ETH-USDT", "ETH/USDT", period)
}

func (kl *KLineLogic) getKlineData(instId string, symbol string, period string) {
	kl.wg.Add(1)
	defer kl.wg.Done()

	// 逻辑代码
	// 准备header
	header := make(map[string]string)
	timestamp := tool.ISO(time.Now())
	sign := tool.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/candles?instId="+instId+"&bar="+period, kl.svc.Config.KLineConf.SecretKey)
	header["OK-ACCESS-KEY"] = kl.svc.Config.KLineConf.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = kl.svc.Config.KLineConf.Pass
	// 发起请求
	api := kl.svc.Config.KLineConf.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	respBytes, err := tool.GetWithHeader(api, header, kl.svc.Config.KLineConf.Proxy)
	if err != nil {
		kl.Logger.Errorf("get okx website msg err. err=%+v", err)
		return
	}
	// json.Unmarshal
	type KLineResult struct {
		Code string     `json:"code"`
		Msg  string     `json:"msg"`
		Data [][]string `json:"data"`
	}
	res := new(KLineResult)
	err = json.Unmarshal(respBytes, res)
	if err != nil {
		kl.Logger.Errorf("unmarshal json data err. err=%+v", err)
		return
	}

	if res.Code == "0" && len(res.Data) > 0 {
		// 存储数据到mongo中
		// 生成kline
		var klineList []any // 最新数据(klineList[0])=>最久数据(klineList[len(klineList)-1])
		for i := 0; i < len(res.Data); i++ {
			kline := model.NewKline(res.Data[i], period)
			klineList = append(klineList, kline)
		}
		tableName := model.GetKlineTableName(symbol, period)
		// 删除mongo中的数据
		latestTime := klineList[len(klineList)-1].(*model.Kline).Time
		many, err := kl.svc.MongoClient.Database(kl.svc.Config.Mongo.Database).Collection(tableName).DeleteMany(kl.ctx, bson.D{{"time", bson.D{{"$gte", latestTime}}}})
		if err != nil {
			kl.Logger.Errorf("删除Table=%s的数据出现异常. err=%+v", tableName, err)
		} else {
			kl.Logger.Infof("删除Table=%s的数据有%d条", tableName, many)
		}
		// 插入数据到mongo
		_, err = kl.svc.MongoClient.Database(kl.svc.Config.Mongo.Database).Collection(tableName).InsertMany(kl.ctx, klineList)
		if err != nil {
			kl.Logger.Errorf("新增Table=%s的数据出现异常. err=%+v", tableName, err)
			return
		} else {
			kl.Logger.Infof("插入klineList成功，插入条数: %d", len(klineList))
		}

		// 把数据放入kafka中
		latestKlineBytes, err := json.Marshal(klineList[0])
		if err != nil {
			kl.Errorf("json.Marshal(klineList[0]) err. err=%+v", err)
			return
		}
		writer := kl.svc.KafkaWriterCreate(constant.Kline1MTopic + "-" + instId)
		defer func() { err = writer.Close(); logx.Infof("关闭kafka writer. err=%+v", err) }()
		err = writer.WriteMessages(kl.ctx, kafka.Message{
			Value: latestKlineBytes,
		})
		if err != nil {
			kl.Errorf("kafka WriteMessages err. err=%+v", err)
		}
	}
}
