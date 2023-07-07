package controllers

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	securityv1alpha1 "github.com/nirmata/kyverno-aws-adapter/api/v1alpha1"
)

func CreateSecurityHubFinding(ctx context.Context, namespace, clusterName, regionName, companyName, accountId, findingType string, cfg aws.Config, kpv securityv1alpha1.KyvernoPolicyViolation) {
	// regionName := "us-west-1"
	// companyName := "Nirmata"
	// accountId := "844333597536"
	log.Println("account id ", accountId)
	log.Println("+=+=+ CREATE SECURITY HUB FINDING +=+=+")
	svc := securityhub.NewFromConfig(cfg)

	findings := []types.AwsSecurityFinding{}
	title := fmt.Sprintf("Kyverno Policy Violation: %s/%s in namespace '%s'", kpv.Policy, kpv.Rule, kpv.Namespace)
	noteText := fmt.Sprintf("APIVersion: %s\nKind: %s\nNamespace: %s\nName: %s\nUID: %s\nPolicy: %s\nRule: %s\nMessage: %s",
		kpv.APIVersion,
		kpv.Kind,
		kpv.Namespace,
		kpv.Name,
		kpv.UID,
		kpv.Policy,
		kpv.Rule,
		kpv.Message)
	description := fmt.Sprintf("Kyverno Policy Violation occurred. Details:\n%s", noteText)
	created := time.Now().UTC().Format(time.RFC3339)
	updated := created
	generatorId := accountId + "/default"
	productArn := "arn:aws:securityhub:" + regionName + ":" + accountId + ":product/" + generatorId
	// uuId := uuid.New()
	// testId := testProductArn + "/finding/" + uuId.String()
	schemaVersion := "2018-10-08"
	resourceType := "AwsEksCluster"
	resourceTypes := []types.Resource{types.Resource{Id: &clusterName, Type: &resourceType}}
	findingId := productArn + "/finding/" + fmt.Sprintf("%x", sha1.Sum([]byte(noteText)))

	severity := new(string)
	criticality := new(int32)
	// complianceStatus := new(types.ComplianceStatus)
	switch kpv.Severity {
	case "info":
		*severity = "INFORMATIONAL"
		*criticality = 0
		// *complianceStatus = types.ComplianceStatusPassed
	case "low":
		*severity = "LOW"
		*criticality = 1
		// *complianceStatus = types.ComplianceStatusWarning
	case "medium":
		*severity = "MEDIUM"
		*criticality = 40
		// *complianceStatus = types.ComplianceStatusFailed
	case "high":
		*severity = "HIGH"
		*criticality = 70
		// *complianceStatus = types.ComplianceStatusFailed
	case "critical":
		*severity = "CRITICAL"
		*criticality = 90
		// *complianceStatus = types.ComplianceStatusFailed
	}
	workflowStatusNew := types.Workflow{Status: types.WorkflowStatusNew}
	newFinding := types.AwsSecurityFinding{
		AwsAccountId:  &accountId,
		Title:         &title,
		Description:   &description,
		CreatedAt:     &created,
		UpdatedAt:     &updated,
		Id:            &findingId,
		GeneratorId:   &generatorId,
		ProductArn:    &productArn,
		SchemaVersion: &schemaVersion,
		Resources:     resourceTypes,
		Criticality:   *criticality,
		Severity: &types.Severity{
			Label: types.SeverityLabel(*severity),
		},
		Types: []string{findingType},
		FindingProviderFields: &types.FindingProviderFields{
			Severity: &types.FindingProviderSeverity{
				Label: types.SeverityLabel(*severity),
			},
			Types: []string{findingType},
		},
		Note: &types.Note{
			Text:      &noteText,
			UpdatedAt: &updated,
			UpdatedBy: &companyName,
		},
		CompanyName: &companyName,
		Workflow:    &workflowStatusNew,
	}
	// Compliance:  &types.Compliance{Status: types.ComplianceStatus(*complianceStatus)},
	findingsList := append(findings, newFinding)
	log.Println("+=+=+ FINDINGS LIST +=+=+ ", findingsList)

	_, err := svc.BatchImportFindings(ctx, &securityhub.BatchImportFindingsInput{Findings: findingsList})
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("importFindingsOutput:\nErrorCode: %v\nErrorMessage: %v\n", *importFindingsOutput.FailedFindings[0].ErrorCode, *importFindingsOutput.FailedFindings[0].ErrorMessage)
}

// AwsAccountId		*string				x
// CreatedAt		*string				x
// Description		*string				x
// GeneratorId		*string				x
// Id				*stringrequired		x
// ProductArn		*string				x
// Resources		[]Resource			~
// SchemaVersion	*string				x
// Title			*string				x
// UpdatedAt		*string				x

func FindingExists(ctx context.Context, cfg aws.Config, findingId string) bool {
	exists := new(bool)
	svc := securityhub.NewFromConfig(cfg)
	findingsOutput, err := svc.GetFindings(ctx, &securityhub.GetFindingsInput{
		Filters: &types.AwsSecurityFindingFilters{
			Id: []types.StringFilter{types.StringFilter{Comparison: types.StringFilterComparisonEquals, Value: &findingId}},
		}})
	if err != nil {
		log.Fatal(err)
	}

	if len(findingsOutput.Findings) > 0 {
		// switch findingsOutput.Findings[0].Workflow.Status {
		// case "NEW", "NOTIFIED", "SUPPRESSED":
		// 	*exists = true
		// }
		*exists = true
	}

	return *exists
}
