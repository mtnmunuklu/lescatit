package classifiers

import (
	"Lescatit/categorizer/models"
	"Lescatit/categorizer/util"
	"bytes"
	"sort"

	"github.com/navossoc/bayesian"
)

// NaiveBayesianClassifier is the interface of the naive bayesian classifer.
type NaiveBayesianClassifier interface {
	Learn(model map[string][]string) (string, error)
	Predict(tokens []string) string
	ReadClassifier(filename string) error
}

// NBCategorizer provides categorizer and classes for naive bayesian classifier.
type NBClassifier struct {
	classifier *bayesian.Classifier
	classes    []bayesian.Class
}

// NewNaiveBayesianClassifer creates a new NaiveBayesianClassifier instance.
func NewNaiveBayesianClassifier() NaiveBayesianClassifier {
	return &NBClassifier{}
}

// Learn provides to create a new classifier model.
func (c *NBClassifier) Learn(model map[string][]string) (string, error) {
	c.classes = make([]bayesian.Class, 0)
	for class := range model {
		c.classes = append(c.classes, bayesian.Class(class))
	}

	c.classifier = bayesian.NewClassifierTfIdf(c.classes...)
	for class, tokens := range model {
		c.classifier.Learn(tokens, bayesian.Class(class))
	}

	c.classifier.ConvertTermsFreqToTfIdf()

	var buffer bytes.Buffer
	err := c.classifier.WriteTo(&buffer)
	if err != nil {
		return "", util.ErrSerializeClassifier
	}

	return buffer.String(), nil
}

// Predict provides to predict a class by naive bayesian classifier model.
func (c *NBClassifier) Predict(tokens []string) string {
	scores, _, _ := c.classifier.LogScores(tokens)
	results := models.Results{}
	for i := 0; i < len(scores); i++ {
		results = append(results, models.Result{ID: i, Score: scores[i]})
	}

	sort.Sort(sort.Reverse(results))

	flags := []string{}
	for i := 0; i < len(results); i++ {
		flags = append(flags, string(c.classes[results[i].ID]))
	}

	return flags[0]
}

// ReadClassifier provides to read naive bayesian classifier model.
func (c *NBClassifier) ReadClassifier(data string) error {
	var buffer bytes.Buffer
	buffer.WriteString(data)

	classifier, err := bayesian.NewClassifierFromReader(&buffer)
	if err != nil {
		return util.ErrDeserializeClassifer
	}

	c.classifier = classifier
	c.classes = c.classifier.Classes

	return nil
}
