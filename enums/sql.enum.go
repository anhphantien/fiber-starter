package enums

type __Order struct {
	ASC  string
	DESC string
}

var _Order = __Order{
	ASC:  "ASC",
	DESC: "DESC",
}

type _Sort struct {
	Order __Order
}

var Sort = _Sort{
	Order: _Order,
}
