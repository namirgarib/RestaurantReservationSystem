package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	InvoiceID      string             `json:"invoice_id"`
	OrderID        string             `json:"order_id"`
	PaymentMethod  *string            `json:"payment_method" validate:"eq=CREDIT_CARD|eq=DEBIT_CARD|eq=CASH|eq="`
	PaymentStatus  *string            `json:"payment_status" validate:"eq=PAID|eq=PENDING"`
	PaymentDueDate time.Time          `json:"payment_due_date"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
