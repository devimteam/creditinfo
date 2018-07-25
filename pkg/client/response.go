package client

import "time"

type PolicyRule struct {
	RuleId      string
	Result      string
	Description string
}

type ScoringAnalysis struct {
	CIPScore             int32
	CIPRiskGrade         string
	MobileScore          int32
	MobileScoreRiskGrade string
	Conclusion           string
	PolicyRules          []PolicyRule
}

type PersonalInformation struct {
	FullName         string
	DateOfBirth      time.Time
	Age              uint32
	Gender           string
	MaritalStatus    string
	EmploymentStatus string
}

type GeneralInformation struct {
	SubjectIDNumber     string
	RequestDate         time.Time
	ReferenceNumber     string
	RecommendedDecision string
	BrokenRules         int32
	CreditLimit         int32
}

type InquiriesAnalysis struct {
	TotalLast7Days       int32
	TotalLast1Month      int32
	NonBankingLast1Month int32
	PolicyRules          []PolicyRule
	Conclusion           *string
}

type CurrentBanking struct {
	Positive      int32
	Negative      int32
	Balance       int32
	BalanceAtRisk int32
}

type CurrentContracts struct {
	CurrentBanking    CurrentBanking
	CurrentNonBanking CurrentBanking
	Total             CurrentBanking
}

type PastDueInformation struct {
	TotalCurrentPastDue                int32
	TotalCurrentDaysPastDue            int32
	MonthsWithoutArrearsLast12Months   int32
	TotalMonthsWithHistoryLast12Months int32
}

type RepaymentInformation struct {
	TotalMonthlyPayment int32
	ClosedContracts     int32
}

type CIP struct {
	RecordList []Record
}

type CIQDetail struct {
	LostStolenRecordsFound           int32
	NumberOfCancelledClosedContracts int32
}

type CIQSummary struct {
	NumberOfFraudAlertsPrimarySubject int32
	NumberOfFraudAlertsThirdParty     int32
}

type CIQ struct {
	Detail  CIQDetail
	Summary CIQSummary
}

type KenCb5Data struct {
	//<BouncedCheques>
	//<ChequeList />
	//</BouncedCheques>
	CIP                 CIP
	CIQ                 CIQ
	ContractSummary     ContractSummary
	Contracts           Contracts
	CurrentRelations    CurrentRelations
	Dashboard           Dashboard
	Disputes            DisputeList
	Individual          Individual
	Inquiries           Inquiries
	Parameters          Parameters
	PaymentIncidentList PaymentIncidentList
}

type Reason struct {
	Code        string
	Description string
}

type Record struct {
	Date                 time.Time
	Grade                string
	ProbabilityOfDefault int32
	ReasonsList          []Reason
	Score                int32
	Trend                string
}

type CurrentBalance struct {
	Currency   string
	LocalValue float64
	Value      float64
}

type Disputes struct {
	ClosedDisputes bool
	//<DisputeList />
	FalseDisputes bool
}
type Contract struct {
	AccountProductType  string
	AccountStatus       string
	CurrentBalance      CurrentBalance
	DateAccountOpened   time.Time
	DaysInArrears       int32
	OriginalAmount      CurrentBalance
	OverdueBalance      CurrentBalance
	PerformingIndicator string
	PhaseOfContract     string
	RoleOfClient        string
	Sector              string

	Branch *string
	// CollateralsList

	//<CollateralsList />
	ContractCode             *string
	CurrencyOfFacility       *string
	DateOfLastPayment        *time.Time
	Disputes                 *Disputes
	InstallmentAmount        *CurrentBalance
	InstallmentsInArrears    *int32
	NegativeStatusOfContract *string
	OutstandingAmount        *CurrentBalance
	PaymentCalendarList      *[]PaymentCalendar
	PaymentFrequency         *string
	//<RelatedSubjectsList />
	Subscriber     *string
	SubscriberType *string
}
type ContractOverview struct {
	ContractList []Contract
}

