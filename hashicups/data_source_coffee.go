package hashicups

import (
  "context"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "time"

  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCoffees() *schema.Resource {
  return &schema.Resource{
    ReadContext: dataSourceCoffeesRead,
    Schema: map[string]*schema.Schema{
      Type: schema.TypeList,
      Computed: true,
      Elem: &schema.Resource{
        Schema: map[string]*schema.Schema{
          "id": &schema.Schema{
            Type: schema.TypeInt,
            Computed: true,
          },
          "name": &schema.Schema{
            Type: schema.TypeString,
            Computed: true,
          },
    },
  }
}
