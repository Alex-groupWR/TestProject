package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StringToNullTime(dateString string) (sql.NullTime, error) {
	layout := "02.01.2006"
	if dateString == "" {
		return sql.NullTime{Valid: false}, nil
	}

	t, err := time.Parse(layout, dateString)
	if err != nil {
		return sql.NullTime{}, err
	}

	return sql.NullTime{
		Time:  t,
		Valid: true,
	}, nil
}

func NullTimeToString(nullTime sql.NullTime, ctx *gin.Context) string {
	layout := "02.01.2006"
	var timeString string
	if nullTime.Valid {
		timeString = nullTime.Time.Format(layout)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parse time to string"})
	}

	return timeString
}
