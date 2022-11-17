package ad

import (
	"context"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAdGroupMembershipV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the group. This can be a GUID, a SID, a Distinguished Name, or the SAM Account Name of the group.",
				ForceNew:    true,
			},
			"group_members": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "A list of member AD Principals. Each principal can be identified by its GUID, SID, Distinguished Name, or SAM Account Name. Only one is required",
				Elem:        &schema.Schema{Type: schema.TypeString},
				MinItems:    1,
			},
		},
	}
}

func resourceAdGroupMembershipStateUpgradeV0(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {

	var (
		groupStateDelimiterOld string = "_"
		groupStateDelimiterNew string = "/"
	)

	log.Printf("[DEBUG] group_membership id before migration: %#v", rawState["id"])

	ids := strings.Split(rawState["id"].(string), groupStateDelimiterOld)

	rawState["id"] = ids[0] + groupStateDelimiterNew + ids[1]

	log.Printf("[DEBUG] group_membership id after migration: %#v", rawState["id"])
	return rawState, nil
}
