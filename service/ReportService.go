package service

import (
	"fmt"
	"math"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
	"sort"
	"strings"
	"time"
)

func GetReport(event model.Event, startTime, endTime, brand, source string) (model.ReportResponse, error) {
	var reportResponse = model.ReportResponse{}
	whereClause := generateWhereClause(event.Id, startTime, endTime, brand, source)
	soldItemDetails, err := repo.GetSoldItemDetailByEventId(setting.DB, whereClause)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.ReportResponse{}, err
	}

	processedData, priceCount, priceList, discountCount, discountList := processSaleRecord(soldItemDetails)
	// core metric
	coreMetrics, err := getCoreMetrics(event.Id, processedData)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Core Metric %s", err.Error()))
		return model.ReportResponse{}, err
	}

	// secondary metric
	secondaryMetrics := getSecondaryMetrics(processedData)

	// distribution
	distribution := getDistribution(priceCount, discountCount, priceList, discountList)

	reportResponse.CoreMetric = coreMetrics
	reportResponse.SecondaryMetric = secondaryMetrics
	reportResponse.Distribution = distribution
	return reportResponse, nil
}

func generateWhereClause(eventId uint64, startTime, endTime, brand, source string) string {
	whereClauses := make([]string, 0)
	if len(startTime) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("order_time >= %s", startTime))
	}

	if len(endTime) == 0 {
		endTime = time.Now().String()
	}
	whereClauses = append(whereClauses, fmt.Sprintf("order_time <= %s", endTime))

	if len(brand) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("i.brand in (%s)", brand))
	}

	if len(source) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("s.source in (%s)", source))
	}

	whereClauses = append(whereClauses, fmt.Sprintf("i.event_id = %d", eventId))

	whereClause := strings.Join(whereClauses, " and ")
	return whereClause
}

func processSaleRecord(records []model.SaleRecordDetail) (map[string]float32, map[string]int, []int, map[string]int, []float64) {
	var data = make(map[string]float32, 0)
	data["item_count"] = float32(len(records))
	var uniqueOrder = make(map[string]bool, 0)
	var uniqueMember = make(map[string]bool, 0)
	var uniqueSku = make(map[string]bool, 0)
	var priceCount = make(map[string]int, 0)
	var discountCount = make(map[string]int, 0)
	var soldAmount = 0
	var returnAmount = 0
	var discountSum float32 = 0
	var salePriceSum = 0
	var maxDiscountSold float64 = 0
	var minDiscountSold float64 = 1
	priceList := make([]int, 0)
	discountList := make([]float64, 0)

	for _, record := range records {
		if record.IsReturn == 1 {
			returnAmount += record.SalePrice
			continue
		}
		soldAmount += record.SalePrice
		uniqueOrder[record.OrderId] = true
		soldAmount += 1
		priceKey := fmt.Sprintf("%d", record.SalePrice)
		if _, ok := priceCount[priceKey]; !ok {
			priceCount[priceKey] = 0
			priceList = append(priceList, record.SalePrice)
		}
		priceCount[priceKey] = priceCount[priceKey] + 1

		discountKey := fmt.Sprintf("%f", record.Discount)
		if _, ok := discountCount[discountKey]; !ok {
			discountCount[discountKey] = 0
			discountList = append(discountList, float64(record.Discount))
		}
		discountCount[discountKey] = discountCount[discountKey] + 1

		if member := uniqueMember[fmt.Sprintf("%d", record.MemberId)]; member {
			continue
		}
		uniqueMember[fmt.Sprintf("%d", record.MemberId)] = true

		if sku := uniqueSku[record.Sku]; sku {
			continue
		}
		uniqueSku[record.Sku] = true

		discountSum += record.Discount
		salePriceSum += record.SalePrice
		maxDiscountSold = math.Max(maxDiscountSold, float64(record.Discount))
		minDiscountSold = math.Min(minDiscountSold, float64(record.Discount))
	}

	data["amount_sold"] = float32(soldAmount)
	data["return_amount"] = float32(returnAmount)
	data["order_sold"] = float32(len(uniqueOrder))
	data["total_member_purchased"] = float32(len(uniqueMember))
	data["sku_sold"] = float32(len(uniqueSku))
	data["average_discount"] = discountSum / data["item_count"]
	data["average_price"] = float32(salePriceSum) / data["item_count"]
	data["max_discount"] = float32(maxDiscountSold)
	data["min_discount"] = float32(minDiscountSold)
	data["average_sku"] = float32(len(uniqueSku)) / float32(len(uniqueMember))
	data["average_item"] = data["item_count"] / float32(len(uniqueMember))
	data["average_amount"] = float32(soldAmount) / float32(len(uniqueMember))

	sort.Ints(priceList)
	sort.Float64s(discountList)
	return data, priceCount, priceList, discountCount, discountList
}

