package seller

type Notifier interface {
	StockChanged(oldStock int, newStock int)
}
