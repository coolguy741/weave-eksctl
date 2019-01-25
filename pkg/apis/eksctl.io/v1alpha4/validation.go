package v1alpha4

import (
	"fmt"
)

// ValidateNodeGroup checks compatible fileds of a given nodegroup
func ValidateNodeGroup(i int, ng *NodeGroup) error {
	path := fmt.Errorf("nodegroups[%d]", i)
	if ng.Name == "" {
		return fmt.Errorf("%s.name must be set", path)
	}

	if ng.IAM == nil {
		return nil
	}

	validate := func(value, fieldName string) error {
		if value != "" {
			p := fmt.Sprintf("%s.iam.%s and %s.iam", path, fieldName, path)
			if ng.IAM.InstanceRoleName != "" {
				return fmt.Errorf("%s.instanceRoleName cannot be set at the same time", p)
			}
			if len(ng.IAM.AttachPolicyARNs) != 0 {
				return fmt.Errorf("%s.attachPolicyARNs cannot be set at the same time", p)
			}
			if v := ng.IAM.WithAddonPolicies.AutoScaler; v != nil && *v {
				return fmt.Errorf("%s.withAddonPolicies.autoScaler cannot be set at the same time", p)
			}
			if v := ng.IAM.WithAddonPolicies.ExternalDNS; v != nil && *v {
				return fmt.Errorf("%s.withAddonPolicies.externalDNS cannot be set at the same time", p)
			}
			if v := ng.IAM.WithAddonPolicies.ImageBuilder; v != nil && *v {
				return fmt.Errorf("%s.imageBuilder cannot be set at the same time", p)
			}
		}
		return nil
	}

	if err := validate(ng.IAM.InstanceProfileARN, "instanceProfileARN"); err != nil {
		return err
	}
	if err := validate(ng.IAM.InstanceRoleARN, "instanceRoleARN"); err != nil {
		return err
	}

	return nil
}
