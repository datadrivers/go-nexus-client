package nexus3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamt1997/go-nexus-client/nexus3/schema"
)

func testRoutingRule(name string, mode schema.RoutingRuleMode) *schema.RoutingRule {
	return &schema.RoutingRule{
		Name:        name,
		Mode:        mode,
		Description: fmt.Sprintf("Go client routing rule %s", name),
		Matchers: []string{
			"match1",
		},
	}
}

func TestRoutingRuleModeIsValid(t *testing.T) {
	var mode schema.RoutingRuleMode

	mode = "allow"
	assert.Error(t, mode.IsValid(), "Invalid routing rule mode. Possible values are ALLOW or BLOCK")
	mode = "ALLOW"
	assert.Nil(t, mode.IsValid())
	mode = "BLOCK"
	assert.Nil(t, mode.IsValid())
}

func TestRoutingRuleCreateReadUpdateDelete(t *testing.T) {
	client := getTestClient()
	name := "test-routing-rule"
	createdRoutingRule := testRoutingRule(name, schema.RoutingRuleModeAllow)

	err := client.RoutingRule.Create(createdRoutingRule)
	assert.Nil(t, err)

	readRoutingRule, err := client.RoutingRule.Get(name)
	assert.Equal(t, createdRoutingRule, readRoutingRule)
	assert.Nil(t, err)

	updatedRoutingRule := testRoutingRule(name, schema.RoutingRuleModeBlock)

	err = client.RoutingRule.Update(updatedRoutingRule)
	assert.Nil(t, err)

	readRoutingRule, err = client.RoutingRule.Get(name)
	assert.Equal(t, updatedRoutingRule, readRoutingRule)
	assert.Nil(t, err)

	err = client.RoutingRule.Delete(name)
	assert.Nil(t, err)

	readRoutingRule, err = client.RoutingRule.Get(name)
	assert.Nil(t, readRoutingRule)
	assert.Error(t, err, "Did not find a routing rule with the name 'test-routing-rule'")
}
