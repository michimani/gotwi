package fields

type PlaceField string

const (
	ContainedWithin PlaceField = "contained_within"
	Country         PlaceField = "country"
	CountryCode     PlaceField = "country_code"
	FullName        PlaceField = "full_name"
	PlaceGeo        PlaceField = "geo"
	PlaceID         PlaceField = "id"
	PlaceName       PlaceField = "name"
	PlaceType       PlaceField = "place_type"
)
