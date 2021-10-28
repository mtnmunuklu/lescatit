package classifiers

import (
	"Lescatit/categorizer/models"
	"bytes"

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
func NewNaiveBayesianClassifer() NaiveBayesianClassifier {
	return &NBClassifier{}
}

// Learn provides to create a new classifer model.
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
		return "", err
	}
	return buffer.String(), nil
}

// TODO: check it if it gives the correct result
func (c *NBClassifier) Predict(tokens []string) string {
	scores, _, _ := c.classifier.LogScores(tokens)
	results := models.Results{}
	for i := 0; i < len(scores); i++ {
		results = append(results, models.Result{ID: i, Score: scores[i]})
	}
	flags := []string{}
	for i := 0; i < len(results); i++ {
		flags = append(flags, string(c.classes[results[i].ID]))
	}
	return flags[0]
}

func (c *NBClassifier) ReadClassifier(data string) error {
	var buffer bytes.Buffer
	buffer.WriteString(data)
	classifier, err := bayesian.NewClassifierFromReader(&buffer)
	if err != nil {
		return err
	}
	c.classifier = classifier
	c.classes = c.classifier.Classes
	return nil
}
