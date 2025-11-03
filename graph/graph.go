// Package graph exposes functionality to express Chalk definitions and modify the protograph.
package graph

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/chalk-ai/chalk-go/expr"
	arrowv1 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	expressionv1 "github.com/chalk-ai/chalk-go/gen/chalk/expression/v1"
	graphv1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
	graphv2 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v2"
	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Definitions stores either Feature Set or Stream Resolver definitions
type Definitions struct {
	FeatureSets     []*FeatureSet
	StreamResolvers []*StreamResolver
}

// WithFeatureSets is a variadic function that adds Feature Set definitions
// note it is a chaining method and not in-place
func (d Definitions) WithFeatureSets(fs ...FeatureSet) Definitions {
	for _, f := range fs {
		d.FeatureSets = append(d.FeatureSets, &f)
	}
	return d
}

// WithFeatureSets is a variadic function that adds Stream Resolver definitions
// note it is a chaining method and not in-place
func (d Definitions) WithStreamResolvers(srs ...StreamResolver) Definitions {
	for _, sr := range srs {
		d.StreamResolvers = append(d.StreamResolvers, &sr)
	}
	return d
}

// ToGraph converts a Definitions struct to a protograph
func (d Definitions) ToGraph() (*graphv1.Graph, error) {
	g := graphv1.Graph{}
	err := d.UpdateGraph(&g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

// UpdateGraph will modify protograph g in-place to include Feature Sets and Stream Resolvers from d,
// replacing any with the same name.
func (d Definitions) UpdateGraph(g *graphv1.Graph) error {
	type protoWithExtras struct {
		proto       *graphv1.FeatureSet
		foreignKeys map[string]string
		primaryName string
		primaryType *ScalarFeatureBuilder
		graphIndex  int
		fromUs      bool
		names       hset[string]
	}
	nameToFeatureSet := map[string]protoWithExtras{}

	// index feature sets in graph g
	for i, fs := range g.FeatureSets {
		primaryName := ""
		var primaryType *ScalarFeatureBuilder
		for _, f := range fs.Features {
			scalar := f.GetScalar()
			if scalar != nil && scalar.IsPrimary {
				primaryName = scalar.Name
				switch TypeName(scalar) {
				case "str":
					primaryType = String
				case "int":
					primaryType = Int
				default:
					panic("unreachable state")
				}
			}
		}

		nameToFeatureSet[fs.Name] = protoWithExtras{
			primaryName: primaryName,
			primaryType: primaryType,
			// assume we won't be redefining any feature sets
			// which are the target of a has_one in the target graph
			foreignKeys: nil,
			proto:       fs,
			graphIndex:  i,
			fromUs:      false,
		}
	}

	// index feature sets in definitions
	for _, fs := range d.FeatureSets {
		proto, names, err := fs.ToProto()
		if err != nil {
			return err
		}

		graphIndex := -1

		prev, withSameName := nameToFeatureSet[fs.namespace]
		if withSameName {
			if prev.foreignKeys == nil {
				graphIndex = prev.graphIndex
				log.Printf("   🔄 Replaced existing FeatureSet: %s", prev.proto.Name)
			} else {
				return fmt.Errorf("defined multiple FeatureSets with the same name: %s", fs.namespace)
			}
		}

		nameToFeatureSet[fs.namespace] = protoWithExtras{
			primaryName: fs.primaryName,
			primaryType: fs.primaryType,
			foreignKeys: fs.foreignKeys,
			proto:       proto,
			graphIndex:  graphIndex,
			fromUs:      true,
			names:       names,
		}
	}

	// resolve foreign references
	for namespace, extras := range nameToFeatureSet {
		if !extras.fromUs {
			continue
		}
		for _, f := range extras.proto.Features {
			switch t := f.Type.(type) {
			// create join conditions for any has_many without them
			case *graphv1.FeatureType_HasMany:
				hm := t.HasMany
				if hm.Join == nil {
					foreignExtras, valid := nameToFeatureSet[hm.ForeignNamespace]
					if !valid {
						return fmt.Errorf("dataframe %s referenced nonexistant feature set %s", hm.Name, hm.ForeignNamespace)
					}
					foreignCol := expr.ColIn(hm.ForeignNamespace, foreignExtras.foreignKeys[namespace])
					primaryCol := expr.ColIn(namespace, extras.primaryName)
					join, err := toFilterParsedProto(foreignCol.Eq(primaryCol), "")
					if err != nil {
						return err
					}
					hm.Join = join
				}

			// create groupby's using dataframe feature set's foreign key
			case *graphv1.FeatureType_Scalar:
				s := t.Scalar
				if s.WindowInfo != nil && s.WindowInfo.Aggregation != nil {
					foreignNamespace := s.WindowInfo.Aggregation.Namespace
					foreignExtras := nameToFeatureSet[foreignNamespace]
					foreignCol := foreignExtras.foreignKeys[namespace]
					s.WindowInfo.Aggregation.GroupBy = []*graphv1.FeatureReference{{
						Name:      foreignCol,
						Namespace: foreignNamespace,
					}}
				}
			}
		}

		// materialize foreign keys
		for featureSetName, foreignKeyName := range extras.foreignKeys {
			foreignExtras, exists := nameToFeatureSet[featureSetName]
			if !exists {
				return fmt.Errorf("invalid foreign key %s", foreignKeyName)
			}
			// only add if the feature doesn't already exist
			if !contains(foreignKeyName, extras.names) {
				extras.proto.Features = append(
					extras.proto.Features,
					foreignExtras.primaryType.ToProto(foreignKeyName, namespace),
				)
			}
		}
	}

	streamResolvers := make(map[string]*graphv1.StreamResolver, len(d.StreamResolvers))
	for _, sr := range d.StreamResolvers {
		extras, found := nameToFeatureSet[sr.OutputFeatureSet]
		if !found {
			return fmt.Errorf("feature set %s referenced by stream resolver %s not found", sr.OutputFeatureSet, sr.Name)
		}
		for k := range sr.OutputFeatures {
			if !contains(k, extras.names) {
				return fmt.Errorf("output feature %s in stream resolver %s are not present in the target feature set %s", k, sr.Name, sr.OutputFeatureSet)
			}
		}
		proto, err := sr.ToProto()
		if err != nil {
			return err
		}
		streamResolvers[proto.Fqn] = proto
	}

	log.Printf("Adding %d new FeatureSets", len(nameToFeatureSet)-len(g.FeatureSets))
	newFeatureSets := make([]*graphv1.FeatureSet, len(nameToFeatureSet))
	size := len(g.FeatureSets)
	for name, extras := range nameToFeatureSet {
		if extras.graphIndex == -1 {
			newFeatureSets[size] = extras.proto
			size++
			log.Printf("   ✅ Added new FeatureSet: %s", name)
		} else {
			newFeatureSets[extras.graphIndex] = extras.proto
		}
	}
	g.FeatureSets = newFeatureSets

	for i, sr := range g.StreamResolvers {
		newSr, sameName := streamResolvers[sr.Fqn]
		if sameName {
			g.StreamResolvers[i] = newSr
			delete(streamResolvers, newSr.Fqn)
			log.Printf("   🔄 Updated streaming resolver: %s", sr.Fqn)
		}
	}
	for _, sr := range streamResolvers {
		g.StreamResolvers = append(g.StreamResolvers, sr)
		log.Printf("   ✅ Added streaming resolver: %s", sr.Fqn)
	}

	return nil
}

type FeatureSet struct {
	// metadata fields
	Name               string
	Features           []*graphv1.FeatureType
	IsSingleton        bool
	Tags               []string
	Owner              string
	Doc                string
	EtlOfflineToOnline bool
	MaxStaleness       time.Duration

	// private fields
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

// FeatureName gets the field name of a feature, useful for debugging
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

// With adds a new feature to a Feature Set; roughly equivalent to one line of equivalent Python
func (fs FeatureSet) With(name string, ofType FeatureBuilder) FeatureSet {
	if fs.err != nil {
		return fs
	}
	if len(fs.namespace) == 0 {
		fs.namespace = strcase.ToSnake(fs.Name)
	}
	newFeatures, err := ofType.appendFeatures(fs.Features, name, fs.namespace)
	if err != nil {
		fs.err = err
	} else {
		fs.Features = newFeatures
	}
	return fs
}

type Features map[string]FeatureBuilder

// WithAll lets you define features in bulk
// For example the following line:
//
//	FeatureSet{name: "user"}.WithAll(Features{"name": graph.String, "age": graph.Int})
//
// is equivalent to two .With chained calls
func (fs FeatureSet) WithAll(m Features) FeatureSet {
	if fs.err != nil {
		return fs
	}
	if len(fs.namespace) == 0 {
		fs.namespace = strcase.ToSnake(fs.Name)
	}
	features := make([]*graphv1.FeatureType, len(fs.Features), len(fs.Features)+len(m))
	copy(features, fs.Features)
	// needed because windowed features reference has many features
	for name, ofType := range m {
		_, check := ofType.(*WindowedFeatureBuilder)
		if !check {
			features, fs.err = ofType.appendFeatures(features, name, fs.namespace)
			if fs.err != nil {
				return fs
			}
		}
	}
	for name, ofType := range m {
		_, check := ofType.(*WindowedFeatureBuilder)
		if check {
			features, fs.err = ofType.appendFeatures(features, name, fs.namespace)
			if fs.err != nil {
				return fs
			}
		}
	}
	fs.Features = features
	return fs
}

// WithPrimary defines a primary field, of type Int or String
// Note this must be called separately from WithAll,
// and if you do not manually define one, an implicit `id: Primary[int]` field will be added
func (fs FeatureSet) WithPrimary(name string, ofType FeatureBuilder) FeatureSet {
	if fs.err != nil {
		return fs
	}
	if fs.primaryType != nil {
		fs.err = fmt.Errorf("tried to add primary column %s to %s, when %s already exists", name, fs.Name, fs.primaryName)
		if fs.err != nil {
			return fs
		}
	}
	if len(fs.namespace) == 0 {
		fs.namespace = strcase.ToSnake(fs.Name)
	}
	scalarPtr, ok := ofType.(*ScalarFeatureBuilder)
	if !ok {
		fs.err = fmt.Errorf("primary column %s.%s must be scalar (int or str)", fs.Name, name)
		if fs.err != nil {
			return fs
		}
		return fs
	}
	newFeature := scalarPtr.ToProto(name, fs.namespace)
	newFeature.Type.(*graphv1.FeatureType_Scalar).Scalar.IsPrimary = true
	fs.Features = append(fs.Features, newFeature)
	fs.primaryName = name
	fs.primaryType = scalarPtr
	return fs
}

// WithForeignKey allows specifying fields which refer to primary fields of other Feature Sets
// For example
//
//	FeatureSet{name: "Transaction"}.WithForeignKey("user_id", User)
//
// is equivalent to the following chalkpy
//
//	@features
//	class Transaction:
//	  user_id: 'User'
func (fs FeatureSet) WithForeignKey(name, relation string) FeatureSet {
	if fs.err != nil {
		return fs
	}
	if fs.foreignKeys == nil {
		fs.foreignKeys = make(map[string]string)
	}
	fs.foreignKeys[strcase.ToSnake(relation)] = name
	return fs
}

func getFeatureSetNamed(graph *graphv1.Graph, fsName string) *graphv1.FeatureSet {
	for _, fs := range graph.FeatureSets {
		if fs.Name == fsName {
			return fs
		}
	}
	return nil
}

func CheckHasFeature(graph *graphv1.Graph, fsName, fName string) error {
	fs := getFeatureSetNamed(graph, fsName)
	if fs == nil {
		return fmt.Errorf("could not find feature set %s", fsName)
	}
	return checkHasFeatures(fs, fName)
}

func checkHasFeatures(fsProto *graphv1.FeatureSet, featureNames ...string) error {
	for _, fn := range featureNames {
		found := false
		for _, f2 := range fsProto.Features {
			if fn == FeatureName(f2) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("could not find %s in %s", fn, fsProto.Name)
		}
	}
	return nil
}

type HasOneFeatureBuilder struct {
	foreignNamespace string
	join             expr.Expr
}

// HasOne defines a new has_one feature
// For example
//
//	FeatureSet{name: "Transaction"}.With("user_id", graph.Int).With("user", HasOne("User", expr.Col("User", "id").Eq(expr.Col("Transaction", "user_id"))))
//
// is equivalent to the following chalkpy
//
//	@features
//	class Transaction:
//	  user_id: int
//	  user: has_one('User', lambda: User.id == Transaction.user_id)
func HasOne(relation string, join expr.Expr) *HasOneFeatureBuilder {
	return &HasOneFeatureBuilder{
		foreignNamespace: strcase.ToSnake(relation),
		join:             join,
	}
}

func (ho HasOneFeatureBuilder) appendFeatures(features []*graphv1.FeatureType, fieldName, namespace string) ([]*graphv1.FeatureType, error) {
	exproto, err := toFilterParsedProto(ho.join, "")
	if err != nil {
		return nil, err
	}
	return append(features, &graphv1.FeatureType{
		Type: &graphv1.FeatureType_HasOne{
			HasOne: &graphv1.HasOneFeatureType{
				ForeignNamespace:         ho.foreignNamespace,
				Name:                     fieldName,
				Namespace:                namespace,
				AttributeName:            fieldName,
				UnversionedAttributeName: fieldName,
				Join:                     exproto,
			},
		},
	}), nil
}

// zero size type to define hashset
type hset[T comparable] map[T]struct{}

func add[T comparable](k T, hs hset[T]) {
	hs[k] = struct{}{}
}
func contains[T comparable](k T, hs hset[T]) bool {
	_, ok := hs[k]
	return ok
}

func (fs *FeatureSet) ToProto() (*graphv1.FeatureSet, hset[string], error) {
	// propagate errors (in the future, list all?)
	if fs.err != nil {
		return nil, nil, fs.err
	}

	features := fs.Features

	names := make(hset[string], len(fs.Features))
	featureTimeName := ""
	for _, f := range features {
		name := FeatureName(f)
		if contains(name, names) {
			return nil, nil, fmt.Errorf("duplicate feature %s found in %s", name, fs.Name)
		} else {
			add(name, names)
		}
		ft := f.GetFeatureTime()
		if ft != nil {
			if featureTimeName == "" {
				featureTimeName = name
			} else {
				return nil, nil, fmt.Errorf("%s has more than one feature time features (%s and %s)", fs.Name, featureTimeName, name)
			}
		}
		ho := f.GetHasOne()
		if ho != nil {
			if fs.foreignKeys == nil {
				fs.foreignKeys = make(map[string]string)
			}
			fs.foreignKeys[ho.ForeignNamespace] = ho.Name
		}
	}

	// define autogenerated features
	if fs.primaryType == nil {
		fs.primaryName = "id"
		fs.primaryType = Int
	}

	if _, exists := names[fs.primaryName]; !exists {
		features = append(features, fs.primaryType.ToProto(fs.primaryName, fs.namespace))
		add("id", names)
	}
	if featureTimeName == "" {
		features = append(features, featureTime("__chalk_observed__", fs.namespace))
		add("__chalk_observed__", names)
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
					RichTypeInfo:  richType("int"),
				},
			},
		}
		features = append(features, chalkFkSingleton)
		add("__chalk_fk_singleton__", names)
	}

	return &graphv1.FeatureSet{
		Name:                 fs.namespace,
		Features:             features,
		IsSingleton:          fs.IsSingleton,
		Tags:                 fs.Tags,
		Owner:                maybeStr(fs.Owner),
		Doc:                  maybeStr(fs.Doc),
		MaxStalenessDuration: durationProto(fs.MaxStaleness),
	}, names, nil
}

