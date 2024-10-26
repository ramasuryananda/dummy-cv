package constant

const DefaultTax = 10
const (
	InvoiceUnpaid int8 = iota
	InvoicePaid
	InvoiceCanceled
)

var InvoiceStatus = map[int8]string{
	InvoiceUnpaid:   "Unpaid",
	InvoicePaid:     "Paid",
	InvoiceCanceled: "Canceled",
}