func getCoreMetrics(eventId uint64, data map[string]float32) (model.CoreMetric, error) {
	var metric = model.CoreMetric{}

	metric.ItemSold = int(data["item_count"])
	metric.OrderSold = int(data["order_sold"])
	metric.AmountSold = data["amount_sold"]
	totalItem, err := repo.GetTotalItemCountByEventId(setting.DB, eventId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return metric, err
	}
	conversion := float32(metric.ItemSold) / float32(totalItem)
	metric.Conversion = conversion

	return metric, nil
}

func getSecondaryMetrics(data map[string]float32) model.SecondaryMetric {
	var metric = model.SecondaryMetric{}

	metric.ReturnAmount = data["return_amount"]
	metric.AverageSku = data["average_sku"]
	metric.AverageItem = data["average_item"]
	metric.AverageAmount = data["average_amount"]
	metric.AveragePrice = data["average_price"]
	metric.AverageDiscount = data["average_discount"]
	metric.MaxDiscount = data["max_discount"]
	metric.MinDiscount = data["min_discount"]

	return metric
}

func getDistribution(priceCount, discountCount map[string]int, priceList []int, discountList []float64) model.Distribution {
	var distribution = model.Distribution{}
	var priceDistribution = make([]model.DistributionItem, 0)
	var discountDistribution = make([]model.DistributionItem, 0)

	for _, price := range priceList {
		var distributionItem = model.DistributionItem{}
		priceKey := fmt.Sprintf("%d", price)
		count := priceCount[priceKey]
		distributionItem.X = priceKey
		distributionItem.Y = fmt.Sprintf("%d", count)
		priceDistribution = append(priceDistribution, distributionItem)
	}

	for _, discount := range discountList {
		var distributionItem = model.DistributionItem{}
		discountKey := fmt.Sprintf("%f", discount)
		count := discountCount[discountKey]
		distributionItem.X = discountKey
		distributionItem.Y = fmt.Sprintf("%d", count)
		discountDistribution = append(discountDistribution, distributionItem)
	}

	distribution.PriceDistribution = priceDistribution
	distribution.DiscountDistribution = discountDistribution
	return distribution
}

func GetReportRanking(event model.Event, startTime, endTime, brand, source, dimension, sortBy, order string, page, pageSize int) (model.RankingResponse, error) {
	var reportResponse = model.RankingResponse{}
	var ranks = make([]model.Rank, 0)
	whereClause := generateWhereClause(event.Id, startTime, endTime, brand, source)
	groupBy := generateGroupByClause(dimension)
	sorts := getSortOrder(sortBy, order)
	offset := (page - 1) * pageSize
	selects := generateSelectByClause(groupBy)
	rankRecords, err := repo.GetRankItems(setting.DB, selects, whereClause, groupBy, sorts, offset, pageSize)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Ranking %s", err.Error()))
		return model.RankingResponse{}, err
	}
	var rankNumber = 1
	for _, record := range rankRecords {
		ranks = append(ranks, model.Rank{Rank: fmt.Sprintf("%d", rankNumber), Item: generateRankItem(record), Quantity: fmt.Sprintf("%d", record.Quantity)})
		rankNumber += 1
	}

	reportResponse.Ranks = ranks
	return reportResponse, nil
}

func generateRankItem(rank model.RankRecord) string {
	var itemNames = make([]string, 0)
	if len(rank.Brand) > 0 {
		itemNames = append(itemNames, rank.Brand)
	}

	if len(rank.Name) > 0 {
		itemNames = append(itemNames, rank.Name)
	}

	if len(rank.Sku) > 0 {
		itemNames = append(itemNames, rank.Sku)
	}

	if len(rank.Color) > 0 {
		itemNames = append(itemNames, rank.Color)
	}

	if len(rank.Category) > 0 {
		itemNames = append(itemNames, rank.Category)
	}

	if len(rank.Size) > 0 {
		itemNames = append(itemNames, rank.Size)
	}

	return strings.Join(itemNames, " ")
}

func getSortOrder(sortBy, order string) string {
	if len(sortBy) == 0 {
		sortBy = "quantity"
	}

	if len(order) == 0 {
		order = "desc"
	}

	sortOrder := fmt.Sprintf("%s %s", sortBy, order)
	return sortOrder
}

