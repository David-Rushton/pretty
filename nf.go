package pretty

type prettyNFFormat int

const (
	DoubleHeightTop    prettyNFFormat = 3
	DoubleHeightBottom prettyNFFormat = 4
	SingleWidth        prettyNFFormat = 5
	DoubleWidth        prettyNFFormat = 6
)

func WithDoubleHeightTop() prettyNFFormat {
	return DoubleHeightTop
}

func WithDoubleHeightBottom() prettyNFFormat {
	return DoubleHeightBottom
}

func WithDoubleWidth() prettyNFFormat {
	return DoubleWidth
}

func WithSingleWidth() prettyNFFormat {
	return SingleWidth
}
