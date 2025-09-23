package expr

import (
	"testing"

	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestToProto(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		expr     Expr
		expected *expressionv1.LogicalExprNode
	}{
		// Nil expression
		{
			name:     "nil_expression",
			expr:     nil,
			expected: nil,
		},

		// Literals
		{
			name: "integer_literal",
			expr: Int(42),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_Int64Value{
								Int64Value: 42,
							},
						},
						IsArrowScalarObject: false,
					},
				},
			},
		},
		{
			name: "string_literal",
			expr: String("hello"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_Utf8Value{
								Utf8Value: "hello",
							},
						},
						IsArrowScalarObject: false,
					},
				},
			},
		},
		{
			name: "boolean_literal",
			expr: Bool(true),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_BoolValue{
								BoolValue: true,
							},
						},
						IsArrowScalarObject: false,
					},
				},
			},
		},
		{
			name: "null_literal",
			expr: Null(),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_NullValue{
								NullValue: &arrowv1.ArrowType{},
							},
						},
						IsArrowScalarObject: false,
					},
				},
			},
		},
		{
			name: "float_literal",
			expr: Float64(3.14),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
					LiteralValue: &expressionv1.ExprLiteral{
						Value: &arrowv1.ScalarValue{
							Value: &arrowv1.ScalarValue_Float64Value{
								Float64Value: 3.14,
							},
						},
						IsArrowScalarObject: false,
					},
				},
			},
		},

		// Columns
		{
			name: "simple_column",
			expr: Col("user_id"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Identifier{
					Identifier: &expressionv1.Identifier{
						Name: "user_id",
					},
				},
			},
		},
		{
			name: "column_with_relation",
			expr: Col("name").(*ColumnExpr).WithRelation("user"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Identifier{
					Identifier: &expressionv1.Identifier{
						Name: "user.name",
					},
				},
			},
		},

		// Attribute access
		{
			name: "get_attribute",
			expr: Col("user").Attr("name"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
					GetAttribute: &expressionv1.ExprGetAttribute{
						Parent: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "user",
								},
							},
						},
						Attribute: &expressionv1.Identifier{
							Name: "name",
						},
					},
				},
			},
		},

		// Function calls
		{
			name: "function_call",
			expr: FunctionCall("abs", Col("value")),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "abs",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "value",
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},
		{
			name: "function_call_with_string_arg",
			expr: FunctionCall("upper", String("text")),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "upper",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
									LiteralValue: &expressionv1.ExprLiteral{
										Value: &arrowv1.ScalarValue{
											Value: &arrowv1.ScalarValue_Utf8Value{
												Utf8Value: "text",
											},
										},
										IsArrowScalarObject: false,
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},

		// Binary operations
		{
			name: "binary_operation_add",
			expr: Int(4).Add(Int(5)),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "+",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
									LiteralValue: &expressionv1.ExprLiteral{
										Value: &arrowv1.ScalarValue{
											Value: &arrowv1.ScalarValue_Int64Value{
												Int64Value: 4,
											},
										},
										IsArrowScalarObject: false,
									},
								},
							},
							{
								ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
									LiteralValue: &expressionv1.ExprLiteral{
										Value: &arrowv1.ScalarValue{
											Value: &arrowv1.ScalarValue_Int64Value{
												Int64Value: 5,
											},
										},
										IsArrowScalarObject: false,
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},

		// Comparison operations
		{
			name: "comparison_greater_than",
			expr: Col("age").Gt(Int(18)),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: ">",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "age",
									},
								},
							},
							{
								ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
									LiteralValue: &expressionv1.ExprLiteral{
										Value: &arrowv1.ScalarValue{
											Value: &arrowv1.ScalarValue_Int64Value{
												Int64Value: 18,
											},
										},
										IsArrowScalarObject: false,
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},

		// Complex nested expression
		{
			name: "complex_nested_expression",
			expr: Col("user").Attr("age").Add(Int(1)).Gt(Int(25)),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: ">",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Call{
									Call: &expressionv1.ExprCall{
										Func: &expressionv1.LogicalExprNode{
											ExprForm: &expressionv1.LogicalExprNode_Identifier{
												Identifier: &expressionv1.Identifier{
													Name: "+",
												},
											},
										},
										Args: []*expressionv1.LogicalExprNode{
											{
												ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
													GetAttribute: &expressionv1.ExprGetAttribute{
														Parent: &expressionv1.LogicalExprNode{
															ExprForm: &expressionv1.LogicalExprNode_Identifier{
																Identifier: &expressionv1.Identifier{
																	Name: "user",
																},
															},
														},
														Attribute: &expressionv1.Identifier{
															Name: "age",
														},
													},
												},
											},
											{
												ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
													LiteralValue: &expressionv1.ExprLiteral{
														Value: &arrowv1.ScalarValue{
															Value: &arrowv1.ScalarValue_Int64Value{
																Int64Value: 1,
															},
														},
														IsArrowScalarObject: false,
													},
												},
											},
										},
										Kwargs: make(map[string]*expressionv1.LogicalExprNode),
									},
								},
							},
							{
								ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
									LiteralValue: &expressionv1.ExprLiteral{
										Value: &arrowv1.ScalarValue{
											Value: &arrowv1.ScalarValue_Int64Value{
												Int64Value: 25,
											},
										},
										IsArrowScalarObject: false,
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},

		// Alias expression
		{
			name: "alias_expression",
			expr: Col("name").As("full_name"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Identifier{
					Identifier: &expressionv1.Identifier{
						Name: "name",
					},
				},
			},
		},

		// DataFrame expression
		{
			name: "dataframe_expression",
			expr: DataFrame("users").(Expr),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Identifier{
					Identifier: &expressionv1.Identifier{
						Name: "users",
					},
				},
			},
		},

		// Aggregation expression
		{
			name: "aggregation_expression",
			expr: DataFrame("transactions").Filter(Col("amount").Gt(Int(100))).Agg("sum"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
								GetAttribute: &expressionv1.ExprGetAttribute{
									Attribute: &expressionv1.Identifier{
										Name: "sum",
									},
									Parent: &expressionv1.LogicalExprNode{
										ExprForm: &expressionv1.LogicalExprNode_GetSubscript{
											GetSubscript: &expressionv1.ExprGetSubscript{
												Parent: &expressionv1.LogicalExprNode{
													ExprForm: &expressionv1.LogicalExprNode_Identifier{
														Identifier: &expressionv1.Identifier{
															Name: "transactions",
														},
													},
												},
												Subscript: []*expressionv1.LogicalExprNode{
													{
														ExprForm: &expressionv1.LogicalExprNode_Call{
															Call: &expressionv1.ExprCall{
																Func: &expressionv1.LogicalExprNode{
																	ExprForm: &expressionv1.LogicalExprNode_Identifier{
																		Identifier: &expressionv1.Identifier{
																			Name: ">",
																		},
																	},
																},
																Args: []*expressionv1.LogicalExprNode{
																	{
																		ExprForm: &expressionv1.LogicalExprNode_Identifier{
																			Identifier: &expressionv1.Identifier{
																				Name: "amount",
																			},
																		},
																	},
																	{
																		ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
																			LiteralValue: &expressionv1.ExprLiteral{
																				Value: &arrowv1.ScalarValue{
																					Value: &arrowv1.ScalarValue_Int64Value{
																						Int64Value: 100,
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},

		// Logical operations
		{
			name: "logical_and",
			expr: Col("active").And(Col("verified")),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "AND",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "active",
									},
								},
							},
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "verified",
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},

		// Null checks
		{
			name: "is_null",
			expr: Col("email").IsNull(),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "IS_NULL",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "email",
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},
		{
			name: "is_not_null",
			expr: Col("phone").IsNotNull(),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_Identifier{
								Identifier: &expressionv1.Identifier{
									Name: "IS_NOT_NULL",
								},
							},
						},
						Args: []*expressionv1.LogicalExprNode{
							{
								ExprForm: &expressionv1.LogicalExprNode_Identifier{
									Identifier: &expressionv1.Identifier{
										Name: "phone",
									},
								},
							},
						},
						Kwargs: make(map[string]*expressionv1.LogicalExprNode),
					},
				},
			},
		},
		{
			name: "dataframe_filter_float",
			expr: DataFrame("transactions").Filter(Col("amount").Gt(Float(4.0))).Agg("sum"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
								GetAttribute: &expressionv1.ExprGetAttribute{
									Attribute: &expressionv1.Identifier{
										Name: "sum",
									},
									Parent: &expressionv1.LogicalExprNode{
										ExprForm: &expressionv1.LogicalExprNode_GetSubscript{
											GetSubscript: &expressionv1.ExprGetSubscript{
												Parent: &expressionv1.LogicalExprNode{
													ExprForm: &expressionv1.LogicalExprNode_Identifier{
														Identifier: &expressionv1.Identifier{
															Name: "transactions",
														},
													},
												},
												Subscript: []*expressionv1.LogicalExprNode{
													{
														ExprForm: &expressionv1.LogicalExprNode_Call{
															Call: &expressionv1.ExprCall{
																Func: &expressionv1.LogicalExprNode{
																	ExprForm: &expressionv1.LogicalExprNode_Identifier{
																		Identifier: &expressionv1.Identifier{
																			Name: ">",
																		},
																	},
																},
																Args: []*expressionv1.LogicalExprNode{
																	{
																		ExprForm: &expressionv1.LogicalExprNode_Identifier{
																			Identifier: &expressionv1.Identifier{
																				Name: "amount",
																			},
																		},
																	},
																	{
																		ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
																			LiteralValue: &expressionv1.ExprLiteral{
																				Value: &arrowv1.ScalarValue{
																					Value: &arrowv1.ScalarValue_Float64Value{
																						Float64Value: 4.0,
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "dataframe_filter",
			expr: DataFrame("transactions").Agg("sum"),
			expected: &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_Call{
					Call: &expressionv1.ExprCall{
						Func: &expressionv1.LogicalExprNode{
							ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
								GetAttribute: &expressionv1.ExprGetAttribute{
									Attribute: &expressionv1.Identifier{
										Name: "sum",
									},

									Parent: &expressionv1.LogicalExprNode{
										ExprForm: &expressionv1.LogicalExprNode_Identifier{
											Identifier: &expressionv1.Identifier{
												Name: "transactions",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Handle nil case specially
			if tt.expr == nil {
				actualProto := ToProto(nil)
				if actualProto != nil {
					t.Errorf("Expected nil, got %v", actualProto)
				}
				return
			}

			actualProto := ToProto(tt.expr)
			if !proto.Equal(actualProto, tt.expected) {
				t.Errorf("Proto mismatch for %s.\nExpected: %v\nActual: %v", tt.name, tt.expected, actualProto)
			}
		})
	}
}

func TestUnderscoreExprId(t *testing.T) {
	expr_ := Identifier("_")
	expr__ := Identifier("__")
	proto_1 := ToProto(expr_)
	proto_2 := ToProto(expr_)
	proto__ := ToProto(expr__)

	// expression ids must be distinct
	assert.NotEqual(t, proto_1.ExprId, proto_2.ExprId)
	// double underscore expressions are identical
	assert.Equal(t, proto__.ExprId, "")
}
