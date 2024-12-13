package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Kaaaaazuya/aws-cost-notification/cost"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
)

// createTimePeriod returns the start and end date of the current month
func createTimePeriod() (start, end *string) {
	today := time.Now()
	mf := time.Date(today.Year(), today.Month(), 1, 0, 0, 0,0, time.UTC)
	me := mf.AddDate(0, 0, -1)

	mfs := mf.Format("2006-01-02")
	mes := me.Format("2006-01-02")
	return &mfs, &mes
}

func sumCost(cost *costexplorer.GetCostAndUsageOutput) (total string){
	sum := 0.0
	for _, data := range cost.ResultsByTime[0].Groups {
		amount, _ := strconv.ParseFloat(*data.Metrics["UnblendedCost"].Amount, 64)
		sum = sum + amount
	}
	total = fmt.Sprintf("%.10f", sum)
	return total
}

func  handleRequest(ctx context.Context, event json.RawMessage) error {
	log.Println("コスト取得バッチ 開始")
	c, err := cost.NewExplorerClient("ap-northeast-1")
	if err != nil {
		return err
	}

	s, e := createTimePeriod()
	log.Println("start:", *s, "end:", *e)
	cost, err := cost.GetCostAndUsage(context.TODO(), c, s, e)
	if err != nil {
		return err
	}

	log.Println(cost)
	total := sumCost(cost)
	log.Println("合計コスト:", total)
	log.Println("コスト取得バッチ 終了")
	return nil
}



func main() {
	lambda.Start(handleRequest)
}

