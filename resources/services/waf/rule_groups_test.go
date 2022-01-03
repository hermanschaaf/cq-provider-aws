// +build integration

package waf

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
)

func TestIntegrationWAFRuleGroups(t *testing.T) {
	client.AWSTestHelper(t, WafRuleGroups())
}