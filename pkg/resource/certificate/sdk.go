// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package certificate

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/acm"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/acm-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ACM{}
	_ = &svcapitypes.Certificate{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeCertificateOutput
	resp, err = rm.sdkapi.DescribeCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeCertificate", err)
	if err != nil {
		if reqErr, ok := ackerr.AWSRequestFailure(err); ok && reqErr.StatusCode() == 404 {
			return nil, ackerr.NotFound
		}
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()
	if resp.Certificate.DomainValidationOptions != nil {
		dvs := []*svcapitypes.DomainValidation{}
		for _, dvsiter := range resp.Certificate.DomainValidationOptions {
			dvselem := &svcapitypes.DomainValidation{}
			if dvsiter.DomainName != nil {
				dvselem.DomainName = dvsiter.DomainName
			}
			if dvsiter.ResourceRecord != nil {
				dvselem.ResourceRecord = &svcapitypes.ResourceRecord{}
				if dvsiter.ResourceRecord.Name != nil {
					dvselem.ResourceRecord.Name = dvsiter.ResourceRecord.Name
				}
				if dvsiter.ResourceRecord.Type != nil {
					dvselem.ResourceRecord.Type = dvsiter.ResourceRecord.Type
				}
				if dvsiter.ResourceRecord.Value != nil {
					dvselem.ResourceRecord.Value = dvsiter.ResourceRecord.Value
				}
			}
			if dvsiter.ValidationDomain != nil {
				dvselem.ValidationDomain = dvsiter.ValidationDomain
			}
			if dvsiter.ValidationEmails != nil {
				dvselem.ValidationEmails = dvsiter.ValidationEmails
			}
			if dvsiter.ValidationMethod != nil {
				dvselem.ValidationMethod = dvsiter.ValidationMethod
			}
			if dvsiter.ValidationStatus != nil {
				dvselem.ValidationStatus = dvsiter.ValidationStatus
			}
			dvs = append(dvs, dvselem)
		}
		ko.Status.DomainValidations = dvs
	} else {
		ko.Status.DomainValidations = nil
	}

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Certificate.CertificateArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Certificate.CertificateArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Certificate.CertificateAuthorityArn != nil {
		ko.Spec.CertificateAuthorityARN = resp.Certificate.CertificateAuthorityArn
	} else {
		ko.Spec.CertificateAuthorityARN = nil
	}
	if resp.Certificate.CreatedAt != nil {
		ko.Status.CreatedAt = &metav1.Time{*resp.Certificate.CreatedAt}
	} else {
		ko.Status.CreatedAt = nil
	}
	if resp.Certificate.DomainName != nil {
		ko.Spec.DomainName = resp.Certificate.DomainName
	} else {
		ko.Spec.DomainName = nil
	}
	if resp.Certificate.DomainValidationOptions != nil {
		f4 := []*svcapitypes.DomainValidationOption{}
		for _, f4iter := range resp.Certificate.DomainValidationOptions {
			f4elem := &svcapitypes.DomainValidationOption{}
			if f4iter.DomainName != nil {
				f4elem.DomainName = f4iter.DomainName
			}
			if f4iter.ValidationDomain != nil {
				f4elem.ValidationDomain = f4iter.ValidationDomain
			}
			f4 = append(f4, f4elem)
		}
		ko.Spec.DomainValidationOptions = f4
	} else {
		ko.Spec.DomainValidationOptions = nil
	}
	if resp.Certificate.ExtendedKeyUsages != nil {
		f5 := []*svcapitypes.ExtendedKeyUsage{}
		for _, f5iter := range resp.Certificate.ExtendedKeyUsages {
			f5elem := &svcapitypes.ExtendedKeyUsage{}
			if f5iter.Name != nil {
				f5elem.Name = f5iter.Name
			}
			if f5iter.OID != nil {
				f5elem.OID = f5iter.OID
			}
			f5 = append(f5, f5elem)
		}
		ko.Status.ExtendedKeyUsages = f5
	} else {
		ko.Status.ExtendedKeyUsages = nil
	}
	if resp.Certificate.FailureReason != nil {
		ko.Status.FailureReason = resp.Certificate.FailureReason
	} else {
		ko.Status.FailureReason = nil
	}
	if resp.Certificate.ImportedAt != nil {
		ko.Status.ImportedAt = &metav1.Time{*resp.Certificate.ImportedAt}
	} else {
		ko.Status.ImportedAt = nil
	}
	if resp.Certificate.InUseBy != nil {
		f8 := []*string{}
		for _, f8iter := range resp.Certificate.InUseBy {
			var f8elem string
			f8elem = *f8iter
			f8 = append(f8, &f8elem)
		}
		ko.Status.InUseBy = f8
	} else {
		ko.Status.InUseBy = nil
	}
	if resp.Certificate.IssuedAt != nil {
		ko.Status.IssuedAt = &metav1.Time{*resp.Certificate.IssuedAt}
	} else {
		ko.Status.IssuedAt = nil
	}
	if resp.Certificate.Issuer != nil {
		ko.Status.Issuer = resp.Certificate.Issuer
	} else {
		ko.Status.Issuer = nil
	}
	if resp.Certificate.KeyAlgorithm != nil {
		ko.Spec.KeyAlgorithm = resp.Certificate.KeyAlgorithm
	} else {
		ko.Spec.KeyAlgorithm = nil
	}
	if resp.Certificate.KeyUsages != nil {
		f12 := []*svcapitypes.KeyUsage{}
		for _, f12iter := range resp.Certificate.KeyUsages {
			f12elem := &svcapitypes.KeyUsage{}
			if f12iter.Name != nil {
				f12elem.Name = f12iter.Name
			}
			f12 = append(f12, f12elem)
		}
		ko.Status.KeyUsages = f12
	} else {
		ko.Status.KeyUsages = nil
	}
	if resp.Certificate.NotAfter != nil {
		ko.Status.NotAfter = &metav1.Time{*resp.Certificate.NotAfter}
	} else {
		ko.Status.NotAfter = nil
	}
	if resp.Certificate.NotBefore != nil {
		ko.Status.NotBefore = &metav1.Time{*resp.Certificate.NotBefore}
	} else {
		ko.Status.NotBefore = nil
	}
	if resp.Certificate.Options != nil {
		f15 := &svcapitypes.CertificateOptions{}
		if resp.Certificate.Options.CertificateTransparencyLoggingPreference != nil {
			f15.CertificateTransparencyLoggingPreference = resp.Certificate.Options.CertificateTransparencyLoggingPreference
		}
		ko.Spec.Options = f15
	} else {
		ko.Spec.Options = nil
	}
	if resp.Certificate.RenewalEligibility != nil {
		ko.Status.RenewalEligibility = resp.Certificate.RenewalEligibility
	} else {
		ko.Status.RenewalEligibility = nil
	}
	if resp.Certificate.RenewalSummary != nil {
		f17 := &svcapitypes.RenewalSummary{}
		if resp.Certificate.RenewalSummary.DomainValidationOptions != nil {
			f17f0 := []*svcapitypes.DomainValidation{}
			for _, f17f0iter := range resp.Certificate.RenewalSummary.DomainValidationOptions {
				f17f0elem := &svcapitypes.DomainValidation{}
				if f17f0iter.DomainName != nil {
					f17f0elem.DomainName = f17f0iter.DomainName
				}
				if f17f0iter.ResourceRecord != nil {
					f17f0elemf1 := &svcapitypes.ResourceRecord{}
					if f17f0iter.ResourceRecord.Name != nil {
						f17f0elemf1.Name = f17f0iter.ResourceRecord.Name
					}
					if f17f0iter.ResourceRecord.Type != nil {
						f17f0elemf1.Type = f17f0iter.ResourceRecord.Type
					}
					if f17f0iter.ResourceRecord.Value != nil {
						f17f0elemf1.Value = f17f0iter.ResourceRecord.Value
					}
					f17f0elem.ResourceRecord = f17f0elemf1
				}
				if f17f0iter.ValidationDomain != nil {
					f17f0elem.ValidationDomain = f17f0iter.ValidationDomain
				}
				if f17f0iter.ValidationEmails != nil {
					f17f0elemf3 := []*string{}
					for _, f17f0elemf3iter := range f17f0iter.ValidationEmails {
						var f17f0elemf3elem string
						f17f0elemf3elem = *f17f0elemf3iter
						f17f0elemf3 = append(f17f0elemf3, &f17f0elemf3elem)
					}
					f17f0elem.ValidationEmails = f17f0elemf3
				}
				if f17f0iter.ValidationMethod != nil {
					f17f0elem.ValidationMethod = f17f0iter.ValidationMethod
				}
				if f17f0iter.ValidationStatus != nil {
					f17f0elem.ValidationStatus = f17f0iter.ValidationStatus
				}
				f17f0 = append(f17f0, f17f0elem)
			}
			f17.DomainValidationOptions = f17f0
		}
		if resp.Certificate.RenewalSummary.RenewalStatus != nil {
			f17.RenewalStatus = resp.Certificate.RenewalSummary.RenewalStatus
		}
		if resp.Certificate.RenewalSummary.RenewalStatusReason != nil {
			f17.RenewalStatusReason = resp.Certificate.RenewalSummary.RenewalStatusReason
		}
		if resp.Certificate.RenewalSummary.UpdatedAt != nil {
			f17.UpdatedAt = &metav1.Time{*resp.Certificate.RenewalSummary.UpdatedAt}
		}
		ko.Status.RenewalSummary = f17
	} else {
		ko.Status.RenewalSummary = nil
	}
	if resp.Certificate.RevocationReason != nil {
		ko.Status.RevocationReason = resp.Certificate.RevocationReason
	} else {
		ko.Status.RevocationReason = nil
	}
	if resp.Certificate.RevokedAt != nil {
		ko.Status.RevokedAt = &metav1.Time{*resp.Certificate.RevokedAt}
	} else {
		ko.Status.RevokedAt = nil
	}
	if resp.Certificate.Serial != nil {
		ko.Status.Serial = resp.Certificate.Serial
	} else {
		ko.Status.Serial = nil
	}
	if resp.Certificate.SignatureAlgorithm != nil {
		ko.Status.SignatureAlgorithm = resp.Certificate.SignatureAlgorithm
	} else {
		ko.Status.SignatureAlgorithm = nil
	}
	if resp.Certificate.Status != nil {
		ko.Status.Status = resp.Certificate.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.Certificate.Subject != nil {
		ko.Status.Subject = resp.Certificate.Subject
	} else {
		ko.Status.Subject = nil
	}
	if resp.Certificate.SubjectAlternativeNames != nil {
		f24 := []*string{}
		for _, f24iter := range resp.Certificate.SubjectAlternativeNames {
			var f24elem string
			f24elem = *f24iter
			f24 = append(f24, &f24elem)
		}
		ko.Spec.SubjectAlternativeNames = f24
	} else {
		ko.Spec.SubjectAlternativeNames = nil
	}
	if resp.Certificate.Type != nil {
		ko.Status.Type = resp.Certificate.Type
	} else {
		ko.Status.Type = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return (r.ko.Status.ACKResourceMetadata == nil || r.ko.Status.ACKResourceMetadata.ARN == nil)

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeCertificateInput, error) {
	res := &svcsdk.DescribeCertificateInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	if err = validatePublicValidationOptions(desired); err != nil {
		ackcondition.SetTerminal(
			desired,
			corev1.ConditionTrue,
			&domainValidationOptionsExceededMsg,
			nil,
		)
		return desired, nil
	}

	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}
	input.SetValidationMethod("DNS")

	var resp *svcsdk.RequestCertificateOutput
	_ = resp
	resp, err = rm.sdkapi.RequestCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "RequestCertificate", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CertificateArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CertificateArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.RequestCertificateInput, error) {
	res := &svcsdk.RequestCertificateInput{}

	if r.ko.Spec.CertificateAuthorityARN != nil {
		res.SetCertificateAuthorityArn(*r.ko.Spec.CertificateAuthorityARN)
	}
	if r.ko.Spec.DomainName != nil {
		res.SetDomainName(*r.ko.Spec.DomainName)
	}
	if r.ko.Spec.DomainValidationOptions != nil {
		f2 := []*svcsdk.DomainValidationOption{}
		for _, f2iter := range r.ko.Spec.DomainValidationOptions {
			f2elem := &svcsdk.DomainValidationOption{}
			if f2iter.DomainName != nil {
				f2elem.SetDomainName(*f2iter.DomainName)
			}
			if f2iter.ValidationDomain != nil {
				f2elem.SetValidationDomain(*f2iter.ValidationDomain)
			}
			f2 = append(f2, f2elem)
		}
		res.SetDomainValidationOptions(f2)
	}
	if r.ko.Spec.KeyAlgorithm != nil {
		res.SetKeyAlgorithm(*r.ko.Spec.KeyAlgorithm)
	}
	if r.ko.Spec.Options != nil {
		f4 := &svcsdk.CertificateOptions{}
		if r.ko.Spec.Options.CertificateTransparencyLoggingPreference != nil {
			f4.SetCertificateTransparencyLoggingPreference(*r.ko.Spec.Options.CertificateTransparencyLoggingPreference)
		}
		res.SetOptions(f4)
	}
	if r.ko.Spec.SubjectAlternativeNames != nil {
		f5 := []*string{}
		for _, f5iter := range r.ko.Spec.SubjectAlternativeNames {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetSubjectAlternativeNames(f5)
	}
	if r.ko.Spec.Tags != nil {
		f6 := []*svcsdk.Tag{}
		for _, f6iter := range r.ko.Spec.Tags {
			f6elem := &svcsdk.Tag{}
			if f6iter.Key != nil {
				f6elem.SetKey(*f6iter.Key)
			}
			if f6iter.Value != nil {
				f6elem.SetValue(*f6iter.Value)
			}
			f6 = append(f6, f6elem)
		}
		res.SetTags(f6)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateCertificateOptionsOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateCertificateOptionsWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateCertificateOptions", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateCertificateOptionsInput, error) {
	res := &svcsdk.UpdateCertificateOptionsInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}
	if r.ko.Spec.Options != nil {
		f1 := &svcsdk.CertificateOptions{}
		if r.ko.Spec.Options.CertificateTransparencyLoggingPreference != nil {
			f1.SetCertificateTransparencyLoggingPreference(*r.ko.Spec.Options.CertificateTransparencyLoggingPreference)
		}
		res.SetOptions(f1)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteCertificateOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteCertificateWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteCertificate", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteCertificateInput, error) {
	res := &svcsdk.DeleteCertificateInput{}

	if r.ko.Status.ACKResourceMetadata != nil && r.ko.Status.ACKResourceMetadata.ARN != nil {
		res.SetCertificateArn(string(*r.ko.Status.ACKResourceMetadata.ARN))
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Certificate,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	if syncCondition == nil && onSuccess {
		syncCondition = &ackv1alpha1.Condition{
			Type:   ackv1alpha1.ConditionTypeResourceSynced,
			Status: corev1.ConditionTrue,
		}
		ko.Status.Conditions = append(ko.Status.Conditions, syncCondition)
	}
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameter",
		"InvalidDomainValidationOptionsException",
		"InvalidTagException",
		"TagPolicyException",
		"TooManyTagsException":
		return true
	default:
		return false
	}
}
