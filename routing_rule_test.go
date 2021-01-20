package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testRoutingRule(name string, mode RoutingRuleMode) *RoutingRule {
	return &RoutingRule{
		Name:        name,
		Mode:        mode,
		Description: fmt.Sprintf("Go client routing roule %s", name),
		Matchers: []string{
			"match1",
		},
	}
}

func TestRoutingRuleModeIsValid(t *testing.T) {
	var mode RoutingRuleMode

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
	createdRoutingRule := testRoutingRule(name, RoutingRuleModeAllow)

	err := client.RoutingRuleCreate(createdRoutingRule)
	assert.Nil(t, err)

	readRoutingRule, err := client.RoutingRuleRead(name)
	assert.Equal(t, createdRoutingRule, readRoutingRule)
	assert.Nil(t, err)

	updatedRoutingRule := testRoutingRule(name, RoutingRuleModeBlock)

	err = client.RoutingRuleUpdate(updatedRoutingRule)
	assert.Nil(t, err)

	readRoutingRule, err = client.RoutingRuleRead(name)
	assert.Equal(t, updatedRoutingRule, readRoutingRule)
	assert.Nil(t, err)

	err = client.RoutingRuleDelete(name)
	assert.Nil(t, err)

	readRoutingRule, err = client.RoutingRuleRead(name)
	assert.Nil(t, readRoutingRule)
	assert.Error(t, err, "Did not find a routing rule with the name 'test-routing-rule'")
}
