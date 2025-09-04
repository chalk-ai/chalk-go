package expr

import (
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
)

// ToProto converts an ExprI to a LogicalExprNode proto message
func ToProto(expr ExprI) *expressionv1.LogicalExprNode {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *LiteralExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
				LiteralValue: &expressionv1.ExprLiteral{
					Value:               e.ScalarValue,
					IsArrowScalarObject: false,
				},
			},
		}

	case *ColumnExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Identifier{
				Identifier: &expressionv1.Identifier{
					Name: formatColumnName(e),
				},
			},
		}

	case *GetAttributeExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
				GetAttribute: &expressionv1.ExprGetAttribute{
					Parent: ToProto(e.Parent),
					Attribute: &expressionv1.Identifier{
						Name: e.Attribute,
					},
				},
			},
		}

	case *CallExpr:
		args := make([]*expressionv1.LogicalExprNode, len(e.Args))
		for i, arg := range e.Args {
			args[i] = ToProto(arg)
		}

		kwargs := make(map[string]*expressionv1.LogicalExprNode)
		for key, value := range e.Kwargs {
			kwargs[key] = ToProto(value)
		}

		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Call{
				Call: &expressionv1.ExprCall{
					Func: &expressionv1.LogicalExprNode{
						ExprForm: &expressionv1.LogicalExprNode_Identifier{
							Identifier: &expressionv1.Identifier{
								Name: e.Function,
							},
						},
					},
					Args:   args,
					Kwargs: kwargs,
				},
			},
		}

	case *AliasExpr:
		// Aliases are typically handled at a higher level in the query plan,
		// but we can represent them as the underlying expression for now
		return ToProto(e.Expression)

	case *dataFrameExprImpl:
		// DataFrame reference as identifier
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Identifier{
				Identifier: &expressionv1.Identifier{
					Name: e.Name,
				},
			},
		}

	case *aggregateExprImpl:
		// Represent aggregation as a function call
		args := []*expressionv1.LogicalExprNode{
			ToProto(e.DataFrame),
		}

		// No need to add filters separately since DataFrame.String() includes them
		return args[0]

		//ret := &expressionv1.LogicalExprNode{
		//	ExprForm: &expressionv1.LogicalExprNode_Call{
		//		Call: &expressionv1.ExprCall{
		//			Func: &expressionv1.ExprGetAttribute{
		//				Attribute: "count",
		//				Parent: &expressionv1.ExprCall{
		//					Func: args[0],
		//					//&expressionv1.LogicalExprNode{
		//					//ExprForm: .ExprForm,
		//					//ExprForm: &expressionv1.LogicalExprNode_Identifier{
		//					//	Identifier: &expressionv1.Identifier{
		//					//		Name: e.Function,
		//					//	},
		//					//},
		//					//},
		//					//Args: args,
		//				},
		//			},
		//		},
		//	},
		//}
		//return ret

	default:
		// Fallback for unknown expression types
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Identifier{
				Identifier: &expressionv1.Identifier{
					Name: "unknown",
				},
			},
		}
	}
}

// formatColumnName formats a column expression as a string
func formatColumnName(col *ColumnExpr) string {
	if col.Relation != "" {
		return col.Relation + "." + col.Name
	}
	return col.Name
}
