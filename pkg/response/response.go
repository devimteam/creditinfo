package response

import "time"

type MultiConnectorResponse struct {
	MessageId     *string      `xml:"MessageId,omitempty" json:"MessageId,omitempty" yaml:"MessageId,omitempty"`
	OperationCode *string      `xml:"OperationCode,omitempty" json:"OperationCode,omitempty" yaml:"OperationCode,omitempty"`
	ResponseXml   *ResponseXml `xml:"ResponseXml,omitempty" json:"ResponseXml,omitempty" yaml:"ResponseXml,omitempty"`
	Timestamp     *string      `xml:"Timestamp,omitempty" json:"Timestamp,omitempty" yaml:"Timestamp,omitempty"`
}

type ResponseXml struct {
	Response Response `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Messages/Response response,omitempty" json:"response,omitempty" yaml:"response,omitempty"`
}

type Response struct {
	Connector *ConnectorResponse `xml:"connector,omitempty"`
}

type ConnectorDataResponse struct {
	// Unique ID of request to particular connector.
	Id       *string         `xml:"id,attr,omitempty"`
	Response *ResultResponse `xml:"http://creditinfo.com/schemas/2012/09/MultiConnector/Connectors/INT/IdmStrategy/Response response,omitempty"`
}

type ConnectorResponse struct {
	Data  *ConnectorDataResponse `xml:"data,omitempty"`
	Error *ErrorResponse         `xml:"error,omitempty"`
}

type ResultResponse struct {
	Status               string                `xml:"status,omitempty"`
	Infomsg              string                `xml:"infomsg,omitempty"`
	Currency             string                `xml:"Currency,omitempty"`
	GeneralInformation   *GeneralInformation   `xml:"GeneralInformation,omitempty"`
	PersonalInformation  *PersonalInformation  `xml:"PersonalInformation,omitempty"`
	ScoringAnalysis      *ScoringAnalysis      `xml:"ScoringAnalysis,omitempty"`
	InquiriesAnalysis    *InquiriesAnalysis    `xml:"InquiriesAnalysis,omitempty"`
	CurrentContracts     *CurrentContracts     `xml:"CurrentContracts,omitempty"`
	PastDueInformation   *PastDueInformation   `xml:"PastDueInformation,omitempty"`
	RepaymentInformation *RepaymentInformation `xml:"RepaymentInformation,omitempty"`
	PolicyRules          *PolicyRule           `xml:"PolicyRules,omitempty"`
	KenCb5Data           *KenCb5Data           `xml:"KenCb5_data,omitempty"`
	Strategy             *StrategyResponse     `xml:"Strategy,omitempty"`
}
type GeneralInformation struct {
	SubjectIDNumber     string    `xml:"SubjectIDNumber,omitempty"`
	RequestDate         time.Time `xml:"RequestDate,omitempty"`
	ReferenceNumber     string    `xml:"ReferenceNumber,omitempty"`
	RecommendedDecision string    `xml:"RecommendedDecision,omitempty"`
	BrokenRules         int32     `xml:"BrokenRules,omitempty"`
	CreditLimit         int32     `xml:"CreditLimit,omitempty"`
}

type PersonalInformation struct {
	FullName         string    `xml:"FullName,omitempty"`
	DateOfBirth      time.Time `xml:"DateOfBirth,omitempty"`
	Age              uint32    `xml:"Age,omitempty"`
	Gender           string    `xml:"Gender,omitempty"`
	MaritalStatus    string    `xml:"MaritalStatus,omitempty"`
	EmploymentStatus string    `xml:"EmploymentStatus,omitempty"`
}

type ScoringAnalysis struct {
	CIPScore             int32       `xml:"CIPScore,omitempty"`
	CIPRiskGrade         string      `xml:"CIPRiskGrade,omitempty"`
	MobileScore          int32       `xml:"MobileScore,omitempty"`
	MobileScoreRiskGrade string      `xml:"MobileScoreRiskGrade,omitempty"`
	Conclusion           string      `xml:"Conclusion,omitempty"`
	PolicyRules          *PolicyRule `xml:"PolicyRules,omitempty"`
}