type AffordabilityHistory struct {
	Month                int32
	MonthlyAffordability CurrentBalance
	Year                 int32
}
type AffordabilitySummary struct {
	MonthlyAffordability CurrentBalance
}

type Debtor struct {
	ClosedContracts   int32
	CurrentBalanceSum CurrentBalance
	OpenContracts     int32
	OriginalAmountSum CurrentBalance
	OverdueBalanceSum CurrentBalance
}

type Guarantor struct {
	ClosedContracts bool
	OpenContracts   bool
}

type Overall struct {
	LastDelinquency90PlusDays         time.Time
	MaxDueInstallmentsClosedContracts int32
	MaxDueInstallmentsOpenContracts   int32
	WorstDaysInArrears                int32
	WorstOverdueBalance               CurrentBalance
}

type PaymentCalendar struct {
	ContractsSubmitted    int32
	Date                  time.Time
	CurrentBalance        CurrentBalance
	DaysInArrears         *int32
	InstallmentsInArrears *int32
	OverdueBalance        *CurrentBalance
	DelinquencyStatus     *string
	OutstandingAmount     *CurrentBalance
	PastDueAmount         *CurrentBalance
	PastDueDays           *int32
	PastDueInstallments   *int32
}

type SectorInfo struct {
	DebtorClosedContracts    int32
	DebtorCurrentBalanceSum  CurrentBalance
	DebtorOpenContracts      int32
	DebtorOriginalAmountSum  CurrentBalance
	DebtorOverdueBalanceSum  CurrentBalance
	GuarantorClosedContracts bool
	GuarantorOpenContracts   bool
	Sector                   string
}
type ContractSummary struct {
	AffordabilityHistoryList []AffordabilityHistory
	AffordabilitySummary     AffordabilitySummary
	Debtor                   Debtor
	Guarantor                Guarantor
	Overall                  Overall
	PaymentCalendarList      []PaymentCalendar
	SectorInfoList           []SectorInfo
}

type Contracts struct {
	ContractList []Contract
}

type MainAddress struct {
	AddressLine string
	Country     string
	District    string
	Region      string
	PlotNumber  string
	PostalCode  int32
	Street      string
	Town        string
}

type Contact struct {
	MobileTelephoneNumber string
	Email                 *string
	HomeTelephoneNumber   *string
}

type ContractRelation struct {
	Contact        Contact
	FullName       string
	IDNumber       string
	IDNumberType   string
	MainAddress    MainAddress
	RoleOfCustomer string
	SubjectType    string
	ValidFrom      time.Time
}

type CurrentRelations struct {
	ContractRelationList []ContractRelation
	InvolvementList      []Involvement
	//<RelatedPartyList />
}

type Involvement struct {
	Contact        Contact
	MainAddress    MainAddress
	FullName       string
	IDNumber       string
	IDNumberType   string
	SubjectType    string
	TypeOfRelation string
	ValidFrom      time.Time
}

type BouncedCheques struct {
	NumberOfCheques     int32
	WeekSinceLastCheque int32
}

type DashboardCIP struct {
	Grade string
	Score int32
	Trend string
}

type DashboardCIQ struct {
	FraudAlerts           int32
	FraudAlertsThirdParty int32
}

type Collaterals struct {
	HighestCollateralValue     CurrentBalance
	HighestCollateralValueType string
	NumberOfCollaterals        int32
	TotalCollateralValue       CurrentBalance
}

type DashboardDisputes struct {
	ActiveContractDisputes          bool
	ActiveSubjectDisputes           bool
	ClosedContractDisputes          bool
	ClosedSubjectDisputes           bool
	FalseDisputes                   bool
	NumberOfCourtRegisteredDisputes bool
}

type DashboardInquiries struct {
	InquiriesForLast12Months  int32
	SubscribersInLast12Months int32
}

