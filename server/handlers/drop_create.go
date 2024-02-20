package handlers

import (
	"encoding/json"
	"errors"
	"github.com/vavilen84/nft-project/constants"
	"github.com/vavilen84/nft-project/dto"
	"github.com/vavilen84/nft-project/helpers"
	"github.com/vavilen84/nft-project/models"
	"github.com/vavilen84/nft-project/store"
	"github.com/vavilen84/nft-project/validation"
	"net/http"
)

func (c *DropController) Create(w http.ResponseWriter, r *http.Request) {

	u, ok := r.Context().Value("user").(*models.User)
	if !ok {
		err := errors.New("No logged in user")
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.UnauthorizedError, http.StatusUnauthorized)
		return
	}

	db := store.GetDB()
	dec := json.NewDecoder(r.Body)
	dtoModel := dto.Drop{}
	err := dec.Decode(&dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}
	err = validation.ValidateByScenario(constants.ScenarioCreate, dtoModel)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, constants.BadRequestError, http.StatusBadRequest)
		return
	}

	m := &models.Drop{
		CollectionName:     dtoModel.CollectionName,
		Blockchain:         models.Blockchain(dtoModel.Blockchain),
		PublicSaleDateTime: dtoModel.PublicSaleDateTime,
		TimeZone:           dtoModel.TimeZone,
		PublicSalePrice:    dtoModel.PublicSalePrice,
		TotalSupply:        dtoModel.TotalSupply,
		BillingPlan:        models.BillingPlan(dtoModel.BillingPlan),
		Status:             models.UnPublishedDropStatus,
		UserID:             u.Id,
		WebsiteURL:         dtoModel.WebsiteURL,
		TwitterURL:         dtoModel.TwitterURL,
		DiscordURL:         dtoModel.DiscordURL,
		BlockchainName:     dtoModel.BlockchainName,
	}

	err = models.InsertDrop(db, m)
	if err != nil {
		helpers.LogError(err)
		c.WriteErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	resp := make(dto.ResponseData)
	c.WriteSuccessResponse(w, resp, http.StatusOK)
}
