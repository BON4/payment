CREATE TYPE transaction_status AS ENUM('accepted', 'declined');
CREATE TYPE payment_type AS ENUM('cash', 'card');

CREATE TABLE transacton_history (
	"TransactionId" BIGSERIAL NOT NULL unique PRIMARY KEY, 
	"RequestId" BIGINT NOT NULL, 
	"TerminalId" BIGINT NOT NULL, 
	"PartnerObjectId" BIGINT NOT NULL, 
	"AmountTotal" DECIMAL(12,2) NOT NULL, 
	"AmountOriginal" DECIMAL(12, 2) NOT NULL, 
	"CommissionPS" DECIMAL(12, 2) NOT NULL, 
	"CommissionClient" DECIMAL(12, 2) NOT NULL, 
	"CommissionProvider" DECIMAL(12, 2) NOT NULL, 
	"DateInput"  TIMESTAMP WITHOUT TIME ZONE NOT NULL, 
	"DatePost"  TIMESTAMP WITHOUT TIME ZONE NOT NULL, 
	"Status" transaction_status NOT NULL, 
	"PaymentType" payment_type NOT NULL, 
	"PaymentNumber" TEXT NOT NULL, 
	"ServiceId" BIGINT NOT NULL, 
	"Service" TEXT NOT NULL, 
	"PayeeId" BIGINT NOT NULL, 
	"PayeeName" TEXT NOT NULL, 
	"PayeeBankMfo" BIGINT NOT NULL, 
	"PayeeBankAccount" TEXT NOT NULL, 
	"PaymentNarrative" TEXT NOT NULL
);