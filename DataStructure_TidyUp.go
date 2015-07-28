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