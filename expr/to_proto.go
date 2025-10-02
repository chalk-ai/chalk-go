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

func toProtos(exprs ...Expr) ([]*expressionv1.LogicalExprNode, error) {
	protos := make([]*expressionv1.LogicalExprNode, len(exprs))
	for i, e := range exprs {
		proto, err := ToProto(e)
		if err != nil {
			return nil, err
		}
		protos[i] = proto
	}
	return protos, nil
}

// ToProto converts an ExprI to a LogicalExprNode proto message
func ToProto(expr ExprI) (*expressionv1.LogicalExprNode, error) {
	if expr == nil {
		return nil, nil
	}

	switch e := expr.(type) {
	case *IdentifierExpr:
		return ToIdentifierLiteral(e.Name), nil

	case *LiteralExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_LiteralValue{
				LiteralValue: &expressionv1.ExprLiteral{
					Value:               e.ScalarValue,
					IsArrowScalarObject: false,
				},
			},
		}, nil

	case *ColumnExpr:
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Identifier{
				Identifier: &expressionv1.Identifier{
					Name: formatColumnName(e),
				},
			},
		}, nil

	case *GetAttributeExpr:
		parent, err := ToProto(e.Parent)
		if err != nil {
			return nil, err
		}
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_GetAttribute{
				GetAttribute: &expressionv1.ExprGetAttribute{
					Parent: parent,
					Attribute: &expressionv1.Identifier{
						Name: e.Attribute,
					},
				},
			},
		}, nil

	case *CallExpr:
		args, err := toProtos(e.Args...)
		if err != nil {
			return nil, err
		}

		kwargs := make(map[string]*expressionv1.LogicalExprNode)
		for key, value := range e.Kwargs {
			proto, err := ToProto(value)
			if err != nil {
				return nil, err
			}
			kwargs[key] = proto
		}

		fn, err := ToProto(e.Function)
		if err != nil {
			return nil, err
		}
		return &expressionv1.LogicalExprNode{
			ExprForm: &expressionv1.LogicalExprNode_Call{
				Call: &expressionv1.ExprCall{
					Func:   fn,
					Args:   args,
					Kwargs: kwargs,
				},
			},
		}, nil

	case *AliasExpr:
		// Aliases are typically handled at a higher level in the query plan,
		// but we can represent them as the underlying expression for now
		return ToProto(e.Expression)

	case *DataFrameExprImpl:
		// DataFrame reference as identifier
		// return ToProto(Identifier("_").Attr(e.Name))
		return ToIdentifierLiteral(e.Name), nil

	case *AggregateExprImpl:
		// Build the base DataFrame reference
		dataframeNode, err := ToProto(e.DataFrame)
		if err != nil {
			return nil, err
		}

		// If there are filter conditions or selections, wrap the DataFrame in a GetSubscript
		if len(e.Conditions) > 0 || e.Selection != nil {
			subscriptNodes := make([]*expressionv1.LogicalExprNode, 0, len(e.Conditions)+1)

			// Convert selection to proto node if it exists
			if e.Selection != nil {
				proto, err := ToProto(e.Selection)
				if err != nil {
					return nil, err
				}
				subscriptNodes = append(subscriptNodes, proto)
			}

			// Convert all conditions to proto nodes
			for _, condition := range e.Conditions {
				proto, err := ToProto(condition)
				if err != nil {
					return nil, err
				}
				subscriptNodes = append(subscriptNodes, proto)
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

		args := make([]*expressionv1.LogicalExprNode, 0, 1)
		kwargs := make(map[string]*expressionv1.LogicalExprNode, 1)
		if e.Function == "approx_top_k" || e.Function == "min_by_n" || e.Function == "max_by_n" {
			if len(e.Arguments) != 1 {
				return nil, fmt.Errorf("expecting exactly one argument 'k' to %s, got %d", e.Function, len(e.Arguments))
			}
			if e.Arguments[0] <= 0 {
				return nil, fmt.Errorf("argument 'k' to %s but be nonnegative, was %d", e.Function, e.Arguments[0])
			}
			k, _ := ToProto(Int(int64(e.Arguments[0])))
			args = append(args, k)
			kwargs["k"] = k
		} else if len(e.Arguments) > 0 {
			return nil, fmt.Errorf("aggregation %s must be called with zero arguments", e.Function)
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
					Args:   args,
					Kwargs: kwargs,
				},
			},
		}, nil

	default:
		return nil, fmt.Errorf("unknown (%T)", e)
	}
}

// formatColumnName formats a column expression as a string
func formatColumnName(col *ColumnExpr) string {
	if col.Relation != "" {
		return col.Relation + "." + col.Name
	}
	return col.Name
}
