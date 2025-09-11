package graph

import (
	"fmt"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	"github.com/iancoleman/strcase"
	proto "google.golang.org/protobuf/proto"
)

// shorthand useful for defining underscore expressions
func __(name string) expr.Expr {
	return expr.Identifier("_").Attr(name)
}

type Definitions struct {
	FeatureSets []*FeatureSet
}

func (d Definitions) WithFeatureSets(fs ...FeatureSet) Definitions {
	for _, f := range fs {
		d.FeatureSets = append(d.FeatureSets, &f)
	}
	return d
}

func (d Definitions) ToGraph() *graphv1.Graph {
	g := graphv1.Graph{}
	for _, fs := range d.FeatureSets {
		g.FeatureSets = append(g.FeatureSets, fs.ToProto())
	}
	return &g
}

type FeatureSet struct {
	Name               string
	Features           []*graphv1.FeatureType
	IsSingleton        bool
	Tags               []string
	Owner              string
	Doc                string
	EtlOfflineToOnline bool
}

func (fs FeatureSet) WithFeature(name string, ofType FeatureBuilder) FeatureSet {
	fs.Features = append(fs.Features, ofType.ToProto(name, fs.Name))
	return fs
}

func maybeStr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func (fs *FeatureSet) ToProto() *graphv1.FeatureSet {
	return &graphv1.FeatureSet{
		Name:        strcase.ToSnake(fs.Name),
		Features:    fs.Features,
		IsSingleton: fs.IsSingleton,
		Tags:        fs.Tags,
		Owner:       maybeStr(fs.Owner),
		Doc:         maybeStr(fs.Doc),
	}
}

type FeatureBuilder interface {
	Expr(expr.Expr) FeatureBuilder

	ToProto(string, string) *graphv1.FeatureType
}

type ScalarFeatureBuilder struct {
	proto *graphv1.ScalarFeatureType
}

func (f ScalarFeatureBuilder) ToProto(fieldName string, className string) *graphv1.FeatureType {
	scalar := proto.Clone(f.proto).(*graphv1.ScalarFeatureType)
	scalar.Name = fieldName
	scalar.Namespace = strcase.ToSnake(className)
	scalar.AttributeName = fieldName
	scalar.UnversionedAttributeName = fieldName
	return &graphv1.FeatureType{
		Type: &graphv1.FeatureType_Scalar{
			Scalar: scalar,
		},
	}
}

func (ofType ScalarFeatureBuilder) Expr(expression expr.Expr) FeatureBuilder {
	scalar := proto.Clone(ofType.proto).(*graphv1.ScalarFeatureType)
	scalar.Expression = expr.ToProto(expression)
	return ScalarFeatureBuilder{
		proto: scalar,
	}
}

func Primitive(name string) *graphv1.FeatureRichTypeInfo {
	return &graphv1.FeatureRichTypeInfo{
		RichTypeIsSameAsPrimitiveType: true,
		RichType: &graphv1.FeatureRichType{
			Type: &graphv1.FeatureRichType_ClassType{
				ClassType: &graphv1.RichClassType{
					ModuleName: "builtins",
					Qualname:   name,
				},
			},
		},
		RichTypeName: maybeStr(fmt.Sprintf("<class '%s'>", name)),
	}
}

func Int() *ScalarFeatureBuilder {
	TRUE := true
	return &ScalarFeatureBuilder{
		proto: &graphv1.ScalarFeatureType{
			ArrowType: &arrowv1.ArrowType{
				ArrowTypeEnum: &arrowv1.ArrowType_Int64{},
			},
			CacheStrategy: graphv1.CacheStrategy_CACHE_STRATEGY_ALL,
			StoreOnline:   &TRUE,
			StoreOffline:  &TRUE,
			RichTypeInfo:  Primitive("int"),
		},
	}
}

type HasOneFeatureBuilder struct {
	proto *graphv1.HasOneFeatureType
}

type HasManyFeatureBuilder struct {
	proto *graphv1.HasManyFeatureType
}

type FeatureTimeFeatureBuilder struct {
	proto *graphv1.FeatureTimeFeatureType
}

type WindowedFeatureBuilder struct {
	proto *graphv1.WindowedFeatureType
}

type GroupByFeatureBuilder struct {
	proto *graphv1.GroupByFeatureType
}
