package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	quota "github.com/aws/aws-sdk-go/service/servicequotas"
)

func (p *Aws) GetQuota(ServiceCode string, QuotaCode string) (*quota.ServiceQuota, error) {
	svc := quota.New(p.Sess)
	rt, err := svc.GetServiceQuota(&quota.GetServiceQuotaInput{
		ServiceCode: aws.String(ServiceCode),
		QuotaCode:   aws.String(QuotaCode),
	})
	if err != nil {
		return nil, err
	}
	return rt.Quota, nil
}
func (p *Aws) ChangeQuota(ServiceCode string, QuotaCode string, DesiredValue float64) error {
	svc := quota.New(p.Sess)
	_, err := svc.RequestServiceQuotaIncrease(&quota.RequestServiceQuotaIncreaseInput{
		ServiceCode:  aws.String(ServiceCode),
		QuotaCode:    aws.String(QuotaCode),
		DesiredValue: aws.Float64(DesiredValue),
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Aws) ListChangeQuota(ServiceCode string, QuotaCode string) ([]*quota.RequestedServiceQuotaChange, error) {
	svc := quota.New(p.Sess)
	rt, err := svc.ListRequestedServiceQuotaChangeHistoryByQuota(&quota.ListRequestedServiceQuotaChangeHistoryByQuotaInput{
		ServiceCode: aws.String(ServiceCode),
		QuotaCode:   aws.String(QuotaCode),
	})
	if err != nil {
		return nil, err
	}
	return rt.RequestedQuotas, nil
}
