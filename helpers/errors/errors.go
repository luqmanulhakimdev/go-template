package errors

import "errors"

type SystemError struct {
	Error   error
	Code    int
	Message string
}

var (
	DataNotFound = SystemError{
		Error:   errors.New("data_not_found"),
		Code:    1,
		Message: "Data tidak ditemukan",
	}
	InvalidParam = SystemError{
		Error:   errors.New("invalid_param"),
		Code:    1,
		Message: "Parameter tidak valid",
	}
	InvalidData = SystemError{
		Error:   errors.New("invalid_data"),
		Code:    0,
		Message: "Data tidak valid",
	}
	InvalidPaymentTaxBy = SystemError{
		Error:   errors.New("invalid_payment_tax_by"),
		Code:    0,
		Message: "Payment Tax By tidak valid",
	}
	InvalidDate = SystemError{
		Error:   errors.New("invalid_date"),
		Code:    0,
		Message: "Tanggal tidak valid",
	}
	InvalidType = SystemError{
		Error:   errors.New("invalid_type"),
		Code:    0,
		Message: "Type tidak valid",
	}
	InvalidTransactionType = SystemError{
		Error:   errors.New("invalid_transaction_type"),
		Code:    0,
		Message: "Transaction type tidak valid",
	}
	InvalidServiceName = SystemError{
		Error:   errors.New("invalid_service_name"),
		Code:    0,
		Message: "Service name tidak valid",
	}
	InvalidPaymentChannel = SystemError{
		Error:   errors.New("invalid_payment_channel"),
		Code:    0,
		Message: "Payment channel tidak valid",
	}
	InvalidStatus = SystemError{
		Error:   errors.New("invalid_status"),
		Code:    0,
		Message: "Status tidak valid",
	}
	InvalidName = SystemError{
		Error:   errors.New("invalid_name"),
		Code:    0,
		Message: "Name tidak valid",
	}
	InvalidAmount = SystemError{
		Error:   errors.New("invalid_amount"),
		Code:    0,
		Message: "Amount tidak valid",
	}
	ZeroBalance = SystemError{
		Error:   errors.New("zero_balance"),
		Code:    0,
		Message: "Balance kosong",
	}
	SettingRegistered = SystemError{
		Error:   errors.New("setting_registered"),
		Code:    0,
		Message: "Setting Name telah terdaftar",
	}
	AsanaEmailAssigneeRequired = SystemError{
		Error:   errors.New("asana_email_assignee_required"),
		Code:    0,
		Message: "Asana email assignee required",
	}
	AsanaCommentMentionIDRequired = SystemError{
		Error:   errors.New("asana_comment_mention_id_required"),
		Code:    0,
		Message: "Asana comment mention id required",
	}
	AsanaEmailFollowersRequired = SystemError{
		Error:   errors.New("asana_email_followers_required"),
		Code:    0,
		Message: "Asana email followers required",
	}
	OriginAccountNumberRequired = SystemError{
		Error:   errors.New("origin_account_number_required"),
		Code:    0,
		Message: "Origin account number required",
	}
	DestinationAccountNumberRequired = SystemError{
		Error:   errors.New("destination_account_number_required"),
		Code:    0,
		Message: "Destination account number required",
	}
	OriginDescriptionRequired = SystemError{
		Error:   errors.New("origin_description_required"),
		Code:    0,
		Message: "Origin description required",
	}
	DestinationDescriptionRequired = SystemError{
		Error:   errors.New("destination_description_required"),
		Code:    0,
		Message: "Destination description required",
	}
	DebitAccountNumberRequired = SystemError{
		Error:   errors.New("debit_account_number_required"),
		Code:    0,
		Message: "Debit account number required",
	}
	CreditAccountNumberRequired = SystemError{
		Error:   errors.New("credit_account_number_required"),
		Code:    0,
		Message: "Credit account number required",
	}
	InvalidDataSetting = SystemError{
		Error:   errors.New("invalid_data_setting"),
		Code:    0,
		Message: "Data setting tidak valid",
	}
	DataAlreadyExist = SystemError{
		Error:   errors.New("data_already_exist"),
		Code:    0,
		Message: "Data sudah ada",
	}
	LoanVirtualAccountRegistered = SystemError{
		Error:   errors.New("loan_virtual_account_registered"),
		Code:    0,
		Message: "Loan Virtual Account Type telah terdaftar",
	}
	AmountZero = SystemError{
		Error:   errors.New("amount_is_zero"),
		Code:    0,
		Message: "Total Amount 0",
	}
	ErrorValidateRequest = SystemError{
		Error:   errors.New("error_validate_request"),
		Code:    0,
		Message: "Error Validate Request",
	}
	CannotDecodeJson = SystemError{
		Error:   errors.New("cannot_decode_json"),
		Code:    0,
		Message: "Cannot Decode JSON",
	}
	ErrInternalServerError = SystemError{
		Error:   errors.New("err_internal_server_error"),
		Code:    0,
		Message: "ERR_INTERNAL_SERVER_ERROR",
	}
	ErrorCreateTx = SystemError{
		Error:   errors.New("error_create_tx"),
		Code:    0,
		Message: "Rrror Create TX",
	}
	ErrorFindOne = SystemError{
		Error:   errors.New("error_find_one"),
		Code:    0,
		Message: "Error Find One",
	}
	ErrorSelectAll = SystemError{
		Error:   errors.New("error_select_all"),
		Code:    0,
		Message: "Error Select All",
	}
	ErrorValidationValue = SystemError{
		Error:   errors.New("error_validation_value"),
		Code:    0,
		Message: "Error Validation Value",
	}
	ErrorOpenFile = SystemError{
		Error:   errors.New("error_open_file"),
		Code:    0,
		Message: "Error Open File",
	}
	ErrorWriteFileCSV = SystemError{
		Error:   errors.New("error_write_file_csv"),
		Code:    0,
		Message: "Error Write File CSV",
	}
	ErrorCheckDetails = SystemError{
		Error:   errors.New("error_check_details"),
		Code:    0,
		Message: "Error Check Details",
	}
	ErrorCheckData = SystemError{
		Error:   errors.New("error_check_data"),
		Code:    0,
		Message: "Error Check Data",
	}
	SavingAccountNotFound = SystemError{
		Error:   errors.New("saving_account_not_found"),
		Code:    1,
		Message: "Saving Account tidak ditemukan",
	}
	ErrorGetSavingAccount = SystemError{
		Error:   errors.New("error_get_saving_account"),
		Code:    1,
		Message: "Error Get Saing Account",
	}
	ErrorJournaling = SystemError{
		Error:   errors.New("error_journaling"),
		Code:    1,
		Message: "Error Journaling",
	}
	EmptyAccountNumber = SystemError{
		Error:   errors.New("empty_account_number"),
		Code:    1,
		Message: "Empty Account Number in Setting",
	}
	InsufficientCashbackBudget = SystemError{
		Error:   errors.New("insufficient_cashback_budget"),
		Code:    1,
		Message: "Insufficient Cashback Budget",
	}
)
