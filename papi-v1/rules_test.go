package papi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRule_AddBehavior(t *testing.T) {
	tests := []struct {
		Rule     Rule
		Behavior Behavior
		Expected Rule
	}{
		{
			Rule: Rule{

			},
			Behavior: Behavior{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Behavior: Behavior{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
					&Behavior{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Behavior: Behavior{
				Name: "existing",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.Rule.AddBehavior(&test.Behavior)
		assert.Equal(t, len(test.Rule.Behaviors), len(test.Expected.Behaviors))
		for key, behavior := range test.Rule.Behaviors {
			assert.Equal(t, behavior.Name, test.Expected.Behaviors[key].Name)
			assert.Equal(t, behavior.Options, test.Expected.Behaviors[key].Options)
		}
	}
}

func TestRule_MergeBehavior(t *testing.T) {
	tests := []struct {
		Rule     Rule
		Behavior Behavior
		Expected Rule
	}{
		{
			Rule: Rule{

			},
			Behavior: Behavior{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Behavior: Behavior{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
					&Behavior{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Behavior: Behavior{
				Name: "existing",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
							"otherkey": "won't change",
						},
					},
				},
			},
			Behavior: Behavior{
				Name: "existing",
				Options: OptionValue{
					"existingkey": "newvalue",
					"bar": "baz",
				},
			},
			Expected: Rule {
				Behaviors: []*Behavior {
					&Behavior{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "newvalue",
							"otherkey": "won't change",
							"bar": "baz",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.Rule.MergeBehavior(&test.Behavior)
		assert.Equal(t, len(test.Rule.Behaviors), len(test.Expected.Behaviors))
		for key, behavior := range test.Rule.Behaviors {
			assert.Equal(t, behavior.Name, test.Expected.Behaviors[key].Name)
			assert.Equal(t, behavior.Options, test.Expected.Behaviors[key].Options)
		}
	}
}

func TestRule_AddCriteria(t *testing.T) {
	tests := []struct {
		Rule     Rule
		Criteria Criteria
		Expected Rule
	}{
		{
			Rule: Rule{

			},
			Criteria: Criteria{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Criteria: Criteria{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
					&Criteria{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Criteria: Criteria{
				Name: "existing",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.Rule.AddCriteria(&test.Criteria)
		assert.Equal(t, len(test.Rule.Criteria), len(test.Expected.Criteria))
		for key, criteria := range test.Rule.Criteria {
			assert.Equal(t, criteria.Name, test.Expected.Criteria[key].Name)
			assert.Equal(t, criteria.Options, test.Expected.Criteria[key].Options)
		}
	}
}

func TestRule_MergeCriteria(t *testing.T) {
	tests := []struct {
		Rule     Rule
		Criteria Criteria
		Expected Rule
	}{
		{
			Rule: Rule{

			},
			Criteria: Criteria{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Criteria: Criteria{
				Name: "foo",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
					&Criteria{
						Name: "foo",
						Options: OptionValue{
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
						},
					},
				},
			},
			Criteria: Criteria{
				Name: "existing",
				Options: OptionValue{
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
							"bar": "baz",
						},
					},
				},
			},
		},
		{
			Rule: Rule{
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "existingvalue",
							"otherkey": "won't change",
						},
					},
				},
			},
			Criteria: Criteria{
				Name: "existing",
				Options: OptionValue{
					"existingkey": "newvalue",
					"bar": "baz",
				},
			},
			Expected: Rule {
				Criteria: []*Criteria {
					&Criteria{
						Name: "existing",
						Options: OptionValue{
							"existingkey": "newvalue",
							"otherkey": "won't change",
							"bar": "baz",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.Rule.MergeCriteria(&test.Criteria)
		assert.Equal(t, len(test.Rule.Criteria), len(test.Expected.Criteria))
		for key, criteria := range test.Rule.Criteria {
			assert.Equal(t, criteria.Name, test.Expected.Criteria[key].Name)
			assert.Equal(t, criteria.Options, test.Expected.Criteria[key].Options)
		}
	}
}

func TestRule_AddVariable(t *testing.T) {
	tests := []struct {
		ParentRule Rule
		Variables  []*Variable
		Expected   Rule
	}{
		{
			ParentRule: Rule{
				Name: "Parent Rule",
			},
			Variables: []*Variable{
				&Variable{
					Name: "Test Variable",
					Description: "Test Description",
					Value: "Test Value",
					Hidden: true,
					Sensitive: true,
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Variables: []*Variable {
					&Variable{
						Name: "Test Variable",
						Description: "Test Description",
						Value: "Test Value",
						Hidden: true,
						Sensitive: true,
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Variables: []*Variable {
					&Variable {
						Name: "Existing Variable",
					},
				},
			},
			Variables: []*Variable{
				&Variable{
					Name: "Existing Variable",
					Description: "New Description",
					Value: "New Value",
					Hidden: true,
					Sensitive: true,
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Variables: []*Variable{
					&Variable{
						Name: "Existing Variable",
						Description: "New Description",
						Value: "New Value",
						Hidden: true,
						Sensitive: true,
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Variables: []*Variable {
					&Variable {
						Name: "Existing Variable",
					},
				},
			},
			Variables: []*Variable{
				&Variable {
					Name: "Existing Variable",
					Description: "Updated Description",
				},
				&Variable{
					Name: "New Variable",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Variables: []*Variable {
					&Variable {
						Name: "Existing Variable",
						Description: "Updated Description",
					},
					&Variable{
						Name: "New Variable",
					},
				},
			},
		},
	}

	for _, test := range tests {
		for _, variable := range test.Variables {
			test.ParentRule.AddVariable(variable)
		}
	}
}

func TestRule_AddChildRule(t *testing.T) {
	tests := []struct {
		ParentRule Rule
		ChildRules  []*Rule
		Expected   Rule
	}{
		{
			ParentRule: Rule{
				Name: "Parent Rule",
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
					&Rule{
						Name: "Child Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Existing Child Rule",
							},
						},
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Children: []*Rule{
						&Rule{
							Name: "Sub-Child Child Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Existing Child Rule",
							},
						},
					},
					&Rule{
						Name: "Child Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Child Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
						},
					},
				},
				&Rule{
					Name: "Child Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Child Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
					&Rule{
						Name: "Child Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Child Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
							Children: []*Rule {
								&Rule {
									Name: "Sub-Sub-Child Existing Rule",
								},
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
								Children: []*Rule {
									&Rule {
										Name: "Sub-Sub-Child Existing Rule",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
		},
	}

	for _, test := range tests {
		for _, child := range test.ChildRules {
			test.ParentRule.AddChildRule(child)
		}

		assertRulesMatch(t, &test.ParentRule, &test.Expected)
	}
}

func TestRule_MergeChildRule(t *testing.T) {
	tests := []struct {
		ParentRule Rule
		ChildRules  []*Rule
		Expected   Rule
	}{
		{
			ParentRule: Rule{
				Name: "Parent Rule",
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
					&Rule{
						Name: "Child Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Existing Child Rule",
							},
						},
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Children: []*Rule{
						&Rule{
							Name: "Sub-Child Child Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Existing Child Rule",
							},
						},
					},
					&Rule{
						Name: "Child Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Child Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
						},
					},
				},
				&Rule{
					Name: "Child Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Child Rule",
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule {
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
					&Rule{
						Name: "Child Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Child Rule",
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
							},
						},
					},
				},
			},
			ChildRules: []*Rule {
				&Rule{
					Name: "Existing Rule",
					Children: []*Rule {
						&Rule {
							Name: "Sub-Child Existing Rule",
							Children: []*Rule {
								&Rule {
									Name: "Sub-Sub-Child Existing Rule",
								},
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Existing Rule",
						Children: []*Rule {
							&Rule {
								Name: "Sub-Child Existing Rule",
								Children: []*Rule {
									&Rule {
										Name: "Sub-Sub-Child Existing Rule",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option": "Value",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
		},
		{
			ParentRule: Rule{
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
			ChildRules: []*Rule{
				&Rule{
					Name: "Child Rule",
					Behaviors: []*Behavior {
						&Behavior{
							Name: "Behavior",
							Options: OptionValue{
								"Option2": "Value2",
							},
						},
					},
					Criteria: []*Criteria {
						&Criteria{
							Name: "Behavior",
							Options: OptionValue{
								"Option2": "Value2",
							},
						},
					},
				},
			},
			Expected: Rule {
				Name: "Parent Rule",
				Children: []*Rule {
					&Rule{
						Name: "Child Rule",
						Behaviors: []*Behavior {
							&Behavior{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
									"Option2": "Value2",
								},
							},
						},
						Criteria: []*Criteria {
							&Criteria{
								Name: "Behavior",
								Options: OptionValue{
									"Option": "Value",
									"Option2": "Value2",
								},
							},
						},
					},
					&Rule{
						Name: "Existing Rule",
					},
				},
			},
		},
	}

	for i, test := range tests {
		for _, child := range test.ChildRules {
			test.ParentRule.MergeChildRule(child)
		}

		if !assertRulesMatch(t, &test.ParentRule, &test.Expected) {
			t.Errorf("Data set %d failed!", i+1)
		}
	}
}

func assertRulesMatch(t *testing.T, expected *Rule, actual *Rule) bool {
	valid := true

	if !assert.Equal(t, expected.Name, actual.Name) {
		valid = false
	}
	if !assert.Equal(t, expected.Criteria, actual.Criteria) {
		valid = false
	}

	if !assert.Equal(t, expected.Behaviors, actual.Behaviors) {
		valid = false
	}

	if !assert.Equal(t, expected.Variables, actual.Variables) {
		valid = false
	}

	if !assert.Equal(t, len(expected.Children), len(actual.Children)) {
		valid = false
	}

	if len(expected.Children) > 0 {
		for key := 0; key < len(expected.Children); key++ {
			if !assertRulesMatch(t, expected.Children[key], actual.Children[key]) {
				valid = false
			}
		}
	}

	return valid
}