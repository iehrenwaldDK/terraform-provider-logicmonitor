package logicmonitor

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	lmclient "github.com/logicmonitor/lm-sdk-go/client"
	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

// terraform data source to look up dashboard groups
func dataSourceFindDashboardGroups() *schema.Resource {
	return &schema.Resource{
		Read: findDashboardGroup,

		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"property": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"operator": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_property_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"custom_property_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"size": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"offset": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}

// function to find dashboard groups with certain filters
func findDashboardGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*lmclient.LMSdkGo)
	filter := getFilters(d)
	size := int32(d.Get("size").(int))
	offset := int32(d.Get("offset").(int))

	params := lm.NewGetDashboardGroupListParams()
	params.SetFilter(&filter)
	params.SetOffset(&offset)
	params.SetSize(&size)

	//looks for dashboard group
	restDashboardGroupPaginationResponse, err := client.LM.GetDashboardGroupList(params)
	if err != nil {
		return err
	}

	var dashboardIds []string
	for _, d := range restDashboardGroupPaginationResponse.Payload.Items {
		log.Printf("Found dashboard group with filter %q", filter)
		dashboardIds = append(dashboardIds, strconv.Itoa(int(d.ID)))
	}

	if len(dashboardIds) > 0 {
		//comma separated string of device Ids
		var ids = strings.Join(dashboardIds, ",")
		d.SetId(ids)
	} else {
		err := fmt.Errorf("Found no dashboard groups with filter %s", filter)
		return err
	}
	return nil
}