func generateGroupByClause(dimension string) string {
	dimensions := strings.Split(dimension, ",")
	groupBys := make([]string, 0)
	for _, d := range dimensions {
		if d == "sku" {
			groupBys = append(groupBys, "i.sku")
		}

		if d == "color" {
			groupBys = append(groupBys, "i.color")
		}

		if d == "category" {
			groupBys = append(groupBys, "i.category")
		}

		if d == "size" {
			groupBys = append(groupBys, "i.size")
		}
	}

	return strings.Join(groupBys, ",")
}

func generateSelectByClause(dimension string) string {
	dimensions := strings.Split(dimension, ",")
	selects := make([]string, 0)
	for _, d := range dimensions {
		if d == "sku" {
			selects = append(selects, "i.sku as sku")
		}

		if d == "color" {
			selects = append(selects, "i.color as color")
		}

		if d == "category" {
			selects = append(selects, "i.category as category")
		}

		if d == "size" {
			selects = append(selects, "i.size as size")
		}
	}
	selects = append(selects, "i.brand as brand")
	selects = append(selects, "i.name as name")
	selects = append(selects, "sum(s.quantity) as quantity")
	return strings.Join(selects, ",")
}

func GetReportDailyDetail(event model.Event, startTime, endTime, brand, source string) (model.DailyDetailResponse, error) {
	var reportResponse = model.DailyDetailResponse{}
	var dailyRecords = make([]model.DailyDetail, 0)
	whereClause := generateWhereClause(event.Id, startTime, endTime, brand, source)
	sortBy := "s.order_time"
	order := "desc"
	sorts := getSortOrder(sortBy, order)
	dailySaleDetail, err := repo.GetSoldItemDetailByEventIdWithOrder(setting.DB, whereClause, sorts)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return reportResponse, err
	}
	totalItem, err := repo.GetTotalItemCountByEventId(setting.DB, event.Id)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: total item %s", err.Error()))
		return reportResponse, err
	}

	processedData, dateList, err := processDailySaleRecord(dailySaleDetail, totalItem)

	for _, date := range dateList {
		var dailyRecord = model.DailyDetail{}
		dailyRecord.Date = date
		dailyRecord.ItemSold = int(processedData[date]["item_count"])
		dailyRecord.AmountSold = processedData[date]["amount_sold"]
		dailyRecord.OrderSold = int(processedData[date]["order_sold"])
		dailyRecord.ReturnAmount = processedData[date]["return_amount"]
		dailyRecord.Conversion = processedData[date]["conversion"]
		dailyRecord.Growth = processedData[date]["growth_to_yesterday"]
		dailyRecords = append(dailyRecords, dailyRecord)
	}

	reportResponse.Detail = dailyRecords
	return reportResponse, nil
}

func processDailySaleRecord(details []model.SaleRecordDetail, totalItem int64) (map[string]map[string]float32, []string, error) {
	var data = make(map[string]map[string]float32, 0)
	var dateList = make([]string, 0)
	var dateOrder = make(map[string]map[string]bool, 0)
	for _, detail := range details {
		date, err := time.Parse("2006-01-02", detail.OrderTime.String())
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: Date %s", err.Error()))
			continue
		}
		if _, ok := data[date.String()]; !ok {
			data[date.String()] = make(map[string]float32, 0)
			dateList = append(dateList, date.String())
		}

		dateData := data[date.String()]
		if detail.IsReturn == 1 {
			if _, ok := dateData["return_amount"]; !ok {
				dateData["return_amount"] = 0
			}
			dateData["return_amount"] += float32(detail.SalePrice)
			continue
		}

		if _, ok := dateData["item_count"]; !ok {
			dateData["item_count"] = 0
		}
		dateData["item_count"] += 1

		if _, ok := dateOrder[date.String()]; !ok {
			dateOrder[date.String()] = make(map[string]bool, 0)
		}
		dateOrder[date.String()][detail.OrderId] = true

		if _, ok := dateData["amount_sold"]; !ok {
			dateData["amount_sold"] = 0
		}
		dateData["amount_sold"] += float32(detail.SalePrice)
		data[date.String()] = dateData
	}

	sort.Strings(dateList)
	for i, date := range dateList {
		totalOrder := len(dateOrder[date])
		data[date]["order_sold"] = float32(totalOrder)
		conversion := data[date]["item_count"] / float32(totalItem)
		data[date]["conversion"] = conversion
		if i > 0 {
			yesterday := dateList[i-1]
			yesterdayConversion := data[yesterday]["conversion"]
			data[date]["growth_to_yesterday"] = (conversion - yesterdayConversion) / yesterdayConversion
		}
	}

	return data, dateList, nil
}