type RuleItem struct {
	Id          string `xml:"id,attr,omitempty"`
	Result      string `xml:"Result,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type PolicyRule struct {
	Rule []RuleItem `xml:"Rule,omitempty"`
}

type InquiriesAnalysis struct {
	TotalLast7Days       int32       `xml:"TotalLast7Days,omitempty"`
	TotalLast1Month      int32       `xml:"TotalLast1Month,omitempty"`
	NonBankingLast1Month int32       `xml:"NonBankingLast1Month,omitempty"`
	PolicyRules          *PolicyRule `xml:"PolicyRules,omitempty"`
	Conclusion           string      `xml:"Conclusion,omitempty"`
}

type BankingData struct {
	Positive      int32   `xml:"Positive,omitempty"`
	Negative      int32   `xml:"Negative,omitempty"`
	Balance       float64 `xml:"Balance,omitempty"`
	BalanceAtRisk int32   `xml:"BalanceAtRisk,omitempty"`
}

type CurrentContracts struct {
	CurrentBanking    *BankingData `xml:"CurrentBanking,omitempty"`
	CurrentNonBanking *BankingData `xml:"CurrentNonBanking,omitempty"`
	Total             *BankingData `xml:"Total,omitempty"`
}

type PastDueInformation struct {
	TotalCurrentPastDue                  int32 `xml:"TotalCurrentPastDue,omitempty"`
	TotalCurrentDaysPastDue              int32 `xml:"TotalCurrentDaysPastDue,omitempty"`
	MonthsWithoutArrearsLast12Months     int32 `xml:"MonthsWithoutArrearsLast12Months,omitempty"`
	TotalMonthsWithHistoryLast12Months   int32 `xml:"TotalMonthsWithHistoryLast12Months,omitempty"`
	PercMonthsWithoutArrearsLast12Months int32 `xml:"PercMonthsWithoutArrearsLast12Months,omitempty"`
}

type RepaymentInformation struct {
	TotalMonthlyPayment int32 `xml:"TotalMonthlyPayment,omitempty"`
	ClosedContracts     int32 `xml:"ClosedContracts,omitempty"`
}

type KenCb5Data struct {
	CIP                 *KenCb5DataCIP       `xml:"CIP,omitempty"`
	CIQ                 *KenCb5DataCIQ       `xml:"CIQ,omitempty"`
	ContractOverview    *ContractOverview    `xml:"ContractOverview,omitempty"`
	ContractSummary     *ContractSummary     `xml:"ContractSummary,omitempty"`
	Contracts           *Contracts           `xml:"Contracts,omitempty"`
	CurrentRelations    *CurrentRelations    `xml:"CurrentRelations,omitempty"`
	Dashboard           *Dashboard           `xml:"Dashboard,omitempty"`
	Disputes            *Disputes            `xml:"Disputes,omitempty"`
	Individual          *Individual          `xml:"Individual,omitempty"`
	Inquiries           *Inquiries           `xml:"Inquiries,omitempty"`
	Parameters          *Parameters          `xml:"Parameters,omitempty"`
	PaymentIncidentList *PaymentIncidentList `xml:"PaymentIncidentList,omitempty"`
	ReportInfo          *ReportInfo          `xml:"ReportInfo,omitempty"`
	SubjectInfoHistory  *SubjectInfoHistory  `xml:"SubjectInfoHistory,omitempty"`
}

type Record struct {
	Date                 time.Time    `xml:"Date,omitempty"`
	Grade                string       `xml:"Grade,omitempty"`
	ProbabilityOfDefault float32      `xml:"ProbabilityOfDefault,omitempty"`
	ReasonsList          *ReasonsList `xml:"ReasonsList,omitempty"`
	Score                int32        `xml:"Score,omitempty"`
	Trend                string       `xml:"Trend,omitempty"`
}

type RecordList struct {
	Records []Record `xml:"Record,omitempty"`
}

type KenCb5DataCIP struct {
	RecordList *RecordList `xml:"RecordList,omitempty"`
}

type Reason struct {
	Code        string `xml:"Code,omitempty"`
	Description string `xml:"Description,omitempty"`
}

type ReasonsList struct {
	Reasons []Reason `xml:"Reason,omitempty"`
}

type KenCb5DataCIQDetail struct {
	LostStolenRecordsFound           int32 `xml:"LostStolenRecordsFound,omitempty"`
	NumberOfCancelledClosedContracts int32 `xml:"NumberOfCancelledClosedContracts,omitempty"`
}

type KenCb5DataCIQSummary struct {
	NumberOfFraudAlertsPrimarySubject int32 `xml:"NumberOfFraudAlertsPrimarySubject,omitempty"`
	NumberOfFraudAlertsThirdParty     int32 `xml:"NumberOfFraudAlertsThirdParty,omitempty"`
}

type KenCb5DataCIQ struct {
	Detail  *KenCb5DataCIQDetail  `xml:"Detail,omitempty"`
	Summary *KenCb5DataCIQSummary `xml:"Summary,omitempty"`
}

type Balance struct {
	Currency   string  `xml:"Currency,omitempty"`
	LocalValue float64 `xml:"LocalValue,omitempty"`
	Value      float64 `xml:"Value,omitempty"`
}

type Contract struct {
	AccountProductType       string                       `xml:"AccountProductType,omitempty"`
	AccountStatus            string                       `xml:"AccountStatus,omitempty"`
	CurrentBalance           *Balance                     `xml:"CurrentBalance,omitempty"`
	DateAccountOpened        time.Time                    `xml:"DateAccountOpened,omitempty"`
	DaysInArrears            int32                        `xml:"DaysInArrears,omitempty"`
	OriginalAmount           Balance                      `xml:"OriginalAmount,omitempty"`
	OverdueBalance           Balance                      `xml:"OverdueBalance,omitempty"`
	PerformingIndicator      string                       `xml:"PerformingIndicator,omitempty"`
	PhaseOfContract          string                       `xml:"PhaseOfContract,omitempty"`
	RoleOfClient             string                       `xml:"RoleOfClient,omitempty"`
	Sector                   string                       `xml:"Sector,omitempty"`
	Branch                   *string                      `xml:"Branch,omitempty"`
	ContractCode             *string                      `xml:"ContractCode,omitempty"`
	CurrencyOfFacility       *string                      `xml:"CurrencyOfFacility,omitempty"`
	DateOfLastPayment        *time.Time                   `xml:"DateOfLastPayment,omitempty"`
	Disputes                 *ContractDisputes            `xml:"Disputes,omitempty"`
	InstallmentAmount        *Balance                     `xml:"InstallmentAmount,omitempty"`
	InstallmentsInArrears    int32                        `xml:"InstallmentsInArrears,omitempty"`
	NegativeStatusOfContract string                       `xml:"NegativeStatusOfContract,omitempty"`
	OutstandingAmount        *Balance                     `xml:"OutstandingAmount,omitempty"`
	PaymentCalendarList      *ContractPaymentCalendarList `xml:"PaymentCalendarList,omitempty"`
	PaymentFrequency         string                       `xml:"PaymentFrequency,omitempty"`
	Subscriber               string                       `xml:"Subscriber,omitempty"`
	SubscriberType           string                       `xml:"SubscriberType,omitempty"`
}

type ContractList struct {
	Contracts []Contract `xml:"Contract,omitempty"`
}

type ContractOverview struct {
	ContractList ContractList `xml:"ContractList,omitempty"`
}

type ContractDisputes struct {
	ClosedDisputes int32
	FalseDisputes  int32
}

type PaymentCalendarList struct {
	Calendars []PaymentCalendar `xml:"PaymentCalendar,omitempty"`
}

type ContractPaymentCalendarList struct {
	Calendars []PaymentCalendar `xml:"CalendarItem,omitempty"`
}

type PaymentCalendar struct {
	ContractsSubmitted    int32     `xml:"ContractsSubmitted,omitempty"`
	Date                  time.Time `xml:"Date,omitempty"`
	CurrentBalance        *Balance  `xml:"CurrentBalance,omitempty"`
	DaysInArrears         int32     `xml:"DaysInArrears,omitempty"`
	InstallmentsInArrears int32     `xml:"InstallmentsInArrears,omitempty"`
	OverdueBalance        *Balance  `xml:"OverdueBalance,omitempty"`
	DelinquencyStatus     string    `xml:"DelinquencyStatus,omitempty"`
	OutstandingAmount     *Balance  `xml:"OutstandingAmount,omitempty"`
	PastDueAmount         *Balance  `xml:"PastDueAmount,omitempty"`
	PastDueDays           int32     `xml:"PastDueDays,omitempty"`
	PastDueInstallments   int32     `xml:"PastDueInstallments,omitempty"`
}

type ContractSummary struct {
	AffordabilityHistoryList *AffordabilityHistoryList `xml:"AffordabilityHistoryList,omitempty"`
	AffordabilitySummary     *AffordabilitySummary     `xml:"AffordabilitySummary,omitempty"`
	Debtor                   *Debtor                   `xml:"Debtor,omitempty"`
	Guarantor                *Guarantor                `xml:"Guarantor,omitempty"`
	Overall                  *Overall                  `xml:"Overall,omitempty"`
	PaymentCalendarList      *PaymentCalendarList      `xml:"PaymentCalendarList,omitempty"`
	SectorInfoList           *SectorInfoList           `xml:"SectorInfoList,omitempty"`
}

type SectorInfoList struct {
	Sectors []SectorInfo `xml:"SectorInfo,omitempty"`
}

type AffordabilityHistoryList struct {
	AffordabilityHistory []AffordabilityHistory `xml:"AffordabilityHistory,omitempty"`
}

type Contracts struct {
	ContractList ContractList `xml:"ContractList,omitempty"`
}

type AffordabilityHistory struct {
	Month                int32    `xml:"Month,omitempty"`
	MonthlyAffordability *Balance `xml:"MonthlyAffordability,omitempty"`
	Year                 int32    `xml:"Year,omitempty"`
}

type AffordabilitySummary struct {
	MonthlyAffordability *Balance `xml:"MonthlyAffordability,omitempty"`
}

type Debtor struct {
	ClosedContracts   int32    `xml:"ClosedContracts,omitempty"`
	CurrentBalanceSum *Balance `xml:"CurrentBalanceSum,omitempty"`
	OpenContracts     int32    `xml:"OpenContracts,omitempty"`
	OriginalAmountSum *Balance `xml:"OriginalAmountSum,omitempty"`
	OverdueBalanceSum *Balance `xml:"OverdueBalanceSum,omitempty"`
}

type Guarantor struct {
	ClosedContracts int32 `xml:"ClosedContracts,omitempty"`
	OpenContracts   int32 `xml:"OpenContracts,omitempty"`
}

type Overall struct {
	LastDelinquency90PlusDays         time.Time `xml:"LastDelinquency90PlusDays,omitempty"`
	MaxDueInstallmentsClosedContracts int32     `xml:"MaxDueInstallmentsClosedContracts,omitempty"`
	MaxDueInstallmentsOpenContracts   int32     `xml:"MaxDueInstallmentsOpenContracts,omitempty"`
	WorstDaysInArrears                int32     `xml:"WorstDaysInArrears,omitempty"`
	WorstOverdueBalance               *Balance  `xml:"WorstOverdueBalance,omitempty"`
}

type SectorInfo struct {
	DebtorClosedContracts    int32    `xml:"DebtorClosedContracts,omitempty"`
	DebtorCurrentBalanceSum  *Balance `xml:"DebtorCurrentBalanceSum,omitempty"`
	DebtorOpenContracts      int32    `xml:"DebtorOpenContracts,omitempty"`
	DebtorOriginalAmountSum  *Balance `xml:"DebtorOriginalAmountSum,omitempty"`
	DebtorOverdueBalanceSum  *Balance `xml:"DebtorOverdueBalanceSum,omitempty"`
	GuarantorClosedContracts int32    `xml:"GuarantorClosedContracts,omitempty"`
	GuarantorOpenContracts   int32    `xml:"GuarantorOpenContracts,omitempty"`
	Sector                   string   `xml:"Sector,omitempty"`
}

type ContractRelationList struct {
	ContractRelations []ContractRelation `xml:"ContractRelation,omitempty"`
}

type InvolvementList struct {
	Involvements []Involvement `xml:"Involvement,omitempty"`
}

type CurrentRelations struct {
	ContractRelationList *ContractRelationList `xml:"ContractRelationList,omitempty"`
	InvolvementList      *InvolvementList      `xml:"InvolvementList,omitempty"`
}

type Involvement struct {
	Contact        *Contact     `xml:"Contact,omitempty"`
	MainAddress    *MainAddress `xml:"MainAddress,omitempty"`
	FullName       string       `xml:"FullName,omitempty"`
	IDNumber       string       `xml:"IDNumber,omitempty"`
	IDNumberType   string       `xml:"IDNumberType,omitempty"`
	SubjectType    string       `xml:"SubjectType,omitempty"`
	TypeOfRelation string       `xml:"TypeOfRelation,omitempty"`
	ValidFrom      time.Time    `xml:"ValidFrom,omitempty"`
}

type Contact struct {
	MobileTelephoneNumber string  `xml:"MobileTelephoneNumber,omitempty"`
	Email                 *string `xml:"Email,omitempty"`
	HomeTelephoneNumber   *string `xml:"HomeTelephoneNumber,omitempty"`
}

type MainAddress struct {
	AddressLine string `xml:"AddressLine,omitempty"`
	Country     string `xml:"Country,omitempty"`
	District    string `xml:"District,omitempty"`
	Region      string `xml:"Region,omitempty"`
	PlotNumber  string `xml:"PlotNumber,omitempty"`
	PostalCode  string `xml:"PostalCode,omitempty"`
	Street      string `xml:"Street,omitempty"`
	Town        string `xml:"Town,omitempty"`
}

type ContractRelation struct {
	Contact        *Contact     `xml:"Contact,omitempty"`
	FullName       string       `xml:"FullName,omitempty"`
	IDNumber       string       `xml:"IDNumber,omitempty"`
	IDNumberType   string       `xml:"IDNumberType,omitempty"`
	MainAddress    *MainAddress `xml:"MainAddress,omitempty"`
	RoleOfCustomer string       `xml:"RoleOfCustomer,omitempty"`
	SubjectType    string       `xml:"SubjectType,omitempty"`
	ValidFrom      time.Time    `xml:"ValidFrom,omitempty"`
}

type DashboardCIP struct {
	Grade string `xml:"Grade,omitempty"`
	Score int32  `xml:"Score,omitempty"`
	Trend string `xml:"Trend,omitempty"`
}

type DashboardCIQ struct {
	FraudAlerts           int32 `xml:"FraudAlerts,omitempty"`
	FraudAlertsThirdParty int32 `xml:"FraudAlertsThirdParty,omitempty"`
}

type Collaterals struct {
	HighestCollateralValue     *Balance `xml:"HighestCollateralValue,omitempty"`
	HighestCollateralValueType string   `xml:"HighestCollateralValueType,omitempty"`
	NumberOfCollaterals        int32    `xml:"NumberOfCollaterals,omitempty"`
	TotalCollateralValue       *Balance `xml:"TotalCollateralValue,omitempty"`
}

type DashboardDisputes struct {
	ActiveContractDisputes          int32 `xml:"ActiveContractDisputes,omitempty"`
	ActiveSubjectDisputes           int32 `xml:"ActiveSubjectDisputes,omitempty"`
	ClosedContractDisputes          int32 `xml:"ClosedContractDisputes,omitempty"`
	ClosedSubjectDisputes           int32 `xml:"ClosedSubjectDisputes,omitempty"`
	FalseDisputes                   int32 `xml:"FalseDisputes,omitempty"`
	NumberOfCourtRegisteredDisputes int32 `xml:"NumberOfCourtRegisteredDisputes,omitempty"`
}

type DashboardInquiries struct {
	InquiriesForLast12Months  int32 `xml:"InquiriesForLast12Months,omitempty"`
	SubscribersInLast12Months int32 `xml:"SubscribersInLast12Months,omitempty"`
}

type Dashboard struct {
	BouncedCheques  *BouncedCheques     `xml:"BouncedCheques,omitempty"`
	CIP             *DashboardCIP       `xml:"CIP,omitempty"`
	CIQ             *DashboardCIQ       `xml:"CIQ,omitempty"`
	Collaterals     *Collaterals        `xml:"Collaterals,omitempty"`
	Disputes        *DashboardDisputes  `xml:"Disputes,omitempty"`
	Inquiries       *DashboardInquiries `xml:"Inquiries,omitempty"`
	PaymentsProfile *PaymentsProfile    `xml:"PaymentsProfile,omitempty"`
	Relations       *Relations          `xml:"Relations,omitempty"`
}

type Disputes struct {
	DisputeList *DisputeList        `xml:"DisputeList,omitempty"`
	Summary     *DisputeListSummary `xml:"Summary,omitempty"`
}

type PaymentsProfile struct {
	InstallmentAmountSum            *Balance `xml:"InstallmentAmountSum,omitempty"`
	NumberOfDifferentSubscribers    int32    `xml:"NumberOfDifferentSubscribers,omitempty"`
	PastDueAmountSum                *Balance `xml:"PastDueAmountSum,omitempty"`
	WorstPastDueDaysCurrent         int32    `xml:"WorstPastDueDaysCurrent,omitempty"`
	WorstPastDueDaysForLast12Months int32    `xml:"WorstPastDueDaysForLast12Months,omitempty"`
}

type Relations struct {
	NumberOfInvolvements int32 `xml:"NumberOfInvolvements,omitempty"`
	NumberOfRelations    int32 `xml:"NumberOfRelations,omitempty"`
}

type BouncedCheques struct {
	NumberOfCheques     int32 `xml:"NumberOfCheques,omitempty"`
	WeekSinceLastCheque int32 `xml:"WeekSinceLastCheque,omitempty"`
}

type DisputeItem struct {
	Comment    string    `xml:"Comment,omitempty"`
	Created    time.Time `xml:"Created,omitempty"`
	InCourt    bool      `xml:"InCourt,omitempty"`
	Resolution string    `xml:"Resolution,omitempty"`
	Status     string    `xml:"Status,omitempty"`
	Type       string    `xml:"Type,omitempty"`
}

type DisputeListSummary struct {
	NumberOfActiveDisputesContracts   int32 `xml:"NumberOfActiveDisputesContracts,omitempty"`
	NumberOfActiveDisputesInCourt     int32 `xml:"NumberOfActiveDisputesInCourt,omitempty"`
	NumberOfActiveDisputesSubjectData int32 `xml:"NumberOfActiveDisputesSubjectData,omitempty"`
	NumberOfClosedDisputesContracts   int32 `xml:"NumberOfClosedDisputesContracts,omitempty"`
	NumberOfClosedDisputesSubjectData int32 `xml:"NumberOfClosedDisputesSubjectData,omitempty"`
	NumberOfFalseDisputes             int32 `xml:"NumberOfFalseDisputes,omitempty"`
}

type DisputeList struct {
	Disputes []DisputeItem `xml:"Dispute,omitempty"`
}

type IndividualGeneral struct {
	ClassificationOfIndividual string    `xml:"ClassificationOfIndividual,omitempty"`
	DateOfBirth                time.Time `xml:"DateOfBirth,omitempty"`
	EmployerName               string    `xml:"EmployerName,omitempty"`
	Employment                 string    `xml:"Employment,omitempty"`
	Forename1                  string    `xml:"Forename1,omitempty"`
	Forename2                  string    `xml:"Forename2,omitempty"`
	FullName                   string    `xml:"FullName,omitempty"`
	Gender                     string    `xml:"Gender,omitempty"`
	MaritalStatus              string    `xml:"MaritalStatus,omitempty"`
	Surname                    string    `xml:"Surname,omitempty"`
}

type Individual struct {
	Contact          *Contact           `xml:"Contact,omitempty"`
	General          *IndividualGeneral `xml:"General,omitempty"`
	Identifications  *Identifications   `xml:"Identifications,omitempty"`
	MainAddress      *MainAddress       `xml:"MainAddress,omitempty"`
	SecondaryAddress *MainAddress       `xml:"SecondaryAddress,omitempty"`
}

type Identifications struct {
	AlienRegistration string `xml:"AlienRegistration,omitempty"`
	NationalID        string `xml:"NationalID,omitempty"`
	PassportNumber    string `xml:"PassportNumber,omitempty"`
	ServiceId         string `xml:"ServiceId,omitempty"`
}

type Inquiry struct {
	DateOfInquiry time.Time `xml:"DateOfInquiry,omitempty"`
	Product       string    `xml:"Product,omitempty"`
	Reason        string    `xml:"Reason,omitempty"`
	Sector        string    `xml:"Sector,omitempty"`
	Subscriber    string    `xml:"Subscriber,omitempty"`
}

type InquiriesSummary struct {
	NumberOfInquiriesLast12Months int32 `xml:"NumberOfInquiriesLast12Months,omitempty"`
	NumberOfInquiriesLast1Month   int32 `xml:"NumberOfInquiriesLast1Month,omitempty"`
	NumberOfInquiriesLast24Months int32 `xml:"NumberOfInquiriesLast24Months,omitempty"`
	NumberOfInquiriesLast3Months  int32 `xml:"NumberOfInquiriesLast3Months,omitempty"`
	NumberOfInquiriesLast6Months  int32 `xml:"NumberOfInquiriesLast6Months,omitempty"`
}

type InquiryList struct {
	Inquirys []Inquiry `xml:"Inquiry,omitempty"`
}

type Inquiries struct {
	InquiryList *InquiryList      `xml:"InquiryList,omitempty"`
	Summary     *InquiriesSummary `xml:"Summary,omitempty"`
}

type Parameters struct {
	Consent       bool      `xml:"Consent,omitempty"`
	IDNumber      string    `xml:"IDNumber,omitempty"`
	IDNumberType  string    `xml:"IDNumberType,omitempty"`
	InquiryReason string    `xml:"InquiryReason,omitempty"`
	ReportDate    time.Time `xml:"ReportDate,omitempty"`
	Sections      string    `xml:"Sections,omitempty"`
	SubjectType   string    `xml:"SubjectType,omitempty"`
}

type PaymentIncidentSummary struct {
	DueAmountSummary *Balance `xml:"DueAmountSummary,omitempty"`
}

type PaymentIncidentList struct {
	Summary PaymentIncidentSummary `xml:"Summary,omitempty"`
}

type GeneralList struct {
	Generals []General `xml:"General,omitempty"`
}

type SubjectInfoHistory struct {
	GeneralList GeneralList `xml:"GeneralList,omitempty"`
}

type General struct {
	Item       string `xml:"Item,omitempty"`
	Subscriber string `xml:"Subscriber,omitempty"`
	ValidFrom  string `xml:"ValidFrom,omitempty"`
	ValidTo    string `xml:"ValidTo,omitempty"`
	Value      string `xml:"Value,omitempty"`
}

type ReportInfo struct {
	Created         time.Time `xml:"Created,omitempty"`
	ReferenceNumber string    `xml:"ReferenceNumber,omitempty"`
	ReportStatus    string    `xml:"ReportStatus,omitempty"`
	Version         string    `xml:"Version,omitempty"`
}

type StrategyResponse struct {
	Id           string `xml:"Id,omitempty"`
	Name         string `xml:"Name,omitempty"`
	BeeStrategy  string `xml:"BeeStrategy,omitempty"`
	TemplateName string `xml:"TemplateName,omitempty"`
}

type ErrorResponse struct {
	Type    string `xml:"type,attr,omitempty"`
	Id      string `xml:"id,attr,omitempty"`
	Message string `xml:"message,omitempty"`
}
