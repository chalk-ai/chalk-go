package expr

import (
	"fmt"

	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
)

func ToIdentifierLiteral(name string) *expressionv1.LogicalExprNode {
	return &expressionv1.LogicalExprNode{
		ExprForm: &expressionv1.LogicalExprNode_Identifier{
			Identifier: &expressionv1.Identifier{
				Name: name,
			},
		},
	}
}

// ToProto converts an ExprI to a LogicalExprNode proto message
func ToProto(expr ExprI) *expressionv1.LogicalExprNode {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *IdentifierExpr:
		return ToIdentifierLiteral(e.Name)

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
					Func:   ToProto(e.Function),
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
		return ToIdentifierLiteral(e.Name)

	case *aggregateExprImpl:
		// Build the base DataFrame reference
		dataframeNode := ToProto(e.DataFrame)

		// If there are filter conditions or selections, wrap the DataFrame in a GetSubscript
		if len(e.Conditions) > 0 || e.Selection != nil {
			subscriptNodes := make([]*expressionv1.LogicalExprNode, 0, len(e.Conditions)+1)

			// Convert selection to proto node if it exists
			if e.Selection != nil {
				subscriptNodes = append(subscriptNodes, ToProto(e.Selection))
			}

			// Convert all conditions to proto nodes
			for _, condition := range e.Conditions {
				subscriptNodes = append(subscriptNodes, ToProto(condition))
			}

			// Wrap the DataFrame in a GetSubscript with all filter conditions
			dataframeNode = &expressionv1.LogicalExprNode{
				ExprForm: &expressionv1.LogicalExprNode_GetSubscript{
					GetSubscript: &expressionv1.ExprGetSubscript{
						Parent:    dataframeNode,
						Subscript: subscriptNodes,
					},
				},
			}
		}

		// Apply the aggregation function as a GetAttribute on the (possibly filtered) DataFrame
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Call{
				Call: &expressionv1.ExprCall{
					Func: &expressionv1.LogicalExprNode{
						ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
							GetAttribute: &expressionv1.ExprGetAttribute{
								Attribute: &expressionv1.Identifier{
									Name: e.Function,
								},
								Parent: dataframeNode,
							},
						},
					},
				},
			},
		}

	default:
		// Fallback for unknown expression types
		return ToIdentifierLiteral(fmt.Sprintf("unknown (%T)", e))
	}
}

// formatColumnName formats a column expression as a string
func formatColumnName(col *ColumnExpr) string {
	if col.Relation != "" {
		return col.Relation + "." + col.Name
	}
	return col.Name
}