type PaymentsProfile struct {
	InstallmentAmountSum            CurrentBalance
	NumberOfDifferentSubscribers    int32
	PastDueAmountSum                CurrentBalance
	WorstPastDueDaysCurrent         int32
	WorstPastDueDaysForLast12Months int32
}

type Relations struct {
	NumberOfInvolvements int32
	NumberOfRelations    int32
}

type Dashboard struct {
	BouncedCheques  BouncedCheques
	CIP             DashboardCIP
	CIQ             DashboardCIQ
	Collaterals     Collaterals
	Disputes        DashboardDisputes
	Inquiries       DashboardInquiries
	PaymentsProfile PaymentsProfile
	Relations       Relations
}

type DisputeItem struct {
	Comment    string
	Created    time.Time
	InCourt    bool
	Resolution string
	Status     string
	Type       string
}

type DisputeListSummary struct {
	NumberOfActiveDisputesContracts   int32
	NumberOfActiveDisputesInCourt     int32
	NumberOfActiveDisputesSubjectData int32
	NumberOfClosedDisputesContracts   int32
	NumberOfClosedDisputesSubjectData int32
	NumberOfFalseDisputes             int32
}

type DisputeList struct {
	DisputeList []DisputeItem
	Summary     DisputeListSummary
}

type IndividualGeneral struct {
	ClassificationOfIndividual string
	DateOfBirth                time.Time
	EmployerName               string
	Employment                 string
	Forename1                  string
	Forename2                  string
	FullName                   string
	Gender                     string
	MaritalStatus              string
	Surname                    string
}

type Individual struct {
	Contact          Contact
	General          IndividualGeneral
	Identifications  Identifications
	MainAddress      MainAddress
	SecondaryAddress MainAddress
}

type Identifications struct {
	AlienRegistration string
	NationalID        string
	PassportNumber    string
	ServiceId         string
}

type Inquiry struct {
	DateOfInquiry time.Time
	Product       string
	Reason        string
	Sector        string
	Subscriber    string
}

type InquiriesSummary struct {
	NumberOfInquiriesLast12Months int32
	NumberOfInquiriesLast1Month   int32
	NumberOfInquiriesLast24Months int32
	NumberOfInquiriesLast3Months  int32
	NumberOfInquiriesLast6Months  int32
}

type Inquiries struct {
	InquiryList []Inquiry
	Summary     InquiriesSummary
}

type Parameters struct {
	Consent       bool
	IDNumber      string
	IDNumberType  string
	InquiryReason string
	ReportDate    time.Time
	Sections      string
	SubjectType   string
}

type PaymentIncidentSummary struct {
	DueAmountSummary CurrentBalance
}

type PaymentIncidentItem struct {
	Summary PaymentIncidentSummary
}

type PaymentIncidentList struct {
	PaymentIncidentList []PaymentIncidentItem
	ReportInfo          ReportInfo
	SubjectInfoHistory  SubjectInfoHistory
}

type ReportInfo struct {
	Created         time.Time
	ReferenceNumber string
	ReportStatus    string
	Version         string
}

type General struct {
	Item       string
	Subscriber string
	ValidFrom  string
	ValidTo    string
	Value      string
}

type SubjectInfoHistory struct {
	AddressList *[]MainAddress
	ContactList *[]Contact
	GeneralList []General
}

type Strategy struct {
	Id           string
	Name         string
	BeeStrategy  string
	TemplateName string
}

type QueryResponse struct {
	Status               string
	Infomsg              string
	Currency             string
	GeneralInformation   GeneralInformation
	PersonalInformation  PersonalInformation
	ScoringAnalysis      ScoringAnalysis
	InquiriesAnalysis    InquiriesAnalysis
	CurrentContracts     CurrentContracts
	PastDueInformation   PastDueInformation
	RepaymentInformation RepaymentInformation
	PolicyRules          []PolicyRule
	KenCb5Data           KenCb5Data
}
