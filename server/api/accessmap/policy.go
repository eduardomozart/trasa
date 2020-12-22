package accessmap

import (
	"database/sql"
	"github.com/pkg/errors"
	"github.com/seknox/trasa/server/api/policies"
	"github.com/seknox/trasa/server/models"
)

func GetAssignedPolicy(params *models.ConnectionParams) (*models.Policy, bool, error) {

	//try to get normal policy
	policy, adhoc, err := policies.Store.GetAccessPolicy(params.UserID, params.ServiceID, params.Privilege, params.OrgID)
	if errors.Is(err, sql.ErrNoRows) {
		//If not found, get policy from group names (in case of 3rd party IDP)
		policy, adhoc, err = policies.Store.GetAccessPolicyFromGroupNames(params.Groups, params.ServiceID, params.Privilege, params.OrgID)
		//If there is non nil error but not empty rows error
		if errors.Is(err, sql.ErrNoRows) {
			//if service is not assigned to user, create one (only if dynamic access is enabled)
			policy, err = CreateDynamicAccessMap(params.ServiceID, params.UserID, params.TrasaID, params.Privilege, params.OrgID)
			if err != nil {
				return policy, adhoc, errors.Errorf("dynamic access map: %v", err)
			}
		} else if err != nil {
			return policy, adhoc, errors.Errorf("get access policy from group names: %v", err)
		}

	} else if err != nil {
		return policy, adhoc, errors.Errorf("get access policy: %v", err)
	}

	return policy, adhoc, nil
}
