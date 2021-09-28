package validation

import (
	"fmt"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/openshift/installer/pkg/types/ibmcloud"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateMachinePool validates the MachinePool.
func ValidateMachinePool(platform *ibmcloud.Platform, mp *ibmcloud.MachinePool, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	for i, zone := range mp.Zones {
		if !strings.HasPrefix(zone, platform.Region) {
			allErrs = append(allErrs, field.Invalid(path.Child("zones").Index(i), zone, fmt.Sprintf("zone not in configured region (%s)", platform.Region)))
		}
	}

	if mp.BootVolume != nil {
		allErrs = append(allErrs, validateBootVolume(mp.BootVolume, path.Child("bootVolume"))...)
	}
	return allErrs
}

func validateBootVolume(bv *ibmcloud.BootVolume, path *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if bv.EncryptionKey != "" {
		_, parseErr := crn.Parse(bv.EncryptionKey)
		if parseErr != nil {
			allErrs = append(allErrs, field.Invalid(path.Child("encryptionKey"), bv.EncryptionKey, "encryptionKey is not a valid IBM CRN"))
		}
	}
	return allErrs
}
