package service

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/starter/component"
	"rackrock/utils"
	"strconv"
	"strings"
	"time"
)

func ReadEventItemFile(file *excelize.File, eventId string) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	var items = make([]model.EventItem, 0)

	for r := 2; r <= rowHeight; r++ {
		var item = model.EventItem{}
		item.EventId, _ = utils.ConvertStringToUint64(eventId)
		var err error
		for c := 0; c < len(model.ItemColumns); c++ {
			cellName := fmt.Sprintf("%s%d", model.ItemColumns[c], r)

			cellValue := file.GetCellValue(model.SheetName, cellName)
			if len(cellValue) == 0 {
				continue
			}
			switch model.ItemColumns[c] {
			case "A":
				item.Brand = cellValue
				break
			case "B":
				item.Sku = cellValue
				break
			case "C":
				item.Barcode = cellValue
				break
			case "D":
				item.Name = cellValue
				break
			case "E":
				item.Season = cellValue
				break
			case "F":
				item.Category = cellValue
				break
			case "G":
				item.Color = cellValue
				break
			case "H":
				item.Size = cellValue
				break
			case "I":
				item.Gender = cellValue
				break
			case "J":
				item.Inventory, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item inventory conversion %s", err)
				}
				break
			case "K":
				item.RetailPrice, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item retail price conversion %s", err)
				}
				break
			case "L":
				value, err := strconv.ParseFloat(cellValue, 32)
				item.SalePrice = float32(value)
				if err != nil {
					fmt.Sprintf("Error: item sale price conversion %s", err)
				}
				break
			case "M":
				discount, err := strconv.ParseFloat(cellValue, 64)
				fmt.Println(fmt.Sprintf("discount %f", discount))
				item.Discount = float32(discount)
				if err != nil {
					fmt.Sprintf("Error: item discount conversion %s", err)
				}
				break
			}
		}

		items = append(items, item)
	}

	err := repo.BatchInsertEventItems(component.DB, items)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	return nil
}

func ReadEventSoldFile(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	var sales = make([]model.SaleRecord, 0)
	var eventId uint64 = 0
	var notImported = 0
	for r := 1; r < rowHeight; r++ {
		quantityStr := file.GetCellValue(model.SheetName, fmt.Sprintf("G%d", r))
		quantity, _ := strconv.Atoi(quantityStr)
		for i := 0; i < quantity; i++ {
			var sale = model.SaleRecord{}
			var err error
			for c := 0; c < len(model.ItemColumns); c++ {
				cellName := fmt.Sprintf("%s%d", model.ItemColumns[c], r)
				cellValue := file.GetCellValue(model.SheetName, cellName)
				if len(cellValue) == 0 {
					continue
				}
				switch model.ItemColumns[c] {
				case "A":
					sale.OrderId = cellValue
					break
				case "B":
					sale.OrderTime, _ = time.Parse("2006-01-02 15:04:05", cellValue)
					break
				case "G":
					sale.Quantity = 1
					break
				case "I":
					convertedPaid, err := strconv.ParseFloat(cellValue, 32)
					if err != nil {
						fmt.Sprintf("Error: item paid price conversion %s", err)
					}
					sale.PaidPrice = float32(convertedPaid) / float32(quantity)
					break
				case "J":
					sale.Source, err = strconv.Atoi(cellValue)
					if err != nil {
						fmt.Sprintf("Error: item inventory conversion %s", err)
					}
					break
				}
			}
			barcode := file.GetCellValue(model.SheetName, fmt.Sprintf("C%d", r))
			sku := file.GetCellValue(model.SheetName, fmt.Sprintf("D%d", r))
			color := file.GetCellValue(model.SheetName, fmt.Sprintf("E%d", r))
			size := file.GetCellValue(model.SheetName, fmt.Sprintf("F%d", r))
			salePrice := file.GetCellValue(model.SheetName, fmt.Sprintf("H%d", r))
			customerName := file.GetCellValue(model.SheetName, fmt.Sprintf("K%d", r))
			customerPhone := file.GetCellValue(model.SheetName, fmt.Sprintf("L%d", r))

			itemWhereClause := generateItemWhereClause(barcode, sku, color, size, salePrice)
			eventItem, err := repo.GetItemByItemDetail(component.DB, itemWhereClause)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
				notImported += 1
				continue
			} else {
				sale.ItemId = eventItem.Id
				sale.CouponUsed = eventItem.SalePrice - sale.PaidPrice
			}

			if eventId == 0 {
				eventId = eventItem.EventId
			}

			memberWhereClause := generateMemberWhereClause(customerName, customerPhone)
			if len(memberWhereClause) > 0 {
				memberId, err := repo.GetMemberIdByMemberDetail(component.DB, memberWhereClause)
				if err != nil {
					fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
				} else {
					sale.MemberId = memberId
				}
			}

			sales = append(sales, sale)
		}
	}

	fmt.Println(fmt.Sprintf("Not imported: %d", notImported))

	err := repo.BatchInsertEventSales(component.DB, sales)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Batch insert sales %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	err = repo.UpdateEventReportStatusByEventId(component.DB, eventId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Update report status %s", err.Error()))
	}

	return nil
}

