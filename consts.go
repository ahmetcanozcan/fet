package fet

const (
	// Equity Keywords

	// KeywordEq is "$eq" keyword
	KeywordEq string = "$eq"
	// KeywordEq is "$ne" keyword
	KeywordNe string = "$ne"

	// Array Keywords

	// KeywordIn is "$in" keyword
	KeywordIn string = "$in"
	// KeywordNin is "$nin" keyword
	KeywordNin string = "$nin"
	// KeywordAll is "$all" keyword
	KeywordAll string = "$all"
	// KeywordElemMatch is "$elemMatch" keyword
	KeywordElemMatch string = "$elemMatch"
	// KeywordAllElemMatch is "$allElemMatch" keyword
	KeywordAllElemMatch string = "$allElemMatch"

	// Logic Keywords

	// KeywordOr is "$or" keyword
	KeywordOr string = "$or"
	// KeywordAnd is "$and" keyword
	KeywordAnd string = "$and"
	// KeywordNor is "$nor" keyword
	KeywordNor string = "$nor"
	// KeywordNand is "$nand" keyword
	KeywordNand string = "$nand"

	// Comparison Keywords

	// KeywordGt is "$gt" keyword
	KeywordGt string = "$gt"
	// KeywordGte is "$gte" keyword
	KeywordGte string = "$gte"
	// KeywordLt is "$lt" keyword
	KeywordLt string = "$lt"
	// KeywordLte is "$lte" keyword
	KeywordLte string = "$lte"

	// Element Keywords

	// KeywordExists is "$exists" keyword
	KeywordExists string = "$exists"
	// KeywordSize is "$size" keyword
	KeywordSize string = "$size"
	// KeywordType is "$type" keyword
	KeywordType string = "$type"

	// Evaluation Keywords

	// KeywordMod is "$mod" keyword
	KeywordMod string = "$mod"
	// KeywordRegex is "$regex" keyword
	KeywordRegex string = "$regex"
	// KeywordOptions is "$options" keyword
	KeywordOptions string = "$options"
	// KeywordText is "$text" keyword
	KeywordText string = "$text"
	// KeywordSearch is "$search" keyword
	KeywordSearch string = "$search"
	// KeywordComment is "$comment" keyword
	KeywordComment string = "$comment"
	// KeywordMax is "$max" keyword
	KeywordMax string = "$max"
	// KeywordMin is "$min" keyword
	KeywordMin string = "$min"
	// KeywordMaxDistance is "$maxDistance" keyword
	KeywordMaxDistance string = "$maxDistance"

	// Bitwise Keywords

	// KeywordBitwiseAnd is "$bitwiseAnd" keyword
	KeywordBitwiseAnd string = "$bitwiseAnd"
	// KeywordBitwiseOr is "$bitwiseOr" keyword
	KeywordBitwiseOr string = "$bitwiseOr"
	// KeywordBitwiseXor is "$bitwiseXor" keyword
	KeywordBitwiseXor string = "$bitwiseXor"
	// KeywordBitwiseNot is "$bitwiseNot" keyword
	KeywordBitwiseNot string = "$bitwiseNot"

	// Geospatial Keywords

	// KeywordGeoIntersects is "$geoIntersects" keyword
	KeywordGeoIntersects string = "$geoIntersects"
	// KeywordGeoWithin is "$geoWithin" keyword
	KeywordGeoWithin string = "$geoWithin"
	// KeywordNear is "$near" keyword
	KeywordNear string = "$near"
	// KeywordWithin is "$within" keyword
	KeywordWithin string = "$within"

	// KeywordProjection is "$projection" keyword
	KeywordProjection string = "$projection"

	// Update Keywords

	// KeywordUnset is "$unset" keyword
	KeywordUnset string = "$unset"
	// KeywordSet is "$set" keyword
	KeywordSet string = "$set"
	// KeywordSetOnInsert is "$setOnInsert" keyword
	KeywordSetOnInsert string = "$setOnInsert"

	// Array Keywords

	// KeywordArrayFilters is "$arrayFilters" keyword
	KeywordArrayFilters string = "$arrayFilters"

	// Text Keywords

	// KeywordTextScore is "$textScore" keyword
	KeywordTextScore string = "$textScore"

	// KeywordCurrentDate is "$currentDate" keyword
	KeywordCurrentDate string = "$currentDate"
)
