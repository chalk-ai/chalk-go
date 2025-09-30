package graph

import (
	"fmt"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/known/durationpb"
)

type Definitions struct {
	FeatureSets []*FeatureSet
}

func (d Definitions) WithFeatureSets(fs ...FeatureSet) Definitions {
	for _, f := range fs {
		d.FeatureSets = append(d.FeatureSets, &f)
	}
	return d
}

func (d Definitions) ToGraph() (*graphv1.Graph, error) {
	type protoWithExtras struct {
		proto       *graphv1.FeatureSet
		foreignKeys map[string]string
		primaryName string
		primaryType *ScalarFeatureBuilder
	}
	g := graphv1.Graph{}
	nameToFeatureSet := map[string]protoWithExtras{}
	for _, fs := range d.FeatureSets {
		proto, err := fs.ToProto()
		if err != nil {
			return nil, err
		}

		g.FeatureSets = append(g.FeatureSets, proto)
		nameToFeatureSet[fs.namespace] = protoWithExtras{
			primaryName: fs.primaryName,
			primaryType: fs.primaryType,
			foreignKeys: fs.foreignKeys,
			proto:       proto,
		}
	}

	// resolve foreign references
	for _, fs := range g.FeatureSets {
		extras := nameToFeatureSet[fs.Name]
		// create join conditions for any has_many without them
		for _, f := range fs.Features {
			hm := f.GetHasMany()
			if hm != nil && hm.Join == nil {
				foreignExtras := nameToFeatureSet[hm.ForeignNamespace]
				foreignCol := expr.ColIn(hm.ForeignNamespace, foreignExtras.foreignKeys[fs.Name])
				primaryCol := expr.ColIn(fs.Name, extras.primaryName)
				join, err := toFilterParsedProto(foreignCol.Eq(primaryCol), "")
				if err != nil {
					return nil, err
				}
				hm.Join = join
			}
		}
		// materialize foreign keys
		for featureSetName, foreignKeyName := range extras.foreignKeys {
			foreignExtras, exists := nameToFeatureSet[featureSetName]
			if !exists || foreignExtras.primaryType == nil {
				continue
			}
			extras.proto.Features = append(
				extras.proto.Features,
				foreignExtras.primaryType.ToProto(foreignKeyName, fs.Name),
			)
		}
	}

	return &g, nil
}

type FeatureSet struct {
	Name               string
	Features           []*graphv1.FeatureType
	IsSingleton        bool
	Tags               []string
	Owner              string
	Doc                string
	EtlOfflineToOnline bool

	// snakecase version of Name
	namespace string
	// whether errors were encountered during construction
	err error
	// map of feature set -> foreign key
	foreignKeys map[string]string
	// primary column
	primaryName string
	primaryType *ScalarFeatureBuilder
}

func FeatureName(f *graphv1.FeatureType) string {
	switch t := f.Type.(type) {
	case *graphv1.FeatureType_FeatureTime:
		return t.FeatureTime.Name
	case *graphv1.FeatureType_Scalar:
		return t.Scalar.Name
	case *graphv1.FeatureType_HasMany:
		return t.HasMany.Name
	case *graphv1.FeatureType_GroupBy:
		return t.GroupBy.Name
	case *graphv1.FeatureType_Windowed:
		return t.Windowed.Name
	case *graphv1.FeatureType_HasOne:
		return t.HasOne.Name
	default:
		panic(fmt.Sprintf("unknown feature type %T", f.Type))
	}
}

func (fs FeatureSet) With(name string, ofType FeatureBuilder) FeatureSet {
	if len(fs.namespace) == 0 {
		fs.namespace = strcase.ToSnake(fs.Name)
	}
	newFeatures, err := ofType.AppendFeatures(fs.Features, name, fs.namespace)
	if err != nil {
		fs.err = err
	} else {
		fs.Features = newFeatures
	}
	return fs
}

