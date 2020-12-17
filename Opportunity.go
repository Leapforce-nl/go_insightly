package insightly

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

// Opportunity stores Opportunity from Insightly
//
type Opportunity struct {
	OpportunityID       int          `json:"OPPORTUNITY_ID"`
	OpportunityName     string       `json:"OPPORTUNITY_NAME"`
	OpportunityDetails  string       `json:"OPPORTUNITY_DETAILS"`
	OpportunityState    string       `json:"OPPORTUNITY_STATE"`
	ResponsibleUserID   int          `json:"RESPONSIBLE_USER_ID"`
	CategoryID          int          `json:"CATEGORY_ID"`
	ImageURL            string       `json:"IMAGE_URL"`
	BidCurrency         string       `json:"BID_CURRENCY"`
	BidAmount           float32      `json:"BID_AMOUNT"`
	BidType             string       `json:"BID_TYPE"`
	BidDuration         int          `json:"BID_DURATION"`
	ActualCloseDate     DateUTC      `json:"ACTUAL_CLOSE_DATE"`
	DateCreatedUTC      DateUTC      `json:"DATE_CREATED_UTC"`
	DateUpdatedUTC      DateUTC      `json:"DATE_UPDATED_UTC"`
	OpportunityValue    float32      `json:"OPPORTUNITY_VALUE"`
	Probability         int          `json:"PROBABILITY"`
	ForecastCloseDate   DateUTC      `json:"FORECAST_CLOSE_DATE"`
	OwnerUserID         int          `json:"OWNER_USER_ID"`
	LastActivityDateUTC DateUTC      `json:"LAST_ACTIVITY_DATE_UTC"`
	NextActivityDateUTC DateUTC      `json:"NEXT_ACTIVITY_DATE_UTC"`
	PipelineID          int          `json:"PIPELINE_ID"`
	StageID             int          `json:"STAGE_ID"`
	CreatedUserID       int          `json:"CREATED_USER_ID"`
	OrganisationID      int          `json:"ORGANISATION_ID"`
	CustomFields        CustomFields `json:"CUSTOMFIELDS"`
	Tags                []Tag        `json:"TAGS"`
}

func (o *Opportunity) prepareMarshal() interface{} {
	if o == nil {
		return nil
	}

	return &struct {
		OpportunityID      int           `json:"OPPORTUNITY_ID"`
		OpportunityName    string        `json:"OPPORTUNITY_NAME"`
		OpportunityDetails string        `json:"OPPORTUNITY_DETAILS"`
		OpportunityState   string        `json:"OPPORTUNITY_STATE"`
		ResponsibleUserID  int           `json:"RESPONSIBLE_USER_ID"`
		CategoryID         int           `json:"CATEGORY_ID"`
		ImageURL           string        `json:"IMAGE_URL"`
		BidCurrency        string        `json:"BID_CURRENCY"`
		BidAmount          float32       `json:"BID_AMOUNT"`
		BidType            string        `json:"BID_TYPE"`
		BidDuration        int           `json:"BID_DURATION"`
		ActualCloseDate    DateUTC       `json:"ACTUAL_CLOSE_DATE"`
		OpportunityValue   float32       `json:"OPPORTUNITY_VALUE"`
		Probability        int           `json:"PROBABILITY"`
		ForecastCloseDate  DateUTC       `json:"FORECAST_CLOSE_DATE"`
		OwnerUserID        int           `json:"OWNER_USER_ID"`
		PipelineID         int           `json:"PIPELINE_ID"`
		StageID            int           `json:"STAGE_ID"`
		OrganisationID     int           `json:"ORGANISATION_ID"`
		CustomFields       []CustomField `json:"CUSTOMFIELDS"`
	}{
		o.OpportunityID,
		o.OpportunityName,
		o.OpportunityDetails,
		o.OpportunityState,
		o.ResponsibleUserID,
		o.CategoryID,
		o.ImageURL,
		o.BidCurrency,
		o.BidAmount,
		o.BidType,
		o.BidDuration,
		o.ActualCloseDate,
		o.OpportunityValue,
		o.Probability,
		o.ForecastCloseDate,
		o.OwnerUserID,
		o.PipelineID,
		o.StageID,
		o.OrganisationID,
		o.CustomFields,
	}
}