func generateItemWhereClause(barcode, sku, color, size, salePrice string) string {
	var whereClauses = make([]string, 0)
	if len(barcode) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("barcode = '%s'", barcode))
	}

	if len(sku) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("sku = '%s'", sku))
	}

	if len(color) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("color = '%s'", color))
	}

	if len(size) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("size = '%s'", size))
	}

	if len(salePrice) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("sale_price like '%s'", salePrice))
	}

	return strings.Join(whereClauses, " and ")
}

func generateMemberWhereClause(customerName, customerPhone string) string {
	var whereClauses = make([]string, 0)
	if len(customerName) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("name = '%s'", customerName))
	}

	if len(customerPhone) > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("phone = '%s'", customerPhone))
	}

	return strings.Join(whereClauses, " and ")
}

func ReadEventReturnFile(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	updateIds := make([]uint64, 0)
	added := make(map[uint64]bool, 0)
	for r := 1; r < rowHeight; r++ {
		var orderId string
		var paidPrice string
		var quantity int
		var err error
		for c := 0; c < len(model.ReturnColumns); c++ {
			cellName := fmt.Sprintf("%s%d", model.ReturnColumns[c], r)
			cellValue := file.GetCellValue(model.SheetName, cellName)
			if len(cellValue) == 0 {
				continue
			}
			switch model.ItemColumns[c] {
			case "A":
				orderId = cellValue
				break
			case "B":
				paidPrice = cellValue
				break
			case "C":
				quantity, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item gender conversion %s", err)
				}
				break
			}
		}

		saleRecords, err := repo.GetSaleRecordsByOrderId(component.DB, orderId, paidPrice)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: Get sale records error  %s", err.Error()))
			return errors.New(model.SqlQueryError)
		}

		var idCount = 0
		for _, s := range saleRecords {
			if _, ok := added[s.Id]; ok {
				continue
			}
			updateIds = append(updateIds, s.Id)
			added[s.Id] = true
			idCount += 1
			if idCount == quantity {
				break
			}
		}

	}

	err := repo.UpdateReturnStatus(component.DB, updateIds)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: update return status %s", err.Error()))
		return errors.New(model.SqlUpdateError)
	}

	return nil
}

func ReadMember(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	var items = make([]model.Member, 0)
	for r := 1; r < rowHeight; r++ {
		var item = model.Member{}
		var err error
		for c := 0; c < len(model.MemberColumns); c++ {
			cellName := fmt.Sprintf("%s%d", model.MemberColumns[c], r)
			cellValue := file.GetCellValue(model.SheetName, cellName)
			if len(cellValue) == 0 {
				continue
			}
			switch model.ItemColumns[c] {
			case "A":
				item.Name = cellValue
				break
			case "B":
				item.Nickname = cellValue
				break
			case "C":
				item.Gender, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item gender conversion %s", err)
				}
				break
			case "D":
				item.Phone = cellValue
				break
			case "E":
				item.City = cellValue
				break
			case "F":
				item.Dob, err = time.Parse("2006-01-02", cellValue)
				if err != nil {
					fmt.Sprintf("Error: item dob conversion %s", err)
				}
				break
			}
		}

		items = append(items, item)
	}

	err := repo.BatchInsertMembers(component.DB, items)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	return nil
}
