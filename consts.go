package fet

const (
	KeywordEq string = "$eq"
	KeywordNe string = "$ne"

	// Array
	KeywordIn           string = "$in"
	KeywordNin          string = "$nin"
	KeywordAll          string = "$all"
	KeywordElemMatch    string = "$elemMatch"
	KeywordAllElemMatch string = "$allElemMatch"

	// Logic
	KeywordOr   string = "$or"
	KeywordAnd  string = "$and"
	KeywordNor  string = "$nor"
	KeywordNand string = "$nand"

	// Comparison
	KeywordGt  string = "$gt"
	KeywordGte string = "$gte"
	KeywordLt  string = "$lt"
	KeywordLte string = "$lte"

	// Element
	KeywordExists string = "$exists"
	KeywordSize   string = "$size"
	KeywordType   string = "$type"

	// Evaluation
	KeywordMod         string = "$mod"
	KeywordRegex       string = "$regex"
	KeywordOptions     string = "$options"
	KeywordText        string = "$text"
	KeywordSearch      string = "$search"
	KeywordComment     string = "$comment"
	KeywordMax         string = "$max"
	KeywordMin         string = "$min"
	KeywordMaxDistance string = "$maxDistance"

	// Bitwise
	KeywordBitwiseAnd string = "$bitwiseAnd"
	KeywordBitwiseOr  string = "$bitwiseOr"
	KeywordBitwiseXor string = "$bitwiseXor"
	KeywordBitwiseNot string = "$bitwiseNot"

	// Geospatial
	KeywordGeoIntersects string = "$geoIntersects"
	KeywordGeoWithin     string = "$geoWithin"
	KeywordNear          string = "$near"
	KeywordWithin        string = "$within"

	// Projection
	KeywordProjection string = "$projection"

	// Update
	KeywordUnset       string = "$unset"
	KeywordSet         string = "$set"
	KeywordSetOnInsert string = "$setOnInsert"

	// Array
	KeywordArrayFilters string = "$arrayFilters"

	// Text
	KeywordTextScore string = "$textScore"

	KeywordCurrentDate string = "$currentDate"
)