// GetOpportunity returns a specific opportunity
//
func (i *Insightly) GetOpportunity(opportunityID int) (*Opportunity, *errortools.Error) {
	endpoint := fmt.Sprintf("Opportunities/%v", opportunityID)

	opportunity := Opportunity{}

	_, _, e := i.get(endpoint, nil, &opportunity)
	if e != nil {
		return nil, e
	}

	return &opportunity, nil
}

type GetOpportunitiesFilter struct {
	UpdatedAfter *time.Time
	Field        *struct {
		FieldName  string
		FieldValue string
	}
}

// GetOpportunities returns all opportunities
//
func (i *Insightly) GetOpportunities(filter *GetOpportunitiesFilter) (*[]Opportunity, *errortools.Error) {
	searchString := "?"
	searchFilter := []string{}

	if filter != nil {
		if filter.UpdatedAfter != nil {
			from := filter.UpdatedAfter.Format(ISO8601Format)
			searchFilter = append(searchFilter, fmt.Sprintf("updated_after_utc=%s&", from))
		}

		if filter.Field != nil {
			searchFilter = append(searchFilter, fmt.Sprintf("field_name=%s&field_value=%s&", filter.Field.FieldName, filter.Field.FieldValue))
		}
	}

	if len(searchFilter) > 0 {
		searchString = "/Search?" + strings.Join(searchFilter, "&")
	}

	endpointStr := "Opportunities%sskip=%s&top=%s"
	skip := 0
	top := 100
	rowCount := top

	opportunities := []Opportunity{}

	for rowCount >= top {
		endpoint := fmt.Sprintf(endpointStr, searchString, strconv.Itoa(skip), strconv.Itoa(top))
		//fmt.Println(endpoint)

		cs := []Opportunity{}

		_, _, e := i.get(endpoint, nil, &cs)
		if e != nil {
			return nil, e
		}

		opportunities = append(opportunities, cs...)

		rowCount = len(cs)
		//rowCount = 0
		skip += top
	}

	if len(opportunities) == 0 {
		opportunities = nil
	}

	return &opportunities, nil
}

// CreateOpportunity creates a new contract
//
func (i *Insightly) CreateOpportunity(opportunity *Opportunity) (*Opportunity, *errortools.Error) {
	if opportunity == nil {
		return nil, nil
	}

	endpoint := "Opportunities"

	opportunityNew := Opportunity{}

	_, _, e := i.post(endpoint, opportunity.prepareMarshal(), &opportunityNew)
	if e != nil {
		return nil, e
	}

	return &opportunityNew, nil
}

// UpdateOpportunity updates an existing contract
//
func (i *Insightly) UpdateOpportunity(opportunity *Opportunity) (*Opportunity, *errortools.Error) {
	if opportunity == nil {
		return nil, nil
	}

	endpoint := "Opportunities"

	opportunityUpdated := Opportunity{}

	_, _, e := i.put(endpoint, opportunity.prepareMarshal(), &opportunityUpdated)
	if e != nil {
		return nil, e
	}

	return &opportunityUpdated, nil
}

// DeleteOpportunity deletes a specific opportunity
//
func (i *Insightly) DeleteOpportunity(opportunityID int) *errortools.Error {
	endpoint := fmt.Sprintf("Opportunities/%v", opportunityID)

	_, _, e := i.delete(endpoint, nil, nil)
	if e != nil {
		return e
	}

	return nil
}

// GetOpportunityLinks returns links for a specific opportunity
//
func (i *Insightly) GetOpportunityLinks(opportunityID int) (*[]Link, *errortools.Error) {
	endpoint := fmt.Sprintf("Opportunity/%v/Links", opportunityID)

	links := []Link{}

	_, _, e := i.get(endpoint, nil, &links)
	if e != nil {
		return nil, e
	}

	return &links, nil
}