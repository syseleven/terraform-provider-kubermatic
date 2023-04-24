package metakube

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/syseleven/go-metakube/client/project"
)

func dataSourceMetakubeProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: metakubeDataSourceProjectRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

func metakubeDataSourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	meta := m.(*metakubeProviderMeta)

	p := project.NewListProjectsParams().WithContext(ctx)
	res, err := meta.client.Project.ListProjects(p, meta.auth)
	if err != nil {
		return diag.FromErr(err)
	}

	name := d.Get("name").(string)
	for _, r := range res.Payload {
		if r != nil && r.Name == name {
			d.SetId(r.ID)
			return nil
		}
	}
	return nil
}
