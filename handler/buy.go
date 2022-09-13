package handler

import (
	"net/http"
	"time"

	"fmt"
	"mini-project/buy"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

type buyHandler struct {
	buyService buy.Service
}

func NewBuyHandler(buyService buy.Service) *buyHandler {
	return &buyHandler{buyService}
}

func (h *buyHandler) GetPembelian(c *gin.Context) {
	pembelians, err := h.buyService.GetPembelian()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var pembeliansResponse []buy.Pembelian

	for _, u := range pembelians {
		pembelianResponse := convertToPembelianResponse(u)
		pembeliansResponse = append(pembeliansResponse, pembelianResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pembeliansResponse,
	})
}

func (h *buyHandler) GetBuys(c *gin.Context) {
	buys, err := h.buyService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var buysResponse []buy.BuyResponse

	for _, u := range buys {
		buyResponse := convertToBuyResponse(u)
		buysResponse = append(buysResponse, buyResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": buysResponse,
	})
}

func (h *buyHandler) GetBuy(c *gin.Context) {
	idString := c.Param("id")
	//  id, _ := strconv.Atoi(idString)
	id, _ := uuid.Parse(idString)

	u, err := h.buyService.FindByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBuyResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *buyHandler) CreateBuy(c *gin.Context) {
	var buyRequest buy.BuysRequest

	err := c.ShouldBindJSON(&buyRequest)

	if err != nil {

		// log.Fatal(err)
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		// fmt.Println(err)
		return

	}
	buyRequest.Time = time.Now()
	buy, err := h.buyService.Create(buyRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBuyResponse(buy),
		// "title": bookInput.Title,
		// "price": bookInput.Price,
		// "sub_title": bookInput.SubTitle,
	})
}

func (h *buyHandler) UpdateBuy(c *gin.Context) {
	var buyRequest buy.BuysRequest

	err := c.ShouldBindJSON(&buyRequest)

	if err != nil {

		// log.Fatal(err)
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
		})
		// fmt.Println(err)
		return

	}

	idString := c.Param("id")
	// id, _ := strconv.Atoi(idString)
	id, _ := uuid.Parse(idString)
	buy, err := h.buyService.Update(id, buyRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBuyResponse(buy),
		// "title": bookInput.Title,
		// "price": bookInput.Price,
		// "sub_title": bookInput.SubTitle,
	})
}

func (h *buyHandler) DeleteBuy(c *gin.Context) {
	idString := c.Param("id")
	// id, _ := strconv.Atoi(idString)
	id, _ := uuid.Parse(idString)

	u, err := h.buyService.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	userResponse := convertToBuyResponse(u)

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}

func convertToBuyResponse(b buy.Buy) buy.BuyResponse {
	return buy.BuyResponse{
		ID:     b.ID,
		IdUser: b.IdUser,
		IdBuku: b.IdBuku,
		Jumlah: b.Jumlah,
		Time:   b.Time,
	}
}
func convertToPembelianResponse(p buy.Pembelian) buy.Pembelian {
	return buy.Pembelian{
		Id_user:       p.Id_user,
		Name:          p.Name,
		Title:         p.Title,
		Description:   p.Description,
		Total_price:   p.Total_price,
		Book_quantity: p.Book_quantity,
		Payment_date:  p.Payment_date,
	}
}
