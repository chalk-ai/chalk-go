package fixtures

// WideOffer and WideRow mimic the shape of a wide bulk response: one
// namespace with many mostly-dense float columns, a has-one with more floats,
// a few sparse integer counters and some bool/int columns. Names are generic.
type WideOffer struct {
	Rate0      *float64 `json:"rate_0"`
	Rate1      *float64 `json:"rate_1"`
	Rate2      *float64 `json:"rate_2"`
	Rate3      *float64 `json:"rate_3"`
	Rate4      *float64 `json:"rate_4"`
	Rate5      *float64 `json:"rate_5"`
	Rate6      *float64 `json:"rate_6"`
	Rate7      *float64 `json:"rate_7"`
	Rate8      *float64 `json:"rate_8"`
	Rate9      *float64 `json:"rate_9"`
	Rate10     *float64 `json:"rate_10"`
	Rate11     *float64 `json:"rate_11"`
	Rate12     *float64 `json:"rate_12"`
	Rate13     *float64 `json:"rate_13"`
	Rate14     *float64 `json:"rate_14"`
	Rate15     *float64 `json:"rate_15"`
	Rate16     *float64 `json:"rate_16"`
	Rate17     *float64 `json:"rate_17"`
	Rate18     *float64 `json:"rate_18"`
	Rate19     *float64 `json:"rate_19"`
	Restricted *bool    `json:"restricted"`
}

type WideRow struct {
	Id        *string    `json:"id"`
	Score0    *float64   `json:"score_0"`
	Score1    *float64   `json:"score_1"`
	Score2    *float64   `json:"score_2"`
	Score3    *float64   `json:"score_3"`
	Score4    *float64   `json:"score_4"`
	Score5    *float64   `json:"score_5"`
	Score6    *float64   `json:"score_6"`
	Score7    *float64   `json:"score_7"`
	Score8    *float64   `json:"score_8"`
	Score9    *float64   `json:"score_9"`
	Score10   *float64   `json:"score_10"`
	Score11   *float64   `json:"score_11"`
	Score12   *float64   `json:"score_12"`
	Score13   *float64   `json:"score_13"`
	Score14   *float64   `json:"score_14"`
	Score15   *float64   `json:"score_15"`
	Score16   *float64   `json:"score_16"`
	Score17   *float64   `json:"score_17"`
	Score18   *float64   `json:"score_18"`
	Score19   *float64   `json:"score_19"`
	Score20   *float64   `json:"score_20"`
	Score21   *float64   `json:"score_21"`
	Score22   *float64   `json:"score_22"`
	Score23   *float64   `json:"score_23"`
	Score24   *float64   `json:"score_24"`
	Score25   *float64   `json:"score_25"`
	Score26   *float64   `json:"score_26"`
	Score27   *float64   `json:"score_27"`
	Score28   *float64   `json:"score_28"`
	Score29   *float64   `json:"score_29"`
	Count0    *int64     `json:"count_0"`
	Count1    *int64     `json:"count_1"`
	Count2    *int64     `json:"count_2"`
	Count3    *int64     `json:"count_3"`
	Count4    *int64     `json:"count_4"`
	Count5    *int64     `json:"count_5"`
	Count6    *int64     `json:"count_6"`
	Count7    *int64     `json:"count_7"`
	Count8    *int64     `json:"count_8"`
	DayOfWeek *int64     `json:"day_of_week"`
	Offer     *WideOffer `json:"offer"`
}
