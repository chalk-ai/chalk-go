package graph

import (
	"fmt"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	proto "google.golang.org/protobuf/proto"
)

type ScalarFeatureBuilder struct {
	proto *graphv1.ScalarFeatureType
}

func (f ScalarFeatureBuilder) ToProto(fieldName string, namespace string) *graphv1.FeatureType {
	scalar := proto.Clone(f.proto).(*graphv1.ScalarFeatureType)
	scalar.Name = fieldName
	scalar.Namespace = namespace
	scalar.AttributeName = fieldName
	scalar.UnversionedAttributeName = fieldName
	return &graphv1.FeatureType{
		Type: &graphv1.FeatureType_Scalar{
			Scalar: scalar,
		},
	}
}

func (f ScalarFeatureBuilder) ToProtos(fieldName string, namespace string) ([]*graphv1.FeatureType, error) {
	return []*graphv1.FeatureType{
		f.ToProto(fieldName, namespace),
	}, nil
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
		RichTypeName: MaybeStr(fmt.Sprintf("<class '%s'>", name)),
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

func Float() *ScalarFeatureBuilder {
	TRUE := true
	return &ScalarFeatureBuilder{
		proto: &graphv1.ScalarFeatureType{
			ArrowType: &arrowv1.ArrowType{
				ArrowTypeEnum: &arrowv1.ArrowType_Float64{},
			},
			CacheStrategy: graphv1.CacheStrategy_CACHE_STRATEGY_ALL,
			StoreOnline:   &TRUE,
			StoreOffline:  &TRUE,
			RichTypeInfo:  Primitive("float"),
		},
	}
}

func String() *ScalarFeatureBuilder {
	TRUE := true
	return &ScalarFeatureBuilder{
		proto: &graphv1.ScalarFeatureType{
			ArrowType: &arrowv1.ArrowType{
				ArrowTypeEnum: &arrowv1.ArrowType_LargeUtf8{},
			},
			CacheStrategy: graphv1.CacheStrategy_CACHE_STRATEGY_ALL,
			StoreOnline:   &TRUE,
			StoreOffline:  &TRUE,
			RichTypeInfo:  Primitive("str"),
		},
	}
}

func Datetime() *ScalarFeatureBuilder {
	TRUE := true
	return &ScalarFeatureBuilder{
		proto: &graphv1.ScalarFeatureType{
			ArrowType: &arrowv1.ArrowType{
				ArrowTypeEnum: &arrowv1.ArrowType_Timestamp{},
			},
			CacheStrategy: graphv1.CacheStrategy_CACHE_STRATEGY_ALL,
			StoreOnline:   &TRUE,
			StoreOffline:  &TRUE,
			RichTypeInfo: &graphv1.FeatureRichTypeInfo{
				RichTypeIsSameAsPrimitiveType: true,
				RichType: &graphv1.FeatureRichType{
					Type: &graphv1.FeatureRichType_ClassType{
						ClassType: &graphv1.RichClassType{
							ModuleName: "datetime",
							Qualname:   "datetime",
						},
					},
				},
				RichTypeName: MaybeStr("<class 'datetime.datetime'>"),
			},
		},
	}
}
