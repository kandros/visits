package visit

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/kandros/visits/internal/dynamo"
	"github.com/kandros/visits/internal/util"
)

type Visit struct {
	VisitID   string `json:"visit_id"`
	URL       string `json:"url"`
	Timestamp int64  `json:"timestamp"`
	IP        string `json:"ip"`
	CountryID string `json:"country_id"`
	IsMobile  bool   `json:"is_mobile"`
	IsTablet  bool   `json:"is_tablet"`
}

func NewFromRequest(r *http.Request) Visit {
	ip, _ := util.MustGetIp(r)
	countryID := r.Header["Cloudfront-Viewer-Country"][0]
	isTablet, _ := strconv.ParseBool(r.Header["Cloudfront-Is-Tablet-Viewer"][0])
	isMobile, _ := strconv.ParseBool(r.Header["Cloudfront-Is-Mobile-Viewer"][0])

	visitID := uuid.New().String()
	timestamp := time.Now().Unix()

	v := Visit{
		VisitID:   visitID,
		URL:       "https://google.com",
		Timestamp: timestamp,
		IP:        ip,
		CountryID: countryID,
		IsMobile:  isMobile,
		IsTablet:  isTablet,
	}

	return v
}

func (v Visit) Persist() error {
	_, err := dynamo.Store(v)
	if err != nil {
		return err
	}

	return nil
}
