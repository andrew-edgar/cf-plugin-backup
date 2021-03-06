package ccv2

import (
	"encoding/json"

	"code.cloudfoundry.org/cli/api/cloudcontroller"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2/internal"
)

// ServicePlan represents a predefined set of configurations for a Cloud
// Controller service object.
type ServicePlan struct {
	//GUID is the unique identifier of the service plan.
	GUID string

	// Name is the name of the service plan.
	Name string

	// ServiceGUID is the unique identifier of the service that the service
	// plan belongs to.
	ServiceGUID string
}

// UnmarshalJSON helps unmarshal a Cloud Controller Service Plan response.
func (servicePlan *ServicePlan) UnmarshalJSON(data []byte) error {
	var ccServicePlan struct {
		Metadata internal.Metadata
		Entity   struct {
			Name        string `json:"name"`
			ServiceGUID string `json:"service_guid"`
		}
	}
	err := json.Unmarshal(data, &ccServicePlan)
	if err != nil {
		return err
	}

	servicePlan.GUID = ccServicePlan.Metadata.GUID
	servicePlan.Name = ccServicePlan.Entity.Name
	servicePlan.ServiceGUID = ccServicePlan.Entity.ServiceGUID
	return nil
}

// GetServicePlan returns the service plan with the given GUID.
func (client *Client) GetServicePlan(servicePlanGUID string) (ServicePlan, Warnings, error) {
	request, err := client.newHTTPRequest(requestOptions{
		RequestName: internal.GetServicePlanRequest,
		URIParams:   Params{"service_plan_guid": servicePlanGUID},
	})
	if err != nil {
		return ServicePlan{}, nil, err
	}

	var servicePlan ServicePlan
	response := cloudcontroller.Response{
		Result: &servicePlan,
	}

	err = client.connection.Make(request, &response)
	return servicePlan, response.Warnings, err
}
