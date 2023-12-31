package contract

const IDKey = "bxd:id"

type IDService interface {
	NewID() string
}
