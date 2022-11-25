package service

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
	"strconv"
	"strings"
	"time"
)

func ReadEventItemFile(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	var items = make([]model.EventItem, 0)
	for r := 1; r < rowHeight; r++ {
		var item = model.EventItem{}
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
				item.Inventory, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item inventory conversion %s", err)
				}
				break
			case "J":
				item.RetailPrice, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item retail price conversion %s", err)
				}
				break
			case "K":
				item.SalePrice, err = strconv.Atoi(cellValue)
				if err != nil {
					fmt.Sprintf("Error: item sale price conversion %s", err)
				}
				break
			case "L":
				discount, err := strconv.ParseFloat(cellValue, 64)
				item.Discount = float32(discount)
				if err != nil {
					fmt.Sprintf("Error: item discount conversion %s", err)
				}
				break
			}
		}
		items = append(items, item)
	}

	err := repo.BatchInsertEventItems(setting.DB, items)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	return nil
}

func ReadEventSoldFile(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	var items = make([]model.SaleRecord, 0)
	var eventId uint64 = 0
	for r := 1; r < rowHeight; r++ {
		quantityStr := file.GetCellValue(model.SheetName, fmt.Sprintf("G%d", r))
		quantity, _ := strconv.Atoi(quantityStr)
		for i := 0; i < quantity; i++ {
			var item = model.SaleRecord{}
			var err error
			for c := 0; c < len(model.ItemColumns); c++ {
				cellName := fmt.Sprintf("%s%d", model.ItemColumns[c], r)
				cellValue := file.GetCellValue(model.SheetName, cellName)
				if len(cellValue) == 0 {
					continue
				}
				switch model.ItemColumns[c] {
				case "A":
					item.OrderId = cellValue
					break
				case "B":
					item.OrderTime, err = time.Parse("2006-01-02", cellValue)
					break
				case "G":
					item.Quantity = 1
					break
				case "I":
					item.PaidPrice, err = strconv.Atoi(cellValue)
					if err != nil {
						fmt.Sprintf("Error: item paid price conversion %s", err)
					}
					break
				case "J":
					item.Source, err = strconv.Atoi(cellValue)
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
			eventItem, err := repo.GetItemByItemDetail(setting.DB, itemWhereClause)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
				return errors.New(model.SqlQueryError)
			} else {
				item.ItemId = eventItem.Id
				item.CouponUsed = eventItem.SalePrice - item.PaidPrice
			}

			if eventId == 0 {
				eventId = eventItem.EventId
			}

			memberWhereClause := generateMemberWhereClause(customerName, customerPhone)
			memberId, err := repo.GetMemberIdByMemberDetail(setting.DB, memberWhereClause)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
				return errors.New(model.SqlQueryError)
			} else {
				item.MemberId = memberId
			}

			items = append(items, item)
		}
	}

	err := repo.BatchInsertEventSales(setting.DB, items)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Batch insert sales %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	err = repo.UpdateEventReportStatusByEventId(setting.DB, eventId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Update report status %s", err.Error()))
	}

	return nil
}

func generateItemWhereClause(barcode, sku, color, size, salePrice string) string {
	var whereClauses = make([]string, 0)
	if len(barcode) > 0 {
		whereClauses = append(whereClauses, barcode)
	}

	if len(sku) > 0 {
		whereClauses = append(whereClauses, sku)
	}

	if len(color) > 0 {
		whereClauses = append(whereClauses, color)
	}

	if len(size) > 0 {
		whereClauses = append(whereClauses, size)
	}

	if len(salePrice) > 0 {
		whereClauses = append(whereClauses, salePrice)
	}

	return strings.Join(whereClauses, " and ")
}

func generateMemberWhereClause(customerName, customerPhone string) string {
	var whereClauses = make([]string, 0)
	if len(customerName) > 0 {
		whereClauses = append(whereClauses, customerName)
	}

	if len(customerPhone) > 0 {
		whereClauses = append(whereClauses, customerPhone)
	}

	return strings.Join(whereClauses, " and ")
}

func ReadEventReturnFile(file *excelize.File) error {
	rows := file.GetRows(model.SheetName)
	rowHeight := len(rows)
	updateIds := make([]uint64, 0)
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

		saleRecords, err := repo.GetSaleRecordsByOrderId(setting.DB, orderId, paidPrice)

		var idCount = 0
		for _, s := range saleRecords {
			updateIds = append(updateIds, s.Id)
			idCount += 1
			if idCount == quantity {
				break
			}
		}

	}

	err := repo.UpdateReturnStatus(setting.DB, updateIds)
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

	err := repo.BatchInsertMembers(setting.DB, items)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return errors.New(model.SqlInsertionError)
	}

	return nil
}
