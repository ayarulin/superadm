package core


type ActiveIntegrationId int
type AccountId int
type ExtCompanyId string
type CompanyName string


func NewActiveIntegrationId(v int32) ActiveIntegrationId {
	return ActiveIntegrationId(v)
}

func (a ActiveIntegrationId) Int32() int32 { return int32(a) }

func NewAccountId(v int32) AccountId {
	return AccountId(v)
}

func (n AccountId) Int32() int32 {
	return int32(n)
}

func NewExtCompanyId(s string) (ExtCompanyId, error) {
	// err := validators.String.IsNotBlank(s)
	//
	// if err != nil {
	// 	return "", fmt.Errorf("ExtCompanyId: %v", err)
	// }

	return ExtCompanyId(s), nil
}

func (n ExtCompanyId) String() string {
	return string(n)
}

func NewCompanyName(s string) (CompanyName, error) {
	// err := validators.String.IsNotBlank(s)
	//
	// if err != nil {
	// 	return "", fmt.Errorf("CompanyName: %v", err)
	// }

	return CompanyName(s), nil
}

func (n CompanyName) String() string {
	return string(n)
}