type FeatureBuilder interface {
	// chained methods (return the same type)
	//WithMaxStaleness(time.Duration) FeatureBuilder

	appendFeatures([]*graphv1.FeatureType, string, string) ([]*graphv1.FeatureType, error)
}

type HasManyFeatureBuilder struct {
	ForeignNamespace string
	MaxStaleness     time.Duration
}

// DataFrame defines a new dataframe feature
// For example
//
//	FeatureSet{name: "User"}.With("transactions", DataFrame("Transactions")))
//
// is equivalent to the following chalkpy
//
//	@features
//	class User:
//	  transactions: DataFrame[Transactions]
func DataFrame(foreignName string) *HasManyFeatureBuilder {
	return &HasManyFeatureBuilder{
		ForeignNamespace: strcase.ToSnake(foreignName),
		MaxStaleness:     Years(10),
	}
}

func (hm *HasManyFeatureBuilder) appendFeatures(features []*graphv1.FeatureType, fieldName, namespace string) ([]*graphv1.FeatureType, error) {
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
func toFilterParsedProto(expression expr.ExprI, foreignNamespace string) (*expressionv1.LogicalExprNode, error) {
	if expression == nil {
		return nil, nil
	}

	switch e := expression.(type) {
	case *expr.ColumnExpr:
		// treat as an underscore expression
		relation := strcase.ToSnake(e.Relation)
		if relation == "" {
			relation = foreignNamespace
		}
		return toFilterParsedColumnProto(relation, e.Name), nil

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
			proto, err := toFilterParsedProto(e, foreignNamespace)
			if err != nil {
				return nil, err
			}
			args[i] = proto
		}

		return toFilterParsedBinaryProto(op, args), nil

	case *expr.GetAttributeExpr:
		if e.Attribute == "chalk_now" && e.Parent.String() == "_" {
			return toFilterParsedColumnProto("__chalk__", "now"), nil
		}

		if e.Attribute == "chalk_window" && e.Parent.String() == "_" {
			return expr.ToProto(expression)
		}
	}
	return nil, fmt.Errorf("invalid expression type for filter (%T): %s", expression, expression.String())
}