func (fs FeatureSet) WithPrimary(name string, ofType FeatureBuilder) FeatureSet {
	if fs.primaryType != nil {
		fs.err = fmt.Errorf("tried to add primary column %s, when %s already exists", name, fs.primaryName)
	}
	if len(fs.namespace) == 0 {
		fs.namespace = strcase.ToSnake(fs.Name)
	}
	scalarPtr, ok := ofType.(*ScalarFeatureBuilder)
	if !ok {
		fs.err = fmt.Errorf("primary column %s must be scalar (int or str)", name)
		return fs
	}
	scalar := *scalarPtr
	newFeature := scalarPtr.ToProto(name, fs.namespace)
	newFeature.Type.(*graphv1.FeatureType_Scalar).Scalar.IsPrimary = true
	fs.Features = append(fs.Features, newFeature)
	fs.primaryName = name
	fs.primaryType = &scalar
	return fs
}

func (fs FeatureSet) WithForeignKey(name string, relation string) FeatureSet {
	if fs.foreignKeys == nil {
		fs.foreignKeys = make(map[string]string)
	}
	fs.foreignKeys[strcase.ToSnake(relation)] = name
	return fs
}

func MaybeStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// zero size type to define hashset
var Empty struct{}

func (fs *FeatureSet) ToProto() (*graphv1.FeatureSet, error) {
	// propagate errors (in the future, list all?)
	if fs.err != nil {
		return nil, fs.err
	}

	features := fs.Features

	names := make(map[string]struct{}, len(fs.Features))
	for i, f := range features {
		fmt.Printf("%d\n", i)
		name := FeatureName(f)
		_, ok := names[name]
		if ok {
			return nil, fmt.Errorf("duplicate feature %s found in %s", name, fs.Name)
		} else {
			names[name] = Empty
		}
	}
	// define autogenerated features
	if fs.primaryType == nil {
		fs.primaryName = "id"
		fs.primaryType = Int()
		features = append(features, fs.primaryType.ToProto(fs.primaryName, fs.namespace))
	}
	_, hasFeatureCalledTs := names["ts"]
	if !hasFeatureCalledTs {
		features = append(features, featureTime("__chalk_observed__", fs.namespace))
	}

	if !fs.IsSingleton {
		TRUE := true
		ONE := uint64(1)
		chalkFkSingleton := &graphv1.FeatureType{
			Type: &graphv1.FeatureType_Scalar{
				Scalar: &graphv1.ScalarFeatureType{
					Name:            "__chalk_fk_singleton__",
					Namespace:       fs.namespace,
					IsAutogenerated: true,
					AttributeName:   "__chalk_fk_singleton__",
					NoDisplay:       true,
					InternalVersion: &ONE,
					ArrowType: &arrowv1.ArrowType{
						ArrowTypeEnum: &arrowv1.ArrowType_Int8{},
					},
					DefaultValue: &arrowv1.ScalarValue{
						Value: &arrowv1.ScalarValue_Int8Value{
							Int8Value: 111,
						},
					},
					CacheStrategy: graphv1.CacheStrategy_CACHE_STRATEGY_ALL,
					StoreOnline:   &TRUE,
					StoreOffline:  &TRUE,
					RichTypeInfo:  Primitive("int"),
				},
			},
		}
		features = append(features, chalkFkSingleton)
	}

	return &graphv1.FeatureSet{
		Name:        fs.namespace,
		Features:    features,
		IsSingleton: fs.IsSingleton,
		Tags:        fs.Tags,
		Owner:       MaybeStr(fs.Owner),
		Doc:         MaybeStr(fs.Doc),
	}, nil
}

type FeatureBuilder interface {
	// chained methods (return the same type)
	//WithMaxStaleness(time.Duration) FeatureBuilder

	AppendFeatures([]*graphv1.FeatureType, string, string) ([]*graphv1.FeatureType, error)
}

type HasManyFeatureBuilder struct {
	ForeignNamespace string
	MaxStaleness     time.Duration
}

