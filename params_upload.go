package chalk

//
//// UploadFeaturesParamsComplete is the only type of object
//// accepted as an argument to Client.UploadFeatures.
//// UploadFeaturesParamsComplete is obtained by calling a chain
//// of methods starting with any method of [UploadFeaturesParam].
////
//// Example:
////
////				client.UploadFeatures(
////			 		WithInput(Features.User.Card.Id, []string{"card_1", "card_2"}).
////		         	WithInputs(map[any]any{
////			        	Features.User.Card.OwnerName: []string{"owner_1", "owner_2"},
////	                    Features.User.Card.OwnerEmail: []string{"owner_1@example", "owner_2@example"},
////		         	}).
////		         	WithEnvironment("pipkjlfc3gtmn"),
////				)
////
//// Specifying Inputs with [UploadFeaturesParams.WithInput] or [UploadFeaturesParams.WithInputs]
//// is mandatory. This means they must each be called at
//// least once for UploadFeatures to be returned.
//// Otherwise, an incomplete type will be returned, and it cannot
//// be passed into Client.UploadFeatures.
//type UploadFeaturesParamsComplete struct {
//	underlying UploadFeaturesParams
//}

type UploadFeaturesParams struct {
	Inputs              map[any]any
	BranchOverride      string
	EnvironmentOverride string
	PreviewDeploymentId string
}