// StreamResolver represents a Native Streaming Resolver
// see https://docs.chalk.ai/docs/native-streaming
type StreamResolver struct {
	Name  string
	Owner string
	Doc   string

	// must match up to a UI-defined Data Source
	StreamSourceName string
	// "kafka" or "kinesis" or "pubsub"
	StreamSourceType string

	OutputFeatureSet string
	OutputFeatures   map[string]expr.Expr
	// defines the structure of message
	// should be a mapping from JSON key name to value type ("int" or "float" or "bool" or "str")
	MessageType map[string]string
	// if provided, only run on this machine type (defaults to all)
	MachineType string
	// if provided, only run in these environments (defaults to all)
	RunInEnvironments []string
	// if provided, evaluate this expression against incoming messages to either parse or filter
	Parse expr.Expr
}

func (sr StreamResolver) ToProto() (*graphv1.StreamResolver, error) {
	if sr.StreamSourceType != "kafka" &&
		sr.StreamSourceType != "kinesis" &&
		sr.StreamSourceType != "pubsub" {
		return nil, fmt.Errorf("[190] Invalid stream source: Invalid source for stream resolver '%s': expected KafkaSource, KinesisSource, or PubSubSource, got %s", sr.Name, sr.StreamSourceType)
	}

	// Validate name is a valid FQN (basic validation)
	if strings.TrimSpace(sr.Name) == "" {
		return nil, fmt.Errorf("[192] Empty resolver name: Stream resolver name cannot be empty")
	}
	if strings.Contains(sr.Name, ".") {
		return nil, fmt.Errorf("[193] Invalid resolver name format: Stream resolver name '%s' cannot contain dots. Use underscores instead", sr.Name)
	}

	messageSchema := arrowStruct(sr.MessageType)

	namespace := strcase.ToSnake(sr.OutputFeatureSet)
	featureExprs := make(map[string]*graphv1.FeatureExpression, len(sr.OutputFeatures))
	outputs := make([]*graphv1.ResolverOutput, len(sr.OutputFeatures))
	i := 0
	for k, v := range sr.OutputFeatures {
		exproto, err := expr.ToProto(v)
		if err != nil {
			return nil, err
		}
		featureExprs[fmt.Sprintf("%s.%s", namespace, k)] = &graphv1.FeatureExpression{
			Expr: &graphv1.FeatureExpression_UnderscoreExpr{
				UnderscoreExpr: exproto,
			},
		}
		outputs[i] = &graphv1.ResolverOutput{
			Annotation: &graphv1.ResolverOutput_Feature{
				Feature: &graphv1.FeatureReference{
					Name:      k,
					Namespace: namespace,
				},
			},
		}
		i++
	}

	var parseInfo *graphv1.ParseInfo
	if sr.Parse != nil {
		exproto, err := expr.ToProto(sr.Parse)
		if err != nil {
			return nil, err
		}
		parseInfo = &graphv1.ParseInfo{
			ParseFunction: &graphv1.FunctionReference{
				Name:               "_dummy_parse_fn",
				Module:             "chalk.features.resolver",
				FileName:           "chalk/chalk/features/resolver.py",
				FunctionDefinition: "    def _dummy_parse_fn(*args: Any, **kwargs: Any):\n        raise ValueError(\n            f\"Stream resolver '{name}' has expression-based parse function so it can't be called directly.\"\n        )\n",
			},
			ParseFunctionInputType: &arrowv1.ArrowType{
				ArrowTypeEnum: &arrowv1.ArrowType_LargeBinary{},
			},
			ParseFunctionOutputType:       messageSchema,
			IsParseFunctionOutputOptional: TRUE,
			ParseExpression: &graphv1.ParseInfo_UnderscoreExpr{
				UnderscoreExpr: exproto,
			},
		}
	}

	return &graphv1.StreamResolver{
		Fqn:                             strcase.ToSnake(sr.Name),
		Environments:                    sr.RunInEnvironments,
		Doc:                             maybeStr(sr.Doc),
		MachineType:                     maybeStr(sr.MachineType),
		Owner:                           maybeStr(sr.Owner),
		UpdatesMaterializedAggregations: true,
		ExplicitSchema:                  messageSchema,
		FeatureExpressions:              featureExprs,
		SourceV2: &graphv2.StreamSourceReference{
			Name:       sr.StreamSourceName,
			SourceType: sr.StreamSourceType,
		},
		Outputs: outputs,
		Params: []*graphv1.StreamResolverParam{
			{
				Type: &graphv1.StreamResolverParam_Message{
					Message: &graphv1.StreamResolverParamMessage{
						Name:      "message",
						ArrowType: messageSchema,
						StructType: &graphv1.StreamResolverParamMessage_Struct{
							Struct: &graphv1.FunctionGlobalCapturedStruct{
								Module:  "chalk-go",
								Name:    fmt.Sprintf("%s_message", sr.OutputFeatureSet),
								PaDtype: messageSchema,
							},
						},
					},
				},
			},
		},
		ParseInfo: parseInfo,
	}, nil
}