func DataFrame(foreignName string) *HasManyFeatureBuilder {
	return &HasManyFeatureBuilder{
		ForeignNamespace: strcase.ToSnake(foreignName),
		MaxStaleness:     Years(10),
	}
}

func (hm *HasManyFeatureBuilder) AppendFeatures(features []*graphv1.FeatureType, fieldName string, namespace string) ([]*graphv1.FeatureType, error) {
	return append(features,
		&graphv1.FeatureType{
			Type: &graphv1.FeatureType_HasMany{
				HasMany: &graphv1.HasManyFeatureType{
					ForeignNamespace:         hm.ForeignNamespace,
					Name:                     fieldName,
					Namespace:                namespace,
					AttributeName:            fieldName,
					UnversionedAttributeName: fieldName,
					MaxStalenessDuration:     durationpb.New(hm.MaxStaleness),
				},
			},
		},
	), nil
}

func (hm *HasManyFeatureBuilder) WithMaxStaleness(d time.Duration) *HasManyFeatureBuilder {
	hm.MaxStaleness = d
	return hm
}

func toFilterParsedColumnProto(relation string, name string) *expressionv1.LogicalExprNode {
	return &expressionv1.LogicalExprNode{
		ExprType: &expressionv1.LogicalExprNode_Column{
			Column: &expressionv1.Column{
				Name: name,
				Relation: &expressionv1.ColumnRelation{
					Relation: relation,
				},
			},
		},
	}
}

func toFilterParsedBinaryProto(op string, operands []*expressionv1.LogicalExprNode) *expressionv1.LogicalExprNode {
	return &expressionv1.LogicalExprNode{
		ExprType: &expressionv1.LogicalExprNode_BinaryExpr{
			BinaryExpr: &expressionv1.BinaryExprNode{
				Op:       op,
				Operands: operands,
			},
		},
	}
}

// ToProto converts an ExprI to a LogicalExprNode proto message using legacy FilterParsed operators
func toFilterParsedProto(expression expr.ExprI, namespace string) (*expressionv1.LogicalExprNode, error) {
	if expression == nil {
		return nil, nil
	}

	switch e := expression.(type) {
	case *expr.ColumnExpr:
		// treat as an underscore expression
		if e.Relation == "" {
			return toFilterParsedProto(expr.Identifier("_").Attr(e.Name), namespace)
		}
		return toFilterParsedColumnProto(e.Relation, e.Name), nil

	case *expr.GetAttributeExpr:
		if e.Attribute == "chalk_now" && e.Parent.String() == "_" {
			return toFilterParsedColumnProto("__chalk__", "now"), nil
		}

		return toFilterParsedBinaryProto("foreign_feature_access", []*expressionv1.LogicalExprNode{
			toFilterParsedColumnProto(namespace, e.Attribute),
		}), nil

	case *expr.LiteralExpr:
		return &expressionv1.LogicalExprNode{
			ExprType: &expressionv1.LogicalExprNode_Literal{
				Literal: e.ScalarValue,
			},
		}, nil

	case *expr.CallExpr:
		op := e.Function.String()
		switch op {
		case "=":
			op = "=="
		case "AND":
			op = "and"
		case "OR":
			op = "or"
		case "NOT":
			op = "not"
		// these have the same symbol
		case "!=":
		case "<":
		case "<=":
		case ">":
		case ">=":
		// any others are not supported
		default:
			return nil, fmt.Errorf("operator %s not allowed in filters", op)
		}

		args := make([]*expressionv1.LogicalExprNode, len(e.Args))
		for i, e := range e.Args {
			proto, err := toFilterParsedProto(e, namespace)
			if err != nil {
				return nil, err
			}
			args[i] = proto
		}

		return toFilterParsedBinaryProto(op, args), nil

	default:
		// Fallback for unknown expression types
		return nil, fmt.Errorf("invalid expression type for filter (%T)", e)
	}
}
