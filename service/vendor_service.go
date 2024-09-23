package service

import (
	"net/http"
	"strconv"

	"github.com/SkiSale/cnswap-legacy-api/domain/dao"
	"github.com/SkiSale/cnswap-legacy-api/domain/dto"
	"github.com/SkiSale/cnswap-legacy-api/lib"
	"github.com/SkiSale/cnswap-legacy-api/lib/constant"
	"github.com/SkiSale/cnswap-legacy-api/repository"
	"github.com/gin-gonic/gin"
)

type VendorService interface {
	GetAllVendors(c *gin.Context)
	GetVendorById(c *gin.Context)
}

type VendorServiceImpl struct {
	VendorRepository repository.VendorRepository
}

func (u VendorServiceImpl) GetVendorById(c *gin.Context) {
	defer lib.PanicHandler(c)

	vendorID, _ := strconv.Atoi(c.Param("vendorId"))

	data, err := u.VendorRepository.FindVendorById(vendorID)
	if err != nil || data == nil {
		lib.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, lib.BuildResponse(constant.Success, mapVendor(data)))
}

func (u VendorServiceImpl) GetAllVendors(c *gin.Context) {
	defer lib.PanicHandler(c)

	p, err := parsePagingParameters(c)
	if err != nil {
		lib.PanicException(constant.UnknownError)
	}

	slice, err := u.VendorRepository.FindAllVendors(p.Offset, p.PerPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.BuildResponse(constant.InvalidRequest, err))
		return
	}

	pagingHeaders(c, u.VendorRepository.Count(), p.PerPage)

	c.JSON(http.StatusOK, lib.BuildResponse(constant.Success, mapCollection(slice, mapVendor)))
}

func mapVendor(v *dao.Vendor) *dto.Vendor {
	return &dto.Vendor{
		ID:           v.ID,
		VendorTypeID: v.VendorTypeID,
		LastName:     v.LastName,
		FirstName:    v.FirstName,
		Address: dto.VendorAddress{
			LineOne:  v.AddressLineOne,
			LineTwo:  v.AddressLineTwo,
			City:     v.AddressCity,
			Province: v.AddressProvince,
			Postcode: v.AddressPostcode,
		},
	}
}

func VendorServiceInit(VendorRepository repository.VendorRepository) *VendorServiceImpl {
	return &VendorServiceImpl{
		VendorRepository: VendorRepository,
	}
}
