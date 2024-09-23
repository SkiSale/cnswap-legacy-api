package service

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/SkiSale/cnswap-legacy-api/domain/dao"
	"github.com/SkiSale/cnswap-legacy-api/domain/dto"
	"github.com/SkiSale/cnswap-legacy-api/lib"
	"github.com/SkiSale/cnswap-legacy-api/lib/constant"
	"github.com/SkiSale/cnswap-legacy-api/repository"
	"github.com/gin-gonic/gin"
)

type EquipmentService interface {
	GetAllEquipment(c *gin.Context)
	GetEquipmentByBarcode(c *gin.Context)
	GetAllEquipmentByVendor(c *gin.Context)
}

type EquipmentServiceImpl struct {
	EquipmentRepository repository.EquipmentRepository
}

func (u EquipmentServiceImpl) GetEquipmentByBarcode(c *gin.Context) {
	defer lib.PanicHandler(c)

	b := c.Param("barcode")

	data, err := u.EquipmentRepository.FindEquipmentByBarcode(b)
	if err != nil || data == nil {
		lib.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, lib.BuildResponse(constant.Success, mapEquipmentItem(data)))
}

func (u EquipmentServiceImpl) GetAllEquipment(c *gin.Context) {
	defer lib.PanicHandler(c)

	p, err := parsePagingParameters(c)
	if err != nil {
		lib.PanicException(constant.UnknownError)
	}

	slice, err := u.EquipmentRepository.FindAllEquipment(p.Offset, p.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.BuildResponse(constant.InvalidRequest, err))
		return
	}

	pagingHeaders(c, u.EquipmentRepository.Count(), p.PerPage)

	c.JSON(http.StatusOK,
		lib.BuildResponse(
			constant.Success,
			mapCollection(slice, mapEquipmentItem),
		),
	)
}

func (u EquipmentServiceImpl) GetAllEquipmentByVendor(c *gin.Context) {
	defer lib.PanicHandler(c)

	vendorID, err := strconv.Atoi(c.Param("vendorId"))
	if err != nil {
		lib.PanicException(constant.UnknownError)
	}

	p, err := parsePagingParameters(c)
	if err != nil {
		lib.PanicException(constant.UnknownError)
	}

	slice, err := u.EquipmentRepository.FindEquipmentByVendor(vendorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.BuildResponse(constant.InvalidRequest, err))
		return
	}

	pagingHeaders(c, u.EquipmentRepository.Count(), p.PerPage)

	c.JSON(http.StatusOK,
		lib.BuildResponse(
			constant.Success,
			mapCollection(slice, mapEquipmentItem),
		),
	)
}

func mapEquipmentItem(m *dao.Equipment) *dto.Equipment {
	return &dto.Equipment{
		Barcode:     m.Barcode,
		VendorId:    m.Owner,
		ItemTypeID:  m.ItemTypeID,
		Description: m.Description,
		ItemSize:    m.ItemSize,
		Price:       m.Price,
		DateSold:    m.DateSold,
		Stock: dto.EquipmentQuantities{
			Registered: int(m.QtyIn),
			CheckedIn:  int(m.QtyCheckedIn),
			Sold:       int(m.QtySold),
		},
		Flags: dto.EquipmentFlags{
			IsNew:    m.IsNew,
			IsRedTag: m.IsRedTag,
		},
		Transactions: []dto.EquipmentTransaction{
			{
				Type: "CR",
				Values: struct {
					Gross float64
					Tax1  float64
					Tax2  float64
				}{
					Gross: m.ValueSold,
					Tax1:  m.ValueTax1,
					Tax2:  m.ValueTax2,
				},
			},
			{
				Type: "DR",
				Values: struct {
					Gross float64
					Tax1  float64
					Tax2  float64
				}{
					Gross: m.PaidValue,
					Tax1:  m.PaidTax1,
					Tax2:  m.PaidTax2,
				},
			},
		},
		TagsPaidCount: m.TagsPaidCount,
		Notes:         strings.Split(m.Notes, "\r"),
	}
}

func EquipmentServiceInit(EquipmentRepository repository.EquipmentRepository) *EquipmentServiceImpl {
	return &EquipmentServiceImpl{
		EquipmentRepository: EquipmentRepository,
	}
}
