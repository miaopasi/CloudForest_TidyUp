package CloudForest_TidyUp

/*
Main Structures

In CloudForest data is stored using the FeatureMatrix struct which contains Features.

The Feature struct  implements storage and methods for both categorical and numerical data and
calculations of impurity etc and the search for the best split.

The Target interface abstracts the methods of Feature that are needed for a feature to be predictable.
This allows for the implementation of alternative types of regression and classification.

Trees are built from Nodes and Splitters and stored within a Forest. Tree has a Grow
implements Brieman and Cutler's method (see extract above) for growing a tree. A GrowForest
method is also provided that implements the rest of the method including sampling cases
but it may be faster to grow the forest to disk as in the growforest utility.

Prediction and Voting is done using Tree.Vote and CatBallotBox and NumBallotBox which implement the
VoteTallyer interface.

*/

/*
ToDo: Figure Out What is DenseCatFeature as input.
ToDo: Figure Out How to call Adaboosttarget.


*/

/*
Analysis of Calling Function:

cattarget := fm.Data[targeti]
classtargets := GetAllClassificationTargets(cattarget.(*DenseCatFeature))

input is the CatFeature loaded from fm.Data[target].
AdaboostTarget receive CatFeature.Copy().(CatFeature) as input.
ToDo: Why need to copy while the others don't




NewAdaBoostTarget is defined in adaboosttarget.go
Input: CatFeature
Output: AdaboostTarget pointer




//NewAdaBoostTarget creates a categorical adaptive boosting target and initializes its weights.

func NewAdaBoostTarget(f CatFeature) (abt *AdaBoostTarget) {
	nCases := f.Length()
	abt = &AdaBoostTarget{f, make([]float64, nCases)}
	for i := range abt.Weights {
		abt.Weights[i] = 1 / float64(nCases)
	}
	return
}




func GetAllClassificationTargets(f CatFeature) []Target {
	costs := make(map[string]float64)
	for _, cat := range f.(*DenseCatFeature).Back {
		costs[cat] = 1.0
	}
	return []Target{f,
		NewEntropyTarget(f),
		NewAdaBoostTarget(f.Copy().(CatFeature)),
		NewWRFTarget(f, costs),
		//regret,
	}

}
 */